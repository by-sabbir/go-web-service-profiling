package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// AppHandler declares the business service structure.
type AppHandler struct {
	Server *http.Server
	Route  *http.ServeMux
}

// NewAppHandler initiates the AppHandler and Maps the internal routes.
func NewAppHandler() *AppHandler {
	h := &AppHandler{}

	h.Route = http.NewServeMux()

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Route,
	}

	h.mapAppHandler()

	return h
}

func (a *AppHandler) mapAppHandler() {
	a.Route.HandleFunc("/fakework", fakeWork)
}

func fakeWork(w http.ResponseWriter, r *http.Request) {

	ch := make(chan map[string]int)
	for i := 0; i < 1000; i++ {
		go work(i, ch)
	}

	for i := 0; i < 1000; i++ {
		mb, _ := json.Marshal(<-ch)
		w.Write(mb)
	}
}

func work(task int, ch chan map[string]int) {
	s := rand.NewSource(time.Now().UnixNano())

	r := rand.New(s)
	ch <- map[string]int{
		"task no":          task,
		"generated number": r.Int(),
	}

}
