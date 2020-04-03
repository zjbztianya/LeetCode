# [140. Word Break II](https://leetcode-cn.com/problems/word-break-ii/)
Given a **non-empty** string _s_ and a dictionary _wordDict_ containing a list of **non-empty** words, add spaces in _s_ to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

**Note:**


- The same word in the dictionary may be reused multiple times in the segmentation.
- You may assume the dictionary does not contain duplicate words.

**Example 1:**


<pre><strong>Input:
</strong>s = &#34;<code>catsanddog</code>&#34;
wordDict = <code>[&#34;cat&#34;, &#34;cats&#34;, &#34;and&#34;, &#34;sand&#34;, &#34;dog&#34;]</code>
<strong>Output:
</strong><code>[
  &#34;cats and dog&#34;,
  &#34;cat sand dog&#34;
]</code>
</pre>

**Example 2:**


<pre><strong>Input:
</strong>s = &#34;pineapplepenapple&#34;
wordDict = [&#34;apple&#34;, &#34;pen&#34;, &#34;applepen&#34;, &#34;pine&#34;, &#34;pineapple&#34;]
<strong>Output:
</strong>[
  &#34;pine apple pen apple&#34;,
  &#34;pineapple pen apple&#34;,
  &#34;pine applepen apple&#34;
]
<strong>Explanation:</strong> Note that you are allowed to reuse a dictionary word.
</pre>

**Example 3:**


<pre><strong>Input:
</strong>s = &#34;catsandog&#34;
wordDict = [&#34;cats&#34;, &#34;dog&#34;, &#34;sand&#34;, &#34;and&#34;, &#34;cat&#34;]
<strong>Output:
</strong>[]</pre>

## 解题思路## 可能的變化