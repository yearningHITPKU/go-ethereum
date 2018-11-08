package p2p

import(
	"fmt"
)

var enable = false

func debugf(format string, args ...interface{}){
	if enable {
		fmt.Printf(format, args...)
	}
}

func debugln(args ...interface{}){
	if enable {
		fmt.Println(args...)
	}
}
