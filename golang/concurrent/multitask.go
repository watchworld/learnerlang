package concurrent

import (
	"fmt"
	"time"
)

//并发程序的调试
func f() {
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("f()")
}

func f1() {
	fmt.Println("f1()")
}

func TestGoMultiTask() {
	go f()
	go f1()
	time.Sleep(time.Duration(10) * time.Second)
	fmt.Println("TestGoMultiTask()")
}

func TestChanel() {
	ch := make(chan int)
	//匿名函数
	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		ch <- 10
	}()
	f := <-ch //无缓存的channel在无法输入时候会阻塞
	fmt.Println(f)
	close(ch) //关闭管道
}

func TestCacheChanel() {
	ch := make(chan int, 3)
	//匿名函数
	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		ch <- 10
		ch <- 100
	}()
	f := <-ch //无缓存的channel在无法输入时候会阻塞
	fmt.Println(f)
	f = <-ch
	fmt.Println(f)
	//close(ch) //关闭管道
}

func TestSelectChanel(y int) {
	ch := make(chan int, 1)
	select {
	case <-ch:
	case x := <-ch:
		fmt.Printf("number:%d", x)
	default:
		fmt.Println("default")
	}
}
