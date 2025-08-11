"""
You are given an nums array containing n integers and an integer k. In a single operation, 
you can choose any index i and increment the nums[i] by 1.

Return the maximum possible frequency of any element of nums after performing at most k operations.

Example 1:

Input: nums = [1, 2, 3], k = 3
Expected Output: 3
Explanation: We can increment the number 1 two times and 2 one time. The final array will be 
[3, 3, 3]. Now, the number 3 appears 3 times in the array [3, 3, 3].
"""

class Solution:
    def quickSort(self, arr):
        if len(arr) <= 1:
            return arr
        
        pivot = arr[0]
        left = [x for x in arr[1:] if x > pivot]
        mid = [x for x in arr if x == pivot]
        right = [x for x in arr[1:] if x < pivot]
        return self.quickSort(left) + mid + self.quickSort(right)
    
    #Brute Force
    def maxFrequency(self, arr, k):
        arr = self.quickSort(arr)
        if len(arr) == 0:
            return None
        maxCount = 0
        for i, num in enumerate(arr):
            count = 0
            j = k
            start = arr[i]
            for num in arr[i:]:
                if num == start:
                    count += 1
                else:
                    diff = start - num 
                    j -= diff
                    if j >= 0:
                        count += 1
            #print(count)
            maxCount = max(maxCount, count)
        return maxCount
    
    #Binary search method - HARD
    # Method to find the maximum length of subarray that can be made equal to `elements[index]` using at most `maxOperations`.
    def findMaxSubarrayLength(self, index, maxOperations, elements, cumulativeSum):
        target = elements[index]  # The target number we want to make others equal to.
        start = 0  # Start of the search range.
        end = index  # End of the search range, we consider subarrays ending at `index`.
        bestLength = index  # This will store the best start position for the longest valid subarray.
        
        while start <= end:
            mid = (start + end) // 2  # Midpoint of the current search range.
            count = index - mid + 1  # Number of elements from `mid` to `index`.
            requiredSum = count * target  # If all elements are `target`, this is the total they would sum to.
            existingSum = cumulativeSum[index] - (cumulativeSum[mid - 1] if mid > 0 else 0)  # Current sum from `mid` to `index`.
            operationsRequired = requiredSum - existingSum  # How many increments are needed to make all equal to `target`.
            
            if operationsRequired > maxOperations:
                start = mid + 1  # If more operations are required than allowed, move the start up.
            else:
                bestLength = mid  # Update bestLength as this is a valid subarray.
                end = mid - 1  # Try for a longer valid subarray.
        
        return index - bestLength + 1  # Return the length of the longest valid subarray.
    
    # Method to calculate the maximum frequency of any element after at most `maxOperations` increments.
    def maxFrequency(self, elements, maxOperations):
        elements.sort()  # Sort the array to facilitate the equalization to the highest element.
        cumulativeSum = [0] * len(elements)  # Array to store cumulative sums.
        cumulativeSum[0] = elements[0]  # Initialize the first element of cumulative sum.
        
        for i in range(1, len(elements)):
            cumulativeSum[i] = elements[i] + cumulativeSum[i - 1]  # Build the cumulative sum array.
        
        maximumFrequency = 0
        for i in range(len(elements)):
            maximumFrequency = max(maximumFrequency, self.findMaxSubarrayLength(i, maxOperations, elements, cumulativeSum))  # Compute max frequency for each end position.
        
        return maximumFrequency  # Return the maximum frequency found.

    #Sliding window
    def slidingWindow(self, nums, k):
        nums.sort()
        l, r = 0, 0
        total = 0
        res = 0

        while r < len(nums):
            total += nums[r]

            while (r - l + 1) * nums[r] > total + k:
                total -= l
                l += 1 # Left always needs to move right because the arr is sorted. We know if we are on the same r we need to get to a point where the window is smaller than total and k
                # we also know even when we move r to right l needs to stay where it is at because r is going to grow which means 0 would fail and it would just end up where it was at originally if we
                # just kept l where it was at

            """
            ✅ Why We Don’t Reset l Every Time:
            When we slide the right pointer r forward (i.e. expand the window), we may end up in a situation where the new window is too expensive — meaning it takes more than k operations to make all elements equal to nums[r].

            Now, if we always reset l to 0, we'd be wasting time rechecking the same invalid windows again and again. But instead, by keeping l where it was, we’re saying:

            “Hey, I already know the window starting at l is either valid or very close to being valid. So let’s just shrink it as needed.”

            This way, l naturally adjusts to the correct position, and we never go backward or waste work.
            """
            res = max(res, (r - l + 1))
        return res



s = Solution()
l = [1,2,4]
print(s.quickSort(l))
print(s.maxFrequency(l, 5))
print(s.slidingWindow(l, 5))
