package cli

import (
	"errors"
	"flag"
	"fmt"
	"golang-learning/t1/internal/geometry"
	"io"
	"strconv"
	"strings"
)

const (
	shapePolygon = "polygon"
	shapeCircle  = "circle"
)

func parsePoint(value string) (geometry.Point, error) {
	xString, yString, found := strings.Cut(value, ",")
	if !found {
		return geometry.Point{}, fmt.Errorf(
			"неверный формат точки %q: ожидается X,Y",
			value,
		)
	}

	x, err := strconv.ParseFloat(strings.TrimSpace(xString), 64)
	if err != nil {
		return geometry.Point{}, fmt.Errorf(
			"некорректная координата X %q: %w",
			xString,
			err,
		)
	}
	y, err := strconv.ParseFloat(strings.TrimSpace(yString), 64)
	if err != nil {
		return geometry.Point{}, fmt.Errorf(
			"некорректная координата Y %q: %w",
			yString,
			err,
		)
	}

	return geometry.Point{X: x, Y: y}, nil
}

func parsePoints(points *[]geometry.Point) func(string) error {
	return func(value string) error {
		point, err := parsePoint(value)
		if err != nil {
			return err
		}

		*points = append(*points, point)

		return nil
	}
}

func parseCenter(center *geometry.Point, centerSet *bool) func(string) error {
	return func(value string) error {
		point, err := parsePoint(value)
		if err != nil {
			return err
		}

		*center = point
		*centerSet = true

		return nil
	}
}

func parseRadius(radius *float64, radiusSet *bool) func(string) error {
	return func(value string) error {
		if value == "" {
			return errors.New("не задан радиус окружности")
		}

		result, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil {
			return fmt.Errorf("некорректный радиус окружности %f: %w", result, err)
		}

		if result <= 0 {
			return fmt.Errorf(
				"радиус должен быть положительным, получено: %v",
				result,
			)
		}

		*radius = result
		*radiusSet = true

		return nil
	}
}

type CalculationResult struct {
	shape  string
	points []geometry.Point
	circle geometry.Circle
}

func parseFigureCalculationArgs(args []string, operation string) (CalculationResult, error) {
	flags := flag.NewFlagSet(operation, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	shape := flags.String("shape", "", "форма фигуры")

	var radius float64
	var points []geometry.Point
	var center geometry.Point
	var radiusSet, centerSet bool

	flags.Func("point", "точка в формате X,Y", parsePoints(&points))
	flags.Func("center", "центр окружности в формате X,Y", parseCenter(&center, &centerSet))
	flags.Func("radius", "радиус окружности", parseRadius(&radius, &radiusSet))

	err := flags.Parse(args)
	if err != nil {
		return CalculationResult{}, fmt.Errorf("не удалось прочитать аргументы: %w", err)
	}

	switch *shape {
	case "":
		return CalculationResult{}, errors.New("не задан аргумент --shape")
	case shapePolygon:
		if len(points) < minPolygonVertices {
			return CalculationResult{}, errors.New("для фигуры polygon требуется передать минимум 3 точки в формате --point=X,Y")
		}

		var result CalculationResult
		result.shape = "polygon"
		result.points = points

		return result, nil
	case shapeCircle:
		if !centerSet {
			return CalculationResult{}, errors.New(
				"необходимо задать центр окружности в формате --center=X,Y",
			)
		}

		if !radiusSet {
			return CalculationResult{}, errors.New(
				"необходимо задать положительный радиус в формате --radius=R",
			)
		}

		var result CalculationResult
		result.shape = "circle"
		result.circle = geometry.Circle{Radius: radius, Center: center}

		return result, nil
	default:
		return CalculationResult{}, fmt.Errorf("некорретная форма %s - введите circle или polygon", *shape)
	}
}
