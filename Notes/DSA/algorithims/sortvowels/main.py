
"""
Given a string s, return an updated string t such that all consonants in the string s stay in their 
original positions while any vowels in the string are reordered according to their ASCII values.

The vowels are 'A', 'E', 'I', 'O', and 'U'. These vowels can appear in lowercase or uppercase. 
All other letters except vowels are consonants.

Examples
Example 1:

Input: "gamE"
Expected Output: "gEma"
Justification: The vowels in "gamE" are 'a' and 'E'. Sorting these by ASCII values, 'E' comes before 'a'. 
Hence, they are placed in the order 'E', 'a' in the output, while consonants remain unchanged.
"""

class Solution:
    def createVowelMap(self, str):
        d = {}
        for i, c in enumerate(str):
            d[c] = i
        return d
    
    def findVowels(self, str, map):
        l = []
        for c in str:
            if c in map:
                l.append(c)
        return l
    
    def sortVowels(self, arr, map):
        if len(arr) <= 1:
            return arr
        
        pivot = arr[0]
        left = [x for x in arr[1:] if map[x] < map[pivot]]
        middle = [x for x in arr if x == pivot]
        right = [x for x in arr[1:] if map[x] > map[pivot]]
        return self.sortVowels(left, map) + middle + self.sortVowels(right, map)
    
    
    def sortString(self, word, vowels, map):
        vowelIdx = 0
        newWord = ""
        for i, c in enumerate(word):
            if c in map:
                newWord += vowels[vowelIdx]
                vowelIdx += 1
            else:
                newWord += c
        return newWord
    
    # For practice here is merge sort
    def mergeSort(self, arr, map):
        if len(arr) <= 1:
            return arr
        
        mid = (len(arr) // 2)
        left = arr[:mid]
        right = arr[mid:]

        self.mergeSort(left, map)
        self.mergeSort(right, map)
        self.merge(arr, left, right, map)
    
    def merge(self, arr, left, right, map):
        i, j, k = 0, 0, 0
        while i < len(left) and j < len(right):
            if map[left[i]] < map[right[j]]:
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
            
        


s = Solution()

vowel = "AEIOUaeiou"
word = "aEiOu"
d = s.createVowelMap(vowel)
l = s.findVowels(word, d)
l2 = s.findVowels(word, d)
sort = s.sortVowels(l, d)
s.mergeSort(l2, d)
print(d)
print(l)
print(sort)
print(l2)
print(s.sortString(word, sort, d))

