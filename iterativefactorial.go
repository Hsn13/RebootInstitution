package piscine

func IterativeFactorial(nb int) int {
	for nb < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}
