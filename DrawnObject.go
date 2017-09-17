package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// DrawnObject : interface for opengl drawable object
type DrawnObject interface {
	Draw()
}

// DrawnObjectData : a struct to hold openGL object data
type DrawnObjectData struct {
	Vao     uint32
	Program uint32
	Points  []float32
	Position
	ModelMatrix int32
	Texture     int32
	DrawnObjectDefaults
}

// DrawnObjectDefaults :
type DrawnObjectDefaults struct {
	XRotation float32
	YRotation float32
}

// New : Create new DrawnObjectData
func (DrawnObjectData) New(position Position, points []float32, program uint32) *DrawnObjectData {

	ptr, free1 := gl.Strs("MODEL")
	defer free1()
	ModelMatrix := gl.GetUniformLocation(program, *ptr)

	ptr, free2 := gl.Strs("tex")
	defer free2()
	Texture := gl.GetUniformLocation(program, *ptr)
	// println(texture)

	return &DrawnObjectData{
		makeVao(points, program),
		program,
		points,
		position,
		ModelMatrix,
		Texture,
		DrawnObjectDefaults{},
	}
}

// Draw : draw the vertecies
func (d *DrawnObjectData) Draw() {

	// translate to obj position
	m := mgl32.Translate3D(d.X, d.Y, d.Z)
	// rotataton
	d.YRotation++
	yrotMatrix := mgl32.HomogRotate3DY(mgl32.DegToRad(d.YRotation))
	xrotMatrix := mgl32.HomogRotate3DX(mgl32.DegToRad(d.XRotation))
	rotation := m.Mul4(yrotMatrix.Mul4(xrotMatrix))

	gl.UseProgram(d.Program)
	gl.UniformMatrix4fv(d.ModelMatrix, 1, false, &rotation[0])
	gl.BindVertexArray(d.Vao)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, boxTexture)

	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}
