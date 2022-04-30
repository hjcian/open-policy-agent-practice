sample-run-disallow:
	opa eval --data opa_rules/official_example.rego --input opa_rules/disallow_example.json --format pretty "data.official_example.allow"

sample-run-allow:
	opa eval --data opa_rules/official_example.rego --input opa_rules/allow_example.json --format pretty "data.official_example.allow"

test-rego:
	opa test -v opa_rules/*.rego

run-basic:
	go run basic_example/main.go