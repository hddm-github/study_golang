package main

import (
	"fmt"
)

// Go 语言学习笔记 - 第二部分：高级特性

// 定义一个结构体
type Person struct {
	Name string
	Age  int
	City string
}

// 为 Person 结构体定义方法
func (p Person) Introduce() string {
	return fmt.Sprintf("大家好，我是%s，今年%d岁，来自%s", p.Name, p.Age, p.City)
}

// 指针接收者的方法
func (p *Person) Grow(years int) {
	p.Age += years
}

// 函数示例：计算两个数的和与差
func calculateSumAndDiff(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return // 命名返回值
}

// 可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 函数作为参数
func process(data []int, f func(int) int) []int {
	var result []int
	for _, v := range data {
		result = append(result, f(v))
	}
	return result
}

// 闭包示例
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 实现 Shape 接口的矩形结构体
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 实现 Shape 接口的圆形结构体
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func main() {
	// 1. 切片 (Slices)
	fmt.Println("=== 切片 (Slices) ===")

	// 创建切片
	var slice1 []int               // 声明一个空切片
	slice2 := []int{1, 2, 3, 4, 5} // 初始化切片

	fmt.Println("空切片:", slice1, "长度:", len(slice1), "容量:", cap(slice1))
	fmt.Println("初始化切片:", slice2, "长度:", len(slice2), "容量:", cap(slice2))

	// 使用 make 创建切片
	slice3 := make([]int, 5)     // 长度和容量都是5
	slice4 := make([]int, 3, 10) // 长度是3，容量是10

	fmt.Println("make创建切片:", slice3, "长度:", len(slice3), "容量:", cap(slice3))
	fmt.Println("make创建切片:", slice4, "长度:", len(slice4), "容量:", cap(slice4))

	// 切片操作
	fmt.Println("\n切片操作:")
	arr := [5]int{10, 20, 30, 40, 50}
	slice5 := arr[1:4] // 包含索引1到3的元素
	fmt.Println("从数组创建切片:", slice5)

	// 切片的切片
	slice6 := slice2[1:3]
	fmt.Println("切片的切片:", slice6)

	// 追加元素
	slice7 := []int{1, 2, 3}
	slice7 = append(slice7, 4, 5)
	fmt.Println("追加元素后:", slice7)

	// 追加另一个切片
	slice8 := []int{6, 7, 8}
	slice7 = append(slice7, slice8...)
	fmt.Println("追加切片后:", slice7)

	// 复制切片
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	fmt.Printf("复制切片: %v, 复制的元素个数: %d\n", dst, copied)

	// 2. 映射 (Maps)
	fmt.Println("\n=== 映射 (Maps) ===")

	// 创建映射
	var map1 map[string]int      // 声明一个映射
	map2 := make(map[string]int) // 使用 make 创建映射
	map3 := map[string]int{      // 初始化映射
		"apple":  5,
		"banana": 8,
		"orange": 7,
	}

	fmt.Println("空映射:", map1)
	fmt.Println("make创建映射:", map2)
	fmt.Println("初始化映射:", map3)

	// 添加和修改元素
	map2["apple"] = 5
	map2["banana"] = 8
	fmt.Println("添加元素后:", map2)

	map2["apple"] = 10
	fmt.Println("修改元素后:", map2)

	// 获取元素
	value, exists := map3["apple"]
	if exists {
		fmt.Println("apple的值:", value)
	}

	// 删除元素
	delete(map3, "orange")
	fmt.Println("删除元素后:", map3)

	// 遍历映射
	fmt.Println("\n遍历映射:")
	for key, value := range map3 {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	// 3. 函数
	fmt.Println("\n=== 函数 ===")

	// 基本函数调用
	sum, diff := calculateSumAndDiff(10, 5)
	fmt.Printf("和: %d, 差: %d\n", sum, diff)

	// 可变参数函数
	//total := sum(1, 2, 3, 4, 5)
	//fmt.Println("总和:", total)
	//
	//numbers := []int{10, 20, 30, 40, 50}
	//total = sum(numbers...) // 展开切片作为可变参数
	//fmt.Println("切片总和:", total)

	// 函数作为参数
	double := func(x int) int { return x * 2 }
	doubled := process([]int{1, 2, 3, 4, 5}, double)
	fmt.Println("加倍后:", doubled)

	// 匿名函数
	square := func(x int) int {
		return x * x
	}
	fmt.Println("5的平方:", square(5))

	// 立即执行的匿名函数
	result := func(x, y int) int {
		return x + y
	}(10, 20)
	fmt.Println("立即执行结果:", result)

	// 闭包
	fmt.Println("\n闭包示例:")
	nextID := counter()
	fmt.Println(nextID()) // 1
	fmt.Println(nextID()) // 2
	fmt.Println(nextID()) // 3

	// 4. 结构体
	fmt.Println("\n=== 结构体 ===")

	// 创建结构体
	p1 := Person{
		Name: "李四",
		Age:  30,
		City: "上海",
	}

	// 简短声明
	p2 := Person{"王五", 25, "广州"}

	fmt.Println("结构体 p1:", p1)
	fmt.Println("结构体 p2:", p2)

	// 访问字段
	fmt.Println("p1的姓名:", p1.Name)
	p1.Age = 31
	fmt.Println("修改后的年龄:", p1.Age)

	// 调用方法
	fmt.Println(p1.Introduce())

	// 指针接收者方法
	p1.Grow(5)
	fmt.Println("5年后的年龄:", p1.Age)

	// 匿名结构体
	point := struct {
		X, Y int
	}{10, 20}
	fmt.Println("匿名结构体:", point)

	// 5. 接口
	fmt.Println("\n=== 接口 ===")

	// 创建实现接口的实例
	rect := Rectangle{Width: 5, Height: 4}
	circ := Circle{Radius: 3}

	// 使用接口变量
	var shape Shape

	shape = rect
	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", shape.Area(), shape.Perimeter())

	shape = circ
	fmt.Printf("圆形面积: %.2f, 周长: %.2f\n", shape.Area(), shape.Perimeter())

	// 接口切片
	shapes := []Shape{rect, circ}
	for _, s := range shapes {
		fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
	}

	// 空接口
	var anything interface{}
	anything = 42
	fmt.Println("空接口存储整数:", anything)
	anything = "Hello"
	fmt.Println("空接口存储字符串:", anything)
	anything = Person{"赵六", 35, "深圳"}
	fmt.Println("空接口存储结构体:", anything)

	// 类型断言
	value1 := anything.(Person)
	fmt.Println("类型断言后:", value1.Name)

	// 安全的类型断言
	if value2, ok := anything.(Person); ok {
		fmt.Println("安全类型断言:", value2.Name)
	}

	// 类型选择
	whatType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("整数: %d\n", v)
		case string:
			fmt.Printf("字符串: %s\n", v)
		case Person:
			fmt.Printf("Person: %s\n", v.Name)
		default:
			fmt.Printf("未知类型\n")
		}
	}

	whatType(42)
	whatType("Hello")
	whatType(Person{"张三", 20, "北京"})
}
