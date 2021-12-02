package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

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
	result := addTwoNumbers(&node1_1, &node2_1)
	PrintFullNode(result)
}

// optimized
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

// rewrite???
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l2 != nil {
			l1.Val += l2.Val
		}
		l1.Val += carry
		carry = l1.Val / 10
		l1.Val = l1.Val % 10
		if l1.Next == nil {
			if l2.Next != nil {
				l1.Next = l2.Next
				l2.Next = &ListNode{0, nil}
			} else if carry != 0 {
				l1.Next = &ListNode{0, nil}
			} else {
				break
			}
		}
		l1 = l1.Next
		if l2.Next == nil {
			l2.Next = &ListNode{0, nil}
		}
		l2 = l2.Next
	}
	return head
}

func PrintFullNode(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
