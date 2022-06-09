package shapes

import "game/internal/engine"

var (
	cubeVertices = []float32{

		0.5, 0.5, 0.5, 1, 1, 1, 1.0, 1.0, // front top right      0
		0.5, -0.5, 0.5, 1, 1, 1, 1.0, 0, // front bottom right    1
		-0.5, -0.5, 0.5, 1, 1, 1, 0, 0, // front bottom left      2
		-0.5, 0.5, 0.5, 1, 1, 1, 0, 1.0, // front top left        3

		0.5, 0.5, -0.5, 1, 1, 1, 0, 1.0, // back top right        4
		0.5, -0.5, -0.5, 1, 1, 1, 0, 0, // back bottom right      5
		-0.5, -0.5, -0.5, 1, 1, 1, 1.0, 1.0, // back bottom left  6
		-0.5, 0.5, -0.5, 1, 1, 1, 1.0, 0, // back top left        7
	}

	cubeIndices = []uint32{
		// Front Face
		0, 1, 2, // First Triangle
		0, 2, 3, // Second Triangle

		// Left Face
		3, 2, 6, // First Triangle
		6, 7, 3, // Second Triangle

		// Right Face
		0, 5, 1, // First Triangle
		5, 0, 4, // Second Triangle

		// Top Face
		0, 3, 4, // First Triangle
		4, 3, 7, // Second Triangle

		// Bottom Face
		1, 6, 2, // First Triangle
		6, 5, 2, // Second Triangle

		// Back Face
		6, 4, 7, // First Triangle
		6, 5, 4, // Second Triangle
	}
)

func NewCube() *engine.Object {
	return engine.MakeObject(cubeVertices, cubeIndices)
}
