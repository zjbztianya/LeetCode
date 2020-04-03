# [190. Reverse Bits](https://leetcode-cn.com/problems/reverse-bits/)
Reverse bits of a given 32 bits unsigned integer.



**Example 1:**


<pre><strong>Input:</strong> 00000010100101000001111010011100
<strong>Output:</strong> 00111001011110000010100101000000
<strong>Explanation: </strong>The input binary string <b>00000010100101000001111010011100</b> represents the unsigned integer 43261596, so return 964176192 which its binary representation is <b>00111001011110000010100101000000</b>.
</pre>

**Example 2:**


<pre><strong>Input:</strong> 11111111111111111111111111111101
<strong>Output:</strong> 10111111111111111111111111111111
<strong>Explanation: </strong>The input binary string <strong>11111111111111111111111111111101</strong> represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is <strong>10111111111111111111111111111111</strong>.</pre>



**Note:**


- Note that in some languages such as Java, there is no unsigned integer type. In this case, both input and output will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using [2&#39;s complement notation](https://en.wikipedia.org/wiki/Two%27s_complement). Therefore, in **Example 2** above the input represents the signed integer <code>-3</code> and the output represents the signed integer <code>-1073741825</code>.



**Follow up**:

If this function is called many times, how would you optimize it?
## 解题思路## 可能的變化