package tzdatastructs

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"strconv"
	"strings"
)

type TimeZoneDataDto struct {
	GroupName                 string
	TzName                    string
	TzAliasValue              string
	TzCanonicalValue          string
	TzValue                   string
	TzSortValue               string
	FuncSelfReferenceVariable string
	FuncType                  string
	FuncName                  string
	FuncReturnType            string
	FuncReturnValue           string
	SourceFileNameExt         string

	// 0 = Unknown
	// 1 = Canonical
	// 2 = Alias
	// 3 = Artificial
	TzClass                   TimeZoneClass

	// 0 = Unknown
	// 1 = Standard
	// 2 = Group
	// 3 = SubZone
	TzType                    TimeZoneType


	DeprecationStatus          TimeZoneDeprecationStatus // 0 = Unknown
	// 1 = Deprecated
	// 2 = Alias
	// 3 = Valid, Current Time Zone

	FuncDeclaration            []byte

	isInitialized              bool
}

// CopyOut - Creates and returns a deep copy of the current
// TimeZoneDataDto instance.
//
func (tzDataDto *TimeZoneDataDto) CopyOut() TimeZoneDataDto {

	newTzDto := TimeZoneDataDto{}

	if !tzDataDto.isInitialized {
		return newTzDto
	}

	newTzDto.GroupName = tzDataDto.GroupName
	newTzDto.TzName = tzDataDto.TzName
	newTzDto.TzCanonicalValue = tzDataDto.TzCanonicalValue
	newTzDto.TzAliasValue = tzDataDto.TzAliasValue
	newTzDto.TzValue = tzDataDto.TzValue
	newTzDto.TzSortValue = tzDataDto.TzSortValue
	newTzDto.FuncSelfReferenceVariable = tzDataDto.FuncSelfReferenceVariable
	newTzDto.FuncType = tzDataDto.FuncType
	newTzDto.FuncName = tzDataDto.FuncName
	newTzDto.FuncReturnType = tzDataDto.FuncReturnType
	newTzDto.FuncReturnValue = tzDataDto.FuncReturnValue
	newTzDto.SourceFileNameExt = tzDataDto.SourceFileNameExt
	newTzDto.TzClass = tzDataDto.TzClass
	newTzDto.TzType = tzDataDto.TzType
	newTzDto.DeprecationStatus = tzDataDto.DeprecationStatus
	newTzDto.isInitialized = tzDataDto.isInitialized

	return newTzDto
}

// CopyIn - Receives an input parameter TimeZoneDataDto instance
// copies all of the data fields to the current TimeZoneDataDto instance.
// When complete, both TimeZoneDataDto instances are equivalent.
//
func (tzDataDto *TimeZoneDataDto) CopyIn(
	inTzDataDto *TimeZoneDataDto) {

	tzDataDto.GroupName = inTzDataDto.GroupName
	tzDataDto.TzName = inTzDataDto.TzName
	tzDataDto.TzCanonicalValue = inTzDataDto.TzCanonicalValue
	tzDataDto.TzAliasValue = inTzDataDto.TzAliasValue
	tzDataDto.TzValue = inTzDataDto.TzValue
	tzDataDto.TzSortValue = inTzDataDto.TzSortValue
	tzDataDto.FuncSelfReferenceVariable = inTzDataDto.FuncSelfReferenceVariable
	tzDataDto.FuncType = inTzDataDto.FuncType
	tzDataDto.FuncName = inTzDataDto.FuncName
	tzDataDto.FuncReturnType = inTzDataDto.FuncReturnType
	tzDataDto.FuncReturnValue = inTzDataDto.FuncReturnValue
	tzDataDto.SourceFileNameExt = inTzDataDto.SourceFileNameExt
	tzDataDto.TzClass = inTzDataDto.TzClass
	tzDataDto.TzType = inTzDataDto.TzType
	tzDataDto.DeprecationStatus = inTzDataDto.DeprecationStatus
	tzDataDto.isInitialized = inTzDataDto.isInitialized

}

// TimeZoneDataDto - Compares isInitialized, GroupName, SubTzName,
// TzName, TzAliasValue TzCanonicalValue and TzValue data elements
// encapsulated by input parameter 'tzDDto' and the current
// TimeZoneDataDto instance.  If these values are identical,
// this method returns 'true'.
//
// Note that data values SourceFileNameExt, TzClass,
// and DeprecationStatus are NOT compared.
//
func (tzDataDto *TimeZoneDataDto) EqualValues( tzDDto TimeZoneDataDto) bool {

	if tzDataDto.isInitialized == tzDDto.isInitialized &&
		tzDataDto.GroupName == tzDDto.GroupName &&
		tzDataDto.TzName == tzDDto.TzName &&
		tzDataDto.TzAliasValue == tzDDto.TzAliasValue &&
		tzDataDto.TzCanonicalValue == tzDDto.TzCanonicalValue &&
		tzDataDto.TzValue == tzDDto.TzValue {
		return true
	}

	return false
}

// EqualClass - Compares the TzClass values for input parameter
// 'TzDDto' and the current TimeZoneDataDto instance. If they
// are equivalent, this method returns true.
func (tzDataDto *TimeZoneDataDto) EqualClass(tzDDto TimeZoneDataDto) bool {

	if tzDataDto.TzClass == tzDDto.TzClass {
		return true
	}

	return false
}

// EqualType - Compares the Time Zone Type values for input parameter
// 'TzDDto' and the current TimeZoneDataDto instance. If they
// are equivalent, this method returns true.
func (tzDataDto *TimeZoneDataDto) EqualType(tzDDto TimeZoneDataDto) bool {

	if tzDataDto.TzType == tzDDto.TzType {
		return true
	}

	return false
}

// EqualDeprecationStatus - Compares the DeprecationStatus values for input
// parameter 'TzDDto' and the current TimeZoneDataDto instance. If they are
// equivalent, this method returns true.
func (tzDataDto *TimeZoneDataDto) EqualDeprecationStatus(tzDDto TimeZoneDataDto) bool {

	if tzDataDto.DeprecationStatus == tzDDto.DeprecationStatus {
		return true
	}

	return false
}


// IsInitialized - Returns the value of internal data field
// TimeZoneDataDto.isInitialized .
func (tzDataDto *TimeZoneDataDto) IsInitialized() bool {
	return tzDataDto.isInitialized
}

// New - Creates and returns a new instance of the TimeZoneDataDto Type.
//
func (tzDataDto TimeZoneDataDto) New(
	groupName,
	subTzName,
	tzName,
	tzCanonicalValue,
	tzAliasValue,
	tzValue,
	tzSortValue,
	funcSelfReferenceVariable,
	funcType,
	funcName,
	funcReturnType,
	funcReturnValue,
	srcFileNameExt string,
	tzClass TimeZoneClass,
	tzType TimeZoneType,
	deprecationStatus TimeZoneDeprecationStatus) (TimeZoneDataDto, error) {

	ePrefix := "TimeZoneDataDto.NewTimeZone() - ERROR:\n"
	newTzDto := TimeZoneDataDto{}

	if len(groupName) == 0 {
		return newTzDto,
			errors.New(ePrefix + "Input Parameter 'groupName' is an EMPTY string!\n")
	}

	// subTzName empty strings are allowed.

	if len(tzName) == 0 {
		return newTzDto,
			errors.New(ePrefix + "Input Parameter 'tzName' is an EMPTY string!\n")
	}

	if len(tzCanonicalValue) == 0 {
		return newTzDto,
			errors.New(ePrefix + "Input Parameter 'tzCanonicalValue' is an EMPTY string!\n")
	}

	err := tzClass.ClassIsValid()

	if err != nil {
		return newTzDto,
			fmt.Errorf(ePrefix + "Input Parameter 'tzClass' is INVALID!\n" +
				"tzClass='%v'", int(tzClass))
	}

	err = deprecationStatus.StatusIsValid()

	if err != nil {
		return TimeZoneDataDto{},
			fmt.Errorf(ePrefix + "Input Parameter 'deprecationStatus' is INVALID!\n" +
				"deprecationStatus='%v'", int(deprecationStatus))
	}

	newTzDto.GroupName = groupName
	newTzDto.TzName = tzName
	newTzDto.TzCanonicalValue = tzCanonicalValue
	newTzDto.TzAliasValue = tzAliasValue
	newTzDto.TzValue = tzValue
	newTzDto.TzSortValue = tzSortValue
	newTzDto.FuncSelfReferenceVariable = funcSelfReferenceVariable
	newTzDto.FuncType = funcType
	newTzDto.FuncName = funcName
	newTzDto.FuncReturnType = funcReturnType
	newTzDto.FuncReturnValue = funcReturnValue
	newTzDto.SourceFileNameExt = srcFileNameExt
	newTzDto.TzClass = tzClass
	newTzDto.TzType = tzType
	newTzDto.DeprecationStatus = deprecationStatus
	newTzDto.isInitialized = true

	return newTzDto, nil
}

// NewSortValue - Creates and returns a new time zone sort value
// based on a time zone value passed in parameter, 'tzValue'.
//
func (tzDataDto TimeZoneDataDto) NewSortValue(tzValue string) string {

	numStrProfile,
	err := strops.StrOps{}.ExtractNumericDigits(
		tzValue,
		0,
		"",
		"",
		"")

	if err != nil {
		return tzValue
	}

	if numStrProfile.NumStrLen < 1 {
		return tzValue
	}

	str1 := tzValue[:numStrProfile.FirstNumCharIndex]
	str2 := tzValue[numStrProfile.FirstNumCharIndex + numStrProfile.NumStrLen:]

	number, err := strconv.Atoi(numStrProfile.NumStr)

	if err != nil {
		return tzValue
	}

	sortName := fmt.Sprintf(str1 + "%02d" + str2, number)

	return sortName
}

// SetIsInitialized - Sets the value of internal data field
// TimeZoneDataDto.isInitialized .
func (tzDataDto *TimeZoneDataDto) SetIsInitialized(isInitialized bool) {
	tzDataDto.isInitialized = isInitialized
}

//SelectTzDto - Select from an array TimeZoneDataDto objects.
type SelectTzDto []TimeZoneDataDto

// MajorGroupExists - Performs a search for on TimeZoneDataDto array
// for a match on TimeZoneDataDto.GroupName. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of false is returned and the
// integer index value is set to -1.
//
func (selTzDto SelectTzDto) MajorGroupExists(majorGroupName string, useLwrCase bool) (bool, int){

	if useLwrCase {
		majorGroupName = strings.ToLower(majorGroupName)
	}

	for i:=0; i < len(selTzDto); i++ {

		if useLwrCase {

			if strings.ToLower(selTzDto[i].GroupName) == majorGroupName {
				return true, i
			}
		} else if selTzDto[i].GroupName == majorGroupName {

			return true, i
		}

	}

	return false, -1
}

// TzNameExists - Performs a search for on TimeZoneDataDto array
// for a match on TimeZoneDataDto.TzName. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of false is returned and the
// integer index value is set to -1.
//
func (selTzDto SelectTzDto) TzNameExists(
	tzName string, useLwrCase bool) (bool,int) {

	if useLwrCase {
		tzName = strings.ToLower(tzName)
	}

	for i:=0; i < len(selTzDto); i++ {

		if useLwrCase {
			if strings.ToLower(selTzDto[i].TzName) == tzName{
				return true, i
			}
		} else if selTzDto[i].TzName == tzName {

			return true, i
		}

	}
	return false, -1
}


// TzValueExists - Performs a search for on TimeZoneDataDto array
// for a match on TimeZoneDataDto.TzCanonicalValue. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of false is returned and the
// integer index value is set to -1.
//
func (selTzDto SelectTzDto) TzValueExists(
	tzValue string, useLwrCase bool) (bool, int) {

	if useLwrCase {
		tzValue = strings.ToLower(tzValue)
	}

	for i:=0; i < len(selTzDto); i++ {

		if useLwrCase {
			if tzValue == strings.ToLower(selTzDto[i].TzCanonicalValue) {
				return true, i
			}
		} else if selTzDto[i].TzCanonicalValue == tzValue {
			return true, i
		}

	}
	return false, -1
}

