class TrieNode:
    def __init__(self):
        self.children = {}
        self.isEnd = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def findPrefix(self, prefix: str):
        node = self.root
        products = []

        for c in prefix:
            if c not in node.children:
                return []
            node = node.children[c]
        
        def dfs(node: TrieNode, path: str):
            if len(products) >= 3:
                return
            if node.isEnd:
                products.append(path)

            """
            # Recursively search for all possible words
            for ch in 'abcdefghijklmnopqrstuvwxyz':
            if ch in node.children:
                self.dfs(node.children[ch], prefix + ch, list)

            This loops for all, the other will skip right to the next char in order, so if only z exist in children it will find it instantly instead of last.
            """
            for c in sorted(node.children):  # Sorted to keep lexicographical order
                dfs(node.children[c], path + c)
 
        dfs(node, prefix)
        return products

    def insert(self, word: str):
        node = self.root
        for c in word:
            if c not in node.children:
                node.children[c] = TrieNode()
            node = node.children[c]
        node.isEnd = True

class Solution:
    def suggestionSystem(self, products: list[str], searchWord: str):
        t = Trie()
        for product in products:
            t.insert(product)
        
        return t.findPrefix(searchWord)

if __name__ == "__main__":
    s = Solution()
    products1 = ["apple", "apricot", "application"]
    searchWord1 = "app"
    print("Example 1:", s.suggestionSystem(products1, searchWord1))
