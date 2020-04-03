# [61. Rotate List](https://leetcode-cn.com/problems/rotate-list/)
Given a linked list, rotate the list to the right by _k_ places, where _k_ is non-negative.

**Example 1:**


<pre><strong>Input:</strong> 1-&gt;2-&gt;3-&gt;4-&gt;5-&gt;NULL, k = 2
<strong>Output:</strong> 4-&gt;5-&gt;1-&gt;2-&gt;3-&gt;NULL
<strong>Explanation:</strong>
rotate 1 steps to the right: 5-&gt;1-&gt;2-&gt;3-&gt;4-&gt;NULL
rotate 2 steps to the right: 4-&gt;5-&gt;1-&gt;2-&gt;3-&gt;NULL
</pre>

**Example 2:**


<pre><strong>Input:</strong> 0-&gt;1-&gt;2-&gt;NULL, k = 4
<strong>Output:</strong> <code>2-&gt;0-&gt;1-&gt;NULL</code>
<strong>Explanation:</strong>
rotate 1 steps to the right: 2-&gt;0-&gt;1-&gt;NULL
rotate 2 steps to the right: 1-&gt;2-&gt;0-&gt;NULL
rotate 3 steps to the right: <code>0-&gt;1-&gt;2-&gt;NULL</code>
rotate 4 steps to the right: <code>2-&gt;0-&gt;1-&gt;NULL</code></pre>

## 解题思路## 可能的變化