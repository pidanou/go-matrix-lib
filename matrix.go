package matrix

import (
	"errors"
	"fmt"
)

func RotateMatrixClockwise[K comparable](Matrix [][]K) [][]K {

	var rows = len(Matrix)
	var cols = len(Matrix[0])

	rotatedMatrix := make([][]K, 0)

	tmp := make([]K, 0)

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

// faire la fonction CClockwise

func VerticalSplitMatrix[K any](index int, matrix [][]K) ([][]K, [][]K, error) {

	if index > len(matrix) || index < 0 {
		return nil, nil, errors.New("index out range")
	}

	splitMatrixLeft := make([][]K, 0)
	splitMatrixRight := make([][]K, 0)

	for i := 0; i < len(matrix); i++ {
		splitMatrixLeft = append(splitMatrixLeft, matrix[i][:index])
		splitMatrixRight = append(splitMatrixRight, matrix[i][index:])
	}

	return splitMatrixLeft, splitMatrixRight, nil

}

func HorizontalSplitMatrix[K any](index int, matrix [][]K) ([][]K, [][]K, error) {

	if index > len(matrix) || index < 0 {
		return nil, nil, errors.New("index out range")
	}

	splitMatrixTop := matrix[:index]
	splitMatrixBottom := matrix[index:]

	return splitMatrixTop, splitMatrixBottom, nil

}

func VerticalMirrorMatrix[K any](matrix [][]K) [][]K {

	for item := 0; item < len(matrix); item++ {
		for i, j := 0, len(matrix[item])-1; i < j; i, j = i+1, j-1 {
			matrix[item][i], matrix[item][j] = matrix[item][j], matrix[item][i]

		}
	}

	return matrix
}

func HorizontalMirrorMatrix[K any](matrix [][]K) [][]K {

	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	return matrix
}
