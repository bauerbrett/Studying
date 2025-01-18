## Security Engineering Notes

These are notes for studying. They are primarily used for coding, data structures, and algorithims right now because I need to improve the most on those.

## Topics

### 1. Coding Data Stuctures
#### Dictionaries/Maps/HashTables
- Concepts:
    -  A hash table stores data in key-value pairs.
    -  Keys are converted into a unique index using a hash function.
    -  The index maps to a location in an underlying array.
- Key Operations:
    - Insert/Update: Add a key-value pair to the table.
    - Search: Retrieve a value using its key.
    - Delete: Remove a key-value pair.
    - Examples: Item #1 in main.go

Let's explore some examples of these data structures.

Primitive Data Structures
Imagine you're making a shopping list. What do you need? Probably fruits, vegetables, snacks, maybe some cleaning supplies. Each of these items can be considered an individual data element - a building block of your list. Primitive data structures in programming are similar. They are the simplest form of data structures and include:

    Integers: Think of these as whole numbers, like the number of apples you need to buy.
    Float: These are real numbers, like the weight of the bananas you're planning to buy.
    Character: A single letter, number, or symbol, like 'A' or '3'.
    Boolean: This type holds either a true or false value.
    Non-Primitive Data Structures
    In contrast, non-primitive data structures are more complex. They are derived from the primitive data structures and can be further divided into linear and non-linear types.

    Linear Data Structures
    In linear data structures, elements are arranged in a sequential manner. Think of them as a line of people waiting for a bus. Some examples are:

    Arrays: Think of an array like a row of lockers. Each locker holds an item and is identified by a unique number.
    Linked Lists: Imagine a chain where each link knows about the link following it.
    Stacks and Queues: Stacks are like a pile of books, where you can only interact with the top book. Queues, on the other hand, are like a line at a grocery store - the first one in line gets served first.

Non-Linear Data Structures
Non-linear data structures, as the name implies, do not maintain a particular sequence for the arrangement of elements. Examples include:

    Trees: Think of a family tree where each person is a node that has connections to their parents and children.
    Graphs: These are like social networks, where each person (node) can be connected to multiple other people (nodes).
    We've only scratched the surface of data structures here. As you delve deeper into each type, you'll gain a stronger understanding of how they work and where they can be applied.

We will start with the Arrays data structure but before that let's see what is Big-O.

What is Big-O?
    Big-O notation is a way of expressing the time (or space) complexity of an algorithm. It provides a rough estimate of how long an algorithm takes to run (or how much memory it uses), based on the size of the input. For example, an algorithm with a time complexity of  means that the running time increases linearly with the size of the input.

    What is time complexity?
    Time complexity is a measure of how long an algorithm takes to run, based on the size of the input. It is expressed using Big-O notation, which provides a rough estimate of the running time. An algorithm with a lower time complexity will generally be faster than an algorithm with a higher time complexity.

    What is space complexity?
    Space complexity is a measure of how much memory an algorithm requires, based on the size of the input. Like time complexity, it is expressed using Big-O notation. An algorithm with a lower space complexity will generally require less memory than an algorithm with a higher space complexity.

    Examples of time complexity
    Here are some examples of how different time complexities are expressed using Big-O notation:

    - O(1):Constant time. The running time is independent of the size of the input.
    - O(n): Linear time. The running time increases linearly with the size of the input.
    - O(n2): Quadratic time. The running time is proportional to the square of the size of the input.
    - O(logn): Logarithmic time. The running time increases logarithmically with the size of the input.
    - O(2^n): Exponential time. The running time increases exponentially in the size of the input.





#### Array/Slice
- Concepts:
    - An array is a collection of elements stored in contiguous memory.
    - Arrays usually have a fixed size in most programming languages (except Python, where lists are dynamic).
- Key Operations:
    - Access by Index: O(1).
    - Search: O(n) for unsorted arrays, O(log n) for sorted arrays (using binary search).
    - Insert/Delete:
    - End: O(1).
    - Middle: O(n) (requires shifting).
    - Examples: Item #2 in main.go