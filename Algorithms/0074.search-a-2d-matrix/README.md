# [74. Search a 2D Matrix](https://leetcode-cn.com/problems/search-a-2d-matrix/)
Write an efficient algorithm that searches for a value in an _m_ x _n_ matrix. This matrix has the following properties:


- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.

**Example 1:**


<pre><strong>Input:</strong>
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
<strong>Output:</strong> true
</pre>

**Example 2:**


<pre><strong>Input:</strong>
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
<strong>Output:</strong> false</pre>

## 解题思路## 可能的變化