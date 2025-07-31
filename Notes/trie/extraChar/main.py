class Solution:
    def __init__(self):
        self.memo = None # Memoization array to store results for each index
        self.wordSet = None # Set for quick dictionary lookup

    def minExtraCha(self, s, dictionary):
        length = len(s)
        self.memo = [None] * length # Initialize memoization array
        self.wordSet = set(dictionary) # Populate Set with dictionary words
        return self.solve(0, length, s)
    
    # l = 5
    #s = "abcde"
    #dictionary = ["a", "bc"]
    #memo = []

    '''
    🧠 How You Should Think:

    Always consider the default (bad) case: 1 extra char

    Loop forward from current position:

    If substring matches word, try jumping over it with no cost

    Take the minimum of all those paths

    Use memo to avoid recomputation

    🧱 Visual Example:
    Let’s say:

    s = "applepie"
    dictionary = ["apple", "pie"]
    Call stack looks like:

    dfs(0):

    Try "a" → not in dict

    ...

    Try "apple" → in dict → go to dfs(5)

    dfs(5) tries "p", "pi", "pie" → match → go to dfs(8)

    dfs(8) = end of string → return 0

    So dfs(5) = 0

    So dfs(0) = 0

    No extra chars!
'''

    def solve(self, index, length, s):
        if index == length:
            return 0
        
        if self.memo[index] is not None:
            return self.memo[index]

        minExtra = self.solve(index + 1, length, s) + 1 # This will find worse case scenario where we dont use any char 
        for end in range(index, length): # for end in range(0, 7):  # s has length 7 || This is example of the syntax
            substring = s[index:end+1]
            if substring in self.wordSet:
                minExtra = min(minExtra, self.solve(end +1, length, s))

        self.memo[index] = minExtra
        return minExtra






