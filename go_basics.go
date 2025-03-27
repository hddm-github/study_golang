package main

import (
	"fmt"
	"time"
)

// Go 语言学习笔记 - 第一部分：基础知识

func main() {
	// 1. 变量声明
	fmt.Println("=== 变量声明 ===")

	// 完整声明
	var name string = "张三"
	var age int = 25

	// 类型推断
	var city = "北京"

	// 简短声明（只能在函数内部使用）
	salary := 8000.50

	fmt.Println("姓名:", name)
	fmt.Println("年龄:", age)
	fmt.Println("城市:", city)
	fmt.Println("薪资:", salary)

	// 多变量声明
	var a, b, c int = 1, 2, 3
	x, y, z := "Go", 100, true
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)

	// 2. 常量
	fmt.Println("\n=== 常量 ===")
	const PI = 3.14159
	const (
		StatusOK       = 200
		StatusNotFound = 404
	)

	fmt.Println("PI:", PI)
	fmt.Println("状态码:", StatusOK, StatusNotFound)

	// iota 常量生成器
	const (
		Monday    = iota + 1 // 1
		Tuesday              // 2
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
	)
	fmt.Println("星期几:", Monday, Tuesday, Wednesday, Thursday, Friday)

	// 3. 基本数据类型
	fmt.Println("\n=== 基本数据类型 ===")

	// 整数类型
	var intNum int = 42
	var int8Num int8 = 127 // -128 到 127
	var uintNum uint = 100 // 无符号整数

	// 浮点类型
	var float32Num float32 = 3.14
	var float64Num float64 = 3.141592653589793

	// 布尔类型
	var isTrue = true
	var isFalse = false

	// 字符串类型
	var message = "你好，Go语言！"

	fmt.Printf("整数: %d, 类型: %T\n", intNum, intNum)
	fmt.Printf("int8: %d, 类型: %T\n", int8Num, int8Num)
	fmt.Printf("无符号整数: %d, 类型: %T\n", uintNum, uintNum)
	fmt.Printf("float32: %f, 类型: %T\n", float32Num, float32Num)
	fmt.Printf("float64: %f, 类型: %T\n", float64Num, float64Num)
	fmt.Printf("布尔值: %t 和 %t, 类型: %T\n", isTrue, isFalse, isTrue)
	fmt.Printf("字符串: %s, 类型: %T\n", message, message)

	// 4. 类型转换
	fmt.Println("\n=== 类型转换 ===")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("i = %d, f = %f, u = %d\n", i, f, u)

	// 字符串转换
	var num int = 65
	var str string = string(num) // 将ASCII码转为字符
	fmt.Printf("ASCII 65 对应的字符: %s\n", str)

	// 5. 控制结构
	fmt.Println("\n=== 控制结构 ===")

	// if-else
	score := 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// if 初始化语句
	if num := 10; num > 5 {
		fmt.Println("num 大于 5")
	}

	// for 循环
	fmt.Println("\n基本for循环:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 类似 while 的 for 循环
	fmt.Println("\n类似while的for循环:")
	j := 0
	for j < 5 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()

	// 无限循环
	fmt.Println("\n带break的循环:")
	k := 0
	for {
		fmt.Printf("%d ", k)
		k++
		if k >= 5 {
			break
		}
	}
	fmt.Println()

	// for-range 循环
	fmt.Println("\nfor-range循环:")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("索引:%d 值:%d\n", index, value)
	}

	// switch 语句
	fmt.Println("\nswitch语句:")
	day := "星期三"
	switch day {
	case "星期一":
		fmt.Println("开始新的一周")
	case "星期三":
		fmt.Println("一周过半")
	case "星期五":
		fmt.Println("周末快到了")
	default:
		fmt.Println("普通的一天")
	}

	// 不带表达式的 switch
	fmt.Println("\n不带表达式的switch:")
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("上午好！")
	case hour < 18:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}

	// 6. 数组
	fmt.Println("\n=== 数组 ===")
	var arr1 [5]int // 声明长度为5的整型数组
	arr1[0] = 100
	arr1[1] = 200
	fmt.Println("arr1:", arr1)

	// 数组初始化
	arr2 := [3]string{"Go", "Java", "Python"}
	fmt.Println("arr2:", arr2)

	// 自动计算长度
	arr3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println("arr3:", arr3, "长度:", len(arr3))

	// 多维数组
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("二维数组:", matrix)
}
