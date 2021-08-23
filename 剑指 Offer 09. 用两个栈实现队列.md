#用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

 

示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
通过次数233,446提交次数325,216

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

···
type CQueue struct {
    s1, s2 *list.List
}

func Constructor() CQueue {
    return CQueue{
        s1: list.New(),
        s2: list.New(),
    }
}

func (this *CQueue) AppendTail(value int)  {
    this.s1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
    // 如果第二个栈为空
    if this.s2.Len() == 0 {
        for this.s1.Len() > 0 {
            this.s2.PushBack(this.s1.Remove(this.s1.Back()))
        }
    }
    if this.s2.Len() != 0 {
        e := this.s2.Back()
        this.s2.Remove(e)
        return e.Value.(int)
    }
    return -1
}
···