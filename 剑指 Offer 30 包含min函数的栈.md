定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

 

示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.
 

提示：

各函数的调用总次数不超过 20000 次

来源：力扣（LeetCode）

链接：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



```
type MinStack struct {
    stack []int
    min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        min: []int{},
    }
}


func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if len(this.min) == 0 || x <= this.min[len(this.min)-1] { // 没有元素或者x小于最小值的时候，append到min最后
        this.min = append(this.min, x)
    }else {
        this.min = append(this.min, this.min[len(this.min)-1]) // 否则继续将上一个最小值append到min
    }
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
    this.min = this.min[:len(this.min)-1]
}


func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
        return 0
    }
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) Min() int {
    if len(this.min) > 0 {
        return this.min[len(this.min)-1]
    }
    return 0
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

```

解题思路：
普通栈的 push() 和 pop() 函数的复杂度为 O(1)O(1) ；而获取栈最小值 min() 函数需要遍历整个栈，复杂度为 O(N)O(N) 。

本题难点： 将 min() 函数复杂度降为 O(1)O(1) ，可通过建立辅助栈实现；
数据栈 AA ： 栈 AA 用于存储所有元素，保证入栈 push() 函数、出栈 pop() 函数、获取栈顶 top() 函数的正常逻辑。
辅助栈 BB ： 栈 BB 中存储栈 AA 中所有 非严格降序 的元素，则栈 AA 中的最小元素始终对应栈 BB 的栈顶元素，即 min() 函数只需返回栈 BB 的栈顶元素即可。
因此，只需设法维护好 栈 BB 的元素，使其保持非严格降序，即可实现 min() 函数的 O(1)O(1) 复杂度。


函数设计：
push(x) 函数： 重点为保持栈 BB 的元素是 非严格降序 的。

将 xx 压入栈 AA （即 A.add(x) ）；
若 ① 栈 BB 为空 或 ② xx 小于等于 栈 BB 的栈顶元素，则将 xx 压入栈 BB （即 B.add(x) ）。
pop() 函数： 重点为保持栈 A, BA,B 的 元素一致性 。

执行栈 AA 出栈（即 A.pop() ），将出栈元素记为 yy ；
若 yy 等于栈 BB 的栈顶元素，则执行栈 B 出栈（即 B.pop() ）。
top() 函数： 直接返回栈 AA 的栈顶元素即可，即返回 A.peek() 。

min() 函数： 直接返回栈 BB 的栈顶元素即可，即返回 B.peek() 。


复杂度分析：
时间复杂度 O(1)O(1) ： push(), pop(), top(), min() 四个函数的时间复杂度均为常数级别。
空间复杂度 O(N)O(N) ： 当共有 NN 个待入栈元素时，辅助栈 BB 最差情况下存储 NN 个元素，使用 O(N)O(N) 额外空间。
