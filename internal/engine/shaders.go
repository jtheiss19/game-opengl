package engine

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sirupsen/logrus"
)

type Shader struct {
	programID  glObjectReference
	fragmentID glObjectReference
	vertexID   glObjectReference

	filePathFragment string
	filePathVertex   string

	modTime time.Time
}

func CreateShaderProgram(pathToFragmentShader, pathToVertexShader string) (*Shader, error) {
	logrus.Info("creating new shader from file")
	prog := gl.CreateProgram()

	logrus.Trace("compiling new fragment shader file")
	fragmentID, err := compileShader(pathToFragmentShader, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, errors.New("could not compile shader file: " + err.Error())
	}

	logrus.Trace("compiling new vertex shader file")
	vertexID, err := compileShader(pathToVertexShader, gl.VERTEX_SHADER)
	if err != nil {
		return nil, errors.New("could not compile shader file: " + err.Error())
	}

	modtime := getLastModTime(pathToFragmentShader, pathToVertexShader)

	logrus.Trace("creating new shader object")
	returnShader := &Shader{
		programID:        glObjectReference(prog),
		fragmentID:       glObjectReference(fragmentID),
		vertexID:         glObjectReference(vertexID),
		filePathFragment: pathToFragmentShader,
		filePathVertex:   pathToVertexShader,
		modTime:          modtime,
	}

	returnShader.compileProgram()

	return returnShader, nil
}

func (shade *Shader) compileProgram() {
	logrus.Debug("compililing shader program")
	gl.AttachShader(uint32(shade.programID), uint32(shade.fragmentID))
	gl.AttachShader(uint32(shade.programID), uint32(shade.vertexID))
	gl.LinkProgram(uint32(shade.programID))
}

func (shade *Shader) Use() {
	gl.UseProgram(uint32(shade.programID))
}

func getLastModTime(pathToFragmentShader string, pathToVertexShader string) time.Time {
	initialStatFragment, err := os.Stat(pathToFragmentShader)
	if err != nil {
		return time.Time{}
	}

	initialStatVertex, err := os.Stat(pathToVertexShader)
	if err != nil {
		return time.Time{}
	}

	var modtime time.Time
	if initialStatFragment.ModTime().After(initialStatVertex.ModTime()) {
		modtime = initialStatFragment.ModTime()
	} else {
		modtime = initialStatVertex.ModTime()
	}
	return modtime
}

func (shade *Shader) CheckForChanges() {

	logrus.Trace("Reading file changes")
	modtime := getLastModTime(shade.filePathFragment, shade.filePathVertex)

	logrus.Trace("comparing shader changes to past values")
	if modtime != shade.modTime {
		shade.modTime = modtime

		logrus.Trace("updating shader values and recompiling")

		logrus.Trace("compiling new fragment shader file")
		fragmentID, err := compileShader(shade.filePathFragment, gl.FRAGMENT_SHADER)
		if err != nil {
			logrus.Error("could not compile shader file: " + err.Error())
			return
		}

		logrus.Trace("compiling new vertex shader file")
		vertexID, err := compileShader(shade.filePathVertex, gl.VERTEX_SHADER)
		if err != nil {
			logrus.Error("could not compile shader file: " + err.Error())
			return
		}

		logrus.Trace("swapping out shader object reference on GPU")
		gl.DetachShader(uint32(shade.programID), uint32(shade.fragmentID))
		gl.DetachShader(uint32(shade.programID), uint32(shade.vertexID))
		gl.DeleteShader(uint32(shade.fragmentID))
		gl.DeleteShader(uint32(shade.vertexID))
		shade.fragmentID = fragmentID
		shade.vertexID = vertexID
		shade.compileProgram()
	}
}

func compileShader(filePath string, shaderType uint32) (glObjectReference, error) {
	logrus.Trace("reading new shader file")
	shaderBytes, err := os.ReadFile(filePath)
	if err != nil {
		return 0, errors.New("could not read shader file: " + err.Error())
	}

	source := string(shaderBytes)

	source = source + "\x00"

	logrus.Trace("creating shader")
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	defer free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return glObjectReference(shader), nil
}

func (shade *Shader) Delete() {
	logrus.Info("releasing held reasources from shader")
	gl.DeleteProgram(uint32(shade.fragmentID))
	gl.DeleteProgram(uint32(shade.vertexID))
	gl.DeleteProgram(uint32(shade.programID))
}

func (shade *Shader) SetUniformMat4(name string, mat mgl32.Mat4) {
	name_cstr := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(uint32(shade.programID), name_cstr)
	m4 := [16]float32(mat)
	gl.UniformMatrix4fv(location, 1, false, &m4[0])
}
