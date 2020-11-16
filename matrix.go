package matrixmath

import "fmt"

//Matrix struct to intiate matrix
type Matrix struct {
	Rows, Columns int
	Data          [][]int
}

//ShowMatrix function to print the matrix
func (m Matrix) ShowMatrix() {
	fmt.Println(m.Data)
}

//Multiply function multiplies a scalar to every element of the matrix
func (m *Matrix) Multiply(num int) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			m.Data[i][j] *= num
		}
	}
}

//MatAdd function adds two matrices and returns a new matrix
func MatAdd(m1, m2 Matrix) Matrix {
	temparr := Matrix{Rows: m1.Rows, Columns: m1.Columns, Data: m1.Data}
	if len(m1.Data) == len(m2.Data) && len(m1.Data[0]) == len(m2.Data[0]) {

		for i := 0; i < m1.Rows; i++ {
			for j := 0; j < m1.Columns; j++ {
				temparr.Data[i][j] = m1.Data[i][j] + m2.Data[i][j]
			}
		}
	} else {
		fmt.Println("matrices don't have common dimnesions")
	}
	return temparr
}

// func (m Matrix) getRows() int {
// 	const arr = const (m.rows)
// 	return arr
// }

//MatMultiply function multiplies two matrices (matrix product) and returns a matrix
func MatMultiply(m1, m2 Matrix) Matrix {
	arr := make([][]int, m1.Rows)
	for i := 0; i < m1.Rows; i++ {
		arr[i] = make([]int, m2.Columns)
	}
	if len(m1.Data[0]) == len(m2.Data) {

		for i := 0; i < m1.Rows; i++ {
			for j := 0; j < m2.Columns; j++ {
				for k := 0; k < m1.Columns; k++ {
					arr[i][j] += m1.Data[i][k] * m2.Data[k][j]
				}
			}
		}
	} else {
		fmt.Println("matrices don't have common dimnesions")
	}
	temparr := Matrix{Rows: m1.Rows, Columns: m2.Columns, Data: arr}
	return temparr
}
