"""
Given an array nums sorted in increasing order, return the maximum between the count of positive integers and the count of negative integers.

Note: 0 is neither positive nor negative.

Examples
Example 1:

Input: nums = [-4, -3, -1, 0, 1, 3, 5, 7]
Expected Output: 4
Justification: The array contains three negative integers (-4, -3, -1) and four positive integers (1, 3, 5, 7). 
The maximum count between negatives and positives is 4.

Since its sorted, just use binary and if < 0 neg + 1 and if > 0 positive + 1
"""

class Solution:

    def maximumCount(self, arr):

        if len(arr) == 0:
            return 0
    
        positive = 0
        neg = 0
        for num in arr:
            if num < 0:
                neg += 1
            if num > 0:
                positive += 1
        return max(positive, neg)
    
    def binarySearch(self, arr):
        if len(arr) == 0:
            return 0
        
        left = 0
        right = len(arr) - 1

        pos = 0
        neg = 0
        while left <= right:
            mid = (left + right) // 2
            if arr[mid] > 0:
                pos = len(arr) - mid # From mid to end
                right = mid - 1
            else:
                left = mid + 1
        left, right = 0, len(arr) - 1
        while left <= right:
            mid = (left + right) // 2
            if arr[mid] < 0:
                neg = mid + 1 # 	From 0 to mid
                left = mid + 1
            else:
                right = mid - 1
        print(neg, pos)
        return max(neg, pos)
        
    
s = Solution()
l =  [-4, -3, -1, 0, 1, 3, 5, 7]
print(s.binarySearch(l))