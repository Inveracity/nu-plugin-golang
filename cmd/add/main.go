package main

import (
	"encoding/json"
	"log"

	"github.com/tidwall/gjson"

	nu "github.com/Inveracity/nu-plugin-golang/pkg/v1"
)

func ProcessCall(r gjson.Result) ([]byte, error) {
	pos := r.Get("call.positional")
	res := nu.Ints{}
	err := json.Unmarshal([]byte(pos.String()), &res)

	if err != nil {
		return []byte{}, err
	}

	endResult := 0
	for _, item := range res {
		endResult += item.Int.Val
	}

	response := nu.Response{}
	response.Value = nu.Value{Int: nu.Int{Val: endResult, InternalSpan: nu.InternalSpan{Start: 0, End: 0}}}

	jsonOutput, err := json.Marshal(response)
	if err != nil {
		return []byte{}, err
	}
	return jsonOutput, nil
}

func main() {
	err := nu.Plugin(ProcessCall, Signatures, false)
	if err != nil {
		log.Fatal(err.Error())
	}
}
