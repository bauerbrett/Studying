"""
Given string s, determine whether it's possible to make a given string palindrome by removing at most one character.

A palindrome is a word or phrase that reads the same backward as forward.

Examples
Example 1:

Input: "racecar"
Expected Output: true
Justification: The string is already a palindrome, so no removals are needed.
"""

class Solution:
    def validPalindrome(self, s):
        if len(s) <= 1:
            return True
        
        l = 0
        r = len(s) - 1
        while l <= r:
            if s[l] == s[r]:
                l += 1
                r -= 1
            else:
                return self.helperCheck(s, l + 1, r) or self.helperCheck(s, l, r - 1)
        return True

    def helperCheck(self, s, l, r):
        if len(s) <= 1:
            return True
        while l <= r:
            if s[l] == s[r]:
                l += 1
                r -= 1
            else:
                return False
        return True
    
s = Solution()
p = "abcdef"
print(s.validPalindrome(p))



