// Cookbook解法1：二分法 查找区间[0,x]，找到最后一个满足 n^2 <= x 的整数n
func mySqrt(x int) int {
	left, right := 0, x
	for left < right { // 不用 <= ，若 x=0，不用单独判断，直接return left
		// mid := left + (right-left)/2
		mid := (left + right + 1) / 2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}

/* *解法2：牛顿迭代法
   求根号 x，即求满足 x^2 - n = 0 方程的所有解*/
/* func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
} */

/* 我的解法：二分查找区间[1,x]， x = 0需要单独考虑return 0
   判断mid*mid == x 没必要，只要不大于，<= 都直接输出mid */
/* func mySqrt(x int) int {
	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {           // 只考虑mid*mid > x { return mid - 1即可
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			if (mid+1)*(mid+1) > x {
				return mid          // 写成if mid*mid > x { return mid-1更简洁
			} else {
				left = mid + 1
			}
		}
	}
	return 0 //注意边界 因为left=1所以不包括x=0的情况 输入0 输出为0
} */