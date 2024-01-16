package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
)

// This Plugins Specific Signature, telling NuShell what this plugin is expecting to receive.
func Signatures() Signature {
	return Signature{
		Sig: SignatureDetails{
			Name:       "nu-golang",
			Usage:      "Signature test for golang",
			ExtraUsage: "",
			InputType:  "Any",
			OutputType: "Any",
			RequiredPositional: []PositionalArg{
				{
					Name:  "a",
					Desc:  "required integer value",
					Shape: "Int",
					VarID: nil,
				},
				{
					Name:  "b",
					Desc:  "required integer value",
					Shape: "Int",
					VarID: nil,
				},
			},
			OptionalPositional: []PositionalArg{
				{
					Name:  "opt",
					Desc:  "Optional number",
					Shape: "Int",
					VarID: nil,
				},
			},
			RestPositional: PositionalArg{
				Name:  "rest",
				Desc:  "rest value string",
				Shape: "String",
				VarID: nil,
			},
			Named: []NamedArg{
				{
					Long:     "help",
					Short:    "h",
					Arg:      nil,
					Required: false,
					Desc:     "Display the help message for this command",
					VarID:    nil,
				},
				{
					Long:     "flag",
					Short:    "f",
					Arg:      nil,
					Required: false,
					Desc:     "a flag for the signature",
					VarID:    nil,
				},
				{
					Long:     "named",
					Short:    "n",
					Arg:      "String",
					Required: false,
					Desc:     "named string",
					VarID:    nil,
				},
			},
			InputOutputTypes:             [][]string{{"Any", "Any"}},
			AllowVariantsWithoutExamples: true,
			SearchTerms:                  []string{"golang", "Example"},
			IsFilter:                     false,
			CreatesScope:                 false,
			AllowsUnknownArgs:            false,
			Category:                     "Experimental",
		},
		Examples: []interface{}{},
	}
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

func ProcessCall(r gjson.Result) []byte {
	pos := r.Get("call.positional")

	Debug([]byte(pos.String()))

	res := Ints{}
	err := json.Unmarshal([]byte(pos.String()), &res)

	if err != nil {
		Debug([]byte(err.Error()))
	}

	endResult := 0
	for _, item := range res {
		Debug([]byte(fmt.Sprint(item.Int.Val)))

		endResult += item.Int.Val
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
	Sendencoding()

	// Get input from nushell
	var input string
	fmt.Scan(&input)
	res := gjson.Parse(input)

	// If NuShell is requesting the signature, return the signature.
	// This is called when nushell calls `register` on it.
	if input == "\"Signature\"" {
		signature := Signatures()
		signatureJSON, _ := json.Marshal(map[string]interface{}{"Signature": []Signature{signature}})
		Send(signatureJSON)
		return
	}

	// Handle plugin input
	if res.Get("CallInfo.name").Str == "nu-golang" {
		Debug([]byte(res.String()))
		response := ProcessCall(res.Get("CallInfo"))
		// responseJSON, _ := json.Marshal(response)
		Send(response)
		return
	}

	// TODO: Figure out how to handle error and return them nicely, probably make a function out of this
	errorMsg := Error{
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
