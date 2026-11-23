package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func cloudNamedEnum(value json.RawMessage, numbers map[string]string) (*string, error) {
	value = bytes.TrimSpace(value)
	if len(value) == 0 || bytes.Equal(value, []byte("null")) {
		return nil, nil
	}
	if value[0] == '"' {
		var name string
		if err := json.Unmarshal(value, &name); err != nil {
			return nil, err
		}
		return &name, nil
	}
	if name, ok := numbers[string(value)]; ok {
		return &name, nil
	}
	return nil, fmt.Errorf("unsupported numeric enum value %s", value)
}
