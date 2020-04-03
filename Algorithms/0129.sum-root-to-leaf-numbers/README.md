# [129. Sum Root to Leaf Numbers](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/)
Given a binary tree containing digits from <code>0-9</code> only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path <code>1-&gt;2-&gt;3</code> which represents the number <code>123</code>.

Find the total sum of all root-to-leaf numbers.

**Note:** A leaf is a node with no children.

**Example:**


<pre><strong>Input:</strong> [1,2,3]
    1
   / \
  2   3
<strong>Output:</strong> 25
<strong>Explanation:</strong>
The root-to-leaf path <code>1-&gt;2</code> represents the number <code>12</code>.
The root-to-leaf path <code>1-&gt;3</code> represents the number <code>13</code>.
Therefore, sum = 12 + 13 = <code>25</code>.</pre>

**Example 2:**


<pre><strong>Input:</strong> [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
<strong>Output:</strong> 1026
<strong>Explanation:</strong>
The root-to-leaf path <code>4-&gt;9-&gt;5</code> represents the number 495.
The root-to-leaf path <code>4-&gt;9-&gt;1</code> represents the number 491.
The root-to-leaf path <code>4-&gt;0</code> represents the number 40.
Therefore, sum = 495 + 491 + 40 = <code>1026</code>.</pre>

## 解题思路## 可能的變化