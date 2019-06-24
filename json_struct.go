package pipeline_msg

import "time"

type (
	PipelineMessage struct {
		Fields []Var `json:"fields"`

		// Possible values:
		//   * "trace"
		//   * "debug"
		//   * "info"
		//   * "warn"
		//   * "error"
		//   * "fatal"
		Log_Level string `json:"log_level"`

		Log_Msg string `json:"log_msg"`

		Metrics []Var1 `json:"metrics"`

		Producer_ID string `json:"producer_id"`

		Product string `json:"product"`

		Service string `json:"service"`

		Team string `json:"team"`

		Timestamp time.Time `json:"timestamp"`

		// Possible values:
		//   * 0
		Version int64 `json:"version"`

		// Possible values:
		//   * "is24"
		//   * "as24"
		//   * "s24"
		Vertical string `json:"vertical"`
	}

	Var struct {
		Key string `json:"key"`

		Value string `json:"value"`
	}

	Var1 struct {
		Name string `json:"name"`

		// Possible values:
		//   * "COUNTER"
		//   * "GAUGE"
		Type string `json:"type"`

		Value float64 `json:"value"`
	}
)

