/*
   While prototype means something different in software development, here it does not mean shitty version of code you will write
   Prototype pattern provides a refernce or base for object creation, which can be used to initiate object with a template state.
   Here, state of object consists of parameters and method an object contains.
   Rather than just cloning, you can also use dymanic parameters to keep count of some sort or recursive value. But better to not use it in interview.
   In nutshell, this is a template loader for classes with added feature of abstraction.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	loader := PluginLoader{}
	loader.RegisterPlugin("Vignette", &VignetteFilter{0.5, 0.5, 0})
	originalImage := GetBlankImage(512, 512, 0.1, 0.25, 0.8)
	fmt.Println("Original Image Pixel: " + originalImage.pixels[10][10].ToString())
	vignetteFilter := loader.GetFilter("Vignette")
	vignetteFilter.ApplyFilter(originalImage)
	fmt.Println("Filtered Image Pixel: " + originalImage.pixels[10][10].ToString())
}

type ImageFilter interface {
	ApplyFilter(img Image)
	Copy() ImageFilter
}

type VignetteFilter struct {
	darkIntensity float64
	width float64
	filterCount int
}

func (v *VignetteFilter) ApplyFilter(img Image) {
	centerX := float64(img.width) / 2.0
	centerY := float64(img.height) / 2.0
	maxDistance := math.Sqrt(centerX*centerX + centerY*centerY)
	startThreshold := maxDistance * v.width

	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			dx := float64(x) - centerX
			dy := float64(y) - centerY
			distance := math.Sqrt(dx*dx + dy*dy)
			if distance > startThreshold {
				factor := (distance - startThreshold) / (maxDistance - startThreshold)
				if factor > 1.0 {
					factor = 1.0
				}
				multiplier := 1.0 - (factor * v.darkIntensity)
				pixel := &img.pixels[y][x]
				pixel.R *= multiplier
				pixel.G *= multiplier
				pixel.B *= multiplier
			}
		}
	}
}

func (v *VignetteFilter) Copy() ImageFilter {
	return &VignetteFilter{
		darkIntensity: v.darkIntensity,
		width: v.width,
		filterCount: v.filterCount + 1,
	}
}

type Image struct {
	pixels [][]Pixel
	width int
	height int
}

func GetBlankImage(width int, height int, r float64, g float64, b float64) Image {
	pixels := make([][]Pixel, height)
	for i := range pixels {
        pixels[i] = make([]Pixel, width)
        for j := range pixels[i] {
            pixels[i][j] = Pixel{
                R: r,
                G: g,
                B: b,
            }
        }
    }
	img := Image{
		width: width,
		height: height,
		pixels: pixels,
	}
	return img
}

type Pixel struct {
	R float64
	G float64
	B float64
}

func (p *Pixel) ToString() string {
	return fmt.Sprintf("R%f G%f B%f", p.R, p.G, p.B)
}

type PluginLoader struct {
	prototypes map[string]ImageFilter
}

func (p *PluginLoader) RegisterPlugin(name string, prototype ImageFilter) {
	if p.prototypes == nil {
		p.prototypes = map[string]ImageFilter{}
	}
	p.prototypes[name] = prototype
}

func (p *PluginLoader) GetFilter(name string) ImageFilter {
	if proto, ok := p.prototypes[name]; ok {
		return proto.Copy()
	}
	return nil
}