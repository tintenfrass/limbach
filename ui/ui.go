package ui

import (
	"strconv"

	"github.com/gen2brain/iup-go/iup"
)

var colors = []iup.Ihandle{}

func BuildAndRun() {
	//Control
	iup.Open()
	defer iup.Close()

	personField := iup.List().SetAttribute("SIZE", "235x370").SetHandle("person")
	searchField := iup.Text().SetAttribute("SIZE", "240x").SetHandle("search")

	eventTypes := iup.List().SetHandle("eventtypes").SetAttribute("SIZE", "60x66")
	initFixedEventtypes()
	msg := iup.Text().SetAttribute("SIZE", "60x").SetHandle("msg")
	places := iup.List().SetHandle("places").SetAttribute("SIZE", "85x50")
	initFixedPlaces()
	places1 := iup.Text().SetHandle("places1").SetAttribute("SIZE", "85x")
	places2 := iup.List().SetHandle("places2").SetAttribute("SIZE", "85x205")
	occupations := iup.List().SetHandle("occupations").SetAttribute("SIZE", "85x222")
	personalEvents := iup.List().SetAttribute("SIZE", "240x222").SetHandle("personalEvents")
	family := iup.List().SetAttribute("SIZE", "240x235").SetHandle("family")
	selectButton := iup.Button("Select").SetAttribute("SIZE", "30x15")

	sex := iup.Radio(
		iup.Vbox(
			iup.Toggle("male").SetHandle("m"),
			iup.Toggle("female").SetHandle("f"),
			iup.Toggle("unknown").SetHandle("u"),
		),
	).SetHandle("sex")
	inputgname := iup.Text().SetAttribute("SIZE", "60x").SetHandle("inputgname")
	inputfname := iup.Text().SetAttribute("SIZE", "60x").SetHandle("inputfname")
	changeButton := iup.Button("Change").SetAttribute("SIZE", "35x15")
	newButton := iup.Button("New").SetAttribute("SIZE", "30x15")
	deleteButton := iup.Button("Delete").SetAttribute("SIZE", "30x15")

	//Color
	for i := 0; i < 46; i++ {
		key := strconv.Itoa(i)
		colors = append(colors, iup.Label("").SetHandle("color"+key))
		create("color"+key, defaultColor)
	}

	married := iup.Toggle("verheiratet").SetHandle("married")
	marriedPlace := iup.Label("").SetHandle("marriedPlace").SetAttribute("SIZE", "80x")
	marriedMsg := iup.Label("").SetHandle("marriedMsg").SetAttribute("SIZE", "80x")
	marriedSet := iup.Button("Set").SetAttribute("SIZE", "40x15")
	eventId := iup.Label("").SetAttribute("SIZE", "20x").SetHandle("eventId")
	eventDate := iup.Text().SetAttribute("SIZE", "50x").SetHandle("eventDate")
	grandparent1 := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("grandparent1").SetAttribute("ALIGNMENT", ":ATOP")
	grandparent2 := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("grandparent2").SetAttribute("ALIGNMENT", ":ATOP")
	grandparent3 := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("grandparent3").SetAttribute("ALIGNMENT", ":ATOP")
	grandparent4 := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("grandparent4").SetAttribute("ALIGNMENT", ":ATOP")
	grandparent1Set := iup.Button("Set").SetAttribute("SIZE", "50x15")
	grandparent1Remove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	grandparent2Set := iup.Button("Set").SetAttribute("SIZE", "50x15")
	grandparent2Remove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	grandparent3Set := iup.Button("Set").SetAttribute("SIZE", "50x15")
	grandparent3Remove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	grandparent4Set := iup.Button("Set").SetAttribute("SIZE", "50x15")
	grandparent4Remove := iup.Button("Remove").SetAttribute("SIZE", "50x15")

	parent1 := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("parent1").SetAttribute("ALIGNMENT", ":ATOP")
	parent2 := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("parent2").SetAttribute("ALIGNMENT", ":ATOP")
	parent1Set := iup.Button("Set").SetAttribute("SIZE", "50x15")
	parent1Remove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	parent2Set := iup.Button("Set").SetAttribute("SIZE", "50x15")
	parent2Remove := iup.Button("Remove").SetAttribute("SIZE", "50x15")

	child := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("child").SetAttribute("ALIGNMENT", ":ATOP")
	childSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	childRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")

	parentA := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("parentA").SetAttribute("ALIGNMENT", ":ATOP")
	parentASet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	parentARemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	spouseA := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("spouseA").SetAttribute("ALIGNMENT", ":ATOP")
	spouseASet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	spouseARemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	childA := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("childA").SetAttribute("ALIGNMENT", ":ATOP")
	childASet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	childARemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	parentB := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("parentB").SetAttribute("ALIGNMENT", ":ATOP")
	parentBSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	parentBRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	spouseB := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("spouseB").SetAttribute("ALIGNMENT", ":ATOP")
	spouseBSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	spouseBRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	childB := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("childB").SetAttribute("ALIGNMENT", ":ATOP")
	childBSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	childBRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	parentC := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("parentC").SetAttribute("ALIGNMENT", ":ATOP")
	parentCSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	parentCRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	spouseC := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("spouseC").SetAttribute("ALIGNMENT", ":ATOP")
	spouseCSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	spouseCRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	childC := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("childC").SetAttribute("ALIGNMENT", ":ATOP")
	childCSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	childCRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	parentD := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("parentD").SetAttribute("ALIGNMENT", ":ATOP")
	parentDSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	parentDRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	spouseD := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("spouseD").SetAttribute("ALIGNMENT", ":ATOP")
	spouseDSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	spouseDRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")
	childD := iup.Label("").SetAttribute("SIZE", "100x42").SetHandle("childD").SetAttribute("ALIGNMENT", ":ATOP")
	childDSet := iup.Button("Set").SetAttribute("SIZE", "50x15")
	childDRemove := iup.Button("Remove").SetAttribute("SIZE", "50x15")

	reset := iup.Button("Reset").SetAttribute("SIZE", "50x15")
	del := iup.Button("Delete").SetAttribute("SIZE", "50x15")

	exportButton := iup.Button("Export").SetAttribute("SIZE", "80x15")
	exitButton := iup.Button("Exit").SetAttribute("SIZE", "80x15")

	//UI
	content := iup.Hbox(
		iup.Vbox(
			iup.Space().SetAttribute("SIZE", "x198"),
			iup.Vbox(
				colors...,
			),
		).SetAttributes("MARGIN=1x1, GAP=1"),
		iup.Vbox(
			iup.Space().SetAttribute("SIZE", "x197"),
			personField,
		).SetAttributes("MARGIN=1x1, GAP=2"),
		iup.Vbox(
			personalEvents,
			iup.Hbox(
				iup.Vbox(
					iup.Space().SetAttribute("SIZE", "x72"),
					selectButton,
				),
				iup.Space().SetAttribute("SIZE", "10x"),
				iup.Vbox(
					iup.Frame(
						iup.Vbox(
							iup.Hbox(
								changeButton,
								deleteButton,
							),
							sex,
							newButton,
							iup.Hbox(
								inputgname,
								inputfname,
							),
						),
					).SetAttribute("TITLE", "Person"),
				),
				iup.Vbox(
					eventTypes,
					iup.Frame(
						msg,
					).SetAttribute("TITLE", "").SetHandle("desc"),
				),
			),
			iup.Space().SetAttribute("SIZE", "x3"),
			searchField,
			family,
		).SetAttributes("MARGIN=1x1, GAP=2"),
		iup.Vbox(
			occupations,
			places,
			places1,
			places2,
			exportButton,
			exitButton,
		).SetAttributes("MARGIN=1x1, GAP=2"),
		iup.Space().SetAttribute("SIZE", "50x"),
		iup.Vbox(
			iup.Space().SetAttribute("SIZE", "x30"),
			iup.Hbox(
				iup.Frame(
					iup.Vbox(
						grandparent1,
						iup.Hbox(
							grandparent1Set,
							grandparent1Remove,
						),
					),
				),
				iup.Space().SetAttribute("SIZE", "2x"),
				iup.Frame(
					iup.Vbox(
						grandparent2,
						iup.Hbox(
							grandparent2Set,
							grandparent2Remove,
						),
					),
				),
				iup.Space().SetAttribute("SIZE", "10x"),
				iup.Frame(
					iup.Vbox(
						grandparent3,
						iup.Hbox(
							grandparent3Set,
							grandparent3Remove,
						),
					),
				),
				iup.Space().SetAttribute("SIZE", "2x"),
				iup.Frame(
					iup.Vbox(
						grandparent4,
						iup.Hbox(
							grandparent4Set,
							grandparent4Remove,
						),
					),
				),
			),
			iup.Space().SetAttribute("SIZE", "x10"),
			iup.Hbox(
				iup.Space().SetAttribute("SIZE", "70x"),
				iup.Frame(
					iup.Vbox(
						parent1,
						iup.Hbox(
							parent1Set,
							parent1Remove,
						),
					),
				),
				iup.Space().SetAttribute("SIZE", "2x"),
				iup.Space().SetAttribute("SIZE", "x30"),
				iup.Frame(
					iup.Vbox(
						married,
						iup.Space().SetAttribute("SIZE", "x10"),
						marriedPlace,
						marriedMsg,
						iup.Hbox(
							iup.Space().SetAttribute("SIZE", "18x"),
							marriedSet,
						),
					),
				).SetAttribute("SIZE", "80x").SetAttribute("TITLE", "oo"),
				iup.Space().SetAttribute("SIZE", "2x"),
				iup.Frame(
					iup.Vbox(
						parent2,
						iup.Hbox(
							parent2Set,
							parent2Remove,
						),
					),
				),
			),
			iup.Space().SetAttribute("SIZE", "x10"),
			iup.Hbox(
				iup.Space().SetAttribute("SIZE", "164x"),
				iup.Frame(
					iup.Vbox(
						child,
						iup.Hbox(
							childSet,
							childRemove,
						),
					),
				),
			),
			iup.Space().SetAttribute("SIZE", "x10"),
			iup.Frame(
				iup.Hbox(
					iup.Space().SetAttribute("SIZE", "1x"),
					iup.Vbox(
						iup.Space().SetAttribute("SIZE", "x2"),
						eventId,
					),
					iup.Space().SetAttribute("SIZE", "10x"),
					iup.Vbox(
						iup.Space().SetAttribute("SIZE", "x1"),
						eventDate,
					),
					iup.Space().SetAttribute("SIZE", "10x"),
					reset,
				),
			).SetAttribute("TITLE", "Event"),
			iup.Space().SetAttribute("SIZE", "x5"),
			iup.Hbox(
				iup.Space().SetAttribute("SIZE", "2x"),
				iup.Frame(
					iup.Vbox(
						parentA,
						iup.Hbox(
							parentASet,
							parentARemove,
						),
					),
				),
				iup.Frame(
					iup.Vbox(
						spouseA,
						iup.Hbox(
							spouseASet,
							spouseARemove,
						),
					),
				),
				iup.Space().SetAttribute("SIZE", "10x"),
				iup.Frame(
					iup.Vbox(
						parentB,
						iup.Hbox(
							parentBSet,
							parentBRemove,
						),
					),
				),
				iup.Frame(
					iup.Vbox(
						spouseB,
						iup.Hbox(
							spouseBSet,
							spouseBRemove,
						),
					),
				),
			),
			iup.Hbox(
				iup.Space().SetAttribute("SIZE", "2x"),
				iup.Frame(
					iup.Vbox(
						childA,
						iup.Hbox(
							childASet,
							childARemove,
						),
					),
				),
				iup.Space().SetAttribute("SIZE", "114x"),
				iup.Frame(
					iup.Vbox(
						childB,
						iup.Hbox(
							childBSet,
							childBRemove,
						),
					),
				),
			),
			iup.Space().SetAttribute("SIZE", "x10"),
			iup.Hbox(
				iup.Space().SetAttribute("SIZE", "2x"),
				iup.Frame(
					iup.Vbox(
						parentC,
						iup.Hbox(
							parentCSet,
							parentCRemove,
						),
					),
				),
				iup.Frame(
					iup.Vbox(
						spouseC,
						iup.Hbox(
							spouseCSet,
							spouseCRemove,
						),
					),
				),
				iup.Space().SetAttribute("SIZE", "10x"),
				iup.Frame(
					iup.Vbox(
						parentD,
						iup.Hbox(
							parentDSet,
							parentDRemove,
						),
					),
				),
				iup.Frame(
					iup.Vbox(
						spouseD,
						iup.Hbox(
							spouseDSet,
							spouseDRemove,
						),
					),
				),
			),
			iup.Hbox(
				iup.Space().SetAttribute("SIZE", "2x"),
				iup.Frame(
					iup.Vbox(
						childC,
						iup.Hbox(
							childCSet,
							childCRemove,
						),
					),
				),
				iup.Space().SetAttribute("SIZE", "114x"),
				iup.Frame(
					iup.Vbox(
						childD,
						iup.Hbox(
							childDSet,
							childDRemove,
						),
					),
				),
			),
			iup.Space().SetAttribute("SIZE", "x20"),
			iup.Hbox(
				iup.Space().SetAttribute("SIZE", "375x"),
				del,
			),
		),
	).SetAttribute("FONT", "Consolas, 9")

	dlg := iup.Dialog(content).SetAttributes(utf82ui(`TITLE="Lympach"`))
	dlg.SetHandle("dlg").SetAttributes("SIZE=1285x560")

	iup.SetCallback(searchField, "VALUECHANGED_CB", iup.ValueChangedFunc(search))
	iup.SetCallback(personField, "VALUECHANGED_CB", iup.ValueChangedFunc(showPerson))
	iup.SetCallback(selectButton, "ACTION", iup.ActionFunc(selectPerson))
	iup.SetCallback(changeButton, "ACTION", iup.ActionFunc(changePerson))
	iup.SetCallback(newButton, "ACTION", iup.ActionFunc(newPerson))
	iup.SetCallback(deleteButton, "ACTION", iup.ActionFunc(deletePerson))
	iup.SetCallback(occupations, "VALUECHANGED_CB", iup.ValueChangedFunc(occupation))
	iup.SetCallback(places, "VALUECHANGED_CB", iup.ValueChangedFunc(place))
	iup.SetCallback(places2, "VALUECHANGED_CB", iup.ValueChangedFunc(place))
	iup.SetCallback(inputgname, "VALUECHANGED_CB", iup.ValueChangedFunc(gname))
	iup.SetCallback(msg, "VALUECHANGED_CB", iup.ValueChangedFunc(msgChange))

	//Events
	iup.SetCallback(personalEvents, "VALUECHANGED_CB", iup.ValueChangedFunc(loadEvent))
	iup.SetCallback(eventDate, "VALUECHANGED_CB", iup.ValueChangedFunc(changeDate))
	iup.SetCallback(married, "VALUECHANGED_CB", iup.ValueChangedFunc(changeMarried))
	iup.SetCallback(marriedSet, "ACTION", iup.ActionFunc(marriedDetails))

	iup.SetCallback(grandparent1Set, "ACTION", iup.ActionFunc(setGP1))
	iup.SetCallback(grandparent2Set, "ACTION", iup.ActionFunc(setGP2))
	iup.SetCallback(grandparent3Set, "ACTION", iup.ActionFunc(setGP3))
	iup.SetCallback(grandparent4Set, "ACTION", iup.ActionFunc(setGP4))
	iup.SetCallback(parent1Set, "ACTION", iup.ActionFunc(setP1))
	iup.SetCallback(parent2Set, "ACTION", iup.ActionFunc(setP2))
	iup.SetCallback(childSet, "ACTION", iup.ActionFunc(setC))
	iup.SetCallback(parentASet, "ACTION", iup.ActionFunc(setPA))
	iup.SetCallback(spouseASet, "ACTION", iup.ActionFunc(setSA))
	iup.SetCallback(childASet, "ACTION", iup.ActionFunc(setCA))
	iup.SetCallback(parentBSet, "ACTION", iup.ActionFunc(setPB))
	iup.SetCallback(spouseBSet, "ACTION", iup.ActionFunc(setSB))
	iup.SetCallback(childBSet, "ACTION", iup.ActionFunc(setCB))
	iup.SetCallback(parentCSet, "ACTION", iup.ActionFunc(setPC))
	iup.SetCallback(spouseCSet, "ACTION", iup.ActionFunc(setSC))
	iup.SetCallback(childCSet, "ACTION", iup.ActionFunc(setCC))
	iup.SetCallback(parentDSet, "ACTION", iup.ActionFunc(setPD))
	iup.SetCallback(spouseDSet, "ACTION", iup.ActionFunc(setSD))
	iup.SetCallback(childDSet, "ACTION", iup.ActionFunc(setCD))

	iup.SetCallback(grandparent1Remove, "ACTION", iup.ActionFunc(removeGP1))
	iup.SetCallback(grandparent2Remove, "ACTION", iup.ActionFunc(removeGP2))
	iup.SetCallback(grandparent3Remove, "ACTION", iup.ActionFunc(removeGP3))
	iup.SetCallback(grandparent4Remove, "ACTION", iup.ActionFunc(removeGP4))
	iup.SetCallback(parent1Remove, "ACTION", iup.ActionFunc(removeP1))
	iup.SetCallback(parent2Remove, "ACTION", iup.ActionFunc(removeP2))
	iup.SetCallback(childRemove, "ACTION", iup.ActionFunc(removeC))
	iup.SetCallback(parentARemove, "ACTION", iup.ActionFunc(removePA))
	iup.SetCallback(spouseARemove, "ACTION", iup.ActionFunc(removeSA))
	iup.SetCallback(childARemove, "ACTION", iup.ActionFunc(removeCA))
	iup.SetCallback(parentBRemove, "ACTION", iup.ActionFunc(removePB))
	iup.SetCallback(spouseBRemove, "ACTION", iup.ActionFunc(removeSB))
	iup.SetCallback(childBRemove, "ACTION", iup.ActionFunc(removeCB))
	iup.SetCallback(parentCRemove, "ACTION", iup.ActionFunc(removePC))
	iup.SetCallback(spouseCRemove, "ACTION", iup.ActionFunc(removeSC))
	iup.SetCallback(childCRemove, "ACTION", iup.ActionFunc(removeCC))
	iup.SetCallback(parentDRemove, "ACTION", iup.ActionFunc(removePD))
	iup.SetCallback(spouseDRemove, "ACTION", iup.ActionFunc(removeSD))
	iup.SetCallback(childDRemove, "ACTION", iup.ActionFunc(removeCD))
	iup.SetCallback(reset, "ACTION", iup.ActionFunc(resetEvent))
	iup.SetCallback(del, "ACTION", iup.ActionFunc(deleteEvent))

	iup.SetCallback(exportButton, "ACTION", iup.ActionFunc(export))
	iup.SetCallback(exitButton, "ACTION", iup.ActionFunc(exit))

	//Run
	iup.Map(dlg)
	iup.Show(dlg)
	iup.MainLoop()
}
