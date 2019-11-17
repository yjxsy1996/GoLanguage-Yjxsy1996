# 一、Go语言简介

1. **Go语言歧视是谷歌公司在2009年11月发布的编程语言，其中团队成员有三位都是大人物**
2. **官方的详解：Go语言也成为GoLang语言，是一种开源的程序设计语言，其设计之初就是为了大家能够比较方便简单的开发软件，其中还包含一些c++、c、Java一类的特性。**

# 二、Go语言的一些优势

- 语法来说还是相对简单，类似于js的语法

  ```go
  var a,b = 1,2
  a,b = b,a
  fmt.Println(a,b)
  
  ```

  <img src="C:\Users\yuanfeng\Desktop\QQ截图20191112174548.png" style="zoom: 67%;" />

- 可以直接编译成机器码
- 静态的数据类型和编译语言

  

```go
a := 1
b := false

```

<img src="C:\Users\yuanfeng\Desktop\QQ截图20191112174903.png" style="zoom:67%;" />

- 内置并发

  ```go
  go func(){
     //do something
  }
  ```

  

- 内置垃圾回收
- 部署简单
- 强大的标准库