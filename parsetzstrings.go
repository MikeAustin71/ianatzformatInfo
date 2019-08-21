package main

import (
  "errors"
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


var tzDataArray = make([]TimeZoneDataDto, 0, 700)

var subTzArray = make([]TimeZoneDataDto, 0, 700)

var tzLinkArray = make([]TimeZoneDataDto, 0, 300)

/*
  tzMajorGroupArray Format
    Africa,
    America,
    Antarctica,
    Artic,
    Asia,
    Atlantic,
    Australia,
      ...
    UCT,
    W_SU,
    WET,
    Zulu
    ------------------------------------------

    tzDataArray Format
      Standard Time Zone -
      ====================
          MajorGroup  = America
          SubTzName   = ""
          TzName      = Chicago
          TzValue     = America/Chicago
          TzClass     = 1
                        // 0 = Unknown
                        // 1 = Canonical
                        // 2 = Alias
                        // 3 = Sub-Group
          Deprecated  = false

        Sub-Zone Place Holder -
        =======================
          MajorGroup  = America
          SubTzName   = Argentina
          TzName      = Argentina
          TzValue     = America/Argentina
          TzClass     = 3
                        // 0 = Unknown
                        // 1 = Canonical
                        // 2 = Alias
                        // 3 = Sub-Group
          Deprecated  = false


*/

type ParseIanaTzData struct {
  input string
  output string

}

// ParseTzAndLinks - Parses Time Zone Data from
// IANA Time Zone files.
func (parseTz *ParseIanaTzData) ParseTzAndLinks(
  dirFileInfo pathfileops.FileMgrCollection) (
                      []TimeZoneMajorGroupDto, // Time Zone Major Group array
                      []TimeZoneDataDto, // Time Zone Data Dto Array
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

func (parseTz *ParseIanaTzData) extractDataElement(
  rawString string) (dataElement string,
                      lengthDataElement int,
                      lastIdx int,
                      err error) {

  ePrefix := "ParseIanaTzData.extractDataElement() "

  dataElement = ""
  lengthDataElement = 0
  lastIdx = 0
  err = nil

  rawRunes := []rune(rawString)
  rawRunesLen := len(rawRunes)

  if rawRunesLen == 0 {
    err = errors.New(ePrefix +
      "Input parameter 'rawString' is EMPTY! Zero string length!\n")
    return dataElement, lengthDataElement, lastIdx, err
  }

  dataElementRunes := make([]rune,0, 30)

  isFieldSeparator := false
  startString := false

  for i:=0; i < rawRunesLen; i++ {

    isFieldSeparator = false

    for j:= 0; j < fieldSeparatorsLen; j++ {
      if rawRunes[i] == FieldSeparators[j] {
        isFieldSeparator = true
        break
      }
    }

    if isFieldSeparator && !startString {
      continue
    } else if isFieldSeparator && startString {
      lastIdx = i
      break
    }

    // isFieldSeparator == false
    // Capture rune which is part of
    // data element.
    startString = true

    dataElementRunes = append(dataElementRunes, rawRunes[i])
  }

  dataElement = string(dataElementRunes)
  lengthDataElement = len(dataElement)

  return dataElement, lengthDataElement, lastIdx, err
}

// extractLink - Extracts link data from IANA Time Zone files.
// Format for Link:
// Link -> Canonical -> Alias
// Link  America/Panama America/Cayman
func (parseTz *ParseIanaTzData) extractLink(rawString string) error {

  ePrefix := "ParseIanaTzData.extractLink() "

  lenRawStr := len(rawString)

  if lenRawStr < lenLinkLabel {
    return nil
  }

  linkIdx := strings.Index(rawString, LinkLabel)

  if linkIdx == -1 {
    return nil
  }

  commentIdx := strings.Index(rawString, CommentCharStr)

  if commentIdx > -1 &&
    commentIdx < linkIdx {
    return nil
  }

  rawStrLen := len(rawString)
  lastRawStrIdx := rawStrLen - 1

  linkIdx += len(LinkLabel)

  if linkIdx >= lastRawStrIdx {
    return nil
  }

  rawString = rawString[linkIdx:]

  rawRunes := []rune(rawString)

  linkRunes := make([]rune, 0, 30)

  isFieldSeparator := false
  startString := false

  rawStrLen = len(rawRunes)

  for i:=0; i < rawStrLen; i++ {
    isFieldSeparator = false

    for j:=0; j < fieldSeparatorsLen; j++ {
      if rawRunes[i] == FieldSeparators[j] {
        isFieldSeparator = true
        break
      }
    }

    if isFieldSeparator && !startString {
      continue
    } else if isFieldSeparator && startString {
      linkIdx = i
      break
    }

    startString = true
    linkRunes = append(linkRunes, rawRunes[i])
  }

  linkStr := string(linkRunes)

  if len(linkStr) == 0 {
    return fmt.Errorf(ePrefix +
      "Invalid Linked Time Zone!\n" +
      "Raw Link String: %v", rawString)

  }

  linkZoneArray := strings.Split(linkStr, "/")
  lenLinkZoneArray := len(linkZoneArray)


  if lenLinkZoneArray == 0 {
    return fmt.Errorf(ePrefix +
      "Link Zone Elements Equal ZERO!\n" +
      "Raw Link Zone String: %v\n", rawString)
  }

  if lenLinkZoneArray == 1 {
    return fmt.Errorf(ePrefix +
      "Link Zone has ONLY one element!\n" +
      "Raw Link Zone String: %v\n", rawString)
  }

  if lenLinkZoneArray > 2 {
    return fmt.Errorf(ePrefix +
      "Link Zone has more than 2 elements!\n" +
      "Raw Link Zone String: %v\n", rawString)
  }

  rawString = rawString[linkIdx:]
  rawRunes = []rune(rawString)
  rawStrLen = len(rawRunes)
  lastRawStrIdx = rawStrLen - 1
  linkRunes = make([]rune, 0, 30)

  for i:=0; i < rawStrLen; i++ {
    isFieldSeparator = false

    for j:=0; j < fieldSeparatorsLen; j++ {
      if rawRunes[i] == FieldSeparators[j] {
        isFieldSeparator = true
        break
      }
    }

    if isFieldSeparator && !startString {
      continue
    } else if isFieldSeparator && startString {
      linkIdx = i
      break
    }

    startString = true
    linkRunes = append(linkRunes, rawRunes[i])
  }


  // lenLinkZoneArray



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

  if lenRawStr < lenZoneLabel {
    return nil
  }

  zoneIdx := strings.Index(rawString, ZoneLabel)

  if zoneIdx == -1 {
    return nil
  }

  commentIdx := strings.Index(rawString, CommentCharStr)

  if commentIdx > -1 &&
    commentIdx < zoneIdx {
    return nil
  }

  rawStrLen := len(rawString)
  lastRawStrIdx := rawStrLen - 1

  zoneIdx += len(ZoneLabel)

  if zoneIdx >= lastRawStrIdx {
    return nil
  }

  rawString = rawString[zoneIdx:]

  if len(rawString) == 0 {
    return errors.New(ePrefix +
      "'rawString' is an empty string!\n")
  }

  zoneStr,
  lenZoneStr,
  lastIdx,
  err := parseTz.extractDataElement(rawString)

  if err != nil {
    return fmt.Errorf("Zone Extraction Error!\n" +
      "rawString='%v'\nError='%v'\n",
      rawString, err.Error())
  }


  if len(zoneStr) == 0 {
    return fmt.Errorf(ePrefix + "Invalid Time Zone!\n" +
      "Raw Zone String: %v", rawString)
  }

  zoneArray := strings.Split(zoneStr, "/")

  lenZoneArray := len(zoneArray)

  if lenZoneArray < 2 ||
    lenZoneArray > 3 {
    return fmt.Errorf(ePrefix +
      "Invalid Time Zone!\n" +
      "Raw Zone String: %v\n", rawString)
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

    tzDataDto.SubTzName = zoneArray[1] // Argentina

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

  for k:=0; k < len(skipTzFiles); k++ {
    if fileName == strings.ToLower(skipTzFiles[k]) {
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

    cmtIdx := strings.Index(extractedString, CommentCharStr)

    zoneIdx := strings.Index(extractedString, ZoneLabel)

    linkIdx := strings.Index(extractedString, LinkLabel)

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