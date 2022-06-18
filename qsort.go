package main

import "fmt"

func less(i, j int) bool {
	return (myarray[i] < myarray[j])
}

func swap(i, j int) {
	var num int
	num = myarray[i]
	myarray[i] = myarray[j]
	myarray[j] = num
}

func partition(less func(i, j int) bool, swap func(i, j int), low int, high int) int {
	var i, j int
	i = low
	j = low
	//на слайдах в презентации был while (2 презентация с первого семестра, слайд 71), но в Go while нет, поэтому тут for :(
	for t := 0; t < 1; t++ {
		if less(j, high) {
			swap(i, j)
			i++
		}
		j++
		if j < high {
			t = -2
		} else {
			t = 2
		}
	}
	swap(i, high)
	return i
}

func qsort(len int, less func(i, j int) bool, swap func(i, j int)) {
	qsortrec(0, len-1, less, swap)
}

func qsortrec(low int, high int, less func(i, j int) bool, swap func(i, j int)) {
	var q int
	q = 0
	if low < high {
		q = partition(less, swap, low, high)
		qsortrec(low, q-1, less, swap)
		qsortrec(q+1, high, less, swap)
	}
}

var len = 10
var myarray [10]int

func main() {

	myarray[0] = 3
	myarray[1] = 4
	myarray[2] = 6
	myarray[3] = 7
	myarray[4] = 8
	myarray[5] = 1
	myarray[6] = 2
	myarray[7] = 9
	myarray[8] = 0
	myarray[9] = 10

	qsort(len, less, swap)
	for i := 0; i < 10; i++ {
		fmt.Println(myarray[i])
	}
}
