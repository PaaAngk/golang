package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Task struct {
	taskTime time.Duration
	taskType string
}

var mu sync.Mutex
var count int

var wg sync.WaitGroup

var (
	lock sync.Mutex
	//taskQueue map[int]Task
	taskQueue chan Task = make(chan Task)
	//taskQueue chan time.Duration = make(chan time.Duration)*sync.Mutex
)
var ex chan bool = make(chan bool)

func main() {
	//var taskQueue chan Task = make(chan Task)

	// wg.Add(1)
	// go func() {
	// 	for {
	// 		lock.Lock()
	// 		val, ok := <-taskQueue
	// 		if !ok {
	// 			fmt.Println("break")
	// 			break
	// 		}
	// 		fmt.Println("1 Sleep for ", val, " taskQueue ", len(taskQueue))
	// 		time.Sleep(val)
	// 		fmt.Println("!!! Sleep done")
	// 		lock.Unlock()
	// 	}
	// 	wg.Done()
	// }()

	http.HandleFunc("/schedule", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			mu.Lock()
			count++
			mu.Unlock()

			fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
		} else {
			fmt.Fprintf(w, "Sorry, only Get methods this link are supported.")
		}

	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || r.Method == "PUT" || r.Method == "DELETE" {
			fmt.Fprintf(w, "Sorry, only POST methods this link are supported.")
			return
		}
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		var queryTime, _ = time.ParseDuration(string(r.FormValue("time")))
		var queryType = string(r.FormValue("type"))
		var task = Task{taskTime: queryTime, taskType: queryType}

		fmt.Println("Post from website! time=", task.taskTime, " type=", task.taskType)

		if queryType == "async" {
			//fmt.Fprintf(w, "Отображение выбранной задачи с queryType %d...", queryType)
			w.WriteHeader(200)
			fmt.Println("WriteHeader")
			go func() {
				taskQueue <- task
			}()
			handq(taskQueue)
		} else {
			//mu.Lock()
			taskQueue <- task
			handq(taskQueue)
			<-taskQueue
			w.WriteHeader(200)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
func handq(taskQueue chan Task) {
	lock.Lock()
	var ta Task
	for r := range taskQueue {
		fmt.Println("1 Sleep for ", r, " taskQueue ", len(taskQueue))
		time.Sleep(r.taskTime)
		ta = r
		fmt.Println("!!! Sleep done")
	}
	lock.Unlock()
	taskQueue <- ta
}
