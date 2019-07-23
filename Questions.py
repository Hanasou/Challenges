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
