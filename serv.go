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
	SessionsMapMutex sync.Mutex
	SessionsMap      map[string]*sync.Mutex
)

func RequestWorker(session_id string) {
	var mutex *Mutex
	var ok bool
	SessionsMapMutex.Lock()
	if mutex, ok = SessionsMap[session_id]; !ok {
		mutex = &Mutex{}
		SessionsMap[session_id] = mutex
	}
	SessionsMapMutex.Unlock()
	mutex.Lock()
	defer mutex.Unlock()
}

var (
	lock      sync.Mutex
	taskQueue map[int]Task
	//taskQueue chan time.Duration = make(chan time.Duration)*sync.Mutex
)

func main() {

	wg.Add(1)
	go func() {
		for {
			lock.Lock()
			task
			fmt.Println("1 Sleep for ", taskQueue, " taskQueue ", len(taskQueue))
			time.Sleep(taskQueue)
			fmt.Println("!!! Sleep done")
			lock.Unlock()
		}
		wg.Done()
	}()

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello World!") })
	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	})
	http.HandleFunc("/co", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count %d\n", count)
		mu.Unlock()
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

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
			taskQueue[len(taskQueue)] = task
			go func() { taskQueue <- queryTime }()
		} else {
			//mu.Lock()
			taskQueue <- queryTime
			//mu.Unlock()
			//wg.Wait()
			w.WriteHeader(200)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
