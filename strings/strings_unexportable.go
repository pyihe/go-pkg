package strings

func isUpper(s string) bool {
	for i := range s {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	for i := range s {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			return false
		}
	}
	return true
}
