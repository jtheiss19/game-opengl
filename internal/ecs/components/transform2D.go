package components

import (
	"game/internal/ecs"

	"github.com/go-gl/mathgl/mgl32"
)

type Transform2D struct {
	*ecs.BaseComponent

	WorldPosition mgl32.Vec2
	WorldRotation float32
	WorldScale    mgl32.Vec2
}

func (tf *Transform2D) AddWorldPosition(positionToAdd mgl32.Vec2) {
	tf.WorldPosition = tf.WorldPosition.Add(positionToAdd)
}

func (tf *Transform2D) AddWorldRotation(angle float32) {
	tf.WorldRotation = tf.WorldRotation + angle
}

func (tf *Transform2D) GetModelMatrix() mgl32.Mat3 {
	modelMatrix := mgl32.Ident3()

	// Rotation

	// Scale

	// Translation
	modelMatrix = mgl32.Translate2D(tf.WorldPosition.Elem()).Mul3(modelMatrix)

	return modelMatrix
}
