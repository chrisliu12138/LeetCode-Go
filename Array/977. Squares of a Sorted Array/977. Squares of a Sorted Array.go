// 双指针法，时间复杂度O(n)
// 注意：
// 可以说使用额外数组target存储结果，空间复杂度O(n)
// 也可以说除了存储答案的数组以外，只需要维护常量空间，空间复杂度：O(1)
func sortedSquares(nums []int) []int {
	left, right, i := 0, len(nums)-1, len(nums)-1
	// var target [i+1]int   // 报错non-constant array bound i + 1
	target := make([]int, len(nums)) // 注意要用make声明，长度为i+1不是i
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			target[i] = nums[right] * nums[right]
			i--
			right--
		} else {
			target[i] = nums[left] * nums[left]
			i--
			left++
		}
	}
	return target
}

/* 不能不使用target而只在原数组上操作的原因：
当第一次比较时，如果最左边的胜出，那么nums[n-1]就被nums[0]^2覆盖掉了，
第二次比较时，应该是nums[1]和nums[n-1]比较，但此时的nums[n-1]就不是之前的那个数了，被覆盖掉了
*/

// 解法2：暴力(不推荐)，平方后快排，时间复杂度O(n+log n)->0(log n)
/* func sortedSquares1(nums []int) []int {
	for i, value := range nums {
		nums[i] = value * value
	}
	sort.Ints(nums)
	return nums
}
对于[]int, []float, []string这种元素类型是基础类型的切片使用sort包提供的下面几个函数进行排序
sort.Ints
sort.Floats
sort.Strings */

/* 思路错误，比较过程可能会进行多次平方运算，结果和预期平方排序不同
func sortedSquares(nums []int) []int {
    slow := 0
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast]*nums[fast] < nums[slow]*nums[slow]{
            nums[slow], nums[fast] = nums[fast]*nums[fast], nums[slow]*nums[slow]
            // fast++
        } else {
            slow++
        }
    }
    return nums
} */