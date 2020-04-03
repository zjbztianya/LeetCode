# [10. Regular Expression Matching](https://leetcode-cn.com/problems/regular-expression-matching/)
Given an input string (<code>s</code>) and a pattern (<code>p</code>), implement regular expression matching with support for <code>&#39;.&#39;</code> and <code>&#39;*&#39;</code>.


<pre>&#39;.&#39; Matches any single character.
&#39;*&#39; Matches zero or more of the preceding element.
</pre>

The matching should cover the **entire** input string (not partial).

**Note:**


- <code>s</code> could be empty and contains only lowercase letters <code>a-z</code>.
- <code>p</code> could be empty and contains only lowercase letters <code>a-z</code>, and characters like <code>.</code> or <code>*</code>.

**Example 1:**


<pre><strong>Input:</strong>
s = &#34;aa&#34;
p = &#34;a&#34;
<strong>Output:</strong> false
<strong>Explanation:</strong> &#34;a&#34; does not match the entire string &#34;aa&#34;.
</pre>

**Example 2:**


<pre><strong>Input:</strong>
s = &#34;aa&#34;
p = &#34;a*&#34;
<strong>Output:</strong> true
<strong>Explanation:</strong> &#39;*&#39; means zero or more of the preceding element, &#39;a&#39;. Therefore, by repeating &#39;a&#39; once, it becomes &#34;aa&#34;.
</pre>

**Example 3:**


<pre><strong>Input:</strong>
s = &#34;ab&#34;
p = &#34;.*&#34;
<strong>Output:</strong> true
<strong>Explanation:</strong> &#34;.*&#34; means &#34;zero or more (*) of any character (.)&#34;.
</pre>

**Example 4:**


<pre><strong>Input:</strong>
s = &#34;aab&#34;
p = &#34;c*a*b&#34;
<strong>Output:</strong> true
<strong>Explanation:</strong> c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches &#34;aab&#34;.
</pre>

**Example 5:**


<pre><strong>Input:</strong>
s = &#34;mississippi&#34;
p = &#34;mis*is*p*.&#34;
<strong>Output:</strong> false
</pre>

## 解题思路## 可能的變化