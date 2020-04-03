# [79. Word Search](https://leetcode-cn.com/problems/word-search/)
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where &#34;adjacent&#34; cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

**Example:**


<pre>board =
[
  [&#39;A&#39;,&#39;B&#39;,&#39;C&#39;,&#39;E&#39;],
  [&#39;S&#39;,&#39;F&#39;,&#39;C&#39;,&#39;S&#39;],
  [&#39;A&#39;,&#39;D&#39;,&#39;E&#39;,&#39;E&#39;]
]

Given word = &#34;<strong>ABCCED</strong>&#34;, return <strong>true</strong>.
Given word = &#34;<strong>SEE</strong>&#34;, return <strong>true</strong>.
Given word = &#34;<strong>ABCB</strong>&#34;, return <strong>false</strong>.
</pre>



**Constraints:**


- <code>board</code> and <code>word</code> consists only of lowercase and uppercase English letters.
- <code>1 &lt;= board.length &lt;= 200</code>
- <code>1 &lt;= board[i].length &lt;= 200</code>
- <code>1 &lt;= word.length &lt;= 10^3</code>
## 解题思路## 可能的變化