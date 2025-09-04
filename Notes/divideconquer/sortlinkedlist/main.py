"""
Given a head of the linked list, return the list after sorting it in ascending order.

Examples
Example 1:

Input: [3, 1, 2]
Expected Output: [1, 2, 3]
Justification: The list is sorted in ascending order, with 1 coming before 2, and 2 before 3.

To split in have we need to go through the linked list nodes next one at a time for slow and 2 at a time for fast
this way slow will be middle and fast will be tail at the end. This basically splits in half 

After we basically do a merge sort with the nodes
"""

class Node:
    def __init__(self, x):
        self.value = x
        self.next = None
class Solution:
    def sortList(self, head):
        if not head or not head.next:
            return head
        
        #Same as in regular merge sort, we need to find first half and second half and merge them
        slow, fast, prev = head, head, None
        while fast and fast.next:
            slow, fast, prev = slow.next, fast.next.next, slow # Prev is the trailer 
        prev.next = None

        l1 = self.sortList(head) # First half
        l2 = self.sortList(slow) # Second half

        return self.merge(l1, l2)
    
    def merge(self, l1, l2):
        dummy = tail = Node(0) # Need a dummy for the beginning and the tail for the end. The tail is what keeps gets pushed back and is on the end.

        # This is the same as a regular merge sort but with a linked list. 
        while l1 and l2:
            if l1.val < l2.val:
                tail.next = l1
                l1 = l1.next 
            else:
                tail.next = l2
                l2 = l2.next
            tail = tail.next

        """
        Why does this ending tail.next work below?

        Because both input lists are already sorted.

        If one list is exhausted, everything left in the other is already greater than what we’ve added so far.

        So we don’t need to compare anymore — just tack it all on.

        So to answer your question directly: ✅ Yes, it means if any nodes are left over, they all get linked onto the tail of the merged list.
        """
        tail.next = l1 if l1 else l2
        return dummy.next
    """
    Mini-example:
    Merging l1 = 1 -> 4 and l2 = 2 -> 3

    Start: dummy -> None, tail = dummy

    Compare 1 vs 2: attach 1 → dummy -> 1, tail = 1

    Compare 4 vs 2: attach 2 → dummy -> 1 -> 2, tail = 2

    Compare 4 vs 3: attach 3 → dummy -> 1 -> 2 -> 3, tail = 3

    Now only 4 is left, so: tail.next = 4

    Final result: 1 -> 2 -> 3 -> 4 (starting from dummy.next).
    """
        

