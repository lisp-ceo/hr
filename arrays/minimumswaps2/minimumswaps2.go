package minimumswaps2

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	return cycles(arr)
}

// Finding cycles. The minimum number of swaps will equal the number of cycles in out of position elements.
func cycles(q []int32) int32{
	var cycles int32
	visited := make([]bool, len(q))

	for i := int32(0); i < int32(len(q)); i++ {
		if !visited[i] {
			visited[i] = true

			if q[i] != i+1 {
				nxt := q[i]

				for !visited[nxt-1] {
					visited[nxt-1] = true
					nxt = q[nxt-1]
					cycles++
				}
			}
		}
	}

	return cycles
}
