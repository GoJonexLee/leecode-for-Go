package easy

func flipAndInvertImage(A [][]int) [][]int {

	var invert = func(num *[]int) {
		for i := 0; i < len(*num); i++ {
			if (*num)[i] == 1 {
				(*num)[i] = 0
			} else {
				(*num)[i] = 1
			}
		}
	}

	for i := 0; i < len(A); i++ {
		size := len(A[i])
		for j := 0; j < size/2; j++ {
			A[i][j], A[i][size-j-1] = A[i][size-j-1], A[i][j]
		}
		invert(&A[i])
	}

	return A
}
