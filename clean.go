package main

import (
	"fmt"
	"regexp"
	"strings"
)

func cleanData(value []string) (cleanValue []string) {

	var re1 = regexp.MustCompile(`^S\w+.\w+y\B`)
	cleanValue = append(cleanValue, re1.ReplaceAllString(value[0], ""))

	var re2 = regexp.MustCompile(`\s`)
	rawDuration := re2.ReplaceAllString(value[1], ",")
	durationSlice := strings.Split(rawDuration, ",")

	if durationSlice[1] == "Week" || durationSlice[1] == "Weeks" {

		tempVal := durationSlice[0] + " (" + durationSlice[1] + ")"
		cleanValue = append(cleanValue, tempVal)
	} else {
		cleanValue = append(cleanValue, durationSlice[0])
	}

	var re3 = regexp.MustCompile(`\s\/|\s|\s\+\s`)
	rawStripend := re3.ReplaceAllString(value[2], ",")
	stripendSlice := strings.Split(rawStripend, ",")

	incentive := ""
	val := ""

	if len(stripendSlice) >= 2 {
		if stripendSlice[1] == "week" {

			val = stripendSlice[0] + " (weekly)"

		} else if stripendSlice[1] == "month" {

			val = stripendSlice[0]
		}
	} else {
		val = stripendSlice[0]
	}

	cleanValue = append(cleanValue, val)

	if len(stripendSlice) == 5 {
		incentive = "Yes"
	}

	cleanValue = append(cleanValue, incentive)

	var re4 = regexp.MustCompile(`Part.`)
	partSplit := re4.ReplaceAllString(value[3], ",")

	cleanPartSplit := strings.Split(partSplit, ",")

	partTime := ""

	if len(cleanPartSplit) == 2 {

		partTime = "Yes"
		cleanValue = append(cleanValue, cleanPartSplit[0])
	} else {
		cleanValue = append(cleanValue, value[3])
	}

	cleanValue = append(cleanValue, partTime)

	fmt.Println(cleanValue)
	fmt.Println("***")
	return
}
