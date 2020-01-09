package learn3

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//简单语法测试
func Test4() {
	const (
		USD int = iota // 美元
		EUR            // 欧元
		GBP            // 英镑
		RMB            // 人民币
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
	fmt.Printf("%T", symbol)
}

type Wheel struct {
}

//defer特性可以在代码阅读上更清晰
//文件的打卡和关闭可以在相邻地方执行
func DeferOp() {
	counts := make(map[string]int)
	f, err := os.Open("test1.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, "dup2:%d")
		return
	}
	defer f.Close() //关闭文件
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line) //格式化打印
		}
	}
}

//'对象'方法声明
//struct定义的属性是小写开头的，不是public的，这样是不能跨包调用的！
type Point struct{ X, Y float64 }

func Distance(p1, p2 Point) float64 {
	return math.Hypot(p2.X-p1.X, p2.Y-p1.Y)
}

//对象声明 私有、公共
func (p Point) Distance(p2 Point) float64 {
	return math.Hypot(p2.X-p.X, p2.Y-p.Y)
}
