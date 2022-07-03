package main

import (
	"fmt"
)

func main() {
	var n, m, start int
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%d\n", &m)
	fmt.Scanf("%d\n", &start)
	var states = make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j != m-1 {
				fmt.Scanf("%d", &states[i*m+j])
			} else {
				fmt.Scanf("%d\n", &states[i*m+j])
			}
		}
	}
	var outputs = make([]string, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j != m-1 {
				fmt.Scanf("%s", &outputs[i*m+j])
			} else {
				fmt.Scanf("%s\n", &outputs[i*m+j])
			}
		}
	}

	fmt.Println("digraph {")
	fmt.Println("    rankdir = LR")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf(`    %d -> %d [label = "%c(%s)"]`, i, states[i*m+j], 97+j, outputs[i*m+j])
			fmt.Print("\n")
		}
	}
	fmt.Println("}")
}
