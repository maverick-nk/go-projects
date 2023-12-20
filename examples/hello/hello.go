package main

import (
	"fmt"
	"maps"
	"time"
)

// constants
const val = "something"

// function: passed by value
func add(a int, b int) int {
	return a + b
}

// function: passed by reference
func change(a *int) {
	*a = 20
}

// variadic functions
func sumAndCount(nums ...int) (int, int) {
	sum := 0
	count := 0
	for _, num := range nums {
		sum += num
		count += 1
	}
	return sum, count
}

// closures: iterator
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// Basic
	var helloVar = "Hello"
	fmt.Println(helloVar + "Hello, World!" + val)
	fmt.Println(val)
	a := time.Now()
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, World! " + val)
	}

	// Array and slices
	var arr []string
	fmt.Println(arr)

	// Maps
	m := make(map[string]float32)
	m["key"] = 1.0
	val, prs := m["k2"]
	fmt.Println(val)
	fmt.Println(prs)
	fmt.Println(m)

	map1 := map[string]int{"k1": 1, "k2": 3}
	map2 := map[string]int{"k1": 1, "k2": 1}
	fmt.Println(maps.Equal(map1, map2))

	// For loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// Range
	nums := [5]int{1, 2, 3, 4, 5}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	kvs := map[string]int{"k1": 1, "k2": 2}
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	// only keys
	for k := range kvs {
		fmt.Println(k)
	}
	for _, v := range kvs {
		fmt.Println(v)
	}

	// function invoke
	fmt.Println(add(1, 2))

	// variadic function and multiple return

	sum, count := sumAndCount(1, 2, 3, 4, 5)
	fmt.Println(sum, count)

	// With slice calling a variadic function
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s, c := sumAndCount(numbers...)

	fmt.Println(s, c)

	// Closure invocation: iterator example
	nextSeq := intSeq()

	fmt.Println(nextSeq()) // 1
	fmt.Println(nextSeq()) // 2
	fmt.Println(nextSeq()) // 3

	// function pass by reference
	pv := 1
	change(&pv)
	fmt.Println(pv) // 20
}
