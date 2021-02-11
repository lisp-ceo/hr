package countingvalleys

func countValleys(_ int, p string) int {
	// Valleys is the number of valleys we've traversed
	valleys := 0

	// Levels tracks our altitude
	level := 0

	for _, char := range p {
		if char == 'U' {
			level += 1
		}
		if char == 'D' {
			if level == 0 {
				valleys += 1
			}
			level -= 1
		}
	}

	return valleys
}
