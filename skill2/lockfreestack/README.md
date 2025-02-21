# 无锁栈
## push
假设初始栈为空，top为nil。

插入item1：item1.next = nil，top指向item1。

插入item2：item2.next = item1，top指向item2。

插入item3：item3.next = item2，top指向item3。

最终栈的结构为：item3 -> item2 -> item1 -> nil。

## pop 
(*Node)(top) 是 Go 语言中的类型转换语法，将 top（一个 unsafe.Pointer 类型）显式转换为 *Node 类型。 .(*Node) 是 Go 语言中的类型断言语法，用于从接口类型中提取具体类型。
在这里，top 是一个 unsafe.Pointer，而不是接口类型，因此需要使用类型转换 (*Node)(top)，而不是类型断言 .(*Node)。


## unittest
```
go test -v                      
=== RUN   TestNewLockFreeStack
--- PASS: TestNewLockFreeStack (0.00s)
=== RUN   TestPush
--- PASS: TestPush (0.00s)
=== RUN   TestPop
--- PASS: TestPop (0.00s)
PASS
ok      skill2/lockfreestack    0.670s
```

### 用例分析

* NewLockFreeStack:
  * 测试用例：创建一个新的无锁栈实例，检查返回的指针是否为非空。
* Push:
  * 测试用例1：向空栈中压入一个元素，检查栈顶元素是否为压入的元素。
  * 测试用例2：向非空栈中压入多个元素，检查栈顶元素是否为最后压入的元素。
* Pop:
  * 测试用例1：从空栈中弹出元素，检查返回的布尔值是否为false。
  * 测试用例2：从非空栈中弹出元素，检查返回的元素是否为栈顶元素，并且栈顶元素被正确移除。

### 代码注释

* TestNewLockFreeStack: 测试NewLockFreeStack函数，确保返回的栈实例不为空。
* TestPush: 测试Push函数，包括向空栈和非空栈中压入元素的情况。
* TestPop: 测试Pop函数，包括从空栈和非空栈中弹出元素的情况，并检查栈顶元素是否被正确移除。