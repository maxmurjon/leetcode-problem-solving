package medianfunc

import (
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	first := len(nums1)/2 - 1
	last := len(nums1) / 2
	if int(float64(len(nums1))/2*10)%10 == 0 {
		return float64((nums1[first] + nums1[last])) / 2
	}
	return float64(nums1[first+1])
}
