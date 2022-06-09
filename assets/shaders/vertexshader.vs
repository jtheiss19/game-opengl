#version 410
layout (location = 0) in vec3 PosCoord;
layout (location = 1) in vec3 ColorCoord;
layout (location = 2) in vec2 TexCoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

out vec3 aColorCoord;
out vec2 aTexCoord;

void main() {
	gl_Position = projection * view * model * vec4(PosCoord, 1.0);
	aColorCoord = ColorCoord;
	aTexCoord = vec2(TexCoord.x, 1- TexCoord.y);
}
