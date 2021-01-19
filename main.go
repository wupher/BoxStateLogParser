package main

func main() {
	logFile := ""
	initData()
	file, isGzip, _ := getFile(logFile)
	defer file.Close()
	_ = process(file, isGzip)
	outPut()
}
