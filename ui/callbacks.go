package ui

import (
	"lympach/data"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/iup-go/iup"
)

func loadEvent(ih iup.Ihandle) int {
	val := iup.GetHandle("personalEvents").GetAttribute("VALUESTRING")
	eId, _, _, _, _ := data.SplitEvent(val)
	if len(eId) < 3 {
		return iup.DEFAULT
	}
	id, err := strconv.Atoi(eId[2:])
	if err != nil {
		return iup.DEFAULT
	}

	loadFullEvent(id)

	return iup.DEFAULT
}

func changeDate(ih iup.Ihandle) int {
	saveEvent("eventDate")

	return iup.DEFAULT
}

func changeMarried(ih iup.Ihandle) int {
	saveEvent("married")

	return iup.DEFAULT
}

func marriedDetails(ih iup.Ihandle) int {
	saveEvent("marriedDetails")
	return iup.DEFAULT
}

func search(ih iup.Ihandle) int {
	val := ih.GetAttribute("VALUE")

	persons := data.FindPerson(strings.TrimSpace((ui2utf8(val))))

	iup.SetAttribute(iup.GetHandle("person"), "REMOVEITEM", "ALL")
	if len(val) > 0 {
		for i, person := range persons {
			iup.SetAttribute(iup.GetHandle("person"), strconv.Itoa(i+1), utf82ui(person.Data))
			if i > 100 {
				break
			}
		}
	}

	createColorBoxes(persons)

	return iup.DEFAULT
}

func gname(ih iup.Ihandle) int {
	data.CheckVnName(ui2utf8(ih.GetAttribute("VALUE")))
	return iup.DEFAULT
}

func showPerson(ih iup.Ihandle) int {
	updatePersonalEvents()
	updateFamily()

	return iup.DEFAULT
}

func selectPerson(ih iup.Ihandle) int {
	value := iup.GetHandle("family").GetAttribute("VALUESTRING")
	pos := strings.Index(value, "I-")
	if pos >= 0 && len(value) > pos+2 {
		iup.GetHandle("search").SetAttribute("VALUE", value[pos+2:])
		search(iup.GetHandle("search"))
		iup.GetHandle("person").SetAttribute("VALUE", 1)
		showPerson(iup.GetHandle("person"))
	}

	return iup.DEFAULT
}

func changePerson(ih iup.Ihandle) int {
	id := validatePerson(iup.GetHandle("person").GetAttribute("VALUESTRING"))
	if id == 0 {
		return iup.DEFAULT
	}
	gn := ui2utf8(iup.GetHandle("inputgname").GetAttribute("VALUE"))
	fn := ui2utf8(iup.GetHandle("inputfname").GetAttribute("VALUE"))
	sex := iup.GetHandle("sex").GetAttribute("VALUE")
	data.ChangePerson(id, gn, fn, sex)

	iup.GetHandle("search").SetAttribute("VALUE", id)
	search(iup.GetHandle("search"))
	iup.GetHandle("person").SetAttribute("VALUE", 1)
	showPerson(iup.GetHandle("person"))

	return iup.DEFAULT
}

func newPerson(ih iup.Ihandle) int {
	gn := ui2utf8(iup.GetHandle("inputgname").GetAttribute("VALUE"))
	fn := ui2utf8(iup.GetHandle("inputfname").GetAttribute("VALUE"))
	sex := iup.GetHandle("sex").GetAttribute("VALUE")
	id := data.NewPerson(gn, fn, sex)

	iup.GetHandle("search").SetAttribute("VALUE", id)
	search(iup.GetHandle("search"))
	iup.GetHandle("person").SetAttribute("VALUE", 1)
	showPerson(iup.GetHandle("person"))

	return iup.DEFAULT
}

func deletePerson(ih iup.Ihandle) int {
	id := validatePerson(iup.GetHandle("person").GetAttribute("VALUESTRING"))
	if id == 0 {
		return iup.DEFAULT
	}
	valid, _, _, _, _, _, ev := data.GetPersonalData(id)
	if valid && len(ev) == 0 {
		data.DeletePerson(id)
		data.UpdatePersonalDataStorage([]int{id})

		iup.GetHandle("inputgname").SetAttribute("VALUE", "")
		iup.GetHandle("inputfname").SetAttribute("VALUE", "")
		iup.GetHandle("search").SetAttribute("VALUE", "")
		search(iup.GetHandle("search"))

		data.UpdateStorage()
	}

	return iup.DEFAULT
}

func occupation(ih iup.Ihandle) int {
	iup.GetHandle("msg").SetAttribute("VALUE", ih.GetAttribute("VALUESTRING"))

	return iup.DEFAULT
}

func placesTop(ih iup.Ihandle) int {
	iup.GetHandle("placesBottom").SetAttribute("VALUE", 0)
	iup.GetHandle("placesMiddle").SetAttribute("VALUE", ih.GetAttribute("VALUESTRING"))

	return iup.DEFAULT
}

func placeMiddle(ih iup.Ihandle) int {
	iup.GetHandle("placesTop").SetAttribute("VALUE", 0)
	iup.GetHandle("placesBottom").SetAttribute("VALUE", 0)

	return iup.DEFAULT
}

func placeBottom(ih iup.Ihandle) int {
	iup.GetHandle("placesTop").SetAttribute("VALUE", 0)
	iup.GetHandle("placesMiddle").SetAttribute("VALUE", ih.GetAttribute("VALUESTRING"))

	return iup.DEFAULT
}

func msgChange(ih iup.Ihandle) int {
	iup.GetHandle("desc").SetAttribute("TITLE", "")
	msg := ih.GetAttribute("VALUE")
	age := ""
	if len(msg) > 1 && msg[1:2] == "J" {
		age = msg[:1]
	} else if len(msg) > 2 && msg[2:3] == "J" {
		age = msg[:2]
	}
	ageInt, err := strconv.Atoi(age)
	if err == nil && ageInt > 0 {
		now := iup.GetHandle("eventDate").GetAttribute("VALUE")
		parsed, err := time.Parse(dateLayout, now)
		if err == nil {
			iup.GetHandle("desc").SetAttribute("TITLE", "ca. "+strconv.Itoa(parsed.Year()-ageInt))
		}
	}

	return iup.DEFAULT
}

func short(ih iup.Ihandle) int {
	prefix := ui2utf8(ih.GetAttribute("VALUESTRING"))
	if len(prefix) == 0 {
		prefix = " "
	}
	updatePlaces(prefix)
	updateQuick(prefix)

	return iup.DEFAULT
}

func quickName(ih iup.Ihandle) int {
	search(iup.GetHandle("search").SetAttribute("VALUE", "? "+ui2utf8(ih.GetAttribute("VALUESTRING"))))

	return iup.DEFAULT
}

func exit(ih iup.Ihandle) int {
	data.Save()
	return iup.CLOSE
}

func export(ih iup.Ihandle) int {
	data.Export()
	return iup.DEFAULT
}
