package image

import "strings"

func imageLoading(value string) string {
	switch strings.TrimSpace(value) {
	case "eager":
		return "eager"
	default:
		return "lazy"
	}
}
