package main

import (
	"fmt"
	"log"
)

func main() {

	// initialize profiler service
	debug := NewProfilerHandler()

	go func() {
		log.Println("profiler at http://localhost:3030/debug/pprof")
		fmt.Println(debug.Server.ListenAndServe())
	}()

	// initialize web app
	app := NewAppHandler()

	log.Println("webapp at http://localhost:8080/fakework")
	log.Println(app.Server.ListenAndServe())

}
