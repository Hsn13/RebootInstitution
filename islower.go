package piscine

func IsLower(s string) bool {
	for _, s := range s {
		if s < 97 || s > 122 {
			return false
		}
	}
	return true
}
