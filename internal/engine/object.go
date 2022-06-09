package engine

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Object struct {
	vaoID         glObjectReference
	vboID         glObjectReference
	eboID         glObjectReference
	texture       *Texture
	worldPosition mgl32.Vec3
	worldRotation mgl32.Vec3
	data          []float32
	indiciesCount int32
}

func MakeObject(points []float32, indicies []uint32) *Object {
	var vboRef uint32
	// creates vbo
	gl.GenBuffers(1, &vboRef)

	// set gl.ARRAY_BUFFER to use data at vboRef
	gl.BindBuffer(gl.ARRAY_BUFFER, vboRef)

	// we put data in the buffer
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var eboRef uint32
	// creates ebo
	gl.GenBuffers(1, &eboRef)

	// bind it
	gl.BindBuffer(gl.ARRAY_BUFFER, eboRef)

	// put our data in the buffer
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(indicies), gl.Ptr(indicies), gl.STATIC_DRAW)

	var vaoRef uint32
	// creates vao
	gl.GenVertexArrays(1, &vaoRef)

	// sets the vao to be used
	gl.BindVertexArray(vaoRef)

	// bind vbo
	gl.BindBuffer(gl.ARRAY_BUFFER, vboRef)

	// Tell it how our data looks in vbo
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, nil)               // Position
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4)) // Color
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4)) // Texture
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.EnableVertexAttribArray(2)

	return &Object{
		vaoID:         glObjectReference(vaoRef),
		vboID:         glObjectReference(vboRef),
		eboID:         glObjectReference(eboRef),
		worldPosition: [3]float32{0, 0, 0},
		data:          points,
		indiciesCount: int32(len(indicies)),
	}
}

func (obj *Object) SetTexutre(newTexture *Texture) {
	obj.texture = newTexture
}

func (obj *Object) SetPosition(newWorldPosition mgl32.Vec3) {
	obj.worldPosition = newWorldPosition
}

func (obj *Object) SetRotation(newWorldRotation mgl32.Vec3) {
	obj.worldRotation = newWorldRotation
}

func (obj *Object) Draw() {
	// bind
	gl.BindTexture(gl.TEXTURE_2D, uint32(obj.texture.textureID))
	gl.BindVertexArray(uint32(obj.vaoID))

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, uint32(obj.eboID))

	// draw
	gl.DrawElements(gl.TRIANGLES, obj.indiciesCount, gl.UNSIGNED_INT, nil)

	// clear
	gl.BindVertexArray(0)
}

func (obj *Object) GetModelMatrix() mgl32.Mat4 {
	modelMatrix := mgl32.Ident4()

	// Rotation

	// Scale

	// Translation
	modelMatrix = mgl32.Translate3D(obj.worldPosition.X(), obj.worldPosition.Y(), obj.worldPosition.Z()).Mul4(modelMatrix)

	return modelMatrix
}
