# [63. Unique Paths II](https://leetcode-cn.com/problems/unique-paths-ii/)
A robot is located at the top-left corner of a _m_ x _n_ grid (marked &#39;Start&#39; in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked &#39;Finish&#39; in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?

![](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)

An obstacle and empty space is marked as <code>1</code> and <code>0</code> respectively in the grid.

**Note:**_m_ and _n_ will be at most 100.

**Example 1:**


<pre><strong>Input:
</strong>[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
<strong>Output:</strong> 2
<strong>Explanation:</strong>
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -&gt; Right -&gt; Down -&gt; Down
2. Down -&gt; Down -&gt; Right -&gt; Right
</pre>

## 解题思路## 可能的變化