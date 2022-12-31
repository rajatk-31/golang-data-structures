package main

func main() {

}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color != newColor {
		changeColor(image, sr, sc, &color, &newColor)
	}
	return image
}

func changeColor(image [][]int, x, y int, color, newColor *int) {
	if image[x][y] == *color {
		image[x][y] = *newColor
		if x >= 1 {
			changeColor(image, x-1, y, color, newColor)
		}
		if y >= 1 {
			changeColor(image, x, y-1, color, newColor)
		}
		if x+1 < len(image) {
			changeColor(image, x+1, y, color, newColor)
		}
		if y+1 < len(image[0]) {
			changeColor(image, x, y+1, color, newColor)
		}
	}
}
