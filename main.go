package main

import (
	"fmt"
	"reflect"
	"math/big"
)

type Printable interface {
	Print()
}

func print(x Printable) {
	x.Print()
}

func printGeneric[T Printable](x T) {
	x.Print()
}

type mytype struct {}

func (t *mytype) Print() {
	fmt.Println("noot")
}

func reverse[T any](in []T) []T {
    reversed := make([]T, len(in))
    for i, item := range in {
    	reversed[len(in)-i-1] = item
    }
    return reversed
}

func reverseNotGeneric(in []interface{}) []interface{} {
    reversed := make([]interface{}, len(in))
    for i, item := range in {
    	reversed[len(in)-i-1] = item
    }
    return reversed
}

func main() {	
	// reversing []int using generic func
	ints := []int{0, 1, 2}
	fmt.Println("before:", ints)
	reversed := reverse(ints)
	fmt.Println("after:", reversed)
	fmt.Println("type:", reflect.TypeOf(reversed))

	// reversing []int using non-generic func - requires converting each
	// int into interface{}
	interfaceInts := make([]interface{}, len(ints))
	for i, item := range ints {
		interfaceInts[i] = item
	}
	fmt.Println("before:", interfaceInts)
	reversedInterfaceInts := reverseNotGeneric(interfaceInts)
	fmt.Println("after:", reversedInterfaceInts)
	fmt.Println("type:", reflect.TypeOf(reversedInterfaceInts))

	// reversing []*big.Int using generic func
	bigints := []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)}
	fmt.Println("before:", bigints)
	bigintsReversed := reverse(bigints)
	fmt.Println("after:", bigintsReversed)

	// reversing []string using generic func
	strs := []string{"a", "b", "c"}
	fmt.Println("before:", strs)
	strsReversed := reverse(strs)
	fmt.Println("after:", strsReversed)

	// using interfaces to print
	t := new(mytype)
	print(t)

	// using generics to print
	printGeneric(t)
}