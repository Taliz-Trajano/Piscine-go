package piscine

func AppendRange(min, max int) []int {
<<<<<<< HEAD
	if min >= max {
		var tab []int = nil
		return tab
	} else {
		tab := []int{}

		for i := min; i < max; i++ {
			tab = append(tab, i)
		}
		return tab
	}
}
=======
        if min >= max {
                var tab []int = nil
                return tab
        }else{
                tab := []int{}
		
                for i:= min;i < max;i++ {
                        tab = append(tab,i)
                }
                return tab
        }
}
>>>>>>> d0fdce4a882dfa96c1d2f9ad9692d94809dc69a6
