package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	// "github.com/gobwas/glob/util/strings"
)

var internship struct {
	JobType, Company, Location, Duration, Stripend, Incentives, ApplyBy, PartTime, Start string
}

func readcsv() {

	csvfile, err := os.Open("data.csv")
	if err != nil {
		log.Fatal("Error is : ", err)
	}

	r := csv.NewReader(csvfile)

	file_name := "data2.csv"
	file, err := os.Create(file_name)

	if err != nil {
		log.Fatalf("Error Occured :%q while creatig file %s", err, file_name)
	}

	defer file.Close()
	writer := csv.NewWriter(file)

	writer.Write([]string{
		"Job-Type", "Comany", "Location", "Duration (Months)", "Stripend (monthly)", "Incentives", "Apply-By", "Part-Time", "Start",
	})

	for {
		record, err := r.Read()

		if err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			log.Fatal("Error 2 is : ", err)
		}

		var re = regexp.MustCompile(`\s\s\s+`)

		row1 := re.ReplaceAllString(record[0], ",")
		rowSlice1 := strings.Split(row1, ",")

		internship.JobType = rowSlice1[0]
		internship.Company = rowSlice1[1]
		internship.Location = rowSlice1[2]

		row2 := re.ReplaceAllString(record[2], ",")
		rowSlice2 := strings.Split(row2, ",")

		cleanRowSlice2 := cleanData(rowSlice2)

		internship.Start = cleanRowSlice2[0]
		internship.Duration = cleanRowSlice2[1]
		internship.Stripend = cleanRowSlice2[2]
		internship.Incentives = cleanRowSlice2[3]
		internship.ApplyBy = cleanRowSlice2[4]
		internship.PartTime = cleanRowSlice2[5]

		writer.Write([]string{
			internship.JobType, internship.Company, internship.Location, internship.Duration,
			internship.Stripend, internship.Incentives, internship.ApplyBy, internship.PartTime, internship.Start,
		})
	}

	writer.Flush()

	fmt.Println("*** Cleaning Completed ***")
}
