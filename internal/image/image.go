package image

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)


func GetImagePixels(path string) ([][]color.Color, error) {
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    img, _, err := image.Decode(f)
    if err != nil {
        return nil, err
    }

    boundsMin := img.Bounds().Min
    boundsMax := img.Bounds().Max
    pixels := make([][]color.Color, boundsMax.Y-boundsMin.Y)

    for y := boundsMin.Y; y < boundsMax.Y; y++ {  
        pixels[y] = make([]color.Color, boundsMax.X-boundsMin.X)
        for x := boundsMin.X; x < boundsMax.X; x++ {
            pixels[y][x] = img.At(x, y)
        }
    }

    return pixels, nil
}
