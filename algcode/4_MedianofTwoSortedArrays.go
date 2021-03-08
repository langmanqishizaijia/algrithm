package algcode

/*

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

//1
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	len1, len2 := len(nums1), len(nums2)

	maxlen := int((len1+len2)/2) + 1

	lena := 0
	lenb := 0
	slice := make([]int, 0)
	for i := 0; i < maxlen; i++ {

		if lena < len1 && lenb < len2 {
			if nums1[lena] < nums2[lenb] {
				slice = append(slice, nums1[lena])
				lena++
			} else {
				slice = append(slice, nums2[lenb])
				lenb++
			}
		} else if lena < len1 {
			slice = append(slice, nums1[lena])
			lena++
		} else {
			slice = append(slice, nums2[lenb])
			lenb++
		}
	}

	if (len1+len2)%2 == 0 {
		var tmp float64 = float64(slice[(len1+len2)/2] + slice[(len1+len2)/2-1])
		return float64(tmp / 2)
	} else {
		return float64(slice[(len1+len2)/2])
	}

}

//2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	slice := make([]int, len(nums1)+len(nums2))
	copy(slice, nums1)
	copy(slice[len(nums1):], nums2)

	sort.Ints(slice)

	lenth := len(slice)
	if lenth%2 == 0 {
		fmt.Println("%+v", slice)
		return (float64)(slice[lenth/2]+slice[(lenth-1)/2]) / 2
	} else {
		return float64(slice[lenth/2])
	}

}
