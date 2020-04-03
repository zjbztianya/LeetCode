# [640. Solve the Equation](https://leetcode-cn.com/problems/solve-the-equation/)
Solve a given equation and return the value of <code>x</code> in the form of string &#34;x=#value&#34;. The equation contains only &#39;+&#39;, &#39;-&#39; operation, the variable <code>x</code> and its coefficient.

If there is no solution for the equation, return &#34;No solution&#34;.

If there are infinite solutions for the equation, return &#34;Infinite solutions&#34;.

If there is exactly one solution for the equation, we ensure that the value of <code>x</code> is an integer.

**Example 1:**


<pre><b>Input:</b> &#34;x+5-3+x=6+x-2&#34;
<b>Output:</b> &#34;x=2&#34;
</pre>



**Example 2:**


<pre><b>Input:</b> &#34;x=x&#34;
<b>Output:</b> &#34;Infinite solutions&#34;
</pre>



**Example 3:**


<pre><b>Input:</b> &#34;2x=x&#34;
<b>Output:</b> &#34;x=0&#34;
</pre>



**Example 4:**


<pre><b>Input:</b> &#34;2x+3x-6x=x+2&#34;
<b>Output:</b> &#34;x=-1&#34;
</pre>



**Example 5:**


<pre><b>Input:</b> &#34;x=x+2&#34;
<b>Output:</b> &#34;No solution&#34;
</pre>


## 解题思路## 可能的變化