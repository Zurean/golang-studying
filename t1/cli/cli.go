package cli

import (
	"errors"
	"fmt"
)

func Run(args []string) error {
	if len(args) == 0 {
		return errors.New("не задана команда")
	}

	switch args[0] {
		case "distance":
			return distance(args[1:])
		case "perimeter":
			return perimeter(args[1:])
		case "area":
			return area(args[1:])
		case "contains":
			return contains(args[1:])
		default:
			return fmt.Errorf("неизвестная команда: %s", args[0])
	}
}


