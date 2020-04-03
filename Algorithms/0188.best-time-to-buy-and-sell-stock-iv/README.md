# [188. Best Time to Buy and Sell Stock IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)
Say you have an array for which the _i-_th element is the price of a given stock on day _i_.

Design an algorithm to find the maximum profit. You may complete at most **k** transactions.

**Note:**You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

**Example 1:**


<pre><strong>Input:</strong> [2,4,1], k = 2
<strong>Output:</strong> 2
<strong>Explanation:</strong> Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
</pre>

**Example 2:**


<pre><strong>Input:</strong> [3,2,6,5,0,3], k = 2
<strong>Output:</strong> 7
<strong>Explanation:</strong> Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
             Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
</pre>

## 解题思路## 可能的變化