package data

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/antzucaro/matchr"
)

type ResultData struct {
	Data   string
	Events []string
}

func FindPerson(input string) (resultList []ResultData) {
	findById := false
	findByShort := false
	nr, err := strconv.Atoi(input)
	if err == nil && nr > 0 {
		findById = true
	} else if len(input) == 2 {
		findByShort = true
		input = input[0:1] + " " + input[1:2]
	}

	results := make(map[int][]ResultData)

	vn, fn := SplitName(strings.ToLower(input))

	for _, indi := range Data.Individual {
		sex := indi.Sex
		if sex == "male" {
			sex = "m"
		} else if sex == "female" {
			sex = "f"
		} else if sex == "unknown" {
			sex = "u"
		}

		pd := getPersonalData(indi.Xref)

		if findById && strconv.Itoa(indi.Xref) == input {
			resultList = []ResultData{pd}
			return
		}

		indiVn := strings.ToLower(indi.GName)
		indiFn := strings.ToLower(indi.FName)

		//Initialien (nur männliche Personen)
		if findByShort {
			if sex != "m" {
				continue
			}
			if len(indiVn) > 0 {
				indiVn = indiVn[0:1]
			}
			if len(indiFn) > 0 {
				indiFn = indiFn[0:1]
			}
		}

		distance := 0
		if len(fn) == 0 {
			//Simple Search
			distance = matchr.DamerauLevenshtein(input, indiVn+" "+indiFn)
		} else {
			//Double Search
			distanceVn := 0
			distanceFn := 0
			//? matched auf alles
			if vn != "?" {
				distanceVn = matchr.DamerauLevenshtein(vn, indiVn)
			}
			if fn != "?" {
				distanceFn = matchr.DamerauLevenshtein(fn, indiFn)
			}

			//1/3 der Differenz abziehen => damit werden Matches leicht bevorteilt, wo ein Part sehr gut matched
			distance = distanceVn + distanceFn - int(math.Round(0.3*math.Abs(float64(distanceVn)-float64(distanceFn))))
			if distance < 0 {
				distance = 0
			}
		}
		results[distance] = append(results[distance], pd)
	}

	for i := 0; i < 20; i++ {
		if len(results[i]) > 0 {
			sort.Sort(customSort(results[i]))
			results[i] = append(results[i], ResultData{Data: ""})
		}
		resultList = append(resultList, results[i]...)
	}

	return
}

//Events, Lebensdaten etc.
func getPersonalData(id int) (pd ResultData) {

	valid, gn, fn, sex, age1, age2, events := GetPersonalData(id)
	if !valid {
		return
	}

	if len(age1) == 0 {
		age1 = "          "
	}
	if len(age2) == 0 {
		age2 = "          "
	}

	pd = ResultData{
		Data:   fmt.Sprintf("%s %s | (%s-%s) %s | I-%d", gn, fn, age1, age2, sex, id),
		Events: events,
	}

	return
}

func SplitName(full string) (vn, fn string) {
	searchParts := strings.Split(full, " ")
	if len(searchParts) < 2 {
		vn = full
		return
	}
	fn = searchParts[len(searchParts)-1]
	searchParts = searchParts[:len(searchParts)-1]
	vn = strings.Join(searchParts, " ")

	return
}
