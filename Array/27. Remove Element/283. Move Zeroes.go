func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		/* if (nums[slow] != 0 && nums[fast] == 0)||(nums[slow] == 0 && nums[fast] != 0 ) {
		    nums[slow], nums[fast] = nums[fast], nums[slow]
		} */
		if nums[slow] == 0 {
			if nums[fast] != 0 {
				nums[slow], nums[fast] = nums[fast], nums[slow] // Go的独特交换方式
				slow++
			}
		} else {
			slow++
		}
	}
	// return nums[slow+1]
}

/* 方法2：
func moveZeroes(nums []int)  {
    slow := 0
    for fast := 0; fast < len(nums); fast++ {
        if nums[fast] != 0 {
	    if fast != slow {
	        nums[slow], nums[fast] = nums[fast], nums[slow]
	    }
	slow++
} */

/* 这个加了会提高效率吗？
if len(nums) == 0 {
	return
} */
