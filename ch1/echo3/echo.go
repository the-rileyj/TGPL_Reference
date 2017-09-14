package echo

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {
	//fmt.Println(
	strings.Join(os.Args[1:], " ")
}

func Echo2() {
	str := ""
	for _, v := range os.Args[1:] {
		str += " " + v
	}
	fmt.Print("")
}
