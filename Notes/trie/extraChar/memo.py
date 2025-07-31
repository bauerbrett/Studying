class Solution:
    def minChar(self, s: str, dictionary: list[str]) -> int:
        words = set(dictionary)
        length = len(s)
        memo = {}

        def dfs(i):
            if i == length:
                return 0
            
            if i in memo:
                return memo[i]
            
            minchar = 1 + dfs(i + 1)
            for end in range(i, length):
                substr = s[i:end+1]
                if substr in words:
                    minchar = min(minchar, dfs(end+1))

            memo[i] = minchar
            return minchar
        return dfs(0)
    

d = ["race", "car"]
s = "amazingracecar"

print(Solution().minChar(s, d))

                

            

