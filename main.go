package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	Path string = "config.json"
	conf        = new(Addrs)
)

type Addrs struct {
	Addr []Connect
}

func main() {
	StartAction(conf.Addr)
	select {}
}

func init() {
	initParams()
	loadConfig()
}

func loadConfig() {
	by, err := ioutil.ReadFile(Path)
	if err != nil {
		log.Println("读取配置文件失败", err.Error())
		os.Exit(1)
	}
	if err = json.Unmarshal(by, conf); err != nil {
		log.Println("解析配置文件失败", err.Error())
		os.Exit(1)
	}
	if len(conf.Addr) < 1 {
		log.Println("无代理配置")
		os.Exit(1)
	}
}

func initParams() {
	var help bool
	flag.BoolVar(&help, "h", false, "帮助说明")
	flag.Int64Var(&ReTime, "r", 10, "连接中断后重新连接的等待时间")
	flag.Int64Var(&HeartBeatTime, "t", 3, "ssh连接状态检测")
	flag.StringVar(&Path, "c", "conf.json", "系统配置文件")
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	Fileabs(Path)
	log.Println("参数解析完成")
}

func Fileabs(cpath string) string {
	Appath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println("当前路径获取失败...", err.Error())
		os.Exit(1)
	}
	if !strings.HasPrefix(cpath, "/") {
		cpath = filepath.Join(Appath, cpath)
	}
	return cpath
}
