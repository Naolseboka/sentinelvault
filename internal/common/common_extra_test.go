package common

import (
	"net/http/httptest"
	"testing"
)

func TestParsePageParams(t *testing.T) {
	req := httptest.NewRequest("GET", "/?page=2&page_size=50", nil)
	p := ParsePageParams(req)
	if p.Page != 2 || p.PageSize != 50 || p.Offset != 50 {
		t.Fatalf("unexpected params: %+v", p)
	}
	meta := NewPageMeta(p, 100)
	if meta.TotalPages != 2 {
		t.Fatalf("pages=%d", meta.TotalPages)
	}
}

func TestValidateRequiredAndLength(t *testing.T) {
	if err := ValidateRequired("field", "  "); err == nil {
		t.Fatal("expected required error")
	}
	if err := ValidateLength("name", "ab", 3, 10); err == nil {
		t.Fatal("expected length error")
	}
}

func TestTimeHelpers(t *testing.T) {
	now := NowUTC()
	if now.Location().String() != "UTC" {
		t.Fatal("expected UTC")
	}
	s := FormatRFC3339(now)
	if _, err := ParseRFC3339(s); err != nil {
		t.Fatal(err)
	}
	day := TruncateToDay(now)
	if day.Hour() != 0 {
		t.Fatal("expected midnight")
	}
}
