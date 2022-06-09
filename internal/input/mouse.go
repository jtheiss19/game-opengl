package input

import "github.com/go-gl/glfw/v3.3/glfw"

var (
	mouseLastMeasuredX = 0.0
	mouseLastMeasuredY = 0.0

	sensitivity = 15.0
)

func CurrentMousePosition() (x, y float64) {
	return glfw.GetCurrentContext().GetCursorPos()
}

func MousePositionDifference() (float64, float64) {
	x, y := CurrentMousePosition()

	x = x / sensitivity
	y = y / sensitivity

	returnX := mouseLastMeasuredX - x
	returnY := mouseLastMeasuredY - y

	mouseLastMeasuredX = x
	mouseLastMeasuredY = y

	return returnX, returnY
}
