package newyearchaos

import (
	"fmt"
	"io"
	"os"
)

var (
	stdout io.Writer = os.Stdout
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	allBribes := make([]int32, len(q)+1)
	bs(q, allBribes)

	var bribes int32
	for n := range allBribes {
		if allBribes[n] > 2{
			_, _ = fmt.Fprint(stdout, "Too chaotic\n")
			return
		}
		bribes += allBribes[n]
	}

	_, _ = fmt.Fprintf(stdout, "%d\n", bribes)
}

// bs implements a bubblesort to track bribes.
func bs(q []int32, bribes []int32) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(q); i++ {
			if q[i-1] > q[i] {
				swap(q,i-1, i, bribes)
				swapped = true
			}
		}
	}
}

// qs implements a simple quicksort.
// XXX: cannot be used to sort for bribes as we lose track of the bribe chain.
func qs(q []int32, lo int, hi int, bribes []int32) {
	if lo < hi {
		p := partition(q, lo, hi, bribes)
		if p > 0 {
			qs(q, lo, p-1, bribes)
		}
		qs(q, p+1, hi, bribes)
	}
}

func partition(q []int32, lo int, hi int, bribes []int32) int {
	pivot := q[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if q[j] < pivot {
			swap(q, i, j, bribes)
			i += 1
		}
	}
	swap(q, i, hi, bribes)

	return i
}

// swap swaps the ith and jth elements of q and records bribes.
func swap(q []int32, i int, j int, bribes []int32) {
	if i == j {
		return
	}
	// the bribe we've found
	justinBriber := q[i]
	// record the bribe
	bribes[justinBriber] += 1
	// swap places
	tmp := q[i]
	q[i] = q[j]
	q[j] = tmp
}
