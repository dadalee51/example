package main

import (
	"example/greetings"
	"fmt"
)

func main() {
	fmt.Println("hello from runmain")
	var retmsg = greetings.Hello("oh this is a module call")
	fmt.Println(retmsg)
}
