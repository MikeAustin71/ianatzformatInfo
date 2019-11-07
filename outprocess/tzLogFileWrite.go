package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/textlinebuilder"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
	"time"
)

type TzLogOps struct {
	Input string
	Output string
	maxLineLen              int
	dashLineBreakStr        textlinebuilder.LineSpec
	equalLineBreakStr       textlinebuilder.LineSpec
	leftMarginLength        int
	leftMargin              textlinebuilder.MarginSpec
	newLine                 textlinebuilder.NewLineSpec
}

func (tzLog *TzLogOps) WriteLogFile(
		outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.WriteLogFile() "

	var err error

	tzLog.InitializeLogOps()

	err = tzLog.WriteLogHeader(outputFileMgr, tzStats, ePrefix)

	if err != nil {
		_ = outputFileMgr.CloseThisFile()
		return err
	}

	err = tzLog.WriteSummaryTotals(outputFileMgr, tzStats, ePrefix)

	if err != nil {
		_ = outputFileMgr.CloseThisFile()
		return err
	}

	err = tzLog.WriteIanaRegionalTotals(outputFileMgr, tzStats, ePrefix)

	if err != nil {
		_ = outputFileMgr.CloseThisFile()
		return err
	}

	err = outputFileMgr.CloseThisFile()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.CloseThisFile()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

// Initialize the TzLogOps base data fields
func (tzLog *TzLogOps) InitializeLogOps() {

	tzLog.leftMarginLength = 2
	tzLog.maxLineLen = 65

	tzLog.dashLineBreakStr = textlinebuilder.LineSpec{
		LineChar:         '-',
		LineLength:       tzLog.maxLineLen,
		LineFieldLength:  tzLog.maxLineLen,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.LeftJustify(),
	}

	tzLog.equalLineBreakStr = textlinebuilder.LineSpec{
		LineChar:         '=',
		LineLength:       tzLog.maxLineLen,
		LineFieldLength:  tzLog.maxLineLen,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.LeftJustify(),
	}

	tzLog.leftMargin = textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: tzLog.leftMarginLength,
		MarginChar:   ' ',
	}

	tzLog.newLine.AddNewLine = true

}

// WriteLogHeader - Writes Log title, header and timing
// information at top of the Log File.
func (tzLog *TzLogOps) WriteLogHeader(
	outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.WriteLogHeader() "
	var b strings.Builder
	b.Grow(2048)

	leadingBlankLines := textlinebuilder.BlankLinesSpec{
		NumBlankLines: 2,
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
		textlinebuilder.BlankLinesSpec{NumBlankLines:2})

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
		StrValue:       tzStats.IanaVersion,
		StrFieldLength: len(tzStats.IanaVersion) + 1,
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
		textlinebuilder.BlankLinesSpec{NumBlankLines:2})

	if err != nil {
		return err
	}


	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Starting Date Time: ",
		StrFieldLength: 25,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	tzdatastructs.ApplicationEndDateTime = time.Now()
	currDateTimeStr := tzdatastructs.ApplicationStartDateTime.Format(tzdatastructs.FmtDateTime)

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       currDateTimeStr,
		StrFieldLength: tzLog.maxLineLen - 25,
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
		StrFieldLength: 25,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	endDateTimeStr := tzdatastructs.ApplicationEndDateTime.Format(tzdatastructs.FmtDateTime)

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       endDateTimeStr,
		StrFieldLength: tzLog.maxLineLen - 25,
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
		StrFieldLength: 25,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	strSpec2 = textlinebuilder.StringSpec{
		StrValue:       elapsedTimeStr,
		StrFieldLength: tzLog.maxLineLen - 25,
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
		textlinebuilder.BlankLinesSpec{NumBlankLines:2})

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

func (tzLog *TzLogOps) WriteIanaRegionalTotals(
	outputFileMgr pathfileops.FileMgr,
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
		NumericValue:       tzStats.TotalZones,
		NumericFieldSpec:   "%4d",
		NumericFieldLength: 4,
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
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

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

// WriteSummaryTotals - Writes the time zone element totals
// to the log file.
//
func (tzLog *TzLogOps) WriteSummaryTotals(
	outputFileMgr pathfileops.FileMgr,
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
	int2FieldSpec := "%" + string(int2FieldLen) + "d"

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
		spacerFieldLen - 2

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
		intSpec2,
		textlinebuilder.BlankLinesSpec{NumBlankLines:3})

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
