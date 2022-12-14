package main

import "fmt"

func main() {

	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := 0
	last := 0
	seconds := 0
	for _, val := range timeSeries {
		fmt.Println(val, seconds, last, total)
		if seconds > 0 {
			if val > seconds {

				total = total + last
			} else {
				left := seconds - val + 1
				fmt.Println("left--", left)
				total = total + last - left

			}
			last = duration
			seconds = val + duration - 1

		} else {
			total = total + last
			seconds = val + duration - 1
			last = duration
		}

	}
	return total + last

	/*
		 var ret int

	    for i := 0; i < len(timeSeries)-1; i++ {
	        x := timeSeries[i+1] - timeSeries[i]
			if duration > x {
	            ret += x
	        } else {
	            ret += duration
	        }
	    }

	    ret += duration
	    return ret
	*/
}
