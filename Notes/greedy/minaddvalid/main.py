"""
Given a string str containing '(' and ')' characters, find the minimum number of parentheses that 
need to be added to a string of parentheses to make it valid.

A valid string of parentheses is one where each opening parenthesis '(' has a corresponding closing parenthesis ')' 
and vice versa. The goal is to determine the least amount of additions needed to achieve this balance.

Examples
Example 1:

Input: "(()"
Expected Output: 1
Justification: The string has two opening parentheses and one closing parenthesis. Adding one closing parenthesis at the end will balance it.
"""

class Solution:
    # First solution with hashmap 
    def minAddToMakeValid(self, s):
        if s is None:
            return None
        
        count = 0
        opens = {"(": 0}
        for char in s:
            if char == ")" and opens["("] == 0:
                count += 1
            elif char == ")" and opens["("] > 0:
                opens["("] -= 1
                count -= 1
            else:
                opens[char] += 1
                count += 1

        return count
    
    def secondSolution(self, s):
        if s is None:
            return None
        
        balance = 0
        count = 0
        for char in s:
            if char == "(":
                balance += 1
            elif char == ")" and balance > 0:
                balance -= 1
            else:
                count += 1
        return count + balance # We add these because balance == the unused ( and count == the unused ) because there was no ( to match it.

s = Solution()

chars = "))(("

"""
( = 1
( = 2
) = 1
( = 2
) = 1
) = 2
"""

print(s.minAddToMakeValid(chars))
print(s.secondSolution(chars))
            