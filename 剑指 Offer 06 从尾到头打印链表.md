输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

 

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
 

限制：

0 <= 链表长度 <= 10000

来源：力扣（LeetCode）

链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {

}
```

补充知识：

Go 语言数组

Go语言提供了数组类型的数据结构。

数组是具有相同唯一类型的一组以编号且长度固定的数据项序列，这种类型可以是人一的原始类型：例如整型、字串或者自定义类型。

相对于去声明 number0,number1,...,number99 的变量，使用数组形式 numbers[0],numbers[1]...,numbers[99] 更加方便且易于扩展。

数组元素可以通过索引(位置)来读取(或者修改)，索引从 0 开始，第二个索引为 1 ，依次类推。

![GO 语言数组](https://www.runoob.com/wp-content/uploads/2015/06/goarray.png "GO 语言数组")