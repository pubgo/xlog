package internal

func Default(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func Set(data ...string) map[string]struct{} {
	var set = make(map[string]struct{})
	for i := range data {
		set[data[i]] = struct{}{}
	}
	return set
}
