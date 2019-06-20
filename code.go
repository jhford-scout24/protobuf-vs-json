package pipeline_msg

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/qri-io/jsonschema"
	"io/ioutil"
	"os"
	"strings"
)

type Tester struct {
	rs *jsonschema.RootSchema
}

func New() *Tester {
	t := &Tester{
		rs: &jsonschema.RootSchema{},
	}

	sf, err := os.Open("schema.json")
	if err != nil {
		panic(err)
	}
	defer sf.Close()

	sb, err := ioutil.ReadAll(sf)

	if err := json.Unmarshal(sb, t.rs); err != nil {
		panic(err)
	}

	return t
}

func (t *Tester) ParseProtobuf(buf []byte) *ProtobufPipelineMessage {
	msg := &ProtobufPipelineMessage{}
	err := proto.Unmarshal(buf, msg)
	if err != nil {
		panic(err)
	}
	return msg
}

func (t *Tester) SerialiseProtobuf(message *ProtobufPipelineMessage) []byte {
	out, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}
	return out
}

func (t *Tester) ParseJSON(buf []byte) *PipelineMessage {
	msg := &PipelineMessage{}
	err := json.Unmarshal(buf, msg)
	if err != nil {
		panic(err)
	}
	return msg
}

func (t *Tester) SerialiseJSON(message *PipelineMessage) []byte {
	out, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return out
}

func (t *Tester) ParseValidateJSON(buf []byte) []byte {
	var doc interface{}
	errors := []jsonschema.ValError{}

	doc = t.ParseJSON(buf)

	t.rs.Validate("/", doc, &errors)

	if errors, _ := t.rs.ValidateBytes(buf); len(errors) > 0 {
		var errorMsg strings.Builder
		for _, error := range errors {
			errorMsg.WriteString(error.Error())
		}
		panic(errorMsg)
	}

	return buf
}
