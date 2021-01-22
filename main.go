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
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "logs",
				Aliases: []string{"L"},
				Usage:   "Load logs from `FILES`",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() < 1 {
				fmt.Printf("请输入房态日志文件名来进行解析 \n")
				_ = cli.Exit("必须提供日志文件路径", 1)
			}
			logs := c.StringSlice("logs")
			fmt.Printf("解析日志文件： %v \n", logs)
			processLogFiles(logs)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func processLogFiles(logFiles []string) {
	initData()
	log.Printf("Start process %d log files: \n", len(logFiles))
	for _, f := range logFiles {
		log.Printf("process log file: %v \n", f)
		file, isGzip, _ := getFile(f)
		_ = process(file, isGzip)
		_ = file.Close()
		log.Printf("finish the log File \n")
	}
	outPut()
	log.Printf("All Done \n")
}
