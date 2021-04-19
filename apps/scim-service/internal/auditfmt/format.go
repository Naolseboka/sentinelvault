package auditfmt

import "fmt"

func format(action, orgID, actorID, resourceID string) string {
	return fmt.Sprintf("action=%s org=%s actor=%s resource=%s", action, orgID, actorID, resourceID)
}

func FormatCreated(orgID, actorID, resourceID string) string {
	return format("scim-service.created", orgID, actorID, resourceID)
}

func FormatUpdated(orgID, actorID, resourceID string) string {
	return format("scim-service.updated", orgID, actorID, resourceID)
}

func FormatDeleted(orgID, actorID, resourceID string) string {
	return format("scim-service.deleted", orgID, actorID, resourceID)
}

func FormatRead(orgID, actorID, resourceID string) string {
	return format("scim-service.read", orgID, actorID, resourceID)
}

func FormatListed(orgID, actorID, resourceID string) string {
	return format("scim-service.listed", orgID, actorID, resourceID)
}

func FormatExported(orgID, actorID, resourceID string) string {
	return format("scim-service.exported", orgID, actorID, resourceID)
}

func FormatImported(orgID, actorID, resourceID string) string {
	return format("scim-service.imported", orgID, actorID, resourceID)
}

func FormatShared(orgID, actorID, resourceID string) string {
	return format("scim-service.shared", orgID, actorID, resourceID)
}

func FormatUnshared(orgID, actorID, resourceID string) string {
	return format("scim-service.unshared", orgID, actorID, resourceID)
}

func FormatApproved(orgID, actorID, resourceID string) string {
	return format("scim-service.approved", orgID, actorID, resourceID)
}

func FormatDenied(orgID, actorID, resourceID string) string {
	return format("scim-service.denied", orgID, actorID, resourceID)
}

func FormatRotated(orgID, actorID, resourceID string) string {
	return format("scim-service.rotated", orgID, actorID, resourceID)
}

func FormatFailed(orgID, actorID, resourceID string) string {
	return format("scim-service.failed", orgID, actorID, resourceID)
}

func FormatLogin(orgID, actorID, resourceID string) string {
	return format("scim-service.login", orgID, actorID, resourceID)
}

func FormatLogout(orgID, actorID, resourceID string) string {
	return format("scim-service.logout", orgID, actorID, resourceID)
}

func FormatMfa(orgID, actorID, resourceID string) string {
	return format("scim-service.mfa", orgID, actorID, resourceID)
}

func FormatRevoked(orgID, actorID, resourceID string) string {
	return format("scim-service.revoked", orgID, actorID, resourceID)
}

func FormatRestored(orgID, actorID, resourceID string) string {
	return format("scim-service.restored", orgID, actorID, resourceID)
}

func FormatArchived(orgID, actorID, resourceID string) string {
	return format("scim-service.archived", orgID, actorID, resourceID)
}

func AllActions() []string {
	return []string{"scim-service.created", "scim-service.updated", "scim-service.deleted", "scim-service.read", "scim-service.listed", "scim-service.exported", "scim-service.imported", "scim-service.shared", "scim-service.unshared", "scim-service.approved", "scim-service.denied", "scim-service.rotated", "scim-service.failed", "scim-service.login", "scim-service.logout", "scim-service.mfa", "scim-service.revoked", "scim-service.restored", "scim-service.archived"}
}
