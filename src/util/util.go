package util

import "fmt"

type Dimension struct {
	height int
	width  int
	length int
}

func (d *Dimension) area() int {
	d.height = 15
	return d.length * d.width
}

func (d Dimension) volume() int {
	d.height = 55
	return d.height * d.length * d.width
}

func Calc() string {
	d := Dimension{10, 25, 36}
	return fmt.Sprintf("Calculation on: %v\nArea: %v, Volume: %v", d, d.area(), d.volume())
}
