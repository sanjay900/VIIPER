package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"viiper/pkg/apiclient"
	"viiper/pkg/device/mouse"
)

// Minimal example: ensure a bus, create a mouse device, stream inputs, clean up on exit.
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: virtual_mouse <api_addr>")
		fmt.Println("Example: virtual_mouse localhost:3242")
		os.Exit(1)
	}

	addr := os.Args[1]
	api := apiclient.New(addr)

	busesResp, err := api.BusList()
	if err != nil {
		fmt.Printf("BusList error: %v\n", err)
		os.Exit(1)
	}
	var busID uint32
	createdBus := false
	if len(busesResp.Buses) == 0 {
		var createErr error
		for try := uint32(1); try <= 100; try++ {
			if r, err := api.BusCreate(try); err == nil {
				busID = r.BusID
				createdBus = true
				break
			}
			createErr = err
		}
		if busID == 0 {
			fmt.Printf("BusCreate failed: %v\n", createErr)
			os.Exit(1)
		}
		fmt.Printf("Created bus %d\n", busID)
	} else {
		busID = busesResp.Buses[0]
		for _, b := range busesResp.Buses[1:] {
			if b < busID {
				busID = b
			}
		}
		fmt.Printf("Using existing bus %d\n", busID)
	}

	addResp, err := api.DeviceAdd(busID, "mouse")
	if err != nil {
		fmt.Printf("DeviceAdd error: %v\n", err)
		if createdBus {
			_, _ = api.BusRemove(busID)
		}
		os.Exit(1)
	}
	deviceBusId := addResp.ID
	devId := deviceBusId
	if i := strings.Index(deviceBusId, "-"); i >= 0 && i+1 < len(deviceBusId) {
		devId = deviceBusId[i+1:]
	}
	createdDevice := true
	fmt.Printf("Created device %s on bus %d\n", devId, busID)

	defer func() {
		if createdDevice {
			if _, err := api.DeviceRemove(busID, devId); err != nil {
				fmt.Printf("DeviceRemove error: %v\n", err)
			} else {
				fmt.Printf("Removed device %s\n", deviceBusId)
			}
		}
		if createdBus {
			if _, err := api.BusRemove(busID); err != nil {
				fmt.Printf("BusRemove error: %v\n", err)
			} else {
				fmt.Printf("Removed bus %d\n", busID)
			}
		}
	}()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Stream dial error: %v\n", err)
		return
	}
	defer conn.Close()
	if _, err := fmt.Fprintf(conn, "bus/%d/%s\n", busID, devId); err != nil {
		fmt.Printf("Stream activate error: %v\n", err)
		return
	}
	fmt.Printf("Stream activated for %s\n", devId)

	stop := make(chan struct{})
	go func() {
		// Mouse has no feedback channel, but we monitor for connection close
		buf := make([]byte, 1)
		if _, err := io.ReadFull(conn, buf); err != nil {
			close(stop)
		}
	}()

	// Send a short movement once every 3 seconds for easy local testing.
	// Followed by a short click and a single scroll notch.
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Alternate direction to keep the pointer near its origin.
	dir := int8(1)
	const step = int8(50) // move diagonally by 50 px in X and Y
	fmt.Println("Every 3s: move diagonally by 50px (X and Y), then click and scroll. Press Ctrl+C to stop.")
	for {
		select {
		case <-ticker.C:
			// Move diagonally: (+step,+step) then (-step,-step) next tick
			dx := step * dir
			dy := step * dir
			dir *= -1

			// One-shot movement report (diagonal)
			move := mouse.InputState{DX: dx, DY: dy}
			pkt, _ := move.MarshalBinary()
			if _, err := conn.Write(pkt); err != nil {
				fmt.Printf("Write error (move): %v\n", err)
				close(stop)
				return
			}
			fmt.Printf("→ Moved mouse dx=%d dy=%d\n", dx, dy)

			// Zero state shortly after to keep movement one-shot (harmless safety)
			time.Sleep(30 * time.Millisecond)
			zero := mouse.InputState{}
			zpkt, _ := zero.MarshalBinary()
			if _, err := conn.Write(zpkt); err != nil {
				fmt.Printf("Write error (zero after move): %v\n", err)
				close(stop)
				return
			}

			// Simulate a short left click: press then release
			time.Sleep(50 * time.Millisecond)
			press := mouse.InputState{Buttons: 0x01}
			ppkt, _ := press.MarshalBinary()
			if _, err := conn.Write(ppkt); err != nil {
				fmt.Printf("Write error (press): %v\n", err)
				close(stop)
				return
			}
			time.Sleep(60 * time.Millisecond)
			rel := mouse.InputState{Buttons: 0x00}
			rpkt, _ := rel.MarshalBinary()
			if _, err := conn.Write(rpkt); err != nil {
				fmt.Printf("Write error (release): %v\n", err)
				close(stop)
				return
			}
			fmt.Printf("→ Clicked (left)\n")

			// Simulate a short scroll: one notch upwards
			time.Sleep(50 * time.Millisecond)
			scr := mouse.InputState{Wheel: 1}
			spkt, _ := scr.MarshalBinary()
			if _, err := conn.Write(spkt); err != nil {
				fmt.Printf("Write error (scroll): %v\n", err)
				close(stop)
				return
			}
			time.Sleep(30 * time.Millisecond)
			scr0 := mouse.InputState{}
			spkt0, _ := scr0.MarshalBinary()
			if _, err := conn.Write(spkt0); err != nil {
				fmt.Printf("Write error (zero after scroll): %v\n", err)
				close(stop)
				return
			}
			fmt.Printf("→ Scrolled (wheel=+1)\n")
		case <-sigCh:
			fmt.Println("Signal received, stopping…")
			return
		case <-stop:
			return
		}
	}
}
