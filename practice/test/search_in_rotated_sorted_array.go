package main

import "fmt"

func main() {
	fmt.Println("Search in Rotated Sorted Array")
	doTest([]int{5, 1, 2, 3, 4}, 4)
	doTest([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	doTest([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	doTest([]int{14, 15, 16, 17, 18, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0)
	doTest([]int{4, 5, 6, 7, 0, 1, 2}, 3)
}

func doTest(nums []int, target int) {
	fmt.Println("Input: ", nums)
	res := search(nums, target)
	fmt.Printf("Target: %d, Result= %d\n\n", target, res)
}

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}

	start, end := 0, l-1

	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		}

		if nums[mid] < nums[end] {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if target < nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}

func searchOld(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}

	start, end := 0, l-1

	if nums[start] > nums[end] {
		return binarySearchPivot(nums, target, start, end)
	}
	return binarySearchNormal(nums, target, start, end)
}

func binarySearchPivot(nums []int, target, start, end int) int {
	l := len(nums) - 1
	for nums[start] > nums[end] {
		mid := start + (end-start)/2
		// fmt.Printf("startP=%d, endP=%d, midP=%d\n", start, end, mid)

		if nums[mid] > nums[mid+1] {
			if nums[mid+1] <= target && target <= nums[l] {
				return binarySearchNormal(nums, target, mid+1, l)
			}
			return binarySearchNormal(nums, target, 0, mid)
		}
		if nums[mid-1] > nums[mid] {
			if nums[mid] <= target && target <= nums[l] {
				return binarySearchNormal(nums, target, mid, l)
			}
			return binarySearchNormal(nums, target, 0, mid-1)
		}
		if nums[mid] < nums[end] {
			end = mid
		} else {
			start = mid
		}
	}

	return -1
}

func binarySearchNormal(nums []int, target, start, end int) int {

	for start <= end {
		mid := start + (end-start)/2
		// fmt.Printf("start=%d, end=%d, mid=%d\n", start, end, mid)

		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
		// fmt.Printf("Updated start=%d, end=%d, mid=%d\n", start, end, mid)
	}

	return -1

}
