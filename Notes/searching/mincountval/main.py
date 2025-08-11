"""
Given two sorted arrays nums1 and nums2 containing integers only, return the smallest integer that 
appears in both arrays. If there isn't any integer that exists in both arrays, the function should return -1.

Examples
Example 1:

input: nums1 = [1, 3, 5, 7], nums2 = [3, 4, 5, 6, 8, 10]
expectedOutput: 3
Justification: Both arrays share the integers 3 and 5, but the smallest common integer is 3.

"""
class Solution:
    def binarySearch(self, arr, target):
        if len(arr) == 0:
            return False
        left = 0
        right = len(arr) - 1

        while left <= right:
            mid = (left + right) // 2
            if arr[mid] < target:
                left = mid + 1
            elif arr[mid] > target:
                right = mid - 1
            else:
                return True
        return False
    
    def getCommon(self, arr1, arr2):
        small = []
        large = []
        if len(arr1) > len(arr2):
            large = arr1
            small = arr2
        else:
            large = arr2
            small = arr1
            
        for num in small:
            if self.binarySearch(large, num):
                return num
        
        return -1
    

s = Solution()
l1 = [1,2,3,6]
l2 = [2,3,4,5]
print(s.getCommon(l1, l2))