package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	tmpA := ***a
	tmpB := *b
	tmpC := *******c
	tmpD := ****d

	***a = tmpB     // b → a
	*b = tmpD       // d → b
	*******c = tmpA // a → c
	****d = tmpC    // c → d
}
