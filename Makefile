test:
	go test -bench=. -test.benchmem -v -timeout 20m

pipeline_msg.pb.go: pipeline_msg.proto
	protoc -I=. --go_out=. $<
