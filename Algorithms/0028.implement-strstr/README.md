# [28. Implement strStr()](https://leetcode-cn.com/problems/implement-strstr/)
Implement [strStr()](http://www.cplusplus.com/reference/cstring/strstr/).

Return the index of the first occurrence of needle in haystack, or **-1** if needle is not part of haystack.

**Example 1:**


<pre><strong>Input:</strong> haystack = &#34;hello&#34;, needle = &#34;ll&#34;
<strong>Output:</strong> 2
</pre>

**Example 2:**


<pre><strong>Input:</strong> haystack = &#34;aaaaa&#34;, needle = &#34;bba&#34;
<strong>Output:</strong> -1
</pre>

**Clarification:**

What should we return when <code>needle</code> is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when <code>needle</code> is an empty string. This is consistent to C&#39;s [strstr()](http://www.cplusplus.com/reference/cstring/strstr/) and Java&#39;s [indexOf()](https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)).
## 解题思路## 可能的變化