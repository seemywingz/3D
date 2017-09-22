package gg

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// LightLogic :
type LightLogic func(l *Light)

// Light : struct to hold light data
type Light struct {
	Iamb    *float32
	Idif    *float32
	Ispec   *float32
	LPosID  int32
	IambID  int32
	IdifID  int32
	IspecID int32
	SceneObjectData
}

// NewLight :
func NewLight(pos Position, iamb, idif, ispec []float32, program uint32) *Light {

	LPosID := gl.GetUniformLocation(program, gl.Str("lightPos\x00"))
	IambID := gl.GetUniformLocation(program, gl.Str("Iamb\x00"))
	IdifID := gl.GetUniformLocation(program, gl.Str("Idif\x00"))
	IspecID := gl.GetUniformLocation(program, gl.Str("Ispec\x00"))

	l := Light{
		&iamb[0],
		&idif[0],
		&ispec[0],
		LPosID,
		IambID,
		IdifID,
		IspecID,
		SceneObjectData{},
	}
	l.Position = pos
	l.Program = program
	return &l
}

// Draw :
func (l *Light) Draw() {

	if l.SceneLogic != nil {
		l.SceneLogic(&l.SceneObjectData)
	}

	gl.UseProgram(l.Program)
	gl.Uniform3fv(l.LPosID, 1, &[]float32{l.X, l.Y, l.Z}[0])
	gl.Uniform3fv(l.IambID, 1, l.Iamb)
	gl.Uniform3fv(l.IdifID, 1, l.Idif)
	gl.Uniform3fv(l.IspecID, 1, l.Ispec)
}
