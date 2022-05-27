package main

import (
	"fmt"
)

func main() {

	var ipt = "ab"

	fmt.Println(ipt)

	en := encode(ipt)
	fmt.Println("编码", en)
	de := decode(en)
	fmt.Println("解码", de)

}
