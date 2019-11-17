package main

import (
	"fmt"
	"strconv"
)

func main() {

	//基本类型转为string
	//var num int = 99
	//var num2 float64 = 23.43
	//var b bool = true
	//var str string
	//var myChar byte = 'h'

	//使用Sprintf
	////s := fmt.Sprintf("%d", num)
	//s := fmt.Sprintf("%f", num2)
	//fmt.Printf("s type %T str = %v %d",s,s,unsafe.Sizeof(s))
	//fmt.Printf("s = %T\n,s=%q",s,s)

	//字符转换为字符串
	//sprintf := fmt.Sprintf("%c", myChar)
	//fmt.Printf("s = %T\ns=%q",sprintf,sprintf)

	//var num3 int = 90
	//var num4 float64 = 23.97
	//var result bool = false
	////formatInt := strconv.FormatInt(int64(num3), 10)
	////formatInt := strconv.FormatFloat(num4, 'f',10,64)
	//formatBool := strconv.FormatBool(result)
	//fmt.Printf("str type is %T ,str is %q",formatBool,formatBool)

	//var num5 int64 = 21
	//itoa := strconv.Itoa(int(num5))
	//fmt.Printf("%T %q",itoa,itoa)

	var str string = "true"
	var str2 string = "122222"
	var n1 int64
	var parseBool bool
	parseBool, _ = strconv.ParseBool(str)
	fmt.Printf("%T %#v\n", parseBool, parseBool)

	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("%T %#v", n1, n1)
}
