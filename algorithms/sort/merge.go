package sort

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	p := len(nums) / 2
	left := MergeSort(nums[0:p])
	right := MergeSort(nums[p:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)

	var result []int

	for {
		if i >= m || j >= n {
			break
		}

		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	if i != m {
		for ; i < m; i++ {
			result = append(result, left[i])
		}
	}

	if j != n {
		for ; j < n; j++ {
			result = append(result, right[j])

		}
	}

	return result
}
