package main

func main() {

}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 0, n

	ans := 1
	for ans != 0 {
		m := l + (r-l)/2
		ans = guess(m)
		if ans == 0 {
			return m
		} else if ans == 1 {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}

func guess(n int) int {
	return 0
}
