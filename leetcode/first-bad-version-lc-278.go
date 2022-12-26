package main

func main() {

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}
	left, right := 1, n
	for left != right {
		mid := ((left + right) / 2)
		check := isBadVersion(mid)
		if check {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return right

}

// Just for example
func isBadVersion(n int) bool {
	return true
}
