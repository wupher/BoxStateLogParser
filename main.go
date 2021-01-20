package main

import (
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "box_log_count",
		Usage: "计数包厢房态",
		Action: func(c *cli.Context) error {
			if c.Args().Len() < 1 {
				fmt.Printf("请输入房态日志文件名来进行解析 \n")
				_ = cli.Exit("必须提供日志文件路径", 1)
			}
			fileName := c.Args().Get(0)
			fmt.Printf("解析日志文件： %v \n", fileName)
			processLogFile(fileName)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func processLogFile(logFile string) {
	initData()
	file, isGzip, _ := getFile(logFile)
	defer file.Close()
	_ = process(file, isGzip)
	outPut()
}
