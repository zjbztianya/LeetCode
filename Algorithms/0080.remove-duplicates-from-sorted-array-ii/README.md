# [80. Remove Duplicates from Sorted Array II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/)
Given a sorted array _nums_, remove the duplicates [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm) such that duplicates appeared at most _twice_ and return the new length.

Do not allocate extra space for another array, you must do this by **modifying the input array [in-place](https://en.wikipedia.org/wiki/In-place_algorithm)** with O(1) extra memory.

**Example 1:**


<pre>Given <em>nums</em> = <strong>[1,1,1,2,2,3]</strong>,

Your function should return length = <strong><code>5</code></strong>, with the first five elements of <em><code>nums</code></em> being <strong><code>1, 1, 2, 2</code></strong> and <strong>3</strong> respectively.

It doesn&#39;t matter what you leave beyond the returned length.</pre>

**Example 2:**


<pre>Given <em>nums</em> = <strong>[0,0,1,1,1,1,2,3,3]</strong>,

Your function should return length = <strong><code>7</code></strong>, with the first seven elements of <em><code>nums</code></em> being modified to <strong><code>0</code></strong>, <strong>0</strong>, <strong>1</strong>, <strong>1</strong>, <strong>2</strong>, <strong>3</strong> and <strong>3</strong> respectively.

It doesn&#39;t matter what values are set beyond the returned length.
</pre>

**Clarification:**

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by **reference**, which means modification to the input array will be known to the caller as well.

Internally you can think of this:


<pre>// <strong>nums</strong> is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to <strong>nums</strong> in your function would be known by the caller.
// using the length returned by your function, it prints the first <strong>len</strong> elements.
for (int i = 0; i &lt; len; i++) {
    print(nums[i]);
}
</pre>

## 解题思路## 可能的變化