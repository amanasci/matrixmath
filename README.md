# matrixmath

A *not-so-efficient, simple, and fun* Matrix Math library that I am trying to create while learning Go. 

The final goal of this library is to help me in creating a toy neural network or a perceptron so I am customizing it as per my needs.

There are many, much more efficient libraries available for use. But, I'll use this, cause why not? 

Learning is fun. And applying things as you learn is even better.

> Feel free to use it and help in adding more functions.

## Example: 


```Go
import "github.com/amanasci/matrixmath"

func main() {
    array := [][]int{{2, 3, 1}, {4, 5, 0}, {3, 4, 2}}
    
    array2 := [][]int{{1, 2}, {0, 5}, {2, 3}}
    
    mat := matrixmath.Matrix{Rows: 3, Columns: 3, Data: array}
    
    mat2 := matrixmath.Matrix{Rows: 3, Columns: 2, Data: array2}
    
    mat3 := matrixmath.MatMultiply(mat, mat2)
    
	fmt.Println(mat3.Data)
}

```
