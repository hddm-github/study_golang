package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type Point struct {
	X, Y int
}

func (p *Point) setX(newX int) (suc bool) {
	p.X = newX
	return true
}

type Vec2 struct {
	X, Y float32
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

func main() {

	//p := Point{}
	//p.Y = 3
	//fmt.Println(p)
	//
	//p2 := new(Point)
	//p2.X = 4
	//fmt.Println(p2)
	//
	//p3 := Point{X: 1, Y: 2}
	//
	//p3.setX(3)
	//fmt.Println(p3)

	//file, err := os.ReadFile("main.go")
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(file))

	//reader := strings.NewReader("hello world")
	//p := make([]byte, 4)
	//for {
	//	n, err := reader.Read(p)
	//	if err != nil {
	//		if err == io.EOF {
	//			log.Printf("read %d bytes, err: %v", n, err)
	//			break
	//		}
	//
	//		log.Printf("read %d bytes, err: %v", n, err)
	//		return
	//	}
	//	log.Printf("[读取道德字节数为%d][内容:%v]", n, string(p[:n]))
	//}

	//open, err := os.Open("main.go")
	//if err != nil {
	//	log.Println("open file error")
	//	return
	//}
	//defer func(open *os.File) {
	//	err := open.Close()
	//	if err != nil {
	//		log.Println("close file error")
	//	}
	//}(open)
	//var buf [128]byte
	//var content []byte
	//for {
	//	n, err := open.Read(buf[:])
	//	if err == io.EOF {
	//		log.Println("读取结束")
	//		break
	//	}
	//	if err != nil {
	//		log.Println("read file error", err)
	//		return
	//	}
	//	content = append(content, buf[:n]...)
	//}
	//log.Println(string(content))
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		open, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "open file error:%s\n%v\n", flag.Arg(i), err)
			return
		}
		cat(bufio.NewReader(open))

	}
}
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}
