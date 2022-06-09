package systems

import (
	"game/internal/ecs"
	"game/internal/ecs/components"
	"game/internal/wrapper"
	"reflect"

	_ "image/png"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/sirupsen/logrus"
)

var (
	ebitenTexturePath = `C:\Users\Jthei\OneDrive\Desktop\game\assets\textures\tiles.png`
)

// -------------------- SPECIAL COMPONENTS -------------------------------------------//
type EbitenSpriteRendererRequiredComponents struct {
	Entities []*EbitenSpriteRendererComponents
	Camera   []*SpecialEbitenSpriteRendererComponents
}

type SpecialEbitenSpriteRendererComponents struct {
	Camera          *components.Camera
	CameraTransform *components.Transform3D
}

type EbitenSpriteRendererComponents struct {
	SpriteComp   *components.Sprite
	TrasformComp *components.Transform3D
}

// -------------------- Main Component -------------------------------------------//

type EbitenSpriteRenderer struct {
	trackedEntities *EbitenSpriteRendererRequiredComponents

	texture *wrapper.Texture
}

func NewEbitenSpriteRenderer() *EbitenSpriteRenderer {
	return &EbitenSpriteRenderer{
		trackedEntities: &EbitenSpriteRendererRequiredComponents{},
	}
}

// -------------------- Custom Functionality -------------------------------------------//

func (ts *EbitenSpriteRenderer) Update(dt float32) {
	logrus.Trace("updating sprite renderer system")

	ts.texture.Draw(
		mgl32.Vec2{0, 0},
		mgl32.Vec2{1, 1},
		0,
	)
}

func (ts *EbitenSpriteRenderer) Initilizer() {
	logrus.Trace("initilizing sprite renderer system")
	ts.texture = wrapper.NewTexture(ebitenTexturePath)
}

// -------------------- BoilerPlate Code -------------------------------------------//

func (ts *EbitenSpriteRenderer) GetRequiredComponents() []reflect.Type {
	reqComponentsStruct := &EbitenSpriteRendererRequiredComponents{}

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

func (ts *EbitenSpriteRenderer) AddEntity(comps []ecs.Component) {
	logrus.Info("adding entity to sprite renderer system")

	for _, reqComp := range ts.GetRequiredComponents() {
		if ecs.SatisfySystemRequirements(comps, reqComp) {

			f := reflect.ValueOf(ts.trackedEntities).Elem()
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

func (ts *EbitenSpriteRenderer) RemoveEntity(id ecs.ID) {
	logrus.Info("removing entity from sprite renderer system")

	var delete int = -1
	for index, entity := range ts.trackedEntities.Entities {
		if entity.SpriteComp.GetComponentID() == id {
			delete = index
		}
	}
	if delete >= 0 {
		logrus.Info("removing now from sprite renderer system")
		ts.trackedEntities.Entities = append(ts.trackedEntities.Entities[:delete], ts.trackedEntities.Entities[delete+1:]...)
	} else {
		var delete int = -1
		for index, entity := range ts.trackedEntities.Camera {
			if entity.Camera.GetComponentID() == id {
				delete = index
			}
		}
		if delete >= 0 {
			logrus.Info("removing now from sprite renderer system")
			ts.trackedEntities.Camera = append(ts.trackedEntities.Camera[:delete], ts.trackedEntities.Camera[delete+1:]...)
		}
	}
}
