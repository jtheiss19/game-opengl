package objects

import (
	"game/internal/ecs"
	"game/internal/ecs/components"
)

func NewPlayerCamera(world *ecs.World) {
	compCameraTransform := &components.Transform3D{
		BaseComponent: &ecs.BaseComponent{},
		WorldPosition: [3]float32{0, 0, 0},
		WorldRotation: [3]float32{1.57, 0, 0},
		WorldScale:    [3]float32{1, 1, 1},
	}
	compCamera := &components.Camera{
		BaseComponent: &ecs.BaseComponent{},
		IsActive:      true,
	}
	compInput := &components.Input{
		BaseComponent: &ecs.BaseComponent{},
	}
	world.AddEntity([]ecs.Component{compCamera, compCameraTransform, compInput})
}
