.PHONY:

cpu-profile:
	curl -o prof.pb.gz http://localhost:3030/debug/pprof/profile\?seconds\=30

memory-profile:
	curl -o prof.pb.gz http://localhost:3030/debug/pprof/heap\?seconds\=30

trace:
	curl -o trace.out http://localhost:3030/debug/pprof/trace\?seconds\=5

test-load:
	while true; do curl http://0.0.0.0:8080/fakework | jq ; sleep .5; done

run:
	go run ./...