package arrays

type BinarySearch struct {
	Array  []int `json:"array"`
	Target int   `json:"target"`
}

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
