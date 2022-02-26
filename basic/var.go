package basic

import "fmt"

const (
	name         = "zhang"
	_    float64 = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func Var() {

	//申明变量 不赋值默认为0
	var a int
	fmt.Println("a = ", a)

	//申明变量  b赋值
	var b int = 100
	fmt.Println("b = ", b)

	//声明变量 自动推断类型
	var c = 200
	fmt.Println("c = ", c)

	//省略var
	e := 900
	fmt.Println("e = ", e)

	//多个变量
	var name, age = "zhangsan", 20
	fmt.Println("name: "+name, "age: ", age)

	//多行多变量
	var (
		id     int    = 1
		hight  string = "2.334m"
		isGirl bool   = true
	)
	fmt.Println("id: ", id, "hight: ", hight, "isGirl: ", isGirl)

	//全局变量
	fmt.Println(id)
	fmt.Println(name)
	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)
	println(TiB)
	println(PiB)
	println(EiB)
	println(ZiB)
	println(YiB)

}
