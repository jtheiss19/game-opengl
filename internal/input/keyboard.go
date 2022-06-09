package input

import "github.com/go-gl/glfw/v3.3/glfw"

func GetKeyState(keyToObserve glfw.Key) bool {
	return glfw.GetCurrentContext().GetKey(keyToObserve) == glfw.Press
}
