package hal

type MouseButton int
type MouseCursor int

const (
	Mouse_Button_Left MouseButton = 1 + iota
	Mouse_Button_Right
)

const (
	MouseCursor_Arrow MouseCursor = 1 + iota
	MouseCursor_Horiz_Resize
	MouseCursor_Vert_Resize
	MouseCursor_IBeam
	MouseCursor_CrossHair
	MouseCursor_Hand
)




