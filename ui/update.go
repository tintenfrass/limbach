package ui

import (
	"lympach/data"
	"sort"
	"strconv"
	"strings"

	"github.com/gen2brain/iup-go/iup"
)

func initFixedEventtypes() {
	for k, eventtype := range data.Data.Eventtypes {
		iup.SetAttribute(iup.GetHandle("eventtypes"), strconv.Itoa(k+1), utf82ui(eventtype))
	}
}

func initFixedPlaces() {
	for k, place := range data.Data.Places {
		iup.SetAttribute(iup.GetHandle("placesTop"), strconv.Itoa(k+1), utf82ui(place))
	}
}

func updateOccupations(events []string) {
	iup.GetHandle("occupations").SetAttribute("REMOVEITEM", "ALL")
	occ := make(map[string]struct{})
	for _, event := range events {
		_, _, _, _, msg := data.SplitEvent(event)
		if len(msg) > 0 {
			occ[msg] = struct{}{}
		}
	}

	o := []string{}
	for oc := range occ {
		o = append(o, oc)
	}
	sort.Strings(o)
	for i, oc := range o {
		i++
		iup.GetHandle("occupations").SetAttribute(strconv.Itoa(i), utf82ui(oc))
	}
}

func updatePlaces(prefix string) {
	storage := data.PlaceStorage
	for _, pl := range data.Data.Places {
		delete(storage, pl)
	}
	places := []string{}
	for pl, _ := range storage {
		if len(pl) > 0 && (prefix == " " || strings.HasPrefix(pl, prefix)) {
			places = append(places, pl)
		}
	}
	sort.Strings(places)
	iup.SetAttribute(iup.GetHandle("placesBottom"), "REMOVEITEM", "ALL")
	for k, place := range places {
		iup.SetAttribute(iup.GetHandle("placesBottom"), strconv.Itoa(k+1), utf82ui(place))
	}
}

func updateQuick(prefix string) {
	list := make(map[string]bool)
	for _, indi := range data.Data.Individual {
		fn := indi.FName
		if len(fn) > 0 {
			if prefix == " " || strings.HasPrefix(fn, prefix) {
				list[fn] = true
			}
		}
	}
	slice := []string{}
	for key := range list {
		slice = append(slice, key)
	}

	sort.Strings(slice)
	slice = append([]string{"?"}, slice...)

	iup.SetAttribute(iup.GetHandle("quick"), "REMOVEITEM", "ALL")
	for k, name := range slice {
		iup.SetAttribute(iup.GetHandle("quick"), strconv.Itoa(k+1), utf82ui(name))
	}
}
