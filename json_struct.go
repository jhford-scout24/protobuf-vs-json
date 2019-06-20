package pipeline_msg

import "time"

type (
	PipelineMessage struct {
		Fields []Var `json:"fields,omitempty"`

		// Possible values:
		//   * "trace"
		//   * "debug"
		//   * "info"
		//   * "warn"
		//   * "error"
		//   * "fatal"
		Log_Level string `json:"log_level,omitempty"`

		Log_Msg string `json:"log_msg,omitempty"`

		Metric []Var1 `json:"metric,omitempty"`

		Producer_ID string `json:"producer_id,omitempty"`

		Product string `json:"product,omitempty"`

		Service string `json:"service,omitempty"`

		Team string `json:"team,omitempty"`

		Timestamp time.Time `json:"timestamp,omitempty"`

		// Possible values:
		//   * 0
		Version int64 `json:"version,omitempty"`

		// Possible values:
		//   * "is24"
		//   * "as24"
		//   * "s24"
		Vertical string `json:"vertical,omitempty"`
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

