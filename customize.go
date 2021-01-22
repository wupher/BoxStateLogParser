package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var gMutex sync.Mutex
var theSet StatusSet
var companySet CompanySet
var errCompany map[string]int
var occurAt map[string]int
var logCount = 0

var targetModel = "BoxStateHashCode"

func initData() {
	theSet = make(map[int]struct{}, 0)
	companySet = make(map[string]struct{}, 0)
	errCompany = make(map[string]int, 0)
	occurAt = make(map[string]int, 0)
}

func lineProcess(line string) {
	//pre process
	r, _ := regexp.Compile(targetModel)
	if !r.MatchString(line) {
		return
	}

	logData, err := UnmarshallLog(line)
	if err != nil {
		return
	}
	//post process

	//Update Global Data
	gMutex.Lock()
	if logData.CompanyCode == "" { //居然没有商家编码
		errCompany[""]++
	}

	if logData.Model == targetModel {
		if theSet.Has(logData.Status) { //出现重复数据
			errCompany[logData.CompanyCode]++
			hour := logData.RequestTime.Hour()
			occurAt[strconv.Itoa(hour)]++
		} else {
			theSet.Add(logData.Status)
		}
		logCount++
		companySet.Add(logData.CompanyCode)
	}
	gMutex.Unlock()
}

func outPut() {
	fmt.Printf("共处理了 %d 条房态 HashCode 日志数据 \n", logCount)
	errSize := len(errCompany)
	companySize := len(companySet)
	per := float64(errSize) / float64(companySize) * 100
	fmt.Printf("其中有 %d 个商家出现房态重复，共记：%d 个商家， 占比 %0.2f", errSize, companySize, per)
	fmt.Println("%")

	errCount := 0
	for _, count := range errCompany {
		errCount += count
	}

	fmt.Printf("共有 %d 条重复房态消息 ", errCount)
	percent := float64(errCount) / float64(logCount) * 100

	fmt.Printf("占比： %0.2f", percent)
	fmt.Println("%")

	sortErrCompany := sortByErrCount(errCompany)
	for _, v := range sortErrCompany {
		fmt.Printf("商家编码：%v , 重复次数: %d \n", v.CompanyCode, v.ErrCount)
	}

	sortOccurAt := sortByErrCount(occurAt)
	for _, v := range sortOccurAt {
		fmt.Printf("时间：%v, 重复次数：%d \n", v.CompanyCode, v.ErrCount)
	}
}

func UnmarshallLog(line string) (logData LogSt, err error) {
	//pre process
	err = json.Unmarshal([]byte(line), &logData)
	if err != nil {
		log.Fatalf("Log Content: %q process erro! %q \n", line, err)
		return
	}
	//post process
	return
}

type LogSt struct {
	CompanyCode string    `json:"companycode"`
	Model       string    `json:"model"`
	RequestTime time.Time `json:"request_time"`
	Msg         string    `json:"msg"`
	Status      int       `json:"status"`
}

func (log LogSt) Convert() []string {
	on := log.RequestTime.Format("2006-01-02 15:04:05")

	status := strconv.Itoa(log.Status)
	return []string{log.CompanyCode, log.Model, on, log.Msg, status}
}
