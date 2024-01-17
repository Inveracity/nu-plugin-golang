package main

type InternalSpan struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Int struct {
	Val          int          `json:"val"`
	InternalSpan InternalSpan `json:"internal_span"`
}

type Ints []struct {
	Int Int `json:"Int"`
}

type Value struct {
	Int Int `json:"Int"`
}

type Response struct {
	Value Value `json:"Value"`
}
