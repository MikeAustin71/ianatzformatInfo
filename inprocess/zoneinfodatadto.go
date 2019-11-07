package inprocess

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/fileops"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
)


type ZoneInfoDataDto struct {
	ZoneInfoInputDir        string
	AppInputPathFileNameExt string
	AppOutputDir            string
	AppLogPathFileNameExt   string
	IanaTimeZoneVersion     string
	ZoneInfoDirTreeInfo     pathfileops.DirectoryTreeInfo
	ZoneInfoDirMgr          pathfileops.DirMgr
	AppOutputDirMgr         pathfileops.DirMgr
	AppLogFileMgr           pathfileops.FileMgr
}

// AcquireZoneInfo - Reads, parses and returns all information
// contained in the base data input file, 'targettzdata.txt'.
// This base data input file must reside in a directory named
//'input' which is located immediately underneath the directory
// containing this application executable.
func (zInDto ZoneInfoDataDto) AcquireZoneInfo(
	baseDataDirMgr pathfileops.DirMgr,
	baseDataFileNameExt,
	ePrefix string) (ZoneInfoDataDto, error) {
		
		ePrefix += "ZoneInfoDataDto.AcquireZoneInfo() "

		zoneInfoDto := ZoneInfoDataDto{}
		
	baseDataInputFileMgr, err2 :=
			fileops.FileOps{}.CreateOpenFile(baseDataDirMgr, baseDataFileNameExt, ePrefix)
			//zInDto.createInputFileMgr(baseDataFileNameExt, ePrefix)
	
	if err2 != nil {
		return zoneInfoDto, err2
	}

	zoneInfoDto.AppInputPathFileNameExt = baseDataInputFileMgr.GetAbsolutePathFileName()

	zoneInfoDto.ZoneInfoInputDir,
	zoneInfoDto.AppOutputDir,
	zoneInfoDto.IanaTimeZoneVersion,
	err2 =
			zInDto.readBaseDataInput(baseDataInputFileMgr, ePrefix)
	
	if err2 != nil {
		return ZoneInfoDataDto{}, err2
	}

	zoneInfoDto.ZoneInfoDirTreeInfo, err2 =
		zInDto.getZoneInfoDirFileInfo(zoneInfoDto.ZoneInfoInputDir, ePrefix)

	if err2 != nil {
		return ZoneInfoDataDto{}, err2
	}

	zoneInfoDto.AppOutputDirMgr, err2 = pathfileops.DirMgr{}.New(zoneInfoDto.AppOutputDir)
	
	if err2 != nil {
		return ZoneInfoDataDto{}, err2
	}

	zoneInfoDto.ZoneInfoDirMgr, err2 = pathfileops.DirMgr{}.New(zoneInfoDto.ZoneInfoInputDir)

	if err2 != nil {
		return ZoneInfoDataDto{}, err2
	}

	zoneInfoDto.AppLogFileMgr, err2 =
		zoneInfoDto.createOpenLogOutputFile(
			zoneInfoDto.AppOutputDirMgr,
			ePrefix)

	if err2 != nil {
		return ZoneInfoDataDto{}, err2
	}

	zoneInfoDto.AppLogPathFileNameExt =
		zoneInfoDto.AppLogFileMgr.GetAbsolutePathFileName()

	return zoneInfoDto, nil
}

// createInputFileMgr - Creates a File Manager to base data input file,
// 'targettzdata.txt' which must reside in a directory named 'input'
// which is located immediately underneath the directory containing
// this application executable.
//
func (zInDto ZoneInfoDataDto) createInputFileMgr(
	baseDataInputPathFileName, ePrefix string) (baseDataInputFileMgr pathfileops.FileMgr, err error) {
	
	ePrefix += "ZoneInfoDataDto.createInputFileMgr() "
	err = nil
	baseDataInputFileMgr = pathfileops.FileMgr{}
	
	var err2 error
	
	baseDataInputFileMgr, err2 = pathfileops.FileMgr{}.New(baseDataInputPathFileName)

	if err2 != nil {

		err = fmt.Errorf(ePrefix+"\n%v\n", err2.Error())

		return baseDataInputFileMgr, err
	}

	var fileDoesExist bool

	fileDoesExist, err2 = baseDataInputFileMgr.DoesThisFileExist()

	if err2 != nil {

		err = fmt.Errorf(ePrefix+"\n%v\n", err2.Error())

		return baseDataInputFileMgr, err
	}

	if !fileDoesExist {
		err = fmt.Errorf(ePrefix+
			"\nBase Data Input File DOES NOT EXIST!\n" +
			"baseDataInputPathFileName='%v'\n" +
			"baseDataInputFileMgr='%v'\n",
			baseDataInputPathFileName, baseDataInputFileMgr.GetAbsolutePathFileName())

		return baseDataInputFileMgr, err
	}

	err = nil

	return baseDataInputFileMgr, err
}

// createOpenLogOutputFile - Generates the log path and
// file name then creates and opens the file.
func (zInDto ZoneInfoDataDto) createOpenLogOutputFile(
	outputPathDirMgr pathfileops.DirMgr,
	ePrefix string) (pathfileops.FileMgr, error) {

	ePrefix += "ZoneInfoDataDto.createOpenLogOutputFile() "

	fmtDateTimeSecondStr := "20060102150405"
	currDateTimeStr := tzdatastructs.ApplicationStartDateTime.Format(fmtDateTimeSecondStr)

	fileNameExt :=   currDateTimeStr +"_ianaformatInfoLog" +".txt"

	return fileops.FileOps{}.CreateOpenFile(outputPathDirMgr, fileNameExt, ePrefix)

}

// getZoneInfoDirFileInfo - Receives the Zone Info directory path as input and returns
// a collection of file managers for every file in the Zone Info directory tree.
//
func (zInDto ZoneInfoDataDto) getZoneInfoDirFileInfo(
	zoneInfoInputDir,
	ePrefix string) (zoneInfoDirFileInfo pathfileops.DirectoryTreeInfo, err error) {

	ePrefix += "ZoneInfoDataDto.getZoneInfoDirFileInfo() "
	zoneInfoDirFileInfo = pathfileops.DirectoryTreeInfo{}
	err = nil


	if len(zoneInfoInputDir) == 0 {
		err = errors.New(ePrefix + "\n" +
			"Input Parameter 'zoneInfoInputDir' is an EMPTY string!\n")
		return zoneInfoDirFileInfo, err
	}

	zoneInfoInputDir = strings.ToLower(zoneInfoInputDir)

	zoneInfoInputDirMgr, err2 := pathfileops.DirMgr{}.New(zoneInfoInputDir)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
		return zoneInfoDirFileInfo, err
	}

	var baseDMgrDoesExist bool

	baseDMgrDoesExist, err2 = zoneInfoInputDirMgr.DoesThisDirectoryExist()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n'zoneInfoInputDirMgr' Non-Path Error.\n" +
			"%v\n", err2.Error())
		return zoneInfoDirFileInfo, err
	}

	if !baseDMgrDoesExist {
		err = fmt.Errorf(ePrefix+"'zoneInfoInputDirMgr' Path DOES NOT EXIST!\nzoneInfoInputDirMgr= %v\n",
			zoneInfoInputDirMgr.GetAbsolutePath())
		return zoneInfoDirFileInfo, err
	}

	fileSelectCriteria := pathfileops.FileSelectionCriteria{}

	var errs []error

	zoneInfoDirFileInfo, errs = zoneInfoInputDirMgr.FindDirectoryTreeFiles(fileSelectCriteria)

	err2 = pathfileops.FileHelper{}.ConsolidateErrors(errs)

	if err2 != nil {

		err =fmt.Errorf(ePrefix +
			"\nError returned by baseDirMgr.FindDirectoryTreeFiles(fileSelectCriteria)\n" +
			"Error='%v'\n", err2.Error())

		return zoneInfoDirFileInfo, err
	}


	if zoneInfoDirFileInfo.FoundFiles.GetNumOfFileMgrs() < 1 {
		err = fmt.Errorf(ePrefix + "Error: No files located in target 'baseDirMgr'.\n" +
			"'baseDirMgr='%v'\n", zoneInfoInputDirMgr.GetAbsolutePath())
	}

	err = nil
	return zoneInfoDirFileInfo, err
}

// readBaseDataInput - Reads base data input file 'targettzdata.txt' to
// identify and return time zone input directory, output directory for
// application products and the Iana Time Zone Version.
//
// Example Input File: 'targettzdata.txt'
//
//   InputDirectory: D:\tz2\zoneinfo
//   OutputDirectory: D:\GoProjects\ianatzformatInfo\app\output
//   Iana Time Zone Version: 2019c
//
func (zInDto ZoneInfoDataDto) readBaseDataInput(
	baseDataInputFileMgr pathfileops.FileMgr,
	ePrefix string) (inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion string, err error) {

	inputZoneInfoDir = ""
	appOutputDir = ""
	ianaTimeZoneVersion = ""
	err = nil

	ePrefix += "ZoneInfoDataDto.readBaseDataInput() "

	err2 := baseDataInputFileMgr.OpenThisFileReadWrite()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"\n%v\n", err2.Error())
		return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
	}

	var bArray []byte
	var errArray []error

	bArray, err2 = baseDataInputFileMgr.ReadAllFile()

	if err2 != nil {

		err3 := fmt.Errorf(ePrefix+
			"\n" +
			"targettzdata.txt Path: %v\n" +
			"Error: %v\n",
			baseDataInputFileMgr.GetAbsolutePathFileName(),
			err2.Error())

		errArray = append(errArray, err3)

		err2 = baseDataInputFileMgr.CloseThisFile()

		if err2 != nil {
			err3 := fmt.Errorf("\n" + ePrefix +
				"\n%v", err2)
			errArray = append(errArray, err3)
		}

		if len(errArray) > 0 {
			err = pathfileops.FileHelper{}.ConsolidateErrors(errArray)
		}

		return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
	}

	err2 = baseDataInputFileMgr.CloseThisFile()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error closing 'baseDataInputFileMgr' (targettzdata.txt Path).\n" +
			"baseDataInputFileMgr='%v'\n" +
			"Error='%v'\n", baseDataInputFileMgr.GetAbsolutePathFileName(), err2.Error())

		return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
	}

	lBArray := len(bArray)

	if lBArray < 3 {
		err = fmt.Errorf(ePrefix +
			"Error: Read only %v bytes from file 'baseDataInputFileMgr' (targettzdata.txt Path)\n" +
			"baseDataInputFileMgr='%v'\n",
			lBArray,
			baseDataInputFileMgr.GetAbsolutePathFileName())

		return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
	}

	startIndex := 0
	readDir := ""

	// Read InputDirectory Field
	readDir, startIndex = strops.StrOps{}.ReadStringFromBytes(bArray, startIndex)

	inputFieldName := "InputDirectory:"

	idx := strings.Index(readDir, inputFieldName)

	if idx == -1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Directory Field Name '%v' was not found in first\n" +
			"string read from 'baseDataInputFileMgr' (targettzdata.txt).\n" +
			"baseDataInputFileMgr='%v'\n", inputFieldName, baseDataInputFileMgr.GetAbsolutePathFileName())

		return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
	}

	idx += len(inputFieldName)
	inputZoneInfoDir = readDir[idx:]
	inputZoneInfoDir = strings.TrimLeft(inputZoneInfoDir, " ")
	inputZoneInfoDir = strings.TrimRight(inputZoneInfoDir, " ")
	inputZoneInfoDir = strings.ToLower(inputZoneInfoDir)

	doesExist := pathfileops.FileHelper{}.DoesFileExist(inputZoneInfoDir)

	if !doesExist {
		err = fmt.Errorf(ePrefix +
			"\nTarget Zone Input Directory DOES NOT EXIST!\n" +
			"inputZoneInfoDir='%v'\n" +
			"baseDataInputFileMgr (targettzdata.txt)='%v'\n",
			inputZoneInfoDir, baseDataInputFileMgr.GetAbsolutePathFileName())

		return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
	}

	// Read OutputDirectory Field
	readDir, startIndex = strops.StrOps{}.ReadStringFromBytes(bArray, startIndex)

	outputFieldName := "OutputDirectory:"

	idx = strings.Index(readDir, outputFieldName)

	if idx == -1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Application Output Directory Field Name '%v' " +
			"was not found in second\n" +
			"string read from 'baseDataInputFileMgr' (targettzdata.txt).\n" +
			"baseDataInputFileMgr='%v'\n",
			inputFieldName, baseDataInputFileMgr.GetAbsolutePathFileName())

		return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
	}

	idx += len(outputFieldName)
	appOutputDir = readDir[idx:]
	appOutputDir = strings.TrimLeft(appOutputDir, " ")
	appOutputDir = strings.TrimRight(appOutputDir, " ")

	doesExist = pathfileops.FileHelper{}.DoesFileExist(appOutputDir)

	if !doesExist {
		err = fmt.Errorf(ePrefix +
			"\nApplication Output Directory DOES NOT EXIST!\n" +
			"appOutputDir='%v'\n" +
			"baseDataInputFileMgr (targettzdata.txt)='%v'",
			appOutputDir, baseDataInputFileMgr.GetAbsolutePathFileName())

		return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
	}

	// Read Iana Time Zone Version Field
	readDir, _ = strops.StrOps{}.ReadStringFromBytes(bArray, startIndex)

	ianaVersionFieldName := "Iana Time Zone Version:"

	idx = strings.Index(readDir, ianaVersionFieldName)

	if idx == -1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Iana Time Zone Version Field Name '%v' was not found in the third\n" +
			"string read from 'baseDataInputFileMgr'.\n" +
			"baseDataInputFileMgr () ='%v'\n",
			ianaVersionFieldName, baseDataInputFileMgr.GetAbsolutePathFileName())

		return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
	}

	idx += len(ianaVersionFieldName)
	ianaTimeZoneVersion = readDir[idx:]
	ianaTimeZoneVersion = strings.TrimLeft(ianaTimeZoneVersion, " ")
	ianaTimeZoneVersion = strings.TrimRight(ianaTimeZoneVersion, " ")


	return inputZoneInfoDir, appOutputDir, ianaTimeZoneVersion, err
}

