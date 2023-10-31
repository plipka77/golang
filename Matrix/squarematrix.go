package main

import (
	"errors"
	"fmt"
)

type SquareMatrix struct {
	Size int
	Rows [][]int
}

func (m *SquareMatrix) instantiate(values ...int) (bool, error) {
	s := m.Size
	rows := make([][]int, s, s)
	if len(values) != (s * s) {
		fmt.Printf("Please provide %d values...", s*s)
		return false, errors.New("Invalid input length")
	}

	index := 0
	for i := 0; i < s; i++ {
		row := make([]int, s, s)
		rows[i] = row
		for j := 0; j < s; j++ {
			rows[i][j] = values[index]
			index++
		}
	}
	m.Rows = rows
	return true, nil
}

func (m1 *SquareMatrix) multiply(m2 *SquareMatrix) (*SquareMatrix, error) {
	s1 := m1.Size
	s2 := m2.Size
	if s1 != s2 {
		fmt.Println("Please provide a square matrix with a compatible size")
		return new(SquareMatrix), errors.New("Incompatible input size")
	}
	value := 0
	rows := make([][]int, s1, s1)
	for i := 0; i < s1; i++ {
		row := make([]int, s1, s1)
		rows[i] = row
		for j := 0; j < s1; j++ {
			value = 0
			for k := 0; k < s1; k++ {
				value += m1.Rows[i][k] * m2.Rows[k][j]
			}
			rows[i][j] = value
		}
	}
	return &SquareMatrix{s1, rows}, nil
}
