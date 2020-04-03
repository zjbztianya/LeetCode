# [460. LFU Cache](https://leetcode-cn.com/problems/lfu-cache/)
Design and implement a data structure for [Least Frequently Used (LFU)](https://en.wikipedia.org/wiki/Least_frequently_used) cache. It should support the following operations: <code>get</code> and <code>put</code>.

<code>get(key)</code> - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.<code>put(key, value)</code> - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least frequently used item before inserting a new item. For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency), the least **recently** used key would be evicted.

Note that the number of times an item is used is the number of calls to the <code>get</code> and <code>put</code> functions for that item since it was inserted. This number is set to zero when the item is removed.



**Follow up:**Could you do both operations in **O(1)** time complexity?



**Example:**


<pre>LFUCache cache = new LFUCache( 2 /* capacity */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.get(3);       // returns 3.
cache.put(4, 4);    // evicts key 1.
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
</pre>


## 解题思路## 可能的變化