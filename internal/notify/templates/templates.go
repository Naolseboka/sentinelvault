package templates

import (
	"fmt"
	"strings"
)

func formatTemplate(channel, event, orgName, actor, resource string) string {
	base := fmt.Sprintf("[SentinelVault/%s] event=%s org=%s actor=%s resource=%s", strings.ToUpper(channel), event, orgName, actor, resource)
	switch channel {
	case "email":
		return "Subject: Security notification\n\n" + base
	case "slack", "teams", "discord":
		return "{" + fmt.Sprintf(`"text":%q`, base) + "}"
	case "sms":
		if len(base) > 140 {
			return base[:137] + "..."
		}
		return base
	default:
		return base
	}
}

func AllEvents() []string {
	return []string{"login_success", "login_failure", "mfa_enrolled", "mfa_challenged", "session_revoked", "vault_item_created", "vault_item_updated", "vault_item_deleted", "vault_item_shared", "vault_item_exported", "secret_rotated", "rotation_failed", "approval_requested", "approval_granted", "approval_denied", "emergency_access", "delegation_created", "risk_high", "threat_detected", "compliance_gap", "backup_completed", "backup_failed", "license_expiring", "seat_limit", "policy_violation", "device_untrusted", "impossible_travel"}
}

func AllChannels() []string {
	return []string{"email", "slack", "teams", "discord", "sms", "webhook", "pagerduty", "opsgenie"}
}

func Template_login_success_email(orgName, actor, resource string) string {
	return formatTemplate("email", "login_success", orgName, actor, resource)
}
func Template_login_success_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "login_success", orgName, actor, resource)
}
func Template_login_success_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "login_success", orgName, actor, resource)
}
func Template_login_success_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "login_success", orgName, actor, resource)
}
func Template_login_success_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "login_success", orgName, actor, resource)
}
func Template_login_success_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "login_success", orgName, actor, resource)
}
func Template_login_success_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "login_success", orgName, actor, resource)
}
func Template_login_success_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "login_success", orgName, actor, resource)
}
func Template_login_failure_email(orgName, actor, resource string) string {
	return formatTemplate("email", "login_failure", orgName, actor, resource)
}
func Template_login_failure_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "login_failure", orgName, actor, resource)
}
func Template_login_failure_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "login_failure", orgName, actor, resource)
}
func Template_login_failure_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "login_failure", orgName, actor, resource)
}
func Template_login_failure_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "login_failure", orgName, actor, resource)
}
func Template_login_failure_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "login_failure", orgName, actor, resource)
}
func Template_login_failure_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "login_failure", orgName, actor, resource)
}
func Template_login_failure_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "login_failure", orgName, actor, resource)
}
func Template_mfa_enrolled_email(orgName, actor, resource string) string {
	return formatTemplate("email", "mfa_enrolled", orgName, actor, resource)
}
func Template_mfa_enrolled_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "mfa_enrolled", orgName, actor, resource)
}
func Template_mfa_enrolled_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "mfa_enrolled", orgName, actor, resource)
}
func Template_mfa_enrolled_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "mfa_enrolled", orgName, actor, resource)
}
func Template_mfa_enrolled_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "mfa_enrolled", orgName, actor, resource)
}
func Template_mfa_enrolled_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "mfa_enrolled", orgName, actor, resource)
}
func Template_mfa_enrolled_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "mfa_enrolled", orgName, actor, resource)
}
func Template_mfa_enrolled_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "mfa_enrolled", orgName, actor, resource)
}
func Template_mfa_challenged_email(orgName, actor, resource string) string {
	return formatTemplate("email", "mfa_challenged", orgName, actor, resource)
}
func Template_mfa_challenged_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "mfa_challenged", orgName, actor, resource)
}
func Template_mfa_challenged_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "mfa_challenged", orgName, actor, resource)
}
func Template_mfa_challenged_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "mfa_challenged", orgName, actor, resource)
}
func Template_mfa_challenged_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "mfa_challenged", orgName, actor, resource)
}
func Template_mfa_challenged_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "mfa_challenged", orgName, actor, resource)
}
func Template_mfa_challenged_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "mfa_challenged", orgName, actor, resource)
}
func Template_mfa_challenged_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "mfa_challenged", orgName, actor, resource)
}
func Template_session_revoked_email(orgName, actor, resource string) string {
	return formatTemplate("email", "session_revoked", orgName, actor, resource)
}
func Template_session_revoked_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "session_revoked", orgName, actor, resource)
}
func Template_session_revoked_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "session_revoked", orgName, actor, resource)
}
func Template_session_revoked_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "session_revoked", orgName, actor, resource)
}
func Template_session_revoked_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "session_revoked", orgName, actor, resource)
}
func Template_session_revoked_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "session_revoked", orgName, actor, resource)
}
func Template_session_revoked_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "session_revoked", orgName, actor, resource)
}
func Template_session_revoked_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "session_revoked", orgName, actor, resource)
}
func Template_vault_item_created_email(orgName, actor, resource string) string {
	return formatTemplate("email", "vault_item_created", orgName, actor, resource)
}
func Template_vault_item_created_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "vault_item_created", orgName, actor, resource)
}
func Template_vault_item_created_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "vault_item_created", orgName, actor, resource)
}
func Template_vault_item_created_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "vault_item_created", orgName, actor, resource)
}
func Template_vault_item_created_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "vault_item_created", orgName, actor, resource)
}
func Template_vault_item_created_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "vault_item_created", orgName, actor, resource)
}
func Template_vault_item_created_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "vault_item_created", orgName, actor, resource)
}
func Template_vault_item_created_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "vault_item_created", orgName, actor, resource)
}
func Template_vault_item_updated_email(orgName, actor, resource string) string {
	return formatTemplate("email", "vault_item_updated", orgName, actor, resource)
}
func Template_vault_item_updated_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "vault_item_updated", orgName, actor, resource)
}
func Template_vault_item_updated_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "vault_item_updated", orgName, actor, resource)
}
func Template_vault_item_updated_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "vault_item_updated", orgName, actor, resource)
}
func Template_vault_item_updated_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "vault_item_updated", orgName, actor, resource)
}
func Template_vault_item_updated_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "vault_item_updated", orgName, actor, resource)
}
func Template_vault_item_updated_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "vault_item_updated", orgName, actor, resource)
}
func Template_vault_item_updated_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "vault_item_updated", orgName, actor, resource)
}
func Template_vault_item_deleted_email(orgName, actor, resource string) string {
	return formatTemplate("email", "vault_item_deleted", orgName, actor, resource)
}
func Template_vault_item_deleted_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "vault_item_deleted", orgName, actor, resource)
}
func Template_vault_item_deleted_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "vault_item_deleted", orgName, actor, resource)
}
func Template_vault_item_deleted_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "vault_item_deleted", orgName, actor, resource)
}
func Template_vault_item_deleted_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "vault_item_deleted", orgName, actor, resource)
}
func Template_vault_item_deleted_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "vault_item_deleted", orgName, actor, resource)
}
func Template_vault_item_deleted_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "vault_item_deleted", orgName, actor, resource)
}
func Template_vault_item_deleted_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "vault_item_deleted", orgName, actor, resource)
}
func Template_vault_item_shared_email(orgName, actor, resource string) string {
	return formatTemplate("email", "vault_item_shared", orgName, actor, resource)
}
func Template_vault_item_shared_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "vault_item_shared", orgName, actor, resource)
}
func Template_vault_item_shared_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "vault_item_shared", orgName, actor, resource)
}
func Template_vault_item_shared_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "vault_item_shared", orgName, actor, resource)
}
func Template_vault_item_shared_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "vault_item_shared", orgName, actor, resource)
}
func Template_vault_item_shared_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "vault_item_shared", orgName, actor, resource)
}
func Template_vault_item_shared_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "vault_item_shared", orgName, actor, resource)
}
func Template_vault_item_shared_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "vault_item_shared", orgName, actor, resource)
}
func Template_vault_item_exported_email(orgName, actor, resource string) string {
	return formatTemplate("email", "vault_item_exported", orgName, actor, resource)
}
func Template_vault_item_exported_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "vault_item_exported", orgName, actor, resource)
}
func Template_vault_item_exported_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "vault_item_exported", orgName, actor, resource)
}
func Template_vault_item_exported_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "vault_item_exported", orgName, actor, resource)
}
func Template_vault_item_exported_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "vault_item_exported", orgName, actor, resource)
}
func Template_vault_item_exported_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "vault_item_exported", orgName, actor, resource)
}
func Template_vault_item_exported_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "vault_item_exported", orgName, actor, resource)
}
func Template_vault_item_exported_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "vault_item_exported", orgName, actor, resource)
}
func Template_secret_rotated_email(orgName, actor, resource string) string {
	return formatTemplate("email", "secret_rotated", orgName, actor, resource)
}
func Template_secret_rotated_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "secret_rotated", orgName, actor, resource)
}
func Template_secret_rotated_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "secret_rotated", orgName, actor, resource)
}
func Template_secret_rotated_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "secret_rotated", orgName, actor, resource)
}
func Template_secret_rotated_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "secret_rotated", orgName, actor, resource)
}
func Template_secret_rotated_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "secret_rotated", orgName, actor, resource)
}
func Template_secret_rotated_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "secret_rotated", orgName, actor, resource)
}
func Template_secret_rotated_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "secret_rotated", orgName, actor, resource)
}
func Template_rotation_failed_email(orgName, actor, resource string) string {
	return formatTemplate("email", "rotation_failed", orgName, actor, resource)
}
func Template_rotation_failed_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "rotation_failed", orgName, actor, resource)
}
func Template_rotation_failed_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "rotation_failed", orgName, actor, resource)
}
func Template_rotation_failed_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "rotation_failed", orgName, actor, resource)
}
func Template_rotation_failed_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "rotation_failed", orgName, actor, resource)
}
func Template_rotation_failed_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "rotation_failed", orgName, actor, resource)
}
func Template_rotation_failed_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "rotation_failed", orgName, actor, resource)
}
func Template_rotation_failed_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "rotation_failed", orgName, actor, resource)
}
func Template_approval_requested_email(orgName, actor, resource string) string {
	return formatTemplate("email", "approval_requested", orgName, actor, resource)
}
func Template_approval_requested_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "approval_requested", orgName, actor, resource)
}
func Template_approval_requested_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "approval_requested", orgName, actor, resource)
}
func Template_approval_requested_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "approval_requested", orgName, actor, resource)
}
func Template_approval_requested_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "approval_requested", orgName, actor, resource)
}
func Template_approval_requested_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "approval_requested", orgName, actor, resource)
}
func Template_approval_requested_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "approval_requested", orgName, actor, resource)
}
func Template_approval_requested_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "approval_requested", orgName, actor, resource)
}
func Template_approval_granted_email(orgName, actor, resource string) string {
	return formatTemplate("email", "approval_granted", orgName, actor, resource)
}
func Template_approval_granted_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "approval_granted", orgName, actor, resource)
}
func Template_approval_granted_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "approval_granted", orgName, actor, resource)
}
func Template_approval_granted_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "approval_granted", orgName, actor, resource)
}
func Template_approval_granted_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "approval_granted", orgName, actor, resource)
}
func Template_approval_granted_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "approval_granted", orgName, actor, resource)
}
func Template_approval_granted_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "approval_granted", orgName, actor, resource)
}
func Template_approval_granted_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "approval_granted", orgName, actor, resource)
}
func Template_approval_denied_email(orgName, actor, resource string) string {
	return formatTemplate("email", "approval_denied", orgName, actor, resource)
}
func Template_approval_denied_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "approval_denied", orgName, actor, resource)
}
func Template_approval_denied_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "approval_denied", orgName, actor, resource)
}
func Template_approval_denied_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "approval_denied", orgName, actor, resource)
}
func Template_approval_denied_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "approval_denied", orgName, actor, resource)
}
func Template_approval_denied_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "approval_denied", orgName, actor, resource)
}
func Template_approval_denied_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "approval_denied", orgName, actor, resource)
}
func Template_approval_denied_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "approval_denied", orgName, actor, resource)
}
func Template_emergency_access_email(orgName, actor, resource string) string {
	return formatTemplate("email", "emergency_access", orgName, actor, resource)
}
func Template_emergency_access_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "emergency_access", orgName, actor, resource)
}
func Template_emergency_access_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "emergency_access", orgName, actor, resource)
}
func Template_emergency_access_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "emergency_access", orgName, actor, resource)
}
func Template_emergency_access_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "emergency_access", orgName, actor, resource)
}
func Template_emergency_access_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "emergency_access", orgName, actor, resource)
}
func Template_emergency_access_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "emergency_access", orgName, actor, resource)
}
func Template_emergency_access_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "emergency_access", orgName, actor, resource)
}
func Template_delegation_created_email(orgName, actor, resource string) string {
	return formatTemplate("email", "delegation_created", orgName, actor, resource)
}
func Template_delegation_created_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "delegation_created", orgName, actor, resource)
}
func Template_delegation_created_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "delegation_created", orgName, actor, resource)
}
func Template_delegation_created_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "delegation_created", orgName, actor, resource)
}
func Template_delegation_created_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "delegation_created", orgName, actor, resource)
}
func Template_delegation_created_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "delegation_created", orgName, actor, resource)
}
func Template_delegation_created_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "delegation_created", orgName, actor, resource)
}
func Template_delegation_created_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "delegation_created", orgName, actor, resource)
}
func Template_risk_high_email(orgName, actor, resource string) string {
	return formatTemplate("email", "risk_high", orgName, actor, resource)
}
func Template_risk_high_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "risk_high", orgName, actor, resource)
}
func Template_risk_high_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "risk_high", orgName, actor, resource)
}
func Template_risk_high_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "risk_high", orgName, actor, resource)
}
func Template_risk_high_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "risk_high", orgName, actor, resource)
}
func Template_risk_high_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "risk_high", orgName, actor, resource)
}
func Template_risk_high_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "risk_high", orgName, actor, resource)
}
func Template_risk_high_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "risk_high", orgName, actor, resource)
}
func Template_threat_detected_email(orgName, actor, resource string) string {
	return formatTemplate("email", "threat_detected", orgName, actor, resource)
}
func Template_threat_detected_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "threat_detected", orgName, actor, resource)
}
func Template_threat_detected_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "threat_detected", orgName, actor, resource)
}
func Template_threat_detected_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "threat_detected", orgName, actor, resource)
}
func Template_threat_detected_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "threat_detected", orgName, actor, resource)
}
func Template_threat_detected_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "threat_detected", orgName, actor, resource)
}
func Template_threat_detected_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "threat_detected", orgName, actor, resource)
}
func Template_threat_detected_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "threat_detected", orgName, actor, resource)
}
func Template_compliance_gap_email(orgName, actor, resource string) string {
	return formatTemplate("email", "compliance_gap", orgName, actor, resource)
}
func Template_compliance_gap_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "compliance_gap", orgName, actor, resource)
}
func Template_compliance_gap_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "compliance_gap", orgName, actor, resource)
}
func Template_compliance_gap_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "compliance_gap", orgName, actor, resource)
}
func Template_compliance_gap_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "compliance_gap", orgName, actor, resource)
}
func Template_compliance_gap_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "compliance_gap", orgName, actor, resource)
}
func Template_compliance_gap_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "compliance_gap", orgName, actor, resource)
}
func Template_compliance_gap_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "compliance_gap", orgName, actor, resource)
}
func Template_backup_completed_email(orgName, actor, resource string) string {
	return formatTemplate("email", "backup_completed", orgName, actor, resource)
}
func Template_backup_completed_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "backup_completed", orgName, actor, resource)
}
func Template_backup_completed_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "backup_completed", orgName, actor, resource)
}
func Template_backup_completed_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "backup_completed", orgName, actor, resource)
}
func Template_backup_completed_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "backup_completed", orgName, actor, resource)
}
func Template_backup_completed_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "backup_completed", orgName, actor, resource)
}
func Template_backup_completed_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "backup_completed", orgName, actor, resource)
}
func Template_backup_completed_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "backup_completed", orgName, actor, resource)
}
func Template_backup_failed_email(orgName, actor, resource string) string {
	return formatTemplate("email", "backup_failed", orgName, actor, resource)
}
func Template_backup_failed_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "backup_failed", orgName, actor, resource)
}
func Template_backup_failed_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "backup_failed", orgName, actor, resource)
}
func Template_backup_failed_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "backup_failed", orgName, actor, resource)
}
func Template_backup_failed_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "backup_failed", orgName, actor, resource)
}
func Template_backup_failed_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "backup_failed", orgName, actor, resource)
}
func Template_backup_failed_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "backup_failed", orgName, actor, resource)
}
func Template_backup_failed_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "backup_failed", orgName, actor, resource)
}
func Template_license_expiring_email(orgName, actor, resource string) string {
	return formatTemplate("email", "license_expiring", orgName, actor, resource)
}
func Template_license_expiring_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "license_expiring", orgName, actor, resource)
}
func Template_license_expiring_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "license_expiring", orgName, actor, resource)
}
func Template_license_expiring_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "license_expiring", orgName, actor, resource)
}
func Template_license_expiring_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "license_expiring", orgName, actor, resource)
}
func Template_license_expiring_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "license_expiring", orgName, actor, resource)
}
func Template_license_expiring_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "license_expiring", orgName, actor, resource)
}
func Template_license_expiring_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "license_expiring", orgName, actor, resource)
}
func Template_seat_limit_email(orgName, actor, resource string) string {
	return formatTemplate("email", "seat_limit", orgName, actor, resource)
}
func Template_seat_limit_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "seat_limit", orgName, actor, resource)
}
func Template_seat_limit_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "seat_limit", orgName, actor, resource)
}
func Template_seat_limit_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "seat_limit", orgName, actor, resource)
}
func Template_seat_limit_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "seat_limit", orgName, actor, resource)
}
func Template_seat_limit_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "seat_limit", orgName, actor, resource)
}
func Template_seat_limit_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "seat_limit", orgName, actor, resource)
}
func Template_seat_limit_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "seat_limit", orgName, actor, resource)
}
func Template_policy_violation_email(orgName, actor, resource string) string {
	return formatTemplate("email", "policy_violation", orgName, actor, resource)
}
func Template_policy_violation_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "policy_violation", orgName, actor, resource)
}
func Template_policy_violation_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "policy_violation", orgName, actor, resource)
}
func Template_policy_violation_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "policy_violation", orgName, actor, resource)
}
func Template_policy_violation_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "policy_violation", orgName, actor, resource)
}
func Template_policy_violation_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "policy_violation", orgName, actor, resource)
}
func Template_policy_violation_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "policy_violation", orgName, actor, resource)
}
func Template_policy_violation_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "policy_violation", orgName, actor, resource)
}
func Template_device_untrusted_email(orgName, actor, resource string) string {
	return formatTemplate("email", "device_untrusted", orgName, actor, resource)
}
func Template_device_untrusted_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "device_untrusted", orgName, actor, resource)
}
func Template_device_untrusted_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "device_untrusted", orgName, actor, resource)
}
func Template_device_untrusted_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "device_untrusted", orgName, actor, resource)
}
func Template_device_untrusted_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "device_untrusted", orgName, actor, resource)
}
func Template_device_untrusted_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "device_untrusted", orgName, actor, resource)
}
func Template_device_untrusted_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "device_untrusted", orgName, actor, resource)
}
func Template_device_untrusted_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "device_untrusted", orgName, actor, resource)
}
func Template_impossible_travel_email(orgName, actor, resource string) string {
	return formatTemplate("email", "impossible_travel", orgName, actor, resource)
}
func Template_impossible_travel_slack(orgName, actor, resource string) string {
	return formatTemplate("slack", "impossible_travel", orgName, actor, resource)
}
func Template_impossible_travel_teams(orgName, actor, resource string) string {
	return formatTemplate("teams", "impossible_travel", orgName, actor, resource)
}
func Template_impossible_travel_discord(orgName, actor, resource string) string {
	return formatTemplate("discord", "impossible_travel", orgName, actor, resource)
}
func Template_impossible_travel_sms(orgName, actor, resource string) string {
	return formatTemplate("sms", "impossible_travel", orgName, actor, resource)
}
func Template_impossible_travel_webhook(orgName, actor, resource string) string {
	return formatTemplate("webhook", "impossible_travel", orgName, actor, resource)
}
func Template_impossible_travel_pagerduty(orgName, actor, resource string) string {
	return formatTemplate("pagerduty", "impossible_travel", orgName, actor, resource)
}
func Template_impossible_travel_opsgenie(orgName, actor, resource string) string {
	return formatTemplate("opsgenie", "impossible_travel", orgName, actor, resource)
}
