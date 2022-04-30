package main

import (
	"context"
	"fmt"

	"github.com/hjcian/open-policy-agent-practice/opa_rules"
	"github.com/open-policy-agent/opa/rego"
)

var defaultQuery = "data.official_example.allow"

type input struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Object string `json:"object"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	q, err := rego.New(
		rego.Query(defaultQuery),
		rego.Module("official_example", string(opa_rules.ReadPolicy())),
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
