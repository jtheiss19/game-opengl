package shapes

import (
	"game/internal/engine"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	squareFullColored = []float32{
		// Positions      // Colors       // Texture Coords
		0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0, // Top Right
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, // Bottom Right
		-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, // Bottom Left
		-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, 0.0, 1.0, // Top Left
	}

	squareFull = []float32{
		// Positions      // Colors       // Texture Coords
		0.5, 0.5, 0, 1.0, 1.0, 1.0, 1.0, 1.0, // Top Right
		0.5, -0.5, 0, 1.0, 1.0, 1.0, 1.0, 0.0, // Bottom Right
		-0.5, -0.5, 0.0, 1.0, 1.0, 1.0, 0.0, 0.0, // Bottom Left
		-0.5, 0.5, 0.0, 1.0, 1.0, 1.0, 0.0, 1.0, // Top Left
	}

	squarePos = []float32{
		// Positions
		0.5, 0.5, 0,
		0.5, -0.5, 0,
		-0.5, -0.5, 0.0,
		-0.5, 0.5, 0.0,
	}

	squareColor = []float32{
		1.0, 1.0, 1.0,
		1.0, 1.0, 1.0,
		1.0, 1.0, 1.0,
		1.0, 1.0, 1.0,
	}

	squareTex = []float32{
		1.0, 1.0,
		1.0, 0.0,
		0.0, 0.0,
		0.0, 1.0,
	}

	indices = []uint32{ // Note that we start from 0!
		0, 1, 3, // First Triangle
		1, 2, 3, // Second Triangle
	}
)

func NewSquare() *engine.Object {
	return engine.MakeObject(squareFull, indices)
}

func NewSquareInstance(modelMatricies []mgl32.Mat4) *engine.InstancedObject {
	return engine.MakeInstancedObject(modelMatricies, squareFull, indices)
}
