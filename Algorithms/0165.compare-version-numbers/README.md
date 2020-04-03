# [165. Compare Version Numbers](https://leetcode-cn.com/problems/compare-version-numbers/)
Compare two version numbers _version1_ and _version2_.If <code><em>version1</em> &gt; <em>version2</em></code> return <code>1;</code> if <code><em>version1</em> &lt; <em>version2</em></code> return <code>-1;</code>otherwise return <code>0</code>.

You may assume that the version strings are non-empty and contain only digits and the <code>.</code> character.

The <code>.</code> character does not represent a decimal point and is used to separate number sequences.

For instance, <code>2.5</code> is not &#34;two and a half&#34; or &#34;half way to version three&#34;, it is the fifth second-level revision of the second first-level revision.

You may assume the default revision number for each level of a version number to be <code>0</code>. For example, version number <code>3.4</code> has a revision number of <code>3</code> and <code>4</code> for its first and second level revision number. Its third and fourth level revision number are both <code>0</code>.



**Example 1:**


<pre><strong>Input:</strong> <code><em>version1</em></code> = &#34;0.1&#34;, <code><em>version2</em></code> = &#34;1.1&#34;
<strong>Output:</strong> -1</pre>

**Example 2:**


<pre><strong>Input: </strong><code><em>version1</em></code> = &#34;1.0.1&#34;, <code><em>version2</em></code> = &#34;1&#34;
<strong>Output:</strong> 1</pre>

**Example 3:**


<pre><strong>Input:</strong> <code><em>version1</em></code> = &#34;7.5.2.4&#34;, <code><em>version2</em></code> = &#34;7.5.3&#34;
<strong>Output:</strong> -1</pre>

**Example 4:**


<pre><strong>Input:</strong> <code><em>version1</em></code> = &#34;1.01&#34;, <code><em>version2</em></code> = &#34;1.001&#34;
<strong>Output:</strong> 0
<strong>Explanation:</strong> Ignoring leading zeroes, both “01” and “001&#34; represent the same number “1”</pre>

**Example 5:**


<pre><strong>Input:</strong> <code><em>version1</em></code> = &#34;1.0&#34;, <code><em>version2</em></code> = &#34;1.0.0&#34;
<strong>Output:</strong> 0
<strong>Explanation:</strong> The first version number does not have a third level revision number, which means its third level revision number is default to &#34;0&#34;</pre>



**Note:**

- Version strings are composed of numeric strings separated by dots <code>.</code> and this numeric strings **may** have leading zeroes. 
- Version strings do not start or end with dots, and they will not be two consecutive dots.## 解题思路## 可能的變化