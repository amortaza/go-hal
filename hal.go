package hal

type ButtonAction int

const (
	Button_Action_Down ButtonAction = 1 + iota
	Button_Action_Up
)

type Canvas interface {
	Begin()
	End()

	GetWidth() int
	GetHeight() int

	Clear(red, green, blue float32)

	Paint(seeThru bool, left, top int, alphas []float32)

	Free()
}

type Graphics interface {
	Clear(red, green, blue float32)

	PushView(width, height int)

	PopView()

	NewCanvas(width, height int) Canvas
}

type HAL interface {

	Start(
		title string,

		width, height int,

		onAfterGL,
		onLoop,
		onBeforeDeleteWindow func(),

		onWindowResize,

		onMouseMove func(int,int),
		onMouseButton func(MouseButton, ButtonAction),

		onKey func(KeyboardKey, ButtonAction, bool, bool, bool))

	GetWindowDim()(width, height int)

	GetMousePos()(x,y int)

	SetMouseCursor(cursor MouseCursor)

	GetGraphics() Graphics
}

