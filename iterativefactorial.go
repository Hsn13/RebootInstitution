package piscine

func IterativeFactorial(nb int) int {

	if nb <= 0 || nb > 24 {
		return 0
	}
	a := nb - 1
	for i := a; i >= 1; i-- {
		nb *= int(i)
	}
	return nb
}
