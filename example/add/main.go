package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/gjson"

	nu "github.com/Inveracity/nu-plugin-golang/pkg/v1"
)

func ProcessCall(r gjson.Result) []byte {
	pos := r.Get("call.positional")

	// Debug([]byte(pos.String()))

	res := nu.Ints{}
	err := json.Unmarshal([]byte(pos.String()), &res)

	if err != nil {
		nu.Debug([]byte(err.Error()))
	}

	endResult := 0
	for _, item := range res {
		// Debug([]byte(fmt.Sprint(item.Int.Val)))

		endResult += item.Int.Val
	}

	newint := nu.Response{
		Value: Int{
			Val:          endResult,
			InternalSpan: nu.InternalSpan{Start: 0, End: 1},
		},
	}
	return []byte(`{
		"Value": {
		  "Int": {
			"val": ` + fmt.Sprint(endResult) + `,
			"internal_span": {
			  "start": 100953,
			  "end": 100957
			}
		  }
		}
	  }`)
}

func Plugin() {
	// Tell NuShell it's a plugin
	nu.Sendencoding()

	// Get input from nushell
	var input string
	fmt.Scan(&input)
	res := gjson.Parse(input)

	// If NuShell is requesting the signature, return the signature.
	// This is called when nushell calls `register` on it.
	if input == "\"Signature\"" {
		signature := Signatures()
		signatureJSON, _ := json.Marshal(map[string]interface{}{"Signature": []nu.Signature{signature}})
		nu.Send(signatureJSON)
		return
	}

	// Handle plugin input
	if res.Get("CallInfo.name").Str == "nu-golang" {
		// Debug([]byte(res.String()))
		response := ProcessCall(res.Get("CallInfo"))
		// responseJSON, _ := json.Marshal(response)
		nu.Send(response)
		return
	}

	// TODO: Figure out how to handle error and return them nicely, probably make a function out of this
	errorMsg := nu.Error{
		Label: "ERROR from plugin",
		Msg:   "error message pointing to call head span",
		Span:  map[string]int{"start": 0, "end": 1},
	}

	errorJSON, _ := json.Marshal(errorMsg)
	fmt.Print(string(errorJSON))
	os.Stdout.Sync()
}

func main() {
	Plugin()
}
