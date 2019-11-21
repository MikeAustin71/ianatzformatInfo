package main


import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/xapptest/libs"
	"time"
)

// 2019-11-15 Friday 09:37:19.061195900 +0930 ACST
// https://yourbasic.org/golang/format-parse-string-time-date-example/

func main() {

	mTest{}.parseTime03()

	return
}

var xFmtDateTimeTzNanoYMD = "2006-01-02 Monday 15:04:05.000000000 -0700 MST"

type mTest struct {
	input  string
	output string
}

func (mT mTest) parseTime03() {

	tz := "America/Buenos_Aires"
	tzLocation := tz

	location1, err := time.LoadLocation(tzLocation)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzLocation).\n"+
			"tzLocation='%v'\n"+
			"Error:'%v'\n", tzLocation, err.Error())
		return
	}

	// time.Parse("2006 01 02 15 04", "2015 11 11 16 50")
	tSummer, err := time.ParseInLocation("2006-01-02 15:04:00", "2019-07-15 16:50:00", location1)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(\"2006-01-02 15:04:00\", \"2015-11-11 16:50:00\").\n" +
			"Error:'%v'\n", err.Error())
		return
	}

	fmt.Println("mTest.parseTime03()")

	fmt.Println("Location: ", tz)
	fmt.Println("Summer Time: ", tSummer.Format(xFmtDateTimeTzNanoYMD))

	tWinter, err := time.ParseInLocation("2006-01-02 15:04:00", "2019-12-15 16:50:00", location1)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(\"2006-01-02 15:04:00\", \"2015-11-11 16:50:00\").\n" +
			"Error:'%v'\n", err.Error())
		return
	}

	fmt.Println("Winter Time: ", tWinter.Format(xFmtDateTimeTzNanoYMD))

}

func (mT mTest) abbrvTest02() {

	fmt.Println("mTest.abbrvTest02()")

	tzLocation, err := time.LoadLocation(libs.TZones.Australia.Darwin())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.Australia.Darwin()).\n" +
			"Error:'%v'\n", err.Error())
		return
	}

	now := time.Now()

	ausyTime := now.In(tzLocation).Format(xFmtDateTimeTzNanoYMD)

	fmt.Println("Austrialian Darwin Time: ", ausyTime)

}

func (mT mTest) ectTest01() {
	tzLocation := "Etc/GMT-1"

	location1, err := time.LoadLocation(tzLocation)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzLocation).\n" +
			"tzLocation='%v'\n" +
			"Error:'%v'\n", tzLocation, err.Error())
		return
	}

	nowTime := time.Now()

	fmt.Println("mTest.ectTest01()")
	fmt.Println(nowTime.In(location1))
	fmt.Printf("Attemped to load location1 '%v'\n",tzLocation )
	fmt.Printf("Actual Location='%v'\n", location1.String())

	return
}