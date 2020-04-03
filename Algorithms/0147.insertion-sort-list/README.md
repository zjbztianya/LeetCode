# [147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
Sort a linked list using insertion sort.

![](https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif)<small>A graphical example of insertion sort. The partial sorted list (black) initially contains only the first element in the list.With each iteration one element (red) is removed from the input data and inserted in-place into the sorted list</small>

**Algorithm of Insertion Sort:**

- Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
- At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
- It repeats until no input elements remain.
**Example 1:**


<pre><strong>Input:</strong> 4-&gt;2-&gt;1-&gt;3
<strong>Output:</strong> 1-&gt;2-&gt;3-&gt;4
</pre>

**Example 2:**


<pre><strong>Input:</strong> -1-&gt;5-&gt;3-&gt;4-&gt;0
<strong>Output:</strong> -1-&gt;0-&gt;3-&gt;4-&gt;5
</pre>

## 解题思路## 可能的變化