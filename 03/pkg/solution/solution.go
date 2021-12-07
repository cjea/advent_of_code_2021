package solution

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type Result struct {
	Rows      []string
	Slots     [12]int
	TotalRows int
}

func Parse(ss []string) Result {
	var rows = []string{}
	var slots [12]int
	var total int

	for _, s := range ss {
		if len(s) == 0 {
			continue
		}
		for i, bit := range s {
			slots[i] += int(bit - '0')
		}
		rows = append(rows, s)
		total++
	}

	return Result{
		Rows:      rows,
		Slots:     slots,
		TotalRows: total,
	}
}

func Part1(res Result) int64 {
	var gamma, epsilon int64
	threshold := res.TotalRows / 2
	for pwr, cardinality := range res.Slots {
		val := int64(1 << (11 - pwr))
		if cardinality >= threshold {
			gamma += val
		} else {
			epsilon += val
		}
	}
	return gamma * epsilon
}

func ParseBinary(bits string) int64 {
	var res int64
	basePwr := len(bits) - 1
	for pwr, b := range bits {
		val := int64(1 << (basePwr - pwr))
		if b == '1' {
			res += val
		}
	}
	return res
}

func Part2(res Result) int64 {
	var gamma, epsilon int64
	threshold := res.TotalRows / 2
	for pwr, cardinality := range res.Slots {
		val := int64(1 << (11 - pwr))
		if cardinality >= threshold {
			gamma += val
		} else {
			epsilon += val
		}
	}
	return gamma * epsilon
}
