package v1

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

func Plugin(processer func(r gjson.Result) ([]byte, error), signatures func() Signature, debug bool) (err error) {
	err = nil
	Sendencoding() // Tell NuShell it's a plugin

	// Get input from nushell
	var input string
	fmt.Scan(&input)
	res := gjson.Parse(input)

	// If NuShell is requesting the signature, return the signature. This is called when nushell calls `register` on it.
	if input == `"Signature"` {
		signature := signatures()
		signatureJSON, err := json.Marshal(map[string]interface{}{"Signature": []Signature{signature}})
		if debug {
			Debug(signatureJSON)
		}
		Send(signatureJSON)
		return err
	}

	// Handle plugin input
	if res.Get("CallInfo.name").Exists() {
		if debug {
			Debug([]byte(res.String()))
		}
		response, err := processer(res.Get("CallInfo"))
		if debug {
			Debug(response)
		}
		Send(response)
		return err
	}

	if err != nil {
		err_msg, err := NewError("Plugin Error!", err.Error())
		if err != nil {
			return err
		}

		Send(err_msg)
	}

	return err
}
