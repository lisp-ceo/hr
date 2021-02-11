package jumpingonclouds

// minJumps calculates the minimum number of jumps required to traverse the clouds where 0 is a cloud we can travel on
// and 1 is a cloud we cannot.
func minJumps(clouds []int32) int32 {
	numClouds := int32(len(clouds))
	max2s := numClouds / 2
	paths := allPaths(max2s, numClouds)
	validPaths := filterInvalid(paths, clouds)

	return minPathLength(validPaths)
}

// minPathLength returns the smallest number of jumps required in the list of valid paths.
func minPathLength(paths [][]int32) int32 {
	min := int32(1 << 31 - 1)
	for _, path := range paths {
		pathLength := len(path)
		asInt32 := int32(pathLength)
		if asInt32 < min {
			min = asInt32
		}
	}
	return min
}

// allPaths returns all combinations of 1 and 2 space jumps that can be used to travel numSpaces.
func allPaths(max2SpaceJumps int32, numSpaces int32) [][]int32 {
	paths := [][]int32{}
	for i := max2SpaceJumps; i >= 0; i-- {
		paths = append(paths, genPaths(i, numSpaces)...)
	}

	return paths

}

// genPaths generates all possible combinations of 1 and 2 space jumps.
func genPaths(twoSpaceJumps int32, spaces int32) [][]int32 {
	paths := [][]int32{}
	// base case - return single jumps for each remaining segment. Or, if we can't fit all the two space jumps into the
	// spaces, there will be no valid combinations.
	if twoSpaceJumps == 0 || spaces < twoSpaceJumps*2 {
		path := []int32{}
		for i := int32(0); i < spaces; i++ {
			path = append(path, 1)
		}
		paths = append(paths, path)
		return paths
	}
	// recursive case - insert a 2 space jump where we can and ask the recursive call to fit the rest in.
	for i := int32(0); i < spaces-1; i++ {
		path := []int32{}

		// prefill with 1's
		for j := int32(0); j < i; j++ {
			path = append(path, 1)
		}

		// allocate the 2 space jump
		path = append(path, 2)

		// remaining spaces we need to fill in the recusive call
		remaining := spaces - count(path)

		// allocate the rest of the 2 space jumps
		combos := genPaths(twoSpaceJumps-1, remaining)

		// expand each combination, prefixed with the path we've calculated so far.
		for _, combo := range combos {
			paths = append(paths, append(path, combo...))
		}
	}

	return paths

}

// count returns the number of jumps in a path.
func count(path []int32) int32 {
	sum := int32(0)

	for n := range path {
		sum += path[n]
	}
	return sum
}

// filterInvalid removes paths that fall into thunderheads.
func filterInvalid(paths [][]int32, clouds []int32) [][]int32 {
	filtered := [][]int32{}
Next:
	for _, path := range paths {
		position := int32(0)
		for i := 0; i < len(path); i++ {
			position += path[i]
			// We've hit a thunderhead. Try the next one.
			if clouds[position] == 1 {
				continue Next
			}
		}
		filtered = append(filtered, path)
	}

	return filtered

}
