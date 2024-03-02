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
	//You can format a time.Time object into a string using time.Format
	formatNow := now.Format("2006-01-02") //now.Format(time.Layout)
	fmt.Println("Formatted Date:", formatNow)
	//Manipulating dates
	twoDaysLater := now.AddDate(0, 0, 2)
	oneMonthAhead := now.AddDate(0, 1, 0)
	timeDiff := oneMonthAhead.Sub(twoDaysLater)
	result := timeDiff.Hours() / 24
	fmt.Println("Two Days Later : ", twoDaysLater)
	fmt.Println("Time Difference: ", result)
	fmt.Println(now)

	//You can parse time with time zones , convert time from different time zones
	dateString = "2024-03-02T15:04:05-07:00"
	parsedDate, err = time.Parse(time.RFC3339, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	fmt.Println("Parsed Date:", parsedDate)

	//Converting From one Time Zone to Another
	timeZoneLocation, _ := time.LoadLocation("Africa/Nairobi")
	currentDate := time.Now().In(timeZoneLocation)
	fmt.Println("Current Time in Kenya is : ", currentDate.Format("2006-01-02 15:04:05"))
	//You can use the time duration for working with time
	duration := time.Hour * 2 //Two hours ahead
	futureTime := time.Now().Add(time.Duration(duration))
	fmt.Println("Future Time is :", futureTime)

	//You can execute code after a certain duration using the time.Timer
	timer := time.NewTimer(time.Second * 10)
	<-timer.C
	fmt.Println("This waits till timer expired")

}
