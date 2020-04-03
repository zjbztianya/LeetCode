# [75. Sort Colors](https://leetcode-cn.com/problems/sort-colors/)
Given an array with _n_ objects colored red, white or blue, sort them **[in-place](https://en.wikipedia.org/wiki/In-place_algorithm)**so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

**Note:** You are not suppose to use the library&#39;s sort function for this problem.

**Example:**


<pre><strong>Input:</strong> [2,0,2,1,1,0]
<strong>Output:</strong> [0,0,1,1,2,2]</pre>

**Follow up:**


- A rather straight forward solution is a two-pass algorithm using counting sort.First, iterate the array counting number of 0&#39;s, 1&#39;s, and 2&#39;s, then overwrite array with total number of 0&#39;s, then 1&#39;s and followed by 2&#39;s.
- Could you come up with a one-pass algorithm using only constant space?
## 解题思路## 可能的變化