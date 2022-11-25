package ui

import (
	"lympach/data"

	"github.com/gen2brain/iup-go/iup"
)

func getSelectedPerson() int {
	return validatePerson(iup.GetHandle("person").GetAttribute("VALUESTRING"))
}

func getSelectedDetails() (details data.Details) {
	details = data.Details{
		Type:  ui2utf8(iup.GetHandle("eventtypes").GetAttribute("VALUESTRING")),
		Msg:   ui2utf8(iup.GetHandle("msg").GetAttribute("VALUE")),
		Place: ui2utf8(iup.GetHandle("places1").GetAttribute("VALUE")),
	}
	iup.GetHandle("msg").SetAttribute("VALUE", "")
	return
}

func setPerson(handle string) {
	xref := getSelectedPerson()
	if xref > 0 {
		SetPerson(handle, data.PersonRecord{
			Xref:    xref,
			Details: getSelectedDetails(),
		})
		saveEvent(handle)
	}
}
