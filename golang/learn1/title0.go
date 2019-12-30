package learn1

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sync"
)

//获取命令输入行信息
func TestRunCmd() {
	var s, sep string

	if len(os.Args) == 0 {
		fmt.Println("参加输入为空")
	}
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " " //空格
	}
	fmt.Println(s)
}

//查找重复行
func FindSameLine() {
	counts := make(map[string]int)
	f, err := os.Open("test1.txt")
	input := bufio.NewScanner(f)
	if err != nil {
		fmt.Fprint(os.Stderr, "dup2:%d")
	}
	for input.Scan() {
		counts[input.Text()]++
	}
	f.Close() //关闭文件

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line) //格式化打印
		}
	}
}

const (
	whiteIndex = 0
	blackIndex = 1
)

var palette = []color.Color{color.White, color.Black}

//gif生成
func GifBuild() {
	f, err := os.Create("test.gif")
	if err != nil {
		fmt.Println(err)
		return
	}
	lissajous(f)
	f.Close() //关闭文件
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors

}

//http请求
func HttpRequest() {
	resp, err := http.Get("https://www.leapmotor.com/news/news-detail.html?id=320")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", resp.Status)
	resp.Body.Close()
	//fmt.Printf("%s\n",b)
}

//简单的http服务器
var mu sync.Mutex
var count int

func TestWeb() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("127.0.0.1:11111", nil))
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count %d\n", count)
	mu.Unlock()
}

func handle(writer http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(writer,"URL.Path = %q \n",request.URL.Path)
	mu.Lock()
	count++
	mu.Unlock()
}
