"""
Given a collection of pairs where each pair contains two elements [a, b] and a < b, find the maximum length of a chain you can form using pairs.

A pair [a, b] can follow another pair [c, d] in the chain if b < c.

You can select pairs in any order and don't need to use all the given pairs.

Examples
Example 1:

Input: [[1,2], [3,4], [2,3]]
Expected Output: 2
Justification: The longest chain is [1,2] -> [3,4]. The chain [1,2] -> [2,3] is invalid because 2 is not smaller than 2.


Basically, we need to sort by the second value in pairs. That way if 1,100 is a pair it is last. We know now that if we compare current chain end to 
next possible chain start we can see if the chain fits becaues they are sorted. If the next chain beginning is not bigger move to the next chain. If it is make that
end on the chain the current end to check the next chain.

"""
# [[1, 2], [2, 3], [3, 5], [4, 5], [5, 6], [7, 8]]
class Solution:
    def findLongestChain(self, pairs):
        if len(pairs) == 0:
            return 0
        pairs.sort(key=lambda x: x[1])
        # or pairs = sorted(pairs, key=lambda x: x[1])

        
        count = 0
        current_end = -20**20
        for l in pairs:
            if l[0] > current_end:
                count += 1
                current_end = l[1]
            else:
                continue

        return count

       






        

s = Solution()

l = [[7,8], [5,6], [1,100], [3,5], [4,5], [2,3]]
#print(sorted(l))
print(s.findLongestChain(l))