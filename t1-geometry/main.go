package main

import (
	"fmt"
	"golang-learning/t1-geometry/geometry"
)

func main() {
	other := geometry.Point{X: 10, Y: 20}
	p := geometry.Point{X: 30, Y: 40}

	fmt.Println(p.DistanceTo(other))
}
