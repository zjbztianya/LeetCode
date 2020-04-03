# [8. String to Integer (atoi)](https://leetcode-cn.com/problems/string-to-integer-atoi/)
Implement <code>atoi</code> which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

**Note:**


- Only the space character <code>&#39; &#39;</code> is considered as whitespace character.
- Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

**Example 1:**


<pre><strong>Input:</strong> &#34;42&#34;
<strong>Output:</strong> 42
</pre>

**Example 2:**


<pre><strong>Input:</strong> &#34;   -42&#34;
<strong>Output:</strong> -42
<strong>Explanation:</strong> The first non-whitespace character is &#39;-&#39;, which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
</pre>

**Example 3:**


<pre><strong>Input:</strong> &#34;4193 with words&#34;
<strong>Output:</strong> 4193
<strong>Explanation:</strong> Conversion stops at digit &#39;3&#39; as the next character is not a numerical digit.
</pre>

**Example 4:**


<pre><strong>Input:</strong> &#34;words and 987&#34;
<strong>Output:</strong> 0
<strong>Explanation:</strong> The first non-whitespace character is &#39;w&#39;, which is not a numerical 
             digit or a +/- sign. Therefore no valid conversion could be performed.</pre>

**Example 5:**


<pre><strong>Input:</strong> &#34;-91283472332&#34;
<strong>Output:</strong> -2147483648
<strong>Explanation:</strong> The number &#34;-91283472332&#34; is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−2<sup>31</sup>) is returned.</pre>

## 解题思路## 可能的變化