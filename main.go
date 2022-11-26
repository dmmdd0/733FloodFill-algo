// sourse
// https://github.com/deepshig/leetcode-solutions/blob/master/easy/array/flood_fill/flood_fill.go
package main

import (
	"fmt"
)

func floodFillRecursion(image [][]int, sr, sc, oldColor, newColor int) [][]int {
	image[sr][sc] = newColor

	if sr > 0 && image[sr-1][sc] == oldColor {
		image = floodFillRecursion(image, sr-1, sc, oldColor, newColor)
	}

	if sr < len(image)-1 && image[sr+1][sc] == oldColor {
		image = floodFillRecursion(image, sr+1, sc, oldColor, newColor)
	}

	if sc > 0 && image[sr][sc-1] == oldColor {
		image = floodFillRecursion(image, sr, sc-1, oldColor, newColor)
	}

	if sc < len(image[0])-1 && image[sr][sc+1] == oldColor {
		image = floodFillRecursion(image, sr, sc+1, oldColor, newColor)
	}

	return image
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	return floodFillRecursion(image, sr, sc, oldColor, newColor)
}

func main() {
	fmt.Println("Solution 1 : ", floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	fmt.Println("Solution 2 : ", floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1))
}
