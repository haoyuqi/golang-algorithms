package sort

func QuickSort(nums []int, p int, r int) {
	if p >= r {
		return
	}
	q := partition(nums, p, r)

	QuickSort(nums, p, q-1)
	QuickSort(nums, q+1, r)
}

func partition(nums []int, p int, r int) int {
	pivot := nums[r]

	i := p

	for j := p; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[r] = pivot, nums[i]

	return i
}
