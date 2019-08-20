package main

import (
  "fmt"
  "github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
  "github.com/MikeAustin71/stringopsgo/strops/v2"
  "strings"
)

/*
  IANA Time Zone Notes

  How to Read the tz Database Source Files
    https://data.iana.org/time-zones/tz-how-to.html

	See File  zic.8.txt, in the code subdirectory.


      Input files use the format described in this section; output files use
      tzfile(5) format.

      Input files should be text files, that is, they should be a series of
      zero or more lines, each ending in a newline byte and containing at
      most 511 bytes, and without any NUL bytes.  The input text's encoding
      is typically UTF-8 or ASCII; it should have a unibyte representation
      for the POSIX Portable Character Set (PPCS) <http://pubs.opengroup.org/
      onlinepubs/9699919799/basedefs/V1_chap06.html> and the encoding's non-
      unibyte characters should consist entirely of non-PPCS bytes.  Non-PPCS
      characters typically occur only in comments: although output file names
      and time zone abbreviations can contain nearly any character, other
      software will work better if these are limited to the restricted syntax
      described under the -v option.


      Input lines are made up of fields.  Fields are separated from one
      another by one or more white space characters.  The white space
      characters are space, form feed, carriage return, newline, tab, and
      vertical tab.  Leading and trailing white space on input lines is
      ignored.  An unquoted sharp character (#) in the input introduces a
      comment which extends to the end of the line the sharp character
      appears on.  White space characters and sharp characters may be
      enclosed in double quotes (") if they're to be used as part of a field.
      Any line that is blank (after comment stripping) is ignored.  Non-blank
      lines are expected to be of one of three types: rule lines, zone lines,
      and link lines.

 */

type TimeZoneMajorGroups []string

type TimeZoneDataDto struct {
  MajorGroup string
  SubTzName string
  TzName string
  TzValue string
  TzClass int       // 0 = Unknown
  // 1 = Canonical
  // 2 = Alias
  // 3 = Sub-Group
}

func (tzDataDto TimeZoneDataDto) NewTimeZone(
  majorGroup,
  tzName,
  tzValue string,
  tzClass int) (TimeZoneDataDto, error) {

  ePrefix := "TimeZoneDataDto.NewTimeZone() - ERROR:\n"

  if tzClass < 1 || tzClass > 3 {
    return TimeZoneDataDto{},
      fmt.Errorf(ePrefix + "Input Parameter tzClass is out of bounds and INVALID!\n" +
        "Valid values are 1-3!\ntzClass='%v'", tzClass)
  }

  tzDto := TimeZoneDataDto{}
  tzDto.MajorGroup = majorGroup
  tzDto.TzName = tzName
  tzDto.TzValue = tzValue
  tzDto.TzClass = tzClass

  return tzDto, nil
}

func (tzDataDto TimeZoneDataDto) NewSubTimeZone(
  majorGroup,
  subTimeZoneName,
  tzName,
  tzValue string,
  tzClass int) (TimeZoneDataDto, error) {

  ePrefix := "TimeZoneDataDto.NewSubTimeZone() - ERROR:\n"

  if tzClass < 1 || tzClass > 3 {
    return TimeZoneDataDto{},
      fmt.Errorf(ePrefix + "Input Parameter tzClass is out of bounds and INVALID!\n" +
        "Valid values are 1-3!\ntzClass='%v'", tzClass)
  }

  tzDto := TimeZoneDataDto{}
  tzDto.MajorGroup = majorGroup
  tzDto.SubTzName = subTimeZoneName
  tzDto.TzName = tzName
  tzDto.TzValue = tzValue
  tzDto.TzClass = tzClass

  return tzDto, nil

}

//ByTzDtoName - Sort by MajorGroup, TzName
type ByTzDtoName []TimeZoneDataDto

func (byTzDtoName ByTzDtoName) Len() int {
  return len(byTzDtoName)
}

func (byTzDtoName ByTzDtoName) Swap(i, j int) {
  byTzDtoName[i], byTzDtoName[j] = byTzDtoName[j], byTzDtoName[i]
}

func (byTzDtoName ByTzDtoName) Less(i, j int) bool {

  if byTzDtoName[i].MajorGroup == byTzDtoName[j].MajorGroup {
    return byTzDtoName[i].TzName < byTzDtoName[j].TzName
  }

  return byTzDtoName[i].MajorGroup < byTzDtoName[j].MajorGroup
}

//SelectTzDto - Select from array TimeZoneDataDto
type SelectTzDto []TimeZoneDataDto

func (selTzDto SelectTzDto) GroupExists(majorGroupName string) bool{

  for i:=0; i < len(selTzDto); i++ {

    if selTzDto[i].MajorGroup == majorGroupName {
      return true
    }

  }

  return false
}

func (selTzDto SelectTzDto) TzNameExists(tzName string) bool {

  for i:=0; i < len(selTzDto); i++ {

    if selTzDto[i].TzName == tzName {
      return true
    }

  }
  return false
}


const commentCharStr = "#"
const zoneLabel = "Zone"
const linkLabel = "Link"

// For IANA Time Zone Files the white space characters which delimit fields
// are space, form feed, carriage return, newline, tab, and  vertical tab.
//
// ' ' = space
// \f = form feed
// \r = carriage return
// \n = new line
// \t = tab
// \v = vertical tab

var fieldSeparators = []rune{
  ' ',
  '\f',
  '\r',
  '\n',
  '\t',
  '\v'}

var fieldSeparatorsLen = len(fieldSeparators)

var tzMajorGroupArray TimeZoneMajorGroups = make([]string, 0, 100)

var tzMajorGroupMap = make(map[string]string)

var tzDataArray = make([]TimeZoneDataDto, 0, 700)

var subTzArray = make([]TimeZoneDataDto, 0, 700)

var tzLinkArray = make([]TimeZoneDataDto, 0, 300)


type ParseIanaTzData struct {
  input string
  output string

}

// ParseTzAndLinks - Parses Time Zone Data from
// IANA Time Zone files.
func (parseTz *ParseIanaTzData) ParseTzAndLinks(
  dirFileInfo pathfileops.FileMgrCollection) (
                      TimeZoneMajorGroups, // Time Zone Group string array [] string
                      []TimeZoneDataDto, // Time Zone Data Array
                      []TimeZoneDataDto, // Sub-Zone Array
                      []TimeZoneDataDto, // Alias Link Array
                      error)  {




  ePrefix := "parsetzdata.parseTzAndLinks() "

  numOfFiles := dirFileInfo.GetNumOfFileMgrs()
  fmt.Println("Number of Target Files: ", numOfFiles)

  for i:=0; i < numOfFiles; i++ {

    fmgr, err := dirFileInfo.PeekFileMgrAtIndex(i)

    if err != nil {
      return tzMajorGroupArray,
              tzDataArray,
              subTzArray,
              tzLinkArray,
              fmt.Errorf(ePrefix+"%v\n", err.Error())
    }

    isSkipFile, err := parseTz.isSkipFile(fmgr)

    if err != nil {
      return tzMajorGroupArray,
        tzDataArray,
        subTzArray,
        tzLinkArray,
        fmt.Errorf(ePrefix+"%v\n", err.Error())
    }
    
    if isSkipFile {
      continue
    }

    fmt.Println("Valid File: ", fmgr.GetFileNameExt())

    err =  parseTz.processFileBytes(fmgr)

    if err != nil {
      return tzMajorGroupArray,
        tzDataArray,
        subTzArray,
        tzLinkArray,
        fmt.Errorf(ePrefix+
          "File Name: %v\n" +
          "Error=%v\n",
          fmgr.GetAbsolutePathFileName(),  err.Error())
    }

  }



  return tzMajorGroupArray,
    tzDataArray,
    subTzArray,
    tzLinkArray,
    nil

}


// extractLink - Extracts link data from IANA Time Zone files.
// Format for Link:
// Link -> Canonical -> Alias
// Link  America/Panama America/Cayman
func (parseTz *ParseIanaTzData) extractLink(rawString string) {

  lenRawStr := len(rawString)

  if lenRawStr < 3 {
    return
  }

  linkIdx := strings.Index(rawString, linkLabel)

  if linkIdx == -1 {
    return
  }

  commentIdx := strings.Index(rawString, commentCharStr)

  if commentIdx > -1 &&
    commentIdx < linkIdx {
    return
  }

  str1Status := 0
  str2Status := 0

  sb1 := strings.Builder{}
  sb1.Grow(lenRawStr + 10)

  sb2 := strings.Builder{}
  sb2.Grow(lenRawStr + 10)

  i:= 0

  if linkIdx > 0 {

    poundIdx := strings.Index(rawString, "#")

    if poundIdx < linkIdx {
      return
    }

    if linkIdx > 5 {
      return
    }

    if linkIdx >= lenRawStr {
      return
    }

  }

  for i= linkIdx; i < lenRawStr; i ++ {

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

// extractZone - Extracts standard time zones and sub time zones.
// Data is stored in tzMajorGroupMap, tzDataArray and
// or subTzArray.
func (parseTz *ParseIanaTzData) extractZone(rawString string) error  {

  ePrefix := "ParseIanaTzData.extractZone() "

  lenRawStr := len(rawString)

  if lenRawStr < 3 {
    return nil
  }

  idx := strings.Index(rawString, zoneLabel)

  if idx == -1 {
    return nil
  }

  poundIdx := strings.Index(rawString, commentCharStr)

  if poundIdx < idx {
    return nil
  }

  rawStrLen := len(rawString)
  lastRawStrIdx := rawStrLen - 1

  idx += len(zoneLabel)

  if idx >= lastRawStrIdx {
    return nil
  }

  rawString = rawString[idx:]

  rawRunes := []rune(rawString)

  rawStrLen = len(rawRunes)

  zoneRunes := make([]rune,0, 30)

  isFieldSeparator := false
  startString := false

  for i:=0; i < rawStrLen; i++ {
    isFieldSeparator = false

    for j:= 0; j < fieldSeparatorsLen; j++ {
      if rawRunes[i] == fieldSeparators[j] {
        isFieldSeparator = true
        break
      }
    }

    if isFieldSeparator && !startString {
      continue
    } else if isFieldSeparator && startString {
      break
    }

    startString = true

    zoneRunes = append(zoneRunes, rawRunes[i])
  }

  zoneStr := string(zoneRunes)

  if len(zoneStr) == 0 {
    return fmt.Errorf(ePrefix + "Invalid Time Zone!\n" +
      "Raw Zone String: %v", rawString)
  }

  zoneArray := strings.Split(zoneStr, "/")

  lenZoneArray := len(zoneArray)

  if lenZoneArray < 2 ||
    lenZoneArray > 3 {
    return fmt.Errorf(ePrefix + "Invalid Time Zone!\n" +
      "Raw Zone String: %v", rawString)
  }

  _, ok := tzMajorGroupMap[zoneArray[0]]

  if !ok {
    // The major group has not been captured yet.
    // Add it to the tzMajorGroupMap
    tzMajorGroupMap[zoneArray[0]] = zoneArray[0]
  }

  if lenZoneArray == 2 {

    tzDataDto, err := TimeZoneDataDto{}.NewTimeZone(
      zoneArray[0], // America
      zoneArray[1], // Chicago
      zoneStr, // America/Chicago
      1)
    // tzClass
    // 0 = Unknown
    // 1 = Canonical
    // 2 = Alias
    // 3 = Sub-Group

    if err != nil {
      return fmt.Errorf(ePrefix + "Zone String: %v\n" +
        "Error: %v\n", zoneStr, err.Error())
    }

    tzDataArray = append(tzDataArray, tzDataDto)

    return nil
  }

  // lenZoneArray must == 3
  // This is a sub zone
  zoneFound := false
  // America/Argentina
  zoneSubValue := zoneArray[0] + "/" + zoneArray[1]

  for i:=0; i < len(tzDataArray); i++ {

    if zoneSubValue == tzDataArray[i].TzValue {
      zoneFound = true
      break
    }
  }

  if !zoneFound {
    // Add reference to this group of time zones
    // in the main Time Zone Data Array
    // Example IANA Time Zones for Argentina
    tzDataDto, err := TimeZoneDataDto{}.NewTimeZone(
      zoneArray[0], // America
      zoneArray[1], // Argentina
      zoneSubValue, // America/Argentina
      3)

    if err != nil {
      return fmt.Errorf(ePrefix +
        "Zone Not Found - SubZone Master Zone String: %v\n" +
        "Error: %v\n", zoneStr, err.Error())
    }

    tzDataArray = append(tzDataArray, tzDataDto)
  }

  // Finally, add the Sub Time Zone to the
  // Sub Time Zone Array (subTzArray)
  //America/Argentina/Buenos_Aires
  tzDataDtoSubTz, err := TimeZoneDataDto{}.NewSubTimeZone(
    zoneArray[0],   // America
    zoneArray[1],   // Argentina
    zoneArray[2],   // Buenos_Aires
    zoneStr,        // America/Argentina/Buenos_Aires
    3)

  if err != nil {
    return fmt.Errorf(ePrefix +
      "Sub Array Addition Error - Zone String: %v\n" +
      "Error: %v\n", zoneStr, err.Error())
  }

  subTzArray = append(subTzArray, tzDataDtoSubTz)

  return nil
}

// isSkipFile - Examines the file name of a time zone data
// file and determines whether the file should be skipped 
// for processing.
//
func (parseTz *ParseIanaTzData) isSkipFile(fMgr pathfileops.FileMgr) (bool, error) {

  ePrefix := "ParseIanaTzData.isSkipFile() "
  
  err := fMgr.IsFileMgrValid(ePrefix)

  if err != nil {
    return false, err
  }
  
  if fMgr.GetFileExt() != "" {
    return true, nil
  }

  fileName := strings.ToLower(fMgr.GetFileName())
  isSkipFile := false

  for k:=0; k < len(skipFiles); k++ {
    if fileName == strings.ToLower(skipFiles[k]) {
      isSkipFile = true
      break
    }
  }

  return isSkipFile, nil
}

// ProcessFileBytes - Process all the bytes in a time zone file
//
func (parseTz *ParseIanaTzData) processFileBytes(fMgr pathfileops.FileMgr) error {

  ePrefix := "ParseIanaTzData.processFileBytes() "
  
  err := fMgr.OpenThisFileReadOnly()

  if err != nil {
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  bytes, err := fMgr.ReadAllFile()

  if err != nil {
    _ = fMgr.CloseThisFile()
    return fmt.Errorf(ePrefix+"%v\n", err.Error())
  }

  err = fMgr.CloseThisFile()

  if err != nil {
    return fmt.Errorf(ePrefix+"Error closing file. File='%v' Error='%v'\n",
        fMgr.GetAbsolutePathFileName(), err.Error())
  }

  nextStartIdx := 0
  extractedString := ""
  cntr := 1
  for nextStartIdx > -1 {

    extractedString, nextStartIdx = strops.StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)
    fmt.Printf("str No %v: %v\n", cntr, extractedString)
    cntr++

    cmtIdx := strings.Index(extractedString, commentCharStr)

    zoneIdx := strings.Index(extractedString, zoneLabel)

    linkIdx := strings.Index(extractedString, linkLabel)

    if zoneIdx > -1 {

      if cmtIdx > -1 &&
        cmtIdx < zoneIdx {

        continue
      }

      parseTz.extractZone(extractedString)
      continue
    }

    if linkIdx > -1 {
      if cmtIdx > -1 &&
        cmtIdx < linkIdx {

        continue
      }

      parseTz.extractLink(extractedString)
    }
  }

  return nil
}