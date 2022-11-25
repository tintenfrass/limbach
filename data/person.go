package data

import (
	"fmt"
)

const dateLayout = "02.01.2006"

type personalEvent struct {
	date    string
	details Details
}

func GetPersonalData(id int) (valid bool, gn, fn, sex, age1, age2 string, events, appearances []string) {
	if pd, ok := PersonalDateStorage[id]; ok {
		return true, pd.gn, pd.fn, pd.sex, pd.age1, pd.age2, pd.events, pd.appearances
	}

	return
}

func GetFamily(id int) (valid bool, family []string) {
	if id == 0 {
		return
	}
	valid = true

	//Parents
	for _, fam := range FamilyStorage {
		hit := false
		for _, child := range fam.Childen {
			if child == id {
				hit = true
				break
			}
		}
		if hit {
			if ok, gn, fn, _, age1, age2, _, _ := GetPersonalData(fam.Father); ok {
				if len(age1) == 0 {
					age1 = "          "
				}
				if len(age2) == 0 {
					age2 = "          "
				}
				family = append(family, fmt.Sprintf("%s %s | father | (%s-%s) I-%d", gn, fn, age1, age2, fam.Father))
			}
			if ok, gn, fn, _, age1, age2, _, _ := GetPersonalData(fam.Mother); ok {
				if len(age1) == 0 {
					age1 = "          "
				}
				if len(age2) == 0 {
					age2 = "          "
				}
				family = append(family, fmt.Sprintf("%s %s | mother | (%s-%s) I-%d", gn, fn, age1, age2, fam.Mother))
			}
			for _, child := range fam.Childen {
				if ok, gn, fn, _, age1, age2, _, _ := GetPersonalData(child); ok && child != id {
					if len(age1) == 0 {
						age1 = "          "
					}
					if len(age2) == 0 {
						age2 = "          "
					}
					family = append(family, fmt.Sprintf("%s %s | sibling | (%s-%s) I-%d", gn, fn, age1, age2, child))
				}
			}
			family = append(family, "")
		}
	}

	//Children
	for _, fam := range FamilyStorage {
		hit := false
		if fam.Father == id {
			if ok, gn, fn, _, age1, age2, _, _ := GetPersonalData(fam.Mother); ok {
				if len(age1) == 0 {
					age1 = "          "
				}
				if len(age2) == 0 {
					age2 = "          "
				}
				if fam.Married {
					family = append(family, fmt.Sprintf("%s %s | wife | (%s-%s) I-%d", gn, fn, age1, age2, fam.Mother))
				} else {
					family = append(family, fmt.Sprintf("%s %s | partner | (%s-%s) I-%d", gn, fn, age1, age2, fam.Mother))
				}
			}
			hit = true
		}
		if fam.Mother == id {
			if ok, gn, fn, _, age1, age2, _, _ := GetPersonalData(fam.Father); ok {
				if len(age1) == 0 {
					age1 = "          "
				}
				if len(age2) == 0 {
					age2 = "          "
				}
				if fam.Married {
					family = append(family, fmt.Sprintf("%s %s | husband | (%s-%s) I-%d", gn, fn, age1, age2, fam.Father))
				} else {
					family = append(family, fmt.Sprintf("%s %s | partner | (%s-%s) I-%d", gn, fn, age1, age2, fam.Father))
				}
			}
			hit = true
		}
		if hit {
			for _, child := range fam.Childen {
				if ok, gn, fn, _, age1, age2, _, _ := GetPersonalData(child); ok {
					if len(age1) == 0 {
						age1 = "          "
					}
					if len(age2) == 0 {
						age2 = "          "
					}
					family = append(family, fmt.Sprintf("%s %s | child | (%s-%s) I-%d", gn, fn, age1, age2, child))
				}
			}
			family = append(family, "")
		}
	}

	return
}

func ChangePerson(id int, gn, fn, sex string) {
	for key, indi := range Data.Individual {
		if indi.Xref == id {
			person := IndividualRecord{
				Xref:  id,
				GName: gn,
				FName: fn,
				Sex:   sex,
			}
			Data.Individual[key] = person
			UpdatePersonalDataStorage([]int{id})
			AddName(gn, sex)
		}
	}
	UpdateStorage()
}

func NewPerson(gn, fn, sex string) (id int) {
	id = Data.CounterI
	Data.CounterI++
	Data.Individual[id] = IndividualRecord{
		Xref:  id,
		GName: gn,
		FName: fn,
		Sex:   sex,
	}
	UpdatePersonalDataStorage([]int{id})
	AddName(gn, sex)
	return
}

func DeletePerson(id int) {
	delete(Data.Individual, id)
}
