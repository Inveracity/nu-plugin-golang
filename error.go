package main

type Error struct {
	Label string         `json:"label"`
	Msg   string         `json:"msg"`
	Span  map[string]int `json:"span"`
}
