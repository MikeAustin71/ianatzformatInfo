package main

import (
  "errors"
  "fmt"
  "strings"
)

type TimeZoneMajorGroupDto struct {
  MajorGroupName string
  SourceFileNameExt string
  MajorGroupType  MajorTimeZoneGroupType
  DeprecationStatus TimeZoneDeprecationStatus
  IsInitialized bool
}

func (tzMajorGrp TimeZoneMajorGroupDto) New(
  majorGroupName,
  sourceFileNameExt string,
  majorGroupType MajorTimeZoneGroupType,
  deprecationStatus TimeZoneDeprecationStatus) (TimeZoneMajorGroupDto, error) {

  ePrefix := "TimeZoneMajorGroupDto.New() "

  newTzMajorGroup := TimeZoneMajorGroupDto{}

  if len(majorGroupName) == 0 {
    return newTzMajorGroup,
      errors.New(ePrefix + "Input Parameter 'majorGroupName' is an EMPTY string!\n")
  }
  
  err := majorGroupType.StatusIsValid()
  
  if err != nil {
    return TimeZoneMajorGroupDto{},
      fmt.Errorf(ePrefix + "Input Parameter 'majorGroupType' is INVALID!\n" +
        "majorGroupType='%v'", int(majorGroupType))
  }

  err = deprecationStatus.StatusIsValid()

  if err != nil {
    return TimeZoneMajorGroupDto{},
      fmt.Errorf(ePrefix + "Input Parameter 'deprecationStatus' is INVALID!\n" +
        "deprecationStatus='%v'", int(deprecationStatus))
  }

  newTzMajorGroup.MajorGroupName = majorGroupName
  newTzMajorGroup.MajorGroupType = majorGroupType
  newTzMajorGroup.DeprecationStatus = deprecationStatus
  newTzMajorGroup.IsInitialized = true
  
  return newTzMajorGroup, nil
}


// CopyOut - Creates and returns a deep copy of the current
// TimeZoneMajorGroupDto instance.
//
func (tzMajorGrp *TimeZoneMajorGroupDto) CopyOut() TimeZoneMajorGroupDto {

  newTzMjrGrp := TimeZoneMajorGroupDto{}

  if !tzMajorGrp.IsInitialized {
    return newTzMjrGrp
  }

  newTzMjrGrp.MajorGroupName    = tzMajorGrp.MajorGroupName
  newTzMjrGrp.MajorGroupType    = tzMajorGrp.MajorGroupType
  newTzMjrGrp.SourceFileNameExt = tzMajorGrp.SourceFileNameExt
  newTzMjrGrp.DeprecationStatus = tzMajorGrp.DeprecationStatus
  newTzMjrGrp.IsInitialized     = true
  
  return newTzMjrGrp
}

// CopyIn - Receives an input parameter TimeZoneMajorGroupDto instance
// copies all of the data fields to the current TimeZoneMajorGroupDto
// instance.
//
// When complete, both TimeZoneMajorGroupDto instances are equivalent.
//
func (tzMajorGrp *TimeZoneMajorGroupDto) CopyIn(
  inMajorGrpDto *TimeZoneMajorGroupDto) {

  tzMajorGrp.MajorGroupName    = inMajorGrpDto.MajorGroupName
  tzMajorGrp.MajorGroupType    = inMajorGrpDto.MajorGroupType
  tzMajorGrp.SourceFileNameExt = inMajorGrpDto.SourceFileNameExt
  tzMajorGrp.DeprecationStatus = inMajorGrpDto.DeprecationStatus
  tzMajorGrp.IsInitialized     = inMajorGrpDto.IsInitialized
}

// SortByTzMajorGroupName - This type provides support methods for
// sorting Time Zone Major Group Dto Arrays by Major Group Name.
//
// Example Usage:
//    sort.Sort(SortByTzMajorGroupName(tzMajorGroupDtoArray))
//
type SortByTzMajorGroupName []TimeZoneMajorGroupDto

// Len - Required by the sort.Interface
func (sortMjrGrpName SortByTzMajorGroupName) len() int {
  return len(sortMjrGrpName)
}

// Swap - Required by the sort.Interface
func (sortMjrGrpName SortByTzMajorGroupName) Swap(i, j int) {
  sortMjrGrpName[i], sortMjrGrpName[j] = sortMjrGrpName[j], sortMjrGrpName[i]
}

// Less - Required by the sort.Interface
func (sortMjrGrpName SortByTzMajorGroupName) Less(i, j int) bool {
 return sortMjrGrpName[i].MajorGroupName < sortMjrGrpName[j].MajorGroupName
}

// TzMajorGrpMgr - This type provides methods for managing arrays of
// TimeZoneMajorGroupDto instances.
//
type TzMajorGrpMgr []TimeZoneMajorGroupDto


// MajorGroupNameExists - Searches MajorGroupName for a match on input parameter 'majorGrpName'.
//
// If input parameter 'useLowerCase' is 'true', a case insensitive search is employed. This means
// that both strings are first converted to lower case before the comparison is executed.
//
// If input parameter 'useLowerCase' is false, a case sensitive search is employed. This means
// that upper and lower case characters are significant.
//
// If the search is successful, a boolean value of 'true' is returned along with the array index
// of the found TimeZoneMajorGroupDto instance.
//
func (tzMjrGrpMgr TzMajorGrpMgr) MajorGroupNameExists(
  majorGrpName string,
  useLowerCase bool) (exists bool, foundIndex int) {

  exists = false
  foundIndex = -1

  if useLowerCase {
    majorGrpName = strings.ToLower(majorGrpName)
  }

  for i:=0; i < len(tzMjrGrpMgr); i++ {

    if useLowerCase {

      if majorGrpName == strings.ToLower(tzMjrGrpMgr[i].MajorGroupName) {
        exists = true
        foundIndex = i
        break
      }

    } else if majorGrpName == tzMjrGrpMgr[i].MajorGroupName {
      exists = true
      foundIndex = i
      break
    }
  }
    
  return exists, foundIndex
}

