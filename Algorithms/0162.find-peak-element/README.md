# [162. Find Peak Element](https://leetcode-cn.com/problems/find-peak-element/)
A peak element is an element that is greater than its neighbors.

Given an input array <code>nums</code>, where <code>nums[i] ≠ nums[i+1]</code>, find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that <code>nums[-1] = nums[n] = -∞</code>.

**Example 1:**


<pre><strong>Input:</strong> <strong>nums</strong> = <code>[1,2,3,1]</code>
<strong>Output:</strong> 2
<strong>Explanation:</strong> 3 is a peak element and your function should return the index number 2.</pre>

**Example 2:**


<pre><strong>Input:</strong> <strong>nums</strong> = <code>[</code>1,2,1,3,5,6,4]
<strong>Output:</strong> 1 or 5 
<strong>Explanation:</strong> Your function can return either index number 1 where the peak element is 2, 
             or index number 5 where the peak element is 6.
</pre>

**Note:**

Your solution should be in logarithmic complexity.
## 解题思路## 可能的變化