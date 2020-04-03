# [142. Linked List Cycle II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)
Given a linked list, return the node where the cycle begins. If there is no cycle, return <code>null</code>.

To represent a cycle in the given linked list, we use an integer <code>pos</code> which represents the position (0-indexed) in the linked list where tail connects to. If <code>pos</code> is <code>-1</code>, then there is no cycle in the linked list.

**Note:** Do not modify the linked list.



**Example 1:**


<pre><strong>Input: </strong>head = [3,2,0,-4], pos = 1
<strong>Output: </strong>tail connects to node index 1
<strong>Explanation:</strong> There is a cycle in the linked list, where tail connects to the second node.
</pre>

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

**Example 2:**


<pre><strong>Input: </strong>head = [1,2], pos = 0
<strong>Output: </strong>tail connects to node index 0
<strong>Explanation:</strong> There is a cycle in the linked list, where tail connects to the first node.
</pre>

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)

**Example 3:**


<pre><strong>Input: </strong>head = [1], pos = -1
<strong>Output: </strong>no cycle
<strong>Explanation:</strong> There is no cycle in the linked list.
</pre>

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)



**Follow-up**:Can you solve it without using extra space?
## 解题思路## 可能的變化