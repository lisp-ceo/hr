package salesbymatch

func matchingSocks(socks []int) int {
	pairs := 0

	got := make(map[int]interface{})

	for _, sock := range socks {
		_, ok := got[sock]
		if ok {
			delete(got, sock)
			pairs += 1
		} else {
			got[sock] = nil
		}
	}

	return pairs
}
