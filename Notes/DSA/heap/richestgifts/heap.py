import heapq
import math


#just to show how easy it is in python because it handles everything under the hood

def gift_pile(gifts, k):
    # Convert gifts into a max heap by negating the values
    max_heap = [-gift for gift in gifts]
    heapq.heapify(max_heap)

    for _ in range(k):
        current = -heapq.heappop(max_heap)  # Pop max gift pile
        reduced = int(math.sqrt(current))
        heapq.heappush(max_heap, -reduced)  # Push reduced pile back

    return -sum(max_heap)  # Sum all remaining gifts (negate back)

# Example usage
gifts = [4, 9, 16]
k = 2
print(gift_pile(gifts, k))  # Output: 5