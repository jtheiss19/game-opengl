package components

import (
	"game/internal/ecs"
)

type Sprite struct {
	*ecs.BaseComponent

	Image   string
	Created bool

	SpriteSheetX, SpriteSheetY int
}
