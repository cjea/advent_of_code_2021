package part2

import (
	"fmt"
	"strconv"
	"strings"
)

type Direction uint8

const (
	FORWARD        Direction = iota
	DOWN                     = iota
	UP                       = iota
	NUM_DIRECTIONS           = iota
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type Op struct {
	Direction Direction
	Qty       int
}

func ParseOp(raw string) Op {
	parts := strings.SplitN(raw, " ", 2)
	if len(parts) != 2 {
		must(fmt.Errorf("Invalid operation: %s", raw))
	}
	qty, err := strconv.Atoi(parts[1])
	must(err)
	if NUM_DIRECTIONS != 3 {
		must(fmt.Errorf("exhaustive handling of directions"))
	}

	var dir Direction
	switch parts[0] {
	case "forward":
		dir = FORWARD
	case "up":
		dir = UP
	case "down":
		dir = DOWN
	default:
		must(fmt.Errorf("Invalid direction: %s", parts[0]))
	}

	return Op{
		Direction: dir,
		Qty:       qty,
	}
}

type Coords struct {
	Depth int
	Horiz int
	Aim   int
}

func (c *Coords) Apply(op Op) {
	switch op.Direction {
	case FORWARD:
		c.Horiz += op.Qty
		c.Depth += (c.Aim * op.Qty)
	case UP:
		c.Aim -= op.Qty
	case DOWN:
		c.Aim += op.Qty
	default:
		must(fmt.Errorf("Invalid direction: %s", op.Direction))
	}
}
