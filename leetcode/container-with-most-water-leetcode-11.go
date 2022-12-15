package main

import "math"

func main() {

}

func maxArea(height []int) int {
	maxArea := float64(0)
	initial := 0
	last := len(height) - 1
	for initial < last {
		minHeight := math.Min(float64(height[initial]), float64(height[last]))
		area := minHeight * float64((last - initial))
		maxArea = math.Max(area, maxArea)
		if height[initial] > height[last] {
			last--
		} else {
			initial++
		}
	}
	return int(maxArea)
}
