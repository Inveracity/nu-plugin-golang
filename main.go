package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func ProcessCall(r Request) Response {

	pos := r.CallInfo.Call.Positional

	return_value := 0
	for _, values := range pos {
		return_value = values.Int.Val
	}

	return Response{
		Value{
			Int{
				Val: return_value,
				InternalSpan: InternalSpan{
					Start: 0, End: 1,
				},
			},
		},
	}
}

func Plugin() {
	// Tell NuShell it's a plugin
	Sendencoding()

	// Get input from nushell
	var input string
	fmt.Scan(&input)

	// Convert JSON to struct
	pluginCall := Request{}
	json.Unmarshal([]byte(input), &pluginCall)

	// If NuShell is requesting the signature, return the signature.
	// This is called when nushell calls `register` on it.
	if input == "\"Signature\"" {
		signature := Signatures()
		signatureJSON, _ := json.Marshal(map[string]interface{}{"Signature": []Signature{signature}})
		Return(signatureJSON)
		return
	}

	// Handle plugin input
	if pluginCall.CallInfo.Name == "nu-golang" {
		response := ProcessCall(pluginCall)
		responseJSON, _ := json.Marshal(response)
		Return(responseJSON)
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
