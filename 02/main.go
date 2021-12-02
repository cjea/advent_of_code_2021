package main

import (
	"bufio"
	"fmt"
	"os"
  "strconv"
	"strings"
)

type Direction uint8

const (
	FORWARD Direction = iota
	DOWN              = iota
	UP                = iota
  NUM_DIRECTIONS    = iota
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
    Qty: qty,
  }
}

type Coords struct {
  Depth int
  Horiz int
}

func (c *Coords) Apply(op Op) {
  switch op.Direction {
  case FORWARD:
    c.Horiz += op.Qty
  case UP:
    c.Depth -= op.Qty
  case DOWN:
    c.Depth += op.Qty
  default:
    must(fmt.Errorf("Invalid direction: %s", op.Direction))
  }
}

func main() {
	f, err := os.Open("input")
	must(err)
	scanner := bufio.NewScanner(f)
  coords := &Coords{Depth: 0, Horiz: 0}
	for scanner.Scan() {
		raw := scanner.Text()
    op := ParseOp(raw)
    coords.Apply(op) 
	}
  fmt.Printf("%+v\n", coords)
  fmt.Printf("%d\n", coords.Depth * coords. Horiz)
}
