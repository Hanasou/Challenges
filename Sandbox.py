
#%%
# longest palindromic substring
# Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
# e.g: Input: "babad". Output: "bab". "aba" is also valid
def longestPalindrome(s):
    if s == '':
        return ''
        
    current_pal = ''
    max_pal = ''
    for i in range(0,len(s)):
        for j in range(0,len(s))[::-1]:
            current_pal = s[i:j + 1]
            if current_pal == current_pal[::-1]:
                if len(current_pal) > len(max_pal):
                    max_pal = current_pal
                break
            if j == i:
                current_pal = ''

    return max_pal

def longestPalindrome2(s):
    if s == '':
        return ''
        
    current_pal = ''
    max_pal = ''
    for i in range(0,len(s)):
        for j in range(i, len(s)):
            current_pal += s[j]
            if current_pal == current_pal[::-1]:
                if len(current_pal) > len(max_pal):
                    max_pal = current_pal
            if j == len(s) - 1:
                current_pal = ''

    return max_pal
longestPalindrome('babad')


#%%
# Find the max two dimensional area given a list of numbers.
# Area is x * y. x = distance between two numbers. y = min(n1,n2), n1 and n2 are two numbers from the list.
def maxArea(height) -> int:
    max_area = 0
    for i in range(0, len(height)):
        for j in range(i, len(height)):
            x = j - i
            y = min((height[i], height[j]))
            max_area = max(x * y,max_area)
    return max_area


#%%
# 3sum. Given an array of n integers, find solution set of triplets where all three values sum to 0. Triplets must be unique.
def three_sum(nums):
    if len(nums) < 3:
            return []
    if nums == [0] * len(nums):
        return [[0,0,0]]
    solutions = []
    nums.sort()
    for i in range(0,len(nums)):
        j = i + 1
        k = len(nums) - 1
        while j < k:
            triple = (nums[i],nums[j],nums[k])
            sol = sum(triple)
            if (sol == 0):
                solutions.append(triple)
                j += 1
            elif (sol > 0):
                k -= 1
            elif (sol < 0):
                j += 1
    new_solutions = [list(ele) for ele in list(set(solutions))]
    return new_solutions
solution = three_sum([-1,0,1,2,-1,-4])
solution


#%%
# Given an array nums of n numbers, find the sum of 3 integers that sums closest to target. Return the sum of 3 integers.
# Assume that each input has exactly one solution
def threeSumClosest(nums, target):
    if nums == [0] * len(nums):
        return 0
    solutions = []
    nums.sort()
    solution = float("inf")
    for i in range(0,len(nums)):
        j = i + 1
        k = len(nums) - 1
        while j < k:
            triple = (nums[i],nums[j],nums[k])
            sol = sum(triple)
            solutions.append(sol)
            if (sol > target):
                k -= 1
            elif (sol < target):
                j += 1
            else:
                return sol
    min_val = float("inf")
    for n in solutions:
        prev_min = min_val
        min_val = min(min_val,abs(n - target))
        if min_val != prev_min:
            solution = n
    return solution
threeSumClosest([1,1,1,0],-100)


#%%
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


#%%
# Add two numbers. Given two non-empty linked lists representing two non-negative integers. Digits are stored in
# reverse order, each node contains a single digit. Add two numbers and return it as a linked list.
# Example
# Input: l1: (2 -> 4 -> 3) l2: (5 -> 6 -> 4)
# Output: 7 -> 0 -> 8
# Explanation: 342 + 465 = 807 (Lists represent numbers in reverse order.)
def addTwoNumbers(l1, l2):
    pass


#%%
# Given a linked list, removed the nth node from the end of the list and return its head.
# e.g head = 1 -> 2 -> 3 -> 4 -> 5, n = 2. The element that is 2nd from the end of the list is 4, so the list
# becomes 1 -> 2 -> 3 -> 5.
def removeNthFromEnd(head, n):
    pass


#%%
# Suppose a sorted array was rotated at some pivot unknown to you. e.g [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2].
# Given a target value to search, if found in the array, return its index. If not, return -1.
# Assume no duplicates exist in the array. Solution must be O(logn)
# e.g. nums = [4,5,6,7,0,1,2], target = 0. Output = 4
# e.g. nums = [4,5,6,7,0,1,2], target = 3. Output = -1
def search(nums, target):
    pass


#%%
# Power of Three
import math
def power_of_three(n):
    if n < 1:
        return False

    return (math.log(n) / math.log(3)) % 1 == 0
power_of_three(243)




#%%
# Reverse String
def rev_string(s):
    for i in range(0, len(s) // 2):
        temp = s[-i - 1]
        s[-i - 1] = s[i]
        s[i] = temp
var = ["h","e","l","l","o"]
rev_string(var)
print(var)

#%%
# Intersection of two arrays
def intersection(nums1, nums2):
    intersection = []
    for i in nums1:
        for j in nums2:
            if i == j:
                if i not in intersection:
                    intersection.append(i)
                break
    return intersection

#%%
string = 'apple'
string[0:len(string)]

#%%
# Unique Character
def unique_char(s):
    index = 0
    counts = {}
    for char in s:
        if char in counts.keys():
            counts[char] += 1
        else:
            counts[char] = 1
                
    for char in s:
        if counts[char] == 1:
            return index
        else:
            index += 1
    return -1

#%%
