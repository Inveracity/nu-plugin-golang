package main

import (
	nu "github.com/Inveracity/nu-plugin-golang/pkg/v1"
)

// This Plugins Specific Signature, telling NuShell what this plugin is expecting to receive.
func Signatures() nu.Signature {
	return nu.Signature{
		Sig: nu.SignatureDetails{
			Name:       "addv2",
			Usage:      "Signature test for golang",
			ExtraUsage: "",
			InputType:  "Nothing",
			OutputType: "Nothing",
			RequiredPositional: []nu.PositionalArg{
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
			OptionalPositional: []nu.PositionalArg{},
			RestPositional:     nil,
			Named: []nu.NamedArg{
				{
					Long:     "help",
					Short:    "h",
					Arg:      nil,
					Required: false,
					Desc:     "Display the help message for this command",
					VarID:    nil,
				},
			},
			InputOutputTypes:             [][]string{{"Nothing", "Nothing"}},
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
