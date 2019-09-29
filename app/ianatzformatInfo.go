package main

/*
- strops -
Use this command to down load and install the *pathfileops* package
locally.

    go get github.com/MikeAustin71/stringopsgo/strops/v2

To update the package run:

    go get -u github.com/MikeAustin71/stringopsgo/strops/v2


- pathfileops -
Use this command to down load and install the *pathfileops* package
locally. Note: Version 2+ supports *Go* modules.

    go get github.com/MikeAustin71/pathfileopsgo/pathfileops/v2

To update the package run:

    go get -u github.com/MikeAustin71/pathfileopsgo/pathfileops/v2

*/

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/inprocess"
	"local.com/amarillomike/ianatzformatInfo/outprocess"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)



func main() {

	ePrefix := "ianaTzFormatInfo.main() "

	tzdatastructs.DEBUG = 0

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

	tzdatastructs.TimeZoneGroups,
	tzdatastructs.TimeZones,
		err =
		 inprocess.ParseIanaTzData{}.ParseTzAndLinks(dirFileInfo, ePrefix)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	var f pathfileops.FileMgr

	f, err = outprocess.TzOutProcess{}.CreateOpenOutputFile(
		outputFileDirMgr,
		tzdatastructs.OutputFileName)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = outprocess.TzOutProcess{}.WriteHeadersToOutputFile(f)

	if err != nil {
		_ = f.CloseThisFile()
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = outprocess.TzOutProcess{}.WriteTimeZoneArrayToOutputFile(f)

	if err != nil {
		_ = f.CloseThisFile()
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = outprocess.TzOutProcess{}.WriteLinkMapToOutputFile(f)

	if err != nil {
		_ = f.CloseThisFile()
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = f.CloseThisFile()

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

}
