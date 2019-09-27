package passwordpolicy

import (
	"fmt"
	"unicode"
)

type Policy struct {
	MinLength int
	RequireUpper bool
	RequireLower bool
	RequireDigit bool
	RequireSpecial bool
	MaxLength int
	DisallowWhitespace bool
}

func Default() Policy {
	return Policy{MinLength: 12, MaxLength: 128, RequireUpper: true, RequireLower: true, RequireDigit: true, RequireSpecial: true, DisallowWhitespace: true}
}

func (p Policy) Validate(pw string) []string {
	errs := make([]string, 0)
	if len(pw) < p.MinLength {
		errs = append(errs, fmt.Sprintf("min length %d", p.MinLength))
	}
	if p.MaxLength > 0 && len(pw) > p.MaxLength {
		errs = append(errs, fmt.Sprintf("max length %d", p.MaxLength))
	}
	var upper, lower, digit, special, space bool
	for _, r := range pw {
		switch {
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsLower(r):
			lower = true
		case unicode.IsDigit(r):
			digit = true
		case unicode.IsSpace(r):
			space = true
		default:
			special = true
		}
	}
	if p.RequireUpper && !upper {
		errs = append(errs, "uppercase required")
	}
	if p.RequireLower && !lower {
		errs = append(errs, "lowercase required")
	}
	if p.RequireDigit && !digit {
		errs = append(errs, "digit required")
	}
	if p.RequireSpecial && !special {
		errs = append(errs, "special required")
	}
	if p.DisallowWhitespace && space {
		errs = append(errs, "whitespace not allowed")
	}
	return errs
}

func (p Policy) OK(pw string) bool { return len(p.Validate(pw)) == 0 }
