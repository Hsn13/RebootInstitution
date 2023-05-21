package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	if nb > 0 {
		for i := 1; i <= nb; i++ {
			result *= i
		}
	}
	return result
}
