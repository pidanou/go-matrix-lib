package matrix

import "fmt"

func rotateMatrix(times int, Matrix [][]string) [][]string {

	var rows = len(Matrix)
	var cols = len(Matrix[0])

	rotatedMatrix := make([][]string, 0)

	tmp := make([]string, 0)

	for col := 0; col < cols; col++ {
		for row := rows - 1; row >= 0; row-- {
			fmt.Println(Matrix[row][col])
			tmp = append(tmp, Matrix[row][col])
		}

		rotatedMatrix = append(rotatedMatrix, tmp)
		tmp = nil
	}

	return rotatedMatrix

}
