package hal

// This is common to Mouse button and Keyboard keys (aka Keyboard buttons)
// since they can both be pressed Up or Down.
type ButtonAction int

const (
	Button_Action_Down ButtonAction = 1 + iota
	Button_Action_Up
)

