package geometry_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang-learning/t1/geometry"
)

const (
	floatTolerance             = 1e-9
	emptyPolygonCase           = "empty polygon"
	fewerThanThreeVerticesCase = "fewer than three vertices"
)

func TestPointDistanceTo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		from geometry.Point
		to   geometry.Point
		want float64
	}{
		{
			name: "triangle 3-4-5",
			from: geometry.Point{X: 0, Y: 0},
			to:   geometry.Point{X: 3, Y: 4},
			want: 5,
		},
		{
			name: "same point",
			from: geometry.Point{X: 2.5, Y: -7},
			to:   geometry.Point{X: 2.5, Y: -7},
			want: 0,
		},
		{
			name: "negative coordinates",
			from: geometry.Point{X: -1, Y: -1},
			to:   geometry.Point{X: 2, Y: 3},
			want: 5,
		},
		{
			name: "fractional coordinates",
			from: geometry.Point{X: 0.5, Y: 1.5},
			to:   geometry.Point{X: 1.5, Y: 2.5},
			want: math.Sqrt(2),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := test.from.DistanceTo(test.to)

			assert.InDelta(t, test.want, got, floatTolerance)
		})
	}
}

//nolint:funlen // Keeping all area scenarios in one table makes them easier to compare.
func TestPolygonArea(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		polygon geometry.Polygon
		want    float64
	}{
		{
			name:    emptyPolygonCase,
			polygon: geometry.Polygon{Points: nil},
			want:    0,
		},
		{
			name: fewerThanThreeVerticesCase,
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
			}},
			want: 0,
		},
		{
			name: "right triangle",
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 4, Y: 0},
				{X: 0, Y: 3},
			}},
			want: 6,
		},
		{
			name: "rectangle",
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 4, Y: 0},
				{X: 4, Y: 3},
				{X: 0, Y: 3},
			}},
			want: 12,
		},
		{
			name: "clockwise vertices",
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 0, Y: 3},
				{X: 4, Y: 3},
				{X: 4, Y: 0},
			}},
			want: 12,
		},
		{
			name: "concave polygon",
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 4, Y: 0},
				{X: 4, Y: 1},
				{X: 1, Y: 1},
				{X: 1, Y: 4},
				{X: 0, Y: 4},
			}},
			want: 7,
		},
		{
			name: "collinear vertices",
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
				{X: 2, Y: 2},
			}},
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := test.polygon.Area()

			assert.InDelta(t, test.want, got, floatTolerance)
		})
	}
}

func TestPolygonPerimeter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		polygon geometry.Polygon
		want    float64
	}{
		{
			name:    emptyPolygonCase,
			polygon: geometry.Polygon{Points: nil},
			want:    0,
		},
		{
			name: fewerThanThreeVerticesCase,
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
			}},
			want: 0,
		},
		{
			name: "triangle",
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 4},
			}},
			want: 12,
		},
		{
			name: "rectangle",
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 4, Y: 0},
				{X: 4, Y: 3},
				{X: 0, Y: 3},
			}},
			want: 14,
		},
		{
			name: "collinear vertices",
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 3, Y: 0},
			}},
			want: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := test.polygon.Perimeter()

			assert.InDelta(t, test.want, got, floatTolerance)
		})
	}
}

//nolint:funlen // Keeping all containment scenarios in one table makes them easier to compare.
func TestPolygonContains(t *testing.T) {
	t.Parallel()

	square := geometry.Polygon{Points: []geometry.Point{
		{X: 0, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 4},
		{X: 0, Y: 4},
	}}
	concave := geometry.Polygon{Points: []geometry.Point{
		{X: 0, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 4},
		{X: 0, Y: 4},
	}}

	tests := []struct {
		name    string
		polygon geometry.Polygon
		point   geometry.Point
		want    bool
	}{
		{
			name:    emptyPolygonCase,
			polygon: geometry.Polygon{Points: nil},
			point:   geometry.Point{X: 0, Y: 0},
			want:    false,
		},
		{
			name: fewerThanThreeVerticesCase,
			polygon: geometry.Polygon{Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
			}},
			point: geometry.Point{X: 0.5, Y: 0.5},
			want:  false,
		},
		{
			name:    "inside convex polygon",
			polygon: square,
			point:   geometry.Point{X: 2, Y: 2},
			want:    true,
		},
		{
			name:    "outside convex polygon",
			polygon: square,
			point:   geometry.Point{X: 5, Y: 2},
			want:    false,
		},
		{
			name:    "on horizontal edge",
			polygon: square,
			point:   geometry.Point{X: 2, Y: 0},
			want:    true,
		},
		{
			name:    "on vertical edge",
			polygon: square,
			point:   geometry.Point{X: 0, Y: 2},
			want:    true,
		},
		{
			name:    "on vertex",
			polygon: square,
			point:   geometry.Point{X: 0, Y: 0},
			want:    true,
		},
		{
			name:    "inside concave polygon",
			polygon: concave,
			point:   geometry.Point{X: 0.5, Y: 3},
			want:    true,
		},
		{
			name:    "inside concave cutout",
			polygon: concave,
			point:   geometry.Point{X: 2, Y: 2},
			want:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := test.polygon.Contains(test.point)

			assert.Equal(t, test.want, got)
		})
	}
}

func TestCircleArea(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{
			name:   "zero radius",
			radius: 0,
			want:   0,
		},
		{
			name:   "unit circle",
			radius: 1,
			want:   math.Pi,
		},
		{
			name:   "fractional radius",
			radius: 2.5,
			want:   math.Pi * 2.5 * 2.5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			circle := geometry.Circle{
				Center: geometry.Point{X: 0, Y: 0},
				Radius: test.radius,
			}

			assert.InDelta(t, test.want, circle.Area(), floatTolerance)
		})
	}
}

func TestCirclePerimeter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{
			name:   "zero radius",
			radius: 0,
			want:   0,
		},
		{
			name:   "unit circle",
			radius: 1,
			want:   2 * math.Pi,
		},
		{
			name:   "fractional radius",
			radius: 2.5,
			want:   5 * math.Pi,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			circle := geometry.Circle{
				Center: geometry.Point{X: 0, Y: 0},
				Radius: test.radius,
			}

			assert.InDelta(t, test.want, circle.Perimeter(), floatTolerance)
		})
	}
}

func TestCircleContains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		circle geometry.Circle
		point  geometry.Point
		want   bool
	}{
		{
			name:   "center",
			circle: geometry.Circle{Center: geometry.Point{X: 1, Y: 2}, Radius: 3},
			point:  geometry.Point{X: 1, Y: 2},
			want:   true,
		},
		{
			name:   "inside",
			circle: geometry.Circle{Center: geometry.Point{X: 1, Y: 2}, Radius: 3},
			point:  geometry.Point{X: 2, Y: 3},
			want:   true,
		},
		{
			name:   "on boundary",
			circle: geometry.Circle{Center: geometry.Point{X: 1, Y: 2}, Radius: 3},
			point:  geometry.Point{X: 4, Y: 2},
			want:   true,
		},
		{
			name:   "outside",
			circle: geometry.Circle{Center: geometry.Point{X: 1, Y: 2}, Radius: 3},
			point:  geometry.Point{X: 4.1, Y: 2},
			want:   false,
		},
		{
			name:   "zero radius contains center",
			circle: geometry.Circle{Center: geometry.Point{X: -2, Y: 5}, Radius: 0},
			point:  geometry.Point{X: -2, Y: 5},
			want:   true,
		},
		{
			name:   "zero radius excludes other point",
			circle: geometry.Circle{Center: geometry.Point{X: -2, Y: 5}, Radius: 0},
			point:  geometry.Point{X: -2, Y: 5.1},
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := test.circle.Contains(test.point)

			assert.Equal(t, test.want, got)
		})
	}
}
