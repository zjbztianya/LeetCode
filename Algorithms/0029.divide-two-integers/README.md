# [29. Divide Two Integers](https://leetcode-cn.com/problems/divide-two-integers/)
Given two integers <code>dividend</code> and <code>divisor</code>, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing <code>dividend</code> by <code>divisor</code>.

The integer division should truncate toward zero, which means losing its fractional part. For example, <code>truncate(8.345) = 8</code> and <code>truncate(-2.7335) = -2</code>.

**Example 1:**


<pre><strong>Input:</strong> dividend = 10, divisor = 3
<strong>Output:</strong> 3
<strong>Explanation:</strong> 10/3 = truncate(3.33333..) = 3.
</pre>

**Example 2:**


<pre><strong>Input:</strong> dividend = 7, divisor = -3
<strong>Output:</strong> -2
<strong>Explanation:</strong> 7/-3 = truncate(-2.33333..) = -2.
</pre>

**Note:**


- Both dividend and divisor will be 32-bit signed integers.
- The divisor will never be 0.
- Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For the purpose of this problem, assume that your function **returns 231 − 1 when the division result overflows**.
## 解题思路## 可能的變化