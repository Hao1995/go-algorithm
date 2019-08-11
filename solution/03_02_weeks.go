package solution

import (
	"fmt"
	"time"
)

// 37%
// case
// (2019,"August", "September", "Thursday")
// (2018,"April", "May", "Monday")
// (2018,"September", "November", "Monday")
// (2019,"April", "May", "Thursday")
// (2019,"May", "June", "Thursday")
func Weeks(Y int, A string, B string, W string) int {
	// write your code in Go 1.4
	monDict := make(map[string]time.Month)
	monDict["January"] = time.January
	monDict["February"] = time.February
	monDict["March"] = time.March
	monDict["April"] = time.April
	monDict["May"] = time.May
	monDict["June"] = time.June
	monDict["July"] = time.July
	monDict["August"] = time.August
	monDict["September"] = time.September
	monDict["November"] = time.November
	monDict["December"] = time.December

	A1thDay := time.Date(Y, monDict[A], 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(A1thDay)
	A1thWeekday := int(A1thDay.Weekday())
	Start := A1thDay
	if A1thWeekday != 1 {
		Start = A1thDay.AddDate(0, 0, 7-A1thWeekday+1)
	}
	fmt.Println(Start)

	BlastDay := time.Date(Y, monDict[B], 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0).AddDate(0, 0, -1)
	B1thWeekday := int(BlastDay.Weekday())
	End := BlastDay.AddDate(0, 0, -B1thWeekday)
	fmt.Println(End)

	duration := End.AddDate(0, 0, 1).Sub(Start) // Add 1 day is in order to include last data
	return int(duration.Hours() / 24 / 7)
}
