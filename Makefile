test:
	go test -bench=. -test.benchmem -test.benchtime=5s -v

pipeline_msg.pb.go: pipeline_msg.proto
	protoc -I=. --go_out=. $<
