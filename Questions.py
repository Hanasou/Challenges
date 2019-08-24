# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
    head = ListNode(0)
    number = head
    carry = 0
        
    while l1 != None or l2 != None:
        if l1 == None:
            l1 = ListNode(0)
        if l2 == None:
            l2 = ListNode(0)
        int_digit = (l1.val + l2.val + carry) % 10
        number.next = ListNode(int_digit)
        number = number.next
        carry = (l1.val + l2.val + carry) // 10
        if l1 != None:
            l1 = l1.next
        if l2 != None:
            l2 = l2.next
        
    if carry == 1:
        number.next = ListNode(carry)
 
    return head.next
# Determine if a string has all unique characters
# What if you cannot use additional data structures?
def is_unique(s):
    for i in range(0, len(s)):
        for j in range(i, len(s)):
            if s[i] == s[j]:
                return False
    return True

# Given two strings, write a method to decide if one is a permutation of the other
def check_perm(s1, s2):
    s1 = s1.sort()
    s2 = s2.sort()
    return s1 == s2

# Given a string, write a function to check if it is a permutation of a palindrome
# e.g. Input: Tact Coa
# Output: True (permutations: 'taco cat', 'atco cta', etc)
def pal_perm(s):
    pass
    counter = -1
    s = list(s)
    s = [x for x in s if not ' ']
    for i in range(0, len(s)):
        for j in range(i + 1, len(s)):
            pass
            if s[i] == s[j]:
                s[j], s[counter] = s[counter], s[j]
                counter -= 1
                break
    return s == s[::-1]

def area_container(a):
    left = 0
    right = len(a) - 1
    max_area = 0
    while right > left:
        pass
        length = right - left
        max_area = max(max_area, length * min(a[right], a[left]))
        if a[right] > a[left]:
            left += 1
        else:
            right -= 1
    return max_area
# Given a set of integers and a target value, find a full solution set of numbers that sum up to the target
# The same number in the set can be repeated multiple times
# eg. [1,2,3,4,5], target = 5. Solution set = (1,4),(2,3),(5),(1,2,2),(1,1,3),(1,1,1,1,1)
def combination(set, target):
    pass

# Given a sorted list and a value, return the index of the value in the list if it was found.
# If not found, return the index of the value if it was inserted into the list
def insert_pos(nums, target):
    pass
    left = 0
    right = len(nums) - 1

    while left <= right:
        mid = left + (right - left) // 2
        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    
    return left

array = [[1,5,7,8, 9, 11, 15, 22]]
print(11 // 2)

# Board is a list of nine lists. Blank spaces are periods.
# Brute force this first.
def valid_sud(board):
    pass
    horizontal = []
    vertical = []
    boxes = []

    # Loop through all the rows to find repeats
    for l in board:
        horizontal = []
        for n in l:
            if n != '.':
                horizontal.append(n)
            if n in horizontal:
                return False
    
    for i in range(0,9):
        vertical = []
        for j in range(0,9):
            n = board[j][i]
            if n != '.':
                vertical.append(n)
            if n in vertical:
                return False
    
    for i in range(0,9):
        boxes = []
        for j in range((i // 3) * 3, (i // 3) * 3 + 3):
            for k in range(0,3):
                pass
                if i in (0,3,6):
                    n = board[j][k]
                elif i in (1,4,7):
                    n = board[j][k + 3]
                else:
                    n = board[j][k + 6]
                
                if n != '.':
                    boxes.append(n)
                if n in boxes:
                    return False
    
    return True


# Given array of integers and a target value sorted in ascending order, find starting
# and ending position of target value.
# Must run in O(log n)
# e.g. nums = [3,3,4,4,5,5,5,6,6,6,7]; target = 6
# Output = [7,9]
def search_range(nums, target):
    pass
    left = 0
    right = len(nums) - 1

    while left <= right:
        mid = left + (right - left) // 2
        if nums[mid] == target:
            break
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    
    if nums[mid] != target:
        return [-1,-1]

    li, ri = mid, mid

    while li > 0:
        if nums[li - 1] == target:
            li -= 1
        else:
            break
    while ri < len(nums) - 1:
        if nums[ri + 1] == target:
            ri += 1
        else:
            break

    return [li, ri]



# Search a sorted array that is rotated at some pivot. Return index of element.
# No duplicates
# Solution must be O(log n)
# e.g. nums = [4,5,6,0,1,2,3]; target = 6
# Output = 2
def search_rot(nums, target):
    pass
    left, right = 0, len(nums) - 1

    while left <= right:
        mid = left + (right - left) // 2
        if nums[mid] == target:
            return mid
        
        if nums[left] <= nums[mid]:
            if nums[left] <= target < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if nums[mid] < target <= nums[right]:
                left = mid + 1
            else:
                right = mid - 1
    
    return -1

# Given a string only containing parantheses, brackets, or curly braces
# Returns true if string is valid. False otherwise.
# Input string is valid if all open brackets are closed by the same type of brackets
# and open brackets must be closed in the correct order
# Empty strings are also considered valid
# Valid: '()', '()[]{}', '{[]}'
# Invalid: '(]', '([)]'
def valid_par(s):
    pass
    par_d = {'(': ')', '[': ']', '{': '}'}
    listing = []

    for par in s:
        if par == '(' or par == '[' or par == '{':
            listing.append(par)
        else:
            if len(listing) == 0 or listing.pop(-1) != par_d[par]:
                return False
    
    return len(listing) == 0
