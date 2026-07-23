package validators

import "testing"

func TestCreditCardValidatorTable(t *testing.T) {
	v := CreditCardValidator{}
	cases := []struct {
		name    string
		ok      bool
		payload map[string]any
	}{
		{name: "ok_0", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "000"}},
		{name: "bad_0", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_1", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "001"}},
		{name: "bad_1", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_2", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "002"}},
		{name: "bad_2", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_3", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "003"}},
		{name: "bad_3", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_4", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "004"}},
		{name: "bad_4", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_5", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "005"}},
		{name: "bad_5", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_6", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "006"}},
		{name: "bad_6", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_7", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "007"}},
		{name: "bad_7", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_8", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "008"}},
		{name: "bad_8", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_9", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "009"}},
		{name: "bad_9", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_10", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "010"}},
		{name: "bad_10", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_11", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "011"}},
		{name: "bad_11", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_12", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "012"}},
		{name: "bad_12", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_13", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "013"}},
		{name: "bad_13", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_14", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "014"}},
		{name: "bad_14", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_15", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "015"}},
		{name: "bad_15", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_16", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "016"}},
		{name: "bad_16", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_17", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "017"}},
		{name: "bad_17", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_18", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "018"}},
		{name: "bad_18", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_19", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "019"}},
		{name: "bad_19", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_20", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "020"}},
		{name: "bad_20", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_21", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "021"}},
		{name: "bad_21", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_22", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "022"}},
		{name: "bad_22", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_23", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "023"}},
		{name: "bad_23", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_24", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "024"}},
		{name: "bad_24", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_25", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "025"}},
		{name: "bad_25", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_26", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "026"}},
		{name: "bad_26", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_27", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "027"}},
		{name: "bad_27", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_28", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "028"}},
		{name: "bad_28", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_29", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "029"}},
		{name: "bad_29", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_30", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "030"}},
		{name: "bad_30", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_31", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "031"}},
		{name: "bad_31", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_32", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "032"}},
		{name: "bad_32", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_33", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "033"}},
		{name: "bad_33", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_34", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "034"}},
		{name: "bad_34", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_35", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "035"}},
		{name: "bad_35", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_36", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "036"}},
		{name: "bad_36", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_37", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "037"}},
		{name: "bad_37", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_38", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "038"}},
		{name: "bad_38", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_39", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "039"}},
		{name: "bad_39", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_40", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "040"}},
		{name: "bad_40", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_41", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "041"}},
		{name: "bad_41", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_42", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "042"}},
		{name: "bad_42", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_43", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "043"}},
		{name: "bad_43", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_44", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "044"}},
		{name: "bad_44", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_45", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "045"}},
		{name: "bad_45", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_46", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "046"}},
		{name: "bad_46", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_47", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "047"}},
		{name: "bad_47", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_48", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "048"}},
		{name: "bad_48", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
		{name: "ok_49", ok: true, payload: map[string]any{"number": "4111111111111111", "exp_month": "12", "exp_year": "2030", "cvv": "049"}},
		{name: "bad_49", ok: false, payload: map[string]any{"number": "", "exp_month": "", "exp_year": "", "cvv": ""}},
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
