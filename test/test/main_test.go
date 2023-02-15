package main

import (
	"testing"
)

// Test twosum
func TestTwoSum(t *testing.T) {

	type twoSumTC struct {
		caseName string
		input    []int
		target   int
		want     []int
	}
	a := []twoSumTC{
		{
			caseName: "TS1",
			input:    []int{2, 7, 11, 15},
			target:   9,
			want:     []int{0, 1},
		},
		{
			caseName: "TS2",
			input:    []int{3, 2, 4},
			target:   6,
			want:     []int{1, 2},
		},
		{
			caseName: "TS3",
			input:    []int{3, 3},
			target:   6,
			want:     []int{0, 1},
		},
	}

	for _, v := range a {
		result := TwoSum(v.input, v.target)
		for i := range v.want {
			if v.want[i] != result[i] {
				t.Errorf("%s Testcase not accepted", v.caseName)
			}
		}

	}
}

func TestIsPalindrome(t *testing.T) {
	type IsPalindromeTC struct {
		caseName string
		input    int
		want     bool
	}

	a := []IsPalindromeTC{
		{
			caseName: "tc1",
			input:    121,
			want:     true,
		},
		{
			caseName: "tc2",
			input:    -121,
			want:     false,
		},
		{
			caseName: "tc3",
			input:    10,
			want:     false,
		},
	}

	for _, v := range a {
		result := IsPalindrome(v.input)
		if result != v.want {
			t.Errorf("%v not accepted", v.caseName)
		}

	}
}

// func TestAddTwoNumbers(t *testing.T) {
// 	type testcaseName struct {
// 		caseName          string
// 		input1, input2 *ListNode
// 		want           *ListNode
// 	}
// 	a := []testcaseName{
// 		{
// 			caseName:  "tc1",
// 		},
// 	}
// }

func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		caseName string
		input    []int
		want     int
	}
	a := []testCase{
		{caseName: "tc1", input: []int{1, 1, 2}, want: 2},
		{caseName: "tc1", input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
	}
	for _, v := range a {
		result := RemoveDuplicates(v.input)
		if result != v.want {
			t.Errorf("%v testCase not accepted", v.caseName)
		}
	}
}

func TestMySqrt(t *testing.T) {
	type Testcase struct {
		caseName string
		input    int
		want     int
	}

	a := []Testcase{
		{caseName: "tc1", input: 4, want: 2},
		{caseName: "tc2", input: 8, want: 2},
	}

	for _, v := range a {
		result := MySqrt(v.input)
		if result != v.want {
			t.Errorf("%v testCase is not accepted", v.caseName)
		}
	}
}

func TestNumJewelsInStones(t *testing.T) {
	type Testcase struct {
		caseName       string
		input1, input2 string
		want           int
	}

	a := []Testcase{
		{caseName: "tc1", input1: "aA", input2: "aAAbbbb", want: 3},
		{caseName: "tc2", input1: "z", input2: "ZZ", want: 0},
	}

	for _, v := range a {
		results := NumJewelsInStones(v.input1, v.input2)
		if results != v.want {
			t.Errorf("%v test case not accepted.", v.caseName)
		}
	}
}

func TestDefangIPaddr(t *testing.T) {
	type testcase struct {
		caseName string
		input    string
		want     string
	}

	a := []testcase{
		{caseName: "tc1", input: "1.1.1.1", want: "1[.]1[.]1[.]1"},
		{caseName: "tc2", input: "255.100.50.0", want: "255[.]100[.]50[.]0"},
	}

	for _, v := range a {
		results := DefangIPaddr(v.input)
		if results != v.want {
			t.Errorf("%v testcase not accepted", v.caseName)
		}
	}
}

func TestBalancedStringSplit(t *testing.T) {
	type testcase struct {
		caseName string
		input    string
		want     int
	}

	a := []testcase{
		{caseName: "tc1", input: "RLRRLLRLRL", want: 4},
		{caseName: "tc2", input: "RLRRRLLRLL", want: 2},
	}
	for _, v := range a {
		results := BalancedStringSplit(v.input)
		if results != v.want {
			t.Errorf("%v testcase not accepted", v.caseName)
		}
	}
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	type testcase struct {
		caseName string
		input    []int
		want     []int
	}

	a := []testcase{
		{caseName: "tc1", input: []int{8, 1, 2, 2, 3}, want: []int{4, 0, 1, 1, 3}},
		{caseName: "tc2", input: []int{6, 5, 4, 8}, want: []int{2, 1, 0, 3}},
		{caseName: "tc3", input: []int{7, 7, 7, 7}, want: []int{0, 0, 0, 0}},
	}
	for _, v := range a {
		result := SmallerNumbersThanCurrent(v.input)
		for i := range v.want {
			if v.want[i] != result[i] {
				t.Errorf("%s Testcase not accepted", v.caseName)
			}
		}
	}
}

func TestKidsWithCandies(t *testing.T) {
	type testcase struct {
		caseName string
		input1   []int
		input2   int
		want     []bool
	}

	a := []testcase{
		{caseName: "tc1", input1: []int{2, 3, 5, 1, 3}, input2: 3, want: []bool{true, true, true, false, true}},
		{caseName: "tc2", input1: []int{4, 2, 1, 1, 2}, input2: 1, want: []bool{true, false, false, false, false}},
		{caseName: "tc3", input1: []int{12, 1, 12}, input2: 10, want: []bool{true, false, true}},
	}
	for _, v := range a {
		result := KidsWithCandies(v.input1, v.input2)
		for i := range v.want {
			if v.want[i] != result[i] {
				t.Errorf("%s Testcase not accepted", v.caseName)
			}
		}
	}
}

func TestShuffle(t *testing.T) {
	type testcase struct {
		caseName string
		input1   []int
		input2   int
		want     []int
	}

	a := []testcase{
		{caseName: "tc1", input1: []int{2, 5, 1, 3, 4, 7}, input2: 3, want: []int{2, 3, 5, 4, 1, 7}},
		{caseName: "tc2", input1: []int{1, 2, 3, 4, 4, 3, 2, 1}, input2: 4, want: []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{caseName: "tc3", input1: []int{1, 1, 2, 2}, input2: 2, want: []int{1, 2, 1, 2}},
	}

	for _, v := range a {
		result := Shuffle(v.input1, v.input2)
		for i := range v.want {
			if v.want[i] != result[i] {
				t.Errorf("%s Testcase not accepted", v.caseName)
			}
		}
	}
}

func TestRunningSum(t *testing.T) {
	type testcase struct {
		caseName string
		input    []int
		want     []int
	}

	a := []testcase{
		{caseName: "tc1", input: []int{1, 2, 3, 4}, want: []int{1, 3, 6, 10}},
		{caseName: "tc2", input: []int{1, 1, 1, 1, 1}, want: []int{1, 2, 3, 4, 5}},
		{caseName: "tc3", input: []int{3, 1, 2, 10, 1}, want: []int{3, 4, 6, 16, 17}},
	}

	for _, v := range a {
		result := RunningSum(v.input)
		for i := range v.want {
			if v.want[i] != result[i] {
				t.Errorf("%s Testcase not accepted", v.caseName)
			}
		}
	}
}

func TestNumIdenticalPairs(t *testing.T) {
	type testcase struct {
		caseName string
		input    []int
		want     int
	}

	a := []testcase{
		{caseName: "tc1", input: []int{1, 2, 3, 1, 1, 3}, want: 4},
		{caseName: "tc2", input: []int{1, 1, 1, 1}, want: 6},
		{caseName: "tc3", input: []int{1, 2, 3}, want: 0},
	}

	for _, v := range a {
		result := NumIdenticalPairs(v.input)
		if v.want != result {
			t.Errorf("%s Testcase not accepted", v.caseName)
		}
	}
}

func TestRestoreString(t *testing.T) {
	type testcase struct {
		caseName string
		input1   string
		input2   []int
		want     string
	}

	a := []testcase{
		{caseName: "tc1", input1: "codeleet", input2: []int{4, 5, 6, 7, 0, 2, 1, 3}, want: "leetcode"},
		{caseName: "tc2", input1: "abc", input2: []int{0, 1, 2}, want: "abc"},
	}

	for _, v := range a {
		result := RestoreString(v.input1, v.input2)
		if v.want != result {
			t.Errorf("%s Testcase not accepted", v.caseName)
		}
	}
}

func TestArrayStringsAreEqual(t *testing.T) {
	type testcase struct {
		caseName string
		input1   []string
		input2   []string
		want     bool
	}

	a := []testcase{
		{caseName: "tc1", input1: []string{"ab", "c"}, input2: []string{"a", "bc"}, want: true},
		{caseName: "tc2", input1: []string{"a", "cb"}, input2: []string{"ab", "c"}, want: false},
		{caseName: "tc3", input1: []string{"abc", "d", "defg"}, input2: []string{"abcddefg"}, want: true},
	}

	for _, v := range a {
		result := ArrayStringsAreEqual(v.input1, v.input2)
		if v.want != result {
			t.Errorf("%s Testcase not accepted", v.caseName)
		}
	}

}

func TestMaximumWealth(t *testing.T) {
	type testcase struct {
		caseName string
		input    [][]int
		want     int
	}

	a := []testcase{
		{caseName: "tc1", input: [][]int{{1, 2, 3}, {3, 2, 1}}, want: 6},
		{caseName: "tc2", input: [][]int{{1, 5}, {7, 3}, {3, 5}}, want: 10},
		{caseName: "tc3", input: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, want: 17},
	}

	for _, v := range a {
		result := MaximumWealth(v.input)
		if v.want != result {
			t.Errorf("%s Testcase not accepted", v.caseName)
		}
	}
}
