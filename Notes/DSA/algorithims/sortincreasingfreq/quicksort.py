class Solution:

    def getFreq(self, arr):
        d = {}
        for num in arr:
            d[num] = d.get(num, 0) + 1
        return d
    
    def quickSort(self, arr, freq):
        if len(arr) <= 1:
            return arr
        
        pivot = arr[0]
        left = [x for x in arr[1:] if freq[x] < freq[pivot] or (freq[x] == freq[pivot] and x > pivot)]
        middle = [x for x in arr if x == pivot] # This is different than the org quick sort because we need to find dupes of pivot. In future problems for this just use this
        right = [x for x in arr[1:] if freq[x] > freq[pivot] or (freq[x] == freq[pivot] and x < pivot)]
        return self.quickSort(left, freq) + middle + self.quickSort(right, freq)
    

s = Solution()

l = [1, 3, 5, 7, 4, 2, 2, 4, 10]
d = s.getFreq(l)

print(s.quickSort(l, d))