package main

import "fmt"

type Number interface {
	int64 | float64
}

// https://go.dev/doc/tutorial/generics
func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloat[string, int64](ints),
		SumIntsOrFloat[string, float64](floats),
	)

	// remove type arguments
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloat(ints),
		SumIntsOrFloat(floats),
	)

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
}

func SumIntsOrFloat[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
