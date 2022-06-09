package engine

import (
	"reflect"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type InstancedObject struct {
	bufferID glObjectReference
	vboID    glObjectReference
	eboID    glObjectReference
	vaoID    glObjectReference

	indiciesCount int32
	models        []mgl32.Mat4
}

func MakeInstancedObject(modelMatrices []mgl32.Mat4, objectPoints []float32, indicies []uint32) *InstancedObject {

	var eboRef uint32
	// creates ebo
	gl.GenBuffers(1, &eboRef)

	// bind it
	gl.BindBuffer(gl.ARRAY_BUFFER, eboRef)

	// put our data in the buffer
	dataSizeIndicies := int(reflect.TypeOf(indicies).Elem().Size())
	gl.BufferData(gl.ARRAY_BUFFER, dataSizeIndicies*len(indicies), gl.Ptr(indicies), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	var vboRef uint32
	// creates vbo
	gl.GenBuffers(1, &vboRef)

	// set gl.ARRAY_BUFFER to use data at vboRef
	gl.BindBuffer(gl.ARRAY_BUFFER, vboRef)

	// we put data in the buffer
	dataSizePoints := int32(reflect.TypeOf(objectPoints).Elem().Size())
	dataSizePointsInt := int(dataSizePoints)
	gl.BufferData(gl.ARRAY_BUFFER, dataSizePointsInt*len(objectPoints), gl.Ptr(objectPoints), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	var vaoRef uint32
	// creates vao
	gl.GenVertexArrays(1, &vaoRef)

	// sets the vao to be used
	gl.BindVertexArray(vaoRef)

	// bind vbo
	gl.BindBuffer(gl.ARRAY_BUFFER, vboRef)

	// Tell it how our data looks in vbo
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*dataSizePoints, nil)                               // Position
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*dataSizePoints, gl.PtrOffset(3*dataSizePointsInt)) // Color
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*dataSizePoints, gl.PtrOffset(6*dataSizePointsInt)) // Texture
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.EnableVertexAttribArray(2)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	// ---------- Handle Instancing
	var buffer uint32
	// creates buffer
	gl.GenBuffers(1, &buffer)

	// set gl.ARRAY_BUFFER to use data at vboRef
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)

	// we put data in the buffer
	dataSizeModels := int32(reflect.TypeOf(modelMatrices).Elem().Size())
	dataSizeModelsInt := int(dataSizeModels)
	gl.BufferData(gl.ARRAY_BUFFER, dataSizeModelsInt*len(modelMatrices), gl.Ptr(modelMatrices), gl.STATIC_DRAW)

	vec4Size := int32(reflect.TypeOf(mgl32.Vec4{}).Size())
	vec4SizeInt := int(vec4Size)

	gl.VertexAttribPointer(3, 4, gl.FLOAT, false, dataSizeModels, nil)
	gl.VertexAttribPointer(4, 4, gl.FLOAT, false, dataSizeModels, gl.PtrOffset(1*vec4SizeInt))
	gl.VertexAttribPointer(5, 4, gl.FLOAT, false, dataSizeModels, gl.PtrOffset(2*vec4SizeInt))
	gl.VertexAttribPointer(6, 4, gl.FLOAT, false, dataSizeModels, gl.PtrOffset(3*vec4SizeInt))
	gl.EnableVertexAttribArray(3)
	gl.EnableVertexAttribArray(4)
	gl.EnableVertexAttribArray(5)
	gl.EnableVertexAttribArray(6)
	gl.VertexAttribDivisor(3, 1)
	gl.VertexAttribDivisor(4, 1)
	gl.VertexAttribDivisor(5, 1)
	gl.VertexAttribDivisor(6, 1)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return &InstancedObject{
		bufferID:      glObjectReference(buffer),
		vboID:         glObjectReference(vboRef),
		eboID:         glObjectReference(eboRef),
		vaoID:         glObjectReference(vaoRef),
		indiciesCount: int32(len(indicies)),
		models:        modelMatrices,
	}
}

func (obj *InstancedObject) Draw() {

	// bind
	gl.BindVertexArray(uint32(obj.vaoID))
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, uint32(obj.eboID))

	// draw
	gl.DrawElementsInstanced(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil, int32(len(obj.models)))

	// clear
	gl.BindVertexArray(0)
}
