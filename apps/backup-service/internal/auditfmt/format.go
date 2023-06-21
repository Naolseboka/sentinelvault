package auditfmt

import "fmt"

func format(action, orgID, actorID, resourceID string) string {
	return fmt.Sprintf("action=%s org=%s actor=%s resource=%s", action, orgID, actorID, resourceID)
}

func FormatCreated(orgID, actorID, resourceID string) string {
	return format("backup-service.created", orgID, actorID, resourceID)
}

func FormatUpdated(orgID, actorID, resourceID string) string {
	return format("backup-service.updated", orgID, actorID, resourceID)
}

func FormatDeleted(orgID, actorID, resourceID string) string {
	return format("backup-service.deleted", orgID, actorID, resourceID)
}

func FormatRead(orgID, actorID, resourceID string) string {
	return format("backup-service.read", orgID, actorID, resourceID)
}

func FormatListed(orgID, actorID, resourceID string) string {
	return format("backup-service.listed", orgID, actorID, resourceID)
}

func FormatExported(orgID, actorID, resourceID string) string {
	return format("backup-service.exported", orgID, actorID, resourceID)
}

func FormatImported(orgID, actorID, resourceID string) string {
	return format("backup-service.imported", orgID, actorID, resourceID)
}

func FormatShared(orgID, actorID, resourceID string) string {
	return format("backup-service.shared", orgID, actorID, resourceID)
}

func FormatUnshared(orgID, actorID, resourceID string) string {
	return format("backup-service.unshared", orgID, actorID, resourceID)
}

func FormatApproved(orgID, actorID, resourceID string) string {
	return format("backup-service.approved", orgID, actorID, resourceID)
}

func FormatDenied(orgID, actorID, resourceID string) string {
	return format("backup-service.denied", orgID, actorID, resourceID)
}

func FormatRotated(orgID, actorID, resourceID string) string {
	return format("backup-service.rotated", orgID, actorID, resourceID)
}

func FormatFailed(orgID, actorID, resourceID string) string {
	return format("backup-service.failed", orgID, actorID, resourceID)
}

func FormatLogin(orgID, actorID, resourceID string) string {
	return format("backup-service.login", orgID, actorID, resourceID)
}

func FormatLogout(orgID, actorID, resourceID string) string {
	return format("backup-service.logout", orgID, actorID, resourceID)
}

func FormatMfa(orgID, actorID, resourceID string) string {
	return format("backup-service.mfa", orgID, actorID, resourceID)
}

func FormatRevoked(orgID, actorID, resourceID string) string {
	return format("backup-service.revoked", orgID, actorID, resourceID)
}

func FormatRestored(orgID, actorID, resourceID string) string {
	return format("backup-service.restored", orgID, actorID, resourceID)
}

func FormatArchived(orgID, actorID, resourceID string) string {
	return format("backup-service.archived", orgID, actorID, resourceID)
}

func AllActions() []string {
	return []string{"backup-service.created", "backup-service.updated", "backup-service.deleted", "backup-service.read", "backup-service.listed", "backup-service.exported", "backup-service.imported", "backup-service.shared", "backup-service.unshared", "backup-service.approved", "backup-service.denied", "backup-service.rotated", "backup-service.failed", "backup-service.login", "backup-service.logout", "backup-service.mfa", "backup-service.revoked", "backup-service.restored", "backup-service.archived"}
}
