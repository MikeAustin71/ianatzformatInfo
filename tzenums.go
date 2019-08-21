package main

import (
  "errors"
  "fmt"
  "strings"
)

var mTimeZoneClassStringToCode = map[string]TimeZoneClass{
  "Unknown"       : TimeZoneClass(0).Unknown(),
  "Canonical"     : TimeZoneClass(0).Canonical(),
  "Alias"         : TimeZoneClass(0).Alias(),
  "SubGroup"      : TimeZoneClass(0).SubGroup(),
  "Artificial"      : TimeZoneClass(0).Artificial(),
}

var mTimeZoneClassLwrCaseStringToCode = map[string]TimeZoneClass{
  "unknown"       : TimeZoneClass(0).Unknown(),
  "canonical"     : TimeZoneClass(0).Canonical(),
  "alias"         : TimeZoneClass(0).Alias(),
  "subgroup"      : TimeZoneClass(0).SubGroup(),
  "artificial"      : TimeZoneClass(0).Artificial(),
}

var mTimeZoneClassToString = map[TimeZoneClass]string{
  TimeZoneClass(0).Unknown()        : "Unknown",
  TimeZoneClass(0).Canonical()      : "Canonical",
  TimeZoneClass(0).Alias()          : "Alias",
  TimeZoneClass(0).SubGroup()       : "SubGroup",
  TimeZoneClass(0).Artificial()       : "Artificial",
}


type TimeZoneClass int

// Unknown - 0 = Unknown or uninitialized Time Zone
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzClass TimeZoneClass) Unknown() TimeZoneClass { return TimeZoneClass(0)}

// Canonical - 1 = Canonical  Standard IANA Current Valid Time Zone
func (tzClass TimeZoneClass) Canonical() TimeZoneClass { return TimeZoneClass(1)}

// Alias - 2 = Outdated Time Zones which references to a current Valid IANA Time Zone
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzClass TimeZoneClass) Alias() TimeZoneClass { return TimeZoneClass(2)}

// SubGroup() = This Time Zone is a place holder referencing a group of valid 
// time zones.
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzClass TimeZoneClass) SubGroup() TimeZoneClass {return TimeZoneClass(3)}

// Artificial() = This Time Zone is a valid time that is not part of the standard
// IANA collection of time zones. One example is Military Time Zones which are
// valid but not included in the IANA collection of time zones.
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzClass TimeZoneClass) Artificial() TimeZoneClass {return TimeZoneClass(4)}

// =============================================================================
// Utility Methods
// =============================================================================

// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of TimeZoneClass is returned set to
// the value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case sensitive and will require an
//	                       exact match. Therefore, 'valid' will NOT
//	                       match the enumeration name, 'Valid'.
//
//	                       If 'false' a case insensitive search is
//	                       conducted for the enumeration name. In
//	                       this case, 'valid' will match the
//	                       enumeration name 'Valid'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	TimeZoneClass           - Upon successful completion, this method will return a new
//	                          instance of TimeZoneClass set to the value of the
//	                          enumeration matched by the string search performed on
//	                          input parameter,'valueString'.
//
//	error                   - If this method completes successfully, the returned error
//	                          Type is set equal to 'nil'. If an error condition is encountered,
//	                          this method will return an error Type which encapsulates an
//	                          appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//  t, err := TimeZoneClass(0).ParseString("Canonical", true)
//                            OR
//  t, err := PathValidityStatusCode(0).ParseString("Canonical()", true)
//                            OR
//  t, err := PathValidityStatusCode(0).ParseString("canonical", false)
//
//  For all of the cases shown above,
//  t is now equal to TimeZoneClass(0).Canonical()
//
func (tzClass TimeZoneClass) ParseString(
  valueString string,
  caseSensitive bool) (TimeZoneClass, error) {

  ePrefix := "TimeZoneClass.ParseString() "

  lenValueStr := len(valueString)

  if strings.HasSuffix(valueString, "()") {
    valueString = valueString[0 : lenValueStr-2]
    lenValueStr -= 2
  }

  if lenValueStr < 3 {
    return TimeZoneClass(0).Unknown(),
      fmt.Errorf(ePrefix+
        "Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
        "valueString='%v'\n", valueString)
  }

  var ok bool

  var tzClassCode TimeZoneClass

  if caseSensitive {

    tzClassCode, ok = mTimeZoneClassStringToCode[valueString]

    if !ok {
      return TimeZoneClass(0).Unknown(),
        errors.New(ePrefix + "Invalid Permission Code!")
    }

  } else {

    valueString = strings.ToLower(valueString)

    tzClassCode, ok = mTimeZoneClassLwrCaseStringToCode[valueString]

    if !ok {
      return TimeZoneClass(0).Unknown(),
        errors.New(ePrefix + "Invalid Permission Code!")
    }

  }

  return tzClassCode, nil
}


// StatusIsValid - If the value of the current TimeZoneClass instance
// is 'invalid', this method will return an error.
//
// If the TimeZoneClass is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (tzClass TimeZoneClass) StatusIsValid() error {

  _, ok := mTimeZoneClassToString[tzClass]

  if !ok {
    ePrefix := "TimeZoneClass.StatusIsValid()\n"
    return fmt.Errorf(ePrefix+
      "Error: The current TimeZoneClass is INVALID! "+
      "TimeZoneClass Value='%v'", int(tzClass))
  }

  return nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneClass'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneClass value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneClass(0).Canonical()
//	str := t.String()
//	    str is now equal to "Canonical"
//
func (tzClass TimeZoneClass) String() string {

  label, ok := mTimeZoneClassToString[tzClass]

  if !ok {
    return ""
  }

  return label
}

// Value - Returns the value of the PathFileTypeCode instance
// as type PathFileTypeCode.
//
func (tzClass TimeZoneClass) Value() TimeZoneClass {

  return tzClass
}

// TZClassCode - public global variable of
// type TimeZoneClass.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneClass
// values.
//
// Usage:
//  TZClassCode.Unknown()
//  TZClassCode.Canonical()
//  TZClassCode.Alias()
//
type TZClassCode TimeZoneClass


var mTimeZoneDeprecationStatusStringToCode = map[string]TimeZoneDeprecationStatus{
  "Unknown"       : TimeZoneDeprecationStatus(0).Unknown(),
  "Deprecated"    : TimeZoneDeprecationStatus(0).Deprecated(),
  "Alias"         : TimeZoneDeprecationStatus(0).Alias(),
  "Valid"         : TimeZoneDeprecationStatus(0).Valid() }

var mTimeZoneDeprecationStatusLwrCaseStringToCode = map[string]TimeZoneDeprecationStatus{
  "unknown"       : TimeZoneDeprecationStatus(0).Unknown(),
  "deprecated"    : TimeZoneDeprecationStatus(0).Deprecated(),
  "alias"         : TimeZoneDeprecationStatus(0).Alias(),
  "valid"         : TimeZoneDeprecationStatus(0).Valid() }

var mTimeZoneDeprecationStatusToString = map[TimeZoneDeprecationStatus]string{
  TimeZoneDeprecationStatus(0).Unknown()        : "Unknown",
  TimeZoneDeprecationStatus(0).Deprecated()     : "Deprecated",
  TimeZoneDeprecationStatus(0).Alias()          : "Alias",
  TimeZoneDeprecationStatus(0).Valid()          : "Valid" }


type TimeZoneDeprecationStatus int

// Unknown - 0 = Unknown or uninitialized Time Zone
//
// This method is part of the standard TimeZoneDeprecationStatus enumeration.
//
func (depStatus TimeZoneDeprecationStatus) Unknown() TimeZoneDeprecationStatus {
  return TimeZoneDeprecationStatus(0)
}

// Deprecated - 1 = Time Zone is outdated and no longer supported. Time Zone
// may be dropped at some point in the future. 'Deprecated' Time Zones are
// linked to valid current Time Zones.
//
// This method is part of the standard TimeZoneDeprecationStatus enumeration.
//
func (depStatus TimeZoneDeprecationStatus) Deprecated() TimeZoneDeprecationStatus {
  return TimeZoneDeprecationStatus(1)
}

// Alias - 2 = Outdated Time Zone which is NOT deprecated and is still supported
// to a limited agree. However, 'Alias' time zones are linked to valid current
// time zones.
//
// In terms of status ranking, 'Alias is between 'Valid' and 'Deprecated'.
//
// This method is part of the standard TimeZoneDeprecationStatus enumeration.
//
func (depStatus TimeZoneDeprecationStatus) Alias() TimeZoneDeprecationStatus { 
  return TimeZoneDeprecationStatus(2)
}

// Valid() = This Time Zone is NOT deprecated and as such, is a current
// valid time zone.
//
// This method is part of the standard TimeZoneDeprecationStatus enumeration.
//
func (depStatus TimeZoneDeprecationStatus) Valid() TimeZoneDeprecationStatus {
  return TimeZoneDeprecationStatus(3)
}

// =============================================================================
// Utility Methods
// =============================================================================

// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of TimeZoneDeprecationStatus is returned set to
// the value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case sensitive and will require an
//	                       exact match. Therefore, 'valid' will NOT
//	                       match the enumeration name, 'Valid'.
//
//	                       If 'false' a case insensitive search is
//	                       conducted for the enumeration name. In
//	                       this case, 'valid' will match the
//	                       enumeration name 'Valid'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	TimeZoneDeprecationStatus           - Upon successful completion, this method will return a new
//	                          instance of TimeZoneDeprecationStatus set to the value of the
//	                          enumeration matched by the string search performed on
//	                          input parameter,'valueString'.
//
//	error                   - If this method completes successfully, the returned error
//	                          Type is set equal to 'nil'. If an error condition is encountered,
//	                          this method will return an error Type which encapsulates an
//	                          appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//  t, err := TimeZoneDeprecationStatus(0).ParseString("Canonical", true)
//                            OR
//  t, err := PathValidityStatusCode(0).ParseString("Canonical()", true)
//                            OR
//  t, err := PathValidityStatusCode(0).ParseString("canonical", false)
//
//  For all of the cases shown above,
//  t is now equal to TimeZoneDeprecationStatus(0).Canonical()
//
func (depStatus TimeZoneDeprecationStatus) ParseString(
  valueString string,
  caseSensitive bool) (TimeZoneDeprecationStatus, error) {

  ePrefix := "TimeZoneDeprecationStatus.ParseString() "

  lenValueStr := len(valueString)

  if strings.HasSuffix(valueString, "()") {
    valueString = valueString[0 : lenValueStr-2]
    lenValueStr -= 2
  }

  if lenValueStr < 3 {
    return TimeZoneDeprecationStatus(0).Unknown(),
      fmt.Errorf(ePrefix+
        "Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
        "valueString='%v'\n", valueString)
  }

  var ok bool

  var depStatusCode TimeZoneDeprecationStatus

  if caseSensitive {

    depStatusCode, ok = mTimeZoneDeprecationStatusStringToCode[valueString]

    if !ok {
      return TimeZoneDeprecationStatus(0).Unknown(),
        errors.New(ePrefix + "Invalid Permission Code!")
    }

  } else {

    valueString = strings.ToLower(valueString)

    depStatusCode, ok = mTimeZoneDeprecationStatusLwrCaseStringToCode[valueString]

    if !ok {
      return TimeZoneDeprecationStatus(0).Unknown(),
        errors.New(ePrefix + "Invalid Permission Code!")
    }

  }

  return depStatusCode, nil
}

// StatusIsValid - If the value of the current TimeZoneDeprecationStatus instance
// is 'invalid', this method will return an error.
//
// If the TimeZoneDeprecationStatus is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (depStatus TimeZoneDeprecationStatus) StatusIsValid() error {

  _, ok := mTimeZoneDeprecationStatusToString[depStatus]

  if !ok {
    ePrefix := "TimeZoneDeprecationStatus.StatusIsValid()\n"
    return fmt.Errorf(ePrefix+
      "Error: The current TimeZoneDeprecationStatus is INVALID! "+
      "TimeZoneDeprecationStatus Value='%v'", int(depStatus))
  }

  return nil
}


// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneDeprecationStatus'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneDeprecationStatus value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneDeprecationStatus(0).Canonical()
//	str := t.String()
//	    str is now equal to "Canonical"
//
func (depStatus TimeZoneDeprecationStatus) String() string {

  label, ok := mTimeZoneDeprecationStatusToString[depStatus]

  if !ok {
    return ""
  }

  return label
}

// Value - Returns the value of the PathFileTypeCode instance
// as type PathFileTypeCode.
//
func (depStatus TimeZoneDeprecationStatus) Value() TimeZoneDeprecationStatus {

  return depStatus
}

// DepStatusCode - public global variable of
// type TimeZoneDeprecationStatus.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneDeprecationStatus
// values.
//
// Usage:
//  DepStatusCode.Unknown()
//  DepStatusCode.Deprecated()
//  DepStatusCode.Alias()
//
type DepStatusCode TimeZoneDeprecationStatus


var mMajorTimeZoneGroupTypeStringToCode = map[string]MajorTimeZoneGroupType{
  "Unknown"       : MajorTimeZoneGroupType(0).Unknown(),
  "IANA"          : MajorTimeZoneGroupType(0).IANA(),
  "Artificial"    : MajorTimeZoneGroupType(0).Artificial(),
  "SubGroup"      : MajorTimeZoneGroupType(0).SubGroup() }

var mMajorTimeZoneGroupTypeLwrCaseStringToCode = map[string]MajorTimeZoneGroupType{
  "unknown"       : MajorTimeZoneGroupType(0).Unknown(),
  "iana"          : MajorTimeZoneGroupType(0).IANA(),
  "artificial"    : MajorTimeZoneGroupType(0).Artificial(),
  "subGroup"      : MajorTimeZoneGroupType(0).SubGroup() }

var mMajorTimeZoneGroupTypeToString = map[MajorTimeZoneGroupType]string{
  MajorTimeZoneGroupType(0).Unknown()         : "Unknown",
  MajorTimeZoneGroupType(0).IANA()            : "IANA",
  MajorTimeZoneGroupType(0).Artificial()      : "Artificial",
  MajorTimeZoneGroupType(0).SubGroup()        : "SubGroup" }


type MajorTimeZoneGroupType int

// Unknown - 0 = Unknown or uninitialized Time Zone
//
// This method is part of the standard MajorTimeZoneGroupType enumeration.
//
func (majgrpstat MajorTimeZoneGroupType) Unknown() MajorTimeZoneGroupType {
  return MajorTimeZoneGroupType(0)
}

// IANA - 1 = This Time Zone Group is a standard collection of IANA Time
// Zones.
//
// This method is part of the standard MajorTimeZoneGroupType enumeration.
//
func (majgrpstat MajorTimeZoneGroupType) IANA() MajorTimeZoneGroupType {
  return MajorTimeZoneGroupType(1)
}

// Artificial - 2 = A Non-IANA Time Zone Group which was constructed to
// add usable time zone information. For example, "Military" Time Zones
// are added but are not recognized by the IANA.
//
// This method is part of the standard MajorTimeZoneGroupType enumeration.
//
func (majgrpstat MajorTimeZoneGroupType) Artificial() MajorTimeZoneGroupType {
  return MajorTimeZoneGroupType(2)
}

// SubGroup() = This Time Zone is NOT an IANA group but represents another
// means of classifying sub-groups of time zones.
//
// This method is part of the standard MajorTimeZoneGroupType enumeration.
//
func (majgrpstat MajorTimeZoneGroupType) SubGroup() MajorTimeZoneGroupType {
  return MajorTimeZoneGroupType(3)
}

// =============================================================================
// Utility Methods
// =============================================================================

// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of MajorTimeZoneGroupType is returned set to
// the value of the associated enumeration.
//
// This is a standard utility method and is not part of the subGroup
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case sensitive and will require an
//	                       exact match. Therefore, 'subGroup' will NOT
//	                       match the enumeration name, 'SubGroup'.
//
//	                       If 'false' a case insensitive search is
//	                       conducted for the enumeration name. In
//	                       this case, 'subGroup' will match the
//	                       enumeration name 'SubGroup'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	MajorTimeZoneGroupType           - Upon successful completion, this method will return a new
//	                          instance of MajorTimeZoneGroupType set to the value of the
//	                          enumeration matched by the string search performed on
//	                          input parameter,'valueString'.
//
//	error                   - If this method completes successfully, the returned error
//	                          Type is set equal to 'nil'. If an error condition is encountered,
//	                          this method will return an error Type which encapsulates an
//	                          appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//  t, err := MajorTimeZoneGroupType(0).ParseString("Canonical", true)
//                            OR
//  t, err := PathSubGroupityStatusCode(0).ParseString("Canonical()", true)
//                            OR
//  t, err := PathSubGroupityStatusCode(0).ParseString("canonical", false)
//
//  For all of the cases shown above,
//  t is now equal to MajorTimeZoneGroupType(0).Canonical()
//
func (majgrpstat MajorTimeZoneGroupType) ParseString(
  valueString string,
  caseSensitive bool) (MajorTimeZoneGroupType, error) {

  ePrefix := "MajorTimeZoneGroupType.ParseString() "

  lenValueStr := len(valueString)

  if strings.HasSuffix(valueString, "()") {
    valueString = valueString[0 : lenValueStr-2]
    lenValueStr -= 2
  }

  if lenValueStr < 3 {
    return MajorTimeZoneGroupType(0).Unknown(),
      fmt.Errorf(ePrefix+
        "Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
        "valueString='%v'\n", valueString)
  }

  var ok bool

  var majorTZGroupType MajorTimeZoneGroupType

  if caseSensitive {

    majorTZGroupType, ok = mMajorTimeZoneGroupTypeStringToCode[valueString]

    if !ok {
      return MajorTimeZoneGroupType(0).Unknown(),
        errors.New(ePrefix + "InsubGroup Permission Code!")
    }

  } else {

    valueString = strings.ToLower(valueString)

    majorTZGroupType, ok = mMajorTimeZoneGroupTypeLwrCaseStringToCode[valueString]

    if !ok {
      return MajorTimeZoneGroupType(0).Unknown(),
        errors.New(ePrefix + "InsubGroup Permission Code!")
    }

  }

  return majorTZGroupType, nil
}

// StatusIsValid - If the value of the current MajorTimeZoneGroupType instance
// is 'invalid', this method will return an error.
//
// If the MajorTimeZoneGroupType is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the subGroup enumerations
// for this type.
//
func (majgrpstat MajorTimeZoneGroupType) StatusIsValid() error {

  _, ok := mMajorTimeZoneGroupTypeToString[majgrpstat]

  if !ok {
    ePrefix := "MajorTimeZoneGroupType.StatusIsValid()\n"
    return fmt.Errorf(ePrefix+
      "Error: The current MajorTimeZoneGroupType is INVALID! "+
      "MajorTimeZoneGroupType Value='%v'", int(majgrpstat))
  }

  return nil
}


// String - Returns a string with the name of the enumeration associated
// with this instance of 'MajorTimeZoneGroupType'.
//
// This is a standard utility method and is not part of the subGroup enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the MajorTimeZoneGroupType value is insubGroup, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= MajorTimeZoneGroupType(0).Canonical()
//	str := t.String()
//	    str is now equal to "Canonical"
//
func (majgrpstat MajorTimeZoneGroupType) String() string {

  label, ok := mMajorTimeZoneGroupTypeToString[majgrpstat]

  if !ok {
    return ""
  }

  return label
}

// Value - Returns the value of the PathFileTypeCode instance
// as type PathFileTypeCode.
//
func (majgrpstat MajorTimeZoneGroupType) Value() MajorTimeZoneGroupType {

  return majgrpstat
}

// DepStatusCode - public global variable of
// type MajorTimeZoneGroupType.
//
// This variable serves as an easier, short hand
// technique for accessing MajorTimeZoneGroupType
// values.
//
// Usage:
//  DepStatusCode.Unknown()
//  DepStatusCode.IANA()
//  DepStatusCode.Artifical()
//
type MajorTZGroup MajorTimeZoneGroupType

