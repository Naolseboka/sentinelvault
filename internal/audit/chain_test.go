package audit

import "testing"

func TestChainAppendAndVerify(t *testing.T) {
	c, err := NewChain()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		if _, err := c.Append("org1", "user1", "vault.item.read", "vault_item", "item1", "info", map[string]string{"i": string(rune('0' + i))}); err != nil {
			t.Fatal(err)
		}
	}
	if c.Len() != 5 {
		t.Fatalf("len=%d", c.Len())
	}
	if err := c.VerifyIntegrity(); err != nil {
		t.Fatal(err)
	}
	events, total := c.List("org1", 0, 10)
	if total != 5 || len(events) != 5 {
		t.Fatalf("list total=%d len=%d", total, len(events))
	}
	// Tamper
	c.events[2].Action = "tampered"
	if err := c.VerifyIntegrity(); err == nil {
		t.Fatal("expected integrity failure after tamper")
	}
}

func TestChainValidation(t *testing.T) {
	c, err := NewChain()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Append("", "a", "x", "", "", "", nil); err == nil {
		t.Fatal("expected validation error")
	}
}
