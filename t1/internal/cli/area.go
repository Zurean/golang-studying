package cli

import (
	"errors"
	"fmt"
	"golang-learning/t1/internal/geometry"
	"io"
)

func area(args []string, output io.Writer) error {
	result, err := parseFigureCalculationArgs(args, "area")
	if err != nil {
		return fmt.Errorf("ошибка парсинга: %w", err)
	}

	switch result.shape {
	case "":
		return errors.New("не задан аргумент --shape")
	case shapePolygon:
		_, err = fmt.Fprintf(output, "Площадь: %.2f\n",
			(geometry.Polygon{Points: result.points}).Area(),
		)
	case shapeCircle:
		_, err = fmt.Fprintf(output, "Площадь: %.2f\n",
			result.circle.Area(),
		)

	default:
		return fmt.Errorf("некорретная форма %s - введите circle или polygon", result.shape)
	}

	if err != nil {
		return fmt.Errorf("не удалось вывести площадь: %w", err)
	}

	return nil
}
