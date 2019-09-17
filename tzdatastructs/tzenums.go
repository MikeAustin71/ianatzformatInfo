package tzdatastructs

import (
	"errors"
	"fmt"
	"strings"
)

var mTimeZoneClassStringToCode = map[string]TimeZoneClass{
	"Unknown"             : TimeZoneClass(0).Unknown(),
	"Canonical"           : TimeZoneClass(0).Canonical(),
	"Alias"               : TimeZoneClass(0).Alias(),
	"SubTimeZone"         : TimeZoneClass(0).SubTimeZone(),
	"SubGroup"            : TimeZoneClass(0).SubGroup(),
	"Artificial"          : TimeZoneClass(0).Artificial(),
}

var mTimeZoneClassLwrCaseStringToCode = map[string]TimeZoneClass{
	"unknown"             : TimeZoneClass(0).Unknown(),
	"canonical"           : TimeZoneClass(0).Canonical(),
	"alias"               : TimeZoneClass(0).Alias(),
	"subtimezone"         : TimeZoneClass(0).SubTimeZone(),
	"subgroup"            : TimeZoneClass(0).SubGroup(),
	"artificial"          : TimeZoneClass(0).Artificial(),
}

var mTimeZoneClassToString = map[TimeZoneClass]string{
	TimeZoneClass(0).Unknown()          : "Unknown",
	TimeZoneClass(0).Canonical()        : "Canonical",
	TimeZoneClass(0).Alias()            : "Alias",
	TimeZoneClass(0).SubTimeZone()      : "SubTimeZone",
	TimeZoneClass(0).SubGroup()         : "SubGroup",
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

// SubTimeZone() = This Time Zone is a subsidiary of a SubGroup.
// Example: 'America/Argentina/Buenos_Aires' is a Sub-Time Zone.
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzClass TimeZoneClass) SubTimeZone() TimeZoneClass {return TimeZoneClass(3)}

// SubGroup() = This Time Zone element is a place holder referencing a group containing
// multiple valid time zones.
//
// Example: 'America/Argentina/Buenos_Aires' is a 'Sub-Time Zone' which would have
// a time zone Sub-Group place holder named, 'America/Argentina'. This Sub-Group
// would reference all instances of sub-groups classified under 'America/Argentina'
// such as 'America/Argentina/Buenos_Aires', 'America/Argentina/Catamarca' and
// 'America/Argentina/Cordoba'.
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzClass TimeZoneClass) SubGroup() TimeZoneClass {return TimeZoneClass(4)}



// Artificial() = This Time Zone is a valid time that is not part of the standard
// IANA collection of time zones. One example is Military Time Zones which are
// valid but not included in the IANA collection of time zones.
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzClass TimeZoneClass) Artificial() TimeZoneClass {return TimeZoneClass(5)}

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
//  t, err := TimeZoneClass(0).ParseString("Canonical()", true)
//                            OR
//  t, err := TimeZoneClass(0).ParseString("canonical", false)
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


// ClassIsValid - If the value of the current TimeZoneClass instance
// is 'invalid', this method will return an error.
//
// If the TimeZoneClass is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (tzClass TimeZoneClass) ClassIsValid() error {

	_, ok := mTimeZoneClassToString[tzClass]

	if !ok {
		ePrefix := "TimeZoneClass.TypeIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current TimeZoneClass is INVALID! "+
			"TimeZoneClass StatusValue='%v'", int(tzClass))
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
// Return StatusValue:
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

// StatusValue - Returns the value of the TimeZoneClass instance
// as type TimeZoneClass.
//
func (tzClass TimeZoneClass) ClassValue() TimeZoneClass {

	return tzClass
}

// TZClass - public global variable of
// type TimeZoneClass.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneClass
// values.
//
// Usage:
//  TZClass.Unknown()
//  TZClass.Canonical()
//  TZClass.Alias()
//
var TZClass TimeZoneClass


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

// TypeIsValid - If the value of the current TimeZoneDeprecationStatus instance
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
		ePrefix := "TimeZoneDeprecationStatus.TypeIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current TimeZoneDeprecationStatus is INVALID! "+
			"TimeZoneDeprecationStatus StatusValue='%v'", int(depStatus))
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
// Return StatusValue:
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

// StatusValue - Returns the value of the PathFileTypeCode instance
// as type PathFileTypeCode.
//
func (depStatus TimeZoneDeprecationStatus) StatusValue() TimeZoneDeprecationStatus {

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
//  DepStatusCode.Valid()
//
var DepStatusCode TimeZoneDeprecationStatus


var mMajorTimeZoneGroupTypeStringToCode = map[string]TimeZoneGroupType{
	"Unknown"       : TimeZoneGroupType(0).Unknown(),
	"IANA"          : TimeZoneGroupType(0).IANA(),
	"Artificial"    : TimeZoneGroupType(0).Artificial(),
	"SubGroup"      : TimeZoneGroupType(0).SubGroup() }

var mMajorTimeZoneGroupTypeLwrCaseStringToCode = map[string]TimeZoneGroupType{
	"unknown"       : TimeZoneGroupType(0).Unknown(),
	"iana"          : TimeZoneGroupType(0).IANA(),
	"artificial"    : TimeZoneGroupType(0).Artificial(),
	"subGroup"      : TimeZoneGroupType(0).SubGroup() }

var mMajorTimeZoneGroupTypeToString = map[TimeZoneGroupType]string{
	TimeZoneGroupType(0).Unknown()         : "Unknown",
	TimeZoneGroupType(0).IANA()            : "IANA",
	TimeZoneGroupType(0).Artificial()      : "Artificial",
	TimeZoneGroupType(0).SubGroup()        : "SubGroup" }


type TimeZoneGroupType int

// Unknown - 0 = Unknown or uninitialized Time Zone
//
// This method is part of the standard TimeZoneGroupType enumeration.
//
func (majgrpstat TimeZoneGroupType) Unknown() TimeZoneGroupType {
	return TimeZoneGroupType(0)
}

// IANA - 1 = This Time Zone Group is a standard collection of IANA Time
// Zones.
//
// This method is part of the standard TimeZoneGroupType enumeration.
//
func (majgrpstat TimeZoneGroupType) IANA() TimeZoneGroupType {
	return TimeZoneGroupType(1)
}

// Artificial - 2 = A Non-IANA Time Zone Group which was constructed to
// add usable time zone information. For example, "Military" Time Zones
// are added but are not recognized by the IANA.
//
// This method is part of the standard TimeZoneGroupType enumeration.
//
func (majgrpstat TimeZoneGroupType) Artificial() TimeZoneGroupType {
	return TimeZoneGroupType(2)
}

// SubTimeZone() = This Time Zone is NOT an IANA group but represents another
// means of classifying sub-groups of time zones.
//
// This method is part of the standard TimeZoneGroupType enumeration.
//
func (majgrpstat TimeZoneGroupType) SubGroup() TimeZoneGroupType {
	return TimeZoneGroupType(3)
}

// =============================================================================
// Utility Methods
// =============================================================================

// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of TimeZoneGroupType is returned set to
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
//	                       match the enumeration name, 'SubTimeZone'.
//
//	                       If 'false' a case insensitive search is
//	                       conducted for the enumeration name. In
//	                       this case, 'subGroup' will match the
//	                       enumeration name 'SubTimeZone'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	TimeZoneGroupType       - Upon successful completion, this method will return a new
//	                          instance of TimeZoneGroupType set to the value of the
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
//  t, err := TimeZoneGroupType(0).ParseString("Canonical", true)
//                            OR
//  t, err := PathSubGroupityStatusCode(0).ParseString("Canonical()", true)
//                            OR
//  t, err := PathSubGroupityStatusCode(0).ParseString("canonical", false)
//
//  For all of the cases shown above,
//  t is now equal to TimeZoneGroupType(0).Canonical()
//
func (majgrpstat TimeZoneGroupType) ParseString(
	valueString string,
	caseSensitive bool) (TimeZoneGroupType, error) {

	ePrefix := "TimeZoneGroupType.ParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 3 {
		return TimeZoneGroupType(0).Unknown(),
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var majorTZGroupType TimeZoneGroupType

	if caseSensitive {

		majorTZGroupType, ok = mMajorTimeZoneGroupTypeStringToCode[valueString]

		if !ok {
			return TimeZoneGroupType(0).Unknown(),
				errors.New(ePrefix + "InsubGroup Permission Code!")
		}

	} else {

		valueString = strings.ToLower(valueString)

		majorTZGroupType, ok = mMajorTimeZoneGroupTypeLwrCaseStringToCode[valueString]

		if !ok {
			return TimeZoneGroupType(0).Unknown(),
				errors.New(ePrefix + "InsubGroup Permission Code!")
		}

	}

	return majorTZGroupType, nil
}

// TypeIsValid - If the value of the current TimeZoneGroupType instance
// is 'invalid', this method will return an error.
//
// If the TimeZoneGroupType is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the subGroup enumerations
// for this type.
//
func (majgrpstat TimeZoneGroupType) TypeIsValid() error {

	_, ok := mMajorTimeZoneGroupTypeToString[majgrpstat]

	if !ok {
		ePrefix := "TimeZoneGroupType.TypeIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current TimeZoneGroupType is INVALID! "+
			"TimeZoneGroupType StatusValue='%v'", int(majgrpstat))
	}

	return nil
}


// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneGroupType'.
//
// This is a standard utility method and is not part of the subGroup enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return StatusValue:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneGroupType value is insubGroup, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneGroupType(0).Canonical()
//	str := t.String()
//	    str is now equal to "Canonical"
//
func (majgrpstat TimeZoneGroupType) String() string {

	label, ok := mMajorTimeZoneGroupTypeToString[majgrpstat]

	if !ok {
		return ""
	}

	return label
}

// TypeValue - Returns the value of the PathFileTypeCode instance
// as type PathFileTypeCode.
//
func (majgrpstat TimeZoneGroupType) TypeValue() TimeZoneGroupType {

	return majgrpstat
}

// MajorTzGrpType - public global variable of
// type TimeZoneGroupType.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneGroupType
// values.
//
// Usage:
//  TzGrpType.Unknown()
//  TzGrpType.IANA()
//  TzGrpType.Artifical()
//  TzGrpType.SubGroup()
//
var TzGrpType TimeZoneGroupType

