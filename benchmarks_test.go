package pipeline_msg

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"math/rand"
	"testing"
	"time"
)

//https://www.calhoun.io/creating-random-strings-in-go/

const charset = `~!@#$%^&*()_+QWERTYUIOP{}|ASDFGHJKL:"ZXCVBNM<>?1234567890-=qwertyuiop[]\asdfghjkl;'zxcvbnm,./ '" ` + "`"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

var tester = New()

type field struct {
	key   string
	value string
}

type metric struct {
	name  string
	value float64
}

func makeMsgs(logMsg string, fields []field, metrics []metric) (*ProtobufPipelineMessage, *PipelineMessage) {

	now := time.Now()

	jsonVersion := &PipelineMessage{
		Version:     0,
		Vertical:    "s24",
		Producer_ID: "testing-producer-01",
		Timestamp:   now,
		Product:     "product-1",
		Team:        "sre-observability",
		Service:     "protobuf-vs-json",
		Log_Level:   "info",
		Log_Msg:     logMsg,
		Fields:      []Var{},
		Metrics:     []Var1{},
	}

	protobufVersion := &ProtobufPipelineMessage{
		Version:    0,
		ProducerId: jsonVersion.Producer_ID,
		Timestamp: &timestamp.Timestamp{
			Seconds: now.Unix(),
		},
		Product:  jsonVersion.Product,
		Team:     jsonVersion.Team,
		Service:  jsonVersion.Service,
		Vertical: ProtobufPipelineMessage_s24,
		LogLevel: ProtobufPipelineMessage_INFO,
		LogMsg:   logMsg,
	}

	for _, field := range fields {
		jsonVersion.Fields = append(jsonVersion.Fields, Var{Key: field.key, Value: field.value})
		protobufVersion.Fields = append(protobufVersion.Fields, &ProtobufPipelineMessage_Field{Key: field.key, Value: field.value})
	}

	for _, metric := range metrics {
		jsonVersion.Metrics = append(jsonVersion.Metrics, Var1{Name: metric.name, Value: metric.value, Type: "GAUGE"})
		protobufVersion.Metrics = append(protobufVersion.Metrics, &ProtobufPipelineMessage_Metric{Key: metric.name, Value: metric.value, Type: ProtobufPipelineMessage_GAUGE})
	}

	return protobufVersion, jsonVersion
}

func TestSerialisedSizes(t *testing.T) {

	type testcase struct {
		msg     string
		fields  []field
		metrics []metric
	}

	tests := map[string]testcase{
		"no_msg": {},
		"small_msg": {
			msg: "this is a small message",
		},
		"128b_msg": {
			msg: String(128),
		},
		"2kb_msg": {
			msg: String(2048),
		},
		"4kb_msg": {
			msg: String(4096),
		},
		"8kb_msg": {
			msg: String(8192),
		},
		"16kb_msg": {
			msg: String(8192 * 2),
		},
		"64kb_msg": {
			msg: String(65536),
		},
	}

	for i := 0; i <= 1024; i = 1 + i*2 {
		tc := testcase{}
		for j := 0; j < i; j++ {
			tc.fields = append(tc.fields, field{key: String(32), value: String(128)})
		}
		tests[fmt.Sprintf("fields,%d", i)] = tc
	}

	for i := 0; i <= 1024; i = 1 + i*2 {
		tc := testcase{}
		for j := 0; j < i; j++ {
			tc.metrics = append(tc.metrics, metric{name: String(32), value: 823749234.323})
		}
		tests[fmt.Sprintf("metrics,%d", i)] = tc
	}

	for test, details := range tests {
		p, j := makeMsgs(details.msg, details.fields, details.metrics)

		pSize := len(tester.SerialiseProtobuf(p))
		jSize := len(tester.SerialiseJSON(j))

		t.Run(fmt.Sprintf("%s,protobuf,size,%d", test, pSize), func(t *testing.T) {})
		t.Run(fmt.Sprintf("%s,json,size,%d", test, jSize), func(t *testing.T) {})
	}
}

func BenchmarkIt(b *testing.B) {

	type testcase struct {
		msg     string
		fields  []field
		metrics []metric
	}

	tests := map[string]testcase{
		"no_msg": {},
		"small_msg": {
			msg: "this is a small message",
		},
		"128b_msg": {
			msg: String(128),
		},
		"2kb_msg": {
			msg: String(2048),
		},
		"4kb_msg": {
			msg: String(4096),
		},
		"8kb_msg": {
			msg: String(8192),
		},
		"16kb_msg": {
			msg: String(8192 * 2),
		},
		"64kb_msg": {
			msg: String(65536),
		},
	}

	for i := 0; i <= 1024; i = 1 + i*2 {
		tc := testcase{}
		for j := 0; j < i; j++ {
			tc.fields = append(tc.fields, field{key: String(32), value: String(128)})
		}
		tests[fmt.Sprintf("fields,%d", i)] = tc
	}

	for i := 0; i <= 1024; i = 1 + i*2 {
		tc := testcase{}
		for j := 0; j < i; j++ {
			tc.metrics = append(tc.metrics, metric{name: String(32), value: 823749234.323})
		}
		tests[fmt.Sprintf("metrics,%d", i)] = tc
	}

	for test, details := range tests {
		p, j := makeMsgs(details.msg, details.fields, details.metrics)

		b.Run(fmt.Sprintf("serialise,protobuf,%s", test), func(b *testing.B) {
			var size int
			for i := 1; i < b.N; i++ {
				result := tester.SerialiseProtobuf(p)
				size += len(result)
			}
			b.SetBytes(int64(size))
		})

		b.Run(fmt.Sprintf("serialise,json,%s", test), func(b *testing.B) {
			var size int
			for i := 1; i < b.N; i++ {
				result := tester.SerialiseJSON(j)
				size += len(result)
			}
			b.SetBytes(int64(size))
		})

		b.Run(fmt.Sprintf("parse,protobuf,%s", test), func(b *testing.B) {
			b.StopTimer()
			buf := tester.SerialiseProtobuf(p)
			b.ResetTimer()
			b.StartTimer()
			var size int
			for i := 1; i < b.N; i++ {
				_ = tester.ParseProtobuf(buf)
				size += len(buf)
			}
			b.SetBytes(int64(len(buf)))
		})

		b.Run(fmt.Sprintf("parse,json,%s", test), func(b *testing.B) {
			b.StopTimer()
			buf := tester.SerialiseJSON(j)
			b.ResetTimer()
			b.StartTimer()
			var size int
			for i := 1; i < b.N; i++ {
				_ = tester.ParseJSON(buf)
				size += len(buf)
			}
			b.SetBytes(int64(len(buf)))
		})

		b.Run(fmt.Sprintf("parse,json-validate,%s", test), func(b *testing.B) {
			b.StopTimer()
			buf := tester.SerialiseJSON(j)
			b.ResetTimer()
			b.StartTimer()
			var size int
			for i := 1; i < b.N; i++ {
				_ = tester.ParseValidateJSON(buf)
				size += len(buf)
			}
			b.SetBytes(int64(len(buf)))
		})
	}
}

/*
func BenchmarkTester_SerialiseProtobuf(b *testing.B) {
	var size int
	for i := 1; i < b.N; i++ {
		result := tester.SerialiseProtobuf(validProtobuf)
		size += len(result)
	}
	b.SetBytes(int64(size))
}

func BenchmarkTester_SerialiseJSON(b *testing.B) {
	var size int
	for i := 1; i < b.N; i++ {
		result := tester.SerialiseJSON(validJSON)
		size += len(result)
	}
	b.SetBytes(int64(size))
}

func BenchmarkTester_ParseProtobuf(b *testing.B) {
	b.StopTimer()
	buf := tester.SerialiseProtobuf(validProtobuf)
	var size int
	b.ResetTimer()
	b.StartTimer()
	for i := 1; i < b.N; i++ {
		tester.ParseProtobuf(buf)
		size += len(buf)
	}
	b.SetBytes(int64(size))
}

func BenchmarkTester_ParseValidateJSON(b *testing.B) {
	var size int
	for i := 1; i < b.N; i++ {
		tester.ParseValidateJSON(validJSON)
		size += len(validJSON)
	}
	b.SetBytes(int64(size))
}

func BenchmarkTester_ParseJSON(b *testing.B) {
	var size int
	for i := 1; i < b.N; i++ {
		tester.ParseJSON(validJSON)
		size += len(validJSON)
	}
	b.SetBytes(int64(size))
}

*/
