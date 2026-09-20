package geometry

import "math"

const minPolygonVertices = 3

var _ Figure = (*Polygon)(nil)
var _ Figure = (*Circle)(nil)

type PointInterface interface {
	DistanceTo(other Point) float64
}

type Figure interface {
	Area() float64
	Perimeter() float64
	Contains(point Point) bool
}

type Point struct {
	X, Y float64
}

func (p *Point) DistanceTo(other Point) float64 {
	return math.Sqrt((p.X-other.X)*(p.X-other.X) + (p.Y-other.Y)*(p.Y-other.Y))
}

type Polygon struct {
	Points []Point
}

func (pl Polygon) Area() float64 {
	n := len(pl.Points)
	if n < minPolygonVertices {
		return 0
	}

	sum := 0.0

	for i := range n {
		j := (i + 1) % n

		sum += pl.Points[i].X*pl.Points[j].Y -
			pl.Points[j].X*pl.Points[i].Y
	}

	return math.Abs(sum) / 2
}

func (pl Polygon) Perimeter() float64 {
	n := len(pl.Points)
	if n < minPolygonVertices {
		return 0
	}

	perimeter := 0.0

	for i := range n {
		j := (i + 1) % n

		dx := pl.Points[j].X - pl.Points[i].X
		dy := pl.Points[j].Y - pl.Points[i].Y

		perimeter += math.Hypot(dx, dy)
	}

	return perimeter
}

func (pl Polygon) Contains(point Point) bool {
	n := len(pl.Points)
	if n < minPolygonVertices {
		return false
	}

	inside := false

	for i := range n {
		j := (i + 1) % n

		first := pl.Points[i]
		second := pl.Points[j]

		if pointOnSegment(point, first, second) {
			return true
		}

		crossesHorizontalLine :=
			(first.Y > point.Y) != (second.Y > point.Y)

		if crossesHorizontalLine {
			intersectionX := first.X +
				(point.Y-first.Y)*(second.X-first.X)/
					(second.Y-first.Y)

			if point.X < intersectionX {
				inside = !inside
			}
		}
	}

	return inside
}

func pointOnSegment(point, first, second Point) bool {
	const epsilon = 1e-9

	crossProduct :=
		(point.X-first.X)*(second.Y-first.Y) -
			(point.Y-first.Y)*(second.X-first.X)

	if math.Abs(crossProduct) > epsilon {
		return false
	}

	dotProduct :=
		(point.X-first.X)*(point.X-second.X) +
			(point.Y-first.Y)*(point.Y-second.Y)

	return dotProduct <= epsilon
}

type Circle struct {
	Center Point
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Contains(point Point) bool {
	return (point.X-c.Center.X)*(point.X-c.Center.X)+(point.Y-c.Center.Y)*(point.Y-c.Center.Y) <= c.Radius*c.Radius
}
