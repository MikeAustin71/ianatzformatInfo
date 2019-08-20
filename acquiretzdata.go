package main

import (
  "errors"
  "fmt"
  "github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
  "github.com/MikeAustin71/stringopsgo/strops/v2"
  "strings"
)


func createInputFileMgr(targetFileName string) (inputFileMgr pathfileops.FileMgr, err error) {

  inputFileMgr = pathfileops.FileMgr{}
  err = nil

  ePrefix := "parsetzdata.createInputFileMgr() Error: "
  var err2 error

  //targetFile := "D:\gowork\src\MikeAustin71\ianatzformatInfo\targettzdata.txt"

  inputFileMgr, err2 = pathfileops.FileMgr{}.NewFromDirMgrFileNameExt(curWorkingDirectory, targetFileName)

  if err2 != nil {

    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())

    return inputFileMgr, err
  }

  ok, err2 := inputFileMgr.DoesThisFileExist()

  if err2 != nil {

    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())

    return inputFileMgr, err
  }

  if !ok {
    err = fmt.Errorf(ePrefix+
      "Target File DOES NOT EXIST! targetFile='%v' ",
      inputFileMgr.GetAbsolutePathFileName())
    return inputFileMgr, err
  }

  err = nil

  return inputFileMgr, err
}

// getDirectoryInfo - Walks the 'baseDir' and returns information on all directories and
// files. Note: Input parameter 'baseDir' MUST BE formatted in all lower case characters.
//
func getDirectoryInfo(baseDir string) (dirFileInfo pathfileops.FileMgrCollection, err error) {

  ePrefix := "parsetzdata.getDirectoryInfo() "

  dirFileInfo = pathfileops.FileMgrCollection{}.New()
  err = nil

  if len(baseDir) < 3 {
    err = fmt.Errorf(ePrefix+"Input parameter 'baseDir' is INVALID! baseDir='%v' ", baseDir)
    return dirFileInfo, err
  }

  baseDirMgr, err2 := pathfileops.DirMgr{}.New(baseDir)

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
    return dirFileInfo, err
  }

  baseDMgrDoesExist, err2 := baseDirMgr.DoesThisDirectoryExist()

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "baseDirMgr Non-Path Error\n" +
      "%v", err2.Error())
    return dirFileInfo, err
  }

  if !baseDMgrDoesExist {
    err = fmt.Errorf(ePrefix+"baseDirMgr Path DOES NOT EXIST! baseDir= %v ", baseDir)
    return dirFileInfo, err
  }

  initialFMgrCol, err2 := baseDirMgr.FindFilesByNamePattern("*")

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v", err2.Error())
    return dirFileInfo, err
  }

  lenInitialFMgrs := initialFMgrCol.GetNumOfFileMgrs()

  if lenInitialFMgrs < 1 {
    err = errors.New(ePrefix + "Error: Zero Files Found in Target Directory!")
    return dirFileInfo, err
  }

  lenSkipFiles := len(skipFiles)

  var fmgr pathfileops.FileMgr

  for i:=0; i< lenInitialFMgrs; i++ {

    fmgr, err2 = initialFMgrCol.PeekFileMgrAtIndex(i)

    if err2 != nil {
      err = fmt.Errorf(ePrefix + "Error returned by initialFMgrCol.PeekFileMgrAtIndex(i)\n" +
        "i='%v'\n" +
        "Error='%v'\n", i, err2.Error())
      return dirFileInfo, err
    }

    fName := fmgr.GetFileName()
    isInvalidFile := false

    if fmgr.GetFileExt() == "" {

      for j:=0; j < lenSkipFiles; j++ {

        if fName == skipFiles[j] {
          isInvalidFile = true
          break
        }
      }

      if !isInvalidFile && fName == "asia" {
        dirFileInfo.AddFileMgr(fmgr.CopyOut())
      }
    }

  }

  /*
    if dirFileInfo.GetNumOfFileMgrs() < 2 {
      err = fmt.Errorf(ePrefix +
        "Error: Number of Valid Files Found in Target Directory is INVALID! " +
        "Number of valid files= '%v' ", dirFileInfo.GetNumOfFileMgrs())

      dirFileInfo = pathfileops.FileMgrCollection{}.New()

      return dirFileInfo, err
    }
  */

  err = nil
  return dirFileInfo, err
}

// getTargetDirectory - Reads the targettzdata.txt file to get the target directory
// where IANA Time Zone data is located.
//
func getTargetDirectory(inputFileMgr pathfileops.FileMgr) (baseDir string, err error) {
  baseDir = ""
  err = nil

  ePrefix := "parsetzdata.getTargetDirectory() Error: "

  err2 := inputFileMgr.OpenThisFileReadWrite()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
    return baseDir, err
  }

  bArray, err2 := inputFileMgr.ReadAllFile()

  if err2 != nil {

    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())

    err2 = inputFileMgr.CloseThisFile()

    if err2 != nil {
      errStr := err.Error()
      errStr += "\n"
      errStr += err2.Error()
      err = errors.New(errStr)
    }

    return baseDir, err
  }

  lBArray := len(bArray)

  if lBArray < 3 {
    err = fmt.Errorf(ePrefix+"Read Zero bytes from file %v", inputFileMgr)
    _ = inputFileMgr.CloseThisFile()
    return baseDir, err
  }

  baseDir, _ = strops.StrOps{}.ReadStringFromBytes(bArray, 0)

  err2 = inputFileMgr.CloseThisFile()

  if err2 != nil {
    err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
    return baseDir, err
  }

  doesExist := pathfileops.FileHelper{}.DoesFileExist(baseDir)

  if !doesExist {
    err = fmt.Errorf("Target Directory: %v DOES NOT EXIST!", baseDir)
    return baseDir, err

  }

  baseDir = strings.ToLower(baseDir)

  return baseDir, err
}

// setCurrentWorkingDirectory - Identifies and
// initializes global variables with the current
// working directory
func setCurrentWorkingDirectory() error {

  ePrefix := "parsetzdata.setCurrentWorkingDirectory() "

  var err error

  crDir := ""

  if DEBUG == 0 {
    // DEBUG is OFF!

    crDir, err = pathfileops.FileHelper{}.GetCurrentDir()

    if err != nil {

      return fmt.Errorf(ePrefix+"%v\n", err.Error())
    }

  } else {

    // DEBUG MUST BE ON
    crDir = homeDir

  }

  curWorkingDirectory, err = pathfileops.DirMgr{}.New(crDir)

  if err != nil {

    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  curWorkDirDoesExist, err := curWorkingDirectory.DoesThisDirectoryExist()

  if err != nil {
    return fmt.Errorf(ePrefix + "Error returned by curWorkingDirectory.DoesThisDirectoryExist()\n" +
      "%v", err.Error())
  }

  if !curWorkDirDoesExist {

    return fmt.Errorf(ePrefix+
      "Current Working Directory DOES NOT EXIST!"+
      "CurWorkingDir: %v", curWorkingDirectory.GetAbsolutePath())

  }


  return nil
}
