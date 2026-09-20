package cli

import (
	"errors"
	"flag"
	"fmt"
	"golang-learning/t1/geometry"
)

func distance(args []string) error {
	flags := flag.NewFlagSet("distance", flag.ContinueOnError)

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

	fmt.Printf("Расстояние: %.2f\n", distance)

	return nil
}
