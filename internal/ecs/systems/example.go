package systems

import (
	"game/internal/ecs"
	"game/internal/ecs/components"
	"reflect"

	"github.com/sirupsen/logrus"
)

type ExampleRequiredComponents struct {
	Entities []*ExampleComponents
}

type ExampleComponents struct {
	ExampleTransform *components.Example
}

type ExampleSystem struct {
	trackedEntities *ExampleRequiredComponents
}

func NewExampleSystem() *ExampleSystem {
	return &ExampleSystem{
		trackedEntities: &ExampleRequiredComponents{},
	}
}

func (es *ExampleSystem) Update(dt float32) {
	// Ran on world update
}

func (es *ExampleSystem) GetRequiredComponents() []reflect.Type {
	reqComponentsStruct := &ExampleRequiredComponents{}

	v := reflect.ValueOf(reqComponentsStruct).Elem()

	returnTypes := []reflect.Type{}
	for j := 0; j < v.NumField(); j++ {
		reqField := v.Field(j)
		switch reqField.Type().Kind() {
		case reflect.Slice:
			returnTypes = append(returnTypes, reqField.Type().Elem())
		case reflect.Ptr:
			returnTypes = append(returnTypes, reqField.Elem().Type())
		default:
			logrus.Error("no field match found")
		}
	}

	return returnTypes
}

func (es *ExampleSystem) AddEntity(comps []ecs.Component) {
	logrus.Info("adding entity to sprite renderer system")

	for _, reqComp := range es.GetRequiredComponents() {
		if ecs.SatisfySystemRequirements(comps, reqComp) {

			f := reflect.ValueOf(es.trackedEntities).Elem()
			for j := 0; j < f.NumField(); j++ {
				reqField := f.Field(j)
				reqFieldType := reqField.Type().Elem()
				if reqFieldType == reqComp {
					newReqFieldEntry := reflect.New(reqFieldType.Elem())
					ecs.Fill(newReqFieldEntry, comps)

					reqFieldElem := reqField

					reqFieldElem.Set(reflect.Append(reqFieldElem, newReqFieldEntry))
				}
			}
		}
	}

}

func (es *ExampleSystem) RemoveEntity(id ecs.ID) {
	// Called when a entity needs removed
}

func (es *ExampleSystem) Initilizer() {
	// some code thats ran on world join
}
