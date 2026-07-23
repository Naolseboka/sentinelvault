package statemachine

import (
	"fmt"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

const (
	StateInit         = "init"
	StateDraft        = "draft"
	StatePending      = "pending"
	StateValidating   = "validating"
	StateApproved     = "approved"
	StateDenied       = "denied"
	StateActive       = "active"
	StateInactive     = "inactive"
	StateSuspended    = "suspended"
	StateLocked       = "locked"
	StateRotating     = "rotating"
	StatePropagating  = "propagating"
	StateCompleted    = "completed"
	StateFailed       = "failed"
	StateArchived     = "archived"
	StateDeleted      = "deleted"
	StateRestoring    = "restoring"
	StateExported     = "exported"
	StateImported     = "imported"
	StateShared       = "shared"
	StateExpired      = "expired"
	StateRevoked      = "revoked"
	StateCompensating = "compensating"
)

var allowedPolicy = map[string]map[string]struct{}{
	StateActive:       {StateApproved: {}, StateDenied: {}, StateExpired: {}, StateImported: {}, StateInit: {}, StatePending: {}, StateRestoring: {}, StateRotating: {}},
	StateApproved:     {StateArchived: {}, StateCompensating: {}, StateLocked: {}},
	StateArchived:     {StateApproved: {}, StateDenied: {}, StateLocked: {}, StateRevoked: {}, StateShared: {}, StateValidating: {}},
	StateCompensating: {StateDenied: {}, StateExported: {}, StateInactive: {}, StatePending: {}, StateRevoked: {}, StateShared: {}, StateValidating: {}},
	StateCompleted:    {StateApproved: {}, StateArchived: {}, StateDenied: {}, StateLocked: {}, StateRevoked: {}, StateShared: {}},
	StateDeleted:      {StateImported: {}, StateInit: {}, StateRestoring: {}, StateRotating: {}},
	StateDenied:       {StateDraft: {}, StateSuspended: {}},
	StateDraft:        {StateArchived: {}, StateDenied: {}, StateLocked: {}, StateRevoked: {}, StateShared: {}, StateValidating: {}},
	StateExpired:      {StateCompleted: {}, StateInit: {}, StateRotating: {}},
	StateExported:     {StateApproved: {}, StateArchived: {}, StateDenied: {}, StateFailed: {}, StateRestoring: {}, StateRevoked: {}},
	StateFailed:       {StateCompleted: {}, StateDraft: {}, StateExpired: {}, StateInit: {}},
	StateImported:     {StateCompensating: {}, StateLocked: {}},
	StateInactive:     {StateImported: {}, StateRestoring: {}},
	StateInit:         {StateApproved: {}, StateArchived: {}, StateDenied: {}, StateExported: {}, StateRevoked: {}, StateShared: {}, StateValidating: {}},
	StateLocked:       {StateCompleted: {}, StateExpired: {}, StateInit: {}, StatePending: {}, StateSuspended: {}},
	StatePending:      {StateDeleted: {}, StatePropagating: {}, StateShared: {}},
	StatePropagating:  {StateCompleted: {}, StateExpired: {}, StateImported: {}, StateInit: {}, StateRestoring: {}, StateRotating: {}},
	StateRestoring:    {StateDeleted: {}, StateExported: {}, StateInactive: {}, StatePropagating: {}},
	StateRevoked:      {StateCompensating: {}, StateLocked: {}},
	StateRotating:     {StateDraft: {}, StateExpired: {}, StateExported: {}, StateInactive: {}, StatePending: {}, StateSuspended: {}},
	StateShared:       {StateActive: {}, StateDeleted: {}, StateInactive: {}, StatePropagating: {}},
	StateSuspended:    {StateArchived: {}, StateDenied: {}, StateInactive: {}, StateRevoked: {}, StateShared: {}, StateValidating: {}},
	StateValidating:   {StateApproved: {}, StateCompleted: {}, StateExpired: {}, StateInit: {}, StatePending: {}, StateRotating: {}, StateSuspended: {}},
}

// CanTransition reports whether from->to is allowed for policy-service.
func CanTransition(from, to string) bool {
	m, ok := allowedPolicy[from]
	if !ok {
		return false
	}
	_, ok = m[to]
	return ok
}

// Transition validates and returns the new state.
func Transition(from, to string) (string, error) {
	if from == "" || to == "" {
		return "", common.New(common.ErrCodeInvalidArgument, "from/to required")
	}
	if !CanTransition(from, to) {
		return "", common.New(common.ErrCodeConflict, fmt.Sprintf("policy: cannot transition %s -> %s", from, to))
	}
	return to, nil
}

// AllStates returns known states.
func AllStates() []string {
	return []string{"init", "draft", "pending", "validating", "approved", "denied", "active", "inactive", "suspended", "locked", "rotating", "propagating", "completed", "failed", "archived", "deleted", "restoring", "exported", "imported", "shared", "expired", "revoked", "compensating"}
}

// ApplySequence applies multiple transitions starting from from.
func ApplySequence(from string, steps []string) (string, error) {
	cur := from
	for _, step := range steps {
		next, err := Transition(cur, step)
		if err != nil {
			return cur, err
		}
		cur = next
	}
	return cur, nil
}
