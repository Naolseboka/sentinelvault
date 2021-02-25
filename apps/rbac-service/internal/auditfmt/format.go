package auditfmt

import "fmt"

func format(action, orgID, actorID, resourceID string) string {
	return fmt.Sprintf("action=%s org=%s actor=%s resource=%s", action, orgID, actorID, resourceID)
}

func FormatCreated(orgID, actorID, resourceID string) string {
	return format("rbac-service.created", orgID, actorID, resourceID)
}

func FormatUpdated(orgID, actorID, resourceID string) string {
	return format("rbac-service.updated", orgID, actorID, resourceID)
}

func FormatDeleted(orgID, actorID, resourceID string) string {
	return format("rbac-service.deleted", orgID, actorID, resourceID)
}

func FormatRead(orgID, actorID, resourceID string) string {
	return format("rbac-service.read", orgID, actorID, resourceID)
}

func FormatListed(orgID, actorID, resourceID string) string {
	return format("rbac-service.listed", orgID, actorID, resourceID)
}

func FormatExported(orgID, actorID, resourceID string) string {
	return format("rbac-service.exported", orgID, actorID, resourceID)
}

func FormatImported(orgID, actorID, resourceID string) string {
	return format("rbac-service.imported", orgID, actorID, resourceID)
}

func FormatShared(orgID, actorID, resourceID string) string {
	return format("rbac-service.shared", orgID, actorID, resourceID)
}

func FormatUnshared(orgID, actorID, resourceID string) string {
	return format("rbac-service.unshared", orgID, actorID, resourceID)
}

func FormatApproved(orgID, actorID, resourceID string) string {
	return format("rbac-service.approved", orgID, actorID, resourceID)
}

func FormatDenied(orgID, actorID, resourceID string) string {
	return format("rbac-service.denied", orgID, actorID, resourceID)
}

func FormatRotated(orgID, actorID, resourceID string) string {
	return format("rbac-service.rotated", orgID, actorID, resourceID)
}

func FormatFailed(orgID, actorID, resourceID string) string {
	return format("rbac-service.failed", orgID, actorID, resourceID)
}

func FormatLogin(orgID, actorID, resourceID string) string {
	return format("rbac-service.login", orgID, actorID, resourceID)
}

func FormatLogout(orgID, actorID, resourceID string) string {
	return format("rbac-service.logout", orgID, actorID, resourceID)
}

func FormatMfa(orgID, actorID, resourceID string) string {
	return format("rbac-service.mfa", orgID, actorID, resourceID)
}

func FormatRevoked(orgID, actorID, resourceID string) string {
	return format("rbac-service.revoked", orgID, actorID, resourceID)
}

func FormatRestored(orgID, actorID, resourceID string) string {
	return format("rbac-service.restored", orgID, actorID, resourceID)
}

func FormatArchived(orgID, actorID, resourceID string) string {
	return format("rbac-service.archived", orgID, actorID, resourceID)
}

func AllActions() []string {
	return []string{"rbac-service.created", "rbac-service.updated", "rbac-service.deleted", "rbac-service.read", "rbac-service.listed", "rbac-service.exported", "rbac-service.imported", "rbac-service.shared", "rbac-service.unshared", "rbac-service.approved", "rbac-service.denied", "rbac-service.rotated", "rbac-service.failed", "rbac-service.login", "rbac-service.logout", "rbac-service.mfa", "rbac-service.revoked", "rbac-service.restored", "rbac-service.archived"}
}
