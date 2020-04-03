# [36. Valid Sudoku](https://leetcode-cn.com/problems/valid-sudoku/)
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated **according to the following rules**:

- Each row must contain the digits <code>1-9</code> without repetition.
- Each column must contain the digits <code>1-9</code> without repetition.
- Each of the 9 <code>3x3</code> sub-boxes of the grid must contain the digits <code>1-9</code> without repetition.
![](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)<small>A partially filled sudoku which is valid.</small>

The Sudoku board could be partially filled, where empty cells are filled with the character <code>&#39;.&#39;</code>.

**Example 1:**


<pre><strong>Input:</strong>
[
  [&#34;5&#34;,&#34;3&#34;,&#34;.&#34;,&#34;.&#34;,&#34;7&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;],
  [&#34;6&#34;,&#34;.&#34;,&#34;.&#34;,&#34;1&#34;,&#34;9&#34;,&#34;5&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;],
  [&#34;.&#34;,&#34;9&#34;,&#34;8&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;6&#34;,&#34;.&#34;],
  [&#34;8&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;6&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;3&#34;],
  [&#34;4&#34;,&#34;.&#34;,&#34;.&#34;,&#34;8&#34;,&#34;.&#34;,&#34;3&#34;,&#34;.&#34;,&#34;.&#34;,&#34;1&#34;],
  [&#34;7&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;2&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;6&#34;],
  [&#34;.&#34;,&#34;6&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;2&#34;,&#34;8&#34;,&#34;.&#34;],
  [&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;4&#34;,&#34;1&#34;,&#34;9&#34;,&#34;.&#34;,&#34;.&#34;,&#34;5&#34;],
  [&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;8&#34;,&#34;.&#34;,&#34;.&#34;,&#34;7&#34;,&#34;9&#34;]
]
<strong>Output:</strong> true
</pre>

**Example 2:**


<pre><strong>Input:</strong>
[
  [&#34;8&#34;,&#34;3&#34;,&#34;.&#34;,&#34;.&#34;,&#34;7&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;],
  [&#34;6&#34;,&#34;.&#34;,&#34;.&#34;,&#34;1&#34;,&#34;9&#34;,&#34;5&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;],
  [&#34;.&#34;,&#34;9&#34;,&#34;8&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;6&#34;,&#34;.&#34;],
  [&#34;8&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;6&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;3&#34;],
  [&#34;4&#34;,&#34;.&#34;,&#34;.&#34;,&#34;8&#34;,&#34;.&#34;,&#34;3&#34;,&#34;.&#34;,&#34;.&#34;,&#34;1&#34;],
  [&#34;7&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;2&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;6&#34;],
  [&#34;.&#34;,&#34;6&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;2&#34;,&#34;8&#34;,&#34;.&#34;],
  [&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;4&#34;,&#34;1&#34;,&#34;9&#34;,&#34;.&#34;,&#34;.&#34;,&#34;5&#34;],
  [&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;.&#34;,&#34;8&#34;,&#34;.&#34;,&#34;.&#34;,&#34;7&#34;,&#34;9&#34;]
]
<strong>Output:</strong> false
<strong>Explanation:</strong> Same as Example 1, except with the <strong>5</strong> in the top left corner being 
    modified to <strong>8</strong>. Since there are two 8&#39;s in the top left 3x3 sub-box, it is invalid.
</pre>

**Note:**


- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.
- The given board contain only digits <code>1-9</code> and the character <code>&#39;.&#39;</code>.
- The given board size is always <code>9x9</code>.
## 解题思路## 可能的變化