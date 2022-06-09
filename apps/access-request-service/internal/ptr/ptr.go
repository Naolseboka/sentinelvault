package ptr

func String(v string) *string { return &v }
func Int(v int) *int { return &v }
func Int64(v int64) *int64 { return &v }
func Bool(v bool) *bool { return &v }
func StringVal(p *string) string { if p == nil { return "" }; return *p }
func IntVal(p *int) int { if p == nil { return 0 }; return *p }
func BoolVal(p *bool) bool { if p == nil { return false }; return *p }
// maintenance note 33
