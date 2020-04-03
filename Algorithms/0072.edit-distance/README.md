# [72. Edit Distance](https://leetcode-cn.com/problems/edit-distance/)
Given two words _word1_ and _word2_, find the minimum number of operations required to convert _word1_ to _word2_.

You have the following 3 operations permitted on a word:

- Insert a character
- Delete a character
- Replace a character
**Example 1:**


<pre><strong>Input:</strong> word1 = &#34;horse&#34;, word2 = &#34;ros&#34;
<strong>Output:</strong> 3
<strong>Explanation:</strong> 
horse -&gt; rorse (replace &#39;h&#39; with &#39;r&#39;)
rorse -&gt; rose (remove &#39;r&#39;)
rose -&gt; ros (remove &#39;e&#39;)
</pre>

**Example 2:**


<pre><strong>Input:</strong> word1 = &#34;intention&#34;, word2 = &#34;execution&#34;
<strong>Output:</strong> 5
<strong>Explanation:</strong> 
intention -&gt; inention (remove &#39;t&#39;)
inention -&gt; enention (replace &#39;i&#39; with &#39;e&#39;)
enention -&gt; exention (replace &#39;n&#39; with &#39;x&#39;)
exention -&gt; exection (replace &#39;n&#39; with &#39;c&#39;)
exection -&gt; execution (insert &#39;u&#39;)
</pre>

## 解题思路## 可能的變化