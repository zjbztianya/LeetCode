# [20. Valid Parentheses](https://leetcode-cn.com/problems/valid-parentheses/)
Given a string containing just the characters <code>&#39;(&#39;</code>, <code>&#39;)&#39;</code>, <code>&#39;{&#39;</code>, <code>&#39;}&#39;</code>, <code>&#39;[&#39;</code> and <code>&#39;]&#39;</code>, determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

**Example 1:**


<pre><strong>Input:</strong> &#34;()&#34;
<strong>Output:</strong> true
</pre>

**Example 2:**


<pre><strong>Input:</strong> &#34;()[]{}&#34;
<strong>Output:</strong> true
</pre>

**Example 3:**


<pre><strong>Input:</strong> &#34;(]&#34;
<strong>Output:</strong> false
</pre>

**Example 4:**


<pre><strong>Input:</strong> &#34;([)]&#34;
<strong>Output:</strong> false
</pre>

**Example 5:**


<pre><strong>Input:</strong> &#34;{[]}&#34;
<strong>Output:</strong> true
</pre>

## 解题思路## 可能的變化