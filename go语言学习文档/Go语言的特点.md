# Go语言的特点

1、Go语言保证了静态语言的安全和性能，并且还具有动态语言开发维护的高效率，其实我们简介点来说Go语言的话，就相当于Go = Python+C语言，这样表名了Go语言语言既代表了C语言静态程序运行的速度，并且还带有其他譬如Python语言的快速开发。

2、其实Go语言继承了C语言中的很多理念，就像我们在大学学的时候的一些控制结构和基本的数据类型、一些表达式语法，最重要的就是指针，其中Go语言有两个比较重要的时刻，就是2009年发布第一个开源版本，2015年移除最后的C语言的特征。

3、其实Go语言也有点类似于Java，其中还是存在包package的概念，每一个Go的文件都隶属于一个package中，并且其中存在的垃圾回收机制是不是类似于Java中的垃圾回收呢？

4、还有一个特点就是Go语言本身支持并发，实现起来难度不大，其中的goroutine是一个轻量级的线程，可以实现大并发的处理，利用多核。基于CPS并发模型（信息物理系统）实现。



5、部分程序代码

```
小试牛刀:
package main
import "fmt"

func main(){
   
   
}
//多个返回值和参数的类型
()中代表的是参数的返回类型
funt getAll(n1 int,n2 int)(int,int){
   
   sum := n1+n2
   sun := n1-n2
   return sum,sub
}




```

```
//分别测试指针和值接收（含有结构体）
package main

//go语言的结构体是值类型
type User struct{
	name string
	age int
}

//使用值接收
func (user User) testValueRecv(){
	 //user.name = "张三"
	 //user.age = 21
	 println("值接收姓名=",user.name)
	 println("值接收年龄 =",user.age)
}

//使用指针接收
func (user *User) testPointerRecv(){
	//user.name = "王思聪"
	//user.age = 32
	println("指针接收姓名=",user.name)
	println("指针接收年龄 =",user.age)
}
//测试Go语言的指针和结构体特性
func main() {

	 //调用值接收函数
	 user1 := User{"张三",21}
	 user1.testValueRecv()

	 //调用指针接收函数
	 user2 := &User{"王思聪",32}
	 user2.testPointerRecv()	
}
```

5.1 以下是部分代码运行结果:

<img src="C:\Users\yuanfeng\AppData\Roaming\Typora\typora-user-images\image-20191116173912603.png" alt="image-20191116173912603" style="zoom:67%;" />

<img src="C:\Users\yuanfeng\AppData\Roaming\Typora\typora-user-images\image-20191116175403761.png" alt="image-20191116175403761" style="zoom:67%;" />



6、介绍以下Go的目录结构

goproject/src/go_code/project01/main--package

(1) go文件目录是.go

(2) package main : 表明该hello.go文件所在的包是main,在go中每个文件都必须属于一个包

(3)import "fmt": 表明引入了fmt可以使用其中的函数

(4)fmt.Println("Hello World !") 

其中go build可以编译go文件为.exe的文件,有时候我们也可以直接执行go run，还是非常可取的！