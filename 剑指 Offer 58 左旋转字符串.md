字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

 
示例 1：
```
输入: s = "abcdefg", k = 2
输出: "cdefgab"
```
示例 2：
```
输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"
```

限制：

1 <= k < s.length <= 10000

来源：力扣（LeetCode）

链接：https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

```
func reverseLeftWords(s string, n int) string {
   return s[n:] + s[:n]
   }
}
```
### 补充知识:字符串

一个字符串是一个不可改变的字节序列，字符串通常是用来包含人类可读的文本数据。和数组不同的是，字符串的元素不可修改，是一个只读的子接数组。每个字符串的长度虽然也是固定的，但是字符串的长度并不是字符串类型的一部分。

Go 语言字符串的底层结构在 reflect.StringHeader 中定义:
```
type StringHeader struct{
   Data uintptr
   Len int
}
```
字符串结构由两个信息组成：第一个是字串指向的底层字节数组，第二个是字符串的字节的长度。字符串其实是一个结构体，因此字符串都得赋值操作也就是 reflect.StringHeader 结构体的复制过程，并不会设计底层字节数组的赋值。

我们看看字符串 "Hello World" 本身对应的内存结构：
![Hello World](https://img2018.cnblogs.com/blog/692143/201812/692143-20181208091330093-555881545.png)

字符串虽然不是切片，但是支持切片操作，不同位置的切片底层也访问的同一块内存数据(因为字符串是只读的，相同的字符串面值通常是对应的同一个字符串常量)：
```
s := "hello world"
hello := s[:5]
world := s[7:]
s1 := "hello world"[:5]
s2 := "hello world"[7:]
```

字符串和数组类似，内置的 len 函数返回字符串的长度。也可以通过 reflect.StringHeader 结构访问字符串的长度
```
fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 12
fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5
```
如果不想解码 UTF-8 字符串，想直接遍历原始的字节码，可以将字符串强制转为 []byte 字节序列后再进行便利(这里的转换一般不会产生运行时开销)：
```
for i, c := range []byte("世界abc") {
fmt.Println(i, c)
}
```