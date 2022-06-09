package objects

import (
	"game/internal/ecs"
	"game/internal/ecs/components"

	"github.com/go-gl/mathgl/mgl32"
)

func NewSprite(world *ecs.World, position mgl32.Vec3) {
	// CUBE
	compTransform := &components.Transform3D{
		BaseComponent: &ecs.BaseComponent{},
		WorldPosition: position,
		WorldRotation: [3]float32{0, 0, 0},
		WorldScale:    [3]float32{1, 1, 1},
	}
	compSprite := &components.Sprite{
		BaseComponent: &ecs.BaseComponent{},
		Image:         "/some/file/path",
	}
	world.AddEntity([]ecs.Component{compSprite, compTransform})
}
