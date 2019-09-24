package fingerprint

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)


func HashUserAgent(v string) string {
	sum := sha256.Sum256([]byte("user_agent:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashLanguage(v string) string {
	sum := sha256.Sum256([]byte("language:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashTimezone(v string) string {
	sum := sha256.Sum256([]byte("timezone:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashScreen(v string) string {
	sum := sha256.Sum256([]byte("screen:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashPlatform(v string) string {
	sum := sha256.Sum256([]byte("platform:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashVendor(v string) string {
	sum := sha256.Sum256([]byte("vendor:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashWebgl(v string) string {
	sum := sha256.Sum256([]byte("webgl:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashCanvas(v string) string {
	sum := sha256.Sum256([]byte("canvas:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashAudio(v string) string {
	sum := sha256.Sum256([]byte("audio:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashFonts(v string) string {
	sum := sha256.Sum256([]byte("fonts:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashPlugins(v string) string {
	sum := sha256.Sum256([]byte("plugins:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashHardwareConcurrency(v string) string {
	sum := sha256.Sum256([]byte("hardware_concurrency:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashDeviceMemory(v string) string {
	sum := sha256.Sum256([]byte("device_memory:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashTouch(v string) string {
	sum := sha256.Sum256([]byte("touch:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashCookies(v string) string {
	sum := sha256.Sum256([]byte("cookies:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashLocalStorage(v string) string {
	sum := sha256.Sum256([]byte("local_storage:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashSessionStorage(v string) string {
	sum := sha256.Sum256([]byte("session_storage:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashIndexedDb(v string) string {
	sum := sha256.Sum256([]byte("indexed_db:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashDoNotTrack(v string) string {
	sum := sha256.Sum256([]byte("do_not_track:"+v))
	return hex.EncodeToString(sum[:8])
}

func HashAdblock(v string) string {
	sum := sha256.Sum256([]byte("adblock:"+v))
	return hex.EncodeToString(sum[:8])
}


func Composite(parts map[string]string) string {
	keys := make([]string, 0, len(parts))
	for k := range parts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(parts[k])
		b.WriteByte('|')
	}
	sum := sha256.Sum256([]byte(b.String()))
	return hex.EncodeToString(sum[:])
}
