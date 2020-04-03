# [12. Integer to Roman](https://leetcode-cn.com/problems/integer-to-roman/)
Roman numerals are represented by seven different symbols: <code>I</code>, <code>V</code>, <code>X</code>, <code>L</code>, <code>C</code>, <code>D</code> and <code>M</code>.


<pre><strong>Symbol</strong>       <strong>Value</strong>
I             1
V             5
X             10
L             50
C             100
D             500
M             1000</pre>

For example, two is written as <code>II</code> in Roman numeral, just two one&#39;s added together. Twelve is written as, <code>XII</code>, which is simply <code>X</code> + <code>II</code>. The number twenty seven is written as <code>XXVII</code>, which is <code>XX</code> + <code>V</code> + <code>II</code>.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not <code>IIII</code>. Instead, the number four is written as <code>IV</code>. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as <code>IX</code>. There are six instances where subtraction is used:


- <code>I</code> can be placed before <code>V</code> (5) and <code>X</code> (10) to make 4 and 9. 
- <code>X</code> can be placed before <code>L</code> (50) and <code>C</code> (100) to make 40 and 90. 
- <code>C</code> can be placed before <code>D</code> (500) and <code>M</code> (1000) to make 400 and 900.

Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

**Example 1:**


<pre><strong>Input:</strong> 3
<strong>Output:</strong> &#34;III&#34;</pre>

**Example 2:**


<pre><strong>Input:</strong> 4
<strong>Output:</strong> &#34;IV&#34;</pre>

**Example 3:**


<pre><strong>Input:</strong> 9
<strong>Output:</strong> &#34;IX&#34;</pre>

**Example 4:**


<pre><strong>Input:</strong> 58
<strong>Output:</strong> &#34;LVIII&#34;
<strong>Explanation:</strong> L = 50, V = 5, III = 3.
</pre>

**Example 5:**


<pre><strong>Input:</strong> 1994
<strong>Output:</strong> &#34;MCMXCIV&#34;
<strong>Explanation:</strong> M = 1000, CM = 900, XC = 90 and IV = 4.</pre>

## 解题思路## 可能的變化