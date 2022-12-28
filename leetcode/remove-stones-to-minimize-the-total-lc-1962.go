package main

import "container/heap"

func main() {

}

type Mine []int

func (m Mine) Len() int {
	return len(m)
}

func (m Mine) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m Mine) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *Mine) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *Mine) Pop() interface{} {
	last := len(*m) - 1
	toDelete := (*m)[last]

	*m = (*m)[:last]

	return toDelete
}

func minStoneSum(piles []int, k int) int {
	data := Mine(piles)
	heap.Init(&data)
	for k > 0 {
		data[0] -= (data[0] / 2)

		heap.Fix(&data, 0)

		k--
	}

	s := 0

	for _, val := range piles {
		s += val
	}

	return s
}
