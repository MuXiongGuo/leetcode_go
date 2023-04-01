package main

func findLengthOfShortestSubarray(arr []int) int {
	firstRight, lastLeft, n := -1, -1, len(arr)
	for i, v := range arr {
		if i != 0 && v < arr[i-1] {
			if firstRight == -1 {
				firstRight = i - 1
			}
			lastLeft = i
		}
	}
	if firstRight == -1 {
		return 0
	}
	if arr[0] <= arr[n-1] {
		c1, c2 := 0, 0
		for cur, i := arr[0], 0; i < n; i++ {

		}
	} else {

	}
}
