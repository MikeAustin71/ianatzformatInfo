package tzdatastructs

import (
	"errors"
	"fmt"
	"strings"
)

var mTimeZoneGroupClassStringToCode = map[string]TimeZoneGroupClass{
	"Unknown"       : TimeZoneGroupClass(0).Unknown(),
	"IANA"          : TimeZoneGroupClass(0).IANA(),
	"Artificial"    : TimeZoneGroupClass(0).Artificial()}

var mTimeZoneGroupClassLwrCaseStringToCode = map[string]TimeZoneGroupClass{
	"unknown"       : TimeZoneGroupClass(0).Unknown(),
	"iana"          : TimeZoneGroupClass(0).IANA(),
	"artificial"    : TimeZoneGroupClass(0).Artificial()}

var mTimeZoneGroupClassToString = map[TimeZoneGroupClass]string{
	TimeZoneGroupClass(0).Unknown()         : "Unknown",
	TimeZoneGroupClass(0).IANA()            : "IANA",
	TimeZoneGroupClass(0).Artificial()      : "Artificial"}


type TimeZoneGroupClass int

// Unknown - 0 = Unknown or uninitialized Time Zone Group
//
// This method is part of the standard TimeZoneGroupClass enumeration.
//
func (grpClass TimeZoneGroupClass) Unknown() TimeZoneGroupClass {
	return TimeZoneGroupClass(0)
}

// IANA - 1 = This Time Zone Group is a standard collection of IANA Time
// Zones.
//
// This method is part of the standard TimeZoneGroupClass enumeration.
//
func (grpClass TimeZoneGroupClass) IANA() TimeZoneGroupClass {
	return TimeZoneGroupClass(1)
}

// Artificial - 2 = A Non-IANA Time Zone Group which was constructed to
// add usable time zone information. For example, "Military" Time Zones
// are added but are not recognized by the IANA.
//
// This method is part of the standard TimeZoneGroupClass enumeration.
//
func (grpClass TimeZoneGroupClass) Artificial() TimeZoneGroupClass {
	return TimeZoneGroupClass(2)
}

// =============================================================================
// Utility Methods
// =============================================================================

// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of TimeZoneGroupClass is returned set to
// the value of the associated enumeration.
//
// This is a standard utility method and is not part of the defined
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
//	                       exact match. Therefore, 'iana' will NOT
//	                       match the enumeration name, 'IANA'.
//
//	                       If 'false' a case insensitive search is
//	                       conducted for the enumeration name. In
//	                       this case, 'iana' will match the
//	                       enumeration name 'IANA'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	TimeZoneGroupClass       - Upon successful completion, this method will return a new
//	                          instance of TimeZoneGroupClass set to the value of the
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
//  t, err := TimeZoneGroupClass(0).ParseString("IANA", true)
//                            OR
//  t, err := TimeZoneGroupClass(0).ParseString("IANA()", true)
//                            OR
//  t, err := TimeZoneGroupClass(0).ParseString("iana", false)
//
//  For all of the cases shown above,
//  t is now equal to TimeZoneGroupClass(0).IANA()
//
func (grpClass TimeZoneGroupClass) ParseString(
	valueString string,
	caseSensitive bool) (TimeZoneGroupClass, error) {

	ePrefix := "TimeZoneGroupClass.ParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 3 {
		return TimeZoneGroupClass(0).Unknown(),
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var tzGroupClassCode TimeZoneGroupClass

	if caseSensitive {

		tzGroupClassCode, ok = mTimeZoneGroupClassStringToCode[valueString]

		if !ok {
			return TimeZoneGroupClass(0).Unknown(),
				errors.New(ePrefix + "Invalid Time Zone Group Class Code!")
		}

	} else {

		valueString = strings.ToLower(valueString)

		tzGroupClassCode, ok = mTimeZoneGroupClassLwrCaseStringToCode[valueString]

		if !ok {
			return TimeZoneGroupClass(0).Unknown(),
				errors.New(ePrefix + "Invalid Time Zone Group Class Code!")
		}

	}

	return tzGroupClassCode, nil
}

// TypeIsValid - If the value of the current TimeZoneGroupClass instance
// is 'invalid', this method will return an error.
//
// If the TimeZoneGroupClass is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the defined enumerations
// for this type.
//
func (grpClass TimeZoneGroupClass) TypeIsValid() error {

	_, ok := mTimeZoneGroupClassToString[grpClass]

	if !ok {
		ePrefix := "TimeZoneGroupClass.TypeIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current TimeZoneGroupClass is INVALID! "+
			"TimeZoneGroupClass StatusValue='%v'", int(grpClass))
	}

	return nil
}


// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneGroupClass'.
//
// This is a standard utility method and is not part of the defined enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return StatusValue:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneGroupClass value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneGroupClass(0).IANA()
//	str := t.String()
//	    str is now equal to "IANA"
//
func (grpClass TimeZoneGroupClass) String() string {

	label, ok := mTimeZoneGroupClassToString[grpClass]

	if !ok {
		return ""
	}

	return label
}

// TypeValue - Returns the value of the TimeZoneGroupClass instance
// as type TimeZoneGroupClass.
//
func (grpClass TimeZoneGroupClass) TypeValue() TimeZoneGroupClass {

	return grpClass
}

// TzGrpClass - public global variable of
// type TimeZoneGroupClass.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneGroupClass
// values.
//
// Usage:
//  TzGrpType.Unknown()
//  TzGrpType.IANA()
//  TzGrpType.Artificial()
//
var TzGrpClass TimeZoneGroupClass

