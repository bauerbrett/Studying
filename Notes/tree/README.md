## Tree

Cover what a tree is and how they are used.

### 1. Introduction - What is a Tree?
In computer science, a tree is a hierarchical data structure that simulates a tree-like structure with a set of connected nodes. Unlike linear data structures such as arrays or linked lists, trees allow for more complex relationships and efficient organization of data. They are particularly useful for representing data that naturally forms a hierarchy, such as file systems, organizational structures, and more.

Visual Representation:
        A
       / \
      B   C
     / \   \
    D   E   F
In the above example:

A is the root node.
B and C are child nodes of A.
D and E are children of B, while F is a child of C.


### 2. Defining a Tree
A tree is made up of nodes that are arranged in a hierarchical structure. At the top of the hierarchy is the root node, which acts as the starting point. All other nodes are connected through edges. The nodes are grouped into levels, and the maximum level of any node in the tree is referred to as the depth.

Here are the key terms related to trees:

- Root Node: The node at the top of the tree. It is the node from where the whole tree originates. For any tree traversal operation, this node serves as the starting point.

- Nodes: All the elements in the tree, including the root, are nodes, and each node has a unique value and may have child nodes connected to it.

- Parent Node: A node in a tree that has one or more child nodes connected to it.

- Child Node: Nodes that are directly connected to a parent node.

- Sibling Nodes: Nodes that share the same parent node.

- Ancestor Nodes: All the nodes in the path from a specific node to the root node.

- Descendant Nodes: All the nodes reachable from a specific node down to the leaves.

- Leaf Node: Nodes in the tree that do not have any children.

- Internal Node: A node with at least one child.

- Edge: The connection between one node and another.

- Path: A sequence of nodes connected by edges.

- Height of a Node: The number of edges on the longest path from the node to a leaf. - Height goes down to leaves- full path - Depth goes up to root from a specific node

- Depth of a Node: The number of edges from the node to the tree's root node.

- Height of the Tree: The height of the root node.

- Degree of a Node: The number of children a node has.

- Subtree: A smaller tree within the main tree, consisting of a node and its descendants.

- The following figure illustrates the different components of a tree.

### 3. Tree Height and Depth
The height of a tree is the number of edges on the longest path from the root node to the leaf node. It represents the depth of the tree from the root. In contrast, the depth of a specific node in the tree is the number of edges from the root node to that particular node.

### 4. Levels in Trees
Levels in a tree are defined based on the distance from the root node. The root node is at level 0; its children are at level 1, and so on. The level of a node indicates its generation within the tree.

### 5. Types of Trees
- Binary Trees:
Binary trees are a type of tree where each node can have at most two children, commonly referred to as the left child and the right child. The structure of a binary tree makes it a fundamental and widely used data structure in computer science.
![alt text](images/image-1.png)

- Full Tree:
In a full tree, every node has either zero children (leaf node) or two children. There are no nodes with only one child. Full trees are also known as proper binary trees.
![alt text](images/image.png)

- Complete Tree:
A complete tree is a binary tree in which all levels are filled, except possibly the last level. The last level must strictly be filled from left to right. Data structures like Heap uses complete binary trees for efficient operations.

![alt text](images/image-5.png)

- Balanced Tree:
Balanced trees are binary trees where the difference in height between the left and right subtrees of any node in the tree is not more than 1. This ensures the tree remains reasonably balanced, preventing skewed structures and maintaining optimal search and insertion times.
![alt text](images/image-3.png)

- Multi-way Tree:
Unlike binary trees, multi-way trees allow nodes to have more than two children. Each node can have multiple branches, making multi-way trees more flexible in representing hierarchical data.
![alt text](images/image-4.png)

### 6. Binary Search Tree (BST)
A Binary Search Tree (BST) is a fundamental data structure that organizes elements in a hierarchical manner, allowing for efficient searching, insertion, deletion, and traversal operations. In this section, we will explore the characteristics of a BST, its benefits, applications, ADT operations, time and space complexity, as well as potential issues and solutions.

Definition and Characteristics of BST
A Binary Search Tree is a binary tree in which each node has at most two children, referred to as the left child and the right child. The BST follows a specific property: for any given node, all nodes in its left subtree have values less than or equal to the node's value, and all nodes in its right subtree have values greater than the node's value.

Example:
Consider the following BST:
![alt text](images/image-bst.png)
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




