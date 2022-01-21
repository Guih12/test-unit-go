package testeavancado

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Width  float64
	Height float64
}

type Circulo struct {
	raio float64
}

func writeArea(f Forma) {
	fmt.Println(f.Area())
}

func (r Retangulo) Area() float64 {
	return r.Height * r.Width
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
