package main

import (
	"fmt"
	"strconv"
)

func main() {
	//e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//e.Logger.Fatal(e.Start(":1323"))

	//for i := 0; i < 2; i++ {
	//	go func(s string) {
	//		fmt.Println("go:Hello, World!", s)
	//	}("world")
	//}
	//
	//for i := 0; i < 2; i++ {
	//	runtime.Gosched()
	//	fmt.Println("main:Hello, World!")
	//}
	//
	//runtime.GOMAXPROCS(12)

	//ch := make(chan int)
	//go send(ch)
	//go recv(ch)
	//fmt.Println("发送成功")
	//time.Sleep(1 * time.Second)

	//c := make(chan int)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		c <- i
	//	}
	//	close(c)
	//}()
	//for {
	//	if data, ok := <-c; ok {
	//		fmt.Println(data)
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println("main结束")
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//// 开启goroutine将0~100的数发送到ch1中
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		ch1 <- i
	//	}
	//	close(ch1)
	//}()
	//// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	//go func() {
	//	for {
	//		i, ok := <-ch1 // 通道关闭后再取值ok=false
	//		if !ok {
	//			break
	//		}
	//		ch2 <- i * i
	//	}
	//	close(ch2)
	//}()
	//// 在主goroutine中从ch2中接收值打印
	//for i := range ch2 { // 通道关闭后会退出for range循环
	//	fmt.Println(i)
	//}

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	go counter1(ch1)
	go counter2(ch2)

	select {
	case i := <-ch1:
		fmt.Println(i)
	case i := <-ch2:
		fmt.Println(i)
	}

}
func counter1(out chan<- string) {
	for i := 0; i < 10; i++ {
		out <- "count11111" + strconv.Itoa(i)
	}
	close(out)
}
func counter2(out chan<- string) {
	for i := 0; i < 10; i++ {
		out <- "count2222" + strconv.Itoa(i)
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
