package opa_rules

import _ "embed"

//go:embed official_example.rego
var policy []byte

func ReadPolicy() []byte {
	return policy
}
