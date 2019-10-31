package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
	"time"
)

type TzLogOps struct {
	Input string
	Output string
	maxLineLen              int
	totalLineLen            int
	leftMarginLen           int
	dashLineBreakStr        string
	equalLineBreakStr       string
	leftMarginStr           string
	totalLineBreakStr       string
	grandTotalLineBreakStr  string
	label1Length            int
	spacer1Length           int
	num1FieldLength         int
	num1TotalLength         int
	total1FieldLength       int
	total1TotalLength       int
}

func (tzLog *TzLogOps) WriteLogFile(
		outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.WriteLogFile() "

	var err error

	tzLog.totalLineLen = 10
	tzLog.maxLineLen = 65
	tzLog.leftMarginLen = 2

	strOps := strops.StrOps{}

	tzLog.dashLineBreakStr, err = strOps.MakeSingleCharString('-', tzLog.maxLineLen)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n'dashLineBreak' error returned by strOps.MakeSingleCharString('-', tzLog.maxLineLen)\n" +
			"maxLineLen='%v'\n" +
			"Error='%v'\n", tzLog.maxLineLen, err.Error())
	}

	tzLog.equalLineBreakStr, err = strOps.MakeSingleCharString('=', tzLog.maxLineLen)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n'equalLineBreak' error returned by strOps.MakeSingleCharString('=', tzLog.maxLineLen)\n" +
			"maxLineLen='%v'\n" +
			"Error='%v'\n", tzLog.maxLineLen, err.Error())
	}

	tzLog.leftMarginStr, err = strOps.MakeSingleCharString(' ', tzLog.leftMarginLen)
	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n'leftMarginStr' error returned by strOps.MakeSingleCharString(' ', tzLog.leftMarginLen)\n" +
			"Error='%v'\n", err.Error())
	}

	tzLog. totalLineBreakStr, err = strOps.MakeSingleCharString('=', tzLog.totalLineLen)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"'totalLineBreak' error returned by strOps.MakeSingleCharString('=', tzLog.totalLineLen)\n" +
			"maxLineLen='%v'\n" +
			"Error='%v'\n", tzLog.maxLineLen, err.Error())
	}

	tzLog.grandTotalLineBreakStr, err = strOps.MakeSingleCharString('*', tzLog.totalLineLen)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"'grandTotalLineBreak' error returned by strOps.MakeSingleCharString('=', tzLog.totalLineLen)\n" +
			"maxLineLen='%v'\n" +
			"Error='%v'\n", tzLog.maxLineLen, err.Error())
	}

	err = tzLog.writeLogHeader(outputFileMgr, tzStats, ePrefix)

	if err != nil {
		_ = outputFileMgr.CloseThisFile()
		return err
	}



	return nil
}

// writeLogHeader - Writes Log title, header and timing
// information at top of the Log File.
func (tzLog *TzLogOps) writeLogHeader(
	outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

		strOps := strops.StrOps{}

	var b strings.Builder
	b.Grow(2048)

	outputStr, err :=
		strOps.StrCenterInStr("ianatzformatInfo.go", tzLog.maxLineLen)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error retuned by strOps.StrCenterInStr(\"ianatzformatInfo.go\", 65)\n" +
			"Error='%v'\n", err.Error())
	}

	b.WriteString(outputStr + "\n")

	b.WriteString(tzLog.dashLineBreakStr + "\n")

	tzdatastructs.ApplicationEndDateTime = time.Now()

	currDateTimeStr := tzdatastructs.ApplicationStartDateTime.Format(tzdatastructs.FmtDateTime)
	endDateTimeStr := tzdatastructs.ApplicationEndDateTime.Format(tzdatastructs.FmtDateTime)
	elapsedTimeStr :=
		TzStrFmt{}.ElapsedTime(
			tzdatastructs.ApplicationStartDateTime,
			tzdatastructs.ApplicationEndDateTime)

	b.WriteString(tzLog.leftMarginStr +
		"    Starting Date-Time: " +
		currDateTimeStr + "\n")

	b.WriteString(tzLog.leftMarginStr +
		"      Ending Date-Time: " +
		endDateTimeStr + "\n")

	b.WriteString(tzLog.leftMarginStr +
		"          Elapsed Time: " +
		elapsedTimeStr + "\n")

	b.WriteString(tzLog.dashLineBreakStr + "\n")

	b.WriteString(tzLog.leftMarginStr +
		"Iana Time Zone Version: " + tzStats.IanaVersion + "\n")
	b.WriteString(tzLog.dashLineBreakStr + "\n\n")

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

// writeSummaryTotals - Writes the time zone element totals
// to the log file.
//
func (tzLog *TzLogOps) writeSummaryTotals(
	outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzLogOps.writeSummaryTotals() "

	strOps := strops.StrOps{}

b := strings.Builder{}
b.Grow(2048)

	outputStr, err :=
		strOps.StrCenterInStr("Time Zones - Summary", tzLog.maxLineLen)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by strOps.StrCenterInStr(\"Time Zones - Summary\"), tzLog.maxLineLen)\n"+
			"Error='%v'\n", err.Error())
	}

	_, err = b.WriteString(outputStr + "\n")

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by b.WriteString(Time Zones - Summary)\n" +
			"Error='%v'\n", err.Error())
	}

	tzLog.label1Length = 35
	tzLog.num1TotalLength = 10
	tzLog.num1FieldLength = 4
	tzLog.total1TotalLength = 10
	tzLog.total1FieldLength = 4

	labelArray := make([]string, 0)

	labelArray = append(labelArray, "Iana Time Zones")

	fieldStr := "Iana Time Zones"

	outputStr, err = TzStrFmt{}.LeftJustifyField(fieldStr, labelLength, true, ePrefix )

	if err != nil{
		return err
	}

	_, err = b.WriteString(tzLog.leftMarginStr + outputStr + "\n")

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by b.WriteString(Iana Time Zones)\n" +
			"Error='%v'\n", err.Error())
	}

	outputStr, err =TzStrFmt{}.RightJustifyNum(
		tzStats.TotalIanaStdTzLinkZones,
		num1FieldLength,
		num1TotalLength,
		true,
		ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by TzStrFmt{}.RightJustifyNum()\n" +
			"Error='%v'\n", err.Error())
	}

	_, err = b.WriteString(outputStr + "\n")

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by b.WriteString(tzStats.TotalIanaStdTzLinkZones)\n" +
			"Error='%v'\n", err.Error())
	}

	fieldStr = "Military Time Zones"

	outputStr, err = TzStrFmt{}.LeftJustifyField(fieldStr, labelLength, true, ePrefix )

	if err != nil{
		return err
	}

	_, err = b.WriteString(tzLog.leftMarginStr + outputStr + "\n")

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by b.WriteString(Military Time Zones)\n" +
			"Error='%v'\n", err.Error())
	}



	//tzStats.NumMilitaryTZones




	outputStr += leftMarginStr +
		"Total Iana Time Zones" +
		fmt.Sprintf("%4d\n\n", tzStats.TotalIanaStdTzLinkZones)

	outputStr += leftMarginStr +
		"NumMilitaryTZones: " +
		fmt.Sprintf("%4d\n", tzStats.NumMilitaryTZones)


	outputStr += leftMarginStr +
		"NumOtherTZones: " +
		fmt.Sprintf("%4d\n", tzStats.NumOtherTZones)

}
