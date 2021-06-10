package algcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/**
	  1. 计算两个数组的长度
	  2. 根据两个长度判断是取1个还是2个
	  3. 同时遍历两个数组
	  4. 找到要找的值后计算中位数
	*/

	l1, l2 := len(nums1), len(nums2)

	if (l1+l2)&1 == 0 { // 偶数个
		k := (l1 + l2) >> 1
		//取 第 k 和 第 k+1
		return (getK(nums1, nums2, k) + getK(nums1, nums2, k+1)) / 2
	} else { //
		k := (l1 + l2 + 1) >> 1
		// 取第 k 个
		return getK(nums1, nums2, k)
	}

}

func getK(nums1, nums2 []int, k int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	}
	if len(nums2) == 0 {
		return float64(nums1[k-1])
	}
	l1, l2 := 0, 0
	ret := nums1[0]
	for k > 0 {
		if l1 < len(nums1) && l2 < len(nums2) {
			if nums1[l1] < nums2[l2] {
				ret = nums1[l1]
				l1++
			} else {
				ret = nums2[l2]
				l2++
			}
		} else if l1 < len(nums1) {
			ret = nums1[l1]
			l1++
		} else if l2 < len(nums2) {
			ret = nums2[l2]
			l2++
		}
		k--
	}
	return float64(ret)
}
