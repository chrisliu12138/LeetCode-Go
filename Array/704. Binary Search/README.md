**二分查找注意：**
* `mid = (left + right) / 2` 容易溢出 (Integer Overflow)！因为 `left+right` 很容易超过int范围2147483647，而 `mid = left + (right-left)/2` 不容易溢出

* 右移运算符 `>>` ，运算结果正好能代替数学上的 `/2` 运算，但是比 `/2` 运算要快

**代码建议：**
* 用 `mid = left + (right-left)>>1` 替代 ~~`mid = (left + right) / 2`~~

* 都使用**左闭右闭**写法：定义 target 是在一个在左闭右闭的区间里，也就是 [left, right] (这个很重要) ，**要写 `for left <= right` `right = mid - 1`**

## 704. 二分查找

### 题目大意
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1

提示：

* 你可以假设 nums 中的所有元素是不重复的
* n 将在 [1, 10000]之间
nums 的每个元素都将在 [-9999, 9999]之间

### 解题思路
给出一个数组，要求在数组中搜索等于 target 的元素的下标。如果找到就输出下标，如果找不到输出 -1

**简单题，二分搜索的裸题**

### 代码实现

```go
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
```

## 35. 搜索插入位置

### 题目大意 
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置，使用O(log n)，不能暴力解法

你可以假设数组中无重复元素

### 解题思路

这一题是**经典的二分搜索的变种题**

**思路1. 在有序数组中找到最后一个比 target 小的元素**

### 思路1 代码实现

```go

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
```
**思路2. `return right + 1`考虑以下四种情况:**

* 目标值在数组所有元素之前
* 目标值等于数组中某一个元素
* 目标值插入数组中的位置
* 目标值在数组所有元素之后

### 思路2 代码实现

```go
func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right{
        mid := left + (right-left)/2   // 防止溢出 等同于(left+right)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid] > target{    // target 在左区间，所以[left, middle - 1]
            right = mid - 1
        }else{
            left = mid + 1              // target 在右区间，所以[middle + 1, right]
        }
    }
	// 分别处理如下四种情况
	// 目标值在数组所有元素之前  [0, -1]
	// 目标值等于数组中某一个元素  return middle
	// 目标值插入数组中的位置 [left, right]，return  right + 1
	// 目标值在数组所有元素之后的情况 [left, right]， 因为是右闭区间，所以 return right + 1
    return right + 1
}
```
## 34. 在排序数组中查找元素的第一个和最后一个位置

### 题目大意
给定一个按照**非递减 (可以有多个相同元素)** 排列的整数数组 nums，和一个目标值 target，找出给定目标值在数组中的开始位置和结束位置，如果数组中不存在目标值，返回 [-1, -1]

算法时间复杂度必须是 O(log n) 级别 (不能暴力循环)

### 解题思路
给出一个有序数组 nums 和一个数 target，要求在数组中找到第一个和这个元素相等的元素下标，最后一个和这个元素相等的元素下标

这一题是**经典的二分搜索变种题**

二分搜索有 **4 大基础变种题：**

1. 查找第一个值等于给定值的元素
2. 查找最后一个值等于给定值的元素
3. 查找第一个大于等于给定值的元素
4. 查找最后一个小于等于给定值的元素

这一题的解题思路可以分别利用变种 1 和变种 2 的解法就可以做出此题，详细代码如下：

### 代码实现

```go
// 不要想着直接return左右两个边界 要用主函数调用两个函数分别return
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
	    return []int{-1, -1}
	}   // 不确定该判断能不能提升性能
	return []int{searchLeftRange(nums, target), searchRightRange(nums, target)}
}

// 变种1：二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchLeftRange(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			// mid = right - 1 开始写错成修改mid
			right = mid - 1
		} else if nums[mid] < target {
			// mid = left + 1
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target { // 找到第一个与target相等的元素直接return
				return mid
			} else {            // 不用else{}也可以
				right = mid - 1 // 这行是核心 一开始我写成mid--
			} 
		}
	}
	return -1
}

// 变种2：二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchRightRange(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			//if (mid == len(nums)-1) || (nums[mid+1] != target) {  // 优先级：+、- 大于 ==、!= 大于||，可以不加括号
			if mid == len(nums)-1 || nums[mid+1] != target { // 找到最后一个与target相等的元素直接return
				return mid
			} else {
				left = mid + 1 // 这行是与二分搜索裸题不同的核心！！
			}
		}
	}
	return -1
}
```

或者用一次变种 1 的方法，然后循环往后找到最后一个与给定值相等的元素。不过后者这种方法可能会使时间复杂度下降到 O(n)，因为有可能数组中 n 个元素都和给定元素相同 (4 大基础变种的实现见代码)

**注意以下常见问题：**

```go
// 注意同时return两个值的写法
return []int{-1, -1}
return []int{ans1, ans2}

// 不确定该判断能不能提升性能
if len(nums) == 0 {
    return []int{-1, -1}
}  

// 优先级：+、- 大于 ==、!= 大于||，可以不加括号
if (mid == len(nums)-1) || (nums[mid+1] != target) {  
if mid == len(nums)-1 || nums[mid+1] != target {
```

**变种题3，4代码：**

```go
// 变种3：二分查找第一个大于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // 找到第一个大于等于 target 的元素
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 变种4：二分查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
func searchLastLessElement(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] > target) { // 找到最后一个小于等于 target 的元素
				return mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
```

## 69.x 的平方根

### 题目大意 

实现 `int sqrt(int x)` 函数。计算并返回 x 的算数平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去

### 解题思路

题目要求求出根号 x

根据题意，根号 x 的取值范围一定在 [0,x] 之间，这个区间内的值是递增有序的，有边界的，可以用下标访问的，满足这三点正好也就满足了二分搜索的 3 大条件

**解题思路1，二分搜索**

### 思路1 代码实现
```go
// 解法1：二分法 查找区间[0,x]，找到最后一个满足 n^2 <= x 的整数n
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
```

***解题思路2，牛顿迭代法。求根号 x，即求满足 x^2 - n = 0 方程的所有解**

### 思路2 代码实现

```go
// *解法2：牛顿迭代法 求根号 x，即求满足 x^2 - n = 0 方程的所有解
func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
```

## 367.有效的完全平方数

### 题目大意

给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False

说明：**不要使用**任何内置的库函数，如 sqrt

### 解题思路

给出一个数，要求判断这个数是不是完全平方数。
可以用二分搜索来解答这道题。判断完全平方数，根据它的定义来，是否能被开根号，即找到一个数的平方是否可以等于待判断的数字。从 [1, n] 区间内进行二分，若能找到则返回 true，找不到就返回 false

### 代码实现

```go
// 经典二分搜索的裸题 简单题
func isPerfectSquare(num int) bool {
    left, right := 1, num
    for left <= right {
        mid := left + (right-left)/2
        if mid*mid == num {
            return true
        } else if mid*mid > num {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return false
}
```

