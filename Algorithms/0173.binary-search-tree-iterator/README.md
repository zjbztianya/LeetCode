# [173. Binary Search Tree Iterator](https://leetcode-cn.com/problems/binary-search-tree-iterator/)
Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling <code>next()</code> will return the next smallest number in the BST.





**Example:**

**![](https://assets.leetcode.com/uploads/2018/12/25/bst-tree.png)**


<pre>BSTIterator iterator = new BSTIterator(root);
iterator.next();    // return 3
iterator.next();    // return 7
iterator.hasNext(); // return true
iterator.next();    // return 9
iterator.hasNext(); // return true
iterator.next();    // return 15
iterator.hasNext(); // return true
iterator.next();    // return 20
iterator.hasNext(); // return false
</pre>



**Note:**


- <code>next()</code> and <code>hasNext()</code> should run in average O(1) time and uses O(_h_) memory, where _h_ is the height of the tree.
- You may assume that <code>next()</code> call will always be valid, that is, there will be at least a next smallest number in the BST when <code>next()</code> is called.
## 解题思路## 可能的變化