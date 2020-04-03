# [130. Surrounded Regions](https://leetcode-cn.com/problems/surrounded-regions/)
Given a 2D board containing <code>&#39;X&#39;</code> and <code>&#39;O&#39;</code> (**the letter O**), capture all regions surrounded by <code>&#39;X&#39;</code>.

A region is captured by flipping all <code>&#39;O&#39;</code>s into <code>&#39;X&#39;</code>s in that surrounded region.

**Example:**


<pre>X X X X
X O O X
X X O X
X O X X
</pre>

After running your function, the board should be:


<pre>X X X X
X X X X
X X X X
X O X X
</pre>

**Explanation:**

Surrounded regions shouldn’t be on the border, which means that any <code>&#39;O&#39;</code> on the border of the board are not flipped to <code>&#39;X&#39;</code>. Any <code>&#39;O&#39;</code> that is not on the border and it is not connected to an <code>&#39;O&#39;</code> on the border will be flipped to <code>&#39;X&#39;</code>. Two cells are connected if they are adjacent cells connected horizontally or vertically.
## 解题思路## 可能的變化