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

## Go 语言数组

Go语言提供了数组类型的数据结构。

数组是具有相同唯一类型的一组以编号且长度固定的数据项序列，这种类型可以是人一的原始类型：例如整型、字串或者自定义类型。

相对于去声明 number0,number1,...,number99 的变量，使用数组形式 numbers[0],numbers[1]...,numbers[99] 更加方便且易于扩展。

数组元素可以通过索引(位置)来读取(或者修改)，索引从 0 开始，第二个索引为 1 ，依次类推。

![test](https://www.runoob.com/wp-content/uploads/2015/06/goarray.png "GO 语言数组")

### 声明数组
Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：

```
var variable_name [size] varaible_type
```
以上为一维数组的定义方式，例如以下定义了数组 balance 长度为 10 类型为 float32:
```
var balance [10] float32
```

### 初始化数组
以下演示了数组初始化：
```
var balance =[5] float32{1000.0 , 2.0 , 3.4 , 7.0 , 50.0}

balace := [5]float32{1000.0 , 2.0 , 3.4 , 7.0 , 50.0}
``` 
如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
```
var balance = [...] float32{1000.0 , 2.0 , 3.4 , 7.0 , 50.0}
或
balace := [...] float32{1000.0 , 2.0 , 3.4 , 7.0 , 50.0}
```


### 访问数组元素
数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：
```
var salary float32 = balance[9]
```
以上实例读取了数组 balance 第 10 个元素的值。

以下演示了数组完整操作（声明、赋值、访问）的实例：
```
package main

import "fmt"

func main() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i,j int

   /* 为数组 n 初始化元素 */        
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}

以上实例执行结果如下：

Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109
```


## Go 语言切片(slice)
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太使用，Go