package main

import "fmt"

func main(){
	//fmt.Println("Hello World !")
	//fmt.Print("你好世界!")

	//go语言中指针,和c语言中的指针类似
	//指针就是操作内存地址空间
	//var a int = 10
	//fmt.Println("你好世界!")
	//fmt.Println("a = ",a)
	//var b *int = &a
	//fmt.Println("b = ",*b)
	//var a,b = 1,2
	//a,b = b,a
	//c :=1
	//d := 2
	//fmt.Println(a,b)
	//fmt.Println(c,d)


	//尝试一下c语言中的指针



	//Go语言返回多个参数
   a,b := returnManyValue(10,6)
   fmt.Println("加法结果是",a)
   fmt.Print("减法结果是",b)


}
//返回多个值
func returnManyValue(n1 int,n2 int)(int,int) {

	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
