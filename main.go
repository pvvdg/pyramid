package main

import (
	"fmt"
)

func main() {
	symbol := "@"
	n := 0
	fmt.Scan(&n)
	var matrix [][]string
	var row []string

	for i := 0; i < n; i++ {
		for j := 0; j < n*2-1; j++ {
			row = append(row, " ")
		}
		matrix = append(matrix, row)
		row = []string{}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j >= ((len(matrix)-1)-i) && (j <= (len(matrix[i])-1)-((len(matrix)-1)-i)) {
				fmt.Print(symbol)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
