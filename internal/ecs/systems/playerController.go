package systems

import (
	"game/internal/ecs"
	"game/internal/ecs/components"
	"game/internal/input"
	"reflect"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sirupsen/logrus"
)

type PlayerControllerComponents struct {
	PlayerTransform *components.Transform3D
	PlayerInput     *components.Input
}

type PlayerControllerRequiredComponents struct {
	Entities []*PlayerControllerComponents
}

type PlayerControllerSystem struct {
	trackedEntities *PlayerControllerRequiredComponents
}

func NewPlayerController() *PlayerControllerSystem {
	return &PlayerControllerSystem{
		trackedEntities: &PlayerControllerRequiredComponents{},
	}
}

func (pc *PlayerControllerSystem) Update(dt float32) {

	// Handle Input Events
	for _, entity := range pc.trackedEntities.Entities {

		mainCamera := entity.PlayerTransform
		// Mouse
		//diffX, diffy := input.MousePositionDifference()
		//mainCamera.AddWorldRotation(diffX, diffy)

		// Keyboard
		moveDirection := mgl32.Vec3{}
		forwardVec := mainCamera.GetForwardVector()

		rightVec := forwardVec.Cross(mgl32.Vec3{0, 1, 0})
		if input.GetKeyState(glfw.KeyW) {
			moveDirection = moveDirection.Add(forwardVec)
		}
		if input.GetKeyState(glfw.KeyS) {
			moveDirection = moveDirection.Add(forwardVec.Mul(-1))
		}
		if input.GetKeyState(glfw.KeyD) {
			moveDirection = moveDirection.Add(rightVec)
		}
		if input.GetKeyState(glfw.KeyA) {
			moveDirection = moveDirection.Add(rightVec.Mul(-1))
		}

		mainCamera.AddWorldPosition(moveDirection.Mul(0.5))
	}
}

func (pc *PlayerControllerSystem) GetRequiredComponents() []reflect.Type {
	reqComponentsStruct := &PlayerControllerRequiredComponents{}

	v := reflect.ValueOf(reqComponentsStruct).Elem()

	returnTyppc := []reflect.Type{}
	for j := 0; j < v.NumField(); j++ {
		reqField := v.Field(j)
		switch reqField.Type().Kind() {
		case reflect.Slice:
			returnTyppc = append(returnTyppc, reqField.Type().Elem())
		case reflect.Ptr:
			returnTyppc = append(returnTyppc, reqField.Elem().Type())
		default:
			logrus.Error("no field match found")
		}
	}

	return returnTyppc
}

func (pc *PlayerControllerSystem) AddEntity(comps []ecs.Component) {
	logrus.Info("adding entity to Player Controller System")

	for _, reqComp := range pc.GetRequiredComponents() {
		if ecs.SatisfySystemRequirements(comps, reqComp) {

			f := reflect.ValueOf(pc.trackedEntities).Elem()
			for j := 0; j < f.NumField(); j++ {
				reqField := f.Field(j)
				reqFieldType := reqField.Type().Elem()
				if reqFieldType == reqComp {
					newReqFieldEntry := reflect.New(reqFieldType.Elem())
					ecs.Fill(newReqFieldEntry, comps)

					reqFieldElem := reqField

					logrus.Debug("Setting Player Controller System entity element")
					reqFieldElem.Set(reflect.Append(reqFieldElem, newReqFieldEntry))
				}
			}
		}
	}

}

func (pc *PlayerControllerSystem) RemoveEntity(id ecs.ID) {
	// Called when a entity needs removed
}

func (pc *PlayerControllerSystem) Initilizer() {
	// some code thats ran on world join
}
