package main

import (
	"auto-chef/model"
	"fmt"
)

func main() {

	p := model.TempData{TmpA: 1, TmpB: 2, TmpC: "yolo"}

	fmt.Println("hello world")
	fmt.Println(p)
}
