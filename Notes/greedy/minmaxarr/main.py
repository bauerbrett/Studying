"""
Determine the minimum number of deletions required to remove the smallest and the largest elements from an array of integers.

In each deletion, you are allowed to remove either the first (leftmost) or the last (rightmost) element of the array.

Examples
Example 1:

Input: [3, 2, 5, 1, 4]
Expected Output: 3
Justification: The smallest element is 1 and the largest is 5. Removing 4, 1, and then 5 (or 5, 4, and then 1) in three moves is the most efficient strategy.
"""

class Solution:
    def minmaxArr(self, arr):
        if arr is None:
            return None
        
        minNum = 20 ** 20
        maxNum = 0
        minIdx = 20 ** 20
        maxIdx = 0

        for i, val in enumerate(arr):
            if val < minNum:
                minNum = val
                minIdx = i
            if val > maxNum:
                maxNum = val
                maxIdx = i

        """
        To find teh min numbers deleted, we need to find the num by taking all possible options and taking the smallest. There are 4 ways to remove elements from min and max.
        So the options are:

        Remove both from the left.

        Remove both from the right.

        Remove min from left + max from right.

        Remove max from left + min from right.

        We hadd the possible ways down I just didnt think to try all and just take the smallest. I thought we just had to find the right one before taking it.
        """
        n = len(arr)
        res = min(
            max(minIdx, maxIdx) + 1,   # both from left | take the max index , add 1 because we start at 0 for these 
            n - min(minIdx, maxIdx),   # both from right | take the min index minus the length of arr - that index to find the count of elements
            (minIdx + 1) + (n - maxIdx), # min left + max right | now we basically combine the other two together since small and large coule be on oposite ends of arr
            (maxIdx + 1) + (n - minIdx)  # max left + min right
        )
        return res
    
s = Solution()

arr = [3, 2, 5, 1, 4]

print(s.minmaxArr(arr))