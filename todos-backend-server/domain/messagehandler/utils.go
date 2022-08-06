package messagehandler

import "strings"

func sanitizedTitle(title string) string {
	return strings.TrimSpace(title)
}
