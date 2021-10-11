package utils

//EmptyString checks whether a string is empty or not
func EmptyString(v string) string {
	if len(v) == 0 {
		v = "-"
	}

	return v
}
