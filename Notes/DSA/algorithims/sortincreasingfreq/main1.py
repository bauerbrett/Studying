class Solution:
    def getFreq(self, arr):
        d = {}
        for num in arr:
            d[num] = d.get(num, 0) + 1
            
        return d

    def mergeSort(self, arr, freq):
        if len(arr) <= 1:
            return arr

        mid = (len(arr) // 2)
        left = arr[:mid]
        right = arr[mid:]

        self.mergeSort(left, freq)
        self.mergeSort(right, freq)
        self.merge(arr, left, right, freq)

    def merge(self, arr, left, right, freq):

        i = 0
        j = 0
        k = 0

        while i < len(left) and j < len(right):
            if freq[left[i]] < freq[right[j]] or freq[left[i]] == freq[right[j]] and left[i] > right[j]: # if freq are diff we take smaller, else we will take the larger if they are same beacuse we are going big to bottom
                arr[k] = left[i]
                i += 1
            else:
                arr[k] = right[j]
                j += 1
            k += 1


        """
        To show this piece on bottom here is breakdown
        Here's a concrete breakdown:
        Suppose you're merging:
        left  = [1, 4, 7]
        right = [2, 3]

        Compare 1 < 2 → put 1 in arr

        Compare 4 > 2 → put 2 in arr

        Compare 4 > 3 → put 3 in arr

        Now right is empty, but left still has [4, 7].
        Since left was already sorted, we can just tack [4, 7] onto the end of arr.

        Since we went from the end where there is only len == 1. we know the leftovers are in order. 7 could not be before 4

        Result:
        arr = [1, 2, 3, 4, 7]
        """
        while i < len(left): #If one of the index in arr runs out we just take the leftovers of the other and add them on. They are already sorted since they got built from one up.
            arr[k] = left[i]
            i += 1
            k += 1
        while j < len(right):
            arr[k] = right[j]
            j += 1 
            k += 1

s = Solution()
arr = [4, 4, 6, 2, 2, 2, 4, 5, 6, 10]
d = s.getFreq(arr)
s.mergeSort(arr, d)
print(arr)