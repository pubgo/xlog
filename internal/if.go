package internal

// If exported
func If(check bool, a, b interface{}) interface{} {
	if check {
		return a
	}
	return b
}
