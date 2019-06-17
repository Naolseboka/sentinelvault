package validators

import "testing"

func TestTokenValidatorTable(t *testing.T) {
	v := TokenValidator{}
	cases := []struct {
		name string
		ok bool
		payload map[string]any
	}{
	{name: "ok_0", ok: true, payload: map[string]any{"token": "v0-token"}},
	{name: "bad_0", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_1", ok: true, payload: map[string]any{"token": "v1-token"}},
	{name: "bad_1", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_2", ok: true, payload: map[string]any{"token": "v2-token"}},
	{name: "bad_2", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_3", ok: true, payload: map[string]any{"token": "v3-token"}},
	{name: "bad_3", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_4", ok: true, payload: map[string]any{"token": "v4-token"}},
	{name: "bad_4", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_5", ok: true, payload: map[string]any{"token": "v5-token"}},
	{name: "bad_5", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_6", ok: true, payload: map[string]any{"token": "v6-token"}},
	{name: "bad_6", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_7", ok: true, payload: map[string]any{"token": "v7-token"}},
	{name: "bad_7", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_8", ok: true, payload: map[string]any{"token": "v8-token"}},
	{name: "bad_8", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_9", ok: true, payload: map[string]any{"token": "v9-token"}},
	{name: "bad_9", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_10", ok: true, payload: map[string]any{"token": "v10-token"}},
	{name: "bad_10", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_11", ok: true, payload: map[string]any{"token": "v11-token"}},
	{name: "bad_11", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_12", ok: true, payload: map[string]any{"token": "v12-token"}},
	{name: "bad_12", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_13", ok: true, payload: map[string]any{"token": "v13-token"}},
	{name: "bad_13", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_14", ok: true, payload: map[string]any{"token": "v14-token"}},
	{name: "bad_14", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_15", ok: true, payload: map[string]any{"token": "v15-token"}},
	{name: "bad_15", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_16", ok: true, payload: map[string]any{"token": "v16-token"}},
	{name: "bad_16", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_17", ok: true, payload: map[string]any{"token": "v17-token"}},
	{name: "bad_17", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_18", ok: true, payload: map[string]any{"token": "v18-token"}},
	{name: "bad_18", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_19", ok: true, payload: map[string]any{"token": "v19-token"}},
	{name: "bad_19", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_20", ok: true, payload: map[string]any{"token": "v20-token"}},
	{name: "bad_20", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_21", ok: true, payload: map[string]any{"token": "v21-token"}},
	{name: "bad_21", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_22", ok: true, payload: map[string]any{"token": "v22-token"}},
	{name: "bad_22", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_23", ok: true, payload: map[string]any{"token": "v23-token"}},
	{name: "bad_23", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_24", ok: true, payload: map[string]any{"token": "v24-token"}},
	{name: "bad_24", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_25", ok: true, payload: map[string]any{"token": "v25-token"}},
	{name: "bad_25", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_26", ok: true, payload: map[string]any{"token": "v26-token"}},
	{name: "bad_26", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_27", ok: true, payload: map[string]any{"token": "v27-token"}},
	{name: "bad_27", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_28", ok: true, payload: map[string]any{"token": "v28-token"}},
	{name: "bad_28", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_29", ok: true, payload: map[string]any{"token": "v29-token"}},
	{name: "bad_29", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_30", ok: true, payload: map[string]any{"token": "v30-token"}},
	{name: "bad_30", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_31", ok: true, payload: map[string]any{"token": "v31-token"}},
	{name: "bad_31", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_32", ok: true, payload: map[string]any{"token": "v32-token"}},
	{name: "bad_32", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_33", ok: true, payload: map[string]any{"token": "v33-token"}},
	{name: "bad_33", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_34", ok: true, payload: map[string]any{"token": "v34-token"}},
	{name: "bad_34", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_35", ok: true, payload: map[string]any{"token": "v35-token"}},
	{name: "bad_35", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_36", ok: true, payload: map[string]any{"token": "v36-token"}},
	{name: "bad_36", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_37", ok: true, payload: map[string]any{"token": "v37-token"}},
	{name: "bad_37", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_38", ok: true, payload: map[string]any{"token": "v38-token"}},
	{name: "bad_38", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_39", ok: true, payload: map[string]any{"token": "v39-token"}},
	{name: "bad_39", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_40", ok: true, payload: map[string]any{"token": "v40-token"}},
	{name: "bad_40", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_41", ok: true, payload: map[string]any{"token": "v41-token"}},
	{name: "bad_41", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_42", ok: true, payload: map[string]any{"token": "v42-token"}},
	{name: "bad_42", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_43", ok: true, payload: map[string]any{"token": "v43-token"}},
	{name: "bad_43", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_44", ok: true, payload: map[string]any{"token": "v44-token"}},
	{name: "bad_44", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_45", ok: true, payload: map[string]any{"token": "v45-token"}},
	{name: "bad_45", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_46", ok: true, payload: map[string]any{"token": "v46-token"}},
	{name: "bad_46", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_47", ok: true, payload: map[string]any{"token": "v47-token"}},
	{name: "bad_47", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_48", ok: true, payload: map[string]any{"token": "v48-token"}},
	{name: "bad_48", ok: false, payload: map[string]any{"token": ""}},
	{name: "ok_49", ok: true, payload: map[string]any{"token": "v49-token"}},
	{name: "bad_49", ok: false, payload: map[string]any{"token": ""}},
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
