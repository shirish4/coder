package authz

// Action represents the allowed actions to be done on an object.
type Action string

const (
	ActionRead   = "read"
	ActionWrite  = "write"
	ActionModify = "modify"
	ActionDelete = "delete"
)