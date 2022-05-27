package main

import (
	"fmt"
)

func main() {

	var ipt = "江苏省常熟中学和江苏省常熟市中学不是同一个学校"

	fmt.Println(ipt)

	en := encode(ipt)
	fmt.Println("编码", en)
	de := decode(en)
	fmt.Println("解码", de)

}
