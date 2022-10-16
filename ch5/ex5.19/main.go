// ex5.19 returns a non-zero value using panic and recover, contradicting the
// function signature.
package main

import (
	"fmt"
)

func value() (ret string) {
	defer func() {
		recover()
		ret = "lets go"
	}()
	panic("panic anyway")
}

func main() {
	fmt.Println(value())
}
