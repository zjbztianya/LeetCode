# [646. Maximum Length of Pair Chain](https://leetcode-cn.com/problems/maximum-length-of-pair-chain/)
You are given <code>n</code> pairs of numbers. In every pair, the first number is always smaller than the second number.

Now, we define a pair <code>(c, d)</code> can follow another pair <code>(a, b)</code> if and only if <code>b &lt; c</code>. Chain of pairs can be formed in this fashion. 

Given a set of pairs, find the length longest chain which can be formed. You needn&#39;t use up all the given pairs. You can select pairs in any order.

**Example 1:**


<pre><b>Input:</b> [[1,2], [2,3], [3,4]]
<b>Output:</b> 2
<b>Explanation:</b> The longest chain is [1,2] -&gt; [3,4]
</pre>



**Note:**

- The number of given pairs will be in the range [1, 1000].

## 解题思路## 可能的變化