package hal

type HAL interface {

	Start(
		title string,

		left, top, width, height int,

		onAfterGL,
		onLoop,
		onBeforeDeleteWindow func(),

		onWindowResize,

		onMouseMove func(x, y int),
		onMouseButton func(MouseButton, ButtonAction),

		onKey func(KeyboardKey, ButtonAction, bool, bool, bool))

	GetWindowDim()(width, height int)

	GetMousePos()(x, y int)

	SetMouseCursor(cursor MouseCursor)

	GetGraphics() Graphics
}

