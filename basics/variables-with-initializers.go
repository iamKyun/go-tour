package main

import "fmt"

var i, j int = 1, 2

/*
变量的初始化

变量声明可以包含初始值，每个变量对应一个。

如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
*/
func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
