package tzdatastructs

import (
	"errors"
	"fmt"
	"strings"
)

var mTimeZoneGroupTypeStringToCode = map[string]TimeZoneGroupType{
	"Unknown"       : TimeZoneGroupType(0).Unknown(),
	"Standard"      : TimeZoneGroupType(0).Standard(),
	"SubGroup"      : TimeZoneGroupType(0).SubGroup() }

var mTimeZoneGroupTypeLwrCaseStringToCode = map[string]TimeZoneGroupType{
	"unknown"       : TimeZoneGroupType(0).Unknown(),
	"standard"      : TimeZoneGroupType(0).Standard(),
	"subgroup"      : TimeZoneGroupType(0).SubGroup() }

var mTimeZoneGroupTypeToString = map[TimeZoneGroupType]string{
	TimeZoneGroupType(0).Unknown()         : "Unknown",
	TimeZoneGroupType(0).Standard()        : "Standard",
	TimeZoneGroupType(0).SubGroup()        : "SubGroup" }


type TimeZoneGroupType int

// Unknown - 0 = Unknown or uninitialized Time Zone
//
// This method is part of the defined TimeZoneGroupType enumeration.
//
func (grpType TimeZoneGroupType) Unknown() TimeZoneGroupType {
	return TimeZoneGroupType(0)
}

// Standard - 1 = This is a Standard Time Zone Group.
//
// This method is part of the defined TimeZoneGroupType enumeration.
//
func (grpType TimeZoneGroupType) Standard() TimeZoneGroupType {
	return TimeZoneGroupType(1)
}

// SubGroup() = This Time Zone Group is a subsidiary group.
//
// This method is part of the defined TimeZoneGroupType enumeration.
//
func (grpType TimeZoneGroupType) SubGroup() TimeZoneGroupType {
	return TimeZoneGroupType(2)
}

// =============================================================================
// Utility Methods
// =============================================================================

// ParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of TimeZoneGroupType is returned set to
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
//  t, err := TimeZoneGroupType(0).ParseString("Standard", true)
//                            OR
//  t, err := TimeZoneGroupType(0).ParseString("Standard()", true)
//                            OR
//  t, err :=  TimeZoneGroupType(0).ParseString("standard", false)
//
//  For all of the cases shown above,
//  t is now equal to TimeZoneGroupType(0).Standard()
//
func (grpType TimeZoneGroupType) ParseString(
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

		majorTZGroupType, ok = mTimeZoneGroupTypeStringToCode[valueString]

		if !ok {
			return TimeZoneGroupType(0).Unknown(),
				errors.New(ePrefix + "Invalid Time Zone Group Type Code!")
		}

	} else {

		valueString = strings.ToLower(valueString)

		majorTZGroupType, ok = mTimeZoneGroupTypeLwrCaseStringToCode[valueString]

		if !ok {
			return TimeZoneGroupType(0).Unknown(),
				errors.New(ePrefix + "Invalid Time Zone Group Type Code!")
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
// This is a standard utility method and is not part of the defined enumerations
// for this type.
//
func (grpType TimeZoneGroupType) TypeIsValid() error {

	_, ok := mTimeZoneGroupTypeToString[grpType]

	if !ok {
		ePrefix := "TimeZoneGroupType.TypeIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current Time Zone Group Type is INVALID! "+
			"TimeZoneGroupType TypeValue='%v'", int(grpType))
	}

	return nil
}


// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneGroupType'.
//
// This is a standard utility method and is not part of the defined enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return StatusValue:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneGroupType value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneGroupType(0).Standard()
//	str := t.String()
//	    str is now equal to "Standard"
//
func (grpType TimeZoneGroupType) String() string {

	label, ok := mTimeZoneGroupTypeToString[grpType]

	if !ok {
		return ""
	}

	return label
}

// TypeValue - Returns the value of the PathFileTypeCode instance
// as type PathFileTypeCode.
//
// This is a standard utility method and is not part of the defined
// enumerations for this type.
//
func (grpType TimeZoneGroupType) TypeValue() TimeZoneGroupType {

	return grpType
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
//  TzGrpType.Artificial()
//  TzGrpType.SubGroup()
//
var TzGrpType TimeZoneGroupType

