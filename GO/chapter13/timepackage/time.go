package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("currentTime : ", currentTime)
	fmt.Printf("Type of currentTime is %T\n", currentTime)
	formatted := currentTime.Format("02-01-2006 15:04:05 Monday")
	fmt.Println("formatted : ", formatted)

	// Corrected layout string to match the date string format
	layoutStr := "02/01/2006"
	dateStr := "25/11/2030"
	formattedTime, err := time.Parse(layoutStr, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	fmt.Println("formatted_time : ", formattedTime)
	//add  1 more  day in the current time
	new_date := currentTime.Add(time.Hour * 24)
	fmt.Println("new_date : ", new_date)
	formatted_newdate := new_date.Format("02-01-2006 15:04:05 Monday")
	fmt.Println("formatted_newdate : ", formatted_newdate)
}
