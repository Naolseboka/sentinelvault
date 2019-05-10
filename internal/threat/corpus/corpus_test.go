package corpus

import "testing"

func TestCorpus(t *testing.T) {
	if Size() < 100 {
		t.Fatalf("size=%d", Size())
	}
	if !IsCommon("password") || !IsCommon("PASSWORD") {
		t.Fatal("expected common")
	}
	if IsCommon("Str0ng!UniqueVaultPassphrase#998877") {
		t.Fatal("unexpected")
	}
}
