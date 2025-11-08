package dfs

import (
	"fmt"
	"io/ioutil"
	"time"
)

var queryDemo = "test"
var matchesDemo int

func DfsDemo() {
	start := time.Now()
	searchDemo("/")
	fmt.Println(matchesDemo, "matches")
	fmt.Println("time:", time.Since(start))
}
func searchDemo(path string) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == queryDemo {
				matchesDemo++
			}
			if file.IsDir() {
				searchDemo(path + name + "/")
			}
		}
	}
}
