package main

func main() {

}

func distributeCandies(candies int, num_people int) []int {
	arr := []int{}
	i := 0
	last := 0
	for candies > 0 {
		if i == num_people {
			i = 0
		}
		lenA := len(arr)
		if candies <= last {
			last = candies
			candies = 0
		} else {
			last++
			candies -= last
		}
		if lenA > i {
			arr[i] += last
		} else {
			arr = append(arr, last)
		}
		i++

	}
	lenArray := len(arr)
	if lenArray < num_people {
		remain := num_people - lenArray
		for remain > 0 {
			arr = append(arr, 0)
			remain--
		}
	}
	return arr
}
