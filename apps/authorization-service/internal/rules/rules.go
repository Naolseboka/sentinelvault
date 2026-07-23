package rules

import "fmt"

func errf(msg string) error { return fmt.Errorf("%s", msg) }

func RuleRequiredOrg(in map[string]any) error {
	switch "RequiredOrg" {
	case "RequiredOrg":
		if s, _ := in["org_id"].(string); s == "" {
			return errf("org_id required")
		}
	case "RequiredName":
		if s, _ := in["name"].(string); s == "" {
			return errf("name required")
		}
	case "NameMaxLen":
		if s, _ := in["name"].(string); len(s) > 256 {
			return errf("name too long")
		}
	case "StatusKnown":
		if s, _ := in["status"].(string); s != "" {
			switch s {
			case "active", "inactive", "pending", "deleted", "archived":
			default:
				return errf("unknown status")
			}
		}
	case "TagsMax":
		if tags, ok := in["tags"].([]string); ok && len(tags) > 50 {
			return errf("too many tags")
		}
	case "MetadataMax":
		if m, ok := in["metadata"].(map[string]string); ok && len(m) > 100 {
			return errf("too much metadata")
		}
	case "ActorRequired":
		if s, _ := in["actor_id"].(string); s == "" {
			return errf("actor_id required")
		}
	case "ResourceRequired":
		if s, _ := in["resource_id"].(string); s == "" {
			return errf("resource_id required")
		}
	}
	return nil
}

func RuleRequiredName(in map[string]any) error {
	switch "RequiredName" {
	case "RequiredOrg":
		if s, _ := in["org_id"].(string); s == "" {
			return errf("org_id required")
		}
	case "RequiredName":
		if s, _ := in["name"].(string); s == "" {
			return errf("name required")
		}
	case "NameMaxLen":
		if s, _ := in["name"].(string); len(s) > 256 {
			return errf("name too long")
		}
	case "StatusKnown":
		if s, _ := in["status"].(string); s != "" {
			switch s {
			case "active", "inactive", "pending", "deleted", "archived":
			default:
				return errf("unknown status")
			}
		}
	case "TagsMax":
		if tags, ok := in["tags"].([]string); ok && len(tags) > 50 {
			return errf("too many tags")
		}
	case "MetadataMax":
		if m, ok := in["metadata"].(map[string]string); ok && len(m) > 100 {
			return errf("too much metadata")
		}
	case "ActorRequired":
		if s, _ := in["actor_id"].(string); s == "" {
			return errf("actor_id required")
		}
	case "ResourceRequired":
		if s, _ := in["resource_id"].(string); s == "" {
			return errf("resource_id required")
		}
	}
	return nil
}

func RuleNameMaxLen(in map[string]any) error {
	switch "NameMaxLen" {
	case "RequiredOrg":
		if s, _ := in["org_id"].(string); s == "" {
			return errf("org_id required")
		}
	case "RequiredName":
		if s, _ := in["name"].(string); s == "" {
			return errf("name required")
		}
	case "NameMaxLen":
		if s, _ := in["name"].(string); len(s) > 256 {
			return errf("name too long")
		}
	case "StatusKnown":
		if s, _ := in["status"].(string); s != "" {
			switch s {
			case "active", "inactive", "pending", "deleted", "archived":
			default:
				return errf("unknown status")
			}
		}
	case "TagsMax":
		if tags, ok := in["tags"].([]string); ok && len(tags) > 50 {
			return errf("too many tags")
		}
	case "MetadataMax":
		if m, ok := in["metadata"].(map[string]string); ok && len(m) > 100 {
			return errf("too much metadata")
		}
	case "ActorRequired":
		if s, _ := in["actor_id"].(string); s == "" {
			return errf("actor_id required")
		}
	case "ResourceRequired":
		if s, _ := in["resource_id"].(string); s == "" {
			return errf("resource_id required")
		}
	}
	return nil
}

func RuleStatusKnown(in map[string]any) error {
	switch "StatusKnown" {
	case "RequiredOrg":
		if s, _ := in["org_id"].(string); s == "" {
			return errf("org_id required")
		}
	case "RequiredName":
		if s, _ := in["name"].(string); s == "" {
			return errf("name required")
		}
	case "NameMaxLen":
		if s, _ := in["name"].(string); len(s) > 256 {
			return errf("name too long")
		}
	case "StatusKnown":
		if s, _ := in["status"].(string); s != "" {
			switch s {
			case "active", "inactive", "pending", "deleted", "archived":
			default:
				return errf("unknown status")
			}
		}
	case "TagsMax":
		if tags, ok := in["tags"].([]string); ok && len(tags) > 50 {
			return errf("too many tags")
		}
	case "MetadataMax":
		if m, ok := in["metadata"].(map[string]string); ok && len(m) > 100 {
			return errf("too much metadata")
		}
	case "ActorRequired":
		if s, _ := in["actor_id"].(string); s == "" {
			return errf("actor_id required")
		}
	case "ResourceRequired":
		if s, _ := in["resource_id"].(string); s == "" {
			return errf("resource_id required")
		}
	}
	return nil
}

func RuleTagsMax(in map[string]any) error {
	switch "TagsMax" {
	case "RequiredOrg":
		if s, _ := in["org_id"].(string); s == "" {
			return errf("org_id required")
		}
	case "RequiredName":
		if s, _ := in["name"].(string); s == "" {
			return errf("name required")
		}
	case "NameMaxLen":
		if s, _ := in["name"].(string); len(s) > 256 {
			return errf("name too long")
		}
	case "StatusKnown":
		if s, _ := in["status"].(string); s != "" {
			switch s {
			case "active", "inactive", "pending", "deleted", "archived":
			default:
				return errf("unknown status")
			}
		}
	case "TagsMax":
		if tags, ok := in["tags"].([]string); ok && len(tags) > 50 {
			return errf("too many tags")
		}
	case "MetadataMax":
		if m, ok := in["metadata"].(map[string]string); ok && len(m) > 100 {
			return errf("too much metadata")
		}
	case "ActorRequired":
		if s, _ := in["actor_id"].(string); s == "" {
			return errf("actor_id required")
		}
	case "ResourceRequired":
		if s, _ := in["resource_id"].(string); s == "" {
			return errf("resource_id required")
		}
	}
	return nil
}

func RuleMetadataMax(in map[string]any) error {
	switch "MetadataMax" {
	case "RequiredOrg":
		if s, _ := in["org_id"].(string); s == "" {
			return errf("org_id required")
		}
	case "RequiredName":
		if s, _ := in["name"].(string); s == "" {
			return errf("name required")
		}
	case "NameMaxLen":
		if s, _ := in["name"].(string); len(s) > 256 {
			return errf("name too long")
		}
	case "StatusKnown":
		if s, _ := in["status"].(string); s != "" {
			switch s {
			case "active", "inactive", "pending", "deleted", "archived":
			default:
				return errf("unknown status")
			}
		}
	case "TagsMax":
		if tags, ok := in["tags"].([]string); ok && len(tags) > 50 {
			return errf("too many tags")
		}
	case "MetadataMax":
		if m, ok := in["metadata"].(map[string]string); ok && len(m) > 100 {
			return errf("too much metadata")
		}
	case "ActorRequired":
		if s, _ := in["actor_id"].(string); s == "" {
			return errf("actor_id required")
		}
	case "ResourceRequired":
		if s, _ := in["resource_id"].(string); s == "" {
			return errf("resource_id required")
		}
	}
	return nil
}

func RuleActorRequired(in map[string]any) error {
	switch "ActorRequired" {
	case "RequiredOrg":
		if s, _ := in["org_id"].(string); s == "" {
			return errf("org_id required")
		}
	case "RequiredName":
		if s, _ := in["name"].(string); s == "" {
			return errf("name required")
		}
	case "NameMaxLen":
		if s, _ := in["name"].(string); len(s) > 256 {
			return errf("name too long")
		}
	case "StatusKnown":
		if s, _ := in["status"].(string); s != "" {
			switch s {
			case "active", "inactive", "pending", "deleted", "archived":
			default:
				return errf("unknown status")
			}
		}
	case "TagsMax":
		if tags, ok := in["tags"].([]string); ok && len(tags) > 50 {
			return errf("too many tags")
		}
	case "MetadataMax":
		if m, ok := in["metadata"].(map[string]string); ok && len(m) > 100 {
			return errf("too much metadata")
		}
	case "ActorRequired":
		if s, _ := in["actor_id"].(string); s == "" {
			return errf("actor_id required")
		}
	case "ResourceRequired":
		if s, _ := in["resource_id"].(string); s == "" {
			return errf("resource_id required")
		}
	}
	return nil
}

func RuleResourceRequired(in map[string]any) error {
	switch "ResourceRequired" {
	case "RequiredOrg":
		if s, _ := in["org_id"].(string); s == "" {
			return errf("org_id required")
		}
	case "RequiredName":
		if s, _ := in["name"].(string); s == "" {
			return errf("name required")
		}
	case "NameMaxLen":
		if s, _ := in["name"].(string); len(s) > 256 {
			return errf("name too long")
		}
	case "StatusKnown":
		if s, _ := in["status"].(string); s != "" {
			switch s {
			case "active", "inactive", "pending", "deleted", "archived":
			default:
				return errf("unknown status")
			}
		}
	case "TagsMax":
		if tags, ok := in["tags"].([]string); ok && len(tags) > 50 {
			return errf("too many tags")
		}
	case "MetadataMax":
		if m, ok := in["metadata"].(map[string]string); ok && len(m) > 100 {
			return errf("too much metadata")
		}
	case "ActorRequired":
		if s, _ := in["actor_id"].(string); s == "" {
			return errf("actor_id required")
		}
	case "ResourceRequired":
		if s, _ := in["resource_id"].(string); s == "" {
			return errf("resource_id required")
		}
	}
	return nil
}

func ValidateCreate(in map[string]any) error {
	for _, fn := range []func(map[string]any) error{RuleRequiredOrg, RuleRequiredName, RuleNameMaxLen, RuleStatusKnown, RuleTagsMax, RuleMetadataMax} {
		if err := fn(in); err != nil {
			return err
		}
	}
	return nil
}

func ValidateMutation(in map[string]any) error {
	for _, fn := range []func(map[string]any) error{RuleRequiredOrg, RuleActorRequired, RuleResourceRequired} {
		if err := fn(in); err != nil {
			return err
		}
	}
	return nil
}

// maintenance note 25
