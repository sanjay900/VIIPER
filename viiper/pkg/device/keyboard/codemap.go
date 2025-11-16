package keyboard

// KeyName maps HID usage codes to human-readable key names.
var KeyName = map[uint8]string{
	// Letters
	KeyA: "A", KeyB: "B", KeyC: "C", KeyD: "D", KeyE: "E", KeyF: "F", KeyG: "G",
	KeyH: "H", KeyI: "I", KeyJ: "J", KeyK: "K", KeyL: "L", KeyM: "M", KeyN: "N",
	KeyO: "O", KeyP: "P", KeyQ: "Q", KeyR: "R", KeyS: "S", KeyT: "T", KeyU: "U",
	KeyV: "V", KeyW: "W", KeyX: "X", KeyY: "Y", KeyZ: "Z",

	// Numbers
	Key1: "1", Key2: "2", Key3: "3", Key4: "4", Key5: "5",
	Key6: "6", Key7: "7", Key8: "8", Key9: "9", Key0: "0",

	// Special keys
	KeyEnter:      "Enter",
	KeyEscape:     "Escape",
	KeyBackspace:  "Backspace",
	KeyTab:        "Tab",
	KeySpace:      "Space",
	KeyMinus:      "Minus",
	KeyEqual:      "Equal",
	KeyLeftBrace:  "LeftBrace",
	KeyRightBrace: "RightBrace",
	KeyBackslash:  "Backslash",
	KeySemicolon:  "Semicolon",
	KeyApostrophe: "Apostrophe",
	KeyGrave:      "Grave",
	KeyComma:      "Comma",
	KeyPeriod:     "Period",
	KeySlash:      "Slash",
	KeyCapsLock:   "CapsLock",

	// Function keys
	KeyF1: "F1", KeyF2: "F2", KeyF3: "F3", KeyF4: "F4", KeyF5: "F5", KeyF6: "F6",
	KeyF7: "F7", KeyF8: "F8", KeyF9: "F9", KeyF10: "F10", KeyF11: "F11", KeyF12: "F12",
	KeyF13: "F13", KeyF14: "F14", KeyF15: "F15", KeyF16: "F16", KeyF17: "F17", KeyF18: "F18",
	KeyF19: "F19", KeyF20: "F20", KeyF21: "F21", KeyF22: "F22", KeyF23: "F23", KeyF24: "F24",

	// Control keys
	KeyPrintScreen: "PrintScreen",
	KeyScrollLock:  "ScrollLock",
	KeyPause:       "Pause",
	KeyInsert:      "Insert",
	KeyHome:        "Home",
	KeyPageUp:      "PageUp",
	KeyDelete:      "Delete",
	KeyEnd:         "End",
	KeyPageDown:    "PageDown",

	// Arrow keys
	KeyRight: "Right",
	KeyLeft:  "Left",
	KeyDown:  "Down",
	KeyUp:    "Up",

	// Numpad
	KeyNumLock:    "NumLock",
	KeyKpSlash:    "Kp/",
	KeyKpAsterisk: "Kp*",
	KeyKpMinus:    "Kp-",
	KeyKpPlus:     "Kp+",
	KeyKpEnter:    "KpEnter",
	KeyKp1:        "Kp1",
	KeyKp2:        "Kp2",
	KeyKp3:        "Kp3",
	KeyKp4:        "Kp4",
	KeyKp5:        "Kp5",
	KeyKp6:        "Kp6",
	KeyKp7:        "Kp7",
	KeyKp8:        "Kp8",
	KeyKp9:        "Kp9",
	KeyKp0:        "Kp0",
	KeyKpDot:      "Kp.",

	// Additional
	KeyApplication: "Application",
	KeyMute:        "Mute",
	KeyVolumeUp:    "VolumeUp",
	KeyVolumeDown:  "VolumeDown",

	// Media control
	KeyMediaPlayPause: "MediaPlayPause",
	KeyMediaStop:      "MediaStop",
	KeyMediaNext:      "MediaNext",
	KeyMediaPrevious:  "MediaPrevious",
}

// CharToKey maps ASCII characters to their corresponding HID usage codes.
// For shifted characters (uppercase, symbols), use with NeedsShift().
var CharToKey = map[byte]uint8{
	// Lowercase letters
	'a': KeyA, 'b': KeyB, 'c': KeyC, 'd': KeyD, 'e': KeyE, 'f': KeyF, 'g': KeyG,
	'h': KeyH, 'i': KeyI, 'j': KeyJ, 'k': KeyK, 'l': KeyL, 'm': KeyM, 'n': KeyN,
	'o': KeyO, 'p': KeyP, 'q': KeyQ, 'r': KeyR, 's': KeyS, 't': KeyT, 'u': KeyU,
	'v': KeyV, 'w': KeyW, 'x': KeyX, 'y': KeyY, 'z': KeyZ,

	// Uppercase letters (same keys, need shift)
	'A': KeyA, 'B': KeyB, 'C': KeyC, 'D': KeyD, 'E': KeyE, 'F': KeyF, 'G': KeyG,
	'H': KeyH, 'I': KeyI, 'J': KeyJ, 'K': KeyK, 'L': KeyL, 'M': KeyM, 'N': KeyN,
	'O': KeyO, 'P': KeyP, 'Q': KeyQ, 'R': KeyR, 'S': KeyS, 'T': KeyT, 'U': KeyU,
	'V': KeyV, 'W': KeyW, 'X': KeyX, 'Y': KeyY, 'Z': KeyZ,

	// Numbers (top row)
	'1': Key1, '2': Key2, '3': Key3, '4': Key4, '5': Key5,
	'6': Key6, '7': Key7, '8': Key8, '9': Key9, '0': Key0,

	// Shifted number row symbols
	'!': Key1, '@': Key2, '#': Key3, '$': Key4, '%': Key5,
	'^': Key6, '&': Key7, '*': Key8, '(': Key9, ')': Key0,

	// Unshifted symbols
	'-':  KeyMinus,
	'=':  KeyEqual,
	'[':  KeyLeftBrace,
	']':  KeyRightBrace,
	'\\': KeyBackslash,
	';':  KeySemicolon,
	'\'': KeyApostrophe,
	'`':  KeyGrave,
	',':  KeyComma,
	'.':  KeyPeriod,
	'/':  KeySlash,

	// Shifted symbols
	'_': KeyMinus,
	'+': KeyEqual,
	'{': KeyLeftBrace,
	'}': KeyRightBrace,
	'|': KeyBackslash,
	':': KeySemicolon,
	'"': KeyApostrophe,
	'~': KeyGrave,
	'<': KeyComma,
	'>': KeyPeriod,
	'?': KeySlash,

	// Whitespace
	' ':  KeySpace,
	'\n': KeyEnter,
	'\r': KeyEnter,
	'\t': KeyTab,
}

// ShiftChars defines which characters require the Shift modifier.
var ShiftChars = map[byte]bool{
	// Uppercase letters
	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true,
	'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true,
	'O': true, 'P': true, 'Q': true, 'R': true, 'S': true, 'T': true, 'U': true,
	'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,

	// Shifted number row
	'!': true, '@': true, '#': true, '$': true, '%': true,
	'^': true, '&': true, '*': true, '(': true, ')': true,

	// Shifted symbols
	'_': true, '+': true, '{': true, '}': true, '|': true,
	':': true, '"': true, '~': true, '<': true, '>': true, '?': true,
}
