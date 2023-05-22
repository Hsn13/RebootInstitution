package piscine

func IsAlpha(s string) bool {
	if s == "" {
		return true
	}
	for _, s := range s {
		if !((s >= 48 && s <= 57) || (s >= 65 && s <= 90) || (s >= 97 && s <= 122)) {
			return false
		}
	}
	return true
}
