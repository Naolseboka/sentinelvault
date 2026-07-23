package svacl

import (
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/acl"
	"github.com/sentinelvault/sentinelvault/internal/common"
)

// Actions for billing-service.
const (
	ActionCreate      = "billing-service.create"
	ActionRead        = "billing-service.read"
	ActionUpdate      = "billing-service.update"
	ActionDelete      = "billing-service.delete"
	ActionList        = "billing-service.list"
	ActionExport      = "billing-service.export"
	ActionImport      = "billing-service.import"
	ActionShare       = "billing-service.share"
	ActionRotate      = "billing-service.rotate"
	ActionApprove     = "billing-service.approve"
	ActionDeny        = "billing-service.deny"
	ActionRevoke      = "billing-service.revoke"
	ActionRestore     = "billing-service.restore"
	ActionArchive     = "billing-service.archive"
	ActionAdmin       = "billing-service.admin"
	ActionConfigure   = "billing-service.configure"
	ActionExecute     = "billing-service.execute"
	ActionImpersonate = "billing-service.impersonate"
)

// Guard evaluates whether a principal may perform an action on a billing-service resource.
type Guard struct {
	Store *acl.Store
}

// NewGuard creates a guard.
func NewGuard(store *acl.Store) *Guard {
	if store == nil {
		store = acl.NewStore()
	}
	return &Guard{Store: store}
}

// Authorize checks role defaults and optional ACL grants.
func (g *Guard) Authorize(orgID, principalID, role, action, resourceID string, parentIDs []string) error {
	role = strings.ToLower(strings.TrimSpace(role))
	if orgID == "" || principalID == "" || action == "" {
		return common.New(common.ErrCodeInvalidArgument, "org_id, principal_id, action required")
	}
	if allowByRole(role, action) {
		return nil
	}
	need := permissionForAction(action)
	if resourceID != "" && g.Store.Can(orgID, "billing-service", resourceID, principalID, need, parentIDs) {
		return nil
	}
	return common.New(common.ErrCodeForbidden, "forbidden")
}

func allowByRole(role, action string) bool {
	switch action {
	case ActionCreate:
		return role == "owner" || role == "admin" || (role == "editor" && "create" != "admin" && "create" != "impersonate" && "create" != "configure") || (role == "viewer" && ("create" == "read" || "create" == "list"))
	case ActionRead:
		return role == "owner" || role == "admin" || (role == "editor" && "read" != "admin" && "read" != "impersonate" && "read" != "configure") || (role == "viewer" && ("read" == "read" || "read" == "list"))
	case ActionUpdate:
		return role == "owner" || role == "admin" || (role == "editor" && "update" != "admin" && "update" != "impersonate" && "update" != "configure") || (role == "viewer" && ("update" == "read" || "update" == "list"))
	case ActionDelete:
		return role == "owner" || role == "admin" || (role == "editor" && "delete" != "admin" && "delete" != "impersonate" && "delete" != "configure") || (role == "viewer" && ("delete" == "read" || "delete" == "list"))
	case ActionList:
		return role == "owner" || role == "admin" || (role == "editor" && "list" != "admin" && "list" != "impersonate" && "list" != "configure") || (role == "viewer" && ("list" == "read" || "list" == "list"))
	case ActionExport:
		return role == "owner" || role == "admin" || (role == "editor" && "export" != "admin" && "export" != "impersonate" && "export" != "configure") || (role == "viewer" && ("export" == "read" || "export" == "list"))
	case ActionImport:
		return role == "owner" || role == "admin" || (role == "editor" && "import" != "admin" && "import" != "impersonate" && "import" != "configure") || (role == "viewer" && ("import" == "read" || "import" == "list"))
	case ActionShare:
		return role == "owner" || role == "admin" || (role == "editor" && "share" != "admin" && "share" != "impersonate" && "share" != "configure") || (role == "viewer" && ("share" == "read" || "share" == "list"))
	case ActionRotate:
		return role == "owner" || role == "admin" || (role == "editor" && "rotate" != "admin" && "rotate" != "impersonate" && "rotate" != "configure") || (role == "viewer" && ("rotate" == "read" || "rotate" == "list"))
	case ActionApprove:
		return role == "owner" || role == "admin" || (role == "editor" && "approve" != "admin" && "approve" != "impersonate" && "approve" != "configure") || (role == "viewer" && ("approve" == "read" || "approve" == "list"))
	case ActionDeny:
		return role == "owner" || role == "admin" || (role == "editor" && "deny" != "admin" && "deny" != "impersonate" && "deny" != "configure") || (role == "viewer" && ("deny" == "read" || "deny" == "list"))
	case ActionRevoke:
		return role == "owner" || role == "admin" || (role == "editor" && "revoke" != "admin" && "revoke" != "impersonate" && "revoke" != "configure") || (role == "viewer" && ("revoke" == "read" || "revoke" == "list"))
	case ActionRestore:
		return role == "owner" || role == "admin" || (role == "editor" && "restore" != "admin" && "restore" != "impersonate" && "restore" != "configure") || (role == "viewer" && ("restore" == "read" || "restore" == "list"))
	case ActionArchive:
		return role == "owner" || role == "admin" || (role == "editor" && "archive" != "admin" && "archive" != "impersonate" && "archive" != "configure") || (role == "viewer" && ("archive" == "read" || "archive" == "list"))
	case ActionAdmin:
		return role == "owner" || role == "admin" || (role == "editor" && "admin" != "admin" && "admin" != "impersonate" && "admin" != "configure") || (role == "viewer" && ("admin" == "read" || "admin" == "list"))
	case ActionConfigure:
		return role == "owner" || role == "admin" || (role == "editor" && "configure" != "admin" && "configure" != "impersonate" && "configure" != "configure") || (role == "viewer" && ("configure" == "read" || "configure" == "list"))
	case ActionExecute:
		return role == "owner" || role == "admin" || (role == "editor" && "execute" != "admin" && "execute" != "impersonate" && "execute" != "configure") || (role == "viewer" && ("execute" == "read" || "execute" == "list"))
	case ActionImpersonate:
		return role == "owner" || role == "admin" || (role == "editor" && "impersonate" != "admin" && "impersonate" != "impersonate" && "impersonate" != "configure") || (role == "viewer" && ("impersonate" == "read" || "impersonate" == "list"))
	default:
		return role == "owner" || role == "admin"
	}
}

func permissionForAction(action string) string {
	switch action {
	case ActionRead, ActionList:
		return acl.PermViewer
	case ActionCreate, ActionUpdate, ActionShare, ActionExport, ActionImport, ActionRestore, ActionArchive, ActionRotate, ActionExecute:
		return acl.PermEditor
	case ActionDelete, ActionApprove, ActionDeny, ActionRevoke, ActionConfigure, ActionAdmin:
		return acl.PermAdmin
	case ActionImpersonate:
		return acl.PermOwner
	default:
		return acl.PermAdmin
	}
}

// AllActions returns known actions.
func AllActions() []string {
	return []string{"billing-service.create", "billing-service.read", "billing-service.update", "billing-service.delete", "billing-service.list", "billing-service.export", "billing-service.import", "billing-service.share", "billing-service.rotate", "billing-service.approve", "billing-service.deny", "billing-service.revoke", "billing-service.restore", "billing-service.archive", "billing-service.admin", "billing-service.configure", "billing-service.execute", "billing-service.impersonate"}
}
