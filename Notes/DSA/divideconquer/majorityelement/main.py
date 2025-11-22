"""
Given an array nums having an n elements, identify the element that appears the majority of the time, meaning more than n/2 times.

Examples
Example 1:

Input: [1, 2, 2, 3, 2]
Expected Output: 2
Justification: Here, '2' appears 3 times in a 5-element array, making it the majority element.
"""

class Solution:
    def majorityelem(self, arr):
        if len(arr) == 1:
            return arr[0]
        
        mid = len(arr) // 2

        elem1 = self.majorityelem(arr[:mid])
        elem2 = self.majorityelem(arr[mid:])

        if elem1 == elem2:
            return elem1
        
        leftCount = arr.count(elem1)
        rightCount = arr.count(elem2)

        return elem1 if leftCount > rightCount else elem2
    
"""
General principle

At the base: you just get raw number.

At each merge step:

If they’re equal → keep it.

If different → count them in the current whole slice → keep the one that appears more.

This guarantees weaker candidates get eliminated as you bubble back up.

Only the “winner” survives to the top unless it gets beaten by a stronger element higher up.
"""

s = Solution()

test = [9, 9, 1, 1, 9, 1, 9, 9]

print(s.majorityelem(test))