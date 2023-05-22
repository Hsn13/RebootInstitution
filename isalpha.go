package piscine

func IsAlpha(s string) bool {
	for _, s := range s {
		if s > 64 && s < 91 || s > 96 && s < 123 || s > 47 && s < 58 || s == 32 {
			return true
		}
	}
	return false
}
