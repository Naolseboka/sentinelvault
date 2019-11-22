package vault

import (
	"testing"
	"time"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

func TestValidatePayloadRequiredFields(t *testing.T) {
	tests := []struct {
		name    string
		typ     ItemType
		payload ItemPayload
		wantErr bool
	}{
		{
			name:    "password ok",
			typ:     ItemTypePassword,
			payload: ItemPayload{"password": "s3cret", "username": "alice"},
			wantErr: false,
		},
		{
			name:    "password missing",
			typ:     ItemTypePassword,
			payload: ItemPayload{"username": "alice"},
			wantErr: true,
		},
		{
			name:    "credit card ok",
			typ:     ItemTypeCreditCard,
			payload: ItemPayload{"number": "4111111111111111", "expiry": "12/28"},
			wantErr: false,
		},
		{
			name:    "db cred ok",
			typ:     ItemTypeDatabaseCredential,
			payload: ItemPayload{"username": "db", "password": "pw", "host": "localhost"},
			wantErr: false,
		},
		{
			name:    "unknown type",
			typ:     ItemType("unknown"),
			payload: ItemPayload{"x": "y"},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePayload(tc.typ, tc.payload)
			if tc.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestValidateMetadataRejectsSecrets(t *testing.T) {
	meta := ItemMetadata{
		CustomFields: map[string]string{"password": "plaintext"},
	}
	err := ValidateMetadata(ItemTypePassword, meta)
	if err == nil || !common.IsCode(err, common.ErrCodeInvalidArgument) {
		t.Fatalf("expected invalid_argument, got %v", err)
	}
}

func TestShareGrantActive(t *testing.T) {
	now := common.NowUTC()
	grant := ShareGrant{ExpiresAt: now.Add(time.Hour)}
	if !grant.IsActive(now) {
		t.Fatal("expected active grant")
	}
	revoked := now
	grant.RevokedAt = &revoked
	if grant.IsActive(now) {
		t.Fatal("expected revoked grant inactive")
	}
}
