package components

import (
	"game/internal/ecs"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Transform3D struct {
	*ecs.BaseComponent

	WorldPosition mgl32.Vec3
	WorldRotation mgl32.Vec3
	WorldScale    mgl32.Vec3
}

func (tf *Transform3D) GetForwardVector() mgl32.Vec3 {

	yaw := tf.WorldRotation.X()
	pitch := tf.WorldRotation.Y()

	yawCos := math.Cos(float64(yaw))
	yawSin := math.Sin(float64(yaw))
	pitchCos := math.Cos(float64(pitch))
	pitchSin := math.Sin(float64(pitch))

	directionX := yawCos * pitchCos
	directionY := pitchSin
	directionZ := yawSin * pitchCos

	lookDirection := mgl32.Vec3{float32(directionX), float32(directionY), float32(directionZ)}
	lookDirection = lookDirection.Normalize()

	return lookDirection
}

func (tf *Transform3D) AddWorldPosition(positionToAdd mgl32.Vec3) {
	tf.WorldPosition = tf.WorldPosition.Add(positionToAdd)
}

func (tf *Transform3D) AddWorldRotation(dYaw, dPitch float64) {

	yaw := mgl32.DegToRad(float32(dYaw))
	pitch := mgl32.DegToRad(float32(dPitch))

	newYaw := tf.WorldRotation.X() - yaw
	newPitch := tf.WorldRotation.Y() + pitch

	if newYaw > 2*math.Pi {
		newYaw = newYaw - 2*math.Pi
	} else if newYaw < 0 {
		newYaw = newYaw + 2*math.Pi
	}

	if newPitch > 2*math.Pi {
		newPitch = newPitch - 2*math.Pi
	} else if newPitch < 0 {
		newPitch = newPitch + 2*math.Pi
	}

	tf.WorldRotation = mgl32.Vec3{
		newYaw,
		newPitch,
		tf.WorldRotation.Z(),
	}
}

func (tf *Transform3D) GetModelMatrix() mgl32.Mat4 {
	modelMatrix := mgl32.Ident4()

	// Rotation

	// Scale

	// Translation
	modelMatrix = mgl32.Translate3D(tf.WorldPosition.Elem()).Mul4(modelMatrix)

	return modelMatrix
}
