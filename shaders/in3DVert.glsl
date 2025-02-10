#version 410 core

uniform mat4 MVP, MODEL, NormalMatrix;

in vec3 vert;
in vec2 vertTexCoord;
in vec3 vertNormal;
in vec3 vertTangent;

out vec3 fragPos;
out vec2 fragTexCoord;
out vec3 fragNormal;
out mat3 fragTBN;

void main(){
  vec4 fragPos4 = MODEL * vec4(vert, 1.0);
  fragTexCoord = vertTexCoord;
  fragPos = fragPos4.xyz / fragPos4.w;
  fragNormal = (NormalMatrix * vec4(vertNormal, 0.0)).xyz;

  vec3 N = normalize((NormalMatrix * vec4(vertNormal, 0.0)).xyz);
  vec3 T = normalize((NormalMatrix * vec4(vertTangent, 0.0)).xyz);
  T = normalize(T - dot(T, N) * N);
  vec3 B = cross(N, T);
  fragTBN = mat3(T, B, N);

  gl_Position = MVP * vec4(fragPos, 1.0);
}
