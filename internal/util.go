package internal

func Default(a, b string) string {
	if a != "" {
		return a
	}
	return b
}
