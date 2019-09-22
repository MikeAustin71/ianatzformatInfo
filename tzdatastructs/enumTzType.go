package tzdatastructs

import (
	"errors"
	"fmt"
	"strings"
)

var mTimeZoneTypeStringToCode = map[string]TimeZoneType{
	"Unknown"          : TimeZoneType(0).Unknown(),
	"Standard"         : TimeZoneType(0).Standard(),
	"Group"            : TimeZoneType(0).Group(),
	"SubZone"          : TimeZoneType(0).SubZone(),
}

var mTimeZoneTypeLwrCaseStringToCode = map[string]TimeZoneType{
	"unknown"          : TimeZoneType(0).Unknown(),
	"standard"         : TimeZoneType(0).Standard(),
	"group"            : TimeZoneType(0).Group(),
	"subzone"          : TimeZoneType(0).SubZone(),
}

var mTimeZoneTypeToString = map[TimeZoneType]string{
	TimeZoneType(0).Unknown()       : "Unknown",
	TimeZoneType(0).Standard()      : "Standard",
	TimeZoneType(0).Group()         : "Group",
	TimeZoneType(0).SubZone()       : "SubZone",
}


type TimeZoneType int


// Unknown - 0 = Unknown or uninitialized Time Zone
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzType TimeZoneType) Unknown() TimeZoneType { return TimeZoneType(0)}


// Standard - 1 = A standard time zone
//
func (tzType TimeZoneType) Standard() TimeZoneType { return TimeZoneType(1)}

// Group - 2 = A place holder time zone which defines a group of 
// subsidiary time zones.
//
func (tzType TimeZoneType) Group() TimeZoneType { return TimeZoneType(2)}


// SubTimeZone - A time zone which is part of a sub-group declared by a Place
// Holder Time Zone. 
//
func (tzType TimeZoneType) SubZone() TimeZoneType { return TimeZoneType(3)}

// =============================================================================
// Utility Methods
// =============================================================================

// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of TimeZoneType is returned set to
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
//	TimeZoneType           - Upon successful completion, this method will return a new
//	                          instance of TimeZoneType set to the value of the
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
//  t, err := TimeZoneType(0).ParseString("Canonical", true)
//                            OR
//  t, err := TimeZoneType(0).ParseString("Canonical()", true)
//                            OR
//  t, err := TimeZoneType(0).ParseString("canonical", false)
//
//  For all of the cases shown above,
//  t is now equal to TimeZoneType(0).Canonical()
//
func (tzType TimeZoneType) ParseString(
	valueString string,
	caseSensitive bool) (TimeZoneType, error) {

	ePrefix := "TimeZoneType.ParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 3 {
		return TimeZoneType(0).Unknown(),
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var tzTypeCode TimeZoneType

	if caseSensitive {

		tzTypeCode, ok = mTimeZoneTypeStringToCode[valueString]

		if !ok {
			return TimeZoneType(0).Unknown(),
				errors.New(ePrefix + "Invalid Type Zone Type Code!")
		}

	} else {

		valueString = strings.ToLower(valueString)

		tzTypeCode, ok = mTimeZoneTypeLwrCaseStringToCode[valueString]

		if !ok {
			return TimeZoneType(0).Unknown(),
				errors.New(ePrefix + "Invalid Permission Code!")
		}

	}

	return tzTypeCode, nil
}


// TypeIsValid - If the value of the current TimeZoneType instance
// is 'invalid', this method will return an error.
//
// If the TimeZoneType is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (tzType TimeZoneType) TypeIsValid() error {

	_, ok := mTimeZoneTypeToString[tzType]

	if !ok {
		ePrefix := "TimeZoneType.TypeIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current TimeZoneType is INVALID! "+
			"TimeZoneType StatusValue='%v'", int(tzType))
	}

	return nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneType'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneType value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneType(0).Standard()
//	str := t.String()
//	    str is now equal to "Standard"
//
func (tzType TimeZoneType) String() string {

	label, ok := mTimeZoneTypeToString[tzType]

	if !ok {
		return ""
	}

	return label
}

// TypeValue - Returns the value of the TimeZoneType instance
// as type TimeZoneType.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (tzType TimeZoneType) TypeValue() TimeZoneType {

	return tzType
}

// TZType - public global variable of
// type TimeZoneType.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneType
// values.
//
// Usage:
//  TZType.Unknown()
//  TZType.Standard()
//  TZType.Group()
//  TZType.SubZone()
//
var TZType TimeZoneType
