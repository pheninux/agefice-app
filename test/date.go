package main

import (
	"fmt"

	"time"
)

func main() {

	const (
		_ = iota
		January
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)

	var month map[string]int = map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12}

	const (
		DATE_DEB = "2020-04-04"
		DATE_FIN = "2020-06-07"
		LAYOUT   = "2006-01-02"
	)

	tdebut, err := time.Parse(LAYOUT, DATE_DEB)
	tfin, err := time.Parse(LAYOUT, DATE_FIN)
	if err != nil {
		fmt.Println(err)
	}

	// The leap year 2016 had 366 days.
	t1 := Date(tdebut.Year(), month[tdebut.Month().String()], tdebut.Day())
	t2 := Date(tfin.Year(), month[tfin.Month().String()], tfin.Day())
	days := t2.Sub(t1).Hours() / 8760
	fmt.Println(days) // 366

	fmt.Println("time since =>", time.Since(tdebut).Hours()/8760)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
