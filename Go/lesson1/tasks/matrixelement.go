package tasks

func MatrixElementSum(matrix [3][4]int) int {
	sum := 0
	for i := 0; i < 1; i++ {
		for j := 0; j < 4; j++ {
			sum += matrix[i][j]
		}
	}
	for i := 1; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if matrix[i-1][j] != 0 {
				sum += matrix[i][j]
			}
		}
	}
	return sum
}