package validators

import "testing"

func TestGenericSecretValidatorTable(t *testing.T) {
	v := GenericSecretValidator{}
	cases := []struct {
		name string
		ok bool
		payload map[string]any
	}{
	{name: "ok_0", ok: true, payload: map[string]any{"value": "v0-value"}},
	{name: "bad_0", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_1", ok: true, payload: map[string]any{"value": "v1-value"}},
	{name: "bad_1", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_2", ok: true, payload: map[string]any{"value": "v2-value"}},
	{name: "bad_2", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_3", ok: true, payload: map[string]any{"value": "v3-value"}},
	{name: "bad_3", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_4", ok: true, payload: map[string]any{"value": "v4-value"}},
	{name: "bad_4", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_5", ok: true, payload: map[string]any{"value": "v5-value"}},
	{name: "bad_5", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_6", ok: true, payload: map[string]any{"value": "v6-value"}},
	{name: "bad_6", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_7", ok: true, payload: map[string]any{"value": "v7-value"}},
	{name: "bad_7", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_8", ok: true, payload: map[string]any{"value": "v8-value"}},
	{name: "bad_8", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_9", ok: true, payload: map[string]any{"value": "v9-value"}},
	{name: "bad_9", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_10", ok: true, payload: map[string]any{"value": "v10-value"}},
	{name: "bad_10", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_11", ok: true, payload: map[string]any{"value": "v11-value"}},
	{name: "bad_11", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_12", ok: true, payload: map[string]any{"value": "v12-value"}},
	{name: "bad_12", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_13", ok: true, payload: map[string]any{"value": "v13-value"}},
	{name: "bad_13", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_14", ok: true, payload: map[string]any{"value": "v14-value"}},
	{name: "bad_14", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_15", ok: true, payload: map[string]any{"value": "v15-value"}},
	{name: "bad_15", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_16", ok: true, payload: map[string]any{"value": "v16-value"}},
	{name: "bad_16", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_17", ok: true, payload: map[string]any{"value": "v17-value"}},
	{name: "bad_17", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_18", ok: true, payload: map[string]any{"value": "v18-value"}},
	{name: "bad_18", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_19", ok: true, payload: map[string]any{"value": "v19-value"}},
	{name: "bad_19", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_20", ok: true, payload: map[string]any{"value": "v20-value"}},
	{name: "bad_20", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_21", ok: true, payload: map[string]any{"value": "v21-value"}},
	{name: "bad_21", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_22", ok: true, payload: map[string]any{"value": "v22-value"}},
	{name: "bad_22", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_23", ok: true, payload: map[string]any{"value": "v23-value"}},
	{name: "bad_23", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_24", ok: true, payload: map[string]any{"value": "v24-value"}},
	{name: "bad_24", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_25", ok: true, payload: map[string]any{"value": "v25-value"}},
	{name: "bad_25", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_26", ok: true, payload: map[string]any{"value": "v26-value"}},
	{name: "bad_26", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_27", ok: true, payload: map[string]any{"value": "v27-value"}},
	{name: "bad_27", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_28", ok: true, payload: map[string]any{"value": "v28-value"}},
	{name: "bad_28", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_29", ok: true, payload: map[string]any{"value": "v29-value"}},
	{name: "bad_29", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_30", ok: true, payload: map[string]any{"value": "v30-value"}},
	{name: "bad_30", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_31", ok: true, payload: map[string]any{"value": "v31-value"}},
	{name: "bad_31", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_32", ok: true, payload: map[string]any{"value": "v32-value"}},
	{name: "bad_32", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_33", ok: true, payload: map[string]any{"value": "v33-value"}},
	{name: "bad_33", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_34", ok: true, payload: map[string]any{"value": "v34-value"}},
	{name: "bad_34", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_35", ok: true, payload: map[string]any{"value": "v35-value"}},
	{name: "bad_35", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_36", ok: true, payload: map[string]any{"value": "v36-value"}},
	{name: "bad_36", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_37", ok: true, payload: map[string]any{"value": "v37-value"}},
	{name: "bad_37", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_38", ok: true, payload: map[string]any{"value": "v38-value"}},
	{name: "bad_38", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_39", ok: true, payload: map[string]any{"value": "v39-value"}},
	{name: "bad_39", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_40", ok: true, payload: map[string]any{"value": "v40-value"}},
	{name: "bad_40", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_41", ok: true, payload: map[string]any{"value": "v41-value"}},
	{name: "bad_41", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_42", ok: true, payload: map[string]any{"value": "v42-value"}},
	{name: "bad_42", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_43", ok: true, payload: map[string]any{"value": "v43-value"}},
	{name: "bad_43", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_44", ok: true, payload: map[string]any{"value": "v44-value"}},
	{name: "bad_44", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_45", ok: true, payload: map[string]any{"value": "v45-value"}},
	{name: "bad_45", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_46", ok: true, payload: map[string]any{"value": "v46-value"}},
	{name: "bad_46", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_47", ok: true, payload: map[string]any{"value": "v47-value"}},
	{name: "bad_47", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_48", ok: true, payload: map[string]any{"value": "v48-value"}},
	{name: "bad_48", ok: false, payload: map[string]any{"value": ""}},
	{name: "ok_49", ok: true, payload: map[string]any{"value": "v49-value"}},
	{name: "bad_49", ok: false, payload: map[string]any{"value": ""}},
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
