package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"strings"
)

// [linked Zone] primary zone
var mapTzLinks map[string]string

var timezoneArray = make([]string, 100, 100)

type TzOutProcess struct {
	input string
}


func (tzOut TzOutProcess) CreateOpenOutputFile(
	outputPathDirMgr pathfileops.DirMgr,
	fileNameExt string) (f pathfileops.FileMgr, err error) {

	ePrefix := "TzOutProcess.CreateOutputFile() Error: "

	f = pathfileops.FileMgr{}
	err = nil
	var err2 error


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
func (tzOut TzOutProcess) WriteHeadersToOutputFile(
	outputFileMgr pathfileops.FileMgr) (err error) {

		err = nil

	ePrefix := "TzOutProcess.WriteHeadersToOutputFile() Error: "

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

func (tzOut TzOutProcess) WriteLinkMapToOutputFile(
	outputFileMgr pathfileops.FileMgr) error {

	ePrefix := "TzOutProcess.WriteTimeZoneArrayToOutputFile() "


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

func (tzOut TzOutProcess) WriteTimeZoneArrayToOutputFile(
	outputFileMgr pathfileops.FileMgr) error {

	ePrefix := "TzOutProcess.WriteTimeZoneArrayToOutputFile() "
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
