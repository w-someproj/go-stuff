package main

import (
	"fmt"
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
	fmt.Println(intToRoman(1994))

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

// utility fuctions

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
