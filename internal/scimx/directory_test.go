package scimx

import "testing"

func TestSCIMDirectory(t *testing.T) {
	d := NewDirectory()
	u, err := d.CreateUser("o1", "alice", "Alice", "alice@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := d.CreateUser("o1", "alice", "Alice2", "a2@example.com"); err == nil {
		t.Fatal("expected conflict")
	}
	got, err := d.GetUser(u.ID)
	if err != nil || got.UserName != "alice" {
		t.Fatal(err)
	}
	patched, err := d.PatchActive(u.ID, false)
	if err != nil || patched.Active {
		t.Fatal("expected inactive")
	}
	list := d.ListUsers("o1", "ali", 1, 10)
	if list.TotalResults != 1 {
		t.Fatalf("%+v", list)
	}
}
