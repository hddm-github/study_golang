package main

import (
	"fmt"
	"net"
)

var level float32 = 1.0

var (
	a int
	b float32
	c string
)
var C3i int

func main() {
	//fmt.Printf("%T \n", a)
	//fmt.Printf("%T \n", b)
	//fmt.Printf("%T \n", c)
	//fmt.Printf("a=%d,b=%f,c=%s", a, b, c)
	//fmt.Printf("%T", i)
	//var level int = 1
	//level2 := 2
	//fmt.Printf("%d",level)
	//var conn net.Conn
	//var err error
	conn, err := net.Dial("tcp", "8.8.8.8")
	fmt.Println(conn, err)
	conn1, err := net.Dial("tcp", "127.0.0.1")
	fmt.Println(conn1, err)
}
