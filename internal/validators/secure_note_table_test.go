package validators

import "testing"

func TestSecureNoteValidatorTable(t *testing.T) {
	v := SecureNoteValidator{}
	cases := []struct {
		name    string
		ok      bool
		payload map[string]any
	}{
		{name: "ok_0", ok: true, payload: map[string]any{"note": "v0-note"}},
		{name: "bad_0", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_1", ok: true, payload: map[string]any{"note": "v1-note"}},
		{name: "bad_1", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_2", ok: true, payload: map[string]any{"note": "v2-note"}},
		{name: "bad_2", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_3", ok: true, payload: map[string]any{"note": "v3-note"}},
		{name: "bad_3", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_4", ok: true, payload: map[string]any{"note": "v4-note"}},
		{name: "bad_4", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_5", ok: true, payload: map[string]any{"note": "v5-note"}},
		{name: "bad_5", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_6", ok: true, payload: map[string]any{"note": "v6-note"}},
		{name: "bad_6", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_7", ok: true, payload: map[string]any{"note": "v7-note"}},
		{name: "bad_7", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_8", ok: true, payload: map[string]any{"note": "v8-note"}},
		{name: "bad_8", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_9", ok: true, payload: map[string]any{"note": "v9-note"}},
		{name: "bad_9", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_10", ok: true, payload: map[string]any{"note": "v10-note"}},
		{name: "bad_10", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_11", ok: true, payload: map[string]any{"note": "v11-note"}},
		{name: "bad_11", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_12", ok: true, payload: map[string]any{"note": "v12-note"}},
		{name: "bad_12", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_13", ok: true, payload: map[string]any{"note": "v13-note"}},
		{name: "bad_13", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_14", ok: true, payload: map[string]any{"note": "v14-note"}},
		{name: "bad_14", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_15", ok: true, payload: map[string]any{"note": "v15-note"}},
		{name: "bad_15", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_16", ok: true, payload: map[string]any{"note": "v16-note"}},
		{name: "bad_16", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_17", ok: true, payload: map[string]any{"note": "v17-note"}},
		{name: "bad_17", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_18", ok: true, payload: map[string]any{"note": "v18-note"}},
		{name: "bad_18", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_19", ok: true, payload: map[string]any{"note": "v19-note"}},
		{name: "bad_19", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_20", ok: true, payload: map[string]any{"note": "v20-note"}},
		{name: "bad_20", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_21", ok: true, payload: map[string]any{"note": "v21-note"}},
		{name: "bad_21", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_22", ok: true, payload: map[string]any{"note": "v22-note"}},
		{name: "bad_22", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_23", ok: true, payload: map[string]any{"note": "v23-note"}},
		{name: "bad_23", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_24", ok: true, payload: map[string]any{"note": "v24-note"}},
		{name: "bad_24", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_25", ok: true, payload: map[string]any{"note": "v25-note"}},
		{name: "bad_25", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_26", ok: true, payload: map[string]any{"note": "v26-note"}},
		{name: "bad_26", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_27", ok: true, payload: map[string]any{"note": "v27-note"}},
		{name: "bad_27", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_28", ok: true, payload: map[string]any{"note": "v28-note"}},
		{name: "bad_28", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_29", ok: true, payload: map[string]any{"note": "v29-note"}},
		{name: "bad_29", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_30", ok: true, payload: map[string]any{"note": "v30-note"}},
		{name: "bad_30", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_31", ok: true, payload: map[string]any{"note": "v31-note"}},
		{name: "bad_31", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_32", ok: true, payload: map[string]any{"note": "v32-note"}},
		{name: "bad_32", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_33", ok: true, payload: map[string]any{"note": "v33-note"}},
		{name: "bad_33", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_34", ok: true, payload: map[string]any{"note": "v34-note"}},
		{name: "bad_34", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_35", ok: true, payload: map[string]any{"note": "v35-note"}},
		{name: "bad_35", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_36", ok: true, payload: map[string]any{"note": "v36-note"}},
		{name: "bad_36", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_37", ok: true, payload: map[string]any{"note": "v37-note"}},
		{name: "bad_37", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_38", ok: true, payload: map[string]any{"note": "v38-note"}},
		{name: "bad_38", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_39", ok: true, payload: map[string]any{"note": "v39-note"}},
		{name: "bad_39", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_40", ok: true, payload: map[string]any{"note": "v40-note"}},
		{name: "bad_40", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_41", ok: true, payload: map[string]any{"note": "v41-note"}},
		{name: "bad_41", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_42", ok: true, payload: map[string]any{"note": "v42-note"}},
		{name: "bad_42", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_43", ok: true, payload: map[string]any{"note": "v43-note"}},
		{name: "bad_43", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_44", ok: true, payload: map[string]any{"note": "v44-note"}},
		{name: "bad_44", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_45", ok: true, payload: map[string]any{"note": "v45-note"}},
		{name: "bad_45", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_46", ok: true, payload: map[string]any{"note": "v46-note"}},
		{name: "bad_46", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_47", ok: true, payload: map[string]any{"note": "v47-note"}},
		{name: "bad_47", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_48", ok: true, payload: map[string]any{"note": "v48-note"}},
		{name: "bad_48", ok: false, payload: map[string]any{"note": ""}},
		{name: "ok_49", ok: true, payload: map[string]any{"note": "v49-note"}},
		{name: "bad_49", ok: false, payload: map[string]any{"note": ""}},
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
