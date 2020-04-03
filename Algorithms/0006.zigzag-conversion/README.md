# [6. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
The string <code>&#34;PAYPALISHIRING&#34;</code> is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)


<pre>P   A   H   N
A P L S I I G
Y   I   R
</pre>

And then read line by line: <code>&#34;PAHNAPLSIIGYIR&#34;</code>

Write the code that will take a string and make this conversion given a number of rows:


<pre>string convert(string s, int numRows);</pre>

**Example 1:**


<pre><strong>Input:</strong> s = &#34;PAYPALISHIRING&#34;, numRows = 3
<strong>Output:</strong> &#34;PAHNAPLSIIGYIR&#34;
</pre>

**Example 2:**


<pre><strong>Input:</strong> s = &#34;PAYPALISHIRING&#34;, numRows = 4
<strong>Output:</strong> &#34;PINALSIGYAHRPI&#34;
<strong>Explanation:</strong>

P     I    N
A   L S  I G
Y A   H R
P     I</pre>

## 解题思路## 可能的變化