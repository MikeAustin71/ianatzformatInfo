package main

import (
  "errors"
  "fmt"
  "github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
  "github.com/MikeAustin71/stringopsgo/strops/v2"
  "strings"
)

var skipFiles = []string{
  "backzone",
  "backward",
  "calendars",
  "CONTRIBUTING",
  "leapseconds",
  "LICENSE",
  "Makefile",
  "NEWS",
  "README",
  "systemv",
  "version"   }

var timezoneArray []string

// [linked Zone] primary zone
var mapTzLinks map[string]string

var DEBUG = 0

const inputFileName = "targettzdata.txt"

const outputfilename = "timezonedata.go"

const homeDir = "D:\\gowork\\src\\MikeAustin71\\ianatzoneinfo\\parsetzdata"

var curWorkingDirectory pathfileops.DirMgr

func main() {
  ePrefix := "parsetzdata.main() "

  timezoneArray = make([]string, 0, 800)

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

  err = parseTzAndLinks(dirFileInfo)
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

func createInputFileMgr(targetFileName string) (inputFileMgr pathfileops.FileMgr, err error) {

  inputFileMgr = pathfileops.FileMgr{}
  err = nil

  ePrefix := "parsetzdata.createInputFileMgr() Error: "
  var err2 error

  //targetFile := "D:/gowork/src/MikeAustin71/ianatzoneinfo/converttimezones/targetzoneinfo.txt"

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

func extractLink(rawString string) {

  lenRawStr := len(rawString)

  if lenRawStr < 3 {
    return
  }

  idx := strings.Index(rawString, "Link")

  if idx == -1 {
    return
  }

  str1Status := 0
  str2Status := 0

  sb1 := strings.Builder{}
  sb1.Grow(lenRawStr + 10)

  sb2 := strings.Builder{}
  sb2.Grow(lenRawStr + 10)

  i:= 0

  if idx > 0 {

    poundIdx := strings.Index(rawString, "#")

    if poundIdx < idx {
      return
    }

    if idx > 5 {
      return
    }

    if idx >= lenRawStr {
      return
    }

  }

  for i=idx ; i < lenRawStr; i ++ {

    b := rawString[i]

    if b=='\t' ||
      b== '\r' ||
      b== '\n' ||
      b== '#'  ||
      b== ' ' {

      if str1Status == 1 {
        str1Status = 2
        continue
      }

      if str2Status == 1 {
        break
      }

      continue
    }

    if (b >= 'a' && b <= 'z') ||
      (b >= 'A' && b <= 'Z')  ||
      (b>= '0' && b <= '9')   ||
      b == '/'                ||
      b == '_'                ||
      b == '-'                {


      if str1Status == 0 ||
        str1Status == 1  {

        str1Status = 1
        sb1.WriteByte(b)
        continue
      }

      if str1Status == 2 {
        str2Status = 1
        sb2.WriteByte(b)
      }

    }

  }

  if sb1.Len() == 0 ||
    sb2.Len() == 0 {
    return
  }

  mapTzLinks[sb1.String()] = sb2.String()

  return
}

func extractZone(rawString string)  {

  lenRawStr := len(rawString)

  if lenRawStr < 3 {
    return
  }

  idx := strings.Index(rawString, "Zone")

  if idx == -1 {
    return
  }

  poundIdx := strings.Index(rawString, "#")

  if poundIdx < idx {
    return
  }

  idx += 4

  startString := false

  sb := strings.Builder{}
  sb.Grow(lenRawStr + 10)

  i:= 0

  if idx > 0 {

    if idx > 15 {
      return
    }

    if idx >= lenRawStr {
      return
    }

  }

  for i=idx ; i < lenRawStr; i ++ {

    b := rawString[i]

    if b=='\t' ||
        b== '\r' ||
        b== '\n' ||
        b== '#'  ||
        b== ' ' {

      if startString {
          break
      }

      continue

    }

    if (b >= 'a' && b <= 'z')   ||
        (b >= 'A' && b <= 'Z')  ||
        (b>= '0' && b <= '9')   ||
        b == '/'                ||
        b == '_'                ||
        b == '-'                {

        sb.WriteByte(b)
        startString = true
    }

  }

  if sb.Len() == 0 {
    return
  }

  timezoneArray = append(timezoneArray, sb.String())

  return
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

func parseTzAndLinks(dirFileInfo pathfileops.FileMgrCollection) error {
  ePrefix := "parsetzdata.parseTzAndLinks() "

  numOfFiles := dirFileInfo.GetNumOfFileMgrs()
  fmt.Println("Number of Target Files: ", numOfFiles)

  for i:=0; i < numOfFiles; i++ {

    fmgr, err := dirFileInfo.PeekFileMgrAtIndex(i)

    if err != nil {
      return fmt.Errorf(ePrefix+"%v\n", err.Error())
    }

    fmt.Println("Valid File: ", fmgr.GetFileNameExt())

    err = fmgr.OpenThisFileReadOnly()

    if err != nil {
      return fmt.Errorf(ePrefix+"%v\n", err.Error())
    }

    bytes, err := fmgr.ReadAllFile()

    if err != nil {
      _ = fmgr.CloseThisFile()
      return fmt.Errorf(ePrefix+"%v\n", err.Error())
    }

    fptr := fmgr.GetFilePtr()

    err = fptr.Close()

    if err != nil {
      return fmt.Errorf(ePrefix+"Error closing file. File='%v' Error='%v'\n",
        fmgr.GetAbsolutePathFileName(), err.Error())
    }

    nextStartIdx := 0
    extractedString := ""
    cntr := 1
    for nextStartIdx > -1 {

      extractedString, nextStartIdx = strops.StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)
      fmt.Printf("str No %v: %v\n", cntr, extractedString)
      cntr++
      extractZone(extractedString)

      extractLink(extractedString)

    }

  }

  return nil
}

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