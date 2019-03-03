package main

import (
	"fmt"
)

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int{
	if (n == 0) {
		return  1
	}

	return  n * fact(n - 1)
}

func main() {
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)

	d, e := vals()

	fmt.Println(d)
	fmt.Println(e)

	sum(1, 2)

	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}

	sum(nums...)

	fmt.Println("==========")

	nexInt := intSeq()

	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())

	newInts := intSeq()

	fmt.Println(newInts())

	fmt.Println("==========")

	fmt.Println(fact(7))
}
