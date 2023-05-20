#version 330


//shader de flou
//pas opti du tout, avec une complexité 0(n²)
//prend comme variable uniforme partagée, "size", le rayon de la matrice de convolution
//réduit également la luminosité en fonction de cette "size"

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;


// Output fragment color
out vec4 finalColor;

// NOTE: Add here your custom variables
uniform float size;

// NOTE: Render size values must be passed from code
const float renderWidth = 1920;
const float renderHeight = 1080;



void main()
{
    // Texel color fetching from texture sampler
    vec3 texelColor = vec3(0.0, 0.0, 0.0);
    int radius = int(size);
    float weightSum = 0.0;
    float maxDistance = size*sqrt(2);
    float weight;

    for (int i = -radius; i <= radius; i++) 
    {
        for (int j = -radius; j <= radius; j++) 
        {
            weight = maxDistance-sqrt(int(i*i+j*j));
            weightSum += weight;
            texelColor += texture(texture0, fragTexCoord + vec2(i/renderWidth, j/renderHeight)).rgb*weight;
        }
    }

    

    finalColor = vec4(texelColor/weightSum, 1.0)*(1-size/30.0);

    if (radius == 0) {
        finalColor = texture(texture0, fragTexCoord);
    } 
    
}