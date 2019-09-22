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
	"Artificial"          : TimeZoneClass(0).Artificial(),
}

var mTimeZoneClassLwrCaseStringToCode = map[string]TimeZoneClass{
	"unknown"             : TimeZoneClass(0).Unknown(),
	"canonical"           : TimeZoneClass(0).Canonical(),
	"alias"               : TimeZoneClass(0).Alias(),
	"artificial"          : TimeZoneClass(0).Artificial(),
}

var mTimeZoneClassToString = map[TimeZoneClass]string{
	TimeZoneClass(0).Unknown()          : "Unknown",
	TimeZoneClass(0).Canonical()        : "Canonical",
	TimeZoneClass(0).Alias()            : "Alias",
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

// Artificial() = This Time Zone is a valid time that is not part of the standard
// IANA collection of time zones. One example is Military Time Zones which are
// valid but not included in the IANA collection of time zones.
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzClass TimeZoneClass) Artificial() TimeZoneClass {return TimeZoneClass(3)}

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
				errors.New(ePrefix + "Invalid Time Zone Class Code!")
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
		ePrefix := "TimeZoneClass.ClassIsValid()\n"
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

// ClassValue - Returns the value of the TimeZoneClass instance
// as type TimeZoneClass.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
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

