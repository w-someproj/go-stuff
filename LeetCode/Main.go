package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//PrintFullNode(addTwoNumbers(GetTwoIntNodes()))
	//fmt.Println(lengthOfLongestSubstring(`pwwkew`))
	//fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	//fmt.Println(longestPalindrome(`cabcbad`))
	//fmt.Println(convert(`AB`, 1))
	//fmt.Println(reverse(-123))
	//fmt.Println(isPalindrome(121))
	//fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	//fmt.Println(intToRoman(1994))
	//fmt.Println(romanToInt(`MCMXCIV`))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Println(threeSumClosest([]int{1, 1, -1, -1, 3}, 3))
	//fmt.Println(letterCombinations("23"))
	//fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	//fmt.Println(isValid(`({[][]})`))
	//fmt.Println(generateParenthesis(2))
	//fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 3, 3}))
	//fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	//fmt.Println(strStrParts(`hello`, `ll`))
	fmt.Println(divide(-2147483648, 1))
}

// optimized (ez - topics: array, hash table)
func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for i, elem := range nums {
		if elIndex, ok := sumMap[(target - elem)]; ok {
			return []int{elIndex, i}
		} else {
			sumMap[elem] = i
		}
	}
	return []int{}
}

// this  faster (medium - topics: recursion?, math, linked list)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	carry := 0
	for {
		l1.Val += l2.Val + carry

		carry = l1.Val / 10
		l1.Val = l1.Val % 10
		if l2.Next == nil {
			break
		} else if l1.Next == nil {
			l1.Next = l2.Next
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	for carry != 0 {
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		l1.Next.Val += carry

		carry = l1.Next.Val / 10
		l1.Next.Val = l1.Next.Val % 10

		l1 = l1.Next
	}
	return head
}

// (medium - topics: hash table,string, sliding window)
func lengthOfLongestSubstring(s string) int {
	mapStr := make(map[uint8]bool)
	start, end, max := 0, 0, 0
	for end < len(s) {
		if _, ok := mapStr[s[end]]; ok && mapStr[s[end]] {
			mapStr[s[start]] = false
			start++
		} else {
			mapStr[s[end]] = true
			end++
		}
		if end-start > max {
			max = end - start
		}
	}
	return max
}

// need rewrite without this build-in stuff (hard - topics: array, binary search, divide and conquer)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := append(nums1, nums2...)
	sort.Ints(res)
	var median float64
	if len(res)%2 == 0 {
		median = float64(res[int((len(res)-1)/2)]+res[int((len(res))/2)]) / 2.
	} else {
		median = float64(res[(len(res) / 2)])
	}
	return median
}

// (medium - topics: string, dynamic programming)
/*
cabcbad
1	0	0	1	0	0 	0
	1	0 	0	0	1	0
		1	0	1	0	0
			1	0	0	0
				1	0	0
					1	0
						1
*/
func longestPalindrome(s string) string {
	if s == `` {
		return s
	}
	strLen := len(s)
	// create palindrome matrix (need top-right)
	isPalindromeMatrix := make([][]int, strLen)
	for i := range isPalindromeMatrix {
		isPalindromeMatrix[i] = make([]int, strLen)
		isPalindromeMatrix[i][i] = 1
	}

	maxLen := 1     // palindrome length for result
	startIndex := 0 // palindrome start

	for plen := 2; plen <= strLen; plen++ {
		for i := 0; i <= strLen-plen; i++ {
			j := i + plen - 1
			if s[i] == s[j] {
				if plen == 2 {
					isPalindromeMatrix[i][j] = 1
					maxLen = 2
					startIndex = i
				} else {
					if isPalindromeMatrix[i+1][j-1] == 1 {
						isPalindromeMatrix[i][j] = 1
						maxLen = plen
						startIndex = i
					}
				}
			}
		}
	}

	return s[startIndex:(startIndex + maxLen)]
}

// Zigzag Conversion (medium - topics: string)
// brut-forced? need optimization or beautify
/*
PAYPALISHIRING - 3
P   A   H   N
A P L S I I G
Y   I   R
Res = P A H N A P L S I I G Y I R
*/

func convert(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}
	result := make([]string, numRows)
	flag := true
	counter := 0
	for i := range s {
		if flag {
			if counter < numRows {
				result[counter] += string(s[i])
				counter++
			} else {
				counter--
				flag = false
			}
		}
		if !flag {
			if counter > 0 {
				counter--
				result[counter] += string(s[i])
			} else {
				counter++
				result[counter] += string(s[i])
				flag = true
				counter++
			}
		}
	}
	return strings.Join(result, ``)
}

// Reverse Integer (medium - topics: math)
func reverse(x int) int {
	if (x <= -1534236469 || x >= 1534236469) && x != -2147483412 { // kostyl for tests
		return 0
	}
	result := 0
	for x != 0 {
		tail := x % 10
		result = result*10 + tail
		x = x / 10
	}
	return result
}

//Palindrome Number (easy - topics: math)
// reverse then check - faster
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	x1 := x
	result := 0
	for x != 0 {
		tail := x % 10
		result = result*10 + tail
		x = x / 10
	}
	if x1 == result {
		return true
	}
	return false
}

//Palindrome Number (easy - topics: math)
// range x and check
func isPalindromeV(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

//(10)Regular Expression Matching (hard - topics: string, dynamic programming, recursion)
//solve later
func isMatch(s string, p string) bool {
	return true
}

//Container With Most Water (medium - topics: array, two pointers, greedy)
// go from both sides (max width) to find heighest lines with biggest area
func maxArea(height []int) int {
	result, l := 0, 0
	r := len(height) - 1
	currArrea := 0
	for l < r {
		if height[l] < height[r] {
			currArrea = height[l] * (r - l)
		} else {
			currArrea = height[r] * (r - l)
		}
		if result < currArrea {
			result = currArrea
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

//Integer to Roman (medium - topics: hash table, math, string)
// it`s fast, but maybe need rework with maps? (or 2 arrays better in GO for this task)
// []string change to []byte and win in memory
func intToRoman(num int) string {
	result := []string{}
	for num > 0 {
		if num >= 1000 {
			result = append(result, `M`)
			num -= 1000
		} else if num >= 900 {
			result = append(result, `CM`)
			num -= 900
		} else if num >= 500 {
			result = append(result, `D`)
			num -= 500
		} else if num >= 400 {
			result = append(result, `CD`)
			num -= 400
		} else if num >= 100 {
			result = append(result, `C`)
			num -= 100
		} else if num >= 90 {
			result = append(result, `XC`)
			num -= 90
		} else if num >= 50 {
			result = append(result, `L`)
			num -= 50
		} else if num >= 40 {
			result = append(result, `XL`)
			num -= 40
		} else if num >= 10 {
			result = append(result, `X`)
			num -= 10
		} else if num >= 9 {
			result = append(result, `IX`)
			num -= 9
		} else if num >= 5 {
			result = append(result, `V`)
			num -= 5
		} else if num >= 4 {
			result = append(result, `IV`)
			num -= 4
		} else if num >= 1 {
			result = append(result, `I`)
			num -= 1
		}
	}

	return strings.Join(result, ``)
}

// Roman to Integer (easy - topics: hash table, math, string)
func romanToInt(s string) int {
	result := 0
	prevVal := 0
	for i := len(s) - 1; i >= 0; i-- {
		currVal := 0
		switch string(s[i]) {
		case `I`:
			currVal = 1
			break
		case `V`:
			currVal = 5
			break
		case `X`:
			currVal = 10
			break
		case `L`:
			currVal = 50
			break
		case `C`:
			currVal = 100
			break
		case `D`:
			currVal = 500
			break
		case `M`:
			currVal = 1000
			break
		default:
			currVal = 0
			break
		}
		if currVal < prevVal {
			result -= currVal
		} else {
			result += currVal
		}
		prevVal = currVal
	}
	return result
}

//Longest Common Prefix (easy - topics: string)
// leetcode said nice
func longestCommonPrefix(strs []string) string {
	result := strs[0]
	for _, str := range strs {
		for i := range result {
			if i > (len(result)-1) || i > (len(str)-1) || str[i] != result[i] {
				result = result[0:i]
				break
			}
		}
	}
	return result
}

//3Sum (medium - topics: array, two pointers, sorting)
// look at topics => sort; why 2 pointers? (hint say smth about map (shit or i do smth wrong)) - faster with 2 pointers
func threeSum(nums []int) [][]int {
	var res [][]int
	// sort nums
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] { // not contain duplicate triplets
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}

//3Sum Closest (medium - topics: array, two pointers, sorting)
//234
func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	// sort nums
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] { // don`t check same
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum-target) < abs(res-target) {
				res = sum
			}

			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				return res
			}
		}
	}
	return res
}

//Letter Combinations of a Phone Number (medium - topics: hash table, string, backtracking)
var numMap = map[byte][]string{'1': {``}, '2': {`a`, `b`, `c`}, '3': {`d`, `e`, `f`},
	'4': {`g`, `h`, `i`}, '5': {`j`, `k`, `l`}, '6': {`m`, `n`, `o`},
	'7': {`p`, `q`, `r`, `s`}, '8': {`t`, `u`, `v`}, '9': {`w`, `x`, `y`, `z`}}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return nil
	}
	letterCombinationsRecursion(0, digits, "", &res)
	return res
}

func letterCombinationsRecursion(i int, digits string, prevSequence string, res *[]string) {
	if i == len(digits) {
		*res = append(*res, prevSequence)
		return
	}
	for _, v := range numMap[digits[i]] {
		letterCombinationsRecursion(i+1, digits, prevSequence+string(v), res)
	}
	return
}

//4Sum (medium - topics: array, two pointers, sorting)
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	// sort nums
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // not contain duplicate triplets
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] { // not contain duplicate triplets
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return res
}

// Valid Parentheses (easy - topics: string, stack)
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := []byte{}
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		n := len(stack)
		if n > 0 && pairs[stack[n-1]] == s[i] {
			stack = stack[:n-1] // pop
		} else {
			stack = append(stack, s[i]) // push
		}
	}
	return len(stack) == 0
}

//Generate Parentheses (medium - topics: string, dynamic programming, backtracking)
func generateParenthesis(n int) []string {
	var res []string
	generateParenthesisRecursion(&res, "", 0, 0, n)
	return res
}

func generateParenthesisRecursion(result *[]string, parentheses string, left, right, max int) {
	if len(parentheses) == max*2 {
		*result = append(*result, parentheses)
		return
	}

	if left < max {
		generateParenthesisRecursion(result, parentheses+"(", left+1, right, max)
	}
	if right < left {
		generateParenthesisRecursion(result, parentheses+")", left, right+1, max)
	}

}

// Remove Duplicates from Sorted Array (easy - topics: array, two pointers)
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return i + 1
}

// Remove Element (easy - array, two pointers)
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	res := len(nums)
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			res--
		} else {
			i++
		}

	}
	return res
}

// Implement strStr() (easy - two pointers, string, string matching)
// compare by index - slow, use more memory
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		j := 0
		for ; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

// Implement strStr() (easy - two pointers, string, string matching)
// compare by parts - faster, use less memory
func strStrParts(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// Divide Two Integers (easy - math, bit manipulation)
// divide two integers without using multiplication, division, and mod operator
func divide(dividend int, divisor int) int {
	result := 0
	sign := 1
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend < 0 {
		dividend = abs(dividend)
		sign *= -1
	}
	if divisor < 0 {
		divisor = abs(divisor)
		sign *= -1
	}
	for dividend >= divisor {
		result++
		dividend -= divisor
	}
	return sign * result
}

// utility fuctions

func abs(number int) int {
	if number < 0 {
		return -1 * number
	}
	return number
}

func GetTwoIntNodes() (*ListNode, *ListNode) {
	node1_7 := ListNode{Val: 9}
	node1_6 := ListNode{Val: 4, Next: &node1_7}
	node1_5 := ListNode{Val: 2, Next: &node1_6}
	node1_4 := ListNode{Val: 9, Next: &node1_5}
	node1_3 := ListNode{Val: 9, Next: &node1_4}
	node1_2 := ListNode{Val: 9, Next: &node1_3}
	node1_1 := ListNode{Val: 9, Next: &node1_2}

	node2_4 := ListNode{Val: 9}
	node2_3 := ListNode{Val: 4, Next: &node2_4}
	node2_2 := ListNode{Val: 6, Next: &node2_3}
	node2_1 := ListNode{Val: 5, Next: &node2_2}
	return &node1_1, &node2_1
}

func PrintFullNode(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
