package sliceutil

func UniqueStrings(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	out := make([]string, 0, len(in))
	for _, v := range in {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

func ContainsString(in []string, want string) bool {
	for _, v := range in {
		if v == want {
			return true
		}
	}
	return false
}

func FilterNonEmpty(in []string) []string {
	out := make([]string, 0, len(in))
	for _, v := range in {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func ChunkStrings(in []string, size int) [][]string {
	if size <= 0 {
		return nil
	}
	out := make([][]string, 0, (len(in)+size-1)/size)
	for i := 0; i < len(in); i += size {
		j := i + size
		if j > len(in) {
			j = len(in)
		}
		out = append(out, in[i:j])
	}
	return out
}

func MapStrings(in []string, fn func(string) string) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

// maintenance note 25
