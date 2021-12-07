package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"submarine-binary/pkg/solution"
)

const FILENAME = "input"

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open(FILENAME)
	must(err)
	reader := bufio.NewReader(f)
	body, err := ioutil.ReadAll(reader)
	must(err)
	rows := strings.Split(string(body), "\n")

	res := solution.Parse(rows)
	sol := solution.Part1(res)
	fmt.Printf("Part 1 total: %v\n", sol)

	oxygen := Oxygen()
	carbon := Carbon()
	fmt.Printf("Part 2 total: %v\n", solution.ParseBinary(oxygen)*solution.ParseBinary(carbon))

}

func Carbon() string {
	f, err := os.Open(FILENAME)
	must(err)
	reader := bufio.NewReader(f)
	body, err := ioutil.ReadAll(reader)
	must(err)
	rows := strings.Split(string(body), "\n")

	res := solution.Parse(rows)

	for i := 0; len(res.Rows) > 1; i++ {
		var bit byte
		if float64(res.Slots[i]) < float64(len(res.Rows))/2.0 {
			bit = '1'
		} else {
			bit = '0'
		}
		filtered := filterRows(res.Rows, i, bit)
		res = solution.Parse(filtered)
	}
	return res.Rows[0]
}

func Oxygen() string {
	f, err := os.Open(FILENAME)
	must(err)
	reader := bufio.NewReader(f)
	body, err := ioutil.ReadAll(reader)
	must(err)
	rows := strings.Split(string(body), "\n")

	res := solution.Parse(rows)

	i := 0
	for len(res.Rows) > 1 {
		var bit byte
		if float64(res.Slots[i]) >= float64(len(res.Rows))/2.0 {
			bit = '1'
		} else {
			bit = '0'
		}
		filtered := filterRows(res.Rows, i, bit)
		res = solution.Parse(filtered)
		i++
	}

	return res.Rows[0]
}

func filterRows(rows []string, i int, bit byte) []string {
	ret := []string{}
	for _, s := range rows {
		if s[i] == bit {
			ret = append(ret, s)
		}
	}
	return ret
}
