package pipeline_msg

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"testing"
	"time"
)

var tester = New()

var validProtobuf = &ProtobufPipelineMessage{
	Version:    0,
	ProducerId: "testing-producer-01",
	Timestamp: &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
	},
	Product:  "product-1",
	Team:     "sre-observability",
	Service:  "protobuf-vs-json",
	Vertical: ProtobufPipelineMessage_s24,
	LogLevel: ProtobufPipelineMessage_INFO,
	LogMsg:   "this is a test log line",
	Fields: []*ProtobufPipelineMessage_Field{
		&ProtobufPipelineMessage_Field{
			Key:   "key1",
			Value: "value1",
		},
		&ProtobufPipelineMessage_Field{
			Key:   "key2",
			Value: "value2",
		},
	},
}

var validJSON = []byte(`{
	"version": 0,
	"producer_id": "testing-producer-01",
	"timestamp": "2008-09-15T15:53:00+05:00",
	"product": "product-1",
	"team": "sre-observability",
	"service": "protobuf-vs-json",
	"vertical": "s24",
	"log_level": "info",
	"log_msg": "this is a test log line",
	"fields": [
		{"key": "key1", "value": "value1"},
		{"key": "key2", "value": "value2"}
    ],
	"metrics": []
}`)

func TestSerialisedSizes(t *testing.T) {
	protoSize := len(tester.SerialiseProtobuf(validProtobuf))
	jsonSize := len(validJSON)

	t.Run(fmt.Sprintf("PROTOBUF: %d bytes", protoSize), func(t *testing.T) {})
	t.Run(fmt.Sprintf("JSON:     %d bytes", jsonSize), func(t *testing.T) {})
}

func BenchmarkTester_SerialiseProtobuf(b *testing.B) {
	for i := 1; i < b.N; i++ {
		result := tester.SerialiseProtobuf(validProtobuf)
		b.SetBytes(int64(len(result)))
	}
}

func BenchmarkTester_SerialiseJSON(b *testing.B) {
	b.StopTimer()
	example := tester.ParseJSON(validJSON)
	b.ResetTimer()
	b.StartTimer()
	for i := 1; i < b.N; i++ {
		result := tester.SerialiseJSON(example)
		b.SetBytes(int64(len(result)))
	}
}

func BenchmarkTester_ParseProtobuf(b *testing.B) {
	b.StopTimer()
	buf := tester.SerialiseProtobuf(validProtobuf)
	b.ResetTimer()
	b.StartTimer()
	for i := 1; i < b.N; i++ {
		tester.ParseProtobuf(buf)
		b.SetBytes(int64(len(buf)))
	}
}

func BenchmarkTester_ParseValidateJSON(b *testing.B) {
	for i := 1; i < b.N; i++ {
		tester.ParseJSON(tester.ValidateJSON(validJSON))
		b.SetBytes(int64(len(validJSON)))
	}
}

func BenchmarkTester_ParseJSON(b *testing.B) {
	for i := 1; i < b.N; i++ {
		tester.ParseJSON(validJSON)
		b.SetBytes(int64(len(validJSON)))
	}
}

func BenchmarkTester_ValidateJSON(b *testing.B) {
	for i := 1; i < b.N; i++ {
		tester.ValidateJSON(validJSON)
		b.SetBytes(int64(len(validJSON)))
	}
}
