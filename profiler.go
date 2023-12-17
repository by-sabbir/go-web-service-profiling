package main

import (
	"net/http"
	"net/http/pprof"
)

// ProfilerHandler declares the profiler service structure.
type ProfilerHandler struct {
	Server *http.Server
	Path   *http.ServeMux
}

// NewProfilerHandler initiates the ProfilerHandler and Maps the profiler routes.
func NewProfilerHandler() *ProfilerHandler {

	h := &ProfilerHandler{}

	h.Path = http.NewServeMux()

	h.Server = &http.Server{
		Addr: "0.0.0.0:3030",
	}
	h.mapProfileHandlers()

	return h
}

func (p *ProfilerHandler) mapProfileHandlers() {

	p.Path.HandleFunc("/debug/pprof/profile", pprof.Profile)
	p.Path.HandleFunc("/debug/pprof/heap", pprof.Index)
	p.Path.HandleFunc("/debug/pprof/trace", pprof.Trace)
	p.Path.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

}
