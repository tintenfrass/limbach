package data

import (
	"strings"

	"github.com/gen2brain/iup-go/iup"
)

type names struct {
	male   map[string]struct{}
	female map[string]struct{}
}

var nameList names

func LoadAllNames() {
	nameList.male = make(map[string]struct{}, 0)
	nameList.female = make(map[string]struct{}, 0)

	for _, indi := range Data.Individual {
		AddName(indi.GName, indi.Sex)
	}
}

func AddName(name, sex string) {
	vn, _ := SplitName(name)
	if len(vn) < 2 {
		return
	}
	if strings.HasPrefix(sex, "m") {
		nameList.male[vn] = struct{}{}
	} else if strings.HasPrefix(sex, "f") {
		nameList.female[vn] = struct{}{}
	}
}

func CheckVnName(name string) {
	vn, _ := SplitName(name)
	if len(vn) < 2 {
		return
	}

	if _, ok := nameList.male[vn]; ok {
		iup.GetHandle("sex").SetAttribute("VALUE", "m")
	} else if _, ok = nameList.female[vn]; ok {
		iup.GetHandle("sex").SetAttribute("VALUE", "f")
	}
}
