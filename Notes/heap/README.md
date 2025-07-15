# 📚 Introduction to Heap

Heaps are a tree-based data structure designed for efficient priority management. They are widely used in priority queues, Dijkstra’s shortest path algorithm, heap sort, and scheduling systems. A heap is a special type of binary tree that maintains a specific order between parent and child nodes.

A heap is always a **complete binary tree**, meaning all levels are fully filled except possibly the last level, which is filled from left to right.

## 🏗️ Two Main Types of Heaps

- **Max Heap**: Every parent node has a value greater than or equal to its children. The largest element is at the root.
- **Min Heap**: Every parent node has a value smaller than or equal to its children. The smallest element is at the root.

> Unlike a Binary Search Tree (BST), heaps do not enforce an ordering between siblings—only between parents and children.

To optimize performance, heaps are commonly stored as arrays.

---

## 📦 Storing a Heap in an Array

While heaps are visually represented as binary trees, they are usually stored in an array to optimize space and operations. This array representation makes insertion, deletion, and access operations more efficient.

### 🔢 Array Representation Rules

For any node at index `i`:

- Left child: `2i + 1`
- Right child: `2i + 2`
- Parent node: `(i - 1) / 2` (integer division)

---

## 🧮 Example: Max Heap as an Array

Level-order traversal gives the array representation:

- `16` (root) → index `0`
- `14`, `10` → indices `1`, `2`
- `8`, `7`, `9`, `3` → indices `3–6`
- `2`, `4` → indices `7–8`

---

## 🧮 Example: Min Heap as an Array

Using the same rules:

- `21` → index `1`
  - Left child → index `3` → `24`
  - Right child → index `4` → `31`
- `26` (index `8`) → Parent at `(8-1)/2 = 3` → `24`

---

# 🔍 Basic Operations on Heap

---

## 1️⃣ Find Maximum/Minimum in a Heap

- **Max Heap** → max at root (index 0)
- **Min Heap** → min at root (index 0)

> 🕒 Time Complexity: **O(1)**

---

## 2️⃣ Inserting Elements into a Heap (Heapify Up)

Insertion occurs at the end to maintain the complete binary tree structure. Then:

### 🔁 Heapify Up Process

- Insert at last index
- Compare with parent:
  - If heap property violated → swap
  - Repeat until heap property restored

> 🕒 Time Complexity: **O(log n)**  
> 💾 Space Complexity: **O(1)**

### ✅ Example: Inserting into Min Heap

```go
package main

import "fmt"

type Solution struct{}

// Heapify Up (Percolate Up) for Min Heap
func heapifyUp(heap []int, i int) {
    for i > 0 {
        parent := (i - 1) / 2
        if heap[i] < heap[parent] {
            heap[i], heap[parent] = heap[parent], heap[i]
            i = parent
        } else {
            break
        }
    }
}
```
### 3️⃣ Deleting Elements from a Heap (Heapify Down)

Removing the root from a heap requires replacing it with the last element and restoring the heap property via **Heapify Down**.

#### 🔁 Heapify Down Steps
1. Start at root (`i = 0`)
2. Compare with left/right children
3. Swap with the smaller (Min Heap) or larger (Max Heap) child
4. Repeat until the heap property is restored

> 🕒 **Time Complexity**: `O(log n)`  
> 💾 **Space Complexity**: `O(1)`

#### ✅ Example: Min Heap Deletion in Go

```go
package main

import "fmt"

type Solution struct{}

// Heapify Down for Min Heap
func (s Solution) heapifyDown(heap *[]int, i int) {
    size := len(*heap)
    for {
        smallest := i
        left := 2*i + 1
        right := 2*i + 2

        if left < size && (*heap)[left] < (*heap)[smallest] {
            smallest = left
        }
        if right < size && (*heap)[right] < (*heap)[smallest] {
            smallest = right
        }

        if smallest != i {
            (*heap)[i], (*heap)[smallest] = (*heap)[smallest], (*heap)[i]
            i = smallest
        } else {
            break
        }
    }
}
```
### 4️⃣ Building a Heap from an Unordered Array

When given an unordered array, we can efficiently transform it into a valid heap using the **Heapify Down** approach. Instead of inserting elements one by one (which takes `O(n log n)` time), we use the **bottom-up heap construction** method, which runs in `O(n)` time.

---

#### 📋 How It Works

- **Start from the last non-leaf node** and call `heapifyDown`.
- **Move upward level by level**, ensuring each subtree satisfies the heap property.
- Continue until the **root is processed**, at which point the entire array becomes a valid heap.

This method is efficient because:
- Heapify is only called on non-leaf nodes (≈ n/2).
- Most nodes are near the bottom and only move 1–2 levels.

---

### 🧠 Algorithm for Building a Heap

#### 🔍 Identify the Last Non-Leaf Node
In a 0-based array:
- Leaf nodes begin at index `n/2`.
- So the last non-leaf node is at index `n/2 - 1`.

#### 🔁 Heapify Down Each Non-Leaf Node in Reverse Order
1. Start from `n/2 - 1` to index `0`.
2. Call `heapifyDown()` on each node to restore the heap property.
3. Once the root is heapified, the array is a valid heap.

---

### ✅ Go Code Example

```go
package main

import "fmt"

type Solution struct{}

// Heapify Down for Min Heap
func (Solution) heapifyDown(heap []int, i int) {
	size := len(heap)
	for {
		smallest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < size && heap[left] < heap[smallest] {
			smallest = left
		}
		if right < size && heap[right] < heap[smallest] {
			smallest = right
		}

		if smallest != i {
			heap[i], heap[smallest] = heap[smallest], heap[i]
			i = smallest
		} else {
			break
		}
	}
}
```