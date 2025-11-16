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
	"viiper/pkg/device/keyboard"
)

// Minimal example: create a keyboard device, type "Hello!" + Enter every 5 seconds, monitor LEDs.
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: virtual_keyboard <api_addr>")
		fmt.Println("Example: virtual_keyboard localhost:3242")
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

	addResp, err := api.DeviceAdd(busID, "keyboard")
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

	// LED monitoring goroutine (bidirectional channel)
	go func() {
		buf := make([]byte, 1)
		for {
			if _, err := io.ReadFull(conn, buf); err != nil {
				close(stop)
				return
			}
			ledByte := buf[0]
			fmt.Printf("→ LEDs: Num=%v Caps=%v Scroll=%v\n",
				ledByte&keyboard.LEDNumLock != 0,
				ledByte&keyboard.LEDCapsLock != 0,
				ledByte&keyboard.LEDScrollLock != 0)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Every 5s: type 'Hello!' + Enter. Press Ctrl+C to stop.")
	for {
		select {
		case <-ticker.C:
			// Type "Hello!" character by character
			states := keyboard.TypeString("Hello!")
			for _, state := range states {
				pkt, _ := state.MarshalBinary()
				if _, err := conn.Write(pkt); err != nil {
					fmt.Printf("Write error: %v\n", err)
					close(stop)
					return
				}
				time.Sleep(100 * time.Millisecond)
			}

			// Press and release Enter
			time.Sleep(100 * time.Millisecond)
			enterPress := keyboard.PressKey(keyboard.KeyEnter)
			pkt, _ := enterPress.MarshalBinary()
			if _, err := conn.Write(pkt); err != nil {
				fmt.Printf("Write error (enter): %v\n", err)
				close(stop)
				return
			}

			time.Sleep(100 * time.Millisecond)
			enterRelease := keyboard.Release()
			pkt, _ = enterRelease.MarshalBinary()
			if _, err := conn.Write(pkt); err != nil {
				fmt.Printf("Write error (release): %v\n", err)
				close(stop)
				return
			}

			fmt.Println("→ Typed: Hello!")
		case <-sigCh:
			fmt.Println("Signal received, stopping…")
			return
		case <-stop:
			return
		}
	}
}
