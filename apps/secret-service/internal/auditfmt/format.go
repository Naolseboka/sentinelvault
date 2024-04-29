package auditfmt

import "fmt"

func format(action, orgID, actorID, resourceID string) string {
	return fmt.Sprintf("action=%s org=%s actor=%s resource=%s", action, orgID, actorID, resourceID)
}

func FormatCreated(orgID, actorID, resourceID string) string {
	return format("secret-service.created", orgID, actorID, resourceID)
}

func FormatUpdated(orgID, actorID, resourceID string) string {
	return format("secret-service.updated", orgID, actorID, resourceID)
}

func FormatDeleted(orgID, actorID, resourceID string) string {
	return format("secret-service.deleted", orgID, actorID, resourceID)
}

func FormatRead(orgID, actorID, resourceID string) string {
	return format("secret-service.read", orgID, actorID, resourceID)
}

func FormatListed(orgID, actorID, resourceID string) string {
	return format("secret-service.listed", orgID, actorID, resourceID)
}

func FormatExported(orgID, actorID, resourceID string) string {
	return format("secret-service.exported", orgID, actorID, resourceID)
}

func FormatImported(orgID, actorID, resourceID string) string {
	return format("secret-service.imported", orgID, actorID, resourceID)
}

func FormatShared(orgID, actorID, resourceID string) string {
	return format("secret-service.shared", orgID, actorID, resourceID)
}

func FormatUnshared(orgID, actorID, resourceID string) string {
	return format("secret-service.unshared", orgID, actorID, resourceID)
}

func FormatApproved(orgID, actorID, resourceID string) string {
	return format("secret-service.approved", orgID, actorID, resourceID)
}

func FormatDenied(orgID, actorID, resourceID string) string {
	return format("secret-service.denied", orgID, actorID, resourceID)
}

func FormatRotated(orgID, actorID, resourceID string) string {
	return format("secret-service.rotated", orgID, actorID, resourceID)
}

func FormatFailed(orgID, actorID, resourceID string) string {
	return format("secret-service.failed", orgID, actorID, resourceID)
}

func FormatLogin(orgID, actorID, resourceID string) string {
	return format("secret-service.login", orgID, actorID, resourceID)
}

func FormatLogout(orgID, actorID, resourceID string) string {
	return format("secret-service.logout", orgID, actorID, resourceID)
}

func FormatMfa(orgID, actorID, resourceID string) string {
	return format("secret-service.mfa", orgID, actorID, resourceID)
}

func FormatRevoked(orgID, actorID, resourceID string) string {
	return format("secret-service.revoked", orgID, actorID, resourceID)
}

func FormatRestored(orgID, actorID, resourceID string) string {
	return format("secret-service.restored", orgID, actorID, resourceID)
}

func FormatArchived(orgID, actorID, resourceID string) string {
	return format("secret-service.archived", orgID, actorID, resourceID)
}

func AllActions() []string {
	return []string{"secret-service.created", "secret-service.updated", "secret-service.deleted", "secret-service.read", "secret-service.listed", "secret-service.exported", "secret-service.imported", "secret-service.shared", "secret-service.unshared", "secret-service.approved", "secret-service.denied", "secret-service.rotated", "secret-service.failed", "secret-service.login", "secret-service.logout", "secret-service.mfa", "secret-service.revoked", "secret-service.restored", "secret-service.archived"}
}
