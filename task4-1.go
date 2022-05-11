package main

import (
	"fmt"
	"math"
)

type Cuboid struct {
	width  float64
	length float64
	height float64
}

func (c Cuboid) Volume() float64 {
	return c.length * c.width * c.height
}

func (c Cuboid) SurfaceArea() float64 {
	area := 2 * c.length * c.width
	area += 2 * c.length * c.height
	area += 2 * c.height * c.width
	return area
}

type Sphere struct {
	radius float64
}

func (s Sphere) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(s.radius, 3)
}

func (s Sphere) SurfaceArea() float64 {
	return 2 * math.Pi * math.Pow(s.radius, 2)
}

type Solid interface {
	Volume() float64
	SurfaceArea() float64
}

func PrintInfo(s Solid) {
	fmt.Println(s)
	fmt.Println("Volume:", s.Volume())
	fmt.Println("Surface Area:", s.SurfaceArea())
}

func main() {
	c := Cuboid{length: 2.5, width: 5.0, height: 10.5}
	PrintInfo(c)
	s := Sphere{radius: 2.0}
	PrintInfo(s)
}