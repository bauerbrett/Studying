"""
Problem Statement
You are given an array apple of size n, where the apple[i] represents the number of apples in ith pack. 
You are also given an array capacity of size m, where capacity[j] is a number of apples that can be stored in the jth box.

Return the minimum number of boxes you need to use to put these all n packs of apples into boxes.

Note: You are allowed to distribute apples from the same pack into different boxes.

Examples
Example 1:

Input: apple = [2, 3, 1], capacity = [4, 2, 5, 1]
Expected Output: 2
Explanation: Box 1 can take apples from packs 1 and 2 partially (totaling 5 apples), and Box 2 can take the rest of 2 apples.
"""

class Solution:
    def appleSort(self, pack, box):
       
       appleSum = sum(pack)
       sortedBoxes = s.quickSort(box)
       capacitySum = 0
       boxCount = 0

       # Add teh capacities until we get to the apple sum
       for box in sortedBoxes:
           capacitySum += box
           boxCount += 1

           if capacitySum >= appleSum:
               return boxCount # If we get to the point the capacity is == sum or > we return the box count
       return -1 # If the capacity sum never gets to apple sum we cannot fit all the apples so we return -1

    # This is easiest way to do quick sort. There is another way with partitions I believe but this is easier
    def quickSort(self, arr):
        if len(arr) <= 1:
            return arr
        pivot = arr[0]
        left = [x for x in arr[1:] if x > pivot]
        middle = [x for x in arr if x == pivot]
        right = [x for x in arr[1:] if x < pivot]

        return self.quickSort(left) + middle + self.quickSort(right) # recursively find the next left and right. They wil keep adding each other together and return up stack

s = Solution()

apple = [1, 2, 5, 6]
capacity = [2, 3, 7, 4, 5, 2, 4]
print(s.appleSort(apple, capacity))