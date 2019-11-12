package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/textlinebuilder"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
	"time"
)


type TzStrFmt struct {
	InputStr     string
	OutputStr    string
}

// ElapsedTime - Calculates and formats time duration. Usually this means
// subtracting the starting time from the ending time and returning the
// duration result as a string.
func (tzFmtStr TzStrFmt) ElapsedTime(starTime, endTime time.Time) string {

	// microSecondNanoseconds - Number of Nanoseconds in a Microsecond
	// 	A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
	microSecondNanoseconds := int64(time.Microsecond)

	// milliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
	//	 A millisecond is 1/1,000 or 1 one-thousandth of a second
	milliSecondNanoseconds := int64(time.Millisecond)

	// secondNanoseconds - Number of Nanoseconds in a Second
	secondNanoseconds := int64(time.Second)

	// minuteNanoseconds - Number of Nanoseconds in a minute
	minuteNanoseconds := int64(time.Minute)

	// hourNanoseconds - Number of Nanoseconds in an hour
	hourNanoseconds := int64(time.Hour)

	dayNanoseconds := int64(time.Hour) * int64(24)

	t2Dur := endTime.Sub(starTime)

	str := ""

	totalNanoseconds := t2Dur.Nanoseconds()

	var numOfDays, numOfHours, numOfMinutes, numOfSeconds, numOfMilliseconds,
		numOfMicroseconds, numOfNanoseconds int64

	if totalNanoseconds >= dayNanoseconds {
		numOfDays = totalNanoseconds / dayNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfDays * dayNanoseconds)
	}

	if totalNanoseconds >= hourNanoseconds {
		numOfHours = totalNanoseconds / hourNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfHours * hourNanoseconds)
	}

	if totalNanoseconds >= minuteNanoseconds {
		numOfMinutes = totalNanoseconds / minuteNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMinutes * minuteNanoseconds)
	}

	if totalNanoseconds >= secondNanoseconds {
		numOfSeconds = totalNanoseconds / secondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfSeconds * secondNanoseconds)
	}

	if totalNanoseconds >= secondNanoseconds {
		numOfSeconds = totalNanoseconds / secondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfSeconds * secondNanoseconds)
	}

	if totalNanoseconds >= milliSecondNanoseconds {
		numOfMilliseconds = totalNanoseconds / milliSecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMilliseconds * milliSecondNanoseconds)
	}

	if totalNanoseconds >= microSecondNanoseconds {
		numOfMicroseconds = totalNanoseconds / microSecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMicroseconds * microSecondNanoseconds)
	}

	numOfNanoseconds = totalNanoseconds

	if numOfDays > 0 {
		str += fmt.Sprintf("%v-Days ", numOfDays)
	}

	if numOfHours > 0 || str != "" {

		str += fmt.Sprintf("%v-Hours ", numOfHours)

	}

	if numOfMinutes > 0 || str != "" {

		str += fmt.Sprintf("%v-Minutes ", numOfMinutes)

	}

	str += fmt.Sprintf("%v-Seconds ", numOfSeconds)

	str += fmt.Sprintf("%v-Milliseconds ", numOfMilliseconds)

	str += fmt.Sprintf("%v-Microseconds ", numOfMicroseconds)

	str += fmt.Sprintf("%v-Nanoseconds", numOfNanoseconds)

	return str
}

// WriteAlphabetizedTimeZoneList - This method formats and
// writes an alphabetized list of time zones to a designated
// file.
func (tzFmtStr TzStrFmt) WriteAlphabetizedTimeZoneList(
	outputFileMgr pathfileops.FileMgr,
	leftMargin textlinebuilder.MarginSpec,
	lineBreakStr textlinebuilder.LineSpec,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {


	ePrefix += "TzLogOps.WriteTimeZones() "

	tzStats.IanaCapturedTimeZones.SortByWorldRegion()

	numOfTimeZones := tzStats.IanaCapturedTimeZones.GetNumberOfTimeZones()

	if numOfTimeZones < 1 {

		return fmt.Errorf(ePrefix +
			"\nError: Number of tzStats.IanaCapturedTimeZones='%v'\n", numOfTimeZones)
	}

	newLine := textlinebuilder.NewLineSpec{AddNewLine:true}

	b := strings.Builder{}

	b.Grow(5120)

	label := "IANA Time Zones Listed in Alphabetical Order"

	strSpec2 := textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: len(label),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	leftSpacer := textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: 10,
		MarginChar:   ' ',
	}

	err := textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		leftSpacer,
		strSpec2,
		newLine,
		leftMargin,
		lineBreakStr,
		newLine,
		leftMargin,
		newLine)

	if err != nil {
		return err
	}

	if tzStats.TotalIanaStdTzLinkZones != numOfTimeZones {
		return fmt.Errorf(ePrefix +
			"\nError: Expected Number of Total Iana Time Zones: %v\n" +
			"Actual Number of Captured Iana Time Zones: \n",
			tzStats.TotalIanaStdTzLinkZones, numOfTimeZones)
	}

	label = "Expected Number of Total Iana Time Zones: "
	label2 := "Actual Number of Captured Iana Time Zones: "
	maxLen := len(label2)

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: maxLen,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}


	intSpec1 := textlinebuilder.IntegerSpec{
		NumericValue:       tzStats.TotalIanaStdTzLinkZones,
		NumericFieldSpec:   "%4d",
		NumericFieldLength: 4,
		NumericPadChar:     ' ',
		NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		strSpec2,
		intSpec1,
		newLine,
		leftMargin,
		newLine)

	if err != nil {
		return err
	}


	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       label2,
		StrFieldLength: maxLen,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}


	intSpec1 = textlinebuilder.IntegerSpec{
		NumericValue:       numOfTimeZones,
		NumericFieldSpec:   "%4d",
		NumericFieldLength: 4,
		NumericPadChar:     ' ',
		NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		strSpec2,
		intSpec1,
		newLine,
		leftMargin,
		newLine,
		leftMargin,
		newLine)

	if err != nil {
		return err
	}

	label = "Item"

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: 6,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	label = "Region"

	strSpec3 := textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: 10,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}


	strSpec4 := textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: 8,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	label = "Time Zone"

	strSpec5 := textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: 18,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	leftSpacer.MarginLength = 5

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		leftSpacer,
		strSpec2,
		strSpec3,
		strSpec4,
		strSpec5,
		newLine)

	if err != nil {
		return err
	}

	label = "No."

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: 6,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	label = "Index"

	strSpec3 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: 9,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify() }

	label = "Name"

	strSpec4 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: 10,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.Center(),
	}

	label = "Canonical Value"

	strSpec5 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: 23,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}


	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		leftSpacer,
		strSpec2,
		strSpec3,
		strSpec4,
		strSpec5,
		newLine,
		leftMargin,
		lineBreakStr,
		newLine,
		leftMargin,
		newLine)

	if err != nil {
		return err
	}


	var intRegion textlinebuilder.IntegerSpec

	writeCnt := 0

	rightSpacer := textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: 5,
		MarginChar:   ' ',
	}

	numOfWorldRegions := len(tzdatastructs.WorldRegions)

	for i :=0; i < numOfTimeZones; i++ {

		tz, err := tzStats.IanaCapturedTimeZones.PeekPtr(i)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by tzStats.IanaCapturedTimeZones.PeekPtr(i)\n" +
				"i='%v'\nError='%v'\n", i, err.Error())
		}

		leftSpacer = textlinebuilder.MarginSpec{
			MarginStr:    "",
			MarginLength: 5,
			MarginChar:   ' ',
		}

		writeCnt++

		intSpec1 = textlinebuilder.IntegerSpec{
			NumericValue:       writeCnt,
			NumericFieldSpec:   "%4d.",
			NumericFieldLength: 5,
			NumericPadChar:     ' ',
			NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
		}

		intRegion = textlinebuilder.IntegerSpec{
			NumericValue:       tz.WorldRegionSortCode,
			NumericFieldSpec:   "%3d",
			NumericFieldLength: 3,
			NumericPadChar:     ' ',
			NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
		}

		strSpec3 = textlinebuilder.StringSpec{
			StrValue:       tz.TzCanonicalValue,
			StrFieldLength: len(tz.TzCanonicalValue),
			StrPadChar:     ' ',
			StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
		}

		if tz.WorldRegionSortCode < 0 ||
			tz.WorldRegionSortCode >= numOfWorldRegions {
			label = "**Error**"
		} else {
			label = tzdatastructs.WorldRegions[tz.WorldRegionSortCode]
		}

		strSpec2 = textlinebuilder.StringSpec{
			StrValue:       label,
			StrFieldLength: 15,
			StrPadChar:     ' ',
			StrPosition:    textlinebuilder.FieldPos.LeftJustify()}

		err = textlinebuilder.TextLineBuilder{}.Build(
			&b,
			ePrefix,
			leftMargin,
			leftSpacer,
			intSpec1,
			rightSpacer,
			intRegion,
			rightSpacer,
			strSpec2,
			strSpec3,
			newLine)

		if err != nil {
			return err
		}
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		newLine,
		leftMargin,
		lineBreakStr,
		newLine,
		leftMargin,
		newLine,
		leftMargin,
		newLine)

	if err != nil {
		return err
	}

	_, err = outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = outputFileMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}
