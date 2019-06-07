package validators

import "testing"

func TestPasswordValidatorTable(t *testing.T) {
	v := PasswordValidator{}
	cases := []struct {
		name string
		ok bool
		payload map[string]any
	}{
	{name: "ok_0", ok: true, payload: map[string]any{"username": "v0-username", "password": "v0-password"}},
	{name: "bad_0", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_1", ok: true, payload: map[string]any{"username": "v1-username", "password": "v1-password"}},
	{name: "bad_1", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_2", ok: true, payload: map[string]any{"username": "v2-username", "password": "v2-password"}},
	{name: "bad_2", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_3", ok: true, payload: map[string]any{"username": "v3-username", "password": "v3-password"}},
	{name: "bad_3", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_4", ok: true, payload: map[string]any{"username": "v4-username", "password": "v4-password"}},
	{name: "bad_4", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_5", ok: true, payload: map[string]any{"username": "v5-username", "password": "v5-password"}},
	{name: "bad_5", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_6", ok: true, payload: map[string]any{"username": "v6-username", "password": "v6-password"}},
	{name: "bad_6", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_7", ok: true, payload: map[string]any{"username": "v7-username", "password": "v7-password"}},
	{name: "bad_7", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_8", ok: true, payload: map[string]any{"username": "v8-username", "password": "v8-password"}},
	{name: "bad_8", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_9", ok: true, payload: map[string]any{"username": "v9-username", "password": "v9-password"}},
	{name: "bad_9", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_10", ok: true, payload: map[string]any{"username": "v10-username", "password": "v10-password"}},
	{name: "bad_10", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_11", ok: true, payload: map[string]any{"username": "v11-username", "password": "v11-password"}},
	{name: "bad_11", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_12", ok: true, payload: map[string]any{"username": "v12-username", "password": "v12-password"}},
	{name: "bad_12", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_13", ok: true, payload: map[string]any{"username": "v13-username", "password": "v13-password"}},
	{name: "bad_13", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_14", ok: true, payload: map[string]any{"username": "v14-username", "password": "v14-password"}},
	{name: "bad_14", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_15", ok: true, payload: map[string]any{"username": "v15-username", "password": "v15-password"}},
	{name: "bad_15", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_16", ok: true, payload: map[string]any{"username": "v16-username", "password": "v16-password"}},
	{name: "bad_16", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_17", ok: true, payload: map[string]any{"username": "v17-username", "password": "v17-password"}},
	{name: "bad_17", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_18", ok: true, payload: map[string]any{"username": "v18-username", "password": "v18-password"}},
	{name: "bad_18", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_19", ok: true, payload: map[string]any{"username": "v19-username", "password": "v19-password"}},
	{name: "bad_19", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_20", ok: true, payload: map[string]any{"username": "v20-username", "password": "v20-password"}},
	{name: "bad_20", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_21", ok: true, payload: map[string]any{"username": "v21-username", "password": "v21-password"}},
	{name: "bad_21", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_22", ok: true, payload: map[string]any{"username": "v22-username", "password": "v22-password"}},
	{name: "bad_22", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_23", ok: true, payload: map[string]any{"username": "v23-username", "password": "v23-password"}},
	{name: "bad_23", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_24", ok: true, payload: map[string]any{"username": "v24-username", "password": "v24-password"}},
	{name: "bad_24", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_25", ok: true, payload: map[string]any{"username": "v25-username", "password": "v25-password"}},
	{name: "bad_25", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_26", ok: true, payload: map[string]any{"username": "v26-username", "password": "v26-password"}},
	{name: "bad_26", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_27", ok: true, payload: map[string]any{"username": "v27-username", "password": "v27-password"}},
	{name: "bad_27", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_28", ok: true, payload: map[string]any{"username": "v28-username", "password": "v28-password"}},
	{name: "bad_28", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_29", ok: true, payload: map[string]any{"username": "v29-username", "password": "v29-password"}},
	{name: "bad_29", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_30", ok: true, payload: map[string]any{"username": "v30-username", "password": "v30-password"}},
	{name: "bad_30", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_31", ok: true, payload: map[string]any{"username": "v31-username", "password": "v31-password"}},
	{name: "bad_31", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_32", ok: true, payload: map[string]any{"username": "v32-username", "password": "v32-password"}},
	{name: "bad_32", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_33", ok: true, payload: map[string]any{"username": "v33-username", "password": "v33-password"}},
	{name: "bad_33", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_34", ok: true, payload: map[string]any{"username": "v34-username", "password": "v34-password"}},
	{name: "bad_34", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_35", ok: true, payload: map[string]any{"username": "v35-username", "password": "v35-password"}},
	{name: "bad_35", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_36", ok: true, payload: map[string]any{"username": "v36-username", "password": "v36-password"}},
	{name: "bad_36", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_37", ok: true, payload: map[string]any{"username": "v37-username", "password": "v37-password"}},
	{name: "bad_37", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_38", ok: true, payload: map[string]any{"username": "v38-username", "password": "v38-password"}},
	{name: "bad_38", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_39", ok: true, payload: map[string]any{"username": "v39-username", "password": "v39-password"}},
	{name: "bad_39", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_40", ok: true, payload: map[string]any{"username": "v40-username", "password": "v40-password"}},
	{name: "bad_40", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_41", ok: true, payload: map[string]any{"username": "v41-username", "password": "v41-password"}},
	{name: "bad_41", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_42", ok: true, payload: map[string]any{"username": "v42-username", "password": "v42-password"}},
	{name: "bad_42", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_43", ok: true, payload: map[string]any{"username": "v43-username", "password": "v43-password"}},
	{name: "bad_43", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_44", ok: true, payload: map[string]any{"username": "v44-username", "password": "v44-password"}},
	{name: "bad_44", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_45", ok: true, payload: map[string]any{"username": "v45-username", "password": "v45-password"}},
	{name: "bad_45", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_46", ok: true, payload: map[string]any{"username": "v46-username", "password": "v46-password"}},
	{name: "bad_46", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_47", ok: true, payload: map[string]any{"username": "v47-username", "password": "v47-password"}},
	{name: "bad_47", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_48", ok: true, payload: map[string]any{"username": "v48-username", "password": "v48-password"}},
	{name: "bad_48", ok: false, payload: map[string]any{"username": "", "password": ""}},
	{name: "ok_49", ok: true, payload: map[string]any{"username": "v49-username", "password": "v49-password"}},
	{name: "bad_49", ok: false, payload: map[string]any{"username": "", "password": ""}},
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
