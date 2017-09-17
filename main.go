package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 800
	height = 600
	title  = "go-gl Boiler"
)

var (
	camera       Camera
	shaders      []uint32
	window       *glfw.Window
	drawnObjects []DrawnObject
	boxTexture   uint32
)

func main() {
	runtime.LockOSThread()

	window = initGlfw(width, height, title)
	defer glfw.Terminate()

	initGL()
	loadShaders()
	loadLights()

	boxTexture = newTexture("textures/square.jpg")
	defer gl.DeleteTextures(1, &boxTexture)

	camera = Camera{}.New(Position{0, 0, 0}, false)

	randObject(100, -200, 200, cubeVertices)
	drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{0, 0, -2}, cubeVertices, shaders[0]))

	for !window.ShouldClose() {
		camera.Update()
		update()
	}
}

func loadLights() {
	// ambient := []float32{0.5, 0.5, 0.5, 1}
	// diffuse := []float32{1, 1, 1, 1}
	// lightPosition := []float32{-5, 5, 10, 0}
	// gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	// gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	// gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	// gl.Enable(gl.LIGHT0)
}

func randObject(numberOfObjects, min, max int, points []float32) {
	for i := 0; i < numberOfObjects; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float32(rand.Intn(max-min) + min)
		y := float32(rand.Intn(max-min) + min)
		z := float32(rand.Intn(max-min) + min)
		drawnObjects = append(drawnObjects, DrawnObjectData{}.New(Position{x, y, z}, points, shaders[0]))
	}
}

func loadShaders() {
	shaders = append(shaders, createGLprogram(basicVertexSRC, basicFragmentSRC))
}

func update() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, obj := range drawnObjects {
		obj.Draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
