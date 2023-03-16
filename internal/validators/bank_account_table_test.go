package validators

import "testing"

func TestBankAccountValidatorTable(t *testing.T) {
	v := BankAccountValidator{}
	cases := []struct {
		name string
		ok bool
		payload map[string]any
	}{
	{name: "ok_0", ok: true, payload: map[string]any{"account_number": "v0-account_number", "routing_number": "v0-routing_number"}},
	{name: "bad_0", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_1", ok: true, payload: map[string]any{"account_number": "v1-account_number", "routing_number": "v1-routing_number"}},
	{name: "bad_1", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_2", ok: true, payload: map[string]any{"account_number": "v2-account_number", "routing_number": "v2-routing_number"}},
	{name: "bad_2", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_3", ok: true, payload: map[string]any{"account_number": "v3-account_number", "routing_number": "v3-routing_number"}},
	{name: "bad_3", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_4", ok: true, payload: map[string]any{"account_number": "v4-account_number", "routing_number": "v4-routing_number"}},
	{name: "bad_4", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_5", ok: true, payload: map[string]any{"account_number": "v5-account_number", "routing_number": "v5-routing_number"}},
	{name: "bad_5", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_6", ok: true, payload: map[string]any{"account_number": "v6-account_number", "routing_number": "v6-routing_number"}},
	{name: "bad_6", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_7", ok: true, payload: map[string]any{"account_number": "v7-account_number", "routing_number": "v7-routing_number"}},
	{name: "bad_7", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_8", ok: true, payload: map[string]any{"account_number": "v8-account_number", "routing_number": "v8-routing_number"}},
	{name: "bad_8", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_9", ok: true, payload: map[string]any{"account_number": "v9-account_number", "routing_number": "v9-routing_number"}},
	{name: "bad_9", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_10", ok: true, payload: map[string]any{"account_number": "v10-account_number", "routing_number": "v10-routing_number"}},
	{name: "bad_10", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_11", ok: true, payload: map[string]any{"account_number": "v11-account_number", "routing_number": "v11-routing_number"}},
	{name: "bad_11", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_12", ok: true, payload: map[string]any{"account_number": "v12-account_number", "routing_number": "v12-routing_number"}},
	{name: "bad_12", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_13", ok: true, payload: map[string]any{"account_number": "v13-account_number", "routing_number": "v13-routing_number"}},
	{name: "bad_13", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_14", ok: true, payload: map[string]any{"account_number": "v14-account_number", "routing_number": "v14-routing_number"}},
	{name: "bad_14", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_15", ok: true, payload: map[string]any{"account_number": "v15-account_number", "routing_number": "v15-routing_number"}},
	{name: "bad_15", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_16", ok: true, payload: map[string]any{"account_number": "v16-account_number", "routing_number": "v16-routing_number"}},
	{name: "bad_16", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_17", ok: true, payload: map[string]any{"account_number": "v17-account_number", "routing_number": "v17-routing_number"}},
	{name: "bad_17", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_18", ok: true, payload: map[string]any{"account_number": "v18-account_number", "routing_number": "v18-routing_number"}},
	{name: "bad_18", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_19", ok: true, payload: map[string]any{"account_number": "v19-account_number", "routing_number": "v19-routing_number"}},
	{name: "bad_19", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_20", ok: true, payload: map[string]any{"account_number": "v20-account_number", "routing_number": "v20-routing_number"}},
	{name: "bad_20", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_21", ok: true, payload: map[string]any{"account_number": "v21-account_number", "routing_number": "v21-routing_number"}},
	{name: "bad_21", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_22", ok: true, payload: map[string]any{"account_number": "v22-account_number", "routing_number": "v22-routing_number"}},
	{name: "bad_22", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_23", ok: true, payload: map[string]any{"account_number": "v23-account_number", "routing_number": "v23-routing_number"}},
	{name: "bad_23", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_24", ok: true, payload: map[string]any{"account_number": "v24-account_number", "routing_number": "v24-routing_number"}},
	{name: "bad_24", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_25", ok: true, payload: map[string]any{"account_number": "v25-account_number", "routing_number": "v25-routing_number"}},
	{name: "bad_25", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_26", ok: true, payload: map[string]any{"account_number": "v26-account_number", "routing_number": "v26-routing_number"}},
	{name: "bad_26", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_27", ok: true, payload: map[string]any{"account_number": "v27-account_number", "routing_number": "v27-routing_number"}},
	{name: "bad_27", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_28", ok: true, payload: map[string]any{"account_number": "v28-account_number", "routing_number": "v28-routing_number"}},
	{name: "bad_28", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_29", ok: true, payload: map[string]any{"account_number": "v29-account_number", "routing_number": "v29-routing_number"}},
	{name: "bad_29", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_30", ok: true, payload: map[string]any{"account_number": "v30-account_number", "routing_number": "v30-routing_number"}},
	{name: "bad_30", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_31", ok: true, payload: map[string]any{"account_number": "v31-account_number", "routing_number": "v31-routing_number"}},
	{name: "bad_31", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_32", ok: true, payload: map[string]any{"account_number": "v32-account_number", "routing_number": "v32-routing_number"}},
	{name: "bad_32", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_33", ok: true, payload: map[string]any{"account_number": "v33-account_number", "routing_number": "v33-routing_number"}},
	{name: "bad_33", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_34", ok: true, payload: map[string]any{"account_number": "v34-account_number", "routing_number": "v34-routing_number"}},
	{name: "bad_34", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_35", ok: true, payload: map[string]any{"account_number": "v35-account_number", "routing_number": "v35-routing_number"}},
	{name: "bad_35", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_36", ok: true, payload: map[string]any{"account_number": "v36-account_number", "routing_number": "v36-routing_number"}},
	{name: "bad_36", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_37", ok: true, payload: map[string]any{"account_number": "v37-account_number", "routing_number": "v37-routing_number"}},
	{name: "bad_37", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_38", ok: true, payload: map[string]any{"account_number": "v38-account_number", "routing_number": "v38-routing_number"}},
	{name: "bad_38", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_39", ok: true, payload: map[string]any{"account_number": "v39-account_number", "routing_number": "v39-routing_number"}},
	{name: "bad_39", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_40", ok: true, payload: map[string]any{"account_number": "v40-account_number", "routing_number": "v40-routing_number"}},
	{name: "bad_40", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_41", ok: true, payload: map[string]any{"account_number": "v41-account_number", "routing_number": "v41-routing_number"}},
	{name: "bad_41", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_42", ok: true, payload: map[string]any{"account_number": "v42-account_number", "routing_number": "v42-routing_number"}},
	{name: "bad_42", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_43", ok: true, payload: map[string]any{"account_number": "v43-account_number", "routing_number": "v43-routing_number"}},
	{name: "bad_43", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_44", ok: true, payload: map[string]any{"account_number": "v44-account_number", "routing_number": "v44-routing_number"}},
	{name: "bad_44", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_45", ok: true, payload: map[string]any{"account_number": "v45-account_number", "routing_number": "v45-routing_number"}},
	{name: "bad_45", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_46", ok: true, payload: map[string]any{"account_number": "v46-account_number", "routing_number": "v46-routing_number"}},
	{name: "bad_46", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_47", ok: true, payload: map[string]any{"account_number": "v47-account_number", "routing_number": "v47-routing_number"}},
	{name: "bad_47", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_48", ok: true, payload: map[string]any{"account_number": "v48-account_number", "routing_number": "v48-routing_number"}},
	{name: "bad_48", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
	{name: "ok_49", ok: true, payload: map[string]any{"account_number": "v49-account_number", "routing_number": "v49-routing_number"}},
	{name: "bad_49", ok: false, payload: map[string]any{"account_number": "", "routing_number": ""}},
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
// maintenance note 96
