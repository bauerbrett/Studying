from typing import List, Optional
from collections import deque

class TreeNode:
    def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

class Solution:
    def buildMaximumBinaryTree(self, nums: List[int]) -> Optional[TreeNode]:
        return self.buildTree(nums, 0, len(nums))

    def buildTree(self, nums: List[int], start: int, end: int) -> Optional[TreeNode]:
        # Base case: if the current subarray is empty, return null
        if start == end:
            return None

        # Find the index of the maximum value in the current subarray
        maxIndex = self.findIndexOfMaximum(nums, start, end)

        # Create a new TreeNode with the maximum value
        root = TreeNode(nums[maxIndex])

        # Recursively build the left subtree from the left subarray
        root.left = self.buildTree(nums, start, maxIndex)

        # Recursively build the right subtree from the right subarray
        root.right = self.buildTree(nums, maxIndex + 1, end)

        return root

    def findIndexOfMaximum(self, nums: List[int], start: int, end: int) -> int:
        maxIndex = start
        # Iterate through the subarray to find the maximum value's index
        for i in range(start, end):
            if nums[i] > nums[maxIndex]:
                maxIndex = i
        return maxIndex

    @staticmethod
    def serializeTree(root: Optional[TreeNode]) -> List[Optional[int]]:
        serialized = []
        if not root:
            return serialized

        queue = deque([root])

        while queue:
            current = queue.popleft()
            if current:
                serialized.append(current.val)
                queue.append(current.left)
                queue.append(current.right)
            else:
                serialized.append(None)

        # Remove trailing nulls for a cleaner representation
        while serialized and serialized[-1] is None:
            serialized.pop()

        return serialized


if __name__ == "__main__":
    solution = Solution()

    # Example 1
    nums1 = [4, 3, 1, 7, 0, 5]
    tree1 = solution.buildMaximumBinaryTree(nums1)
    output1 = Solution.serializeTree(tree1)
    print("Example 1 Output:", output1)
    # Expected Output: [7, 4, 5, None, 3, 0, None, None, 1]

    # Example 2
    nums2 = [1, 4, 3, 2]
    tree2 = solution.buildMaximumBinaryTree(nums2)
    output2 = Solution.serializeTree(tree2)
    print("Example 2 Output:", output2)
    # Expected Output: [4, 1, 3, None, None, None, 2]

    # Example 3
    nums3 = [7, 2, 5, 3, 9, 1]
    tree3 = solution.buildMaximumBinaryTree(nums3)
    output3 = Solution.serializeTree(tree3)
    print("Example 3 Output:", output3)
    # Expected Output: [9, 7, 1, None, 5, None, None, 2, 3]
