package auditfmt

import "fmt"

func format(action, orgID, actorID, resourceID string) string {
	return fmt.Sprintf("action=%s org=%s actor=%s resource=%s", action, orgID, actorID, resourceID)
}

func FormatCreated(orgID, actorID, resourceID string) string {
	return format("device-service.created", orgID, actorID, resourceID)
}

func FormatUpdated(orgID, actorID, resourceID string) string {
	return format("device-service.updated", orgID, actorID, resourceID)
}

func FormatDeleted(orgID, actorID, resourceID string) string {
	return format("device-service.deleted", orgID, actorID, resourceID)
}

func FormatRead(orgID, actorID, resourceID string) string {
	return format("device-service.read", orgID, actorID, resourceID)
}

func FormatListed(orgID, actorID, resourceID string) string {
	return format("device-service.listed", orgID, actorID, resourceID)
}

func FormatExported(orgID, actorID, resourceID string) string {
	return format("device-service.exported", orgID, actorID, resourceID)
}

func FormatImported(orgID, actorID, resourceID string) string {
	return format("device-service.imported", orgID, actorID, resourceID)
}

func FormatShared(orgID, actorID, resourceID string) string {
	return format("device-service.shared", orgID, actorID, resourceID)
}

func FormatUnshared(orgID, actorID, resourceID string) string {
	return format("device-service.unshared", orgID, actorID, resourceID)
}

func FormatApproved(orgID, actorID, resourceID string) string {
	return format("device-service.approved", orgID, actorID, resourceID)
}

func FormatDenied(orgID, actorID, resourceID string) string {
	return format("device-service.denied", orgID, actorID, resourceID)
}

func FormatRotated(orgID, actorID, resourceID string) string {
	return format("device-service.rotated", orgID, actorID, resourceID)
}

func FormatFailed(orgID, actorID, resourceID string) string {
	return format("device-service.failed", orgID, actorID, resourceID)
}

func FormatLogin(orgID, actorID, resourceID string) string {
	return format("device-service.login", orgID, actorID, resourceID)
}

func FormatLogout(orgID, actorID, resourceID string) string {
	return format("device-service.logout", orgID, actorID, resourceID)
}

func FormatMfa(orgID, actorID, resourceID string) string {
	return format("device-service.mfa", orgID, actorID, resourceID)
}

func FormatRevoked(orgID, actorID, resourceID string) string {
	return format("device-service.revoked", orgID, actorID, resourceID)
}

func FormatRestored(orgID, actorID, resourceID string) string {
	return format("device-service.restored", orgID, actorID, resourceID)
}

func FormatArchived(orgID, actorID, resourceID string) string {
	return format("device-service.archived", orgID, actorID, resourceID)
}

func AllActions() []string {
	return []string{"device-service.created", "device-service.updated", "device-service.deleted", "device-service.read", "device-service.listed", "device-service.exported", "device-service.imported", "device-service.shared", "device-service.unshared", "device-service.approved", "device-service.denied", "device-service.rotated", "device-service.failed", "device-service.login", "device-service.logout", "device-service.mfa", "device-service.revoked", "device-service.restored", "device-service.archived"}
}
