package ecs

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

type system interface {
	Update(dt float32)
	GetRequiredComponents() []reflect.Type
	AddEntity(comps []Component)
	RemoveEntity(id ID)
	Initilizer()
}

type RequiredComponents interface{}

func Fill(thingToFill reflect.Value, compsToFillWith []Component) {
	v := thingToFill.Elem()

	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		for _, comp := range compsToFillWith {
			if reflect.TypeOf(comp) == f.Type() {
				logrus.Trace("system field match, setting")
				f.Set(reflect.ValueOf(comp))
			}
		}
	}
}
