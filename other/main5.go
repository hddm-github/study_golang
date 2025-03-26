package main

import (
	"errors"
	"fmt"
)

func main() {

	//fmt.Println(m(1, 2))
	//test(fn)
	//i, d := sum2(2, 4)
	//fmt.Println(i)
	//fmt.Println(d)

	//fu := func(a, b int) float64 {
	//	return math.Sqrt(float64(a))
	//}
	//fmt.Println(fu(2, 3))
	//visit([]int{1, 2, 3, 4, 5}, func(i int) {
	//	fmt.Println(i)
	//})

	//getPlayer := playerGen()
	//player, i := getPlayer("哈哈")
	//fmt.Println(player, i)

	//var watever = [5]int{1, 2, 3, 4, 5}
	//
	//var v2 int
	//for _, v := range watever {
	//
	//	defer func() { fmt.Println(v) }()
	//}

	//start := time.Now()
	//log.Printf("start:%v", start)
	//
	//defer log.Printf("时间差:%v", time.Since(start))
	//defer func() {
	//	log.Printf("时间差:%v", time.Since(start))
	//}()
	//
	//time.Sleep(3 * time.Second)
	//
	//log.Printf("函数结束")

	//test22()
	//defer func() {
	//	fmt.Println(recover())
	//}()
	//switch z, err := div(10, 0); err {
	//case nil:
	//	fmt.Println(z)
	//case ErrDivByZero:
	//	panic(err)
	//}
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

var ErrDivByZero = errors.New("division by zero")

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

func test22() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("panic")
}
func playerGen() func(string) (string, int) {
	hp := 150

	return func(name string) (string, int) {
		return name, hp
	}
}
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
func m(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func test(fn func() int) {
	fmt.Println(fn())
}
func fn() int {
	return 2
}

func sum2(a, b int) (c int, d int) {
	c++
	return
}
