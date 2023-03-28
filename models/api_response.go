package models

import "encoding/json"

func UnmarshalAPIResponse(data []byte) (APIResponse, error) {
	var r APIResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *APIResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type APIResponse struct {
	Status bool   `json:"status"`
	Detail string `json:"detail"`
}
