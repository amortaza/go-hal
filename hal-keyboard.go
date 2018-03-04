package hal

type KeyboardKey int
type KeyBehaviorType int

const (
	Key_Behavior_HOME KeyBehaviorType = 1 + iota
	Key_Behavior_INSERT
	Key_Behavior_PAGE_UP
	Key_Behavior_PAGE_DOWN
	Key_Behavior_DOWN
	Key_Behavior_UP
	Key_Behavior_RIGHT
	Key_Behavior_LEFT
	Key_Behavior_NIL
	Key_Behavior_END
	Key_Behavior_DELETE
	Key_Behavior_ENTER
	Key_Behavior_BACKSPACE
	Key_Behavior_TAB
	Key_Behavior_ESCAPE
	Key_Behavior_UNKNOWN
	Key_Behavior_CHAR
	Key_Behavior_FUNCTION_KEY
)

const (
	Key_APOSTROPHE KeyboardKey = 1 + iota
	Key_SPACE
	Key_COMMA
	Key_MINUS
	Key_PERIOD
	Key_FORWARD_SLASH
	Key_SEMICOLON
	Key_EQUAL
	Key_LEFT_BRACKET
	Key_BACK_SLASH
	Key_RIGHT_BRACKET
	Key_ESCAPE
	Key_ENTER
	Key_TAB
	Key_BACKSPACE
	Key_INSERT
	Key_DELETE

	Key_RIGHT
	Key_LEFT
	Key_UP
	Key_DOWN
	Key_PAGE_UP
	Key_PAGE_DOWN
	Key_HOME
	Key_END

	Key_F1
	Key_F2
	Key_F3
	Key_F4
	Key_F5
	Key_F6
	Key_F7
	Key_F8
	Key_F9
	Key_F10
	Key_F11
	Key_F12

	Key_PAD_0
	Key_PAD_1
	Key_PAD_2
	Key_PAD_3
	Key_PAD_4
	Key_PAD_5
	Key_PAD_6
	Key_PAD_7
	Key_PAD_8
	Key_PAD_9

	Key_PAD_DECIMAL
	Key_PAD_DIVIDE
	Key_PAD_MULTIPLY
	Key_PAD_SUBTRACT
	Key_PAD_ADD
	Key_PAD_ENTER
	Key_PAD_EQUAL

	Key_1
	Key_2
	Key_3
	Key_4
	Key_5
	Key_6
	Key_7
	Key_8
	Key_9
	Key_0

	Key_A
	Key_B
	Key_C
	Key_D
	Key_E
	Key_F
	Key_G
	Key_H
	Key_I
	Key_J
	Key_K
	Key_L
	Key_M
	Key_N
	Key_O
	Key_P
	Key_Q
	Key_R
	Key_S
	Key_T
	Key_U
	Key_V
	Key_W
	Key_X
	Key_Y
	Key_Z
)

func KeyToChar(key KeyboardKey, shift, numlock bool) string {

	if key == Key_SPACE {
		return " "
	}

	if key == Key_A {
		if shift {
			return "A"
		} else {
			return "a"
		}
	}

	if key == Key_B {
		if shift {
			return "B"
		} else {
			return "b"
		}
	}

	if key == Key_C {
		if shift {
			return "C"
		} else {
			return "c"
		}
	}

	if key == Key_D {
		if shift {
			return "D"
		} else {
			return "d"
		}
	}

	if key == Key_E {
		if shift {
			return "E"
		} else {
			return "e"
		}
	}

	if key == Key_F {
		if shift {
			return "F"
		} else {
			return "f"
		}
	}

	if key == Key_G {
		if shift {
			return "G"
		} else {
			return "g"
		}
	}

	if key == Key_H {
		if shift {
			return "H"
		} else {
			return "h"
		}
	}

	if key == Key_I {
		if shift {
			return "I"
		} else {
			return "i"
		}
	}

	if key == Key_J {
		if shift {
			return "J"
		} else {
			return "j"
		}
	}

	if key == Key_K {
		if shift {
			return "K"
		} else {
			return "k"
		}
	}

	if key == Key_L {
		if shift {
			return "L"
		} else {
			return "l"
		}
	}

	if key == Key_M {
		if shift {
			return "M"
		} else {
			return "m"
		}
	}

	if key == Key_N {
		if shift {
			return "N"
		} else {
			return "n"
		}
	}

	if key == Key_O {
		if shift {
			return "O"
		} else {
			return "o"
		}
	}

	if key == Key_P {
		if shift {
			return "P"
		} else {
			return "p"
		}
	}

	if key == Key_Q {
		if shift {
			return "Q"
		} else {
			return "q"
		}
	}

	if key == Key_R {
		if shift {
			return "R"
		} else {
			return "r"
		}
	}

	if key == Key_S {
		if shift {
			return "S"
		} else {
			return "s"
		}
	}

	if key == Key_T {
		if shift {
			return "T"
		} else {
			return "t"
		}
	}

	if key == Key_U {
		if shift {
			return "U"
		} else {
			return "u"
		}
	}

	if key == Key_V {
		if shift {
			return "V"
		} else {
			return "v"
		}
	}

	if key == Key_W {
		if shift {
			return "W"
		} else {
			return "w"
		}
	}

	if key == Key_X {
		if shift {
			return "X"
		} else {
			return "x"
		}
	}

	if key == Key_Y {
		if shift {
			return "Y"
		} else {
			return "y"
		}
	}

	if key == Key_Z {
		if shift {
			return "Z"
		} else {
			return "z"
		}
	}

	if key == Key_RIGHT_BRACKET {
		if shift {
			return "}"
		} else {
			return "]"
		}
	}

	if key == Key_LEFT_BRACKET {
		if shift {
			return "{"
		} else {
			return "["
		}
	}

	if key == Key_BACK_SLASH {
		if shift {
			return "|"
		} else {
			return "\\"
		}
	}

	if key == Key_EQUAL {
		if shift {
			return "+"
		} else {
			return "="
		}
	}

	if key == Key_FORWARD_SLASH {
		if shift {
			return "?"
		} else {
			return "/"
		}
	}

	if key == Key_SEMICOLON {
		if shift {
			return ":"
		} else {
			return ";"
		}
	}

	if key == Key_PERIOD {
		if shift {
			return ">"
		} else {
			return "."
		}
	}

	if key == Key_MINUS {
		if shift {
			return "_"
		} else {
			return "-"
		}
	}

	if key == Key_COMMA {
		if shift {
			return "<"
		} else {
			return ","
		}
	}

	if key == Key_APOSTROPHE {
		if shift {
			return "~"
		} else {
			return "`"
		}
	}

	if key == Key_0 {
		if shift {
			return ")"
		} else {
			return "0"
		}
	}

	if key == Key_1 {
		if shift {
			return "!"
		} else {
			return "1"
		}
	}

	if key == Key_2 {
		if shift {
			return "@"
		} else {
			return "2"
		}
	}

	if key == Key_3 {
		if shift {
			return "#"
		} else {
			return "3"
		}
	}

	if key == Key_4 {
		if shift {
			return "$"
		} else {
			return "4"
		}
	}

	if key == Key_5 {
		if shift {
			return "%"
		} else {
			return "5"
		}
	}

	if key == Key_6 {
		if shift {
			return "^"
		} else {
			return "6"
		}
	}

	if key == Key_7 {
		if shift {
			return "&"
		} else {
			return "7"
		}
	}

	if key == Key_8 {
		if shift {
			return "*"
		} else {
			return "8"
		}
	}

	if key == Key_9 {
		if shift {
			return "("
		} else {
			return "9"
		}
	}

	if key == Key_PAD_1 {

		if numlock {
			return "1"
		}
	}

	if key == Key_PAD_2 {

		if numlock {
			return "2"
		}
	}

	if key == Key_PAD_3 {

		if numlock {
			return "3"
		}
	}

	if key == Key_PAD_4 {

		if numlock {
			return "4"
		}
	}

	if key == Key_PAD_5 {

		if numlock {
			return "5"
		}
	}

	if key == Key_PAD_6 {

		if numlock {
			return "6"
		}
	}

	if key == Key_PAD_7 {

		if numlock {
			return "7"
		}
	}

	if key == Key_PAD_8 {

		if numlock {
			return "8"
		}
	}

	if key == Key_PAD_9 {

		if numlock {
			return "9"
		}
	}

	if key == Key_PAD_0 {

		if numlock {
			return "0"
		}
	}

	if key == Key_PAD_DECIMAL {

		if numlock {
			return "."
		}
	}

	if key == Key_PAD_DIVIDE {
		return "/"
	}

	if key == Key_PAD_MULTIPLY {
		return "*"
	}

	if key == Key_PAD_SUBTRACT {
		return "-"
	}

	if key == Key_PAD_ADD {
		return "+"
	}

	if key == Key_PAD_EQUAL {
		return "="
	}

	return ""
}

func KeyToBehavior(key KeyboardKey, shift, numlock bool) KeyBehaviorType {

	if KeyToChar(key, shift, numlock) != "" {
		return Key_Behavior_CHAR
	}

	if key == Key_PAD_1 {

		if !numlock {
			return Key_Behavior_END
		}
	}

	if key == Key_PAD_2 {

		if !numlock {
			return Key_Behavior_DOWN
		}
	}

	if key == Key_PAD_3 {

		if !numlock {
			return Key_Behavior_PAGE_DOWN
		}
	}

	if key == Key_PAD_4 {

		if !numlock {
			return Key_Behavior_LEFT
		}
	}

	if key == Key_PAD_5 {

		if !numlock {
			return Key_Behavior_NIL
		}
	}

	if key == Key_PAD_6 {

		if !numlock {
			return Key_Behavior_RIGHT
		}
	}

	if key == Key_PAD_7 {

		if !numlock {
			return Key_Behavior_HOME
		}
	}

	if key == Key_PAD_8 {

		if !numlock {
			return Key_Behavior_UP
		}
	}

	if key == Key_PAD_9 {

		if !numlock {
			return Key_Behavior_PAGE_UP
		}
	}

	if key == Key_PAD_0 {

		if !numlock {
			return Key_Behavior_INSERT
		}
	}

	if key == Key_F1 || key == Key_F2 || key == Key_F3 || key == Key_F4 || key == Key_F5 || key == Key_F6 || key == Key_F7 || key == Key_F8 || key == Key_F9 || key == Key_F10 || key == Key_F11 || key == Key_F12 {

		return Key_Behavior_FUNCTION_KEY
	}

	if key == Key_PAD_DECIMAL {

		if !numlock {
			return Key_Behavior_DELETE
		}
	}

	if key == Key_PAD_ENTER {
		return Key_Behavior_ENTER
	}

	if key == Key_RIGHT {
		return Key_Behavior_RIGHT
	}

	if key == Key_LEFT {
		return Key_Behavior_LEFT
	}

	if key == Key_UP {
		return Key_Behavior_UP
	}

	if key == Key_DOWN {
		return Key_Behavior_DOWN
	}

	if key == Key_PAGE_UP {
		return Key_Behavior_PAGE_UP
	}

	if key == Key_PAGE_DOWN {
		return Key_Behavior_PAGE_DOWN
	}

	if key == Key_HOME {
		return Key_Behavior_HOME
	}

	if key == Key_END {
		return Key_Behavior_END
	}

	if key == Key_DELETE {
		return Key_Behavior_DELETE
	}

	if key == Key_INSERT {
		return Key_Behavior_INSERT
	}

	if key == Key_BACKSPACE {
		return Key_Behavior_BACKSPACE
	}

	if key == Key_TAB {
		return Key_Behavior_TAB
	}

	if key == Key_ENTER {
		return Key_Behavior_ENTER
	}

	if key == Key_ESCAPE {
		return Key_Behavior_ESCAPE
	}

	return Key_Behavior_UNKNOWN
}
