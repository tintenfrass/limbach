package data

import (
	"fmt"
	"strings"
)

type event struct {
	typ   string
	date  string
	place string
	msg   string
}

var tempFams []Familie

func getGedcomData(data *FullData) (output string) {
	tempFams = []Familie{}
	i := 0
	for _, fam := range FamilyStorage {
		tempFams = append(tempFams, fam)
		i++
	}

	//header
	output += "0 HEAD\r\n"
	output += "1 CHAR UTF-8\r\n"
	output += "1 SOUR Lÿmpach\r\n"

	//individuals
	for _, indi := range data.Individual {
		valid, _, _, _, _, _, events := GetPersonalData(indi.Xref)
		if !valid {
			continue
		}

		output += fmt.Sprintf("0 %s INDI\r\n", getI(indi.Xref))
		output += fmt.Sprintf("1 NAME %s /%s/\r\n", indi.GName, indi.FName)
		output += fmt.Sprintf("2 GIVN %s\r\n", indi.GName)
		output += fmt.Sprintf("2 SURN %s\r\n", indi.FName)
		output += fmt.Sprintf("1 SEX %s\r\n", strings.ToUpper(indi.Sex[:1]))
		exists, date, place := getBirt(events)
		if exists {
			output += "1 BIRT\r\n"
			output += fmt.Sprintf("2 DATE %s\r\n", date)
			output += fmt.Sprintf("2 PLAC %s\r\n", place)
		}
		exists, date, place = getChr(events)
		if exists {
			output += "1 CHR\r\n"
			output += fmt.Sprintf("2 DATE %s\r\n", date)
			output += fmt.Sprintf("2 PLAC %s\r\n", place)
		}
		exists, date, place = getDeat(events)
		if exists {
			output += "1 DEAT\r\n"
			output += fmt.Sprintf("2 DATE %s\r\n", date)
			output += fmt.Sprintf("2 PLAC %s\r\n", place)
		}
		for _, ev := range getEvents(events) {
			output += fmt.Sprintf("1 EVEN %s\r\n", ev.msg)
			output += fmt.Sprintf("2 TYPE %s\r\n", ev.typ)
			output += fmt.Sprintf("2 DATE %s\r\n", ev.date)
			output += fmt.Sprintf("2 PLAC %s\r\n", ev.place)
		}
		famc := getFamc(indi.Xref)
		if len(famc) > 0 {
			output += fmt.Sprintf("1 FAMC %s\r\n", famc)
		}
		for _, fams := range getFams(indi.Xref) {
			output += fmt.Sprintf("1 FAMS %s\r\n", fams)
		}
	}

	//Familie
	for key, family := range tempFams {
		output += fmt.Sprintf("0 %s FAM\r\n", getF(key))
		if family.Father > 0 {
			output += fmt.Sprintf("1 HUSB %s\r\n", getI(family.Father))
		}
		if family.Mother > 0 {
			output += fmt.Sprintf("1 WIFE %s\r\n", getI(family.Mother))
		}
		for _, child := range family.Childen {
			output += fmt.Sprintf("1 CHIL %s\r\n", getI(child))
		}
		if family.Married {
			output += "1 MARR Y\r\n"
		}
		exists, date, place := getMarr(family.Father, family.Mother)
		if exists {
			output += fmt.Sprintf("2 DATE %s\r\n", date)
			output += fmt.Sprintf("2 PLAC %s\r\n", place)
		}
	}

	//Footer
	output += "0 TRLR"

	return
}

func getI(input int) (i string) {
	return fmt.Sprintf("@I%d@", input)
}

func getF(input int) (f string) {
	return fmt.Sprintf("@F%d@", input)
}

//Geburt
func getBirt(events []string) (exists bool, date, place string) {
	for _, ev := range events {
		_, dt, typ, pl, _ := SplitEvent(ev)
		if typ == "Geburt" {
			return true, dt, pl
		}
	}

	return
}

//Taufe
func getChr(events []string) (exists bool, date, place string) {
	for _, ev := range events {
		_, dt, typ, pl, _ := SplitEvent(ev)
		if typ == "Taufe" {
			return true, dt, pl
		}
	}

	return
}

//Tod
func getDeat(events []string) (exists bool, date, place string) {
	for _, ev := range events {
		_, dt, typ, pl, _ := SplitEvent(ev)
		if typ == "Tod" {
			return true, dt, pl
		}
	}

	return
}

//Heirat
func getMarr(father, mother int) (exists bool, date, place string) {
	for _, ev := range Data.Events {
		if ev.Details.Type == "Trauung" && ev.Parent1.Xref == father && ev.Parent2.Xref == mother {
			return true, ev.Date, ev.Details.Place
		}
	}

	return
}

//Events
func getEvents(events []string) (evs []event) {
	for _, ev := range events {
		_, dt, typ, pl, msg := SplitEvent(ev)
		if typ != "Taufe" && typ != "Geburt" && typ != "Tod" && typ != "Taufe eines Kindes" && typ != "Geburt eines Kindes" {
			evs = append(evs, event{typ, dt, pl, msg})
		}
	}

	return
}

//Family Child
func getFamc(xref int) (famc string) {
	for key, fam := range tempFams {
		for _, child := range fam.Childen {
			if child == xref {
				famc = getF(key)
				return
			}
		}
	}
	return
}

//Family Spouse
func getFams(xref int) (fams []string) {
	for key, fam := range tempFams {
		if fam.Father == xref || fam.Mother == xref {
			fams = append(fams, getF(key))
		}
	}
	return
}
