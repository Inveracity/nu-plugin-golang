package main

type Error struct {
	Label string         `json:"label"`
	Msg   string         `json:"msg"`
	Span  map[string]int `json:"span"`
}

type Ints []struct {
	Int struct {
		Val          int `json:"val"`
		InternalSpan struct {
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"internal_span"`
	} `json:"Int"`
}
