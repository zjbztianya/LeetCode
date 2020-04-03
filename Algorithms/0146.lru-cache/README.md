# [146. LRU Cache](https://leetcode-cn.com/problems/lru-cache/)
Design and implement a data structure for [Least Recently Used (LRU) cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU). It should support the following operations: <code>get</code> and <code>put</code>.

<code>get(key)</code> - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.<code>put(key, value)</code> - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a **positive** capacity.

**Follow up:**Could you do both operations in **O(1)** time complexity?

**Example:**


<pre>LRUCache cache = new LRUCache( 2 /* capacity */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
</pre>


## 解题思路## 可能的變化