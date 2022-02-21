package var

import {
	"fmt"
}
/*
  # 变量声明
*/
func main() {

	//申明变量 不赋值默认为0
	var a int
	fmt.Println("a = " + a)

	//申明变量  b赋值
	var b int = 100
	fmt.Println("b = " + b)

	//声明变量 自动推断类型
	var c = 200
	fmt.Println("c = " + c)

	//省略var
	e := 900
	fmt.Println("e = " + e)

	//多个变量
	var name, age = "zhangsan", 20
	fmt.Println("name: " + name, "age: " + age)

	//多行多变量
	var (
		id int = 1
		hight string = "2.334m"
		isGirl bool = true
	)
	fmt.Println("id: " + id, "isGirl: " + isGirl)
	

}