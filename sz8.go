package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	depth     = 3
	charPools = []string{"常", "熟", "市", "中", "学", "我", "爱", "你"}
	indexMap  = map[string]int{
		"常": 0,
		"熟": 1,
		"市": 2,
		"中": 3,
		"学": 4,
		"我": 5,
		"爱": 6,
		"你": 7,
	}
)

func encode(s string) string {

	var iptBytes = []rune(s)
	var iptBin string
	for _, v := range iptBytes {
		bin := fmt.Sprintf("%016b", v)
		iptBin += bin
	}
	fmt.Println(iptBin)
	var spiltBin = strings.Split(iptBin, "")
	var extraChar = depth - len(spiltBin)%depth
	for i := 0; i < extraChar; i++ {
		spiltBin = append(spiltBin, "0")
	}

	var opt string
	for i := 0; i < len(spiltBin); i += depth {
		var tmp = spiltBin[i : i+depth]
		var tmpStr = strings.Join(tmp, "")
		var index, _ = strconv.ParseUint(tmpStr, 2, 64)
		opt += charPools[index]
	}

	for i := 0; i < extraChar; i++ {
		opt += "="
	}

	return opt

}

func decode(s string) string {
	var extraChar = strings.Count(s, "=")
	var spiltChar = strings.Split(s, "")
	var realChar = spiltChar[:len(spiltChar)-extraChar]
	var bin string
	for _, v := range realChar {
		bin += fmt.Sprintf("%03b", indexMap[v])
	}
	var spiltBin = strings.Split(bin, "")
	var realBin = spiltBin[:len(spiltBin)-extraChar]
	var iptBytes []rune
	for i := 0; i < len(realBin); i += 16 {
		var tmp = realBin[i : i+16]
		var tmpStr = strings.Join(tmp, "")
		var index, _ = strconv.ParseUint(tmpStr, 2, 64)
		iptBytes = append(iptBytes, rune(index))
	}
	return string(iptBytes)
}
