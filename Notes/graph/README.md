# 📊 Graphs: A Fundamental Data Structure

A **graph** is a data structure used to represent relationships or connections between objects. A graph consists of two main parts:

- **Vertices (Nodes)**: Represent the individual objects.
- **Edges (Links)**: Represent the connections between those objects.

A graph is often denoted as `G(V, E)` where:

- `V` is the set of vertices.
- `E` is the set of edges connecting pairs of vertices.

---

## 🧠 Example

Consider a graph with five nodes A, B, C, D, and E, and the following connections:

![alt text](image.png)

This graph forms a network of connections, where you can traverse from one node to another based on the defined edges.

As the graph grows larger, it can become highly interconnected, forming complex relationships and paths.

---

## 🚀 Why Are Graphs Important?

Graphs are a fundamental and flexible data structure used in many real-world scenarios—especially those involving:

- 📍 Routes and maps
- 👥 User connections in social media
- 🌐 Communication networks
- 📅 Dependencies in task scheduling
- 🔗 Web page links and recommendations

---

## 📚 Common Graph Terminologies

Understanding baseline graph terminology is essential for mastering graph concepts.

### 1. **Node (Vertex)**

A **node** (also known as a **vertex**) is a fundamental building block of a graph.

> Nodes represent entities, and edges represent relationships between them.

Example: In a graph with nodes A, B, C, D, and E — each letter is a node.

![alt text](image-1.png)

### 2. **Adjacent Nodes**

Two nodes are **adjacent** if an edge directly connects them.

> The set of nodes adjacent to a given node is called its **neighborhood**.

Example: In a graph with an edge between A and B, A and B are adjacent.

### 3. **Digraph (Directed Graph)**

A **digraph** is a graph in which edges have a direction (represented by arrows).

- An edge from vertex A → B means you can move from A to B, but not necessarily from B to A.

![alt text](image-2.png)

### 4. **Loop**

A **loop** is an edge that connects a vertex to itself.

- In **undirected graphs**, it’s just an edge from a node to itself.

![alt text](image-4.png)

- In **directed graphs**, it's a one-way edge that starts and ends at the same node.

![alt text](image-5.png)

### 5. **Degree of a Node**

The **degree** of a node is the number of edges connected to it.

- In **undirected graphs**:
  - Degree is the count of all connecting edges.
  - A self-loop counts as **2**.

Example:
- Node A with two neighbors → degree 2
- Node C with:
  - 1 edge to B
  - 1 edge to E
  - 1 loop → counts as 2
  - **Total Degree** = 1 + 1 + 2 = 4

  ![alt text](image-6.png)

- In **directed graphs**:
  - **In-degree**: number of incoming edges
  - **Out-degree**: number of outgoing edges
  - In the below graph, the degree of node A is two because two edges are outgoing, and the degree of node C is 1.

  ![alt text](image-7.png)

### 6. **Path**

A **path** is a sequence of vertices where each adjacent pair is connected by an edge.

Example: A → B → C

### 7. **Cycle**

A **cycle** is a path where the first and last nodes are the same.

Example:
- A → B → D → C → A
- This forms a **closed loop** (cycle)

![alt text](image-8.png)

---

> ✅ Graphs are crucial in algorithms, networking, scheduling, and much more. Mastering graph basics is a gateway to solving advanced problems like shortest paths, connectivity, and traversal.

# 📘 Types of Graphs

Graphs can vary based on how their edges behave, whether they carry weights, and how their nodes are connected. Understanding these types helps in choosing the right data structure or algorithm for your problem.

---

## 1. 🔄 Undirected Graph

- Edges have **no direction**.
- If there's an edge between `A` and `B`, you can go **both ways**: A → B and B → A.
- Represents **two-way relationships**.

📌 **Example**: A friendship network — if Alice is friends with Bob, then Bob is also friends with Alice.

![alt text](image-9.png)

---

## 2. 🔁 Directed Graph (Digraph)

- Edges have a **specific direction**.
- A → B does **not** imply B → A.
- Useful for representing **one-way relationships**.

📌 **Example**: Web page links — one site may link to another, but not necessarily the other way around.

![alt text](image-10.png)

---

## 3. ⚖️ Weighted Graph

- Each edge has a **value (weight)** representing **cost, time, distance**, etc.
- Useful when connections are **not equal**.

📌 **Example**: A road map — edges represent travel time or distance.

![alt text](image-11.png)

---

## 4. ➖ Unweighted Graph

- All edges are treated **equally**.
- No weights — just pure connectivity.

📌 **Example**: Family trees or org charts where relationships either **exist or don’t**.

![alt text](image-12.png)

---

## 5. 🔁 Cyclic Graph

- Contains **at least one cycle** (a path that starts and ends at the same node).
- Cycles must involve **three or more nodes** (not just a self-loop).

📌 **Example**: A group of people connected in a way where you can return to the start by following connections.

![alt text](image-13.png)

---

## 6. 🔄 Acyclic Graph

- **No cycles** allowed — once you leave a node, you can’t come back.

### Two Variants:
- **Undirected Acyclic Graph**: If it's connected, it's called a **Tree**; multiple trees form a **Forest**.
- **Directed Acyclic Graph (DAG)**: Has directed edges and no cycles. Can be **topologically sorted**.

📌 **Example**: Task scheduling — some tasks must happen before others.

![alt text](image-14.png)

---

## 7. 🔗 Connected Graph

- Every node is **reachable** from every other node.
- Applies to **undirected graphs**.

📌 **Example**: A fully connected computer network.

![alt text](image-15.png)

---

## 8. ❌ Disconnected Graph

- **Not all nodes are connected**.
- Made up of **independent components**.

📌 **Example**: Two social groups with no connections between them.

![alt text](image-16.png)

---

## 9. 🔄 Strongly Connected Graph

- Applies to **directed graphs**.
- Every node is reachable from **every other node** — in **both directions**.

📌 **Example**: A one-way street network where you can still reach all streets and return.

![alt text](image-17.png)
---

## 📊 Summary Comparison Table

| Graph Type           | Directional? | Weighted? | Cycles Allowed | Example                          |
|----------------------|--------------|-----------|----------------|----------------------------------|
| **Undirected**       | ❌           | Optional  | Yes/No         | Friendship network               |
| **Directed**         | ✅           | Optional | Yes/No          | Web links                        |
| **Weighted**         | ✅ / ❌     | ✅       | Yes/No          | Road map                         |
| **Unweighted**       | ✅ / ❌     | ❌         | Yes/No        | Family tree                      |
| **Cyclic**           | ✅ / ❌       | Optional  | ✅           | Social group forming a circle    |
| **Acyclic (Tree/DAG)**| ✅ / ❌       | Optional  | ❌          | Task scheduling, directory tree  |
| **Connected**        | ❌           | Optional  | Yes/No         | Fully linked network             |
| **Disconnected**     | ✅ / ❌       | Optional  | Yes/No       | Disconnected subgraphs           |
| **Strongly Connected**| ✅           | Optional  | Yes/No        | One-way road system              |

---

> ✅ Understanding the **types of graphs** helps in choosing the right algorithms, especially for problems involving **search, traversal, shortest paths, cycles, and connectivity**.

