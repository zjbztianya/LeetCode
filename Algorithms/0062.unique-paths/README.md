# [62. Unique Paths](https://leetcode-cn.com/problems/unique-paths/)
A robot is located at the top-left corner of a _m_ x _n_ grid (marked &#39;Start&#39; in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked &#39;Finish&#39; in the diagram below).

How many possible unique paths are there?

![](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)<small>Above is a 7 x 3 grid. How many possible unique paths are there?</small>



**Example 1:**


<pre><strong>Input:</strong> m = 3, n = 2
<strong>Output:</strong> 3
<strong>Explanation:</strong>
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -&gt; Right -&gt; Down
2. Right -&gt; Down -&gt; Right
3. Down -&gt; Right -&gt; Right
</pre>

**Example 2:**


<pre><strong>Input:</strong> m = 7, n = 3
<strong>Output:</strong> 28
</pre>



**Constraints:**


- <code>1 &lt;= m, n &lt;= 100</code>
- It&#39;s guaranteed that the answer will be less than or equal to <code>2 * 10 ^ 9</code>.
## 解题思路## 可能的變化