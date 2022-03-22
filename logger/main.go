package main

import (
	"fmt"
	"github.com/gongxulei/go_kit/logger/driver"
	"time"
)

func main() {
	log := driver.NewFileLogger(0, "/Volumes/work/bb_golang/src/gong/logger", "test")
	for true {
		time.Sleep(time.Second * 1)
		log.Log("this is a Log %s", "测试")
	}
	//aaaaa(nil)
	//atomic.AddInt32()
}

func aaaaa(a []*int) {
	fmt.Println(len(a))
}
