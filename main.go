package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strconv"
	"strings"
)

func main() {
	var inTE, outTE *walk.TextEdit
	MainWindow{
		Title:   "计算器·CALCULATOR",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE, Font: Font{PointSize: 15}, Name: "IN"},
					TextEdit{AssignTo: &outTE, ReadOnly: true, Font: Font{PointSize: 20}, Name: "OUT"},
				},
			},
			PushButton{
				Text:    "CALCULATOR",
				Font:    Font{PointSize: 15},
				MinSize: Size{10, 10},
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(strconv.Itoa(calc(inTE.Text()))))
				},
			},
		},
	}.Run()
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func calc(args string) int {
	if strings.Contains(args, "+") {
		strs := strings.Split(args, "+")
		return Atoi(strs[0]) + Atoi(strs[1])
	} else if strings.Contains(args, "-") {
		strs := strings.Split(args, "-")
		return Atoi(strs[0]) - Atoi(strs[1])
	} else if strings.Contains(args, "*") {
		strs := strings.Split(args, "*")
		return Atoi(strs[0]) * Atoi(strs[1])
	} else if strings.Contains(args, "/") {
		strs := strings.Split(args, "/")
		return Atoi(strs[0]) / Atoi(strs[1])
	}
	return 0
}
