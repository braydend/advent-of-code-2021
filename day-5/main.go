package main

import (
	"fmt"
	"github.com/braydend/advent-of-code/day-5/io"
	"github.com/braydend/advent-of-code/day-5/vent"
	"log"
)

func main() {
	vents, err := io.ParseVentsFromFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	ventlocations := vent.FindOverlappingVents(vents, true)
	fmt.Println(vent.CountOverlappingVents(ventlocations, 2))
}
