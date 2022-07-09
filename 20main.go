package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	var data []string = read_file("data.txt")
	var timeArray []time.Duration
	for _, line := range data {
		if len(line) > 1 {
			line = strings.ReplaceAll(line, "\x0d", "")
			line = strings.ReplaceAll(line, " ", "")
			var time_el, _ = time.ParseDuration(line)
			//fmt.Println(time_el, "  ", err)
			timeArray = append(timeArray, time_el)
		}
	}
	fmt.Println(timeArray)
	fmt.Println("Приступаю к выполнению задач")
	for _, sleepTime := range timeArray {
		go workFunc(sleepTime)
	}
	fmt.Scanln()
	fmt.Println("Окончил выполнять задачи")
}

func read_file(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	split_data := strings.Split(string(data), "\n")
	return split_data
}

func workFunc(sleepTime time.Duration) {
	fmt.Println("Выполняю задачу длительностью: ", sleepTime)
	time.Sleep(sleepTime)
	fmt.Println("( ˘⌣˘)♡(˘⌣˘ ) ♡ Окончил выполнять задачу длительностью: ", sleepTime)
}

// 7h3m45s

// 1h2m17s

// 3h32m20s

// 4h35m20s
