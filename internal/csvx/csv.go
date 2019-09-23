package csvx

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/sentinelvault/sentinelvault/internal/common"
)

type Row map[string]string

func Parse(r io.Reader) ([]Row, error) {
	cr := csv.NewReader(r)
	cr.TrimLeadingSpace = true
	cr.ReuseRecord = true
	header, err := cr.Read()
	if err != nil {
		return nil, common.Wrap(common.ErrCodeInvalidArgument, "csv header", err)
	}
	cols := make([]string, len(header))
	for i, h := range header {
		cols[i] = strings.TrimSpace(strings.ToLower(h))
	}
	out := make([]Row, 0, 64)
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, common.Wrap(common.ErrCodeInvalidArgument, "csv row", err)
		}
		row := make(Row, len(cols))
		for i, c := range cols {
			if i < len(rec) {
				row[c] = strings.TrimSpace(rec[i])
			} else {
				row[c] = ""
			}
		}
		out = append(out, row)
	}
	return out, nil
}

func Write(w io.Writer, cols []string, rows []Row) error {
	cw := csv.NewWriter(w)
	if err := cw.Write(cols); err != nil {
		return err
	}
	for _, row := range rows {
		rec := make([]string, len(cols))
		for i, c := range cols {
			rec[i] = row[c]
		}
		if err := cw.Write(rec); err != nil {
			return err
		}
	}
	cw.Flush()
	return cw.Error()
}

func Require(row Row, keys ...string) error {
	for _, k := range keys {
		if strings.TrimSpace(row[k]) == "" {
			return common.New(common.ErrCodeInvalidArgument, "missing column "+k)
		}
	}
	return nil
}
