package keyboard

// Modifier key bitmasks
const (
	ModLeftCtrl   = 0x01
	ModLeftShift  = 0x02
	ModLeftAlt    = 0x04
	ModLeftGUI    = 0x08 // Windows/Command key
	ModRightCtrl  = 0x10
	ModRightShift = 0x20
	ModRightAlt   = 0x40
	ModRightGUI   = 0x80
)

// LED bitmasks
const (
	LEDNumLock    = 0x01
	LEDCapsLock   = 0x02
	LEDScrollLock = 0x04
	LEDCompose    = 0x08
	LEDKana       = 0x10
)

// HID Usage codes for keyboard keys (USB HID Keyboard/Keypad usage page)
const (
	// Letters A-Z
	KeyA = 0x04
	KeyB = 0x05
	KeyC = 0x06
	KeyD = 0x07
	KeyE = 0x08
	KeyF = 0x09
	KeyG = 0x0A
	KeyH = 0x0B
	KeyI = 0x0C
	KeyJ = 0x0D
	KeyK = 0x0E
	KeyL = 0x0F
	KeyM = 0x10
	KeyN = 0x11
	KeyO = 0x12
	KeyP = 0x13
	KeyQ = 0x14
	KeyR = 0x15
	KeyS = 0x16
	KeyT = 0x17
	KeyU = 0x18
	KeyV = 0x19
	KeyW = 0x1A
	KeyX = 0x1B
	KeyY = 0x1C
	KeyZ = 0x1D

	// Numbers 1-0 (top row)
	Key1 = 0x1E
	Key2 = 0x1F
	Key3 = 0x20
	Key4 = 0x21
	Key5 = 0x22
	Key6 = 0x23
	Key7 = 0x24
	Key8 = 0x25
	Key9 = 0x26
	Key0 = 0x27

	// Special keys
	KeyEnter      = 0x28
	KeyEscape     = 0x29
	KeyBackspace  = 0x2A
	KeyTab        = 0x2B
	KeySpace      = 0x2C
	KeyMinus      = 0x2D // - and _
	KeyEqual      = 0x2E // = and +
	KeyLeftBrace  = 0x2F // [ and {
	KeyRightBrace = 0x30 // ] and }
	KeyBackslash  = 0x31 // \ and |
	KeyNonUSHash  = 0x32 // Non-US # and ~
	KeySemicolon  = 0x33 // ; and :
	KeyApostrophe = 0x34 // ' and "
	KeyGrave      = 0x35 // ` and ~
	KeyComma      = 0x36 // , and <
	KeyPeriod     = 0x37 // . and >
	KeySlash      = 0x38 // / and ?
	KeyCapsLock   = 0x39

	// Function keys
	KeyF1  = 0x3A
	KeyF2  = 0x3B
	KeyF3  = 0x3C
	KeyF4  = 0x3D
	KeyF5  = 0x3E
	KeyF6  = 0x3F
	KeyF7  = 0x40
	KeyF8  = 0x41
	KeyF9  = 0x42
	KeyF10 = 0x43
	KeyF11 = 0x44
	KeyF12 = 0x45

	// Control keys
	KeyPrintScreen = 0x46
	KeyScrollLock  = 0x47
	KeyPause       = 0x48
	KeyInsert      = 0x49
	KeyHome        = 0x4A
	KeyPageUp      = 0x4B
	KeyDelete      = 0x4C
	KeyEnd         = 0x4D
	KeyPageDown    = 0x4E

	// Arrow keys
	KeyRight = 0x4F
	KeyLeft  = 0x50
	KeyDown  = 0x51
	KeyUp    = 0x52

	// Numpad
	KeyNumLock    = 0x53
	KeyKpSlash    = 0x54 // Keypad /
	KeyKpAsterisk = 0x55 // Keypad *
	KeyKpMinus    = 0x56 // Keypad -
	KeyKpPlus     = 0x57 // Keypad +
	KeyKpEnter    = 0x58 // Keypad Enter
	KeyKp1        = 0x59 // Keypad 1 and End
	KeyKp2        = 0x5A // Keypad 2 and Down
	KeyKp3        = 0x5B // Keypad 3 and PageDn
	KeyKp4        = 0x5C // Keypad 4 and Left
	KeyKp5        = 0x5D // Keypad 5
	KeyKp6        = 0x5E // Keypad 6 and Right
	KeyKp7        = 0x5F // Keypad 7 and Home
	KeyKp8        = 0x60 // Keypad 8 and Up
	KeyKp9        = 0x61 // Keypad 9 and PageUp
	KeyKp0        = 0x62 // Keypad 0 and Insert
	KeyKpDot      = 0x63 // Keypad . and Delete

	// Additional keys
	KeyNonUSBackslash = 0x64 // Non-US \ and |
	KeyApplication    = 0x65 // Application (Windows Menu key)
	KeyPower          = 0x66 // Power (not commonly used)
	KeyKpEqual        = 0x67 // Keypad =

	// Extended function keys
	KeyF13 = 0x68
	KeyF14 = 0x69
	KeyF15 = 0x6A
	KeyF16 = 0x6B
	KeyF17 = 0x6C
	KeyF18 = 0x6D
	KeyF19 = 0x6E
	KeyF20 = 0x6F
	KeyF21 = 0x70
	KeyF22 = 0x71
	KeyF23 = 0x72
	KeyF24 = 0x73

	// Execution keys
	KeyExecute    = 0x74
	KeyHelp       = 0x75
	KeyMenu       = 0x76
	KeySelect     = 0x77
	KeyStop       = 0x78
	KeyAgain      = 0x79 // Redo
	KeyUndo       = 0x7A
	KeyCut        = 0x7B
	KeyCopy       = 0x7C
	KeyPaste      = 0x7D
	KeyFind       = 0x7E
	KeyMute       = 0x7F
	KeyVolumeUp   = 0x80
	KeyVolumeDown = 0x81

	// Media control keys
	KeyMediaPlayPause = 0xE8 // Play/Pause
	KeyMediaStop      = 0xE9 // Stop
	KeyMediaNext      = 0xEB // Next Track
	KeyMediaPrevious  = 0xEC // Previous Track
)
