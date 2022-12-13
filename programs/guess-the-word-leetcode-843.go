package main

func main() {

}

// func findSecretWord(wordlist []string, master *Master) {

//     // map word to a score
//     wordMap := make(map[string]int)
//     for i := range wordlist {
//         wordMap[wordlist[i]] = 0
//     }

//     // just use first word as first guess
//     guess := wordlist[0]
//     for i := 0; i < 10; i++ {
//         score := master.Guess(guess)
//         if score == -1 {
//             guess = getMaxKey(wordMap) // if there were no matches to guess, then use the less best guess
//         }
//         guess = pruneMap(wordMap, guess, score) // re-update map scores and return best guess with highest score
//     }
// }

// func getMaxKey(m map[string]int) string {
//     ans, max := "", -1
//     for k, v := range m {
//         if v > max {
//             max = v
//             ans = k
//         }
//     }
//     delete(m, ans)
//     return ans
// }

// func pruneMap(m map[string]int, word string, score int) string {
//     maxScore, bestGuess := -1, ""
//     for k := range m {
//         for i := range word {
//             if k[i] == word[i] {
//                 m[k] += score
//                 if score == 0 {
//                     delete(m, k)
//                 }
//             }
//         }
//         if m[k] > maxScore {
//             maxScore = m[k]
//             bestGuess = k
//         }
//     }
//     delete(m, bestGuess)
//     return bestGuess
// }

/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */
func findSecretWord(wordlist []string, master *Master) {
	H := make([][]int, len(wordlist))
	for i := 0; i < len(wordlist); i++ {
		H[i] = make([]int, len(wordlist))
	}
	for i := 0; i < len(wordlist); i++ {
		for j := i; j < len(wordlist); j++ {
			if i == j {
				H[i][j] = 6
			} else {
				match := 0
				for k := 0; k < 6; k++ {
					if wordlist[i][k] == wordlist[j][k] {
						match++
					}
				}
				H[i][j] = match
				H[j][i] = match
			}
		}
	}

	// makes a slice with all indexes
	// eg [0,0] => [0,1]
	possible := make([]int, len(wordlist))
	for i := 0; i < len(wordlist); i++ {
		possible[i] = i
	}
	for len(possible) > 1 {
		// returns a single index "guess"
		guess := solve(possible, H)

		// passes "guess" word to the API
		val := master.Guess(wordlist[guess])

		// if its 6, you found the word
		if val == 6 {
			return
		}

		// If the guess has the same number of matches as the
		// possible index
		// append it to a new list temp
		temp := []int{}
		for i := 0; i < len(possible); i++ {
			if H[guess][possible[i]] == val {
				temp = append(temp, possible[i])
			}
		}
		possible = temp
	}
	master.Guess(wordlist[possible[0]])
}

// Returns a row index number
// the index represents the minimum value
// of the row, whose count is mapped
func solve(possible []int, H [][]int) int {
	minv := len(H)
	minvIdx := -1

	// range over 2d matrix
	// pass the dd matrix, and the current row index
	// pass the "possible" slice of indexes

	// returns a min row Index
	// of a map of all rows,
	// mapping word match occurrances to a count of occurrances
	for i, _ := range H {
		//
		maxl := eachRowMax(possible, i, H)
		if maxl < minv {
			minvIdx = i
			minv = maxl
		}
	}
	return minvIdx
}

// returns an int
func eachRowMax(possible []int, n int, H [][]int) int {
	m := map[int]int{}
	ret := 1

	// range over "possible" slice
	// try to find a maped value for H{i,j}
	// i is a row of H
	// j is a value of "possible" ("possible" index value of H)
	// if you find a mapped value + 1
	// set max of ret

	// otherwise map "possible" for H{i,j} as 1

	// if the original matrix row is [[3 6 2 2]]
	// the ret would be map[2:2, 3:1, 6:1]
	// 2
	for _, i := range possible {

		if val, ok := m[H[n][i]]; ok {
			m[H[n][i]] = val + 1
			ret = max(ret, m[H[n][i]])
		} else {
			m[H[n][i]] = 1
		}

	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
