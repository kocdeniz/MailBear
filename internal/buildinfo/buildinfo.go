package buildinfo

import "strings"

// These values are set at build time with -ldflags -X.
var (
	Version   = "dev"
	BuildDate = "unknown"
)

func IntroLabel() string {
	v := strings.TrimSpace(Version)
	if v == "" {
		v = "dev"
	}
	b := strings.TrimSpace(BuildDate)
	if b == "" || b == "unknown" {
		return "v" + v
	}
	return "v" + v + "  build " + b
}
