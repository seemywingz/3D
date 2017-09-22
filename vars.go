package gg

import "github.com/go-gl/glfw/v3.2/glfw"

// SceneLogic :
type SceneLogic func(s *StdData)

// StdData :
type StdData struct {
	Position
	Program    uint32
	XRotation  float32
	YRotation  float32
	SceneLogic SceneLogic
}

// Position : struct to store 3D coords
type Position struct {
	X float32
	Y float32
	Z float32
}

// NewPosition : create a new Position, for looks
func NewPosition(x, y, z float32) Position {
	return Position{X: x, Y: y, Z: z}
}

// Color : struct to store RGBA values
type Color struct {
	R float32
	G float32
	B float32
	A float32
}

// NewColor : create a new Color, you know, for looks
func NewColor(r, g, b, a float32) Color {
	return Color{R: r, G: g, B: b, A: a}
}

var (

	// unexported
	window       *glfw.Window
	camera       *Camera
	lightManager *LightManager

	// Feature : used to enable features
	Feature map[int]bool

	// Shader : map of gg provided shaders
	Shader map[string]uint32

	// Triangle :
	Triangle = []float32{
		-1.0, -1.0, 0, 1.0, 0.0, 0.0, 0.0, 1.0,
		0, 1.0, 0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
	}

	// Square :
	Square = []float32{
		//  X, Y, Z, U, V, normal(3)
		-1.0, -1.0, 0, 0.0, 1.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 0, 0.0, 0.0, 0.0, 0.0, 1.0,

		-1.0, 1.0, 0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.0, 1.0, 0, 1.0, 0.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
	}

	// CardFront :
	CardFront = []float32{
		//  X, Y, Z, U, V, normal(3)
		-1.25, -1.75, 0, 0.0, 1.0, 0.0, 0.0, 1.0,
		1.25, -1.75, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
		-1.25, 1.75, 0, 0.0, 0.0, 0.0, 0.0, 1.0,

		-1.25, 1.75, 0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.25, 1.75, 0, 1.0, 0.0, 0.0, 0.0, 1.0,
		1.25, -1.75, 0, 1.0, 1.0, 0.0, 0.0, 1.0,
	}

	// CardBack :
	CardBack = []float32{
		-1.25, 1.75, -0.01, 1.0, 0.0, 0.0, 0.0, -1.0, // left top
		-1.25, -1.75, -0.01, 1.0, 1.0, 0.0, 0.0, -1.0, // left bottom
		1.25, -1.75, -0.01, 0.0, 1.0, 0.0, 0.0, -1.0, // right bottom

		-1.25, 1.75, -0.01, 1.0, 0.0, 0.0, 0.0, -1.0, // left top
		1.25, 1.75, -0.01, 0.0, 0.0, 0.0, 0.0, -1.0, // right top
		1.25, -1.75, -0.01, 0.0, 1.0, 0.0, 0.0, -1.0, //right bottom
	}

	// Cube :
	Cube = []float32{
		//  X, Y, Z, U, V, normal(3)
		// Bottom
		-1.0, -1.0, -1.0, 0.0, 0.0, 0.0, -1.0, 0.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 1.0, 0.0, -1.0, 0.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, -1.0, 0.0,
		1.0, -1.0, 1.0, 1.0, 1.0, 0.0, -1.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 1.0, 0.0, -1.0, 0.0,

		// Top
		-1.0, 1.0, -1.0, 0.0, 0.0, 0.0, 1.0, 0.0,
		-1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0,
		1.0, 1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 0.0,
		1.0, 1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 0.0,
		-1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0,
		1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 1.0, 0.0,

		// Front
		-1.0, -1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 1.0,
		1.0, -1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0,
		1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 0.0, 1.0,
		-1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 1.0,

		// Back
		-1.0, -1.0, -1.0, 0.0, 0.0, 0.0, 0.0, -1.0,
		-1.0, 1.0, -1.0, 0.0, 1.0, 0.0, 0.0, -1.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, 0.0, -1.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 0.0, 0.0, -1.0,
		-1.0, 1.0, -1.0, 0.0, 1.0, 0.0, 0.0, -1.0,
		1.0, 1.0, -1.0, 1.0, 1.0, 0.0, 0.0, -1.0,

		// Left
		-1.0, -1.0, 1.0, 0.0, 1.0, -1.0, 0.0, 0.0,
		-1.0, 1.0, -1.0, 1.0, 0.0, -1.0, 0.0, 0.0,
		-1.0, -1.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 1.0, -1.0, 0.0, 0.0,
		-1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 0.0, 0.0,
		-1.0, 1.0, -1.0, 1.0, 0.0, -1.0, 0.0, 0.0,

		// Right
		1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0,
		1.0, -1.0, -1.0, 1.0, 0.0, 1.0, 0.0, 0.0,
		1.0, 1.0, -1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
		1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0,
		1.0, 1.0, -1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
		1.0, 1.0, 1.0, 0.0, 1.0, 1.0, 0.0, 0.0,
	}
)
