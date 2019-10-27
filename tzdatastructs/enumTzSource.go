package tzdatastructs

import (
	"errors"
	"fmt"
	"strings"
)

var mTimeZoneSourceStringToCode = map[string]TimeZoneSource{
	"None"         : TimeZoneSource(0).None(),
	"Iana"         : TimeZoneSource(0).Iana(),
	"Military"     : TimeZoneSource(0).Military(),
	"Other"        : TimeZoneSource(0).Other(),
}

var mTimeZoneSourceLwrCaseStringToCode = map[string]TimeZoneSource{
	"none"         : TimeZoneSource(0).None(),
	"iana"         : TimeZoneSource(0).Iana(),
	"military"     : TimeZoneSource(0).Military(),
	"other"        : TimeZoneSource(0).Other(),
}

var mTimeZoneSourceToString = map[TimeZoneSource]string{
	TimeZoneSource(0).None()     : "None",
	TimeZoneSource(0).Iana()     : "Iana",
	TimeZoneSource(0).Military() : "Military",
	TimeZoneSource(0).Other()    : "Other",
}


type TimeZoneSource int

// None - 0 = Signifies that this entity is not a Time Zone.
//
// This method is part of the standard TimeZoneSource enumeration.
//
func (tzSrc TimeZoneSource) None() TimeZoneSource { return TimeZoneSource(0)}

// TimeZone - 0 = Signifies that this entity is sourced from the IANA 
// time zone database.
//
// This method is part of the standard TimeZoneSource enumeration.
//
func (tzSrc TimeZoneSource) Iana() TimeZoneSource { return TimeZoneSource(1)}

// Military - Signifies this time zone is used by the Military.
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzSrc TimeZoneSource) Military() TimeZoneSource { return TimeZoneSource(2)}

// Military - Signifies this time zone is artificial and provided for
// convenience.
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzSrc TimeZoneSource) Other() TimeZoneSource { return TimeZoneSource(3)}


// =============================================================================
// Utility Methods
// =============================================================================


// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneSource'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneSource value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneSource(0).Iana()
//	str := t.String()
//	    str is now equal to "Iana"
//
func (tzSrc TimeZoneSource) String() string {

	label, ok := mTimeZoneSourceToString[tzSrc]

	if !ok {
		return ""
	}

	return label
}


// UtilityIsValid - If the value of the current TimeZoneSource instance
// is 'invalid', this method will return an error.
//
// If the TimeZoneSource is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (tzSrc TimeZoneSource) UtilityIsValid() error {

	_, ok := mTimeZoneSourceToString[tzSrc]

	if !ok {
		ePrefix := "TimeZoneSource.UtilityIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current TimeZoneSource is INVALID! "+
			"TimeZoneSource Value='%v'", int(tzSrc))
	}

	return nil
}


// UtilityParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of TimeZoneSource is returned set to
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
//	TimeZoneSource        - Upon successful completion, this method will return a new
//	                          instance of TimeZoneSource set to the value of the
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
//  t, err := TimeZoneSource(0).UtilityParseString("TimeZone", true)
//                            OR
//  t, err := TimeZoneSource(0).UtilityParseString("TimeZone()", true)
//                            OR
//  t, err := TimeZoneSource(0).UtilityParseString("timezone", false)
//
//  For all of the cases shown above,
//  t is now equal to TimeZoneSource(0).TimeZone()
//
func (tzClass TimeZoneSource) UtilityParseString(
	valueString string,
	caseSensitive bool) (TimeZoneSource, error) {

	ePrefix := "TimeZoneSource.UtilityParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 4 {
		return TimeZoneSource(0).None(),
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! Length Less than 4-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var tzClassCode TimeZoneSource

	if caseSensitive {

		tzClassCode, ok = mTimeZoneSourceStringToCode[valueString]

		if !ok {
			return TimeZoneSource(0).None(),
				errors.New(ePrefix + "Invalid Time Zone Class Code!")
		}

	} else {

		valueString = strings.ToLower(valueString)

		tzClassCode, ok = mTimeZoneSourceLwrCaseStringToCode[valueString]

		if !ok {
			return TimeZoneSource(0).None(),
				errors.New(ePrefix + "Invalid Permission Code!")
		}

	}

	return tzClassCode, nil
}


// UtilityValue - Returns the value of the TimeZoneSource instance
// as type TimeZoneSource.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (tzSrc TimeZoneSource) UtilityValue() TimeZoneSource {

	return tzSrc
}


// TZSrc - public global variable of
// type TimeZoneClass.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneSource
// values.
//
// Usage:
//  TZSrc.None()
//  TZSrc.Iana()
//  TZSrc.Military()
//  TZSrc.Other()
//
var TZSrc TimeZoneSource
