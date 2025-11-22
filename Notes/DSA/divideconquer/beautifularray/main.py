class Solution:
    def beautifulArray(self, N: int) :
        # Base case: if N is 1, return an array with a single element [1]
        if N == 1:
            return [1]

        # Recursively construct the beautiful array for odd and even parts
        odd = self.beautifulArray((N + 1) // 2)
        even = self.beautifulArray(N // 2)

        # Transform and populate the odd and even parts in the result
        # Each odd element is 2*value - 1 and each even element is 2*value
        return [2 * x - 1 for x in odd] + [2 * x for x in even]


# Testing the solution with examples
solution = Solution()
print("Beautiful Array for N = 4:", solution.beautifulArray(4))
print("Beautiful Array for N = 3:", solution.beautifulArray(3))
print("Beautiful Array for N = 8:", solution.beautifulArray(8))