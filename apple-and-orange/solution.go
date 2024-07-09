package main

// s: start point house
// t: end point house
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var appleCount, orangeCount int32

	for _, apple := range apples {
		if a+apple >= s && a+apple <= t {
			appleCount++
		}
	}

	for _, orange := range oranges {
		if b+orange >= s && b+orange <= t {
			orangeCount++
		}
	}

	println(appleCount)
	println(orangeCount)
}
