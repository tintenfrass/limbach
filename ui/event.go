package ui

import (
	"lympach/data"
	"strconv"
	"strings"

	"github.com/gen2brain/iup-go/iup"
)

func saveEvent(handle string) {
	eventId, err := strconv.Atoi(iup.GetHandle("eventId").GetAttribute("TITLE"))
	if err != nil {
		//New Event
		if data.Data.CounterE == 0 {
			data.Data.CounterE++
		}
		eventId = data.Data.CounterE
		iup.GetHandle("eventId").SetAttribute("TITLE", eventId)
		data.Data.CounterE++
	}

	storage := data.Data.Events[eventId]
	if storage.Additionals == nil {
		storage.Additionals = make(map[string]data.AdditionalRecord)
		storage.Additionals["A"] = data.AdditionalRecord{}
		storage.Additionals["B"] = data.AdditionalRecord{}
		storage.Additionals["C"] = data.AdditionalRecord{}
		storage.Additionals["D"] = data.AdditionalRecord{}
	}

	ids := []int{}
	value := ""
	rec := data.PersonRecord{}
	details := data.Details{}
	if handle == "marriedDetails" {
		details = getSelectedDetails()
		details.Type = "Trauung"
		ids = append(ids, storage.Parent1.Xref, storage.Parent2.Xref)
	} else if handle == "eventDate" || handle == "married" {
		value = ui2utf8(iup.GetHandle(handle).GetAttribute("VALUE"))
	} else {
		value = ui2utf8(iup.GetHandle(handle).GetAttribute("TITLE"))
		xref, err := strconv.Atoi(iup.GetHandle(handle).GetAttribute("VALUE"))
		if err != nil {
			xref = 0
		}
		values := strings.Split(value, "\r\n")

		if xref > 0 && len(values) > 4 {
			rec = data.PersonRecord{
				Xref: xref,
				Details: data.Details{
					Type:  values[2],
					Place: values[3],
					Msg:   values[4],
				},
			}
			ids = append(ids, xref)
		}
	}

	switch handle {
	case "eventDate":
		storage.Date = value
		ids = append(ids,
			storage.Grandparent1.Xref, storage.Grandparent2.Xref, storage.Grandparent3.Xref, storage.Grandparent4.Xref,
			storage.Parent1.Xref, storage.Parent2.Xref, storage.Child.Xref,
		)
		for _, ak := range data.AdditonalKeys {
			ids = append(ids, storage.Additionals[ak].Parent.Xref, storage.Additionals[ak].Spouse.Xref, storage.Additionals[ak].Child.Xref)
		}
	case "married":
		storage.Married = value == "ON"
		ids = append(ids, storage.Parent1.Xref, storage.Parent2.Xref, storage.Child.Xref)
	case "marriedDetails":
		storage.Details = details
		ids = append(ids, storage.Parent1.Xref, storage.Parent2.Xref, storage.Child.Xref)
	case "grandparent1":
		ids = append(ids, storage.Grandparent1.Xref, storage.Grandparent2.Xref, storage.Parent1.Xref)
		storage.Grandparent1 = rec
	case "grandparent2":
		ids = append(ids, storage.Grandparent1.Xref, storage.Grandparent2.Xref, storage.Parent1.Xref)
		storage.Grandparent2 = rec
	case "grandparent3":
		ids = append(ids, storage.Grandparent3.Xref, storage.Grandparent4.Xref, storage.Parent2.Xref)
		storage.Grandparent3 = rec
	case "grandparent4":
		ids = append(ids, storage.Grandparent3.Xref, storage.Grandparent4.Xref, storage.Parent2.Xref)
		storage.Grandparent4 = rec
	case "parent1":
		ids = append(ids, storage.Child.Xref, storage.Parent1.Xref, storage.Parent2.Xref, storage.Grandparent1.Xref, storage.Grandparent2.Xref)
		storage.Parent1 = rec
	case "parent2":
		ids = append(ids, storage.Child.Xref, storage.Parent1.Xref, storage.Parent2.Xref, storage.Grandparent3.Xref, storage.Grandparent3.Xref)
		storage.Parent2 = rec
	case "child":
		ids = append(ids, storage.Child.Xref, storage.Parent1.Xref, storage.Parent2.Xref)
		storage.Child = rec
	case "parentA":
		d := storage.Additionals["A"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Parent = rec
		storage.Additionals["A"] = d
	case "spouseA":
		d := storage.Additionals["A"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Spouse = rec
		storage.Additionals["A"] = d
	case "childA":
		d := storage.Additionals["A"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Child = rec
		storage.Additionals["A"] = d
	case "parentB":
		d := storage.Additionals["B"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Parent = rec
		storage.Additionals["B"] = d
	case "spouseB":
		d := storage.Additionals["B"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Spouse = rec
		storage.Additionals["B"] = d
	case "childB":
		d := storage.Additionals["B"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Child = rec
		storage.Additionals["B"] = d
	case "parentC":
		d := storage.Additionals["C"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Parent = rec
		storage.Additionals["C"] = d
	case "spouseC":
		d := storage.Additionals["C"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Spouse = rec
		storage.Additionals["C"] = d
	case "childC":
		d := storage.Additionals["C"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Child = rec
		storage.Additionals["C"] = d
	case "parentD":
		d := storage.Additionals["D"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Parent = rec
		storage.Additionals["D"] = d
	case "spouseD":
		d := storage.Additionals["D"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Spouse = rec
		storage.Additionals["D"] = d
	case "childD":
		d := storage.Additionals["D"]
		ids = append(ids, d.Parent.Xref, d.Spouse.Xref, d.Child.Xref)
		d.Child = rec
		storage.Additionals["D"] = d
	}

	data.Data.Events[eventId] = storage
	updatePersonalEvents()
	if handle != "eventDate" {
		if len(ids) > 0 {
			data.UpdatePersonalDataStorage(ids)
		}
		data.UpdateStorage()
		loadFullEvent(eventId)
	}
	showPerson(iup.GetHandle("person"))
}

func loadFullEvent(id int) {
	iup.GetHandle("eventId").SetAttribute("TITLE", id)
	if ev, ok := data.Data.Events[id]; ok {
		iup.GetHandle("eventDate").SetAttribute("VALUE", utf82ui(ev.Date))

		married := 0
		if ev.Married || ev.Details.Type == "Trauung" {
			married = 1
		}
		iup.GetHandle("married").SetAttribute("VALUE", married)
		iup.GetHandle("marriedPlace").SetAttribute("TITLE", ev.Details.Place)
		iup.GetHandle("marriedMsg").SetAttribute("TITLE", ev.Details.Msg)

		SetPerson("grandparent1", ev.Grandparent1)
		SetPerson("grandparent2", ev.Grandparent2)
		SetPerson("grandparent3", ev.Grandparent3)
		SetPerson("grandparent4", ev.Grandparent4)
		SetPerson("parent1", ev.Parent1)
		SetPerson("parent2", ev.Parent2)
		SetPerson("child", ev.Child)

		for _, key := range data.AdditonalKeys {
			RemovePerson("parent" + key)
			RemovePerson("spouse" + key)
			RemovePerson("child" + key)
			if add, ok := ev.Additionals[key]; ok {
				SetPerson("parent"+key, add.Parent)
				SetPerson("spouse"+key, add.Spouse)
				SetPerson("child"+key, add.Child)
			}
		}
	}
}

func deleteFullEvent() {
	eventId, err := strconv.Atoi(iup.GetHandle("eventId").GetAttribute("TITLE"))
	if err != nil {
		return
	}
	ids := []int{}
	ev := data.Data.Events[eventId]
	ids = append(ids,
		ev.Grandparent1.Xref, ev.Grandparent2.Xref, ev.Grandparent3.Xref, ev.Grandparent4.Xref,
		ev.Parent1.Xref, ev.Parent2.Xref, ev.Child.Xref,
	)
	for _, ak := range data.AdditonalKeys {
		ids = append(ids, ev.Additionals[ak].Parent.Xref, ev.Additionals[ak].Spouse.Xref, ev.Additionals[ak].Child.Xref)
	}
	delete(data.Data.Events, eventId)
	data.UpdatePersonalDataStorage(ids)
	data.UpdateStorage()
}

func clearEvent() {
	iup.GetHandle("eventId").SetAttribute("TITLE", "")
	iup.GetHandle("eventDate").SetAttribute("VALUE", "")
	iup.GetHandle("married").SetAttribute("VALUE", 0)
	iup.GetHandle("marriedPlace").SetAttribute("TITLE", "")
	iup.GetHandle("marriedMsg").SetAttribute("TITLE", "")
}
