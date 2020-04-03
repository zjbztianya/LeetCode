# [139. Word Break](https://leetcode-cn.com/problems/word-break/)
Given a **non-empty** string _s_ and a dictionary _wordDict_ containing a list of **non-empty** words, determine if _s_ can be segmented into a space-separated sequence of one or more dictionary words.

**Note:**


- The same word in the dictionary may be reused multiple times in the segmentation.
- You may assume the dictionary does not contain duplicate words.

**Example 1:**


<pre><strong>Input:</strong> s = &#34;leetcode&#34;, wordDict = [&#34;leet&#34;, &#34;code&#34;]
<strong>Output:</strong> true
<strong>Explanation:</strong> Return true because <code>&#34;leetcode&#34;</code> can be segmented as <code>&#34;leet code&#34;</code>.
</pre>

**Example 2:**


<pre><strong>Input:</strong> s = &#34;applepenapple&#34;, wordDict = [&#34;apple&#34;, &#34;pen&#34;]
<strong>Output:</strong> true
<strong>Explanation:</strong> Return true because <code>&#34;</code>applepenapple<code>&#34;</code> can be segmented as <code>&#34;</code>apple pen apple<code>&#34;</code>.
             Note that you are allowed to reuse a dictionary word.
</pre>

**Example 3:**


<pre><strong>Input:</strong> s = &#34;catsandog&#34;, wordDict = [&#34;cats&#34;, &#34;dog&#34;, &#34;sand&#34;, &#34;and&#34;, &#34;cat&#34;]
<strong>Output:</strong> false
</pre>

## 解题思路## 可能的變化