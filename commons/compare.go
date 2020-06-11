package commons

// Compare check if ami are same
func Compare(a, b []string) []string {
	for i := len(a) - 1; i >= 0; i-- {
		for _, v := range b {
			if a[i] == v {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
	return a
}
