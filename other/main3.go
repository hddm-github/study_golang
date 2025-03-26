package main

import (
	"flag"
	"fmt"
)

func main() {
	//var a = 100
	//var b = 200
	//var c int
	//c = b
	//b = a
	//a = c
	//fmt.Printf("a=%d,b=%d", a, b)
	//var a = 100
	//var b = 200
	//b, a = a, b
	//
	//fmt.Printf("a=%d,b=%d", a, b)

	//dial, _ := net.Dial("tcp", "8.8.8.8")
	//fmt.Printf("%T", dial)
	//C3i = 3

	//var a = 100
	//var b = 200
	//c := sum(a, b)
	//fmt.Println(c)

	//var ch rune = '\u0031'
	//var ch1 int64 = '\U00000011'
	//fmt.Println(ch)
	//fmt.Println(ch1)
	//fmt.Println(unicode.IsDigit(ch))
	//fmt.Println(unicode.IsLetter(ch))

	//var str = "hello你好啊鲍勃啊\n 溶溶"
	//var str2 = `hello你好啊鲍勃啊\\n
	//			溶溶`
	//fmt.Print(str2)

	//var myStr01 string = "hello,李强"
	//fmt.Printf("len=%d \n", len(myStr01))
	//fmt.Println(myStr01[6])
	//
	//lenX := utf8.RuneCountInString(myStr01)
	//fmt.Printf("%d", lenX)

	//var stringBuilder bytes.Buffer
	//stringBuilder.WriteString("hello")
	//stringBuilder.WriteString("world")
	//fmt.Println(stringBuilder.String())
	//s1 := "hello"
	//s2 := "world,李强"
	//u := string([]rune(s2)[7])
	//fmt.Println(s1 + s2)
	//fmt.Println(s1[0])
	//fmt.Println(u)

	//s2 := "你好"
	//for _, v := range s2 {
	//	result := fmt.Sprintf("%c \n", v)
	//	fmt.Printf(result)
	//}
	//pi := math.Pi
	//fmt.Println(pi)
	//v := 3
	//ptr := &v
	//*ptr = 4
	//
	//fmt.Println(*ptr + 1)
	//fmt.Println(*ptr)

	var mode = flag.String("mode", "", "this is test mode.")
	//point := new(string)
	//fmt.Println(*point)

	flag.Parse()
	fmt.Printf("mode=%s", *mode)

}

func sum(i, i2 int) int {
	return i + i2
}
