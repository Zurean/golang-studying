package cli

import (
	"errors"
	"flag"
	"fmt"
	"golang-learning/t1/geometry"
	"io"
)

func distance(args []string, output io.Writer) error {
	flags := flag.NewFlagSet("distance", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var points []geometry.Point

	flags.Func("point", "точка в формате X,Y", parsePoints(&points))

	err := flags.Parse(args)
	if err != nil {
		return fmt.Errorf("не удалось прочитать аргументы: %w", err)
	}

	if len(points) != 2 {
		return errors.New("для distance необходимо указать ровно две точки")
	}

	distance := points[0].DistanceTo(points[1])

	_, err = fmt.Fprintf(output, "Расстояние: %.2f\n", distance)
	if err != nil {
		return fmt.Errorf("не удалось вывести расстояние: %w", err)
	}

	return nil
}
