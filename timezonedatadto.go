package main

import (
  "errors"
  "fmt"
  "strings"
)

type TimeZoneDataDto struct {
  MajorGroup string
  SubTzName string
  TzName string
  TzValue string
  SourceFileNameExt string
  TzClass TimeZoneClass       // 0 = Unknown
  // 1 = Canonical
  // 2 = Alias
  // 3 = Sub-Group Place Holder

  DeprecationStatus TimeZoneDeprecationStatus   // 0 = Unknown
  // 1 = Deprecated
  // 2 = Alias
  // 3 = Valid, Current Time Zone
  IsInitialized bool
}

// New - Creates and returns a new instance of the TimeZoneDataDto Type.
//
func (tzDataDto TimeZoneDataDto) New(
  majorGroup,
  subTzName,
  tzName,
  tzValue,
  srcFileNameExt string,
  tzClass TimeZoneClass,
  deprecationStatus TimeZoneDeprecationStatus) (TimeZoneDataDto, error) {

  ePrefix := "TimeZoneDataDto.NewTimeZone() - ERROR:\n"
  newTzDto := TimeZoneDataDto{}

  if len(majorGroup) == 0 {
    return newTzDto,
      errors.New(ePrefix + "Input Parameter 'majorGroup' is an EMPTY string!\n")
  }

  // subTzName empty strings are allowed.

  if len(tzName) == 0 {
    return newTzDto,
      errors.New(ePrefix + "Input Parameter 'tzName' is an EMPTY string!\n")
  }

  if len(tzValue) == 0 {
    return newTzDto,
      errors.New(ePrefix + "Input Parameter 'tzValue' is an EMPTY string!\n")
  }

  err := tzClass.StatusIsValid()

  if err != nil {
    return newTzDto,
      fmt.Errorf(ePrefix + "Input Parameter 'tzClass' is INVALID!\n" +
        "tzClass='%v'", int(tzClass))
  }

  err = deprecationStatus.StatusIsValid()

  if err != nil {
    return TimeZoneDataDto{},
      fmt.Errorf(ePrefix + "Input Parameter 'deprecationStatus' is INVALID!\n" +
        "deprecationStatus='%v'", int(deprecationStatus))
  }


  newTzDto.MajorGroup = majorGroup
  newTzDto.SubTzName = subTzName
  newTzDto.TzName = tzName
  newTzDto.TzValue = tzValue
  newTzDto.SourceFileNameExt = srcFileNameExt
  newTzDto.TzClass = tzClass
  newTzDto.DeprecationStatus = deprecationStatus
  newTzDto.IsInitialized = true

  return newTzDto, nil
}

// CopyOut - Creates and returns a deep copy of the current
// TimeZoneDataDto instance.
//
func (tzDataDto *TimeZoneDataDto) CopyOut() TimeZoneDataDto {

  newTzDto := TimeZoneDataDto{}

  if !tzDataDto.IsInitialized {
    return newTzDto
  }

  newTzDto.MajorGroup = tzDataDto.MajorGroup
  newTzDto.SubTzName = tzDataDto.SubTzName
  newTzDto.TzName = tzDataDto.TzName
  newTzDto.TzValue = tzDataDto.TzValue
  newTzDto.SourceFileNameExt = tzDataDto.SourceFileNameExt
  newTzDto.TzClass = tzDataDto.TzClass
  newTzDto.DeprecationStatus = tzDataDto.DeprecationStatus
  newTzDto.IsInitialized = tzDataDto.IsInitialized

  return newTzDto
}

// CopyIn - Receives an input parameter TimeZoneDataDto instance
// copies all of the data fields to the current TimeZoneDataDto instance.
// When complete, both TimeZoneDataDto instances are equivalent.
//
func (tzDataDto *TimeZoneDataDto) CopyIn(
  inTzDataDto *TimeZoneDataDto) {

  tzDataDto.MajorGroup = inTzDataDto.MajorGroup
  tzDataDto.SubTzName = inTzDataDto.SubTzName
  tzDataDto.TzName = inTzDataDto.TzName
  tzDataDto.TzValue = inTzDataDto.TzValue
  tzDataDto.SourceFileNameExt = inTzDataDto.SourceFileNameExt
  tzDataDto.TzClass = inTzDataDto.TzClass
  tzDataDto.DeprecationStatus = inTzDataDto.DeprecationStatus
  tzDataDto.IsInitialized = inTzDataDto.IsInitialized

}

// SortByTzDtoName - Sort by MajorGroup Name, TzName
//
// Example Usage:
//    sort.Sort(SortByTzMajorGroupName(tzMajorGroupDtoArray))
//
type SortByTzDtoName []TimeZoneDataDto

// Len - Required by the sort.Interface
func (sortByTzDtoName SortByTzDtoName) Len() int {
  return len(sortByTzDtoName)
}

// Swap - Required by the sort.Interface
func (sortByTzDtoName SortByTzDtoName) Swap(i, j int) {
  sortByTzDtoName[i], sortByTzDtoName[j] = sortByTzDtoName[j], sortByTzDtoName[i]
}

// Less - required by the sort.Interface
func (sortByTzDtoName SortByTzDtoName) Less(i, j int) bool {

  if sortByTzDtoName[i].MajorGroup == sortByTzDtoName[j].MajorGroup {
    return sortByTzDtoName[i].TzName < sortByTzDtoName[j].TzName
  }

  return sortByTzDtoName[i].MajorGroup < sortByTzDtoName[j].MajorGroup
}

//SelectTzDto - Select from an array TimeZoneDataDto objects.
type SelectTzDto []TimeZoneDataDto

// MajorGroupExists - Performs a search for on TimeZoneDataDto array
// for a match on TimeZoneDataDto.MajorGroup. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of false is returned and the
// integer index value is set to -1.
//
func (selTzDto SelectTzDto) MajorGroupExists(majorGroupName string, useLwrCase bool) (bool, int){

  if useLwrCase {
    majorGroupName = strings.ToLower(majorGroupName)
  }

  for i:=0; i < len(selTzDto); i++ {

    if useLwrCase {

      if strings.ToLower(selTzDto[i].MajorGroup) == majorGroupName {
        return true, i
      }
    } else if selTzDto[i].MajorGroup == majorGroupName {

      return true, i
    }

  }

  return false, -1
}

// SubTzNameExists - Performs a search for on TimeZoneDataDto array
// for a match on TimeZoneDataDto.SubTzName. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of false is returned and the
// integer index value is set to -1.
//
func (selTzDto SelectTzDto) SubTzNameExists(
  subTzName string, useLwrCase bool) (bool, int) {

  if useLwrCase {
    subTzName = strings.ToLower(subTzName)
  }

  for i:=0; i < len(selTzDto); i++ {

    if useLwrCase {
      if strings.ToLower(selTzDto[i].SubTzName) == subTzName{
        return true, i
      }
    } else if selTzDto[i].SubTzName == subTzName {

      return true, i
    }

  }

  return false, -1

}

// SubTzNameExists - Performs a search for on TimeZoneDataDto array
// for a match on TimeZoneDataDto.TzName. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of false is returned and the
// integer index value is set to -1.
//
func (selTzDto SelectTzDto) TzNameExists(
  tzName string, useLwrCase bool) (bool,int) {

  if useLwrCase {
    tzName = strings.ToLower(tzName)
  }

  for i:=0; i < len(selTzDto); i++ {

    if useLwrCase {
      if strings.ToLower(selTzDto[i].TzName) == tzName{
        return true, i
      }
    } else if selTzDto[i].TzName == tzName {

      return true, i
    }

  }
  return false, -1
}


// TzValueExists - Performs a search for on TimeZoneDataDto array
// for a match on TimeZoneDataDto.TzValue. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of false is returned and the
// integer index value is set to -1.
//
func (selTzDto SelectTzDto) TzValueExists(
  tzValue string, useLwrCase bool) (bool, int) {

  if useLwrCase {
    tzValue = strings.ToLower(tzValue)
  }

  for i:=0; i < len(selTzDto); i++ {

    if useLwrCase {
      if tzValue == strings.ToLower(selTzDto[i].TzValue) {
        return true, i
      }
    } else if selTzDto[i].TzValue == tzValue {
      return true, i
    }

  }
  return false, -1
}

