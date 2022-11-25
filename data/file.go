package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const datafile = "data.json"
const gedcomfile = "limbach.ged"

func Load() {
	jsonFile, err := os.Open(datafile)
	if err == nil {
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &Data)
	} else {
		fmt.Println(err)
	}

	PlaceStorage = make(map[string]struct{})
	UpdatePersonalDataStorage([]int{})
	UpdateStorage()
	LoadAllNames()
}

//save to json
func Save() {
	jsonData, _ := json.MarshalIndent(Data, "", " ")
	ioutil.WriteFile(datafile, jsonData, 0644)
}

//write gedcom file
func Export() {
	ioutil.WriteFile(gedcomfile, []byte(getGedcomData(&Data)), 0644)
}
