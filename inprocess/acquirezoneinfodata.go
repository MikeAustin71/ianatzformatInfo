package inprocess

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"strings"
)

type ZoneInfoDataDto struct {
	ZoneInfoInputDir    string
	AppOutputDir        string
	ZoneInfoDirFileInfo pathfileops.FileMgrCollection
	AppOutputDirMgr     pathfileops.DirMgr
}

func (zInDto ZoneInfoDataDto) AcquireZoneInfo(
	baseDataInputPathFileName,
	ePrefix string) (ZoneInfoDataDto, error) {
		
		ePrefix += "ZoneInfoDataDto.AcquireZoneInfo() "

		zoneInfoDto := ZoneInfoDataDto{}
		
	baseDataInputFileMgr, err2 := 
			zInDto.createInputFileMgr(baseDataInputPathFileName, ePrefix)
	
	if err2 != nil {
		return zoneInfoDto, err2
	}

	zoneInfoDto.ZoneInfoInputDir,
	zoneInfoDto.AppOutputDir, 
	err2 =
			zInDto.getTargetDirectories(baseDataInputFileMgr, ePrefix)
	
	if err2 != nil {
		return ZoneInfoDataDto{}, err2
	}

	zoneInfoDto.ZoneInfoDirFileInfo, err2 =
		zInDto.getZoneInfoDirFileInfo(zoneInfoDto.ZoneInfoInputDir, ePrefix)

	if err2 != nil {
		return ZoneInfoDataDto{}, err2
	}

	zoneInfoDto.AppOutputDirMgr, err2 = pathfileops.DirMgr{}.New(zoneInfoDto.AppOutputDir)
	
	if err2 != nil {
		return ZoneInfoDataDto{}, err2
	}
	
	return zoneInfoDto, nil
}

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

func (zInDto ZoneInfoDataDto) getZoneInfoDirFileInfo(
	zoneInfoInputDir, ePrefix string) (zoneInfoDirFileInfo pathfileops.FileMgrCollection, err error) {

	ePrefix += "ZoneInfoDataDto.getZoneInfoDirFileInfo() "
	zoneInfoDirFileInfo = pathfileops.FileMgrCollection{}
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

	zoneInfoDirFileInfo, err2 = zoneInfoInputDirMgr.FindFilesByNamePattern("*")

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
		return zoneInfoDirFileInfo, err
	}

	if zoneInfoDirFileInfo.GetNumOfFileMgrs() < 1 {
		err = fmt.Errorf(ePrefix + "Error: No files located in target 'zoneInfoInputDirMgr'.\n" +
			"'zoneInfoInputDirMgr='%v'\n", zoneInfoInputDirMgr.GetAbsolutePath())
	}

	err = nil
	return zoneInfoDirFileInfo, err
}

func (zInDto ZoneInfoDataDto) getTargetDirectories(
	baseDataInputFileMgr pathfileops.FileMgr,
	ePrefix string) (inputZoneInfoDir, appOutputDir string, err error) {
	
	inputZoneInfoDir = ""
	appOutputDir = ""
	err = nil

	ePrefix += "ZoneInfoDataDto.getTargetDirectories() "


	err2 := baseDataInputFileMgr.OpenThisFileReadWrite()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"\n%v\n", err2.Error())
		return inputZoneInfoDir, appOutputDir, err
	}

	var bArray []byte
	var errArray []error

	bArray, err2 = baseDataInputFileMgr.ReadAllFile()

	if err2 != nil {

		err3 := fmt.Errorf(ePrefix+
			"\n" +
			"baseDataInputFileMgr: %v\n" +
			"Error: %v\n",
			baseDataInputFileMgr.GetAbsolutePath(),
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

		return inputZoneInfoDir, appOutputDir, err
	}

	err2 = baseDataInputFileMgr.CloseThisFile()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error closing 'baseDataInputFileMgr'.\n" +
			"baseDataInputFileMgr='%v'\n" +
			"Error='%v'\n", baseDataInputFileMgr.GetAbsolutePath(), err2.Error())
		return inputZoneInfoDir, appOutputDir, err
	}

	lBArray := len(bArray)

	if lBArray < 3 {
		err = fmt.Errorf(ePrefix +
			"Error: Read only %v bytes from file 'baseDataInputFileMgr'\n" +
			"baseDataInputFileMgr='%v'\n",
			lBArray,
			baseDataInputFileMgr)
		
		return inputZoneInfoDir, appOutputDir, err
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
			"string read from 'baseDataInputFileMgr'.\n" +
			"baseDataInputFileMgr='%v'\n", inputFieldName, baseDataInputFileMgr.GetAbsolutePath())
		return inputZoneInfoDir, appOutputDir, err
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
			"baseDataInputFileMgr='%v'\n", inputZoneInfoDir, baseDataInputFileMgr.GetAbsolutePath())

		return inputZoneInfoDir, appOutputDir, err
	}

	readDir, _ = strops.StrOps{}.ReadStringFromBytes(bArray, startIndex)

	outputFieldName := "OutputDirectory:"

	idx = strings.Index(readDir, outputFieldName)

	if idx == -1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Output Directory Field Name '%v' was not found in second\n" +
			"string read from 'baseDataInputFileMgr'.\n" +
			"baseDataInputFileMgr='%v'\n", inputFieldName, baseDataInputFileMgr.GetAbsolutePath())
		return inputZoneInfoDir, appOutputDir, err
	}

	idx += len(outputFieldName)
	appOutputDir = readDir[idx:]
	appOutputDir = strings.TrimLeft(appOutputDir, " ")
	appOutputDir = strings.TrimRight(appOutputDir, " ")

	doesExist = pathfileops.FileHelper{}.DoesFileExist(appOutputDir)

	if !doesExist {
		err = fmt.Errorf(ePrefix +
			"\nTarget Zone Output Directory DOES NOT EXIST!\n" +
			"appOutputDir='%v'\n" +
			"baseDataInputFileMgr='%v'",
			appOutputDir, baseDataInputFileMgr.GetAbsolutePath())

		return inputZoneInfoDir, appOutputDir, err
	}

	return inputZoneInfoDir, appOutputDir, err

} 