package ecs

type Component interface {
	GetComponentID() ID
	SetComponentID(id ID)
}

type BaseComponent struct {
	entityID ID
}

func (bc *BaseComponent) GetComponentID() ID {
	return bc.entityID
}

func (bc *BaseComponent) SetComponentID(id ID) {
	bc.entityID = id
}
