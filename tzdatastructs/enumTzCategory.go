package tzdatastructs

import (
	"errors"
	"fmt"
	"strings"
)

var mTimeZoneCategoryStringToCode = map[string]TimeZoneCategory{
	"None"             : TimeZoneCategory(0).None(),
	"TimeZone"         : TimeZoneCategory(0).TimeZone(),
	"LinkZone"         : TimeZoneCategory(0).LinkZone(),
}

var mTimeZoneCategoryLwrCaseStringToCode = map[string]TimeZoneCategory{
	"none"             : TimeZoneCategory(0).None(),
	"timeZone"         : TimeZoneCategory(0).TimeZone(),
	"linkZone"         : TimeZoneCategory(0).LinkZone(),
}

var mTimeZoneCategoryToString = map[TimeZoneCategory]string{
	TimeZoneCategory(0).None()     : "None",
	TimeZoneCategory(0).TimeZone() : "TimeZone",
	TimeZoneCategory(0).LinkZone() : "LinkZone",
}


type TimeZoneCategory int

// None - Signifies that this entity is not a Time Zone and not a Link Zone.
//
// This method is part of the standard TimeZoneCategory enumeration.
//
func (tzCat TimeZoneCategory) None() TimeZoneCategory { return TimeZoneCategory(0)}

// TimeZone -Signifies a Standard Time Zone
//
// This method is part of the standard TimeZoneCategory enumeration.
//
func (tzCat TimeZoneCategory) TimeZone() TimeZoneCategory { return TimeZoneCategory(1)}

// Link Zone - Signifies a deprecated time zone which is linked or mapped to
// a valid time zone
//
// This method is part of the standard TimeZoneClass enumeration.
//
func (tzCat TimeZoneCategory) LinkZone() TimeZoneCategory { return TimeZoneCategory(2)}

// =============================================================================
// Utility Methods
// =============================================================================


// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneCategory'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneCategory value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneCategory(0).TimeZone()
//	str := t.String()
//	    str is now equal to "TimeZone"
//
func (tzCat TimeZoneCategory) String() string {

	label, ok := mTimeZoneCategoryToString[tzCat]

	if !ok {
		return ""
	}

	return label
}


// UtilityIsValid - If the value of the current TimeZoneCategory instance
// is 'invalid', this method will return an error.
//
// If the TimeZoneCategory is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (tzCat TimeZoneCategory) UtilityIsValid() error {

	_, ok := mTimeZoneCategoryToString[tzCat]

	if !ok {
		ePrefix := "TimeZoneCategory.UtilityIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current TimeZoneCategory is INVALID! "+
			"TimeZoneCategory Value='%v'", int(tzCat))
	}

	return nil
}


// UtilityParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of TimeZoneCategory is returned set to
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
//	TimeZoneCategory        - Upon successful completion, this method will return a new
//	                          instance of TimeZoneCategory set to the value of the
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
//  t, err := TimeZoneCategory(0).UtilityParseString("TimeZone", true)
//                            OR
//  t, err := TimeZoneCategory(0).UtilityParseString("TimeZone()", true)
//                            OR
//  t, err := TimeZoneCategory(0).UtilityParseString("timezone", false)
//
//  For all of the cases shown above,
//  t is now equal to TimeZoneCategory(0).TimeZone()
//
func (tzClass TimeZoneCategory) UtilityParseString(
	valueString string,
	caseSensitive bool) (TimeZoneCategory, error) {

	ePrefix := "TimeZoneCategory.UtilityParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 4 {
		return TimeZoneCategory(0).None(),
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! Length Less than 4-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var tzClassCode TimeZoneCategory

	if caseSensitive {

		tzClassCode, ok = mTimeZoneCategoryStringToCode[valueString]

		if !ok {
			return TimeZoneCategory(0).None(),
				errors.New(ePrefix + "Invalid Time Zone Class Code!")
		}

	} else {

		valueString = strings.ToLower(valueString)

		tzClassCode, ok = mTimeZoneCategoryLwrCaseStringToCode[valueString]

		if !ok {
			return TimeZoneCategory(0).None(),
				errors.New(ePrefix + "Invalid Permission Code!")
		}

	}

	return tzClassCode, nil
}


// UtilityValue - Returns the value of the TimeZoneCategory instance
// as type TimeZoneCategory.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (tzCat TimeZoneCategory) UtilityValue() TimeZoneCategory {

	return tzCat
}


// TZCat - public global variable of
// type TimeZoneClass.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneCategory
// values.
//
// Usage:
//  TZCat.None()
//  TZCat.TimeZone()
//  TZCat.LinkZone()
//
var TZCat TimeZoneCategory

