package data

import (
	"fmt"
	"sort"
)

const dateLayout = "02.01.2006"

type personalEvent struct {
	eId     int
	date    string
	details Details
}

func GetPersonalData(id int) (valid bool, gn, fn, sex, age1, age2 string, events []string) {
	if pd, ok := PersonalDateStorage[id]; ok {
		return true, pd.gn, pd.fn, pd.sex, pd.age1, pd.age2, pd.events
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
			if ok, gn, fn, _, age1, age2, _ := GetPersonalData(fam.Father); ok {
				if len(age1) == 0 {
					age1 = "          "
				}
				if len(age2) == 0 {
					age2 = "          "
				}
				family = append(family, fmt.Sprintf("%s %s | father | (%s-%s) I-%d", gn, fn, age1, age2, fam.Father))
			}
			if ok, gn, fn, _, age1, age2, _ := GetPersonalData(fam.Mother); ok {
				if len(age1) == 0 {
					age1 = "          "
				}
				if len(age2) == 0 {
					age2 = "          "
				}
				family = append(family, fmt.Sprintf("%s %s | mother | (%s-%s) I-%d", gn, fn, age1, age2, fam.Mother))
			}
			sibs := []string{}
			for _, child := range fam.Childen {
				if ok, gn, fn, _, age1, age2, _ := GetPersonalData(child); ok && child != id {
					if len(age1) == 0 {
						age1 = "          "
					}
					if len(age2) == 0 {
						age2 = "          "
					}
					sibs = append(sibs, fmt.Sprintf("%s %s | sibling | (%s-%s) I-%d", gn, fn, age1, age2, child))
				}
			}
			sort.Sort(personSort(sibs))
			family = append(family, sibs...)
			family = append(family, "")
		}
	}

	//Children
	families := make(map[int][]string)
	for _, fam := range FamilyStorage {
		hit := false
		firstId := 0
		tempFam := []string{}
		if fam.Father == id {
			if ok, gn, fn, _, age1, age2, _ := GetPersonalData(fam.Mother); ok {
				if len(age1) == 0 {
					age1 = "          "
				}
				if len(age2) == 0 {
					age2 = "          "
				}
				if fam.Married {
					tempFam = append(tempFam, fmt.Sprintf("%s %s | wife | (%s-%s) I-%d", gn, fn, age1, age2, fam.Mother))
				} else {
					tempFam = append(tempFam, fmt.Sprintf("%s %s | partner | (%s-%s) I-%d", gn, fn, age1, age2, fam.Mother))
				}
				firstId = fam.Mother
			}
			hit = true
		}
		if fam.Mother == id {
			if ok, gn, fn, _, age1, age2, _ := GetPersonalData(fam.Father); ok {
				if len(age1) == 0 {
					age1 = "          "
				}
				if len(age2) == 0 {
					age2 = "          "
				}
				if fam.Married {
					tempFam = append(tempFam, fmt.Sprintf("%s %s | husband | (%s-%s) I-%d", gn, fn, age1, age2, fam.Father))
				} else {
					tempFam = append(tempFam, fmt.Sprintf("%s %s | partner | (%s-%s) I-%d", gn, fn, age1, age2, fam.Father))
				}
				if firstId == 0 {
					firstId = fam.Father
				}
			}
			hit = true
		}
		if hit {
			children := []string{}
			for _, child := range fam.Childen {
				if ok, gn, fn, _, age1, age2, _ := GetPersonalData(child); ok {
					if len(age1) == 0 {
						age1 = "          "
					}
					if len(age2) == 0 {
						age2 = "          "
					}
					children = append(children, fmt.Sprintf("%s %s | child | (%s-%s) I-%d", gn, fn, age1, age2, child))
					if firstId == 0 {
						firstId = child
					}
				}
			}
			sort.Sort(personSort(children))
			tempFam = append(tempFam, children...)
			tempFam = append(tempFam, "")
		}
		families[firstId] = tempFam
	}

	sortMe := []int{}
	for fId, _ := range families {
		sortMe = append(sortMe, fId)
	}
	sort.Ints(sortMe)
	for _, fId := range sortMe {
		family = append(family, families[fId]...)
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
	if Data.CounterI == 0 {
		Data.CounterI++
	}
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
