package validators

import "testing"

func TestSSHKeyValidatorTable(t *testing.T) {
	v := SSHKeyValidator{}
	cases := []struct {
		name    string
		ok      bool
		payload map[string]any
	}{
		{name: "ok_0", ok: true, payload: map[string]any{"private_key": "v0-private_key"}},
		{name: "bad_0", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_1", ok: true, payload: map[string]any{"private_key": "v1-private_key"}},
		{name: "bad_1", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_2", ok: true, payload: map[string]any{"private_key": "v2-private_key"}},
		{name: "bad_2", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_3", ok: true, payload: map[string]any{"private_key": "v3-private_key"}},
		{name: "bad_3", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_4", ok: true, payload: map[string]any{"private_key": "v4-private_key"}},
		{name: "bad_4", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_5", ok: true, payload: map[string]any{"private_key": "v5-private_key"}},
		{name: "bad_5", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_6", ok: true, payload: map[string]any{"private_key": "v6-private_key"}},
		{name: "bad_6", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_7", ok: true, payload: map[string]any{"private_key": "v7-private_key"}},
		{name: "bad_7", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_8", ok: true, payload: map[string]any{"private_key": "v8-private_key"}},
		{name: "bad_8", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_9", ok: true, payload: map[string]any{"private_key": "v9-private_key"}},
		{name: "bad_9", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_10", ok: true, payload: map[string]any{"private_key": "v10-private_key"}},
		{name: "bad_10", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_11", ok: true, payload: map[string]any{"private_key": "v11-private_key"}},
		{name: "bad_11", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_12", ok: true, payload: map[string]any{"private_key": "v12-private_key"}},
		{name: "bad_12", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_13", ok: true, payload: map[string]any{"private_key": "v13-private_key"}},
		{name: "bad_13", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_14", ok: true, payload: map[string]any{"private_key": "v14-private_key"}},
		{name: "bad_14", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_15", ok: true, payload: map[string]any{"private_key": "v15-private_key"}},
		{name: "bad_15", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_16", ok: true, payload: map[string]any{"private_key": "v16-private_key"}},
		{name: "bad_16", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_17", ok: true, payload: map[string]any{"private_key": "v17-private_key"}},
		{name: "bad_17", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_18", ok: true, payload: map[string]any{"private_key": "v18-private_key"}},
		{name: "bad_18", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_19", ok: true, payload: map[string]any{"private_key": "v19-private_key"}},
		{name: "bad_19", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_20", ok: true, payload: map[string]any{"private_key": "v20-private_key"}},
		{name: "bad_20", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_21", ok: true, payload: map[string]any{"private_key": "v21-private_key"}},
		{name: "bad_21", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_22", ok: true, payload: map[string]any{"private_key": "v22-private_key"}},
		{name: "bad_22", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_23", ok: true, payload: map[string]any{"private_key": "v23-private_key"}},
		{name: "bad_23", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_24", ok: true, payload: map[string]any{"private_key": "v24-private_key"}},
		{name: "bad_24", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_25", ok: true, payload: map[string]any{"private_key": "v25-private_key"}},
		{name: "bad_25", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_26", ok: true, payload: map[string]any{"private_key": "v26-private_key"}},
		{name: "bad_26", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_27", ok: true, payload: map[string]any{"private_key": "v27-private_key"}},
		{name: "bad_27", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_28", ok: true, payload: map[string]any{"private_key": "v28-private_key"}},
		{name: "bad_28", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_29", ok: true, payload: map[string]any{"private_key": "v29-private_key"}},
		{name: "bad_29", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_30", ok: true, payload: map[string]any{"private_key": "v30-private_key"}},
		{name: "bad_30", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_31", ok: true, payload: map[string]any{"private_key": "v31-private_key"}},
		{name: "bad_31", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_32", ok: true, payload: map[string]any{"private_key": "v32-private_key"}},
		{name: "bad_32", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_33", ok: true, payload: map[string]any{"private_key": "v33-private_key"}},
		{name: "bad_33", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_34", ok: true, payload: map[string]any{"private_key": "v34-private_key"}},
		{name: "bad_34", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_35", ok: true, payload: map[string]any{"private_key": "v35-private_key"}},
		{name: "bad_35", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_36", ok: true, payload: map[string]any{"private_key": "v36-private_key"}},
		{name: "bad_36", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_37", ok: true, payload: map[string]any{"private_key": "v37-private_key"}},
		{name: "bad_37", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_38", ok: true, payload: map[string]any{"private_key": "v38-private_key"}},
		{name: "bad_38", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_39", ok: true, payload: map[string]any{"private_key": "v39-private_key"}},
		{name: "bad_39", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_40", ok: true, payload: map[string]any{"private_key": "v40-private_key"}},
		{name: "bad_40", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_41", ok: true, payload: map[string]any{"private_key": "v41-private_key"}},
		{name: "bad_41", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_42", ok: true, payload: map[string]any{"private_key": "v42-private_key"}},
		{name: "bad_42", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_43", ok: true, payload: map[string]any{"private_key": "v43-private_key"}},
		{name: "bad_43", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_44", ok: true, payload: map[string]any{"private_key": "v44-private_key"}},
		{name: "bad_44", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_45", ok: true, payload: map[string]any{"private_key": "v45-private_key"}},
		{name: "bad_45", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_46", ok: true, payload: map[string]any{"private_key": "v46-private_key"}},
		{name: "bad_46", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_47", ok: true, payload: map[string]any{"private_key": "v47-private_key"}},
		{name: "bad_47", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_48", ok: true, payload: map[string]any{"private_key": "v48-private_key"}},
		{name: "bad_48", ok: false, payload: map[string]any{"private_key": ""}},
		{name: "ok_49", ok: true, payload: map[string]any{"private_key": "v49-private_key"}},
		{name: "bad_49", ok: false, payload: map[string]any{"private_key": ""}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, err := v.Validate(tc.payload)
			if tc.ok && err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !tc.ok && err == nil {
				t.Fatal("expected error")
			}
		})
	}
}

// maintenance note 63
