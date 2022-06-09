package systems

import (
	"game/internal/ecs"
	"game/internal/ecs/components"
	"game/internal/engine"
	"game/internal/shapes"
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/sirupsen/logrus"
)

var (
	pathToFrag = `C:\Users\Jthei\OneDrive\Desktop\game\assets\shaders\instancedshader.fs`
	pathToVert = `C:\Users\Jthei\OneDrive\Desktop\game\assets\shaders\instancedshader.vs`

	//pathToFrag = `C:\Users\Jthei\OneDrive\Desktop\game\assets\shaders\fragmentshader.fs`
	// pathToVert = `C:\Users\Jthei\OneDrive\Desktop\game\assets\shaders\vertexshader.vs`

	texturePath = `C:\Users\Jthei\OneDrive\Desktop\game\assets\textures\tiles.png`
)

// -------------------------- Required Components -----------------------------------------//
type SpriteRendererRequiredComponents struct {
	Entities []*spriteRendererComponents
	Camera   []*SpecialSpriteRendererComponents
}

type SpecialSpriteRendererComponents struct {
	Camera          *components.Camera
	CameraTransform *components.Transform3D
}

type spriteRendererComponents struct {
	SpriteComp   *components.Sprite
	TrasformComp *components.Transform3D
}

// -------------------------- Base Class -----------------------------------------//
type SpriteRenderer struct {
	trackedEntities *SpriteRendererRequiredComponents

	cube *engine.Object

	instancedObject *engine.InstancedObject

	texture *engine.Texture
	shader  *engine.Shader
}

func NewSpriteRenderer() *SpriteRenderer {
	return &SpriteRenderer{
		trackedEntities: &SpriteRendererRequiredComponents{},
		cube:            &engine.Object{},
		texture:         &engine.Texture{},
		shader:          &engine.Shader{},
	}
}

// -------------------------- Custom Functionality -----------------------------------------//

func (ts *SpriteRenderer) Update(dt float32) {
	logrus.Trace("updating sprite renderer system")

	if !ts.trackedEntities.Entities[0].SpriteComp.Created {
		logrus.Debug("creating Object Instance")
		modelMatries := []mgl32.Mat4{}
		for _, entity := range ts.trackedEntities.Entities {
			modelMatrix := entity.TrasformComp.GetModelMatrix()
			modelMatries = append(modelMatries, modelMatrix)
		}
		ts.instancedObject = shapes.NewSquareInstance(modelMatries)
		ts.trackedEntities.Entities[0].SpriteComp.Created = true
	}

	viewMatrix := mgl32.Mat4{}
	projectionMatrix := mgl32.Mat4{}

	for _, camera := range ts.trackedEntities.Camera {

		lookPosition := camera.CameraTransform.GetForwardVector()
		viewMatrix = mgl32.LookAtV(
			camera.CameraTransform.WorldPosition,
			camera.CameraTransform.WorldPosition.Add(lookPosition),
			mgl32.Vec3{0, 1, 0},
		)

		projectionMatrix = mgl32.Perspective(mgl32.DegToRad(45), 1, 0.1, 1000)
	}

	ts.shader.Use()
	ts.shader.SetUniformMat4("projection", projectionMatrix)
	ts.shader.SetUniformMat4("view", viewMatrix)

	ts.texture.Use()
	ts.instancedObject.Draw()

	ts.shader.CheckForChanges()
}

func (ts *SpriteRenderer) Initilizer() {
	ts.texture = engine.NewTexture(texturePath)
	// Create shader
	shader, err := engine.CreateShaderProgram(pathToFrag, pathToVert)
	if err != nil {
		logrus.Panic(err)
	}

	ts.shader = shader
}

// -------------------------- Boilerplate Code -----------------------------------------//

func (ts *SpriteRenderer) GetRequiredComponents() []reflect.Type {
	reqComponentsStruct := &SpriteRendererRequiredComponents{}

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

func (ts *SpriteRenderer) AddEntity(comps []ecs.Component) {
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

func (ts *SpriteRenderer) RemoveEntity(id ecs.ID) {
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
