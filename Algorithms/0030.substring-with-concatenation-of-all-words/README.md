# [30. Substring with Concatenation of All Words](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/)
You are given a string, **s**, and a list of words, **words**, that are all of the same length. Find all starting indices of substring(s) in **s** that is a concatenation of each word in **words** exactly once and without any intervening characters.



**Example 1:**


<pre><strong>Input:
  s =</strong> &#34;barfoothefoobarman&#34;,
<strong>  words = </strong>[&#34;foo&#34;,&#34;bar&#34;]
<strong>Output:</strong> <code>[0,9]</code>
<strong>Explanation:</strong> Substrings starting at index 0 and 9 are &#34;barfoo&#34; and &#34;foobar&#34; respectively.
The output order does not matter, returning [9,0] is fine too.
</pre>

**Example 2:**


<pre><strong>Input:
  s =</strong> &#34;wordgoodgoodgoodbestword&#34;,
<strong>  words = </strong>[&#34;word&#34;,&#34;good&#34;,&#34;best&#34;,&#34;word&#34;]
<strong>Output:</strong> <code>[]</code>
</pre>

## 解题思路## 可能的變化