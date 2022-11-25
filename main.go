package main

import (
	"fmt"
	"lympach/data"
	"lympach/ui"
)

func main() {
	fmt.Println("Lympach 0.1")

	data.Load()

	ui.BuildAndRun()
}
