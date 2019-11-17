# Go语言的编译

> 三种平台：Linux、Mac、Windows

在Linux和windows中还是存在一定区别的，我们在Liunx中可以通过go build来进行文件的编译，然后在linux中生成对应的可执行文件，但是首先需要注意到的是和Windows中的区别，Linux中的可惜执行文件不带.exe，windows中是带有.exe的二进制可执行文件，因此我们在Linux中可以使用./hello执行文件，windows中则需要运行hello.exe命令运行；

还有一个问题就是我们可以直接把go build编译好的可执行文件(hello.exe)直接拿到另外一个机器上面运行，并且另外一个机器是不需要有Go SDK的，但是如果我们首先使用的是go run这种形式得到的可执行文件那么就不行了，这个时候必须要求另外一台机器需要满足安装对应的Go环境才行；

## Go语言开发的注意要点：

- Go源文件必须要以.go为文件的扩展名
- Go程序的执行入口为main函数
- Go语言严格区分大小写
- Go语言语句结束不需要分号，但是编译器会自动加上分号。体现出了Go的简介
- Go语言定义的变量或者import引入的包没有使用到，那么会报错，注意和Java的区别
- 最好一行只写一条语句，大括号都是成对出现的
- 一个包里面只能有一个main入口函数(不同包可以 有main函数)



## Go语言的转义字符

-    \t ：一个指标为，表示实现对齐的功能
-    \n :   表示换行
-    \\\:   表示\
-    \r :  还存在一个替换功能

```
	fmt.Println("你好\r厉害")
	输出:厉害，相当于厉害替换了前面两个你好字符
```

 

## GoLang中的开发常见问题

文件名或者路径写错了的问题（The System ....）



## Go语言项目知识回顾

1. GoRoot:指定go sdk的安装目录
2. Path:指令sdk\bin目录: 其中含有go.exe godoc.exe gofmt.exe
3. GoPath:就是golang的工作没目录，我们所有项目都在这个目录下

GoLang变量使用的几种形式:

- 指定变量类型，声明了不赋值那么就会使用默认值
- 声明时候就赋值可以省略数据类型
- 也可以先声明变量，在后面随即赋值即可

```go
//声明不赋值
	//int默认值就是0
	var i int
	fmt.Print("i=",i)
	
输出:i = 0
```

- 根据值自行判定变量类型（类型推导）

```go
//根据值自行判定变量类型（类型推导）
	var num = 11.11
	fmt.Println("num=",num)
	输出：num= 11.11
```

- 多变量和省略var赋值

```go
////省略var 的时候使用:=
	////下面的方式等价于var name string name = "tom"
	//name :="zhangsan"
	//fmt.Println("name = ",name)
	//
	////多变量赋值,变量类型一样
	//var n1,n2,n3 int
	//fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

	//多变量赋值,变量类型不一样
	var n1,name,n3 = 100,"zhangsan",200
	fmt.Println(n1,name,n3)

```

- 全局变量使用的几种形式

  ```go
  //演示全局变量
  var n1 = 100
  var n2 = 200
  var name = "jack"
  
  
  //全局变量简写
  var(
  	n3 = 400
  	n4 = 500
  	n5 = 600
  )
  ```

  `注意：int和小数默认为0，string默认为空串`，浮点数默认为float64

  

  **几种数据类型**(rune和int32类似)

  <img src="C:\Users\yuanfeng\AppData\Roaming\Typora\typora-user-images\image-20191117112957548.png" style="zoom:50%;" />

  > **有符号测试：**

  <img src="C:\Users\yuanfeng\AppData\Roaming\Typora\typora-user-images\image-20191117113943128.png" style="zoom:50%;" />

  

  > **无符号的测试：**

  <img src="C:\Users\yuanfeng\AppData\Roaming\Typora\typora-user-images\image-20191117114130500.png" alt="image-20191117114130500" style="zoom:50%;" />

  > **byte部分测试：**

  <img src="C:\Users\yuanfeng\AppData\Roaming\Typora\typora-user-images\image-20191117114600006.png" alt="image-20191117114600006" style="zoom:50%;" />

  

  查看变量的数据类型:

  > ```go
  > n1 := 12fmt.Println("n1的类型是%T",n1)
  > //查看变量的占用字节大小和数据类型
  > var n2 int64 = 32
  > fmt.Println("n2的类型是%T n2占用的字节数是 %d",n2,unsafe.Sizeof(n2))
  > 
  > 注意：float32和 float64都是有符号的,选择精度准确避免丢失最好选择float64
  > ```

  ## GoLang中“+”的使用

  ```go
  package main
  
  import "fmt"
  
  //测试字符串和数字加号的使用
  func testAdd() {
  
  	var i = 1
  	var j = 2
  	var r = i + j
  	fmt.Println("r = ",r)
  
  	//字符串
  	var str1 = "hello"
  	var str2 = "world"
  	var str3 = str1 + str2
  	fmt.Println("拼接后=",str3)
  
  }
  ```

  ```go
  package main
  
  func main() {
  	
  	testAdd()
  
  }
  
  ```

  

  ## GoLang不同文件的调用问题

  > 同一个包里面，不同文件之间，不需要 import，直接用就好。不同包的话，需要引用包，只能使用大写字母开头的方法 ，变量 等等，小写子母开头的只能包内使用。

  ***简单来说 ：大写字母开头的变量，方法暴露给其他包用的，包内的话可以随便引用\***

​        **注意：在main包中的函数之间不能互相调用，如果非要调用可以同时选中两个.go文件然后右击run运行即可，一般其他包没有这种限制;**



**Golang没有专门的字符类型，我们使用byte存储，传统字符串是由单个字符组成，Golang中字符由字节组成**

字符格式输出：

```go
package main

import "fmt"

func main() {

	var c1 byte = 'a'
	var c2 byte = 'b'
	var c3 byte = '0'

	fmt.Println("c1 = ",c1)
	fmt.Println("c2 = ",c2)
	fmt.Println("c3 = ",c3)

	//按照字符形式输出
	fmt.Print("c2=%c",c2)


}
```

> Go语言允许使用转义字符'\'来将其后的字符转换为特殊字符类型变量，例如:var c3 char = '\n'表示换行符;
>
> Go语言中的编码是使用的utf8(unicode)，其中英文字母占据一个字节，汉字占据3个字节;
>
> 其实我们在Go中直接输出字符的时候其实是输出的对应的UTF-8的编码的码值;
>
> 我们可以直接给变量赋值一个数字，要是这个时候使用%c输出是，会输出该数字对应的unicode字符;
>
> ```go
> var c4 int =2269//22269-->'国'
> fmt.Print("c4=%c\n",c4)
> ```
>
> 字符类型是可以进行运算的，相当于一个整数，其实都是因为都有各自的Unicode编码。
>
> 现在使用Go语言再也没有Java中jsp那种乱码问题了，因为现在在Go语言中只有utf-8



## Go中字符串的注意事项

字符串只要赋值了就不能被改变，这点类似于Java中的String类

```
var str string = "hello"
	str[0] = 'a' //这是错误的
	fmt.Println("str=",str)
```

字符串可以使用""和反引号(``)

反引号实现字符串原始输出:

```go
//反引号实现原始字符串的输出
	str3 := `//var str1 string = "你好世界!"
	//fmt.Println(str1)

	//var str string = "hello"
	//str[0] = 'a'
	//fmt.Println("str=",str)`

	fmt.Println(str3)
	
	输出：//var str1 string = "你好世界!"
	//fmt.Println(str1)

	//var str string = "hello"
	//str[0] = 'a'
	//fmt.Println("str=",str)
```

字符串使用"+"号拼接

注意拼接多行字符串时候，需要把"+"号放在上行

```go
str4 := "hello + hello + hello + hello" +
		"hello" + "hello" +"hello" +
		"hello"+
		"world!"

	fmt.Println("str4 = ",str4)
```

> 在Golang中默认值又被称作零值

> ```go
> 	//看看几种数据类型的默认值
> 	var a int //0
> 	var b float64 //0
> 	var isMe bool //false
> 	var name string //""
> 
> 	//%v表示按照变量值输出
> 	fmt.Print("a=%d,b=%v,isMe=%v,name=%s",a,b,isMe,name)
> ```



**基本数据类型的相互转换**(基本-->string)

不管低-->高，还是高-->低，都需要显示转换

> **Sprintf (将一个指定数据类型转换为你需要的字符串)**

```go
	//基本类型转为string
	//var num int = 99
	//var num2 float64 = 23.43
	//var b bool = true
	//var str string
	var myChar byte = 'h'

	//使用Sprintf
	////s := fmt.Sprintf("%d", num)
	//s := fmt.Sprintf("%f", num2)
	//fmt.Printf("s type %T str = %v %d",s,s,unsafe.Sizeof(s))
	//fmt.Printf("s = %T\n,s=%q",s,s)

	//字符转换为字符串
	sprintf := fmt.Sprintf("%c", myChar)
	fmt.Printf("s = %T\ns=%q",sprintf,sprintf)
```

> **使用strconv进行转换**

```go
	//var num3 int = 90
	//var num4 float64 = 23.97
	var result bool = false
	//formatInt := strconv.FormatInt(int64(num3), 10)
	//formatInt := strconv.FormatFloat(num4, 'f',10,64)
	formatBool := strconv.FormatBool(result)
	fmt.Printf("str type is %T ,str is %q",formatBool,formatBool)
```



> strconv包中有一个函数叫Itoa 
>
> ```go
> 	var num5 int64 = 21
> 	itoa := strconv.Itoa(int(num5))
> 	fmt.Printf("%T %q",itoa,itoa)
> ```
>
> **字符串转为其他基本数据类型**
>
> ```go
>     var str string = "true"
> 	var str2 string = "122222"
> 	var n1 int64
> 	var parseBool bool
> 	parseBool,_ =  strconv.ParseBool(str)
> 	fmt.Printf("%T %#v\n",parseBool,parseBool)
> 
> 	n1,_ = strconv.ParseInt(str2,10,64)
> 	fmt.Printf("%T %#v",n1,n1)
> ```
>
> ***注意***：确保string类型能够转成有效的数据，比如我们不能直接把"hello"转成int,这时候golang直接就会转成0，其他比如string转为bool随机转为false，其为默认值.其他类似。