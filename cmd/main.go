package main

import (
	"fmt"
	"os"

	"github.com/Kaspetti/spe/internal/image"
)


func main() {
    pixels, err := image.GetImagePixels("./norway.jpg")
    if err != nil {
        panic(err)
    }

    f, err := os.Create("processedImage.ppm")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    f.Write([]byte(fmt.Sprintf("P3\n%d %d\n255\n", len(pixels[0]), len(pixels))))

    for _, row := range pixels {
        for _, p := range row {
            r, g, b, _ := p.RGBA()
            f.Write([]byte(fmt.Sprintf("%d %d %d\n", r>>8, g>>8, b>>8)))
        }
    }
}
