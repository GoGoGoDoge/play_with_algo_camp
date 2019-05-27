package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	for i := 0; i <= 50; i++ {
		j := 50 - i
		if 2*i+4*j == 120 {
			fmt.Printf("Chicken #%d, Rabbit #%d\n", i, j)
			return
		}
	}
	fmt.Println("No solution is found")
}
