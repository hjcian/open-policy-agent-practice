package main

import (
	"context"
	"fmt"

	"github.com/hjcian/open-policy-agent-practice/opa_rules"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/open-policy-agent/opa/util"
)

const user_roles = ` {
	"alice": ["engineering", "webdev"],
	"bob": ["hr"]
	}`

const role_permissions = `{
		"engineering": [{"action": "read", "object": "server123"}],
		"webdev": [
			{"action": "read", "object": "server123"},
			{"action": "write", "object": "server123"}
			],
			"hr": [{"action": "read", "object": "database456"}]
			}`

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type input struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Object string `json:"object"`
}

func main() {
	// https://stackoverflow.com/questions/69491963/opa-rego-as-go-lib-how-to-apply-external-data
	// https://pkg.go.dev/github.com/open-policy-agent/opa@v0.33.1/rego#example-Rego.Eval-Storage
	policy := opa_rules.ReadExternalDataExamplePolicy()

	var dataJson map[string]interface{}
	must(util.UnmarshalJSON(
		[]byte(fmt.Sprintf(
			`{
			"user_roles": %[1]s,
			"role_permissions": %[2]s
		}`, user_roles, role_permissions),
		), &dataJson))

	q, err := rego.New(
		rego.Query("data.external_data_example.allow"),
		rego.Module("official_example", string(policy)),
		rego.Store(inmem.NewFromObject(dataJson)),
	).PrepareForEval(context.Background())
	must(err)

	allowedInput := input{
		User:   "bob",
		Action: "read",
		Object: "database456",
	}

	res, err := q.Eval(context.Background(), rego.EvalInput(allowedInput))
	must(err)
	fmt.Println("should be true:", res.Allowed())

	disallowedInput := input{
		User:   "bob",
		Action: "read",
		Object: "server123",
	}
	res, err = q.Eval(context.Background(), rego.EvalInput(disallowedInput))
	must(err)
	fmt.Println("should be false:", res.Allowed())
}
