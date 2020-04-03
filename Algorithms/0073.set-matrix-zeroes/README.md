# [73. Set Matrix Zeroes](https://leetcode-cn.com/problems/set-matrix-zeroes/)
Given a _m_ x _n_ matrix, if an element is 0, set its entire row and column to 0. Do it [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm).

**Example 1:**


<pre><strong>Input:</strong> 
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
<strong>Output:</strong> 
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
</pre>

**Example 2:**


<pre><strong>Input:</strong> 
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
<strong>Output:</strong> 
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
</pre>

**Follow up:**


- A straight forward solution using O(_m__n_) space is probably a bad idea.
- A simple improvement uses O(_m_ + _n_) space, but still not the best solution.
- Could you devise a constant space solution?
## 解题思路## 可能的變化