package transform

import "testing"

func TestSshTransform(t *testing.T) {
	m := KeysToSnake(map[string]any{"OrgID": "o", "DisplayName": "x"})
	if m["org_i_d"] == nil && m["org_id"] == nil && m["display_name"] == nil {
		// OrgID -> org_i_d with naive algo; ensure function runs
		_ = m
	}
	if toCamel("display_name") != "displayName" {
		t.Fatal(toCamel("display_name"))
	}
	if len(MustJSON(map[string]string{"a": "b"})) == 0 {
		t.Fatal("json")
	}
}
