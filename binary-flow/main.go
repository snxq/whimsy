package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	speed    = flag.Int("speed", 100, "流动速度，单位为毫秒")
	color    = flag.String("color", "green", "打印字体颜色, green/red/yellow")
	interval = flag.Int("interval", 1, "字体之间间隔的空格数量, 大小为0到屏幕宽度一半")
)

func main() {
	flag.Parse()
	width, height, err := getTerminalSize()
	if err != nil {
		log.Fatal(err)
	}
	if *interval < 0 || *interval > width/2 {
		log.Fatal("interval 参数不合法。大小应为0到屏幕宽度的一半")
	}

	c, err := GetColor(*color)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan string)
	go print(ch, *c)

	var out string
	for j := 0; j < height; j++ {
		row := flipCoins(width)
		out += row
		ch <- out
	}

	for i := 0; ; i++ {
		clearTerminal()
		out = out[width:] + flipCoins(width)
		ch <- out
		time.Sleep(time.Millisecond * time.Duration(*speed))
	}
}

// flipCoins 抛硬币，随机生成01字符串
func flipCoins(count int) (result string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for ; len(result) < count; result += strings.Repeat(" ", *interval) {
		result += strconv.Itoa(r.Intn(2))
	}
	return result[:count]
}

// getTerminalSize 获取终端大小
func getTerminalSize() (width, height int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin // 第一个拦路虎,需要设置 stdin, 否则执行失败

	output, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	fmt.Sscan(string(output), &height, &width)
	return
}

// clearTerminal 清屏
func clearTerminal() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func print(ch chan string, wc WordColor) {
	for {
		select {
		case out := <-ch:
			fmt.Printf(fmt.Sprintf("\033[1;%sm%%s\033[0m", wc.String()), out)
		default:
			//
		}
	}
}
