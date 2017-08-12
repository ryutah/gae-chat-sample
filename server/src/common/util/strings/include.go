package strings

func Include(s string, tgt []string) bool {
	for _, t := range tgt {
		if t == s {
			return true
		}
	}
	return false
}
