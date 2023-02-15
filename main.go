package main

import "strings"

type ListNode struct {
	Val  int
	Next *ListNode
}

func TwoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				break
			}
		}
		if len(result) > 0 {
			break
		}
	}
	return result
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result

	for l1 != nil || l2 != nil {
		if l1 != nil {
			temp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp.Val += l2.Val
			l2 = l2.Next
		}
		if temp.Val > 9 {
			temp.Val -= 10
			temp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			temp.Next = &ListNode{}
		}
		temp = temp.Next
	}
	return result
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := 0
	cpy := x
	for x != 0 {
		remainder := x % 10
		temp = temp*10 + remainder
		x = x / 10
	}
	return temp == cpy
}

func RemoveDuplicates(nums []int) int {
	newNumber := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[newNumber] = nums[i]
			newNumber++
		}
	}
	return newNumber
}

func MySqrt(x int) int {
	low, high := 0, x

	for low <= high {
		mid := (low + high) / 2

		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

func NumJewelsInStones(jewels string, stones string) int {
	count := 0

	for _, v1 := range jewels {
		for _, v2 := range stones {
			if v1 == v2 {
				count++
			}
		}
	}
	return count
}

func DefangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func BalancedStringSplit(s string) int {

	Count, rCount, lCount := 0, 0, 0

	for _, v := range s {
		if string(v) == "R" {
			rCount++
		} else {
			lCount++
		}

		if rCount == lCount {
			Count++
			rCount = 0
			lCount = 0
		}
	}
	return Count

}

func SmallerNumbersThanCurrent(nums []int) []int {
	var arr = make([]int, len(nums))

	for i := range nums {
		count := 0
		for j := 0; j < len(nums); j++ {
			if j != i && nums[j] < nums[i] {
				count++
			}
		}
		arr[i] = count
	}
	return arr
}

func KidsWithCandies(candies []int, extraCandies int) []bool {
	var result = make([]bool, len(candies))
	max := 0
	for _, v := range candies {
		if max < v {
			max = v
		}
	}

	for i, v := range candies {
		if v+extraCandies >= max {
			result[i] = true
		}
	}
	return result
}

func Shuffle(nums []int, n int) []int {
	var arr = make([]int, len(nums))
	for i := 0; i < n; i++ {
		arr[2*i] = nums[i]
		arr[2*i+1] = nums[n+i]
	}
	return arr
}

func RunningSum(nums []int) []int {
	var runningSum = make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		runningSum[i] = sum
	}
	return runningSum
}

func NumIdenticalPairs(nums []int) int {
	op := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				op++
			}
		}
	}
	return op

}

func RestoreString(s string, indices []int) string {
	xb := make([]byte, len(s))
	for i, v := range indices {
		xb[v] = s[i]
	}
	return string(xb)
}

func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func MaximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, v1 := range accounts {
		personWealth := 0
		for j := range v1 {
			personWealth += v1[j]
		}
		if maxWealth < personWealth {
			maxWealth = personWealth
		}
	}
	return maxWealth
}
