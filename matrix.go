package pandemic

import "fmt"

type Pair struct {
	A int
	B int
}

func NewPair(a, b int) Pair {
	if a < b {
		return Pair{
			A: a,
			B: b,
		}
	}
	return Pair{
		A: b,
		B: a,
	}
}

type Matrix [][]int

func (m Matrix) Average() float64 {
	var sum, count int
	for _, row := range m {
		for _, e := range row {
			sum += e
			count++
		}
	}
	return float64(sum) / float64(count)
}

func (m Matrix) Max() int {
	var max int
	for _, row := range m {
		for _, e := range row {
			if e > max {
				max = e
			}
		}
	}
	return max
}

func (m Matrix) SortedPairs() [][]Pair {
	var max int
	pairs := make(map[Pair]int)
	for ri, row := range m {
		for ci, e := range row {
			p := NewPair(ri, ci)
			if _, ok := pairs[p]; !ok {
				pairs[p] = e
				if e > max {
					max = e
				}
			}
		}
	}
	sortedPairs := make([][]Pair, max+1)
	for p, e := range pairs {
		sortedPairs[e] = append(sortedPairs[e], p)
	}
	return sortedPairs
}

func (m Matrix) String() string {
	var out string
	for _, elements := range m {
		for ci, element := range elements {
			if ci != 0 {
				out += " "
			}
			out += fmt.Sprintf("%2.d", element)
			if ci == len(elements)-1 {
				out += "\n"
			}
		}
	}
	return out
}
