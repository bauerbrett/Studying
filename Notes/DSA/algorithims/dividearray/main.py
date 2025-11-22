"""
You are given an array nums containing n integers and a positive integer k.

Divide the nums into arrays of size 3 such that it satisfies the below conditions:

Each element of nums should be in exactly one array.
The difference between any two elements of a single array should be less than or equal to k.
Return a 2D array of these subarrays. If no such division is possible, return an empty array.

Examples
Example 1:

Input: nums = [2, 6, 4, 9, 3, 7, 3, 4, 1], k = 3
Expected Output: [[1,2,3],[3,4,4],[6,7,9]]
Explanation: The three groups [1, 2, 3], [3, 4, 4] and [6, 7, 9] all have elements with differences ≤ k. 
For the group [1, 2, 3] the maximum difference between elements is 2. For the group [3, 4, 4], 
the maximum difference is 1, and for [6, 7, 9], it's 3, all of which satisfy the condition.
"""

class Solution:
    def mergeSort(self, arr):
        if len(arr) <= 1:
            return arr
        
        mid = len(arr) // 2 # we need an index not an element from array like in quicksort
        left = arr[:mid]
        right = arr[mid:]
        self.mergeSort(left)
        self.mergeSort(right)
        self.merge(left, right, arr)
    
    def merge(self, left, right, arr):
    
        i, j, k = 0, 0, 0
        while i < len(left) and j < len(right):
            if left[i] < right[j]:
                arr[k] = left[i]
                i += 1
            else:
                arr[k] = right[j]
                j += 1
            k += 1

        while i < len(left):
            arr[k] = left[i]
            i += 1
            k += 1
        while j < len(right):
            arr[k] = right[j]
            j += 1
            k += 1
    def diff(self, a, b):
        if a < b:
            return a - b
        return b - a

    def divideArray(self, arr, k):
        l = []
        self.mergeSort(arr)
        new = []
        for num in arr:
            new.append(num)
            if len(new) == 3:
                if (new[-1] - new[0]) > k:
                    return []
                else:
                    l.append(new)
                    new = []
        return l





#l = [2, 6, 4, 9, 3, 7, 3, 4, 1]
s = Solution()
l = [1,3,4,8,7,9,3,5,1]
#s.mergeSort(l)
#print(l)
print(s.divideArray(l, 2))