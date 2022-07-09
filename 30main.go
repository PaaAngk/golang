package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var thread_nubmer int

	var data []string = read_file("data.txt")
	var timeArray []time.Duration
	for _, line := range data {
		if len(line) > 1 {
			line = strings.ReplaceAll(line, "\x0d", "")
			line = strings.ReplaceAll(line, " ", "")
			var time_el, _ = time.ParseDuration(line)
			timeArray = append(timeArray, time_el)
		}
	}

	var wg sync.WaitGroup

	buffer := make(chan time.Duration, len(timeArray))

	wg.Add(1)
	go func() {
		for i := 0; i < len(timeArray); i++ {
			buffer <- timeArray[i]
		}
		close(buffer)
		wg.Done()
	}()

	fmt.Println("Введите количество процессоров")
	fmt.Fscan(os.Stdin, &thread_nubmer)

	for i := 0; i < thread_nubmer; i++ {
		wg.Add(1)
		go func() {
			for {
				val, ok := <-buffer
				if !ok {
					break
				}
				workFunc(val)
			}
			wg.Done()
		}()
	}

	wg.Wait()

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
