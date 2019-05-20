package validators

import "testing"

func TestDatabaseCredentialValidatorTable(t *testing.T) {
	v := DatabaseCredentialValidator{}
	cases := []struct {
		name string
		ok bool
		payload map[string]any
	}{
	{name: "ok_0", ok: true, payload: map[string]any{"username": "v0-username", "password": "v0-password", "host": "v0-host"}},
	{name: "bad_0", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_1", ok: true, payload: map[string]any{"username": "v1-username", "password": "v1-password", "host": "v1-host"}},
	{name: "bad_1", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_2", ok: true, payload: map[string]any{"username": "v2-username", "password": "v2-password", "host": "v2-host"}},
	{name: "bad_2", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_3", ok: true, payload: map[string]any{"username": "v3-username", "password": "v3-password", "host": "v3-host"}},
	{name: "bad_3", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_4", ok: true, payload: map[string]any{"username": "v4-username", "password": "v4-password", "host": "v4-host"}},
	{name: "bad_4", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_5", ok: true, payload: map[string]any{"username": "v5-username", "password": "v5-password", "host": "v5-host"}},
	{name: "bad_5", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_6", ok: true, payload: map[string]any{"username": "v6-username", "password": "v6-password", "host": "v6-host"}},
	{name: "bad_6", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_7", ok: true, payload: map[string]any{"username": "v7-username", "password": "v7-password", "host": "v7-host"}},
	{name: "bad_7", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_8", ok: true, payload: map[string]any{"username": "v8-username", "password": "v8-password", "host": "v8-host"}},
	{name: "bad_8", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_9", ok: true, payload: map[string]any{"username": "v9-username", "password": "v9-password", "host": "v9-host"}},
	{name: "bad_9", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_10", ok: true, payload: map[string]any{"username": "v10-username", "password": "v10-password", "host": "v10-host"}},
	{name: "bad_10", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_11", ok: true, payload: map[string]any{"username": "v11-username", "password": "v11-password", "host": "v11-host"}},
	{name: "bad_11", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_12", ok: true, payload: map[string]any{"username": "v12-username", "password": "v12-password", "host": "v12-host"}},
	{name: "bad_12", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_13", ok: true, payload: map[string]any{"username": "v13-username", "password": "v13-password", "host": "v13-host"}},
	{name: "bad_13", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_14", ok: true, payload: map[string]any{"username": "v14-username", "password": "v14-password", "host": "v14-host"}},
	{name: "bad_14", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_15", ok: true, payload: map[string]any{"username": "v15-username", "password": "v15-password", "host": "v15-host"}},
	{name: "bad_15", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_16", ok: true, payload: map[string]any{"username": "v16-username", "password": "v16-password", "host": "v16-host"}},
	{name: "bad_16", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_17", ok: true, payload: map[string]any{"username": "v17-username", "password": "v17-password", "host": "v17-host"}},
	{name: "bad_17", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_18", ok: true, payload: map[string]any{"username": "v18-username", "password": "v18-password", "host": "v18-host"}},
	{name: "bad_18", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_19", ok: true, payload: map[string]any{"username": "v19-username", "password": "v19-password", "host": "v19-host"}},
	{name: "bad_19", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_20", ok: true, payload: map[string]any{"username": "v20-username", "password": "v20-password", "host": "v20-host"}},
	{name: "bad_20", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_21", ok: true, payload: map[string]any{"username": "v21-username", "password": "v21-password", "host": "v21-host"}},
	{name: "bad_21", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_22", ok: true, payload: map[string]any{"username": "v22-username", "password": "v22-password", "host": "v22-host"}},
	{name: "bad_22", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_23", ok: true, payload: map[string]any{"username": "v23-username", "password": "v23-password", "host": "v23-host"}},
	{name: "bad_23", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_24", ok: true, payload: map[string]any{"username": "v24-username", "password": "v24-password", "host": "v24-host"}},
	{name: "bad_24", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_25", ok: true, payload: map[string]any{"username": "v25-username", "password": "v25-password", "host": "v25-host"}},
	{name: "bad_25", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_26", ok: true, payload: map[string]any{"username": "v26-username", "password": "v26-password", "host": "v26-host"}},
	{name: "bad_26", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_27", ok: true, payload: map[string]any{"username": "v27-username", "password": "v27-password", "host": "v27-host"}},
	{name: "bad_27", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_28", ok: true, payload: map[string]any{"username": "v28-username", "password": "v28-password", "host": "v28-host"}},
	{name: "bad_28", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_29", ok: true, payload: map[string]any{"username": "v29-username", "password": "v29-password", "host": "v29-host"}},
	{name: "bad_29", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_30", ok: true, payload: map[string]any{"username": "v30-username", "password": "v30-password", "host": "v30-host"}},
	{name: "bad_30", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_31", ok: true, payload: map[string]any{"username": "v31-username", "password": "v31-password", "host": "v31-host"}},
	{name: "bad_31", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_32", ok: true, payload: map[string]any{"username": "v32-username", "password": "v32-password", "host": "v32-host"}},
	{name: "bad_32", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_33", ok: true, payload: map[string]any{"username": "v33-username", "password": "v33-password", "host": "v33-host"}},
	{name: "bad_33", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_34", ok: true, payload: map[string]any{"username": "v34-username", "password": "v34-password", "host": "v34-host"}},
	{name: "bad_34", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_35", ok: true, payload: map[string]any{"username": "v35-username", "password": "v35-password", "host": "v35-host"}},
	{name: "bad_35", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_36", ok: true, payload: map[string]any{"username": "v36-username", "password": "v36-password", "host": "v36-host"}},
	{name: "bad_36", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_37", ok: true, payload: map[string]any{"username": "v37-username", "password": "v37-password", "host": "v37-host"}},
	{name: "bad_37", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_38", ok: true, payload: map[string]any{"username": "v38-username", "password": "v38-password", "host": "v38-host"}},
	{name: "bad_38", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_39", ok: true, payload: map[string]any{"username": "v39-username", "password": "v39-password", "host": "v39-host"}},
	{name: "bad_39", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_40", ok: true, payload: map[string]any{"username": "v40-username", "password": "v40-password", "host": "v40-host"}},
	{name: "bad_40", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_41", ok: true, payload: map[string]any{"username": "v41-username", "password": "v41-password", "host": "v41-host"}},
	{name: "bad_41", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_42", ok: true, payload: map[string]any{"username": "v42-username", "password": "v42-password", "host": "v42-host"}},
	{name: "bad_42", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_43", ok: true, payload: map[string]any{"username": "v43-username", "password": "v43-password", "host": "v43-host"}},
	{name: "bad_43", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_44", ok: true, payload: map[string]any{"username": "v44-username", "password": "v44-password", "host": "v44-host"}},
	{name: "bad_44", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_45", ok: true, payload: map[string]any{"username": "v45-username", "password": "v45-password", "host": "v45-host"}},
	{name: "bad_45", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_46", ok: true, payload: map[string]any{"username": "v46-username", "password": "v46-password", "host": "v46-host"}},
	{name: "bad_46", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_47", ok: true, payload: map[string]any{"username": "v47-username", "password": "v47-password", "host": "v47-host"}},
	{name: "bad_47", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_48", ok: true, payload: map[string]any{"username": "v48-username", "password": "v48-password", "host": "v48-host"}},
	{name: "bad_48", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
	{name: "ok_49", ok: true, payload: map[string]any{"username": "v49-username", "password": "v49-password", "host": "v49-host"}},
	{name: "bad_49", ok: false, payload: map[string]any{"username": "", "password": "", "host": ""}},
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
