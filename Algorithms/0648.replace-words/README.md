# [648. Replace Words](https://leetcode-cn.com/problems/replace-words/)
In English, we have a concept called <code>root</code>, which can be followed by some other words to form another longer word - let&#39;s call this word <code>successor</code>. For example, the root <code>an</code>, followed by <code>other</code>, which can form another word <code>another</code>.

Now, given a dictionary consisting of many roots and a sentence. You need to replace all the <code>successor</code> in the sentence with the <code>root</code> forming it. If a <code>successor</code> has many <code>roots</code> can form it, replace it with the root with the shortest length.

You need to output the sentence after the replacement.

**Example 1:**


<pre><b>Input:</b> dict = [&#34;cat&#34;, &#34;bat&#34;, &#34;rat&#34;]
sentence = &#34;the cattle was rattled by the battery&#34;
<b>Output:</b> &#34;the cat was rat by the bat&#34;
</pre>



**Note:**

- The input will only have lower-case letters.
- 1 &lt;= dict words number &lt;= 1000
- 1 &lt;= sentence words number &lt;= 1000
- 1 &lt;= root length &lt;= 100
- 1 &lt;= sentence words length &lt;= 1000

## 解题思路## 可能的變化