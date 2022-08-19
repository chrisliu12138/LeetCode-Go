//不要想着直接return左右两个边界 要用主函数调用两个函数分别return
func searchRange(nums []int, target int) []int {
	/* 	if len(nums) == 0 {
	    return []int{-1, -1}
	} */
	return []int{searchLeftRange(nums, target), searchRightRange(nums, target)}
}

// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchLeftRange(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			//mid = right - 1 开始写错成修改mid
			right = mid - 1
		} else if nums[mid] < target {
			// mid = left + 1
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target { // 找到第一个与target相等的元素直接return
				return mid
			} else {
				right = mid - 1 //这行是核心 一开始我写成mid--
			} //不用else{}也可以
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchRightRange(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			//if (mid == len(nums)-1) || (nums[mid+1] != target) {  //优先级：+、- 大于 ==、!= 大于||，可以不加括号
			if mid == len(nums)-1 || nums[mid+1] != target { // 找到最后一个与target相等的元素直接return
				return mid
			} else {
				left = mid + 1 // 这行是与二分搜索裸题不同的核心！！
			}
		}
	}
	return -1
}

/* 错误记录1：左右一起查找不方便一起return 并且函数修改left right写成修改mid
func searchRange(nums []int, target int) []int {
    return []int{searchLeftRange(nums, target), searchRightRange(nums, target)}
    left, right := 0, len(nums)-1
    var l, r int
    for left <= right {
        mid := left + (right-left)/2
        if (nums[mid] == target) && (nums[mid-1] == target){
            mid--
        } else if (nums[mid] == target) && (nums[mid+1] == target){
            mid++
        } else if nums[mid] > target {
            mid = right - 1
        } else if nums[mid] < target{
            mid = left + 1
        } else if (nums[mid] == target) && (nums[mid-1] != target){
            l = mid
        } else if (nums[mid] == target) && (nums[mid+1] != target){
            r = mid
        }
    }
    // fmt.Printf("[%d,%d]", l, r)
    // return fmt.Printf("[-1,-1]")
    //return fmt.Printf("[%d,%d]", l, r)
    return l,r
}
错误记录2：左右边界查找分开后 函数修改left right写成修改mid
func searchLeftRange(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if (nums[mid] == target) && (nums[mid-1] == target){
            mid--
        } else if (nums[mid] == target) && (nums[mid-1] != target){
            return mid
        } else if nums[mid] < target {
            mid = left + 1
        }
    }
    return -1
}
func searchRightRange(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if (nums[mid] == target) && (nums[mid+1] == target){
            mid++
        } else if (nums[mid] == target) && (nums[mid+1] != target){
            return mid
        } else  if nums[mid] > target {
            mid = right - 1
        }
    }
    return -1
} */

//api选手
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
/*  var searchRange = function(nums, target) {
    let left = nums.indexOf(target)
    let right = nums.lastIndexOf(target)
    return [left,right]
} */