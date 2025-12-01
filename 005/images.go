package main

import (
	"fmt"
	"image"
	"image/color"
)

////////////////////////////////////////////////////
// 1) Basic image creation with image.RGBA
////////////////////////////////////////////////////
func basicImageDemo() {
	fmt.Println("---- Basic Image Demo ----")
	// Create a 100x100 RGBA image
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println("Bounds:", m.Bounds())

	// RGBA returns 4 values: r, g, b, a
	r, g, b, a := m.At(0, 0).RGBA()
	fmt.Printf("Top-left pixel RGBA: r=%v g=%v b=%v a=%v\n", r, g, b, a)

	// Optional: set a pixel to red
	m.Set(0, 0, color.RGBA{255, 0, 0, 255})
	r, g, b, a = m.At(0, 0).RGBA()
	fmt.Printf("After setting pixel: r=%v g=%v b=%v a=%v\n", r, g, b, a)
}

////////////////////////////////////////////////////
// 2) Custom Image Example (gradient)
////////////////////////////////////////////////////
type Image struct{}

// ColorModel returns the Image's color model
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the size of the image
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

// At returns the color of pixel at (x, y) with a simple gradient
func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func customImageDemo() {
	fmt.Println("\n---- Custom Image Demo ----")
	fmt.Println("This demo creates a gradient image but 'pic.ShowImage' is skipped")
	// Normally you would call pic.ShowImage(img) to visualize it
	// pic.ShowImage(Image{})
}

////////////////////////////////////////////////////
// MAIN
////////////////////////////////////////////////////
func main() {
	basicImageDemo()
	customImageDemo()
}
