package validators

import "testing"

func TestAPIKeyValidatorTable(t *testing.T) {
	v := APIKeyValidator{}
	cases := []struct {
		name string
		ok bool
		payload map[string]any
	}{
	{name: "ok_0", ok: true, payload: map[string]any{"key": "v0-key"}},
	{name: "bad_0", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_1", ok: true, payload: map[string]any{"key": "v1-key"}},
	{name: "bad_1", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_2", ok: true, payload: map[string]any{"key": "v2-key"}},
	{name: "bad_2", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_3", ok: true, payload: map[string]any{"key": "v3-key"}},
	{name: "bad_3", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_4", ok: true, payload: map[string]any{"key": "v4-key"}},
	{name: "bad_4", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_5", ok: true, payload: map[string]any{"key": "v5-key"}},
	{name: "bad_5", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_6", ok: true, payload: map[string]any{"key": "v6-key"}},
	{name: "bad_6", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_7", ok: true, payload: map[string]any{"key": "v7-key"}},
	{name: "bad_7", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_8", ok: true, payload: map[string]any{"key": "v8-key"}},
	{name: "bad_8", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_9", ok: true, payload: map[string]any{"key": "v9-key"}},
	{name: "bad_9", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_10", ok: true, payload: map[string]any{"key": "v10-key"}},
	{name: "bad_10", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_11", ok: true, payload: map[string]any{"key": "v11-key"}},
	{name: "bad_11", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_12", ok: true, payload: map[string]any{"key": "v12-key"}},
	{name: "bad_12", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_13", ok: true, payload: map[string]any{"key": "v13-key"}},
	{name: "bad_13", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_14", ok: true, payload: map[string]any{"key": "v14-key"}},
	{name: "bad_14", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_15", ok: true, payload: map[string]any{"key": "v15-key"}},
	{name: "bad_15", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_16", ok: true, payload: map[string]any{"key": "v16-key"}},
	{name: "bad_16", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_17", ok: true, payload: map[string]any{"key": "v17-key"}},
	{name: "bad_17", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_18", ok: true, payload: map[string]any{"key": "v18-key"}},
	{name: "bad_18", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_19", ok: true, payload: map[string]any{"key": "v19-key"}},
	{name: "bad_19", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_20", ok: true, payload: map[string]any{"key": "v20-key"}},
	{name: "bad_20", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_21", ok: true, payload: map[string]any{"key": "v21-key"}},
	{name: "bad_21", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_22", ok: true, payload: map[string]any{"key": "v22-key"}},
	{name: "bad_22", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_23", ok: true, payload: map[string]any{"key": "v23-key"}},
	{name: "bad_23", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_24", ok: true, payload: map[string]any{"key": "v24-key"}},
	{name: "bad_24", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_25", ok: true, payload: map[string]any{"key": "v25-key"}},
	{name: "bad_25", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_26", ok: true, payload: map[string]any{"key": "v26-key"}},
	{name: "bad_26", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_27", ok: true, payload: map[string]any{"key": "v27-key"}},
	{name: "bad_27", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_28", ok: true, payload: map[string]any{"key": "v28-key"}},
	{name: "bad_28", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_29", ok: true, payload: map[string]any{"key": "v29-key"}},
	{name: "bad_29", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_30", ok: true, payload: map[string]any{"key": "v30-key"}},
	{name: "bad_30", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_31", ok: true, payload: map[string]any{"key": "v31-key"}},
	{name: "bad_31", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_32", ok: true, payload: map[string]any{"key": "v32-key"}},
	{name: "bad_32", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_33", ok: true, payload: map[string]any{"key": "v33-key"}},
	{name: "bad_33", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_34", ok: true, payload: map[string]any{"key": "v34-key"}},
	{name: "bad_34", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_35", ok: true, payload: map[string]any{"key": "v35-key"}},
	{name: "bad_35", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_36", ok: true, payload: map[string]any{"key": "v36-key"}},
	{name: "bad_36", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_37", ok: true, payload: map[string]any{"key": "v37-key"}},
	{name: "bad_37", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_38", ok: true, payload: map[string]any{"key": "v38-key"}},
	{name: "bad_38", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_39", ok: true, payload: map[string]any{"key": "v39-key"}},
	{name: "bad_39", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_40", ok: true, payload: map[string]any{"key": "v40-key"}},
	{name: "bad_40", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_41", ok: true, payload: map[string]any{"key": "v41-key"}},
	{name: "bad_41", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_42", ok: true, payload: map[string]any{"key": "v42-key"}},
	{name: "bad_42", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_43", ok: true, payload: map[string]any{"key": "v43-key"}},
	{name: "bad_43", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_44", ok: true, payload: map[string]any{"key": "v44-key"}},
	{name: "bad_44", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_45", ok: true, payload: map[string]any{"key": "v45-key"}},
	{name: "bad_45", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_46", ok: true, payload: map[string]any{"key": "v46-key"}},
	{name: "bad_46", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_47", ok: true, payload: map[string]any{"key": "v47-key"}},
	{name: "bad_47", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_48", ok: true, payload: map[string]any{"key": "v48-key"}},
	{name: "bad_48", ok: false, payload: map[string]any{"key": ""}},
	{name: "ok_49", ok: true, payload: map[string]any{"key": "v49-key"}},
	{name: "bad_49", ok: false, payload: map[string]any{"key": ""}},
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
