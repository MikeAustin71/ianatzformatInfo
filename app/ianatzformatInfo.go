package main

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/inprocess"
	"local.com/amarillomike/ianatzformatInfo/outprocess"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
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

	targetParameterPathFileName :=
		pathfileops.FileHelper{}.JoinPathsAdjustSeparators(
			currWorkingDirMgr.GetAbsolutePath(),
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

	timeZoneGroups,
	timeZones,
		err =
		 inprocess.ParseIanaTzData{}.ParseTzAndLinks(dirFileInfo, ePrefix)

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
