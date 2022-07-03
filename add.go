package main

import (
	"fmt"
)

func add(a, b []int32, p int) []int32 {
	pp := int32(p)
	if len(b) < len(a) {
		a, b = b, a
	}
	result := make([]int32, len(b))
	var toAdd int32
	var i = 0
	for ; i < len(a); i++ {
		if a[i] >= pp || b[i] >= pp {
			panic("wrong input")
		}
		result[i] = (a[i] + b[i] + toAdd) % pp
		toAdd = (a[i] + b[i] + toAdd) / pp
	}
	for ; i < len(b); i++ {
		if b[i] >= pp {
			panic("wrong input")
		}
		result[i] = (b[i] + toAdd) % pp
		toAdd = (b[i] + toAdd) / pp
	}
	for toAdd != 0 {
		result = append(result, toAdd%pp)
		toAdd = toAdd / pp
	}

	return result
}

func reverse(src []int32) []int32 {
	res := make([]int32, len(src))
	for i := range src {
		res[i] = src[len(src)-1-i]
	}
	return res
}
func main() {
	first := []int32{9, 7, 6}
	second := []int32{9, 9, 9}
	base := 10
	fmt.Print(reverse(add(first, second, base)))
}
