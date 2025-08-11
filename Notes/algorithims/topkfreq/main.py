"""
Given an unsorted array of numbers, find the top K frequently occurring numbers in it.

Example 1:

Input: [1, 3, 5, 12, 11, 12, 11], K = 2
Output: [12, 11]
Explanation: Both '11' and '12' apeared twice.

Create freq map and sort by freq. After sort grab the first num and the second num that isnt the first (if dupes of 0 exist)
"""

class Solution:
    def getFreq(self, arr):
        d = {}
        for num in arr:
            d[num] = d.get(num, 0) + 1
        return d
    
    def quickSort(self, arr, freq):
        if len(arr) <= 1:
            return arr
        
        pivot = arr[0] # Since we take 0 always use 1 in the list comprehension because 0 is already used
        left = [x for x in arr[1:] if freq[x] > freq[pivot] or (freq[x] == freq[pivot] and x > pivot)]
        middle = [x for x in arr if x == pivot]
        right = [x for x in arr[1:] if freq[x] < freq[pivot] or (freq[x] == freq[pivot] and x < pivot)]
        return self.quickSort(left, freq) + middle + self.quickSort(right, freq)
    
    def topKFrequent(self, arr, k):
        freq = self.getFreq(arr)
        arr = self.quickSort(arr, freq)
        print(arr)

        l = [arr[0]]
        biggest = arr[0]
        for num in arr:
            if num != biggest:
                l.append(num)
                biggest = num
            if len(l) == k:
                return l


s = Solution()
l = [1, 3, 5, 12, 11, 12, 11]
freq = s.getFreq(l)
print(s.quickSort(l, freq))
print(s.topKFrequent(l, 2))