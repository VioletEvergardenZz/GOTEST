package dfs

import (
	"fmt"
	"io/ioutil"
	"time"
)

var quire = "test"
var matches int
var workerCount = 0
var maxWorkerCount = 16
var searchRequest = make(chan string) //任务队列
var workerDone = make(chan bool)      //工人完成工作
var foundMatch = make(chan bool)      //找到匹配

func Dfs() {
	start := time.Now()
	workerCount = 1
	go search("/", true) //指派任务
	waitForWorkers()
	fmt.Println(matches, "matches")
	fmt.Println("time:", time.Since(start))
}

func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matches++
		}
	}
}
func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == quire {
				foundMatch <- true
			}
			if file.IsDir() {
				if workerCount < maxWorkerCount {
					searchRequest <- path + name + "/"
				} else {
					search(path+name+"/", false)
				}
			}
		}
	}
	if master {
		workerDone <- true
	}
}
