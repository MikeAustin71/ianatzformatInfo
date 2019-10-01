package main

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/inprocess"
	"local.com/amarillomike/ianatzformatInfo/outprocess"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"os"
)


func main() {

	ePrefix := "ianaTzFormatInfo.main() "

	// /////////////////////////////////////////
	//  IMPORTANT!!! SET CORRECT VALUE!!!    //
	// ////////////////////////////////////////
	tzdatastructs.DEBUG = 1

	currWorkingDirMgr, err := inprocess.AcquireTzData{}.SetCurrentWorkingDirectory(ePrefix)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	targetInputDir :=
		currWorkingDirMgr.GetAbsolutePath() + string(os.PathSeparator) + "input"

	targetParameterPathFileName :=
		pathfileops.FileHelper{}.JoinPathsAdjustSeparators(
			targetInputDir,
			tzdatastructs.AppInputParametersFileName)

	var dirFileInfo pathfileops.FileMgrCollection
	var outputFileDirMgr pathfileops.DirMgr

	dirFileInfo, outputFileDirMgr, err =
		inprocess.AcquireTzData{}.AcquireDirectoryInfo(targetParameterPathFileName, ePrefix)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	var timeZoneGroups []tzdatastructs.TimeZoneGroupCollection
	var timeZones []tzdatastructs.TimeZoneDataCollection

parser := inprocess.ParseIanaTzData{}

	timeZoneGroups,
	timeZones,
		err =
		 parser.ParseTzAndLinks(dirFileInfo, ePrefix)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	err = outprocess.TzOutProcess{}.WriteOutput(
		outputFileDirMgr,
		tzdatastructs.OutputFileName,
		timeZoneGroups,
		timeZones,
		ePrefix)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

}

/*
$ go run ianatzformatInfo.go
Number of Target Files:  33
Valid File:  africa
Valid File:  antarctica
Valid File:  asia
Valid File:  australasia
Valid File:  backward
Valid File:  backzone
Valid File:  etcetera
Valid File:  europe
Valid File:  northamerica
Valid File:  pacificnew
Valid File:  southamerica
ianaTzFormatInfo.main() ianaTzFormatInfo.main() TzOutProcess.WriteOutput() TzOutProcess.writeLevelOneTimeZones()

Time Zone Collection is EMPTY!
Parent Group='America'
Group Name='Argentina'
 */