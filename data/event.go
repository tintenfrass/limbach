package data

import "strings"

func SplitEvent(ev string) (id, date, typ, place, msg string) {
	parts := strings.Split(ev, "|")
	if len(parts) > 4 {
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2]), strings.TrimSpace(parts[3]), strings.TrimSpace(parts[4])
	}
	return
}
