package cli

import (
	"errors"
	"fmt"
	"io"
)

func Run(args []string, output io.Writer) error {
	if len(args) == 0 {
		return errors.New("не задана команда")
	}

	switch args[0] {
	case "distance":
		return distance(args[1:], output)
	case "perimeter":
		return perimeter(args[1:], output)
	case "area":
		return area(args[1:], output)
	case "contains":
		return contains(args[1:], output)
	default:
		return fmt.Errorf("неизвестная команда: %s", args[0])
	}
}
