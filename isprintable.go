package piscine

func IsPrintable(s string) bool {
	if s == "" {
		return false
	}
	for _, s := range s {
		if !(s >= 32) {
			return false
		}
	}
	return true
}
