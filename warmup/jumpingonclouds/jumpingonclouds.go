package jumpingonclouds

// jumpingOnClouds calculates the minimum number of jumps required to traverse the clouds where 0 is a cloud we can
// travel on and 1 is a cloud we cannot.
func jumpingOnClouds(clouds []int32) int32 {
	// First jump is assumed to be to the 0th cloud so offset by one as that is guaranteed to be a good cloud.
	numJumpsToGoal := int32(len(clouds) - 1)
	max2SpaceJumps := numJumpsToGoal / 2
	validPaths := validPaths(max2SpaceJumps, numJumpsToGoal, clouds)
	if len(validPaths) == 0 {
		panic("no valid paths generated")
	}

	return minPathLength(validPaths)
}

// minPathLength returns the smallest number of jumps required in the list of valid paths.
func minPathLength(paths [][]int32) int32 {
	min := int32(1<<31 - 1)
	for _, path := range paths {
		pathLength := int32(len(path))
		if pathLength < min {
			min = pathLength
		}
	}
	return min
}

// validPaths returns all combinations of 1 and 2 space jumps that can be used to travel numSpaces.
func validPaths(max2SpaceJumps int32, numSpaces int32, clouds []int32) [][]int32 {
	paths := [][]int32{}
	for i := max2SpaceJumps; i >= 0; i-- {
		paths = append(paths, genPathsWithTermination(i, numSpaces, clouds)...)
	}
	return paths
}

// genPathsWithTermination does not return paths that are known to be invalid.
func genPathsWithTermination(twoSpaceJumps int32, spaces int32, clouds []int32) [][]int32 {
	paths := [][]int32{}
	// base case - return single jumps for each remaining segment. Or, if we can't fit all the two space jumps into the
	// spaces, there will be no valid combinations.
	if twoSpaceJumps == 0 || spaces < twoSpaceJumps*2 {
		path := []int32{}
		for i := int32(0); i < spaces; i++ {
			path = append(path, 1)
		}
		// check if path is valid before adding it to results
		if isSafePath(path, clouds) {
			paths = append(paths, path)
		}
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

		// check if path is valid before recursing for more results
		if isSafePath(path, clouds) {
			travelled := count(path)
			distanceToGoal := spaces
			remaining := distanceToGoal - travelled

			// allocate the rest of the 2 space jumps
			combos := genPathsWithTermination(twoSpaceJumps-1, remaining, clouds[int(travelled):])

			// expand each combination, prefixed with the path we've calculated so far.
			for _, combo := range combos {
				paths = append(paths, append(path, combo...))
			}
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

// isSafePath traverses the clouds using the path provided and returns whether the path is a valid traversal.
func isSafePath(path []int32, clouds []int32) bool {
	position := int32(0)
	lenClouds := int32(len(clouds))
	for i := 0; i < len(path); i++ {
		// We jump first as the first jump is free.
		position += path[i]

		// We've jumped too far and have an invalid path
		if position >= lenClouds {
			return false
		}
		// We've hit a thunderhead. Try the next one
		if clouds[position] == 1 {
			return false
		}
	}
	return true
}
