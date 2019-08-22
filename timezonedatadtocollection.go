package main

import (
  "errors"
  "fmt"
  "io"
  "strings"
)

// TimeZoneDataCollection - is a collection TimeZoneDataDto objects.
// The collection effectively encapsulates a TimeZoneDataDto array.
type TimeZoneDataCollection struct {
  tzDataDtos []TimeZoneDataDto
}

// AddTimeZoneDataDto - Adds a TimeZoneDataDto object to the collection
func (tzDataCol *TimeZoneDataCollection) AddTimeZoneDataDto(tzDataDto TimeZoneDataDto) error {

  ePrefix := "TimeZoneDataCollection.AddTimeZoneDataDto() "

  if !tzDataDto.IsInitialized {
    return errors.New(ePrefix + "Input Parameter 'tzDataDto' is uninitialized and INVALID!\n")
  }

  if tzDataCol.tzDataDtos == nil {

    tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 50)

  }

  tzDataCol.tzDataDtos = append(tzDataCol.tzDataDtos, tzDataDto)

  return nil
}

// MajorGroupExists - Performs a search for on the internal TimeZoneDataDto
// array for a match on TimeZoneDataDto.MajorGroup. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of 'false' is returned and the
// integer index value is set to -1.
//
// If the input parameter 'useLwrCase' is set to 'true', the search for
// TimeZoneDataDto.MajorGroup will be conducted as a case insensitive search.
// This means that both strings are converted to lower case before the comparison
// is performed.
//
// If the input parameter 'useLwrCase' is set to 'false', the search
// is conducted as a case sensitive comparison where upper and lower case
// characters are significant.
//
func (tzDataCol *TimeZoneDataCollection) MajorGroupExists(
  majorGroupName string, useLwrCase bool) (majorGrpExists bool, index int) {

  majorGrpExists = false
  index = -1

  if tzDataCol.tzDataDtos == nil {

    tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 50)

  }

  lenTzDataDtoArray := len(tzDataCol.tzDataDtos)

  if lenTzDataDtoArray == 0 {
    return false, -1
  }

  if useLwrCase {
    majorGroupName = strings.ToLower(majorGroupName)
  }

  for i:=0; i < lenTzDataDtoArray; i++ {

    if useLwrCase {
      if majorGroupName == strings.ToLower(tzDataCol.tzDataDtos[i].MajorGroup) {
        return true, i
      }
    } else if majorGroupName == tzDataCol.tzDataDtos[i].MajorGroup {
        return true, i
    }

  }

  return false, -1
}

// Peek - Returns a deep copy of the TimeZoneDataDto located in the internal
// TimeZoneDataDto array at input parameter 'index'.
//
// The internal array is not altered by this method.
//
func (tzDataCol *TimeZoneDataCollection) Peek(index int) (TimeZoneDataDto, error) {

  ePrefix := "TimeZoneDataCollection.Peek() "

  if tzDataCol.tzDataDtos == nil {

    tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 50)

  }

  lenTzDataDtos := len(tzDataCol.tzDataDtos)

  if lenTzDataDtos == 0 {
    return TimeZoneDataDto{}, io.EOF
  }

  if index < 0 {
    return TimeZoneDataDto{},
      fmt.Errorf(ePrefix +
        "ERROR: Input parameter 'index' is less than zero and INVALID!\n" +
        "index='%v'", index)
  }

  if index > (lenTzDataDtos - 1) {
    return TimeZoneDataDto{},
      fmt.Errorf(ePrefix +
        "ERROR: Input paramter 'index' exceeds array upper boundary.\n" +
        "TimeZoneDataDto Array last index='%v'\n" +
        "Input parameter 'index'='%v'\n ", lenTzDataDtos - 1, index )
  }

  return tzDataCol.tzDataDtos[index].CopyOut(), nil
}
