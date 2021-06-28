package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	lenMatrix, lenRow := n, n*2-1
	for i := 0; i < lenMatrix; i++ {
		for j := 0; j < lenRow; j++ {
			if j >= ((lenMatrix-1)-i) && (j <= (lenRow-1)-((lenMatrix-1)-i)) {
				fmt.Print("@")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
