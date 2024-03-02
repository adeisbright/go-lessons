package topics

import (
	"fmt"
	"time"
)

func Timer() {
	fmt.Println("Working with Time")
	now := time.Now()
	fmt.Println(now, now.Month(), now.Year())
	//Given a date string, you can parse it to a Time.time object
	dateString := "2024-03-02"
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error on Parsing Date : ", err)
		return
	}
	fmt.Println("Parsed Date : ", parsedDate)
}
