package main

import (
  "errors"
  "fmt"
  "strings"
)

// TzMajorGrpMgr - This type provides methods for managing arrays of
// TimeZoneGroupDto instances.
//
type TzMajorGrpMgr []TimeZoneGroupDto


// MajorGroupNameExists - Searches MajorGroupName for a match on input parameter 'majorGrpName'.
//
// If input parameter 'useLowerCase' is 'true', a case insensitive search is employed. This means
// that both strings are first converted to lower case before the comparison is executed.
//
// If input parameter 'useLowerCase' is false, a case sensitive search is employed. This means
// that upper and lower case characters are significant.
//
// If the search is successful, a boolean value of 'true' is returned along with the array index
// of the found TimeZoneGroupDto instance.
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

