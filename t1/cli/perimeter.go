package cli

import (
	"errors"
	"fmt"
	"golang-learning/t1/geometry"
	"io"
)

func perimeter(args []string, output io.Writer) error {
	result, err := parseFigureCalculationArgs(args, "perimeter")
	if err != nil {
		return fmt.Errorf("ошибка парсинга: %w", err)
	}

	switch result.shape {
	case "":
		return errors.New("не задан аргумент --shape")
	case shapePolygon:
		_, err = fmt.Fprintf(output, "Периметр: %.2f\n",
			(geometry.Polygon{Points: result.points}).Perimeter(),
		)
	case shapeCircle:
		_, err = fmt.Fprintf(output, "Периметр: %.2f\n",
			result.circle.Perimeter(),
		)

	default:
		return fmt.Errorf("некорретная форма %s - введите circle или polygon", result.shape)
	}

	if err != nil {
		return fmt.Errorf("не удалось вывести периметр: %w", err)
	}

	return nil
}
