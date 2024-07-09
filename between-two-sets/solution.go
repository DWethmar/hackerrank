package main

import "fmt"

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */
func getTotalX(a []int32, b []int32) int32 {
	// var considerations []int32
	min := a[0]
	for _, l := range a {
		if l > min {
			min = l
		}
	}

	max := b[0]
	for _, l := range b {
		if l < max || max == 0 {
			max = l
		}
	}

	fmt.Print(min, max)

	for i := min; i <= max; i++ {
		// if i is a factor of a and a is a factor of i
		// if i is a factor of b and b is a factor of i
		// considerations = append(considerations, i)
	}

	return 0
}
