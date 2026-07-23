package importers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
	"github.com/sentinelvault/sentinelvault/internal/validators"
)

// CreditCardImportResult is one imported row.
type CreditCardImportResult struct {
	Row      int            `json:"row"`
	Name     string         `json:"name"`
	Payload  map[string]any `json:"payload"`
	Warnings []string       `json:"warnings,omitempty"`
	Error    string         `json:"error,omitempty"`
}

// ImportCreditCardCSV parses CSV rows into validated credit_card payloads.
func ImportCreditCardCSV(r io.Reader) ([]CreditCardImportResult, error) {
	cr := csv.NewReader(r)
	cr.TrimLeadingSpace = true
	headers, err := cr.Read()
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "read csv header", err)
	}
	idx := map[string]int{}
	for i, h := range headers {
		idx[strings.ToLower(strings.TrimSpace(h))] = i
	}
	v := validators.CreditCardValidator{}
	out := make([]CreditCardImportResult, 0)
	rowNum := 1
	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		rowNum++
		if err != nil {
			out = append(out, CreditCardImportResult{Row: rowNum, Error: err.Error()})
			continue
		}
		rec := map[string]string{}
		for k, i := range idx {
			if i < len(row) {
				rec[k] = row[i]
			}
		}
		payload := map[string]any{}
		if v := strings.TrimSpace(rec["number"]); v != "" {
			payload["number"] = v
		}
		if v := strings.TrimSpace(rec["exp_month"]); v != "" {
			payload["exp_month"] = v
		}
		if v := strings.TrimSpace(rec["exp_year"]); v != "" {
			payload["exp_year"] = v
		}
		if v := strings.TrimSpace(rec["cvv"]); v != "" {
			payload["cvv"] = v
		}
		name := strings.TrimSpace(rec["name"])
		if name == "" {
			name = fmt.Sprintf("credit_card-row-%d", rowNum)
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, CreditCardImportResult{Row: rowNum, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, CreditCardImportResult{Row: rowNum, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}

// ImportCreditCardJSON parses a JSON array of objects.
func ImportCreditCardJSON(data []byte) ([]CreditCardImportResult, error) {
	var rows []map[string]any
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "invalid json", err)
	}
	v := validators.CreditCardValidator{}
	out := make([]CreditCardImportResult, 0, len(rows))
	for i, row := range rows {
		name, _ := row["name"].(string)
		if name == "" {
			name = fmt.Sprintf("credit_card-%d", i+1)
		}
		payload := row
		if p, ok := row["payload"].(map[string]any); ok {
			payload = p
		}
		norm, warnings, err := v.Validate(payload)
		if err != nil {
			out = append(out, CreditCardImportResult{Row: i + 1, Name: name, Error: err.Error()})
			continue
		}
		out = append(out, CreditCardImportResult{Row: i + 1, Name: name, Payload: norm, Warnings: warnings})
	}
	return out, nil
}
