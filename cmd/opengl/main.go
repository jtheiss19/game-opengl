package main

import (
	"game/internal/ecs"
	"game/internal/ecs/objects"
	"game/internal/ecs/systems"
	"game/internal/engine"
	"game/internal/input"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sirupsen/logrus"
)

const (
	width      = 1000
	height     = 1000
	windowName = "Test Game"

	fpsCap = 60
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	logrus.Info("starting up...")
	runtime.LockOSThread()

	window := engine.CreateWindow(height, width, windowName)
	defer glfw.Terminate()

	engine.OpenGLInit()

	version := gl.GoStr(gl.GetString(gl.VERSION))
	logrus.Info("OpenGL version", version)

	// Buffer Mouse Input
	glfw.GetCurrentContext().SetCursorPos(width/2, height/2)
	input.MousePositionDifference()

	logrus.Info("Creating base world")
	// WOLRD AND SYSTEMS
	world := ecs.NewWorld()

	logrus.Info("Creating base Systems")
	newRenderer := systems.NewSpriteRenderer()
	world.AddSystem(newRenderer)

	newPlayerController := systems.NewPlayerController()
	world.AddSystem(newPlayerController)

	logrus.Info("Creating base objects")
	objects.NewSprite(world, mgl32.Vec3{0, 0, 8})
	objects.NewSprite(world, mgl32.Vec3{2, 0, 8})
	objects.NewPlayerCamera(world)

	// Run Program on window
	logrus.Info("Running Game...")
	for !window.ShouldClose() {
		timeStart := time.Now()

		{ // Screen Management
			// Remove last drawing
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			// Close Key Held
			if input.GetKeyState(glfw.KeyEscape) {
				window.SetShouldClose(true)
			}

			// Run Systems
			world.Update()

			// Post Draw Operations
			glfw.PollEvents()
			window.SwapBuffers()

		}

		delay := time.Second / fpsCap
		time.Sleep(delay - time.Since(timeStart))
	}
}
