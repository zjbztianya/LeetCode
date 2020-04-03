# [210. Course Schedule II](https://leetcode-cn.com/problems/course-schedule-ii/)
There are a total of _n_ courses you have to take, labeled from <code>0</code> to <code>n-1</code>.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: <code>[0,1]</code>

Given the total number of courses and a list of prerequisite **pairs**, return the ordering of courses you should take to finish all courses.

There may be multiple correct orders, you just need to return one of them. If it is impossible to finish all courses, return an empty array.

**Example 1:**


<pre><strong>Input:</strong> 2, [[1,0]] 
<strong>Output: </strong><code>[0,1]</code>
<strong>Explanation:</strong> There are a total of 2 courses to take. To take course 1 you should have finished   
             course 0. So the correct course order is <code>[0,1] .</code></pre>

**Example 2:**


<pre><strong>Input:</strong> 4, [[1,0],[2,0],[3,1],[3,2]]
<strong>Output: </strong><code>[0,1,2,3] or [0,2,1,3]</code>
<strong>Explanation:</strong> There are a total of 4 courses to take. To take course 3 you should have finished both     
             courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0. 
             So one correct course order is <code>[0,1,2,3]</code>. Another correct ordering is <code>[0,2,1,3] .</code></pre>

**Note:**

- The input prerequisites is a graph represented by **a list of edges**, not adjacency matrices. Read more about [how a graph is represented](https://www.khanacademy.org/computing/computer-science/algorithms/graph-representation/a/representing-graphs).
- You may assume that there are no duplicate edges in the input prerequisites.## 解题思路## 可能的變化