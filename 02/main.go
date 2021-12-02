package main

import (
	"bufio"
	"directions/pkg/part1"
	"directions/pkg/part2"
	"fmt"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("input")
	must(err)
	scanner := bufio.NewScanner(f)
	coords := &part1.Coords{Depth: 0, Horiz: 0}
	for scanner.Scan() {
		raw := scanner.Text()
		op := part1.ParseOp(raw)
		coords.Apply(op)
	}
  fmt.Printf("Part 1 Coords: %+v\n", coords)
  fmt.Printf("Part 1 Product: %d\n", coords.Depth*coords.Horiz)

	f, err = os.Open("input")
	must(err)
	scanner = bufio.NewScanner(f)
  coords2 := &part2.Coords{Depth: 0, Horiz: 0, Aim: 0}
	for scanner.Scan() {
		raw := scanner.Text()
		op := part2.ParseOp(raw)
		coords2.Apply(op)
	}
  fmt.Printf("Part 2 Coords: %+v\n", coords2)
  fmt.Printf("Part 2 Product: %d\n", coords2.Depth*coords.Horiz)
}
