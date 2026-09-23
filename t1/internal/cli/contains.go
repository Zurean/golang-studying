package cli

import (
	"errors"
	"flag"
	"fmt"
	"golang-learning/t1/internal/geometry"
	"io"
)

const minPolygonVertices = 3

type ContainParams struct {
	radius               float64
	points               []geometry.Point
	center               geometry.Point
	radiusSet, centerSet bool
}

func contains(args []string, output io.Writer) error {
	flags := flag.NewFlagSet("contains", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	shape := flags.String("shape", "", "форма фигуры")

	var params ContainParams
	var result bool

	flags.Func("point", "точка в формате X,Y", parsePoints(&params.points))

	flags.Func("center", "центр окружности в формате X,Y", parseCenter(&params.center, &params.centerSet))

	flags.Func("radius", "радиус окружности", parseRadius(&params.radius, &params.radiusSet))

	err := flags.Parse(args)
	if err != nil {
		return fmt.Errorf("не удалось прочитать аргументы: %w", err)
	}

	switch *shape {
	case "":
		return errors.New("не задан аргумент --shape")
	case shapePolygon:
		result, err = polygon(params.points)
	case shapeCircle:
		result, err = circle(params)
	default:
		return fmt.Errorf("некорретная форма фигуры %s - введите circle или polygon", *shape)
	}

	if err != nil {
		return fmt.Errorf("ошибка определения вхождения точки в фигуру: %w", err)
	}

	message := "Точка не входит в фигуру"
	if result {
		message = "Точка входит в фигуру"
	}

	_, err = fmt.Fprintln(output, message)
	if err != nil {
		return fmt.Errorf("не удалось вывести результат проверки: %w", err)
	}

	return nil
}

func polygon(points []geometry.Point) (bool, error) {
	if len(points) < minPolygonVertices+1 {
		return false, errors.New("для определения признака включения точки в многоугольник " +
			"нужно передать минимум 4 точки в формате --point=X,Y")
	}

	return (geometry.Polygon{Points: points[1:]}).Contains(points[0]), nil
}

func circle(params ContainParams) (bool, error) {
	if len(params.points) < 1 {
		return false, errors.New("необходимо задать точку в формате --point=X,Y")
	}

	if !params.centerSet {
		return false, errors.New(
			"необходимо задать центр окружности в формате --center=X,Y",
		)
	}

	if !params.radiusSet {
		return false, errors.New(
			"необходимо задать положительный радиус в формате --radius=R",
		)
	}

	return (geometry.Circle{Center: params.center, Radius: params.radius}).Contains(params.points[0]), nil
}
