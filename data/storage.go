package data

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var FamilyStorage map[string]Familie
var PlaceStorage map[string]struct{}
var PersonalDateStorage map[int]PersonalData

type PersonalData struct {
	gn, fn, sex, age1, age2 string
	events                  []string
}

type Familie struct {
	Father  int
	Mother  int
	Childen []int
	Married bool
}

type span struct {
	begin string
	end   string
}

func createPersonalEvents(id int) (personalEvents map[int][]personalEvent) {
	personalEvents = make(map[int][]personalEvent)

	for eId, ev := range Data.Events {
		PlaceStorage[ev.Details.Place] = struct{}{}
		if id == 0 || ev.Grandparent1.Xref == id {
			personalEvents[ev.Grandparent1.Xref] = append(personalEvents[ev.Grandparent1.Xref], personalEvent{eId, ev.Date, ev.Grandparent1.Details})
		}
		PlaceStorage[ev.Grandparent1.Details.Place] = struct{}{}

		if id == 0 || ev.Grandparent2.Xref == id {
			personalEvents[ev.Grandparent2.Xref] = append(personalEvents[ev.Grandparent2.Xref], personalEvent{eId, ev.Date, ev.Grandparent2.Details})
		}
		PlaceStorage[ev.Grandparent2.Details.Place] = struct{}{}

		if id == 0 || ev.Grandparent3.Xref == id {
			personalEvents[ev.Grandparent3.Xref] = append(personalEvents[ev.Grandparent3.Xref], personalEvent{eId, ev.Date, ev.Grandparent3.Details})
		}
		PlaceStorage[ev.Grandparent3.Details.Place] = struct{}{}

		if id == 0 || ev.Grandparent4.Xref == id {
			personalEvents[ev.Grandparent4.Xref] = append(personalEvents[ev.Grandparent4.Xref], personalEvent{eId, ev.Date, ev.Grandparent4.Details})
		}
		PlaceStorage[ev.Grandparent4.Details.Place] = struct{}{}

		if id == 0 || ev.Parent1.Xref == id {
			personalEvents[ev.Parent1.Xref] = append(personalEvents[ev.Parent1.Xref], personalEvent{eId, ev.Date, ev.Parent1.Details})
			if ev.Parent2.Xref > 0 && ev.Parent2.Details.Type == "Tod" {
				e := ev.Parent2.Details
				e.Type += " der Partnerin"
				personalEvents[ev.Parent1.Xref] = append(personalEvents[ev.Parent1.Xref], personalEvent{eId, ev.Date, e})
			}
		}
		PlaceStorage[ev.Parent1.Details.Place] = struct{}{}

		if id == 0 || ev.Parent2.Xref == id {
			personalEvents[ev.Parent2.Xref] = append(personalEvents[ev.Parent2.Xref], personalEvent{eId, ev.Date, ev.Parent2.Details})
			if ev.Parent1.Xref > 0 && ev.Parent1.Details.Type == "Tod" {
				e := ev.Parent1.Details
				e.Type += " des Partners"
				personalEvents[ev.Parent2.Xref] = append(personalEvents[ev.Parent2.Xref], personalEvent{eId, ev.Date, e})
			}
		}
		PlaceStorage[ev.Parent2.Details.Place] = struct{}{}

		if id == 0 || ev.Parent1.Xref == id || ev.Parent2.Xref == id {
			if ev.Details.Type == "Trauung" {
				personalEvents[ev.Parent1.Xref] = append(personalEvents[ev.Parent1.Xref], personalEvent{eId, ev.Date, ev.Details})
				personalEvents[ev.Parent2.Xref] = append(personalEvents[ev.Parent2.Xref], personalEvent{eId, ev.Date, ev.Details})
			}
			if ev.Child.Xref > 0 && len(ev.Child.Details.Type) > 0 && len(ev.Date) > 0 {
				e := ev.Child.Details
				e.Type += " eines Kindes"
				personalEvents[ev.Parent1.Xref] = append(personalEvents[ev.Parent1.Xref], personalEvent{eId, ev.Date, e})
				personalEvents[ev.Parent2.Xref] = append(personalEvents[ev.Parent2.Xref], personalEvent{eId, ev.Date, e})
			}
		}
		if id == 0 || ev.Child.Xref == id {
			personalEvents[ev.Child.Xref] = append(personalEvents[ev.Child.Xref], personalEvent{eId, ev.Date, ev.Child.Details})
		}
		PlaceStorage[ev.Child.Details.Place] = struct{}{}

		for _, key := range AdditonalKeys {
			if id == 0 || ev.Additionals[key].Parent.Xref == id {
				personalEvents[ev.Additionals[key].Parent.Xref] = append(personalEvents[ev.Additionals[key].Parent.Xref], personalEvent{eId, ev.Date, ev.Additionals[key].Parent.Details})
			}
			PlaceStorage[ev.Additionals[key].Parent.Details.Place] = struct{}{}

			if id == 0 || ev.Additionals[key].Spouse.Xref == id {
				personalEvents[ev.Additionals[key].Spouse.Xref] = append(personalEvents[ev.Additionals[key].Spouse.Xref], personalEvent{eId, ev.Date, ev.Additionals[key].Spouse.Details})
			}
			PlaceStorage[ev.Additionals[key].Spouse.Details.Place] = struct{}{}

			if id == 0 || ev.Additionals[key].Child.Xref == id {
				personalEvents[ev.Additionals[key].Child.Xref] = append(personalEvents[ev.Additionals[key].Child.Xref], personalEvent{eId, ev.Date, ev.Additionals[key].Child.Details})
			}
			PlaceStorage[ev.Additionals[key].Child.Details.Place] = struct{}{}
		}
	}

	return
}

func createPersonalData(id int, personalEvents []personalEvent) (valid bool, gn, fn, sex, age1, age2 string, events []string) {
	if val, ok := Data.Individual[id]; ok {
		gn = val.GName
		fn = val.FName
		sex = val.Sex
		valid = true
	}

	if !valid {
		return
	}

	sort.Sort(eventSort(personalEvents))
	for _, pe := range personalEvents {
		if len(pe.details.Type) == 0 {
			continue
		}
		events = append(events, fmt.Sprintf("E-%d | %s | %s | %s | %s", pe.eId, pe.date, pe.details.Type, pe.details.Place, pe.details.Msg))
		if pe.details.Type == "Taufe" || pe.details.Type == "Geburt" {
			age1 = pe.date
		} else if pe.details.Type == "Tod" {
			age2 = pe.date
			if age1 == "" {
				age := ""
				if len(pe.details.Msg) > 1 && pe.details.Msg[1:2] == "J" {
					age = pe.details.Msg[:1]
				} else if len(pe.details.Msg) > 2 && pe.details.Msg[2:3] == "J" {
					age = pe.details.Msg[:2]
				}
				ageInt, err := strconv.Atoi(age)
				if err == nil && ageInt > 0 {
					parsed, err := time.Parse(dateLayout, pe.date)
					if err != nil {
						return
					}
					if parsed.Year() > 0 {
						age1 = "ca. " + strconv.Itoa(parsed.Year()-ageInt)
					}
				}
			}
		} else if pe.details.Type == "Weiland" && len(age2) == 0 {
			age2 = "Weiland"
		}
	}

	return
}

func UpdatePersonalDataStorage(ids []int) {
	if PersonalDateStorage == nil {
		PersonalDateStorage = make(map[int]PersonalData)
	}

	if len(ids) == 0 {
		allPersonalEvents := createPersonalEvents(0)
		for _, indi := range Data.Individual {
			valid, gn, fn, sex, age1, age2, evs := createPersonalData(indi.Xref, allPersonalEvents[indi.Xref])
			if !valid {
				delete(PersonalDateStorage, indi.Xref)
				continue
			}
			pd := PersonalData{
				gn:     gn,
				fn:     fn,
				sex:    sex,
				age1:   age1,
				age2:   age2,
				events: evs,
			}
			PersonalDateStorage[indi.Xref] = pd
		}
	} else {
		for _, id := range ids {
			if id == 0 {
				continue
			}
			valid, gn, fn, sex, age1, age2, evs := createPersonalData(id, createPersonalEvents(id)[id])
			if !valid {
				delete(PersonalDateStorage, id)
				continue
			}
			pd := PersonalData{
				gn:     gn,
				fn:     fn,
				sex:    sex,
				age1:   age1,
				age2:   age2,
				events: evs,
			}
			PersonalDateStorage[id] = pd
		}
	}
}

func UpdateStorage() {
	children := make(map[int]map[int]struct{}) //childId => parentIds
	marriedWith := make(map[int]map[int]string)
	spans := make(map[int]map[int]span) //fatherId => list of partners

	for _, ev := range Data.Events {
		if ev.Parent1.Xref > 0 && ev.Child.Xref > 0 {
			if children[ev.Child.Xref] == nil {
				children[ev.Child.Xref] = make(map[int]struct{})
			}
			children[ev.Child.Xref][ev.Parent1.Xref] = struct{}{}
		}
		if ev.Parent2.Xref > 0 && ev.Child.Xref > 0 {
			if children[ev.Child.Xref] == nil {
				children[ev.Child.Xref] = make(map[int]struct{})
			}
			children[ev.Child.Xref][ev.Parent2.Xref] = struct{}{}
		}
		if ev.Grandparent1.Xref > 0 && ev.Parent1.Xref > 0 {
			if children[ev.Parent1.Xref] == nil {
				children[ev.Parent1.Xref] = make(map[int]struct{})
			}
			children[ev.Parent1.Xref][ev.Grandparent1.Xref] = struct{}{}
		}
		if ev.Grandparent2.Xref > 0 && ev.Parent1.Xref > 0 {
			if children[ev.Parent1.Xref] == nil {
				children[ev.Parent1.Xref] = make(map[int]struct{})
			}
			children[ev.Parent1.Xref][ev.Grandparent2.Xref] = struct{}{}
		}
		if ev.Grandparent3.Xref > 0 && ev.Parent2.Xref > 0 {
			if children[ev.Parent2.Xref] == nil {
				children[ev.Parent2.Xref] = make(map[int]struct{})
			}
			children[ev.Parent2.Xref][ev.Grandparent3.Xref] = struct{}{}
		}
		if ev.Grandparent4.Xref > 0 && ev.Parent2.Xref > 0 {
			if children[ev.Parent2.Xref] == nil {
				children[ev.Parent2.Xref] = make(map[int]struct{})
			}
			children[ev.Parent2.Xref][ev.Grandparent4.Xref] = struct{}{}
		}
		if ev.Parent1.Xref > 0 && ev.Parent2.Xref > 0 && (ev.Married || ev.Details.Type == "Trauung") {
			if marriedWith[ev.Parent1.Xref] == nil {
				marriedWith[ev.Parent1.Xref] = make(map[int]string)
			}
			if ev.Details.Type == "Trauung" {
				marriedWith[ev.Parent1.Xref][ev.Parent2.Xref] = ev.Date
			} else if _, exists := marriedWith[ev.Parent1.Xref][ev.Parent2.Xref]; !exists {
				marriedWith[ev.Parent1.Xref][ev.Parent2.Xref] = ""
			}
		}

		for _, key := range AdditonalKeys {
			if ev.Additionals[key].Parent.Xref > 0 && ev.Additionals[key].Child.Xref > 0 {
				if children[ev.Additionals[key].Child.Xref] == nil {
					children[ev.Additionals[key].Child.Xref] = make(map[int]struct{})
				}
				children[ev.Additionals[key].Child.Xref][ev.Additionals[key].Parent.Xref] = struct{}{}
			}
			if ev.Additionals[key].Parent.Xref > 0 && ev.Additionals[key].Spouse.Xref > 0 {
				if marriedWith[ev.Additionals[key].Parent.Xref] == nil {
					marriedWith[ev.Additionals[key].Parent.Xref] = make(map[int]string)
				}
				if _, exists := marriedWith[ev.Additionals[key].Parent.Xref][ev.Additionals[key].Spouse.Xref]; !exists {
					marriedWith[ev.Additionals[key].Parent.Xref][ev.Additionals[key].Spouse.Xref] = ""
				}
			}
		}
	}

	for _, pId := range children {
		father := 0
		mother := 0
		for k, _ := range pId {
			if strings.HasPrefix(Data.Individual[k].Sex, "m") {
				father = k
			} else {
				mother = k
			}
		}
		if father > 0 && mother > 0 {
			if spans[father] == nil {
				spans[father] = make(map[int]span)
			}
			_, _, _, _, age1, age2, _ := GetPersonalData(mother)
			if age2 == "Weiland" {
				age2 = ""
			}
			spans[father][mother] = span{
				begin: age1,
				end:   age2,
			}
		}
	}

	familyList := make(map[string]Familie)

	//Married
	for m1, m := range marriedWith {
		for m2, date := range m {
			fam := Familie{}
			father := 0
			mother := 0
			if strings.HasPrefix(Data.Individual[m1].Sex, "m") {
				father = m1
				mother = m2
			} else {
				father = m2
				mother = m1
			}
			familyId := strconv.Itoa(father) + "-" + strconv.Itoa(mother)
			if _, exists := familyList[familyId]; exists {
				fam = familyList[familyId]
			} else {
				fam.Father = father
				fam.Mother = mother
			}
			fam.Married = true
			familyList[familyId] = fam

			//Spans aktualisieren
			_, _, _, _, _, age2, _ := GetPersonalData(mother)
			if age2 == "Weiland" {
				age2 = date
			}

			if spans[father] == nil {
				spans[father] = make(map[int]span)
			}
			sp := spans[father][mother]
			sp.begin = date
			if len(sp.end) == 0 || dateAfter(sp.end, age2) {
				sp.end = age2
			}
			spans[father][mother] = sp
		}
	}

	//Children
	for cId, pId := range children {
		fam := Familie{}
		father := 0
		mother := 0
		for k, _ := range pId {
			if strings.HasPrefix(Data.Individual[k].Sex, "m") {
				father = k
			} else {
				mother = k
			}
		}

		//check for mother
		if father > 0 && mother == 0 {
			_, _, _, _, birth, _, _ := GetPersonalData(cId)
			possibleMothers := []int{}
			for m, s := range spans[father] {
				if len(s.begin) > 0 && len(birth) > 0 && dateAfter(s.begin, birth) {
					continue
				}
				if len(s.end) > 0 && len(birth) > 0 && dateAfter(birth, s.end) {
					continue
				}
				if len(birth) > 0 && len(s.begin) > 0 && len(s.end) > 0 && dateAfter(birth, s.begin) && dateAfter(s.end, birth) {
					possibleMothers = []int{m}
					break
				}
				possibleMothers = append(possibleMothers, m)
			}
			if len(possibleMothers) == 1 {
				mother = possibleMothers[0]
			}
		}

		familyId := strconv.Itoa(father) + "-" + strconv.Itoa(mother)

		if _, exists := familyList[familyId]; exists {
			fam = familyList[familyId]
		} else {
			fam.Father = father
			fam.Mother = mother
		}
		fam.Childen = append(fam.Childen, cId)

		if _, exists := marriedWith[father][mother]; exists {
			fam.Married = true
		}
		if _, exists := marriedWith[mother][father]; exists {
			fam.Married = true
		}

		familyList[familyId] = fam
	}

	FamilyStorage = familyList
}
