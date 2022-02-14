package main

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 0
	} else if nb > 0 {
		i := 1
		a := i
		for ; i <= nb; i++ {
			a = a * i
		}
		return a
	} else if nb < 0 {
		i := -1
		a := -i
		for ; i >= nb; i-- {
			a = a * -1
		}
		return a
	}
	return 1
}
