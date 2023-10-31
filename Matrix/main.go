package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please select a size (# of rows/columns) for the square matrices that will be multiplied")
	readSize, err := reader.ReadBytes('\n')

	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	size, err := strconv.Atoi(string(readSize[:len(readSize)-1]))
	if err != nil || size > 10 {
		fmt.Println("Invalid size input. Input must be a valid integer that is less than or equal to 10")
		return
	}

	fmt.Println("Provide the values for matrix 1")
	m1 := &SquareMatrix{size, make([][]int, 0)}
	createSquareMatrix(m1, reader)

	fmt.Println("Provide the values for matrix 2")
	m2 := &SquareMatrix{size, make([][]int, 0)}
	createSquareMatrix(m2, reader)

	result, err := m1.multiply(m2)
	if err != nil {
		fmt.Println("Something went wrong with the matrix multiplcation")
		return
	}

	fmt.Println("Resulting Matrix:")
	for i := 0; i < size; i++ {
		fmt.Println(result.Rows[i])
	}
}

func createSquareMatrix(m *SquareMatrix, r *bufio.Reader) (*SquareMatrix, error) {
	numValues := m.Size * m.Size
	index := 0
	values := make([]int, numValues, numValues)
	fmt.Printf("Please provide %d arguments\n", numValues)
	for i := 0; i < m.Size; i++ {
		for j := 0; j < m.Size; j++ {
			fmt.Printf("Enter value for r%d,c%d: ", i, j)
			readValue, err := r.ReadBytes('\n')
			if err != nil {
				fmt.Println("Error encountered reading input")
				return &SquareMatrix{}, errors.New("Error reading input")
			}
			value, err := strconv.Atoi(string(readValue[:len(readValue)-1]))
			if err != nil {
				fmt.Println("Please provide a valid integer as an input")
				return &SquareMatrix{}, errors.New("Invalid Integer Input")
			}
			values[index] = value
			index++
		}
	}
	m.instantiate(values...)
	return m, nil
}
