package main

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// If the kangaroo starting behind has a slower or equal speed, they will never meet
	if (x1 < x2 && v1 <= v2) || (x2 < x1 && v2 <= v1) {
		return "NO"
	}

	// Check if they can meet
	if (x2-x1)%(v1-v2) == 0 {
		return "YES"
	}

	return "NO"
}
