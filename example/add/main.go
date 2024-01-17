package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tidwall/gjson"

	nu "github.com/Inveracity/nu-plugin-golang/pkg/v1"
)

func ProcessCall(r gjson.Result) ([]byte, error) {
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
	response := nu.Response{}
	response.Value = nu.Value{Int: nu.Int{Val: endResult, InternalSpan: nu.InternalSpan{Start: 0, End: 0}}}

	jsonOutput, err := json.Marshal(response)
	if err != nil {
		return []byte{}, err
	}
	return jsonOutput, nil
}

func Plugin() (err error) {
	// Tell NuShell it's a plugin
	nu.Sendencoding()

	// Get input from nushell
	var input string
	fmt.Scan(&input)
	res := gjson.Parse(input)

	// If NuShell is requesting the signature, return the signature.
	// This is called when nushell calls `register` on it.
	if input == `"Signature"` {
		signature := Signatures()
		signatureJSON, _ := json.Marshal(map[string]interface{}{"Signature": []nu.Signature{signature}})
		nu.Send(signatureJSON)
		return nil
	}

	// Handle plugin input
	if res.Get("CallInfo.name").Str == "nu-golang" {
		// Debug([]byte(res.String()))
		response, err := ProcessCall(res.Get("CallInfo"))
		if err != nil {
			return err
		}
		// responseJSON, _ := json.Marshal(response)
		nu.Send(response)
		return nil
	}

	err_msg, err := nu.NewError("Plugin Error!", err.Error())
	if err != nil {
		return err
	}

	nu.Send(err_msg)

	return nil
}

func main() {
	err := Plugin()
	if err != nil {
		log.Panic(err)
	}
}
