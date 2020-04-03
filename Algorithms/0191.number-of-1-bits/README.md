# [191. Number of 1 Bits](https://leetcode-cn.com/problems/number-of-1-bits/)
Write a function that takes an unsigned integer and return the number of &#39;1&#39; bits it has (also known as the [Hamming weight](http://en.wikipedia.org/wiki/Hamming_weight)).



**Example 1:**


<pre><strong>Input:</strong> 00000000000000000000000000001011
<strong>Output:</strong> 3
<strong>Explanation: </strong>The input binary string <code><strong>00000000000000000000000000001011</strong> has a total of three &#39;1&#39; bits.</code>
</pre>

**Example 2:**


<pre><strong>Input:</strong> 00000000000000000000000010000000
<strong>Output:</strong> 1
<strong>Explanation: </strong>The input binary string <strong>00000000000000000000000010000000</strong> has a total of one &#39;1&#39; bit.
</pre>

**Example 3:**


<pre><strong>Input:</strong> 11111111111111111111111111111101
<strong>Output:</strong> 31
<strong>Explanation: </strong>The input binary string <strong>11111111111111111111111111111101</strong> has a total of thirty one &#39;1&#39; bits.</pre>



**Note:**


- Note that in some languages such as Java, there is no unsigned integer type. In this case, the input will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using [2&#39;s complement notation](https://en.wikipedia.org/wiki/Two%27s_complement). Therefore, in **Example 3** above the input represents the signed integer <code>-3</code>.



**Follow up**:

If this function is called many times, how would you optimize it?
## 解题思路## 可能的變化