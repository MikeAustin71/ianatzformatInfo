package main


import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/inprocess"
	"local.com/amarillomike/ianatzformatInfo/outprocess"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"time"
)


// main - This application will process IANA source data files and create a series
// of types housed in a source file used by the 'Go' programming language. These
// types will facilitate the use of global time zones in date time operations
// incorporated into 'Go' programs.
//
// In order to function properly, this application expects to find the file
// 'targettzdata.txt' in a sub-directory labeled 'input' which is located directly
// beneath the directory which houses this application executable.
// Example:
// Applzication executable is located in 'D:\myAppDir' .
// File 'targettzdata.txt' MUST BE located in directory: 'D:\myAppDir\input\targettzdata.txt'.
//
// 'targettzdata.txt' is a text file containing two lines of information on the first two
// lines of the text file. Each line must be terminated with a new line character '\n'.
//
// Line 1: The first line designates the 'path' to the IANA time zone
//         data files.
//
// Line 2: The second line designates the 'path' to the output file.
//
// Configure these two lines in accordance with the following example.
//
// Example:
// "InputDirectory: D:\T11\data\2019c\n"
// "OutputDirectory: D:\GoProjects\ianatzformatInfo\app\output\n"
//
// The leading field names, 'InputDirectory:' and 'OutputDirectory:'
// are mandatory.
//
func main() {

	ePrefix := "ianaTzFormatInfo.main() "

	tzdatastructs.ApplicationStartDateTime = time.Now()

	// executableWorkingDirMgr, err := fileops.FileOps{}.GetApplicationDirectory(ePrefix)

	executableWorkingDirMgr, err := pathfileops.DirMgr{}.New("D:\\GoProjects\\ianatzformatInfo\\app")

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}


	fmt.Println()
	fmt.Println("ianatzformatInfo.exe" )
	fmt.Println("--------------------")
	fmt.Printf("Directory Directory:\n     %v\n\n", executableWorkingDirMgr.GetAbsolutePath())
	fmt.Println()

	var baseDataInputDirMgr pathfileops.DirMgr

	baseDataInputDirMgr, err = pathfileops.DirMgr{}.New(
		executableWorkingDirMgr.GetAbsolutePathWithSeparator() + "input")


	var zoneInfoDataDto tzdatastructs.ZoneInfoDataDto


	// tzdatastructs.AppInputParametersFileName = "targettzdata.txt"
	zoneInfoDataDto, err =
		tzdatastructs.ZoneInfoDataDto{}.AcquireZoneInfo(
			baseDataInputDirMgr,
			tzdatastructs.AppInputParametersFileName,
			ePrefix)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	 fmt.Printf("Base Data Input File:\n     %v\n\n", zoneInfoDataDto.AppInputPathFileNameExt)

	tzLog := outprocess.TzLogOps{}

	err = tzLog.InitializeLogOps(&zoneInfoDataDto, ePrefix)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}


	var timeZoneStats tzdatastructs.TimeZoneStatsDto
parser := inprocess.ParseZoneInfoData{}

	timeZoneStats,
		err =
		 parser.ParseZoneInfo(zoneInfoDataDto, ePrefix)

	if err != nil {
		_ = tzLog.WriteError(err, ePrefix)
		_ = tzLog.WriteFooter(&timeZoneStats, ePrefix)
		fmt.Printf("%v\n", err.Error())
		return
	}

	err = outprocess.TzOutProcess{}.WriteOutput(
		zoneInfoDataDto.AppOutputDirMgr,
		tzdatastructs.TimeZoneDataOutputFileName,
		&timeZoneStats,
		ePrefix)


	if err != nil {
		_ = tzLog.WriteError(err, ePrefix)
		_ = tzLog.WriteFooter(&timeZoneStats, ePrefix)
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	fmt.Printf("Output Source File:\n     %v\n\n",
		zoneInfoDataDto.AppOutputDirMgr.GetAbsolutePathWithSeparator() + tzdatastructs.TimeZoneDataOutputFileName)

	fmt.Printf("Number Of Captured Iana Time Zones: %v\n\n",
		timeZoneStats.IanaCapturedTimeZones.GetNumberOfTimeZones())

	err = tzLog.TestCapturedIanaTimeZones(&timeZoneStats, ePrefix)

	if err != nil {
		_ = tzLog.WriteError(err, ePrefix)
		_ = tzLog.WriteFooter(&timeZoneStats, ePrefix)
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = tzLog.WriteTimeZones(&timeZoneStats, ePrefix)

	if err != nil {
		_ = tzLog.WriteError(err, ePrefix)
		_ = tzLog.WriteFooter(&timeZoneStats, ePrefix)
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = tzLog.WriteIanaRegionalTotals(&timeZoneStats, ePrefix)

	if err != nil {
		_ = tzLog.WriteError(err, ePrefix)
		_ = tzLog.WriteFooter(&timeZoneStats, ePrefix)
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = tzLog.WriteSummaryTotals(&timeZoneStats, ePrefix)

	if err != nil {
		_ = tzLog.WriteError(err, ePrefix)
		_ = tzLog.WriteFooter(&timeZoneStats, ePrefix)
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = tzLog.WriteFooter(&timeZoneStats, ePrefix)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}


	fmt.Println("---------------------")
	fmt.Println("Successful Completion")
	fmt.Println("---------------------")
	fmt.Println()

}

