"""
Given a string s, remove all duplicate letters from the input string while maintaining the original order of the letters.

Additionally, the returned string should be the smallest in lexicographical order among all possible results.

A string is in the smallest lexicographical order if it appears first in a dictionary. For example, "abc" is smaller 
than "acb" because "abc" comes first alphabetically.

Examples:

Input: "babac"
Expected Output: "abc"
Justification:
After removing 1 b and 1 a from the input string, we can get bac, and abc strings. The final answer is 'abc', which is the smallest lexicographical string without duplicate letters.
"""
class Solution:
    def removeDupes(self, s):
        if s is None:
            return None
        # "babac"
        # b: 0 | a: 1 | b: 2 | a: 3 | c: 4
        check = {}
        new = ""
        for i, char in enumerate(s):
            if char not in check:
                check[char] = i
                new += char
            elif char in check:
                if  char > s[check[char] + 1]:
                    
                    check[char] = i
                    new += char

                else:
                    continue

    def removeDupe(self, s):
        if s is None:
            return None
        
        count = {}
        present = set()

        """
        present is basically a "membership log" for result.

        Why it exists

        Without present → every time you see a character, you might accidentally add it again and get duplicates in result.

        With present → you instantly know if that character is already in the final string we’re building.
        """


        result = []

        for char in s:
            count[char] = count.get(char, 0) + 1

        for char in s:
            if char not in present:
                while result and char < result[-1] and count[result[-1]] > 0: #If there items in result still and char is < end of result and end of result has more letters coming up, remove the end of letter
                    present.remove(result.pop()) # Removes from both present and result 
                present.add(char) # We know getting to this 
                result.append(char)
                count[char] -= 1
        return "".join(result)

"""
Walkthrough of removeDupe for babac

Step 1 — First 'b'

'b' not in present → check while loop:
result is empty → skip popping.

Add 'b' →
result = ['b']
present = {'b'}

Decrement count['b'] → b: 1

Step 2 — 'a'

'a' not in present

While loop:

result not empty

'a' < 'b' ✅

count['b'] = 1 (meaning another 'b' will come later) ✅
→ pop 'b' from result and present

Now add 'a' →
result = ['a']
present = {'a'}

Decrement count['a'] → a: 1

Step 3 — 'b' (second time)

'b' not in present (because we popped it earlier in Step 2) ✅

While loop:

result not empty

'b' < 'a' ❌ → no pop

Add 'b' →
result = ['a', 'b']
present = {'a', 'b'}

Decrement count['b'] → b: 0

Step 4 — 'a' (second time)

'a' is already in present → skip entirely.

Decrement count['a'] → a: 0

Step 5 — 'c'

'c' not in present

While loop:

'c' < 'b' ❌ → skip

Add 'c' →
result = ['a', 'b', 'c']
present = {'a', 'b', 'c'}

Decrement count['c'] → c: 0

✅ Final output: "abc"
"""




s = Solution()

word = "babac"
print(s.removeDupe(word))


        
