"""
Given a 2D grid of size m x n matrix containing integers, and integer target, return true if target value exists in the matrix. Otherwise, return false.

The matrix has the following properties:

Values in each column are sorted in non-decreasing order from top to bottom.
Values in each row are sorted in non-decreasing order from left to right.
Examples
Example 1:

Input: target = 5, matrix =
[[1,2,3],
 [4,5,6],
 [7,8,9]]
Expected Output: true
Justification: The number 5 is located in the second row and second column of the matrix, thus the output is true.
"""
class Solution:

    def binarySearch(self, target, arr):
        l, r = 0, len(arr) - 1
        while l <= r:
            mid = (r + l) // 2
            if arr[mid] < target:
                l = mid + 1
            elif arr[mid] > target:
                r = mid - 1
            else:
                return True
        return False

    #Wrong
    def searchMatrix2(self, matrix, target):
        l, r = 0, len(matrix) - 1
        while l <= r:
            mid = (r + l) // 2
            #print(mid)
            if matrix[mid][0] == target or matrix[mid][-1] == target:
                return True
            if matrix[mid][0] < target and matrix[mid][-1] > target:
                if self.binarySearch(target, matrix[mid]) is False:
                    return self.binarySearch(target, matrix[mid-1]) or self.binarySearch(target, matrix[mid+1])
                else:
                    return True
            if matrix[mid][0] > target:
                r = mid - 1
            if matrix[mid][0] < target:
                l = mid + 1
        return False
    def searchMatrix1(self, matrix, target):
        for arr in matrix:
            if self.binarySearch(target, arr):
                return True
        return False
    
    def searchMatrix(self, matrix, target):
        if not matrix or not matrix[0]:
            return False
        
        row = 0
        col = len(matrix[0]) - 1 # Start in top right an work in 
        
        # Start from top right corner and move towards the target
        while row < len(matrix) and col >= 0:
            if matrix[row][col] == target:
                # Target found
                return True
            elif matrix[row][col] < target: # If num is < target we need to move row down 
                # Move down
                row += 1
            else: # Else if it is greater, keep moving column to the left
                # Move left
                col -= 1
        
        # Target not found
        return False

s = Solution()
m = [[1,4,7,11,15],
     [2,5,8,12,19],
     [3,6,9,16,22],
     [10,13,14,17,24],
     [18,21,23,26,30]]
print(s.searchMatrix(m, 5))
