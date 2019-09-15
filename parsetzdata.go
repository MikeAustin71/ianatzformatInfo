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
	"strings"
)

var TimeZoneGroups []TimeZoneGroupDto
var TimeZoneData []TimeZoneDataDto
var SubTimeZoneData []TimeZoneDataDto
var AliasTimeZoneData []TimeZoneDataDto




// [linked Zone] primary zone
var mapTzLinks map[string]string

var DEBUG = 0


const inputFileName = "targettzdata.txt"

const outputfilename = "timezonedata.go"

const homeDir = "D:\\gowork\\src\\MikeAustin71\\ianatzoneinfo\\parsetzdata"

var curWorkingDirectory pathfileops.DirMgr

func main() {
	ePrefix := "parsetzdata.main() "

	mapTzLinks = make(map[string]string, 0)

	err := setCurrentWorkingDirectory()

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	inputFileMgr, err := createInputFileMgr(inputFileName)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	baseDir, err := getTargetDirectory(inputFileMgr)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	fmt.Println("baseDir: ", baseDir)

	dirFileInfo, err := getDirectoryInfo(baseDir)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	TimeZoneGroups,
		TimeZoneData,
		SubTimeZoneData,
		AliasTimeZoneData,
		err =
		ParseIanaTzData{}.ParseTzAndLinks(dirFileInfo)
	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	f, err := createOpenOutputFile(inputFileMgr.GetDirMgr(), outputfilename)

	if err != nil {
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = writeHeadersToOutputFile(f)

	if err != nil {
		_ = f.CloseThisFile()
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = writeTimeZoneArrayToOutputFile(f)

	if err != nil {
		_ = f.CloseThisFile()
		fmt.Printf(ePrefix+"%v\n", err.Error())
		return
	}

	err = writeLinkMapToOutputFile(f)

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

func createOpenOutputFile(
	outputPathDirMgr pathfileops.DirMgr,
	fileNameExt string) (f pathfileops.FileMgr, err error) {

	f = pathfileops.FileMgr{}
	err = nil
	var err2 error

	ePrefix := "converttimezones.createOutputFile() Error: "

	f, err2 = pathfileops.FileMgr{}.NewFromDirMgrFileNameExt(outputPathDirMgr, fileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	err = f.IsFileMgrValid(ePrefix)

	if err != nil {
		return f, err
	}

	fileExists, err2 := f.DoesThisFileExist()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	if fileExists {

		err2 = f.DeleteThisFile()

		if err2 != nil {
			err = fmt.Errorf(ePrefix+"%v", err2.Error())
			return f, err
		}

		fileExists, err2 = f.DoesThisFileExist()

		if err2 != nil {
			err = fmt.Errorf(ePrefix+"%v", err2.Error())
			return f, err
		}

		if fileExists {
			err = fmt.Errorf(ePrefix+"Existing Output File FAILED to Delete! "+
				"Output File= '%v' ", f.GetAbsolutePathFileName())
			return f, err
		}

	}

	err2 = f.CreateThisFile()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	err2 = f.OpenThisFileReadWrite()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	err = nil

	return f, err
}


// outputFileMgr MUST be open and ready for Write operations.
func writeHeadersToOutputFile(outputFileMgr pathfileops.FileMgr) (err error) {
	err = nil

	ePrefix := "parsetzdata.writeHeadersToOutputFile() Error: "

	if !outputFileMgr.IsInitialized() {
		err = fmt.Errorf(ePrefix +
			"Input parameter 'outputFileMgr' IS NOT INITIALIZED!")
	}

	if !outputFileMgr.IsFilePointerOpen() {
		err = fmt.Errorf(ePrefix +
			"'outputFileMagr IS NOT OPEN!")
	}

	_, err2 := outputFileMgr.WriteStrToFile("package main\n\n")

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"Line1: %v", err2.Error())
		err2 = outputFileMgr.CloseThisFile()

		if err2 != nil {
			err = pathfileops.FileHelper{}.ConsolidateErrors([]error{err, err2})
		}

		return err
	}

	_, err2 = outputFileMgr.WriteStrToFile("\n\n\n")

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"Line5: %v", err2.Error())
		err2 = outputFileMgr.CloseThisFile()

		if err2 != nil {
			err = pathfileops.FileHelper{}.ConsolidateErrors([]error{err, err2})
		}

		return err
	}

	err = nil
	return err
}

func writeLinkMapToOutputFile(outputFileMgr pathfileops.FileMgr) error {

	ePrefix := "parsetzdata.writeTimeZoneArrayToOutputFile() "


	sb := strings.Builder{}
	lenMapTzLinks := len(mapTzLinks)
	sb.Grow( lenMapTzLinks * 40)

	sb.WriteString("// mapTzLinks - A listing of deprecated time zones with links to active \n")
	sb.WriteString("// IANA time zones. key='deprecated time zone' value='current active time zone'\n")
	sb.WriteString(fmt.Sprintf("// The number of links is: %v\n", lenMapTzLinks))
	// var linkMap = map[string]string{
	//  "America/Buenos_Aires":             "America/Argentina/Buenos_Aires",

	sb.WriteString("var linkMap = map[string]string {\n")

	for key, value := range mapTzLinks {

		sb.WriteString("   \"" + key + "\":         \"" + value +"\", \n")

	}

	sb.WriteString("    }\n\n\n")

	_, err := outputFileMgr.WriteStrToFile(sb.String())

	if err != nil {
		return fmt.Errorf(ePrefix + "%v", err.Error())
	}

	return nil
}

func writeTimeZoneArrayToOutputFile(outputFileMgr pathfileops.FileMgr) error {

	ePrefix := "parsetzdata.writeTimeZoneArrayToOutputFile() "
	lenTzAry := len(timezoneArray)

	sb := strings.Builder{}
	sb.Grow(lenTzAry * 30)

	sb.WriteString("// timeZoneArray - This array contains time zones from the IANA database. \n")
	sb.WriteString(fmt.Sprintf("// The total number of time zones is %v\n", lenTzAry ))
	sb.WriteString("var timeZoneAry = []string {\n")

	for i:=0; i < lenTzAry; i++ {
		if i == lenTzAry - 1 {
			sb.WriteString("           \"" + timezoneArray[i] + "\"}\n\n\n")
		} else {
			sb.WriteString("           \"" + timezoneArray[i] + "\",\n")
		}

	}

	_, err := outputFileMgr.WriteStrToFile(sb.String())

	if err != nil {
		return fmt.Errorf(ePrefix + "%v", err.Error())
	}

	return nil
}