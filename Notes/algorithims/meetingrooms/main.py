"""
Given an array of meeting intervals where intervals[i] = [starti, endi], return the minimum number of meeting rooms needed so that no meetings overlap.

Examples
Example 1:
Input: intervals = [[10, 15], [20, 25], [30, 35]]
Expected Output: 1
Justification: There are no overlapping intervals in the given list. So, only 1 meeting room is enough for all the meetings.

"""

class Solution:
    def mergeSort(self, arr):
        if len(arr) <= 1:
            return arr
        
        mid = len(arr) // 2
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

    def meetingMap(self, arr):
        """method to create a map that marks a meeting time as start or end"""
        d = {}
        for slice in arr:
            start, end = slice[0], slice[-1]
            # This is weird but basically if we ever see teh value in teh map we need to make it say end. The reason is if 
            # to meeting cross over at the same time we also take the end and -1 from count so we need it to say end no matter what.
            # To do this we check if it is in the dict and if it is make it end and make the other index what it should be like we regulary would.
            if start in d:
                d[start] = "end"
                d[end] = "end"
            elif end in d:
                d[end] = "end"
                d[start] = "start"
            else:
                d[start] = "start"
                d[end] = "end"
        return d
         


    def calcMeetings(self, arr):
        meetingMap = self.meetingMap(arr) # Create the map to track if a time is start or end
        print(meetingMap)
        l = []
        for slice in arr: # Create teh new slice that is 1d and not 2d
            l.append(slice[0])
            l.append(slice[-1])
        s.mergeSort(l) # Sort the new arr. This just helps get the numbers that start in front
        largest = 0 # Need this because count will go to 0 by the end of the loop. This will let us know the max meeting rooms used
        count = 0 # Count meetings as we go through 
        for num in l:
            if meetingMap[num] == "start": # add one if start
                count += 1
            if meetingMap[num] == "end": # - 1 if end
                count -= 1
            if count > largest: # if > than largest make largest that count
                largest = count

        return largest
s = Solution()
test = [0 , 10, 4, 5, 9]
s.mergeSort(test)
print(test)

l = [[10, 15], [15, 25], [10, 25]]
print(s.calcMeetings(l))