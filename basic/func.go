package basic

import "fmt"

func init() {
	fmt.Println("init starting ...")
}

func VoidResult() {
	fmt.Println("无参无返回值...")
	return
}

func IntResult() int {
	fmt.Println("无参带返回值...")
	return 1
}

func ParamAndResult(name string) string {
	fmt.Println("传入参数: "+name, "返回值："+name)
	return ""
}

func ParamsAndResults(name string, age int) (string, int) {
	fmt.Println("传入参数: "+name, "返回值："+name)
	return name, age
}
func ParamsTypesAndResults(name, sex, age string) (length int) {
	return len(name) + len(sex) + len(age)
}

func Defer() {
	defer fmt.Println("defer executing 4...")
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
}
