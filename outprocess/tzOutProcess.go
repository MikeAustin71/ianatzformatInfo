package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
)

// [linked Zone] primary zone
var mapTzLinks map[string]string

var timezoneArray = make([]string, 100, 100)

type TzOutProcess struct {
	input string
}

func (tzOut TzOutProcess) WriteOutput(
	outputPathDirMgr pathfileops.DirMgr,
	fileNameExt string,
	tzGroupsAry [] tzdatastructs.TimeZoneGroupCollection, // Array of Time Zone Group Collections
	tzZonesAry [] tzdatastructs.TimeZoneDataCollection,  // Array of Time Zone Data Collections
	ePrefix string) error {

	ePrefix += "TzOutProcess.WriteOutput() "

	f, err := tzOut.createOpenOutputFile(outputPathDirMgr, fileNameExt, ePrefix)

	if err != nil {
		return err
	}

	err = tzOut.writeHeadersToOutputFile(f, ePrefix)

	if err != nil {
		return err
	}


	return nil
}

func (tzOut TzOutProcess) createOpenOutputFile(
	outputPathDirMgr pathfileops.DirMgr,
	fileNameExt, ePrefix string) (f pathfileops.FileMgr, err error) {

	ePrefix += "TzOutProcess.CreateOutputFile() "

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

// writeHeadersToOutputFile - Writes header information to the
// output file. This includes the 'package' statement.
//
// Input parameter 'outputFileMgr' MUST be open and ready for
// Write operations.
//
func (tzOut TzOutProcess) writeHeadersToOutputFile(
	outputFileMgr pathfileops.FileMgr, ePrefix string) (err error) {

		err = nil

	ePrefix += "TzOutProcess.writeHeadersToOutputFile() "

	if !outputFileMgr.IsInitialized() {
		err = fmt.Errorf(ePrefix +
			"Input parameter 'outputFileMgr' IS NOT INITIALIZED!")
	}

	if !outputFileMgr.IsFilePointerOpen() {
		err = fmt.Errorf(ePrefix +
			"'outputFileMagr IS NOT OPEN!")
	}

	var errorArray []error

	_, err2 := outputFileMgr.WriteBytesToFile ([]byte("package main\n\n\n\n\n"))

	if err2 != nil {

		errorArray = append(errorArray, fmt.Errorf(ePrefix+"Line1: %v", err2.Error()))

		err2 = outputFileMgr.CloseThisFile()

		if err2 != nil {
			errorArray = append(errorArray, err2)
			err = pathfileops.FileHelper{}.ConsolidateErrors(errorArray)
		}

		return err
	}


	err = nil
	return err
}

func (tzOut TzOutProcess) writeLevelOneTimeZones(
	tzGroupsAry [] tzdatastructs.TimeZoneGroupCollection, // Array of Time Zone Group Collections
	tzZonesAry [] tzdatastructs.TimeZoneDataCollection,  // Array of Time Zone Data Collections)
	ePrefix string) error {

	ePrefix += "TzOutProcess.writeLevelOneTimeZones() "

	//tzGroupsAry[tzdatastructs.Level_01_Idx].Sort

	return nil
}

func (tzOut TzOutProcess) WriteLinkMapToOutputFile(
	outputFileMgr pathfileops.FileMgr, ePrefix string) error {

	ePrefix += "TzOutProcess.WriteTimeZoneArrayToOutputFile() "


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
	outputFileMgr pathfileops.FileMgr, ePrefix string) error {

	ePrefix += "TzOutProcess.WriteTimeZoneArrayToOutputFile() "
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
