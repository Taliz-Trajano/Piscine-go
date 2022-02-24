package piscine

func RecursiveFactorial(x int) int {
	if x == 0 {
		return 1
	} else if x > 9999 || x < 0 {
		return 0
	} else {
		return x * RecursiveFactorial(x-1)
	}
}
