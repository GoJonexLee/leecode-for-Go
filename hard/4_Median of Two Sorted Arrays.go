package hard

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n1 < n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	var (
		lo = 0
		hi = n2 * 2
	)

	var l1, r1, l2, r2 int64
	for lo <= hi {
		mid2 := (lo + hi) / 2
		mid1 := n1 + n2 - mid2

		if mid1 == 0 {
			l1 = math.MinInt64
		} else {
			l1 = int64(nums1[(mid1-1)/2])
		}

		if mid2 == 0 {
			l2 = math.MinInt64
		} else {
			l2 = int64(nums2[(mid2-1)/2])
		}

		if mid1 == n1*2 {
			r1 = math.MaxInt64
		} else {
			r1 = int64(nums1[mid1/2])
		}

		if mid2 == n2*2 {
			r2 = math.MaxInt64
		} else {
			r2 = int64(nums2[mid2/2])
		}

		if l1 > r2 {
			lo = mid2 + 1
		} else if l2 > r1 {
			hi = mid2 - 1
		} else {
			return (math.Max(float64(l1), float64(l2)) + math.Min(float64(r1), float64(r2))) / 2
		}
	}

	return -1
}
