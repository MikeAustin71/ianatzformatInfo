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