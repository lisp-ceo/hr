package twodarray

func hourglassSum(arr [][]int32) int32 {
	glasses := allGlasses(arr)
	return maxSum(glasses)
}

func allGlasses(arr [][]int32) [][]int32 {
	glasses := [][]int32{}
	for x := 0; x < len(arr) - 2; x++ {
		for y := 0; y < len(arr[x]) - 2; y++ {
			glasses = append(glasses, []int32{
				arr[x][y],
				arr[x][y+1],
				arr[x][y+2],

				arr[x+1][y+1],

				arr[x+2][y],
				arr[x+2][y+1],
				arr[x+2][y+2],
			})
		}
	}
	return glasses
}

func maxSum(glasses [][]int32) int32 {
	var sum int32
	for n, glass := range glasses {
		var glassSum int32
		for i := range glass {
			glassSum += glass[i]
		}
		if glassSum > sum || n == 0{
			sum = glassSum
		}
	}
	return sum
}
