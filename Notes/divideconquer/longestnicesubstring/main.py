"""
Given a string str, return the longest nice substring of a given string.

A substring is considered nice if for every lowercase letter in the substring, its uppercase counterpart is also present, and vice versa.

If no such string exists, return an empty string.

Examples
Example 1:

Input: "BbCcXxY"
Expected Output: "BbCcXx"
Justification: Here, "BbCcXx" is the longest substring where each letter's uppercase and lowercase forms are present.


Scan the string left to right
As long as every character has both cases present, we keep going.
Hit a “bad” character (one missing its upper/lower partner):
That means the whole string can’t be nice.
And — importantly — any substring that includes that bad character can’t be nice either.
Split at that bad character:
Left side = everything before it
Right side = everything after it
Recurse
Check the left substring
Check the right substring
Compare
Take the longer of left vs right and return it
"""

class Solution:
    def longestSubstring(self, s):
        # Base case: If the string has less than 2 characters, it cannot be "nice"
        if len(s) < 2:
            return ""
        
         # Create a set to store all characters in the string
        check = set(s)

        for i, char in enumerate(s):
            if char.upper() in check and char.lower() in check:
                continue

            left = s[:i] # left part of string to check 
            right = s[i + 1:] # right part, need to get past failed char so we add 1

            # If the character doesn't satisfy the "nice" condition, split the string
            # Recursively find the longest nice substring in the left and right parts
            # This will check each substring until we hit < 2 which it will then return "" and compare and return the greater one
            substr1 = self.longestSubstring(left) 
            substr2 = self.longestSubstring(right)

            # Take the bigger
            return substr1 if len(substr1) >= len(substr2) else substr2
        return s

s = Solution()
word = "BbCcXxy"
print(s.longestSubstring(word))

        
        
