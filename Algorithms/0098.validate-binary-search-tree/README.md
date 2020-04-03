# [98. Validate Binary Search Tree](https://leetcode-cn.com/problems/validate-binary-search-tree/)
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:


- The left subtree of a node contains only nodes with keys **less than** the node&#39;s key.
- The right subtree of a node contains only nodes with keys **greater than** the node&#39;s key.
- Both the left and right subtrees must also be binary search trees.



**Example 1:**


<pre>    2
   / \
  1   3

<strong>Input:</strong> [2,1,3]
<strong>Output:</strong> true
</pre>

**Example 2:**


<pre>    5
   / \
  1   4
     / \
    3   6

<strong>Input:</strong> [5,1,4,null,null,3,6]
<strong>Output:</strong> false
<strong>Explanation:</strong> The root node&#39;s value is 5 but its right child&#39;s value is 4.
</pre>

## 解题思路## 可能的變化