#version 410
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec3 ColorCoord;
layout (location = 2) in vec2 TexCoord;
layout (location = 3) in mat4 instanceMatrix;

out vec2 aTexCoord;
out vec3 aColorCoord;

uniform mat4 projection;
uniform mat4 view;

void main()
{
    gl_Position = projection * view * instanceMatrix * vec4(aPos, 1.0); 
    aTexCoord = TexCoord;
    aColorCoord = ColorCoord;
}