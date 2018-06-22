package easy

func peakIndexInMountainArray(A []int) int {

	var i int
	for i = 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			break
		}
	}

	return i
}
