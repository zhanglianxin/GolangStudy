## 指针

指针是一个变量，它的值是其它变量的地址。使用之前必须先声明。

```go
var ip *int // pointer to an interge
var fp *float32 // pointer to a float
```

所有指针的值的实际数据类型都是相同的（long hex number），一个表示内存地址的长十六进制数。不同数据类型的指针之间唯一的区别是指针所指向的变量或常量的数据类型。
