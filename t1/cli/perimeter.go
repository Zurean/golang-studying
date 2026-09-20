package cli

import (
	"errors"
	"fmt"
	"golang-learning/t1/geometry"
)

func perimeter(args []string) error {
	result, err := parseFigureCalculationArgs(args, "perimeter")
	if err != nil {
		return fmt.Errorf("ошибка парсинга: %w", err)
	}

	switch result.shape {
	case "":
		return errors.New("не задан аргумент --shape")
	case shapePolygon:
		fmt.Printf("Периметр: %.2f\n",
			(geometry.Polygon{Points: result.points}).Perimeter(),
		)
	case shapeCircle:
		fmt.Printf("Периметр: %.2f\n",
			result.circle.Perimeter(),
		)

	default:
		return fmt.Errorf("некорретная форма %s - введите circle или polygon", result.shape)
	}

	return nil
}
