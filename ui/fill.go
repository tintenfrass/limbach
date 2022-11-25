package ui

import (
	"lympach/data"
	"strconv"

	"github.com/gen2brain/iup-go/iup"
)

func SetPerson(handle string, pr data.PersonRecord) {
	valid, gn, fn, _, age1, age2, _, _ := data.GetPersonalData(pr.Xref)
	if !valid {
		RemovePerson(handle)
		return
	}
	hdl := iup.GetHandle(handle)

	hdl.SetAttribute("TITLE", utf82ui(gn+" "+fn+"\r\n("+age1+"-"+age2+")\r\n"+pr.Details.Type+"\r\n"+pr.Details.Place+"\r\n"+pr.Details.Msg))
	hdl.SetAttribute("VALUE", pr.Xref)
}

func RemovePerson(handle string) {
	iup.SetAttribute(iup.GetHandle(handle), "TITLE", "")
	iup.SetAttribute(iup.GetHandle(handle), "VALUE", "")
}

func updatePersonalEvents() {
	valid, gn, fn, sex, _, _, events, apps := data.GetPersonalData(validatePerson(iup.GetHandle("person").GetAttribute("VALUESTRING")))
	if !valid {
		return
	}

	iup.SetAttribute(iup.GetHandle("personalEvents"), "REMOVEITEM", "ALL")
	p := 1
	for i := len(events) - 1; i >= 0; i-- {
		iup.SetAttribute(iup.GetHandle("personalEvents"), strconv.Itoa(p), utf82ui(events[i]))
		p++
	}

	iup.SetAttribute(iup.GetHandle("eventList"), "REMOVEITEM", "ALL")
	p = 1
	for i := len(apps) - 1; i >= 0; i-- {
		iup.SetAttribute(iup.GetHandle("eventList"), strconv.Itoa(p), utf82ui(apps[i]))
		p++
	}

	iup.GetHandle("inputgname").SetAttribute("VALUE", utf82ui(gn))
	iup.GetHandle("inputfname").SetAttribute("VALUE", utf82ui(fn))
	iup.GetHandle("sex").SetAttribute("VALUE", sex)

	updateOccupations(events)
	updatePlaces()

	return
}

func updateFamily() {
	valid, families := data.GetFamily(validatePerson(iup.GetHandle("person").GetAttribute("VALUESTRING")))
	if !valid {
		return
	}

	iup.SetAttribute(iup.GetHandle("family"), "REMOVEITEM", "ALL")
	for i, family := range families {
		iup.SetAttribute(iup.GetHandle("family"), strconv.Itoa(i+1), utf82ui(family))
	}
}
