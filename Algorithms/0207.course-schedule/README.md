# [207. Course Schedule](https://leetcode-cn.com/problems/course-schedule/)
There are a total of <code>numCourses</code> courses you have to take, labeled from <code>0</code> to <code>numCourses-1</code>.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: <code>[0,1]</code>

Given the total number of courses and a list of prerequisite **pairs**, is it possible for you to finish all courses?



**Example 1:**


<pre><strong>Input:</strong> numCourses = 2, prerequisites = [[1,0]]
<strong>Output:</strong> true
<strong>Explanation:</strong> There are a total of 2 courses to take.
             To take course 1 you should have finished course 0. So it is possible.
</pre>

**Example 2:**


<pre><strong>Input:</strong> numCourses = 2, prerequisites = [[1,0],[0,1]]
<strong>Output:</strong> false
<strong>Explanation:</strong> There are a total of 2 courses to take.
             To take course 1 you should have finished course 0, and to take course 0 you should
             also have finished course 1. So it is impossible.
</pre>



**Constraints:**


- The input prerequisites is a graph represented by **a list of edges**, not adjacency matrices. Read more about [how a graph is represented](https://www.khanacademy.org/computing/computer-science/algorithms/graph-representation/a/representing-graphs).
- You may assume that there are no duplicate edges in the input prerequisites.
- <code>1 &lt;= numCourses &lt;= 10^5</code>
## 解题思路## 可能的變化