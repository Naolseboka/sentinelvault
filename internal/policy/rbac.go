package policy

import "strings"

// Permission describes an allowed action on a resource type.
type Permission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

// Role groups permissions granted to a subject.
type Role struct {
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

// HasPermission reports whether role grants action on resource.
func HasPermission(role Role, resource, action string) bool {
	for _, p := range role.Permissions {
		if matchPattern(p.Resource, resource) && matchPattern(p.Action, action) {
			return true
		}
	}
	return false
}

// HasAnyRolePermission checks permissions across multiple roles.
func HasAnyRolePermission(roles []Role, resource, action string) bool {
	for _, r := range roles {
		if HasPermission(r, resource, action) {
			return true
		}
	}
	return false
}

func matchPattern(pattern, value string) bool {
	if pattern == "*" || pattern == value {
		return true
	}
	if strings.HasSuffix(pattern, ":*") {
		prefix := strings.TrimSuffix(pattern, ":*")
		return strings.HasPrefix(value, prefix)
	}
	return false
}

// BuiltinRoles returns default platform roles.
func BuiltinRoles() map[string]Role {
	return map[string]Role{
		"admin": {
			Name: "admin",
			Permissions: []Permission{
				{Resource: "*", Action: "*"},
			},
		},
		"member": {
			Name: "member",
			Permissions: []Permission{
				{Resource: "secret", Action: "read"},
				{Resource: "secret", Action: "write"},
				{Resource: "user", Action: "read"},
			},
		},
		"viewer": {
			Name: "viewer",
			Permissions: []Permission{
				{Resource: "secret", Action: "read"},
				{Resource: "user", Action: "read"},
			},
		},
	}
}
