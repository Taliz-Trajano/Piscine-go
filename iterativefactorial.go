package piscine

func IterativeFactorial(nb int) int {
	if nb < -999 || nb > 999 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb > 0 {
		i := 1
		a := i
		for ; i <= nb; i++ {
			a = a * i
		}
		return a
	} else {
		i := -1
		a := -i
		for ; i >= nb; i-- {
			a = a * -1
		}
		return a
	}
}
