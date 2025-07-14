import heapq
from collections import Counter

"""
 frequency (-freq) is placed before the character in the tuple. This matters because Python's heapq is a min-heap and it sorts tuples by their first element by default.

The heap uses the first element of each tuple to determine priority.

By putting -freq first, the heap treats the highest frequencies (largest negative numbers in magnitude) as smallest values.

This effectively turns the min-heap into a max-heap based on frequency.

If you put the character first, the heap would prioritize characters lexicographically, not by frequency.

So, frequency comes first to ensure the heap orders elements by frequency (descending), and the character is the tiebreaker when frequencies are equal.
"""

l = "trersesess"
freq = Counter(l)

maxheap = [(-freq, char) for char, freq in freq.items()]
heapq.heapify(maxheap)

result = []

while maxheap: 
    freq, char = heapq.heappop(maxheap)
    result.append(char * -freq)

print("".join(result))
