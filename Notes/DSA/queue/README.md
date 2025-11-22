## Queue

Cover what a stack is and how they are used.

### 1. Introduction
Imagine you're at your favorite coffee shop. It's rush hour, and there's a line of people waiting to place their orders. This line is a perfect real-life example of a Queue! In this line, the person who has been waiting the longest gets served first (we call this the "First In, First Out" or FIFO principle). Just like this, a Queue in computer science is a type of data structure where the element that enters first is the one that gets accessed first.

Now, why should you care about Queues? Well, Queues are an essential data structure in computer science. They help us solve complex problems and are widely used in various fields, from handling processes in an operating system to managing data packets in networking. By the end of this chapter, you'll see just how valuable understanding Queues can be!

### 2. The FIFO Principle & Terminology
To get comfortable with Queues, we need to grasp some key terms.

- Enqueue: This is like adding a person to the end of our coffee shop line. In Queues, when we add an element at the end, we call it "Enqueue."

- Dequeue: Remember how the first person in line gets their coffee first? Similarly, removing an element from the front of the Queue is called "Dequeue."

- Front: This is the start of the Queue, where the first element resides.

- Rear: This is the end of the Queue, where the last element is placed.

When you think about it, it's pretty simple, right? These are the basics, and once you've got a firm grip on them, everything else about Queues will fall into place!

### 3. Types of Queues
You might think that all Queues are the same, but that's not quite the case. Just like we have different types of lines (a rush-hour coffee shop line versus a less crowded post office line, for example), we also have different types of Queues. Let's look at a few of these:

- Simple Queue: This is the most basic type. Elements are inserted at the rear and removed from the front, following the FIFO principle.

- Deque (Double Ended Queue): Here, elements can be added or removed from both ends of the Queue. Think of it as having a coffee shop line where orders can be taken from both ends!

- Circular Queue: A circular queue is a linear data structure where the last position is connected back to the first, forming a circle. It efficiently utilizes space by reusing positions when elements are dequeued.

- Priority Queue: A priority queue is a data structure where each element has a priority, and elements with higher priority are dequeued before elements with lower priority. It can be implemented using heaps for efficient access to the highest (or lowest) priority element.

- Affinity Queue: In this type, every element has an affinity & is placed with an element having the same affinity; otherwise placed at the rear.

Each of these types of Queues serves different needs and purposes, and we'll explore them in more detail in the upcoming lessons. Understanding them will expand your problem-solving toolbox and make you a stronger programmer.

### 4. Sliding Window Technique
General Steps to Set Up a Sliding Window for Different Problems:
- Identify the window size (k): Determine how many elements you want to process in each sliding window. This will typically be given in the problem statement.

- Define the window start and end:
The window starts at index i-k+1 (this helps you identify when elements are out of the window).
The window ends at index i (as i increases, the window expands to include the current element).

- Handle boundary conditions:
Use i >= k-1 to make sure that you're working with a valid window (i.e., once the index reaches k-1, you have a full window).
Use i-k+1 to identify when an element has fallen out of the window and should be removed from tracking (e.g., deque or array).

- Track the required information: As the window slides:
If you're looking for the maximum, track it using a deque or similar data structure to efficiently maintain the largest element.
If you're tracking the sum, adjust the sum by adding the new element and removing the old one as the window slides.
Store or return results: After processing a full window, store or return the result (e.g., maximum, sum, etc.).