// Package types define the structs used by the SDK
package types

import (
	"encoding/json"
)

type Workflow struct {
	ID          string          `json:"id,omitempty"`
	Name        string          `json:"name"`
	Nodes       json.RawMessage `json:"nodes"`
	Connections json.RawMessage `json:"connections,omitempty"`
	Active      bool            `json:"active,omitempty"`
	Settings    json.RawMessage `json:"settings,omitempty"`
	Tags        []string        `json:"tags,omitempty"`
}

// TODO: implement

func (w Workflow) CreateWorkflow() error {
	return nil
}
