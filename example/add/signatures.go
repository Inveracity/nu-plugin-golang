package main

import (
	nu "github.com/Inveracity/nu-plugin-golang/pkg/v1"
)

// This Plugins Specific Signature, telling NuShell what this plugin is expecting to receive.
func Signatures() nu.Signature {
	return nu.Signature{
		Sig: nu.SignatureDetails{
			Name:       "nu-golang",
			Usage:      "Signature test for golang",
			ExtraUsage: "",
			InputType:  "Any",
			OutputType: "Any",
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
			OptionalPositional: []nu.PositionalArg{
				{
					Name:  "opt",
					Desc:  "Optional number",
					Shape: "Int",
					VarID: nil,
				},
			},
			RestPositional: nu.PositionalArg{
				Name:  "rest",
				Desc:  "rest value string",
				Shape: "String",
				VarID: nil,
			},
			Named: []nu.NamedArg{
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
