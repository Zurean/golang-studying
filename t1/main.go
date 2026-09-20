package main

import (
	"fmt"
	"golang-learning/t1/cli"
	"os"
)

func main() {
	err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
}
