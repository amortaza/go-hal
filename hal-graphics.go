package hal

type Canvas interface {

	Clear(red, green, blue float32)

	GetWidth() int
	GetHeight() int

	Begin()
	End()

	Paint(seeThru bool, left, top int, alphas []float32)

	Free()
}

type Graphics interface {

	Clear(red, green, blue float32)

	PushView(width, height int)

	PopView()

	NewCanvas(width, height int) Canvas
}


