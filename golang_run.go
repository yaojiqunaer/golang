/* @author: zhangxiaodong
 * @desc: golang_run.go
 * @create: 2022/2/26 23:59
 */
package main

import (
	"fmt"
	//匿名导包
	"golang/basic"
	_ "golang/basic"
)

func main() {
	//变量声明测试
	basic.Var()
	//init方法测试
	basic.One()
	basic.VoidResult()
	_ = basic.IntResult()
	var name string = basic.ParamAndResult("zhangshan")
	var _, two = basic.ParamsAndResults("zhangshan", 20)
	fmt.Println(two)
	var one, three = basic.ParamsAndResults("zhangshan", 20)
	fmt.Println(one, three)
	one, three = basic.ParamsAndResults("zhangshan2", 10)
	fmt.Println(one, three)
	length := basic.ParamsTypesAndResults("zhangshan", "男", "34")
	fmt.Println(length)
	fmt.Println(name)
	basic.Defer()
}
