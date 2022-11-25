package ui

import (
	"lympach/data"

	"github.com/gen2brain/iup-go/iup"
)

func removeGP1(ih iup.Ihandle) int {
	RemovePerson("grandparent1")
	saveEvent("grandparent1")
	return iup.DEFAULT
}

func removeGP2(ih iup.Ihandle) int {
	RemovePerson("grandparent2")
	saveEvent("grandparent2")
	return iup.DEFAULT
}

func removeGP3(ih iup.Ihandle) int {
	RemovePerson("grandparent3")
	saveEvent("grandparent3")
	return iup.DEFAULT
}

func removeGP4(ih iup.Ihandle) int {
	RemovePerson("grandparent4")
	saveEvent("grandparent4")
	return iup.DEFAULT
}

func removeP1(ih iup.Ihandle) int {
	RemovePerson("parent1")
	saveEvent("parent1")
	return iup.DEFAULT
}

func removeP2(ih iup.Ihandle) int {
	RemovePerson("parent2")
	saveEvent("parent2")
	return iup.DEFAULT
}

func removeC(ih iup.Ihandle) int {
	RemovePerson("child")
	saveEvent("child")
	return iup.DEFAULT
}

func removePA(ih iup.Ihandle) int {
	RemovePerson("parentA")
	saveEvent("parentA")
	return iup.DEFAULT
}

func removeSA(ih iup.Ihandle) int {
	RemovePerson("spouseA")
	saveEvent("spouseA")
	return iup.DEFAULT
}

func removeCA(ih iup.Ihandle) int {
	RemovePerson("childA")
	saveEvent("childA")
	return iup.DEFAULT
}

func removePB(ih iup.Ihandle) int {
	RemovePerson("parentB")
	saveEvent("parentB")
	return iup.DEFAULT
}

func removeSB(ih iup.Ihandle) int {
	RemovePerson("spouseB")
	saveEvent("spouseB")
	return iup.DEFAULT
}

func removeCB(ih iup.Ihandle) int {
	RemovePerson("childB")
	saveEvent("childB")
	return iup.DEFAULT
}

func removePC(ih iup.Ihandle) int {
	RemovePerson("parentC")
	saveEvent("parentC")
	return iup.DEFAULT
}

func removeSC(ih iup.Ihandle) int {
	RemovePerson("spouseC")
	saveEvent("spouseC")
	return iup.DEFAULT
}

func removeCC(ih iup.Ihandle) int {
	RemovePerson("childC")
	saveEvent("childC")
	return iup.DEFAULT
}

func removePD(ih iup.Ihandle) int {
	RemovePerson("parentD")
	saveEvent("parentD")
	return iup.DEFAULT
}

func removeSD(ih iup.Ihandle) int {
	RemovePerson("spouseD")
	saveEvent("spouseD")
	return iup.DEFAULT
}

func removeCD(ih iup.Ihandle) int {
	RemovePerson("childD")
	saveEvent("childD")
	return iup.DEFAULT
}

func removeAll() {
	RemovePerson("grandparent1")
	RemovePerson("grandparent2")
	RemovePerson("grandparent3")
	RemovePerson("grandparent4")
	RemovePerson("parent1")
	RemovePerson("parent2")
	RemovePerson("child")
	RemovePerson("parentA")
	RemovePerson("spouseA")
	RemovePerson("childA")
	RemovePerson("parentB")
	RemovePerson("spouseB")
	RemovePerson("childB")
	RemovePerson("parentC")
	RemovePerson("spouseC")
	RemovePerson("childC")
	RemovePerson("parentD")
	RemovePerson("spouseD")
	RemovePerson("childD")
}

func resetEvent(ih iup.Ihandle) int {
	removeAll()
	_, _, year := data.SplitDate(iup.GetHandle("eventDate").GetAttribute("VALUE"))
	clearEvent()
	iup.GetHandle("eventDate").SetAttribute("VALUE", year)
	return iup.DEFAULT
}

func deleteEvent(ih iup.Ihandle) int {
	removeAll()
	deleteFullEvent()
	clearEvent()
	showPerson(iup.GetHandle("person"))
	updatePersonalEvents()
	return iup.DEFAULT
}
