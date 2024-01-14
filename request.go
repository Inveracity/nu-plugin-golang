package main

type Request struct {
	CallInfo struct {
		Name string `json:"name"`
		Call struct {
			Head struct {
				Start int `json:"start"`
				End   int `json:"end"`
			} `json:"head"`
			Positional []struct {
				Int struct {
					Val          int `json:"val"`
					InternalSpan struct {
						Start int `json:"start"`
						End   int `json:"end"`
					} `json:"internal_span"`
				} `json:"Int,omitempty"`
				String struct {
					Val          string `json:"val"`
					InternalSpan struct {
						Start int `json:"start"`
						End   int `json:"end"`
					} `json:"internal_span"`
				} `json:"String,omitempty"`
			} `json:"positional"`
			Named []interface{} `json:"named"`
		} `json:"call"`
		Input struct {
			Value struct {
				Nothing struct {
					InternalSpan struct {
						Start int `json:"start"`
						End   int `json:"end"`
					} `json:"internal_span"`
				} `json:"Nothing"`
			} `json:"Value"`
		} `json:"input"`
	} `json:"CallInfo"`
}
