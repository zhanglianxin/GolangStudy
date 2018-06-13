Each variable in Go has a specific type, which determines the size and layout of the variable's memory, the range of values that can be stored within that memory, and the set of operations that can be applied to the variable.

## 变量定义

定义的时候指定类型

初始化时自动判断类型

## 静态类型声明

变量声明只在编译时期有意义，编译器在连接时期需要真实的变量类型。


## 动态类型声明/类型接口

动态类型变量声明需要编译器基于传递的值来解释变量的类型。编译器不需要固定类型的变量作为必备条件。

## 混合变量声明

使用类型推断一次可以声明不同类型的变量。

## 左值和右值

指向一个内存位置的表达式叫“左值”。左值可以出现在等号赋值的两边。

指向存储在内存中的数据值叫“右值”。右值不可以被赋值，只能出现在等号右边。
