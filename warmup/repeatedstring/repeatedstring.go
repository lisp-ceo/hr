package repeatedstring

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	stringLength := int64(len(s))
	repetitions := n / stringLength
	remainder := n % stringLength

	var count int64
	for n := range s {
		if s[n] == 'a' {
			count += repetitions
			if int64(n) < remainder {
				count += 1
			}
		}
	}

	return count
}
