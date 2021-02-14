package leftrotation

// rotLeft rotates the head of array a, d times.
func rotLeft(a []int32, d int32) []int32 {
	for i := d; i > 0; i-- {
		head, tail := a[0], a[1:]
		a = append(tail, head)
	}
	return a
}
