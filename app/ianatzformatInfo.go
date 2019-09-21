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
	"local.com/amarillomike/ianatzformatInfo/inprocess"
	"local.com/amarillomike/ianatzformatInfo/outprocess"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)



func main() {
	ePrefix := "ianaTzFormatInfo.main() "

	tzdatastructs.DEBUG = 0

	err := inprocess.SetCurrentWorkingDirectory()

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	inputFileMgr, err := inprocess.CreateInputFileMgr(tzdatastructs.InputFileName)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	baseDir, err := inprocess.GetTargetDirectory(inputFileMgr)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	fmt.Println("baseDir: ", baseDir)

	dirFileInfo, err := inprocess.GetDirectoryInfo(baseDir)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	tzdatastructs.TimeZoneGroups,
	tzdatastructs.TimeZones,
		err =
		 inprocess.ParseIanaTzData{}.ParseTzAndLinks(dirFileInfo)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	f, err := outprocess.TzOutProcess{}.CreateOpenOutputFile(
		inputFileMgr.GetDirMgr(), tzdatastructs.OutputFileName)

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
