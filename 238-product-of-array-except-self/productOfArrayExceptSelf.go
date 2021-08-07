package productofarrayexceptself

func productExceptSelf(nums []int) []int {

	left := make([]int, len(nums))
	right := make([]int, len(nums))

	output := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			left[i] = 1
			continue
		}

		left[i] = left[i-1] * nums[i-1]
	}

	for i := len(nums) - 1; i >= 0; i-- {

		if i == len(nums)-1 {
			right[i] = 1
			continue
		}
		right[i] = right[i+1] * nums[i+1]

	}

	for i := 0; i < len(nums); i++ {
		output[i] = left[i] * right[i]
	}

	return output

}
