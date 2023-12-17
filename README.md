# Web Service Profiling in Go

## Quick Guide

- Open terminal and execute

```bash
make run
```

this will run the webservice.

- To generate load open another terminal and execute

```bash
make test-load
```

- To download the profile (cpu/memory)

```bash
make {cpu}/{memory}-profile
```

this will record 30 seconds session for the web app.

- To download the trace

```bash
make trace
```

this will record 5 seconds session for the web app.


- To view the profile (cpu/memory)
```bash
go tool pprof -http localhost:9999 prof.pb.gz
```

- To view the traces
```bash
go tool trace trace.out
```
