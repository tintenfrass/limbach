package data

import (
	"strconv"
	"strings"
)

type customSort []ResultData

func (s customSort) Len() int {
	return len(s)
}
func (s customSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s customSort) Less(i, j int) bool {
	a1, a2 := split(s[i].Data)
	b1, b2 := split(s[j].Data)

	if a1 == b1 {
		return a2 > b2
	}

	return strings.Compare(s[i].Data, s[j].Data) < 0
}

func split(person string) (name string, id int) {
	posI := strings.Index(person, "I-")
	if posI < 0 || len(person) <= posI {
		return
	} else {
		pos := strings.Index(person, "|")
		name = strings.TrimSpace(person[:pos])
		id, _ = strconv.Atoi(strings.TrimSpace(person[posI+2:]))
	}
	return
}

type eventSort []personalEvent

func (s eventSort) Len() int {
	return len(s)
}
func (s eventSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s eventSort) Less(i, j int) bool {
	d1, m1, y1 := SplitDate(s[i].date)
	d2, m2, y2 := SplitDate(s[j].date)

	if y1 == y2 {
		if m1 == m2 {
			return d2 > d1
		} else {
			return m2 > m1
		}
	} else {
		return y2 > y1
	}
}

func SplitDate(date string) (d, m, y int) {
	parts := strings.Split(date, ".")
	switch len(parts) {
	case 1:
		y, _ = strconv.Atoi(parts[0])
	case 2:
		m, _ = strconv.Atoi(parts[0])
		y, _ = strconv.Atoi(parts[1])
	case 3:
		d, _ = strconv.Atoi(parts[0])
		m, _ = strconv.Atoi(parts[1])
		y, _ = strconv.Atoi(parts[2])
	}
	return
}

func dateAfter(date1, date2 string) bool {
	d1, m1, y1 := SplitDate(date1)
	d2, m2, y2 := SplitDate(date2)

	if y1 == y2 {
		if m1 == m2 {
			return d1 > d2
		} else {
			return m1 > m2
		}
	} else {
		return y1 > y2
	}
}
