package arrays

func binarySearch(bs BinarySearch) int {
	nums := bs.Array
	target := bs.Target
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle - 1
			continue
		}

		if nums[middle] < target {
			left = middle + 1
		}
	}

	return -1
}

func productExceptSelf(pes ProductExceptSelf) []int {
	nums := pes.Array
	output := make([]int, len(nums))

	prefix := 1
	for i, n := range nums {
		output[i] = prefix
		prefix *= n
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] *= postfix
		postfix *= nums[i]
	}

	return output
}
