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


