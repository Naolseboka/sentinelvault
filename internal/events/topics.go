package events

// Kafka topic names for domain events.
const (
	TopicUserCreated       = "identity.user.created"
	TopicUserUpdated       = "identity.user.updated"
	TopicOrgCreated        = "identity.org.created"
	TopicSessionCreated    = "identity.session.created"
	TopicVaultItemCreated  = "vault.item.created"
	TopicVaultItemUpdated  = "vault.item.updated"
	TopicSecretRotated     = "vault.secret.rotated"
	TopicPolicyEvaluated   = "security.policy.evaluated"
	TopicAuditRecorded     = "audit.event.recorded"
	TopicNotificationSent  = "ops.notification.sent"
)
