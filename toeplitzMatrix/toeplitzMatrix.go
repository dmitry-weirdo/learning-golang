package main

func isToeplitzMatrix(matrix [][]int) bool {
	rows := len(matrix)
	columns := len(matrix[0])

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}

	return true
}

func main() {

}
