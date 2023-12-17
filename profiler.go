package main

import (
	"net/http"
	"net/http/pprof"
)

// ProfilerHandler declares the profiler service structure.
type ProfilerHandler struct {
	Server *http.Server
	Route  *http.ServeMux
}

// NewProfilerHandler initiates the ProfilerHandler and Maps the profiler routes.
func NewProfilerHandler() *ProfilerHandler {

	h := &ProfilerHandler{}

	h.Route = http.NewServeMux()

	h.Server = &http.Server{
		Addr: "0.0.0.0:3030",
	}
	h.mapProfileHandlers()

	return h
}

func (p *ProfilerHandler) mapProfileHandlers() {

	p.Route.HandleFunc("/debug/pprof/profile", pprof.Profile)
	p.Route.HandleFunc("/debug/pprof/heap", pprof.Index)
	p.Route.HandleFunc("/debug/pprof/trace", pprof.Trace)
	p.Route.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

}
