test:
	go test -bench=. -test.benchmem

pipeline_msg.pb.go: pipeline_msg.proto
	protoc -I=. --go_out=. $<
