package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	must(err)
	return i
}

func CountIncreases(buf *bufio.Scanner) int {
  var total, i, prev, first int
	total = 0
  buf.Scan()
  first = mustInt(buf.Text())

  for i, prev = first, first; buf.Scan(); prev = i {
   i = mustInt(buf.Text())
   if i > prev {
     total++
   }
 }

	return total
}

func CountIncreasesThrees(buf *bufio.Scanner) int {
	total := 0
	slot0, slot1, slot2, slot3 := 0, 0, 0, 0
	buf.Scan()
	slot0 = mustInt(buf.Text())

	buf.Scan()
	slot1 = mustInt(buf.Text())

	buf.Scan()
	slot2 = mustInt(buf.Text())

	buf.Scan()
	slot3 = mustInt(buf.Text())

	for ; buf.Scan(); slot3 = mustInt(buf.Text()) {
		if slot3 > slot0 {
			total++
		}
		slot0, slot1, slot2 = slot1, slot2, slot3
	}

	if slot3 > slot0 {
		total++
	}

	return total
}

func main() {
	f, err := os.Open("input")
	must(err)
	buf := bufio.NewScanner(f)
	fmt.Printf("Total increases: %d\n", CountIncreases(buf))

	f, err = os.Open("input")
	must(err)
	buf = bufio.NewScanner(f)
	fmt.Printf("Total increases by threes: %d\n", CountIncreasesThrees(buf))
}
