package data

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/iup-go/iup"

	"github.com/antzucaro/matchr"
)

type ResultData struct {
	Data   string
	Events []string
}

func FindPerson(input string) (resultList []ResultData) {
	findById := false
	findByShort := 0
	nr, err := strconv.Atoi(input)
	if err == nil && nr > 0 {
		findById = true
	} else if len(input) == 2 {
		findByShort = 2
	} else if len(input) == 3 {
		findByShort = 3
	}

	results := make(map[int][]ResultData)

	for _, indi := range Data.Individual {
		sex := indi.Sex
		pd := getPersonalData(indi.Xref)

		if findById && strconv.Itoa(indi.Xref) == input {
			resultList = []ResultData{pd}
			return
		}

		indiVn := strings.ToLower(indi.GName)
		indiFn := strings.ToLower(indi.FName)

		//Initialien (nur männliche Personen)
		if findByShort > 0 {
			if sex != "m" {
				continue
			}
			//short 3
			g2 := indiVn
			if len(indiVn) > 1 {
				g2 = indiVn[0:2]
			}
			f2 := indiFn
			if len(indiFn) > 1 {
				f2 = indiFn[0:2]
			}
			//short 2
			if len(indiVn) > 0 {
				indiVn = indiVn[0:1]
			}
			if len(indiFn) > 0 {
				indiFn = indiFn[0:1]
			}

			if input == indiVn+indiFn || input == indiVn+f2 || input == g2+indiFn {
				//max 100 Jahre
				if len(pd.Events) > 0 {
					_, eventDate, _, _, _ := SplitEvent(pd.Events[len(pd.Events)-1])
					if len(eventDate) > 3 {
						eventYear := eventDate[len(eventDate)-4:]
						ey, _ := strconv.Atoi(eventYear)
						now := iup.GetHandle("eventDate").GetAttribute("VALUE")
						parsed, err := time.Parse(dateLayout, now)
						if err == nil && parsed.Year()-ey > 100 {
							continue
						}
					}
				}
				results[0] = append(results[0], pd)
			}
		} else {
			vn, fn := SplitName(strings.ToLower(input))
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
