#version 460
in vec3 aColorCoord;
in vec2 aTexCoord;

uniform sampler2D aTexture;

out vec4 frag_colour;

void main() {
	vec3 ColorMask = vec3(1, 1, 1);
	frag_colour = texture(aTexture, aTexCoord) * vec4(aColorCoord * ColorMask, 1);
}
