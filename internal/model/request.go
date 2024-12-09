package model

type EvaluateRequest struct {
	Expression string `json:"expression"`
}

type EvaluateResponse struct {
	Result int `json:"result"`
}

type ValidateRequest struct {
	Expression string `json:"expression"`
}

type ValidateResponse struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}
