"""
Given divisor1, divisor2, uniqueCnt1, and uniqueCnt2 integers, find the smallest 
possible maximum integer that could be present in either array after they are filled according to the below conditions.

You can take two arrays arr1 and arr2 which are initially empty.
arr1 contains total uniqueCnt1 different positive integers, each of them is not divisible by divisor1.
arr2 contains total uniqueCnt2 different positive integers, each of them is not divisible by divisor2.
There are no common integers in both arrays.
Examples
Example 1:

Input: uniqueCnt1 = 2, divisor1 = 2, uniqueCnt2 = 2, divisor2 = 3
Expected Output: 4
Explanation: The optimal arrays could be arr1 = [1, 3] (numbers not divisible by 2) and arr2 = [2, 4] (numbers not divisible by 3). 
The maximum number among both arrays is 4.
"""
class Solution:
    def minMax(self, div1, div2, uni1, uni2):
        return 0