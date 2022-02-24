package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	valc := *******c
	*******c = ***a
	vald := ****d
	****d = valc
	valb := *b
	*b = vald
	***a = valb
}

func Decript(a ***int, b *int, c *******int, d ****int) {
	vala := ***a
	***a = *******c
	valb := *b
	*b = vala
	vald := ****d
	****d = valb
	*******c = vald
}
