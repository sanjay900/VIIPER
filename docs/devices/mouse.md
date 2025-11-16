# HID Mouse (virtual)

A standard 5-button mouse with vertical and horizontal scroll wheels. Reports relative motion deltas and supports up to five buttons.

- USB IDs: VID 0x2E8A (Raspberry Pi), PID 0x0011
- Interface/Endpoint: IN 0x81 (mouse input report)
- Device type id (for API add): `mouse`

## HID report format (host-facing)

Input (device → host): 5 bytes

- Byte 0: Buttons bitfield (bits 0..4 for buttons 1..5)
- Byte 1: X delta (int8)
- Byte 2: Y delta (int8)
- Byte 3: Vertical wheel (int8; positive up)
- Byte 4: Horizontal wheel/pan (int8; positive right)

Deltas are consumed after each IN report so motion is truly relative and not repeated across host polls.

## Device stream protocol (client-facing)

Wire format from your client into VIIPER:

- Fixed 5-byte packets matching the HID report layout:
  [Buttons, dX, dY, Wheel, Pan]

Buttons persist until changed; motion/wheel deltas are applied once and reset.

## Adding the device

```text
bus/create
bus/1/add mouse
```

## Examples

A runnable example that periodically moves the mouse a short distance, clicks, and scrolls is provided in `examples/virtual_mouse/`.
