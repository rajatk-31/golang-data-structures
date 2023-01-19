package main

func numWaterBottles(numBottles int, numExchange int) int {
	sum := numBottles
	for numBottles >= numExchange {
		totalNew := numBottles / numExchange
		sum += totalNew
		numBottles = totalNew + (numBottles % numExchange)
	}
	return sum
}
