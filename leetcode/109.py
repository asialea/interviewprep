# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# Given the head of a singly linked list where elements are sorted in ascending order, convert it to a 
# height-balanced
#  binary search tree.

# Ideas -  we can just recursively find mid and set that as root. since it is a singly linke list wewill have to be creative about finding the root
# We can use the fast and slow pointer concept to find he mid of the list
#  space complexity - constant   time conplexity - n logn
class Solution(object):
    # splitList splits a list at the mid and returns the middle node.
    def splitList(self, head):
        if not head:
            return None
        prev = None
        slow = head
        fast = head
        while(fast and fast.next):
            prev = slow
            slow = slow.next
            fast = fast.next.next
        if prev:
            prev.next = None
        return slow
    
    def sortedListToBST(self, head):
        """
        :type head: Optional[ListNode]
        :rtype: Optional[TreeNode]
        """
        if not head:
            return None
        if not head.next:
            return TreeNode(head.val)
        mid = self.splitList(head)
        curr = TreeNode(mid.val)
        curr.right = self.sortedListToBST(mid.next)
        curr.left = self.sortedListToBST(head)
        return curr


s = Solution()
b = ListNode(1, ListNode(2,  ListNode(3,  ListNode(4, ListNode(5)))))
a = ListNode(1, ListNode(2,  ListNode(3,  ListNode(4, ListNode(5, ListNode(6))))))
c = ListNode(1, ListNode(2))
d = ListNode(1)