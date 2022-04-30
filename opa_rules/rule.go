package opa_rules

import _ "embed"

//go:embed official_example.rego
var officialExamplePolicy []byte

func ReadOfficialExamplePolicy() []byte {
	return officialExamplePolicy
}

//go:embed external_data_example.rego
var externalDataExamplePolicy []byte

func ReadExternalDataExamplePolicy() []byte {
	return externalDataExamplePolicy
}
