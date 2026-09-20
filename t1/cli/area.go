package cli

import (
	"errors"
	"fmt"
	"golang-learning/t1/geometry"
)

func area(args []string) error {
	result, err := parseFigureCalculationArgs(args, "area")
	if err != nil {
		return fmt.Errorf("ошибка парсинга: %w", err)
	}

	switch result.shape {
	case "":
		return errors.New("не задан аргумент --shape")
	case shapePolygon:

		fmt.Printf("Площадь: %.2f\n",
			(geometry.Polygon{Points: result.points}).Area(),
		)
	case shapeCircle:
		fmt.Printf("Площадь: %.2f\n",
			result.circle.Area(),
		)

	default:
		return fmt.Errorf("некорретная форма %s - введите circle или polygon", result.shape)
	}

	return nil
}
