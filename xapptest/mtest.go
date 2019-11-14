package main


import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/xapptest/libs"
	"time"
)


func main() {

	mTest{}.abbrvTest02()

	return
}

var xFmtDateTimeTzNanoYMD = "2006-01-02 Monday 15:04:05.000000000 -0700 MST"

type mTest struct {
	input  string
	output string
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