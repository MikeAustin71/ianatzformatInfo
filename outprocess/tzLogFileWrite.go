package outprocess

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/fileops"
	"local.com/amarillomike/ianatzformatInfo/textlinebuilder"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
	"time"
)

type TzLogOps struct {
	Input string
	Output string
	dashLineBreakStr        textlinebuilder.LineSpec
	equalLineBreakStr       textlinebuilder.LineSpec
	errorLineBreakStr       textlinebuilder.LineSpec
	leftMarginLength        int
	leftMargin              textlinebuilder.MarginSpec
	maxLineLen              int
	newLine                 textlinebuilder.NewLineSpec
	outputFileMgr pathfileops.FileMgr
}

// Initialize the TzLogOps base data fields, create
// the Log File Manager and write the Log File Header
func (tzLog *TzLogOps) InitializeLogOps(
	zoneInfoDataDto *tzdatastructs.ZoneInfoDataDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.InitializeLogOps() "

	tzLog.leftMarginLength = 2
	tzLog.maxLineLen = 78

	tzLog.errorLineBreakStr = textlinebuilder.LineSpec{
		LineChar:         '*',
		LineLength:       tzLog.maxLineLen,
		LineFieldLength:  tzLog.maxLineLen,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.LeftJustify(),
	}

	tzLog.dashLineBreakStr = textlinebuilder.LineSpec{
		LineChar:         '-',
		LineLength:       tzLog.maxLineLen,
		LineFieldLength:  tzLog.maxLineLen + tzLog.leftMarginLength,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.RightJustify(),
	}

	tzLog.equalLineBreakStr = textlinebuilder.LineSpec{
		LineChar:         '=',
		LineLength:       tzLog.maxLineLen,
		LineFieldLength:  tzLog.maxLineLen + tzLog.leftMarginLength,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.RightJustify(),
	}

	tzLog.leftMargin = textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: tzLog.leftMarginLength,
		MarginChar:   ' ',
	}

	tzLog.newLine.AddNewLine = true

	var err error

	tzLog.outputFileMgr, err = tzLog.createOpenLogOutputFile(zoneInfoDataDto.AppOutputDirMgr, ePrefix)

	if err != nil {
		return err
	}

	err = tzLog.WriteHeader(zoneInfoDataDto, ePrefix)

	if err != nil {
		_ = tzLog.outputFileMgr.CloseThisFile()
		return err
	}

	return err
}

// TestCapturedIanaTimeZones -  This method attempts to initialize
// captured Iana Time Zones using the current 'Go' zoneinfo.zip
// database. If a time zone fails to load a warning message is
// recorded in the log file.
//
func (tzLog *TzLogOps) TestCapturedIanaTimeZones(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzLogOpsTestCapturedIanaTimeZones() "

	tzStats.IanaCapturedTimeZones.SortByGroupTzName(false)

	numIanaTimeZones := tzStats.IanaCapturedTimeZones.GetNumberOfTimeZones()

	if numIanaTimeZones < 1 {
		return fmt.Errorf(ePrefix +
			"\nError: The number of IANA Captured Time Zones is %v\n",
			numIanaTimeZones)
	}

	for i:=0; i < numIanaTimeZones; i++ {

		tz, err := tzStats.IanaCapturedTimeZones.PeekPtr(i)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by tzStats.IanaCapturedTimeZones.PeekPtr(i)\n" +
				"i='%v'\nError='%v'\n", i, err.Error())
		}

		_, err = time.LoadLocation(tz.TzCanonicalValue)

		if err != nil {
			warningMsg := fmt.Sprintf(
				"\nError Returned loading IANA time zone!\n" +
				"time.LoadLocation(tz.TzCanonicalValue)\n" +
				"Time Zone (tz.TzCanonicalValue)='%v'\n" +
				"Error='%v'\n", tz.TzCanonicalValue, err.Error())

			_ = tzLog.WriteWarning(warningMsg, ePrefix)
		}

	}

	return nil
}

// TestIanaTimeZoneAbbreviations - Tests to ensure that all captured
// iana time zones have valid Time Zone Abbreviations.
func (tzLog *TzLogOps) TestIanaTimeZoneAbbreviations(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

		ePrefix += "TzLogOps.TestIanaTimeZoneAbbreviations() "

	numIanaTimeZones := tzStats.IanaCapturedTimeZones.GetNumberOfTimeZones()

	var err error
	var tz *tzdatastructs.TimeZoneDataDto
	var testLocation *time.Location
	var testTime time.Time
	var testTimeStr, tzAbbrv, utcOffset, tzAbbrvId,
	firstLetter, warningMsg string
	var testTimeSplitAry, tzCanonicalValues, tzAbbrvValues []string
	var ok bool
	var tzAbbrvDto tzdatastructs.TzAbbreviationDto

	for i:=0; i < numIanaTimeZones; i++ {

		tz, err = tzStats.IanaCapturedTimeZones.PeekPtr(i)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by tzStats.IanaCapturedTimeZones.PeekPtr(i)\n" +
				"i='%v'\n" +
				"Error='%v'\n", i, err.Error())
		}

		// FmtDateTimeTzYM

		testLocation, err = time.LoadLocation(tz.TzCanonicalValue)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by time.LoadLocation(tz.TzCanonicalValue).\n"+
				"tzLocation='%v'\n"+
				"Error:'%v'\n", tz.TzCanonicalValue, err.Error())
		}

		// Test Summer Time
		testTime, err =
			time.ParseInLocation("2006-01-02 15:04:00", "2019-06-25 16:50:00", testLocation)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by summer time.Parse(\"2006-01-02 15:04:00\", \"2019-06-25 16:50:00\", " +
				"testLocation).\n" +
				"testLocation='%v'\n"+
				"Error:'%v'\n", tz.TzCanonicalValue, err.Error())
		}

		testTimeStr = testTime.Format(tzdatastructs.FmtDateTimeTzYM)

		testTimeSplitAry = strings.Split(testTimeStr, " ")

		if len(testTimeSplitAry) != 4 {
			return fmt.Errorf(ePrefix +
				"\nError: Expected Test Summer Time to consist of 4-Elements.\n" +
				"Instead, Test Summer Time Splits as %v-Elements!\n", len(testTimeSplitAry))
		}

		tzAbbrv = strings.TrimRight(strings.TrimLeft(testTimeSplitAry[3], " "), " ")

		firstLetter = tzAbbrv[0:1]

		if firstLetter == "+" ||
			firstLetter == "-"{
			goto winterTimeTest
		}

		utcOffset = strings.TrimRight(strings.TrimLeft(testTimeSplitAry[2], " "), " ")

		tzAbbrvId = tzAbbrv+utcOffset

		tzAbbrvDto, ok = tzdatastructs.TzAbbreviationReference[tzAbbrvId]

		if !ok {
			warningMsg = fmt.Sprintf(
				"\nCould not locate time zone abbreviation for time zone: '%v'\n" +
					"Expected Time Zone Abbreviation: '%v'\n" +
					"UTC Offset: '%v'\n" +
					"Summer Time: '%v'",
					tz.TzCanonicalValue, tzAbbrv, utcOffset, testTimeStr)

			_ = tzLog.WriteWarning(warningMsg, ePrefix)

			continue
		}

		tzAbbrvDto.IanaZone = tz.TzCanonicalValue

		tzStats.TzAbbreviations.AddIfNew(tzAbbrvDto)

		tzCanonicalValues, ok = tzStats.MapTzAbbrvsToTimeZones[tzAbbrvId]

		if !ok {
			tzCanonicalValues = make([]string, 0)
		}

		tzCanonicalValues = append(tzCanonicalValues, tz.TzCanonicalValue)
		tzStats.MapTzAbbrvsToTimeZones[tzAbbrvId] = tzCanonicalValues

		tzAbbrvValues, ok = tzStats.MapTimeZonesToTzAbbrvs[tz.TzCanonicalValue]

		if !ok {
			tzAbbrvValues = make([]string, 0)
		}

		tzAbbrvValues = append(tzAbbrvValues, tzAbbrvId)

		tzStats.MapTimeZonesToTzAbbrvs[tz.TzCanonicalValue] = tzAbbrvValues

	winterTimeTest:
		// Test Winter Time
		testTime, err =
			time.ParseInLocation("2006-01-02 15:04:00", "2019-12-23 16:50:00", testLocation)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by summer time.Parse(\"2006-01-02 15:04:00\", \"2019-12-23 16:50:00\", " +
				"testLocation).\n" +
				"testLocation='%v'\n"+
				"Error:'%v'\n", tz.TzCanonicalValue, err.Error())
		}

		testTimeStr = testTime.Format(tzdatastructs.FmtDateTimeTzYM)

		testTimeSplitAry = strings.Split(testTimeStr, " ")

		if len(testTimeSplitAry) != 4 {
			return fmt.Errorf(ePrefix +
				"\nError: Expected Test Winter Time to consist of 4-Elements.\n" +
				"Instead, Test Summer Time Splits as %v-Elements!\n", len(testTimeSplitAry))
		}

		tzAbbrv = strings.TrimRight(strings.TrimLeft(testTimeSplitAry[3], " "), " ")

		firstLetter = tzAbbrv[0:1]

		if firstLetter == "+" ||
			firstLetter == "-"{
			continue
		}

		utcOffset = strings.TrimRight(strings.TrimLeft(testTimeSplitAry[2], " "), " ")

		tzAbbrvId = tzAbbrv+utcOffset

		tzAbbrvDto, ok = tzdatastructs.TzAbbreviationReference[tzAbbrvId]

		if !ok {
			warningMsg = fmt.Sprintf(
				"\nCould not locate time zone abbreviation for time zone: '%v'\n" +
					"Expected Time Zone Abbreviation: '%v'\n" +
					"UTC Offset: '%v'\n" +
					"Summer Time: '%v'",
				tz.TzCanonicalValue, tzAbbrv, utcOffset, testTimeStr)

			_ = tzLog.WriteWarning(warningMsg, ePrefix)

			continue
		}

		tzAbbrvDto.IanaZone = tz.TzCanonicalValue

		tzStats.TzAbbreviations.AddIfNew(tzAbbrvDto)

		tzAbbrvDto.IanaZone = tz.TzCanonicalValue

		tzStats.TzAbbreviations.AddIfNew(tzAbbrvDto)

		tzCanonicalValues, ok = tzStats.MapTzAbbrvsToTimeZones[tzAbbrvId]

		if !ok {
			tzCanonicalValues = make([]string, 0)
		}

		tzCanonicalValues = append(tzCanonicalValues, tz.TzCanonicalValue)
		tzStats.MapTzAbbrvsToTimeZones[tzAbbrvId] = tzCanonicalValues

		tzAbbrvValues, ok = tzStats.MapTimeZonesToTzAbbrvs[tz.TzCanonicalValue]

		if !ok {
			tzAbbrvValues = make([]string, 0)
		}

		tzAbbrvValues = append(tzAbbrvValues, tzAbbrvId)

		tzStats.MapTimeZonesToTzAbbrvs[tz.TzCanonicalValue] = tzAbbrvValues

	}

	return nil
}


// WriteError - Writes an error message to the log
// file
func (tzLog *TzLogOps) WriteError(
	errMsg error,
	ePrefix string) error {

		ePrefix += "TzLogOps.WriteError() "

		errStr := errMsg.Error()
		lenErrStr := len(errStr)

		if lenErrStr == 0 {
			errs := make([]error, 2)

			errs[0] = errMsg
			errs[1] = errors.New("\n" + ePrefix + "Error message is Empty! Zero string length!\n")

			return pathfileops.FileHelper{}.ConsolidateErrors(errs)
		}

		b := strings.Builder{}
		b.Grow(lenErrStr + 5)

		label := "Error"

	strSpec1 := textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: len(label),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	strSpec2 := textlinebuilder.StringSpec{
		StrValue:       errStr,
		StrFieldLength: lenErrStr,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err := textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		tzLog.errorLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.leftMargin,
		textlinebuilder.MarginSpec{
				MarginStr:    "",
				MarginLength: 10,
				MarginChar:   ' ',
			},
		strSpec1,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.errorLineBreakStr,
		tzLog.leftMargin,
		strSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.leftMargin,
		tzLog.errorLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	_, err = tzLog.outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzLog.outputFileMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	tzdatastructs.ErrorCount++

	return nil
}

func (tzLog *TzLogOps) WriteFooter(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.WriteFooter() "
	var b strings.Builder
	b.Grow(1024)

	label1 := "End Of Execution"

	strSpec1 := textlinebuilder.StringSpec{
		StrValue:       label1,
		StrFieldLength: tzLog.maxLineLen,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.Center(),
	}

	strSpec2 := textlinebuilder.StringSpec{
		StrValue:       "ianatzformatInfo.go",
		StrFieldLength: tzLog.maxLineLen,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.Center(),
	}

	err := textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.equalLineBreakStr,
		tzLog.newLine,
		strSpec1,
		tzLog.newLine,
		strSpec2,
		tzLog.newLine,
		tzLog.equalLineBreakStr,
		tzLog.newLine,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2})

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Iana Time Zone Version: ",
		StrFieldLength: 25,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       tzStats.IanaVersion,
		StrFieldLength: len(tzStats.IanaVersion),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		strSpec2,
		tzLog.newLine,
		tzLog.dashLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Errors and Warnings",
		StrFieldLength: 50,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.Center(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		strSpec1,
		tzLog.newLine,
		tzLog.dashLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2})

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Error Count",
		StrFieldLength: 25,
		StrPadChar:     '.',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	intSpec3 := textlinebuilder.IntegerSpec{
		NumericValue:       tzdatastructs.ErrorCount,
		NumericFieldSpec:   "%3d",
		NumericFieldLength: 10,
		NumericPadChar:     '.',
		NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		intSpec3,
		tzLog.newLine)

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Warning Count",
		StrFieldLength: 25,
		StrPadChar:     '.',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	intSpec3 = textlinebuilder.IntegerSpec{
		NumericValue:       tzdatastructs.WarningCount,
		NumericFieldSpec:   "%3d",
		NumericFieldLength: 10,
		NumericPadChar:     '.',
		NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		intSpec3,
		tzLog.newLine,
		tzLog.dashLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Execution Times",
		StrFieldLength: 50,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.Center(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		strSpec1,
		tzLog.newLine,
		tzLog.dashLineBreakStr,
		tzLog.newLine)

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Starting Date Time: ",
		StrFieldLength: 20,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	currDateTimeStr := tzdatastructs.ApplicationStartDateTime.Format(tzdatastructs.FmtDateTimeTzNanoYMD)

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       currDateTimeStr,
		StrFieldLength: len(currDateTimeStr),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		strSpec2,
		tzLog.newLine )

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Ending Date Time: ",
		StrFieldLength: 20,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	tzdatastructs.ApplicationEndDateTime = time.Now()

	endDateTimeStr := tzdatastructs.ApplicationEndDateTime.Format(tzdatastructs.FmtDateTimeTzNanoYMD)

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       endDateTimeStr,
		StrFieldLength: len(endDateTimeStr),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		strSpec2,
		tzLog.newLine )

	if err != nil {
		return err
	}

	elapsedTimeStr :=
		TzStrFmt{}.ElapsedTime(
			tzdatastructs.ApplicationStartDateTime,
			tzdatastructs.ApplicationEndDateTime)

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Elapsed Time: ",
		StrFieldLength: 20,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       elapsedTimeStr,
		StrFieldLength: len(elapsedTimeStr),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		strSpec2,
		tzLog.newLine,
		tzLog.dashLineBreakStr,
		tzLog.newLine,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.equalLineBreakStr,
		tzLog.newLine,
		tzLog.equalLineBreakStr,
		tzLog.newLine)

	if err != nil {
		return err
	}

	_, err = tzLog.outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzLog.outputFileMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzLog.outputFileMgr.CloseThisFile()

	return err
}

// WriteHeader - Writes header information to the
// Log file.
func (tzLog *TzLogOps) WriteHeader(
	zoneInfoDataDto *tzdatastructs.ZoneInfoDataDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.WriteHeader() "
	var b strings.Builder
	b.Grow(2048)

	leadingBlankLines := textlinebuilder.BlankLinesSpec{
		NumBlankLines: 1,
	}

	var strSpec1, strSpec2 textlinebuilder.StringSpec

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "ianatzformatInfo.go",
		StrFieldLength: tzLog.maxLineLen,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.Center(),
	}

	err := textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leadingBlankLines,
		tzLog.leftMargin,
		strSpec1,
		tzLog.newLine,
		tzLog.equalLineBreakStr,
		tzLog.newLine,
		textlinebuilder.BlankLinesSpec{NumBlankLines:1})

	if err != nil {
		return err
	}


	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Execution Start Time: ",
		StrFieldLength: 25,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	startTime := tzdatastructs.ApplicationStartDateTime.Format(tzdatastructs.FmtDateTimeTzNanoYMD)

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       startTime,
		StrFieldLength: len(startTime) + 1,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		strSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.leftMargin,
		tzLog.equalLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	label:= "Base Data Input"

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: len(label),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		textlinebuilder.MarginSpec{
			MarginStr:    "",
			MarginLength: 10,
			MarginChar:   ' ',
		},
		strSpec1,
		tzLog.newLine,
		tzLog.leftMargin,
		tzLog.dashLineBreakStr,
		tzLog.newLine)

	if err != nil {
		return err
	}


	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "IANA Time Zone Version: ",
		StrFieldLength: 25,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       zoneInfoDataDto.IanaTimeZoneVersion,
		StrFieldLength: len(zoneInfoDataDto.IanaTimeZoneVersion) + 1,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		strSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2})

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Base Data Input File: ",
		StrFieldLength: 25,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	label = zoneInfoDataDto.AppInputPathFileNameExt

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: len(label) + 1,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		strSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2})

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Output Source File: ",
		StrFieldLength: 25,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	label = zoneInfoDataDto.AppOutputDirMgr.GetAbsolutePathWithSeparator() +
		tzdatastructs.TimeZoneDataOutputFileName

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: len(label) + 1,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		strSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2})

	if err != nil {
		return err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Log File: ",
		StrFieldLength: 25,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	label = tzLog.outputFileMgr.GetAbsolutePathFileName()

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: len(label) + 1,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		strSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.leftMargin,
		tzLog.dashLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	_, err = tzLog.outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzLog.outputFileMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

// WriteIanaRegionalTotals - Prints totals for IANA
// time zones by Region.
func (tzLog *TzLogOps) WriteIanaRegionalTotals(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.WriteIanaRegionalTotals() "


	b := strings.Builder{}
	b.Grow(2048)

	var strSpec1 textlinebuilder.StringSpec

	label1 := "Iana Time Zones by Region"

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       label1,
		StrFieldLength: tzLog.maxLineLen - 10,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.Center(),
	}

	lineSpec1 := textlinebuilder.LineSpec{
		LineChar:         '-',
		LineLength:       len(label1),
		LineFieldLength:  tzLog.maxLineLen - 10,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.Center(),
	}

	err := textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		tzLog.newLine,
		tzLog.leftMargin,
		lineSpec1,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	leftSpacer := textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: 5,
		MarginChar:   '.',
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "X-Value",
		StrFieldLength: 30,
		StrPadChar:     '.',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	rightSpacer := textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: 5,
		MarginChar:   '.',
	}

	intSpec2 := textlinebuilder.IntegerSpec{
		NumericValue:       -1,
		NumericFieldSpec:   "%4d",
		NumericFieldLength: 4,
		NumericPadChar:     '.',
		NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	lenRegionsAry := len(tzStats.IanaTzRegions)

	for i:=0; i < lenRegionsAry; i++ {

		strSpec1.StrValue = tzStats.IanaTzRegions[i]
		intSpec2.NumericValue = tzStats.IanaTzCounters[i]

		xErr := ePrefix + "\n" +
			"Error returned by TextLineBuilder{}.Build()\n" +
			fmt.Sprintf("Regional Totals Interation %v\n", i)

		err := textlinebuilder.TextLineBuilder{}.Build(
			&b,
			xErr,
			tzLog.leftMargin,
			leftSpacer,
			strSpec1,
			rightSpacer,
			intSpec2,
			tzLog.newLine)

		if err != nil{
			return err
		}

	}

	totalLineStarts :=
		tzLog.leftMargin.MarginLength +
			leftSpacer.MarginLength +
			strSpec1.StrFieldLength +
			rightSpacer.MarginLength +
			intSpec2.NumericFieldLength +
			5 - 2

	totalLineLen := intSpec2.NumericFieldLength + 4

	totalLineSpec := textlinebuilder.LineSpec{
		LineChar:         '-',
		LineLength:       totalLineLen,
		LineFieldLength:  totalLineLen,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		textlinebuilder.MarginSpec{
			MarginStr:    "",
			MarginLength: totalLineStarts,
			MarginChar:   ' ',
		},
		totalLineSpec,
		tzLog.newLine)

	if err != nil {
		return err
	}

	label1 = "Total Iana Time Zones "
	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       label1,
		StrFieldLength: len(label1),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	intSpec2 = textlinebuilder.IntegerSpec{
		NumericValue:       tzStats.TotalIanaStdTzLinkZones,
		NumericFieldSpec:   "%4d",
		NumericFieldLength: 5,
		NumericPadChar:     ' ',
		NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	totalLineStarts = totalLineStarts + 2 -
		len(label1) - 1


	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		textlinebuilder.MarginSpec{
			MarginStr:    "",
			MarginLength: totalLineStarts,
			MarginChar:   ' ',
		},
		strSpec1,
		intSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.dashLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	_, err = tzLog.outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzLog.outputFileMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

// WriteSummaryTotals - Writes the time zone element totals
// to the log file.
//
func (tzLog *TzLogOps) WriteSummaryTotals(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.WriteSummaryTotals() "


b := strings.Builder{}
b.Grow(2048)

var strSpec1 textlinebuilder.StringSpec

label1 := "Time Zones - Summary"

strSpec1 = textlinebuilder.StringSpec{
	StrValue:       label1,
	StrFieldLength: tzLog.maxLineLen - 10,
	StrPadChar:     ' ',
	StrPosition:    textlinebuilder.FieldPos.Center(),
}

lineSpec1 := textlinebuilder.LineSpec{
	LineChar:         '-',
	LineLength:       len(label1),
	LineFieldLength:  tzLog.maxLineLen - 10,
	LineFieldPadChar: ' ',
	LinePosition:     textlinebuilder.FieldPos.Center(),
}

	err := textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		tzLog.newLine,
		tzLog.leftMargin,
		lineSpec1,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	spec1FieldLen := 30
	spacerFieldLen := 5
	int2FieldLen := 4
	int2FieldSpec := "%4d"

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Iana Time Zones",
		StrFieldLength: spec1FieldLen,
		StrPadChar:     '.',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	spacerSpec1 := textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: spacerFieldLen,
		MarginChar:   '.',
	}

	intSpec2 := textlinebuilder.IntegerSpec{
		NumericValue:       tzStats.TotalIanaStdTzLinkZones,
		NumericFieldSpec:   int2FieldSpec,
		NumericFieldLength: int2FieldLen,
		NumericPadChar:     '.',
		NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		spacerSpec1,
		intSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2} )

	if err != nil {
		return err
	}


	strSpec1.StrValue = "Military Time Zones"
	intSpec2.NumericValue = tzStats.NumMilitaryTZones

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		spacerSpec1,
		intSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2} )

	if err != nil {
		return err
	}

	strSpec1.StrValue = "Other Non-Iana Time Zones"
	intSpec2.NumericValue = tzStats.NumOtherTZones


	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		strSpec1,
		spacerSpec1,
		intSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2} )

	if err != nil {
		return err
	}

	totalLineStarts :=
	tzLog.leftMargin.MarginLength +
		spec1FieldLen +
		spacerFieldLen  + 7 - 2

	totalLineLen := int2FieldLen + 4

	totalLineSpec := textlinebuilder.LineSpec{
		LineChar:         '-',
		LineLength:       totalLineLen,
		LineFieldLength:  totalLineLen,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.LeftJustify(),
	}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		textlinebuilder.MarginSpec{
			MarginStr:    "",
			MarginLength: totalLineStarts,
			MarginChar:   ' ',
		},
		totalLineSpec,
		tzLog.newLine)

	if err != nil {
		return err
	}

	label1 = "Total Time Zones "
	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       label1,
		StrFieldLength: len(label1),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	intSpec2 = textlinebuilder.IntegerSpec{
		NumericValue:       tzStats.TotalZones,
		NumericFieldSpec:   int2FieldSpec,
		NumericFieldLength: int2FieldLen,
		NumericPadChar:     ' ',
		NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	totalLineStarts = totalLineStarts + 2 -
		len(label1) - 1


	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		textlinebuilder.MarginSpec{
			MarginStr:    "",
			MarginLength: totalLineStarts,
			MarginChar:   ' ',
		},
		strSpec1,
		textlinebuilder.MarginSpec{
			MarginStr:    "",
			MarginLength: 1,
			MarginChar:   ' ',
		},
		intSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.dashLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	_, err = tzLog.outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzLog.outputFileMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

// WriteTimeZones - Writes all captured IANA Time Zones
// in alphabetical order to the log file.
func (tzLog *TzLogOps) WriteTimeZones(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.WriteTimeZones() "

	lineBreak := textlinebuilder.LineSpec{
		LineChar:         '-',
		LineLength:       78,
		LineFieldLength:  78,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.LeftJustify(),
	}

	return TzStrFmt{}.WriteAlphabetizedTimeZoneList(
		tzLog.outputFileMgr,
		tzLog.leftMargin,
		lineBreak,
		tzStats,
		ePrefix)
}

// WriteWarning - Writes an error message to the log
// file
func (tzLog *TzLogOps) WriteWarning(
	warningMsg string,
	ePrefix string) error {

	ePrefix += "TzLogOps.WriteError() "


	lenWarningStr := len(warningMsg)

	if lenWarningStr == 0 {
		return errors.New("\n" + ePrefix + "Warning message is Empty! Zero string length!\n")
	}

	b := strings.Builder{}
	b.Grow(lenWarningStr + 5)

	label := "Warning"

	strSpec1 := textlinebuilder.StringSpec{
		StrValue:       label,
		StrFieldLength: len(label),
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	strSpec2 := textlinebuilder.StringSpec{
		StrValue:       warningMsg,
		StrFieldLength: lenWarningStr,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	err := textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		tzLog.leftMargin,
		tzLog.errorLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.leftMargin,
		textlinebuilder.MarginSpec{
			MarginStr:    "",
			MarginLength: 10,
			MarginChar:   ' ',
		},
		strSpec1,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.errorLineBreakStr,
		tzLog.leftMargin,
		strSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:2},
		tzLog.leftMargin,
		tzLog.errorLineBreakStr,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

	if err != nil {
		return err
	}

	_, err = tzLog.outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzLog.outputFileMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	tzdatastructs.WarningCount++

	return nil
}

// createOpenLogOutputFile - Generates the log path and
// file name then creates and opens the file.
func (tzLog *TzLogOps) createOpenLogOutputFile(
	outputPathDirMgr pathfileops.DirMgr,
	ePrefix string) (pathfileops.FileMgr, error) {

	ePrefix += "TzLogOps.createOpenLogOutputFile() "

	fmtDateTimeSecondStr := "20060102150405"
	currDateTimeStr := tzdatastructs.ApplicationStartDateTime.Format(fmtDateTimeSecondStr)

	fileNameExt :=   currDateTimeStr +"_ianaformatInfoLog" +".txt"

	return fileops.FileOps{}.CreateOpenFile(outputPathDirMgr, fileNameExt, ePrefix)
}
