package main

func main() {

}

func rotate(matrix [][]int) {
	// Change 90def
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse array
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][i], matrix[j][len(matrix)-1-i] = matrix[j][len(matrix)-1-i], matrix[j][i]
		}
	}
}
