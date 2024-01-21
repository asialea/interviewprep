package main
import "fmt"
// You are given the head of a linked list, which contains a series of integers separated by 0's. The beginning and end of the linked list will have Node.val == 0.

// For every two consecutive 0's, merge all the nodes lying in between them into a single node whose value is the sum of all the merged nodes. The modified list should not contain any 0's.

// Return the head of the modified linked list.

// Beginning and end of list will always start with 0

// Idea, do the fast and slow pointer approach.
// how do we know weve reached end of list? - when fast == nil

// 4->4,5,2,0
// 0,3,1,0,4,5,2,0

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	slow := head
	fast := slow.Next
	for fast != nil {
		count := 0
		for fast.Val != 0 {
			count += fast.Val
			fast = fast.Next
		}
		fast.Val = count
		slow.Next = fast
		slow = fast
		fast = slow.Next
	}
	return head.Next
}

func main() {
	a := &ListNode{Val: 0,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val:  0,
									Next: nil}}}}}}}}
	l := mergeNodes(a)
	for(l != nil){
		fmt.Println(l.Val)
		l=l.Next
	}
}
