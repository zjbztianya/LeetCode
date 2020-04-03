# [155. Min Stack](https://leetcode-cn.com/problems/min-stack/)
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.


- push(x) -- Push element x onto stack.
- pop() -- Removes the element on top of the stack.
- top() -- Get the top element.
- getMin() -- Retrieve the minimum element in the stack.



**Example:**


<pre>MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --&gt; Returns -3.
minStack.pop();
minStack.top();      --&gt; Returns 0.
minStack.getMin();   --&gt; Returns -2.
</pre>


## 解题思路## 可能的變化