package engine

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sirupsen/logrus"
)

func CreateWindow(height, width int, windowName string) *glfw.Window {
	logrus.Trace("initing glfw window")
	if err := glfw.Init(); err != nil {
		logrus.Panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	logrus.Info("creating glfw window")
	window, err := glfw.CreateWindow(width, height, windowName, nil, nil)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Trace("changing window context")
	window.MakeContextCurrent()

	return window
}
