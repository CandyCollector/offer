用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

 

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

1 <= values <= 10000 最多会对 appendTail、deleteHead 进行 10000 次调用

来源：力扣（LeetCode）

链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

来源：力扣（LeetCode）

链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

```
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
```

补充内容 List：
```
list包实现了双向链表。要遍历一个链表：
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}

type Element
type Element struct {
    // 元素保管的值
    Value interface{}
    // 内含隐藏或非导出字段
}
Element类型代表是双向链表的一个元素。

func (*Element) Next
func (e *Element) Next() *Element
Next返回链表的后一个元素或者nil。

func (*Element) Prev
func (e *Element) Prev() *Element
Prev返回链表的前一个元素或者nil。

type List
type List struct {
    // 内含隐藏或非导出字段
}
List代表一个双向链表。List零值为一个空的、可用的链表。

func New
func New() *List
New创建一个链表。

func (*List) Init
func (l *List) Init() *List
Init清空链表。

func (*List) Len
func (l *List) Len() int
Len返回链表中元素的个数，复杂度O(1)。

func (*List) Front
func (l *List) Front() *Element
Front返回链表第一个元素或nil。

func (*List) Back
func (l *List) Back() *Element
Back返回链表最后一个元素或nil。

func (*List) PushFront
func (l *List) PushFront(v interface{}) *Element
PushBack将一个值为v的新元素插入链表的第一个位置，返回生成的新元素。

func (*List) PushFrontList
func (l *List) PushFrontList(other *List)
PushFrontList创建链表other的拷贝，并将拷贝的最后一个位置连接到链表l的第一个位置。

func (*List) PushBack
func (l *List) PushBack(v interface{}) *Element
PushBack将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素。

func (*List) PushBackList
func (l *List) PushBackList(other *List)
PushBack创建链表other的拷贝，并将链表l的最后一个位置连接到拷贝的第一个位置。

func (*List) InsertBefore
func (l *List) InsertBefore(v interface{}, mark *Element) *Element
InsertBefore将一个值为v的新元素插入到mark前面，并返回生成的新元素。如果mark不是l的元素，l不会被修改。

func (*List) InsertAfter
func (l *List) InsertAfter(v interface{}, mark *Element) *Element
InsertAfter将一个值为v的新元素插入到mark后面，并返回新生成的元素。如果mark不是l的元素，l不会被修改。

func (*List) MoveToFront
func (l *List) MoveToFront(e *Element)
MoveToFront将元素e移动到链表的第一个位置，如果e不是l的元素，l不会被修改。

func (*List) MoveToBack
func (l *List) MoveToBack(e *Element)
MoveToBack将元素e移动到链表的最后一个位置，如果e不是l的元素，l不会被修改。

func (*List) MoveBefore
func (l *List) MoveBefore(e, mark *Element)
MoveBefore将元素e移动到mark的前面。如果e或mark不是l的元素，或者e==mark，l不会被修改。

func (*List) MoveAfter
func (l *List) MoveAfter(e, mark *Element)
MoveAfter将元素e移动到mark的后面。如果e或mark不是l的元素，或者e==mark，l不会被修改。

func (*List) Remove
func (l *List) Remove(e *Element) interface{}
Remove删除链表中的元素e，并返回e.Value。
```