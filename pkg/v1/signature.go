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
