package hal

type MouseButton int
type MouseCursor int

const (
	Mouse_Button_Left MouseButton = 1 + iota
	Mouse_Button_Right
)

const (
	Mouse_Cursor_Arrow MouseCursor = 1 + iota
	Mouse_Cursor_Horiz_Resize
	Mouse_Cursor_Vert_Resize
	Mouse_Cursor_IBeam
	Mouse_Cursor_CrossHair
	Mouse_Cursor_Hand
)




