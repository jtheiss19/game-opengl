package ecs

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ID string

type World struct {
	systems      []system
	entityLookup map[ID][]Component
}

func NewWorld() *World {
	return &World{
		systems:      []system{},
		entityLookup: map[ID][]Component{},
	}
}

func (wrld *World) Update() {
	for _, system := range wrld.systems {
		system.Update(0)
	}
}

func (wrld *World) AddComponent(comp Component) {
	logrus.Info("adding component")
	id := comp.GetComponentID()

	if comps, ok := wrld.entityLookup[id]; ok {
		comps = append(comps, comp)
		wrld.entityLookup[id] = comps
		wrld.checkCompsForNewSystemMatch(comps)
	} else {
		comps := []Component{comp}
		wrld.entityLookup[id] = []Component{comp}
		wrld.checkCompsForNewSystemMatch(comps)
	}

}

func (wrld *World) RemoveCompoent(comp Component) {
	logrus.Info("removing component")
	id := comp.GetComponentID()

	if comps, ok := wrld.entityLookup[id]; ok {
		newComps := []Component{}
		for _, compItem := range comps {
			if reflect.TypeOf(compItem) != reflect.TypeOf(comp) {
				newComps = append(newComps, compItem)
			}
		}
		wrld.entityLookup[id] = newComps

		for _, system := range wrld.systems {
			for _, reqComp := range system.GetRequiredComponents() {
				if !SatisfySystemRequirements(newComps, reqComp) {
					system.RemoveEntity(id)
				}
			}
		}
	}
}

func (wrld *World) AddEntity(comps []Component) {
	id := ID(uuid.New().String())
	logrus.Info("Creating and adding new entity: ", id)
	for _, comp := range comps {
		comp.SetComponentID(id)
	}

	if compsExisting, ok := wrld.entityLookup[id]; ok {
		comps = append(compsExisting, comps...)
		wrld.entityLookup[id] = comps
		wrld.checkCompsForNewSystemMatch(comps)
	} else {
		wrld.entityLookup[id] = comps
		wrld.checkCompsForNewSystemMatch(comps)
	}
}

func (wrld *World) RemoveEntity(id ID) {
	logrus.Info("removing entity")
	for _, system := range wrld.systems {
		system.RemoveEntity(id)
	}
}

func (wrld *World) AddSystem(system system) {
	logrus.Info("initing system")
	system.Initilizer()
	logrus.Info("adding system")
	wrld.systems = append(wrld.systems, system)
}

func (wrld *World) checkCompsForNewSystemMatch(comps []Component) {
	logrus.Info("checking for system matches")
	for _, system := range wrld.systems {
		requirementComponents := system.GetRequiredComponents()
		for _, requirement := range requirementComponents {
			if SatisfySystemRequirements(comps, requirement) {
				logrus.Info("found system match")
				system.AddEntity(comps)
			}
		}
	}
}

func SatisfySystemRequirements(comps []Component, reqComps reflect.Type) bool {
	if reqComps.Kind() == reflect.Ptr {
		reqComps = reqComps.Elem()
	}

	for j := 0; j < reqComps.NumField(); j++ {
		f := reqComps.Field(j).Type
		if !func() bool { // return if there is a compoenent that doesn't have a match
			for _, compArrayName := range comps {
				if reflect.TypeOf(compArrayName) == f {
					return true
				}
			}
			return false
		}() {
			return false
		}
	}
	return true
}
