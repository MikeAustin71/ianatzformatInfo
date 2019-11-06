package inprocess

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
)

type AcquireTzData struct {
	input          string
	output         string
}

// AcquireDirectoryInfo - Returns directory and file information on the Time Zone
// source data directory as well as and 'outputDirMgr' describing the output directory
// where output time zone format data will be stored.
//
func (acTzDat AcquireTzData) AcquireDirectoryInfo(
	targetParameterPathFileName,
	ePrefix string) (dirFileInfo pathfileops.FileMgrCollection, outputDirMgr pathfileops.DirMgr, err error) {

	ePrefix += "AcquireTzData.AcquireDirectoryInfo() "
	dirFileInfo = pathfileops.FileMgrCollection{}.New()
	outputDirMgr = pathfileops.DirMgr{}
	inputPathName := ""
	outputPathName := ""

	inputFileMgr, err2 :=
		AcquireTzData{}.createInputFileMgr(targetParameterPathFileName, ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v", err2.Error())
		return dirFileInfo, outputDirMgr, err
	}

	inputPathName, outputPathName, err2 =
		AcquireTzData{}.getTargetDirectories(inputFileMgr, ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v", err2.Error())
		return dirFileInfo, outputDirMgr, err
	}
	fmt.Printf("Time Zone Data Input Directory:\n     %v\n\n",
		inputPathName)

	fmt.Printf("'timezonedata.go' Output Directory:\n     %v\n\n",
		outputPathName)

	dirFileInfo, err2 = AcquireTzData{}.getDirectoryInfo(inputPathName, ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v", err2.Error())
		return dirFileInfo, outputDirMgr, err
	}

	outputDirMgr, err2 = pathfileops.DirMgr{}.New(outputPathName)

	if err2 != nil {
		err = fmt.Errorf("%v", err2.Error())
		return dirFileInfo, outputDirMgr, err
	}

	return dirFileInfo, outputDirMgr, err
}
// createInputFileMgr - Creates the File Manager for the input data
// file.
func (acTzDat AcquireTzData) createInputFileMgr(
	targetPathFileName, ePrefix string) (inputFileMgr pathfileops.FileMgr, err error) {

	inputFileMgr = pathfileops.FileMgr{}
	err = nil

	ePrefix += "AcquireTzData.CreateInputFileMgr() "

	var err2 error

	// Example targetPathFileName
	// "D:\GoProjects\ianatzformatInfo\app\input\targettzdata.txt"

	inputFileMgr, err2 = pathfileops.FileMgr{}.New(targetPathFileName)

	if err2 != nil {

		err = fmt.Errorf(ePrefix+"\n%v\n", err2.Error())

		return inputFileMgr, err
	}

	var fileDoesExist bool

	fileDoesExist, err2 = inputFileMgr.DoesThisFileExist()

	if err2 != nil {

		err = fmt.Errorf(ePrefix+"\n%v\n", err2.Error())

		return inputFileMgr, err
	}

	if !fileDoesExist {
		err = fmt.Errorf(ePrefix+
			"\nTarget File DOES NOT EXIST!\n" +
			"targetPathFileName='%v'\n" +
			"inputFileMgr='%v'\n",
			targetPathFileName, inputFileMgr.GetAbsolutePathFileName())

		return inputFileMgr, err
	}

	err = nil

	return inputFileMgr, err
}

// GetDirectoryInfo - Walks the 'baseDir' and returns information on all directories and
// files found in the target input parameter 'baseDir'.
//
func (acTzDat AcquireTzData) getDirectoryInfo(
	baseDir, ePrefix string) (dirFileInfo pathfileops.FileMgrCollection, err error) {

	ePrefix += "AcquireTzData.GetDirectoryInfo() "
	dirFileInfo = pathfileops.FileMgrCollection{}.New()
	err = nil

	if len(baseDir) == 0 {
		err = errors.New(ePrefix + "\n" +
			"Input Parameter 'baseDir' is an EMPTY string!\n")
		return dirFileInfo, err
	}

	baseDir = strings.ToLower(baseDir)

	baseDirMgr, err2 := pathfileops.DirMgr{}.New(baseDir)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
		return dirFileInfo, err
	}

	var baseDMgrDoesExist bool

	baseDMgrDoesExist, err2 = baseDirMgr.DoesThisDirectoryExist()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n'baseDirMgr' Non-Path Error.\n" +
			"%v\n", err2.Error())
		return dirFileInfo, err
	}

	if !baseDMgrDoesExist {
		err = fmt.Errorf(ePrefix+"'baseDirMgr' Path DOES NOT EXIST!\nbaseDirMgr= %v\n",
			baseDirMgr.GetAbsolutePath())
		return dirFileInfo, err
	}

	var directTreeInfo  pathfileops.DirectoryTreeInfo

	fileSelectCriteria := pathfileops.FileSelectionCriteria{}

	//var errs = make([]error, 0)

	directTreeInfo, err2 = baseDirMgr.FindWalkDirFiles(fileSelectCriteria)

	if err2 != nil {
		err =fmt.Errorf(ePrefix +
			"\nError returned by baseDirMgr.FindWalkDirFiles(fileSelectCriteria)\n")

		return dirFileInfo, err
	}


	/*
	if len(errs) > 0 {
		err2 = pathfileops.FileHelper{}.ConsolidateErrors(errs)

		if err2 != nil {
			err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
		} else {
			err = fmt.Errorf(ePrefix +
				"\nError: FileHelper{}.ConsolidateErrors(errs) failed!")
		}
			return dirFileInfo, err
		}
*/


	dirFileInfo = directTreeInfo.FoundFiles

	if dirFileInfo.GetNumOfFileMgrs() < 1 {
		err = fmt.Errorf(ePrefix + "Error: No files located in target 'baseDirMgr'.\n" +
			"'baseDirMgr='%v'\n", baseDirMgr.GetAbsolutePath())
	}

	err = nil
	return dirFileInfo, err
}

// readBaseDataInput - Reads the input text file and extracts the input directory
// where IANA Time Zone data is located plus the output directory where formatted
// time zones will be created.
//
func (acTzDat AcquireTzData) getTargetDirectories(
	inputFileMgr pathfileops.FileMgr,
	ePrefix string) (inputDir, outputDir string, err error) {

	inputDir = ""
	outputDir = ""
	err = nil

	ePrefix += "AcquireTzData.GetTargetDirectories() "

	err2 := inputFileMgr.OpenThisFileReadWrite()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"\n%v\n", err2.Error())
		return inputDir, outputDir, err
	}

	var bArray []byte
	var errArray []error

	bArray, err2 = inputFileMgr.ReadAllFile()

	if err2 != nil {

		err3 := fmt.Errorf(ePrefix+
			"\n" +
			"inputFileMgr: %v\n" +
			"Error: %v\n",
			inputFileMgr.GetAbsolutePath(),
			err2.Error())

		errArray = append(errArray, err3)

		err2 = inputFileMgr.CloseThisFile()

		if err2 != nil {
			err3 := fmt.Errorf("\n" + ePrefix +
				"\n%v", err2)
			errArray = append(errArray, err3)
		}

		if len(errArray) > 0 {
			err = pathfileops.FileHelper{}.ConsolidateErrors(errArray)
		}

		return inputDir, outputDir, err
	}

	err2 = inputFileMgr.CloseThisFile()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error closing 'inputFileMgr'.\n" +
			"inputFileMgr='%v'\n" +
			"Error='%v'\n", inputFileMgr.GetAbsolutePath(), err2.Error())
		return inputDir, outputDir, err
	}

	lBArray := len(bArray)

	if lBArray < 3 {
		err = fmt.Errorf(ePrefix+"Read Zero bytes from file 'inputFileMgr'\n" +
			"inputFileMgr='%v'\n", inputFileMgr)
		return inputDir, outputDir, err
	}

	startIndex := 0
	readDir := ""

	readDir, startIndex = strops.StrOps{}.ReadStringFromBytes(bArray, startIndex)

	inputFieldName := "InputDirectory:"

	idx := strings.Index(readDir, inputFieldName)

	if idx == -1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Directory Field Name '%v' was not found in first\n" +
			"string read from 'inputFileMgr'.\n" +
			"inputFileMgr='%v'\n", inputFieldName, inputFileMgr.GetAbsolutePath())
		return inputDir, outputDir, err
	}

	idx += len(inputFieldName)
	inputDir = readDir[idx:]
	inputDir = strings.TrimLeft(inputDir, " ")
	inputDir = strings.TrimRight(inputDir, " ")
	inputDir = strings.ToLower(inputDir)

	doesExist := pathfileops.FileHelper{}.DoesFileExist(inputDir)

	if !doesExist {
		err = fmt.Errorf(ePrefix +
			"\nTarget Zone Input Directory DOES NOT EXIST!\n" +
			"inputDir='%v'\n" +
			"inputFileMgr='%v'", inputDir, inputFileMgr.GetAbsolutePath())

		return inputDir, outputDir, err
	}

	readDir, _ = strops.StrOps{}.ReadStringFromBytes(bArray, startIndex)

	outputFieldName := "OutputDirectory:"

	idx = strings.Index(readDir, outputFieldName)

	if idx == -1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Output Directory Field Name '%v' was not found in second\n" +
			"string read from 'inputFileMgr'.\n" +
			"inputFileMgr='%v'\n", inputFieldName, inputFileMgr.GetAbsolutePath())
		return inputDir, outputDir, err
	}

	idx += len(outputFieldName)
	outputDir = readDir[idx:]
	outputDir = strings.TrimLeft(outputDir, " ")
	outputDir = strings.TrimRight(outputDir, " ")

	doesExist = pathfileops.FileHelper{}.DoesFileExist(outputDir)

	if !doesExist {
		err = fmt.Errorf(ePrefix +
			"\nTarget Zone Output Directory DOES NOT EXIST!\n" +
			"outputDir='%v'\n" +
			"inputFileMgr='%v'",
			outputDir, inputFileMgr.GetAbsolutePath())

		return inputDir, outputDir, err
	}

	return inputDir, outputDir, err
}

// SetCurrentWorkingDirectory - Identifies and
// initializes global variables with the current
// working directory
func (acTzDat AcquireTzData) SetCurrentWorkingDirectory(
	ePrefix string) (currWorkingDirMgr pathfileops.DirMgr, err error) {

	ePrefix += "AcquireTzData.SetCurrentWorkingDirectory() "
	currWorkingDirMgr = pathfileops.DirMgr{}
	err = nil
	var err2 error
	var crDir string

	crDir, err2 = pathfileops.FileHelper{}.GetCurrentDir()

		if err2 != nil {
			err = fmt.Errorf(ePrefix+"\n%v\n", err2.Error())
			return currWorkingDirMgr, err
		}


	tzdatastructs.CurWorkingDirectory, err2 = pathfileops.DirMgr{}.New(crDir)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"\n%v\n", err2.Error())
		return currWorkingDirMgr, err
	}

	var curWorkDirDoesExist bool
	curWorkDirDoesExist, err2 = tzdatastructs.CurWorkingDirectory.DoesThisDirectoryExist()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned by curWorkingDirectory.DoesThisDirectoryExist()\n" +
			"Current Working Directory: %v\n" +
			"Error: %v\n",
			tzdatastructs.CurWorkingDirectory.GetAbsolutePath(), err2.Error())
		return currWorkingDirMgr, err
	}

	if !curWorkDirDoesExist {

		err = fmt.Errorf(ePrefix+
			"Current Working Directory DOES NOT EXIST!"+
			"CurWorkingDir: %v", tzdatastructs.CurWorkingDirectory.GetAbsolutePath())
		return currWorkingDirMgr, err
	}

	currWorkingDirMgr = tzdatastructs.CurWorkingDirectory.CopyOut()
	err = nil

	return currWorkingDirMgr, err
}
