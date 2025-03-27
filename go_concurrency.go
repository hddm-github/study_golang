package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Go 语言学习笔记 - 第三部分：并发与高级特性

// 自定义错误
var (
	ErrDivByZero = errors.New("除数不能为零")
	ErrNegative  = errors.New("不能使用负数")
)

// 结构化的错误类型
type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("查询 %q 失败: %v", e.Query, e.Err)
}

// 安全的除法函数，返回错误
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

// 演示 defer 的函数
func deferDemo() {
	fmt.Println("=== defer 示例 ===")

	// defer 语句会在函数返回前执行
	defer fmt.Println("1. 函数结束时执行")

	// 多个 defer 按照 LIFO（后进先出）顺序执行
	defer fmt.Println("2. 这条消息会在 1 之前显示")
	defer fmt.Println("3. 这条消息会在 2 之前显示")

	fmt.Println("函数开始执行...")
}

// 演示 defer 用于资源清理
func deferCleanup() {
	fmt.Println("\n=== defer 用于资源清理 ===")

	// 模拟打开文件
	fmt.Println("打开文件...")

	// 确保文件关闭，即使后面的代码发生错误
	defer func() {
		fmt.Println("关闭文件...")
	}()

	// 处理文件
	fmt.Println("处理文件内容...")
}

// 演示 defer 与匿名函数
func deferWithValues() {
	fmt.Println("\n=== defer 与变量值 ===")

	a := 1
	defer fmt.Println("defer 中的 a =", a) // 会打印 1

	// defer 中的匿名函数可以捕获后续修改的变量值
	defer func() {
		fmt.Println("defer 匿名函数中的 a =", a) // 会打印 2
	}()

	a = 2
	fmt.Println("函数中的 a =", a) // 会打印 2
}

// 演示 defer 用于性能测量
func deferForTiming() {
	fmt.Println("\n=== defer 用于性能测量 ===")

	start := time.Now()

	defer func() {
		fmt.Printf("函数执行时间: %v\n", time.Since(start))
	}()

	// 模拟耗时操作
	time.Sleep(100 * time.Millisecond)
	fmt.Println("耗时操作完成")
}

// 演示 panic 和 recover
func panicAndRecover() {
	fmt.Println("\n=== panic 和 recover ===")

	// 使用 defer 和 recover 捕获 panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("已恢复的 panic:", r)
		}
	}()

	fmt.Println("函数开始执行...")

	// 引发 panic
	panic("发生了严重错误！")

	// 这行代码不会执行
	fmt.Println("这行不会执行")
}

// 演示实际应用中的 panic 和 recover
func divideWithPanic(a, b int) int {
	// 如果除数为零，引发 panic
	if b == 0 {
		panic("除数不能为零")
	}
	return a / b
}

func safeDivideWithRecover(a, b int) (result int) {
	// 设置 defer 来捕获可能的 panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic:", r)
			result = 0 // 设置默认返回值
		}
	}()

	// 调用可能会 panic 的函数
	return divideWithPanic(a, b)
}

// 简单的 goroutine 示例
func goRoutineDemo() {
	fmt.Println("\n=== goroutine 示例 ===")

	// 启动一个 goroutine
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine:", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// 主 goroutine
	for i := 0; i < 3; i++ {
		fmt.Println("主函数:", i)
		time.Sleep(20 * time.Millisecond)
	}

	// 等待 goroutine 完成
	time.Sleep(100 * time.Millisecond)
}

// 使用 WaitGroup 等待 goroutine 完成
func waitGroupDemo() {
	fmt.Println("\n=== WaitGroup 示例 ===")

	var wg sync.WaitGroup

	// 启动 3 个 goroutine
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数器

		go func(id int) {
			defer wg.Done() // 完成时减少计数器

			fmt.Printf("goroutine %d 开始\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("goroutine %d 结束\n", id)
		}(i)
	}

	fmt.Println("等待所有 goroutine 完成...")
	wg.Wait() // 阻塞直到计数器为 0
	fmt.Println("所有 goroutine 已完成")
}

// 使用 channel 进行 goroutine 间通信
func channelDemo() {
	fmt.Println("\n=== channel 示例 ===")

	// 创建一个无缓冲的 channel
	messages := make(chan string)

	// 发送者 goroutine
	go func() {
		fmt.Println("发送者: 发送消息")
		messages <- "你好，channel!" // 发送消息
		fmt.Println("发送者: 消息已发送")
	}()

	// 接收消息
	fmt.Println("接收者: 等待消息")
	msg := <-messages // 接收消息
	fmt.Println("接收者: 收到消息:", msg)
}

// 使用带缓冲的 channel
func bufferedChannelDemo() {
	fmt.Println("\n=== 带缓冲的 channel 示例 ===")

	// 创建一个容量为 2 的缓冲 channel
	messages := make(chan string, 2)

	// 发送消息
	messages <- "消息1"
	messages <- "消息2"

	// 接收消息
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// 使用 channel 实现工作池
func workerPoolDemo() {
	fmt.Println("\n=== 工作池示例 ===")

	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动 worker
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for job := range jobs {
				fmt.Printf("工作者 %d 处理任务 %d\n", id, job)
				time.Sleep(100 * time.Millisecond) // 模拟处理时间
				results <- job * 2                 // 发送结果
			}
		}(w)
	}

	// 发送任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // 关闭 jobs channel

	// 等待所有 worker 完成
	go func() {
		wg.Wait()
		close(results) // 关闭 results channel
	}()

	// 收集结果
	for result := range results {
		fmt.Println("结果:", result)
	}
}

// 使用 select 多路复用 channel
func selectDemo() {
	fmt.Println("\n=== select 示例 ===")

	c1 := make(chan string)
	c2 := make(chan string)

	// 发送数据到 c1
	go func() {
		time.Sleep(50 * time.Millisecond)
		c1 <- "来自 c1 的消息"
	}()

	// 发送数据到 c2
	go func() {
		time.Sleep(100 * time.Millisecond)
		c2 <- "来自 c2 的消息"
	}()

	// 使用 select 接收数据
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("接收:", msg1)
		case msg2 := <-c2:
			fmt.Println("接收:", msg2)
		}
	}
}

// 使用 select 实现超时
func selectWithTimeout() {
	fmt.Println("\n=== select 超时示例 ===")

	c := make(chan string)

	// 尝试发送数据
	go func() {
		time.Sleep(200 * time.Millisecond)
		c <- "操作完成"
	}()

	// 使用 select 实现超时
	select {
	case result := <-c:
		fmt.Println("接收:", result)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("操作超时")
	}
}

// 使用 sync.Mutex 实现互斥锁
func mutexDemo() {
	fmt.Println("\n=== 互斥锁示例 ===")

	var counter = 0
	var mutex sync.Mutex
	var wg sync.WaitGroup

	// 启动 10 个 goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				mutex.Lock()   // 获取锁
				counter++      // 临界区
				mutex.Unlock() // 释放锁
			}
		}()
	}

	wg.Wait()
	fmt.Println("计数器最终值:", counter)
}

// 使用 sync.RWMutex 实现读写锁
func rwMutexDemo() {
	fmt.Println("\n=== 读写锁示例 ===")

	var data = make(map[string]string)
	var rwMutex sync.RWMutex
	var wg sync.WaitGroup

	// 写入 goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			rwMutex.Lock() // 获取写锁
			key := fmt.Sprintf("key%d", id)
			data[key] = fmt.Sprintf("value%d", id)
			fmt.Printf("写入: %s = %s\n", key, data[key])
			time.Sleep(50 * time.Millisecond) // 模拟写入时间
			rwMutex.Unlock()                  // 释放写锁
		}(i)
	}

	// 读取 goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			time.Sleep(10 * time.Millisecond) // 给写入一些时间

			rwMutex.RLock() // 获取读锁
			for k, v := range data {
				fmt.Printf("读取 %d: %s = %s\n", id, k, v)
			}
			time.Sleep(10 * time.Millisecond) // 模拟读取时间
			rwMutex.RUnlock()                 // 释放读锁
		}(i)
	}

	wg.Wait()
}

// 使用 sync.Once 确保函数只执行一次
func onceDemo() {
	fmt.Println("\n=== sync.Once 示例 ===")

	var once sync.Once
	var wg sync.WaitGroup

	// 初始化函数
	initialize := func() {
		fmt.Println("初始化操作 - 只会执行一次")
	}

	// 启动多个 goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("goroutine %d 尝试执行初始化\n", id)
			once.Do(initialize) // 只会执行一次
			fmt.Printf("goroutine %d 继续执行\n", id)
		}(i)
	}

	wg.Wait()
}

// 使用 context 控制 goroutine
func contextDemo() {
	fmt.Println("\n=== context 示例 ===")

	// 创建一个带取消功能的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 启动 worker goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker: 收到取消信号，退出")
				return
			default:
				fmt.Println("worker: 正在工作...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 让 worker 工作一段时间
	time.Sleep(300 * time.Millisecond)

	// 发送取消信号
	fmt.Println("main: 发送取消信号")
	cancel()

	// 给 worker 一些时间来响应取消
	time.Sleep(100 * time.Millisecond)
}

func main() {
	// 1. defer 示例
	deferDemo()
	deferCleanup()
	deferWithValues()
	deferForTiming()

	// 2. panic 和 recover 示例
	panicAndRecover()

	fmt.Println("\n=== 实际应用中的 panic 和 recover ===")
	fmt.Println("安全除法结果:", safeDivideWithRecover(10, 2))
	fmt.Println("安全除法结果:", safeDivideWithRecover(10, 0))

	// 3. 错误处理示例
	fmt.Println("\n=== 错误处理 ===")

	// 基本错误处理
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}

	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}

	// 错误类型检查
	if err == ErrDivByZero {
		fmt.Println("检测到除零错误")
	}

	// 自定义错误类型
	queryErr := &QueryError{
		Query: "SELECT * FROM users",
		Err:   errors.New("数据库连接失败"),
	}
	fmt.Println("查询错误:", queryErr)

	// 4. goroutine 和 channel 示例
	goRoutineDemo()
	waitGroupDemo()
	channelDemo()
	bufferedChannelDemo()
	workerPoolDemo()
	selectDemo()
	selectWithTimeout()

	// 5. 同步原语示例
	mutexDemo()
	rwMutexDemo()
	onceDemo()

	// 6. context 示例
	contextDemo()
}
