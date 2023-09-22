package main

import (
	"flag"
	"fmt"
	"github.com/nathan-osman/go-sunrise"
	"time"
)

func main() {

	var year = flag.Int("year", 2023, "year for sunrise calendar generation")
	flag.Parse()

	monthLastDay := map[time.Month]int{
		time.January:   31,
		time.February:  28,
		time.March:     31,
		time.April:     30,
		time.May:       31,
		time.June:      30,
		time.July:      31,
		time.August:    31,
		time.September: 30,
		time.October:   31,
		time.November:  30,
		time.December:  31,
	}

	loc := time.FixedZone("UTC+3", +3*60*60) // PR timezone

	for month := time.January; month <= time.December; month++ {
		for day := 1; day <= monthLastDay[month]; day++ {

			rise, _ := sunrise.SunriseSunset(
				47.897683, 33.404919, // uk,kr
				*year, month, day,
			)

			fmt.Println(rise.In(loc).Format("20060102: {'hour': 3, 'minute': 4},"))
		}
	}
}
