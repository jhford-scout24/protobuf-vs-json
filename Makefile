pipeline_msg.pb.go: pipeline_msg.proto
	protoc -I=. --go_out=. $<
