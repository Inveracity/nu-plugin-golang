package main

type Signature struct {
	Sig      SignatureDetails `json:"sig"`
	Examples []interface{}    `json:"examples"`
}

type SignatureDetails struct {
	Name                         string          `json:"name"`
	Usage                        string          `json:"usage"`
	ExtraUsage                   string          `json:"extra_usage"`
	InputType                    string          `json:"input_type"`
	OutputType                   string          `json:"output_type"`
	RequiredPositional           []PositionalArg `json:"required_positional"`
	OptionalPositional           []PositionalArg `json:"optional_positional"`
	RestPositional               PositionalArg   `json:"rest_positional"`
	Named                        []NamedArg      `json:"named"`
	InputOutputTypes             [][]string      `json:"input_output_types"`
	AllowVariantsWithoutExamples bool            `json:"allow_variants_without_examples"`
	SearchTerms                  []string        `json:"search_terms"`
	IsFilter                     bool            `json:"is_filter"`
	CreatesScope                 bool            `json:"creates_scope"`
	AllowsUnknownArgs            bool            `json:"allows_unknown_args"`
	Category                     string          `json:"category"`
}

type PositionalArg struct {
	Name  string      `json:"name"`
	Desc  string      `json:"desc"`
	Shape string      `json:"shape"`
	VarID interface{} `json:"var_id"`
}

type NamedArg struct {
	Long     string      `json:"long"`
	Short    string      `json:"short"`
	Arg      interface{} `json:"arg"`
	Required bool        `json:"required"`
	Desc     string      `json:"desc"`
	VarID    interface{} `json:"var_id"`
}

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
