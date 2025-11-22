class TrieNode:
    def __init__(self):
        self.children = {}
        self.isEnd = False
class Trie:
    def __init__(self):
        self.root = TrieNode()
    
    def addWord(self, word):
        node = self.root
        for c in word:
            if c not in node.children:
                node.children[c] = TrieNode()
            node = node.children[c]
        node.isEnd = True

    def search(self, word, root: TrieNode):
        node = root
        for i, c in enumerate(word):
            if c == '.':
                for child in node.children:
                    # Recursively search the rest of the word starting from each child node
                    next_node = node.children[child] # get next child
                    if self.search(word[i+1:], next_node): #put the new word chopping off the already search word and the new node and search it
                        return True
                return False  # None of the paths matched
            if c not in node.children:
                return False
            node = node.children[c]
        return node.isEnd

if __name__ == "__main__":
    t = Trie()

    t.addWord("apple")
    t.addWord("banana")
    print(t.search("......", t.root))
    print(t.search("apple", t.root))
    print(t.search("ace", t.root))
