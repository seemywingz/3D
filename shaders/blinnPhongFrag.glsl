#version 410

uniform vec3 CPOS;
uniform sampler2D tex;
uniform mat4 MVP, MODEL;

// TODO: store light data in go program
vec3 lightPos = vec3(0.0, 0.0, 0.0);
const vec3 Iamb = vec3(0.1, 0.1, 0.1)*1;
const vec3 Idif = vec3(0.1, 0.1, 0.1)*6;
const vec3 Ispec = vec3(0.1, 0.1, 0.1)*10;

in vec3 fragPos;
in vec3 fragNoraml;
in vec2 fragTexCoord;

out vec4 finalColor;

void main() {
  // TODO: Support multiple light sources
  vec3 L = normalize(lightPos - fragPos);
  vec3 N = normalize(fragNoraml);
  vec3 V = normalize(-fragPos);

  float lambertian = max(dot(N,L), 0.0);
  float specular = 0.0;

  if(lambertian > 0.0) {
    vec3 H = normalize(L + V);
    float specAngle = max(dot(H, N), 0.0);
    specular = pow(specAngle, 16.0);
  }
  float diffuse = max(dot(normalize(fragNoraml), normalize(lightPos)), 0.0);

  vec3 texture = texture(tex, fragTexCoord).rgb;
  finalColor = vec4(texture * (Iamb +
                    lambertian * Idif +
                    specular * Ispec ) ,1);
}
