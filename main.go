package main

import "fmt"

//import "github.com/astaxie/beego"

func main() {
	//	beego.run()
	var arr = [3]int{1, 2, 3}
	arr2 := arr
	arr[0] = 10

	fmt.Print(arr, arr2)
}
