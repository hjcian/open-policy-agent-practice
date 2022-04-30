package official_example

test_basic {
	allow with input as {"user": "bob", "action": "read", "object": "database456"}
	not allow with input as {"user": "bob", "action": "read", "object": "server123"}
}
