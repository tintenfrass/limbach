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
		iup.SetAttribute(iup.GetHandle("places"), strconv.Itoa(k+1), utf82ui(place))
	}
}

func updateOccupations(events []string) {
	iup.GetHandle("occupations").SetAttribute("REMOVEITEM", "ALL")
	occ := make(map[string]struct{})
	for _, event := range events {
		parts := strings.Split(event, "|")
		if len(parts) < 4 {
			continue
		}
		parts[3] = strings.TrimSpace(parts[3])
		if len(parts[3]) > 0 {
			occ[parts[3]] = struct{}{}
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

func updatePlaces() {
	storage := data.PlaceStorage
	for _, place := range data.Data.Places {
		delete(storage, place)
	}
	places := []string{}
	for place, _ := range storage {
		if len(place) > 0 {
			places = append(places, place)
		}
	}
	sort.Strings(places)
	for k, place := range places {
		iup.SetAttribute(iup.GetHandle("places2"), strconv.Itoa(k+1), utf82ui(place))
	}
}
