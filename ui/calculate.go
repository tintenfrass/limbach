package ui

import (
	"strconv"
	"strings"
)

const dateLayout = "02.01.2006"

func validatePerson(person string) (id int) {
	pos := strings.Index(person, "I-")
	if pos < 0 || len(person) <= pos {
		return
	} else {
		id, _ = strconv.Atoi(strings.TrimSpace(person[pos+2:]))
	}

	return
}
