"""
Given a string s containing 0 to 9 digits, create the largest possible palindromic number using the string characters. It should not contain leading zeroes.

A palindromic number reads the same backward as forward.

If it's not possible to form such a number using all digits of the given string, you can skip some of them.

Examples
Example 1
Input: s = "323211444"
Expected Output: "432141234"
Justification: This is the largest palindromic number that can be formed from the given digits.

"""

class Solution:
    def largestPalindromicWrong(self, s):
        if s is None:
            return None
        
        count = {}

        for char in s:
            count[char] = count.get(char, 0) + 1
        
        # We can use all of the evens count
        # We can use one odd count
        # Sort take biggest evens count
        # Take biggest odds count 
        even = {}
        largest_odd = 0
        largest_char = ""
        odd = {}
        total_char = 0
        for char, num in count.items():
            if num % 2 == 0:
                even[char] = num
                total_char += num
            else:
                if num > largest_odd:
                    largest_odd = num
                    largest_char = char
        
        odd[largest_char] = largest_odd
        total_char += largest_odd
        longest_pali = [0] * total_char
        print(longest_pali)
        i = 0
        j = total_char - 1
        
        while i <= j:
            for num, numCount in even.items():
                while numCount > 0:
                    longest_pali[i] = num
                    longest_pali[j] = num
                    i += 1
                    j -= 1
                    numCount -= 2
            for num, numCount in odd.items():
                while numCount > 0 and i <= j:
                    longest_pali[i] = num
                    i += 1
        print(longest_pali)
        return "".join(longest_pali)
    
    def largestPalindromic(self, s):
        if s is None:
            return None
        
        freq = [0] * 10
        firstHalf = ""
        middle = -1
        




    
s = Solution()

pali = "444947137"
print(s.largestPalindromic(pali))