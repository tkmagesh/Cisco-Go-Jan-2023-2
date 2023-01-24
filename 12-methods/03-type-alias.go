package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("Excepteur quis exercitation officia velit ad tempor adipisicing adipisicing consequat laboris deserunt. Ex ex labore non anim nulla. Non nulla dolore cillum velit ea sit voluptate nulla eu. Enim tempor velit ullamco nulla incididunt. Ut duis id deserunt mollit velit consequat.")
	fmt.Println(str.Length())
}
