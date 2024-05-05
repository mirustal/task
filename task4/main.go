package main

import "fmt"

func printMultiplicationTable(num int) {
	i := 1
	fmt.Printf(" ")
	for i <= num {
		fmt.Printf("%4d", i)
		i++
	}
	fmt.Println()
	i = 1
	for i <= num { 
		j := 1
		fmt.Printf("%d", i) 
		for j <= num { 
			fmt.Printf("%4d", i*j)
			j++
		}
		fmt.Println()
		i++
	}
}

func main() {
	printMultiplicationTable(5)
}

/*
    1   2   3   4   5
1   1   2   3   4   5
2   2   4   6   8  10
3   3   6   9  12  15
4   4   8  12  16  20
5   5  10  15  20  25
*/