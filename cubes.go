package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type PossibleAnswer struct {
	bases  []float64
	values []float64
	count  int
}

func compute(maxPermutations int) *PossibleAnswer {
	m := make(map[string]*PossibleAnswer)
	for i := float64(1); ; i++ {
		num := math.Pow(i, 3)
		s := strconv.FormatFloat(num, 'f', -1, 64)
		x := strings.Split(s, "")
		sort.Sort(sort.Reverse(sort.StringSlice(x)))
		key := strings.Join(x, "")
		pa, ok := m[key]
		if !ok {
			pa = &PossibleAnswer{bases: make([]float64, maxPermutations), values: make([]float64, maxPermutations), count: 0}
			m[key] = pa
		}
		pa.bases[pa.count] = i
		pa.values[pa.count] = num
		pa.count++

		if pa.count == maxPermutations {
			return pa
		}
	}
}

func main() {
	pa := compute(5)
	for j := 0; j < pa.count; j++ {
		fmt.Printf("%d^3 = %d\n", int64(pa.bases[j]), int64(pa.values[j]))
	}
}
