package engine

import "github.com/go-gl/gl/v4.6-core/gl"

func OpenGLInit() {
	// initOpenGL initializes OpenGL
	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.DEPTH_TEST)

}
