# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

# Add two numbers. Given two non-empty linked lists representing two non-negative integers. Digits are stored in
# reverse order, each node contains a single digit. Add two numbers and return it as a linked list.
# Example
# Input: l1: (2 -> 4 -> 3) l2: (5 -> 6 -> 4)
# Output: 7 -> 0 -> 8
# Explanation: 342 + 465 = 807 (Lists represent numbers in reverse order.)
def addTwoNumbers(l1, l2):
    pass

# Given a linked list, removed the nth node from the end of the list and return its head.
# e.g head = 1 -> 2 -> 3 -> 4 -> 5, n = 2. The element that is 2nd from the end of the list is 4, so the list
# becomes 1 -> 2 -> 3 -> 5.
def removeNthFromEnd(head, n):
    pass

# Suppose a sorted array was rotated at some pivot unknown to you. e.g [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2].
# Given a target value to search, if found in the array, return its index. If not, return -1.
# Assume no duplicates exist in the array. Solution must be O(logn)
# e.g. nums = [4,5,6,7,0,1,2], target = 0. Output = 4
# e.g. nums = [4,5,6,7,0,1,2], target = 3. Output = -1
def search(nums, target):
    pass