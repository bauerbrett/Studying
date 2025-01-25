## Binary Search Tree (BST)

Cover what a BST is and how they are used.

### 1. Introduction - What is a Binary Search Tree (BST)
A Binary Search Tree (BST) is a fundamental data structure that organizes elements in a hierarchical manner, allowing for efficient searching, insertion, deletion, and traversal operations. In this section, we will explore the characteristics of a BST, its benefits, applications, ADT operations, time and space complexity, as well as potential issues and solutions.

Definition and Characteristics of BST
A Binary Search Tree is a binary tree in which each node has at most two children, referred to as the left child and the right child. The BST follows a specific property: for any given node, all nodes in its left subtree have values less than or equal to the node's value, and all nodes in its right subtree have values greater than the node's value.

Example:
Consider the following BST:
![alt text](image.png)
In this example, the tree satisfies the BST property since, for each node, all nodes in its left subtree have values less than the node's value, and all nodes in its right subtree have values greater than the node's value.

Benefits and Applications of BST
BSTs provide efficient searching operations due to their hierarchical structure and the BST property. The search can be performed in O(log n) time in a balanced BST.

Efficient Searching: BSTs provide fast searching capabilities with an average time complexity of O(log n) for balanced trees. This makes them suitable for applications that require quick data retrieval based on a key, such as database indexing and dictionary implementations.
Ordered Data Storage: The inherent order of BSTs makes them suitable for scenarios where data needs to be stored in a sorted manner. For example, BSTs can maintain sorted lists or retrieve elements in a particular order.
Binary Search Tree in Databases: BSTs are used in databases to create indexes that allow for faster data retrieval based on specific fields. Using BSTs for indexing, databases can efficiently handle large datasets and complex queries.
Binary Search Tree in File Systems: File systems often use BSTs to organize and manage directory structures. The hierarchical nature of BSTs aligns well with the hierarchical structure of directories, enabling efficient file retrieval.
BST Abstract Data Type (ADT)
The BST ADT covers various operations to manipulate and manage the tree efficiently. The basic operations typically supported by a BST ADT include the following:

- Insertion
- Update
- Deletion
- Search
- In-order Traversal
- Pre-order Traversal
- Post-order Traversal

Let's start by discussing our data model. We will use a Node class to store a single node of BST. Here is what our Node class looks like:
type Node[T any] struct {
    Data  T
    Left  *Node[T]
    Right *Node[T]
}
The above class represents a single node of a binary tree. It has the following members:
- data: A template member that stores the value of the node.
- left: A pointer to the left child node.
- right: A pointer to the right child node.

The class has two constructors:
- Node(): Default constructor that initializes left and right pointers to null (0).
- Node(int v): Constructor that takes a value v and initializes info with it. It also sets left and right pointers to null.
