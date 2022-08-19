//思路1：在有序数组中找到最后一个比 target 小的元素
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			left = mid + 1
		}
	}
	return 0
}

//思路2：return right + 1 分别处理四种情况
// 目标值在数组所有元素之前  [0, -1]
// 目标值等于数组中某一个元素  return middle
// 目标值插入数组中的位置 [left, right]，return  right + 1
// 目标值在数组所有元素之后的情况 [left, right]， 因为是右闭区间，所以 return right + 1

/* func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right{
        mid := left + (right-left)/2   // 防止溢出 等同于(left+right)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid] > target{    //target 在左区间，所以[left, middle - 1]
            right = mid - 1
        }else{
            left = mid + 1              //target 在右区间，所以[middle + 1, right]
        }
    }
	// 分别处理如下四种情况
	// 目标值在数组所有元素之前  [0, -1]
	// 目标值等于数组中某一个元素  return middle
	// 目标值插入数组中的位置 [left, right]，return  right + 1
	// 目标值在数组所有元素之后的情况 [left, right]， 因为是右闭区间，所以 return right + 1
    return right + 1
} */

/* 错误记录：
func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target{
            right = mid - 1
        } else {
            left = mid + 1
        }
        //break
    }
    if nums[left] < target {
        return left + 1
        } else if left != 0 {
            return left - 1
        } else {
            return 0
        }
}*/
//1.通过4/64 增加else if left != 0的判断 修改right := len(nums) - 1
//2.通过41/64 去掉break
//3.通过46/64 思路完全错误 都没和次小值比较