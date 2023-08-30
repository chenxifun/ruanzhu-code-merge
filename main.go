package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/chenxifun/ruanzhu-code-merge/comm"
	"github.com/chenxifun/ruanzhu-code-merge/core"
)

var config string

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "合并代码到同一个文件\n")
		flag.PrintDefaults()
	}
	flag.StringVar(&config, "c", "config.json", "配置文件")
	flag.Parse()

	fmt.Println("读取配置文件", config)
	confBytes, err := comm.ReadFile(config)
	if err != nil {
		fmt.Println("读取配置文件失败", err.Error())
		return
	}

	c := &core.CodeArrange{}
	err = json.Unmarshal(confBytes, c)
	if err != nil {
		fmt.Println("读取配置文件失败", err.Error())
		return
	}

	err = c.Start()
	if err != nil {
		fmt.Println("合并失败", err.Error())
		return
	}

	fmt.Println("合并成功，输入文件：", c.OutPath)
}
