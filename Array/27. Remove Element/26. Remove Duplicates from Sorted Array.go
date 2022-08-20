func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1 // return slow我研究了几个小时才发现应该是slow+1
}