package main

const (
	basicFragmentSRC = `
	#version 410
	// Hard Code Color for Now
	uniform vec4 inputColour = vec4(1,1,1,1);
	out vec4 fragColour;

	void main() {
	  fragColour = inputColour;
	}` + "\x00"

	basicVertexSRC = `
	#version 330 core

	// Input vertex data, different for all executions of this shader.
	layout(location = 0) in vec3 pos;

	//vales that stay constant for the whole mesh
	uniform mat4 MVP;
	uniform mat4 rotation;
	uniform vec4 translation;

	void main(){
	  gl_Position = MVP * ((vec4(pos,1.0) + translation));
	}` + "\x00"
)