"""
Given an array nums containing the integers, return the resultant array after sorting it in increasing order 
based on the frequency of the values. If two numbers have the same frequency, they should be sorted in descending numerical order.

Examples
Example 1:

Input: nums = [4, 4, 6, 2, 2, 2]
ExpectedOutput: [6, 4, 4, 2, 2, 2]
Justification: Here, '6' appears once, '4' appears twice, and '2' appears three times. 
Thus, numbers are first sorted by frequency and then by value when frequencies tie.
"""

class Solution:
    # Method to sort the array based on frequency
    def frequencySort(self, nums):
        # Step 1: Count frequency
        freqMap = {}
        for num in nums:
            freqMap[num] = freqMap.get(num, 0) + 1
        
        # Step 2: Use merge sort for custom sorting
        self.mergeSort(nums, 0, len(nums) - 1, freqMap)
        
        return nums

    # Merge Sort Function
    def mergeSort(self, nums, left, right, freqMap):
        if left < right:
            mid = (left + right) // 2
            self.mergeSort(nums, left, mid, freqMap)
            self.mergeSort(nums, mid + 1, right, freqMap)
            self.merge(nums, left, mid, right, freqMap)

    # Merge Function
    def merge(self, nums, left, mid, right, freqMap):
        # Temporary arrays to hold the values
        leftArray = nums[left:mid + 1]
        rightArray = nums[mid + 1:right + 1]

        i, j, k = 0, 0, left
        while i < len(leftArray) and j < len(rightArray):
            # Comparing frequencies and values for sorting
            if freqMap[leftArray[i]] < freqMap[rightArray[j]] or (freqMap[leftArray[i]] == freqMap[rightArray[j]] and leftArray[i] > rightArray[j]):
                nums[k] = leftArray[i]
                i += 1
            else:
                nums[k] = rightArray[j]
                j += 1
            k += 1

        # Copy the remaining elements of leftArray, if any
        while i < len(leftArray):
            nums[k] = leftArray[i]
            i += 1
            k += 1

        # Copy the remaining elements of rightArray, if any
        while j < len(rightArray):
            nums[k] = rightArray[j]
            j += 1
            k += 1
    
    #Build in sort
    def ffrequencySort(self, nums):
        # Step 1: Count frequencies
        freq = {}
        for num in nums:
            freq[num] = freq.get(num, 0) + 1

        # Step 2: Sort using frequency asc, value desc
        nums.sort(key=lambda x: (freq[x], -x))
        return nums
       

s = Solution()

print(s.frequencySort([4, 4, 6, 2, 2, 2]))     # [6, 4, 4, 2, 2, 2]
print(s.frequencySort([0, -1, -1, -1, 5]))     # [5, 0, -1, -1, -1]
print(s.frequencySort([10, 10, 10, 20, 20, 30]))  # [30, 20, 20, 10, 10, 10]



