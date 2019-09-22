package tzdatastructs

import (
	"errors"
	"fmt"
	"strings"
)


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


