package mouse

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"viiper/internal/server/api"
	"viiper/pkg/usb"
)

func init() {
	api.RegisterDevice("mouse", &handler{})
}

type handler struct{}

func (r *handler) CreateDevice() usb.Device { return New() }

func (r *handler) StreamHandler() api.StreamHandlerFunc {
	return func(conn net.Conn, devPtr *usb.Device, logger *slog.Logger) error {
		if devPtr == nil || *devPtr == nil {
			return fmt.Errorf("nil device")
		}
		mdev, ok := (*devPtr).(*Mouse)
		if !ok {
			return fmt.Errorf("device is not mouse")
		}

		buf := make([]byte, 5)
		for {
			if _, err := io.ReadFull(conn, buf); err != nil {
				if err == io.EOF {
					logger.Info("client disconnected")
					return nil
				}
				return fmt.Errorf("read input state: %w", err)
			}

			var state InputState
			if err := state.UnmarshalBinary(buf); err != nil {
				return fmt.Errorf("unmarshal input state: %w", err)
			}
			mdev.UpdateInputState(state)
		}
	}
}
