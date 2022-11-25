package ui

import (
	"lympach/data"
	"strconv"
	"strings"

	"github.com/gen2brain/iup-go/iup"
)

var defaultColor = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var pStart = []int{
	0, 2, 4, 6, 8, 10,
	24, 26, 28, 30, 32, 34,
	48, 50, 52, 54, 56, 58,
	72, 74, 76, 78, 80, 82,
	96, 98, 100, 102, 104, 106,
	120, 122, 124, 126, 128, 130,
}

func create(handle string, colors []byte) {
	colorMap := iup.Image(12, 13, colors)
	colorMap.SetAttribute("0", "240 240 240") //Default
	colorMap.SetAttribute("1", "0 0 255")     //Limbach blau
	colorMap.SetAttribute("2", "200 0 0")     //Sora rot
	colorMap.SetAttribute("3", "255 128 0")   //Lampersdorf orange
	colorMap.SetAttribute("4", "0 128 0")     //Birkenhain grün
	colorMap.SetAttribute("5", "200 200 0")   //Lotzen gelb
	colorMap.SetAttribute("6", "100 100 100") //Sonstiges
	iup.SetHandle("colorMap"+handle, colorMap)
	iup.GetHandle(handle).SetAttribute("IMAGE", "colorMap"+handle)
}

func createColorBoxes(person []data.ResultData) {
	for k := 0; k < 46; k++ {
		if k < len(person) {
			create("color"+strconv.Itoa(k), createPersonalColor(person[k].Events))
		} else {
			create("color"+strconv.Itoa(k), createPersonalColor([]string{}))
		}
	}
}

func createPersonalColor(events []string) (colorMap []byte) {
	colorMap = make([]byte, 156)
	if len(events) > len(pStart) {
		events = events[:len(pStart)]
	}
	offset := 12

	i := 0
	for j := len(events) - 1; j >= 0; j-- {

		colorPixel := byte(6)
		if strings.Contains(events[j], "Limbach") {
			colorPixel = byte(1)
		}
		if strings.Contains(events[j], "Sora") {
			colorPixel = byte(2)
		}
		if strings.Contains(events[j], "Lampersdorf") {
			colorPixel = byte(3)
		}
		if strings.Contains(events[j], "Birkenhain") {
			colorPixel = byte(4)
		}
		if strings.Contains(events[j], "Lotzen") {
			colorPixel = byte(5)
		}
		colorMap[pStart[i]+offset] = colorPixel
		colorMap[pStart[i]+1+offset] = colorPixel
		colorMap[pStart[i]+12+offset] = colorPixel
		colorMap[pStart[i]+13+offset] = colorPixel
		i++
	}
	return
}
