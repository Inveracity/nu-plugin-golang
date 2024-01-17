package v1

import "encoding/json"

type Error struct {
	Label string         `json:"label"`
	Msg   string         `json:"msg"`
	Span  map[string]int `json:"span"`
}

func NewError(label string, msg string) ([]byte, error) {
	errorMsg := Error{
		Label: label,
		Msg:   msg,
		Span:  map[string]int{"start": 0, "end": 1}, // TODO: figure this out eventually
	}

	errorJSON, err := json.Marshal(errorMsg)

	if err != nil {
		return []byte{}, err
	}

	return errorJSON, nil

}
