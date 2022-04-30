package external_data_example

# By default, deny requests.
default allow = false

allow {
	# lookup the list of roles for the user
	roles := data.user_roles[input.user]

	# for each role in that list
	r := roles[_]

	# lookup the permissions list for role r
	permissions := data.role_permissions[r]

	# for each permission
	p := permissions[_]

	# check if the permission granted to r matches the user's request
	p == {"action": input.action, "object": input.object}
}
