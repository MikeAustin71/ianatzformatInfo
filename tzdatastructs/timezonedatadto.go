package tzdatastructs

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"strconv"
)

type TimeZoneDataDto struct {
	ParentGroupName           string
	GroupName                 string
	TzName                    string
	TzAliasValue              string
	TzCanonicalValue          string
	TzValue                   string
	TzSortValue               string
	UtcOffset                 []string // Two element array populated with UTC offsets
	                                   // First offset is June 15th offset. Second
	                                   // offset is December 15th offset. The UTC
	                                   // offsets are formatted in accordance with the
	                                   // following examples:
	                                   //      UTC+0500, UTC-0500, UTC+1000, UTC+1000
	ArrayStorageLevel         int    // 0-2
	FuncSelfReferenceVariable string
	FuncType                  string
	FuncName                  string
	FuncReturnType            string
	FuncReturnValue           string
	SourceFileNameExt         string
	WorldRegionSortCode       int

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

	// 0 = None
	// 1 = TimeZone
	// 2 = LinkZone
	TzCategory                TimeZoneCategory

	// 0 = None
	// 1 = Iana
	// 2 = Military
	// 3 = Other
	TzSource                  TimeZoneSource

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

	newTzDto.ParentGroupName = tzDataDto.ParentGroupName
	newTzDto.GroupName = tzDataDto.GroupName
	newTzDto.TzName = tzDataDto.TzName
	newTzDto.TzCanonicalValue = tzDataDto.TzCanonicalValue
	newTzDto.TzAliasValue = tzDataDto.TzAliasValue
	newTzDto.TzValue = tzDataDto.TzValue
	newTzDto.TzSortValue = tzDataDto.TzSortValue
	newTzDto.ArrayStorageLevel = tzDataDto.ArrayStorageLevel
	newTzDto.FuncSelfReferenceVariable = tzDataDto.FuncSelfReferenceVariable
	newTzDto.FuncType = tzDataDto.FuncType
	newTzDto.FuncName = tzDataDto.FuncName
	newTzDto.FuncReturnType = tzDataDto.FuncReturnType
	newTzDto.FuncReturnValue = tzDataDto.FuncReturnValue
	newTzDto.SourceFileNameExt = tzDataDto.SourceFileNameExt
	newTzDto.WorldRegionSortCode = tzDataDto.WorldRegionSortCode
	newTzDto.TzClass = tzDataDto.TzClass
	newTzDto.TzType = tzDataDto.TzType
	newTzDto.TzCategory = tzDataDto.TzCategory
	newTzDto.TzSource = tzDataDto.TzSource
	newTzDto.DeprecationStatus = tzDataDto.DeprecationStatus
	newTzDto.FuncDeclaration = make([]byte, len(tzDataDto.FuncDeclaration))
	copy(newTzDto.FuncDeclaration, tzDataDto.FuncDeclaration)
	newTzDto.isInitialized = tzDataDto.isInitialized

	newTzDto.UtcOffset = make([]string, 2)

	lenTzDataDtoUtcOffset := len(tzDataDto.UtcOffset)

	switch lenTzDataDtoUtcOffset {
	case 0:
		newTzDto.UtcOffset[0] = ""
		newTzDto.UtcOffset[1] = ""
	case 1:
		newTzDto.UtcOffset[0] = tzDataDto.UtcOffset[0]
		newTzDto.UtcOffset[1] = tzDataDto.UtcOffset[0]
	case 2:
		newTzDto.UtcOffset[0] = tzDataDto.UtcOffset[0]
		newTzDto.UtcOffset[1] = tzDataDto.UtcOffset[1]
	default:
		newTzDto.UtcOffset[0] = ""
		newTzDto.UtcOffset[1] = ""
	}

	return newTzDto
}

// CopyIn - Receives an input parameter TimeZoneDataDto instance
// copies all of the data fields to the current TimeZoneDataDto instance.
// When complete, both TimeZoneDataDto instances are equivalent.
//
func (tzDataDto *TimeZoneDataDto) CopyIn(
	inTzDataDto *TimeZoneDataDto) {

	tzDataDto.ParentGroupName = inTzDataDto.ParentGroupName
	tzDataDto.GroupName = inTzDataDto.GroupName
	tzDataDto.TzName = inTzDataDto.TzName
	tzDataDto.TzCanonicalValue = inTzDataDto.TzCanonicalValue
	tzDataDto.TzAliasValue = inTzDataDto.TzAliasValue
	tzDataDto.TzValue = inTzDataDto.TzValue
	tzDataDto.TzSortValue = inTzDataDto.TzSortValue
	tzDataDto.ArrayStorageLevel = inTzDataDto.ArrayStorageLevel
	tzDataDto.FuncSelfReferenceVariable = inTzDataDto.FuncSelfReferenceVariable
	tzDataDto.FuncType = inTzDataDto.FuncType
	tzDataDto.FuncName = inTzDataDto.FuncName
	tzDataDto.FuncReturnType = inTzDataDto.FuncReturnType
	tzDataDto.FuncReturnValue = inTzDataDto.FuncReturnValue
	tzDataDto.SourceFileNameExt = inTzDataDto.SourceFileNameExt
	tzDataDto.WorldRegionSortCode = inTzDataDto.WorldRegionSortCode
	tzDataDto.TzClass = inTzDataDto.TzClass
	tzDataDto.TzType = inTzDataDto.TzType
	tzDataDto.TzCategory = inTzDataDto.TzCategory
	tzDataDto.TzSource = inTzDataDto.TzSource
	tzDataDto.DeprecationStatus = inTzDataDto.DeprecationStatus
	tzDataDto.FuncDeclaration = make([]byte, len(inTzDataDto.FuncDeclaration))
	copy(tzDataDto.FuncDeclaration, inTzDataDto.FuncDeclaration)
	tzDataDto.isInitialized = inTzDataDto.isInitialized

	tzDataDto.UtcOffset = make([]string, 2)

	lenInTzDataDtoUtcOffset := len(inTzDataDto.UtcOffset)

	switch lenInTzDataDtoUtcOffset {
	case 0:
		tzDataDto.UtcOffset[0] = ""
		tzDataDto.UtcOffset[1] = ""
	case 1:
		tzDataDto.UtcOffset[0] = inTzDataDto.UtcOffset[0]
		tzDataDto.UtcOffset[1] = inTzDataDto.UtcOffset[0]
	case 2:
		tzDataDto.UtcOffset[0] = inTzDataDto.UtcOffset[0]
		tzDataDto.UtcOffset[1] = inTzDataDto.UtcOffset[1]
	default:
		tzDataDto.UtcOffset[0] = ""
		tzDataDto.UtcOffset[1] = ""
	}

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

	if tzDataDto.isInitialized != tzDDto.isInitialized ||
		tzDataDto.WorldRegionSortCode != tzDataDto.WorldRegionSortCode ||
		tzDataDto.ParentGroupName != tzDDto.ParentGroupName ||
		tzDataDto.GroupName != tzDDto.GroupName ||
		tzDataDto.TzName != tzDDto.TzName ||
		tzDataDto.TzAliasValue != tzDDto.TzAliasValue ||
		tzDataDto.TzCanonicalValue != tzDDto.TzCanonicalValue ||
		tzDataDto.TzValue != tzDDto.TzValue {
		return false
	}

	if len(tzDataDto.UtcOffset) != len(tzDDto.UtcOffset) {
		return false
	}

	for i:=0; i < len(tzDataDto.UtcOffset); i++ {
		if tzDataDto.UtcOffset[i] != tzDDto.UtcOffset[i] {
			return false
		}
	}

	return true
}

// IsInitialized - Returns the value of internal data field
// TimeZoneDataDto.isInitialized .
func (tzDataDto *TimeZoneDataDto) IsInitialized() bool {
	return tzDataDto.isInitialized
}

// New - Creates and returns a new instance of the TimeZoneDataDto Type.
//
func (tzDataDto TimeZoneDataDto) New(
	parentGroupName,
	groupName,
	tzName,
	tzAliasValue,
	tzCanonicalValue,
	tzValue string,
	tzUtcOffset []string,
	tzSortValue string,
	arrayStorageLevel int,
	funcSelfReferenceVariable,
	funcType,
	funcName,
	funcReturnType,
	funcReturnValue,
	srcFileNameExt string,
	worldRegionSortCode int,
	tzClass TimeZoneClass,
	tzType TimeZoneType,
	tzCategory TimeZoneCategory,
	tzSource TimeZoneSource,
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

	newTzDto.ParentGroupName = parentGroupName
	newTzDto.GroupName = groupName
	newTzDto.TzName = tzName
	newTzDto.TzAliasValue = tzAliasValue
	newTzDto.TzCanonicalValue = tzCanonicalValue
	newTzDto.TzValue = tzValue
	newTzDto.TzSortValue = tzSortValue
	newTzDto.ArrayStorageLevel = arrayStorageLevel
	newTzDto.FuncSelfReferenceVariable = funcSelfReferenceVariable
	newTzDto.FuncType = funcType
	newTzDto.FuncName = funcName
	newTzDto.FuncReturnType = funcReturnType
	newTzDto.FuncReturnValue = funcReturnValue
	newTzDto.SourceFileNameExt = srcFileNameExt
	newTzDto.WorldRegionSortCode = worldRegionSortCode
	newTzDto.TzClass = tzClass
	newTzDto.TzType = tzType
	newTzDto.TzCategory = tzCategory
	newTzDto.TzSource = tzSource
	newTzDto.DeprecationStatus = deprecationStatus
	newTzDto.isInitialized = true

	newTzDto.UtcOffset = make([]string, 2)

	lenTzUtcOffsetUtcOffset := len(tzUtcOffset)

	switch lenTzUtcOffsetUtcOffset {
	case 0:
		newTzDto.UtcOffset[0] = ""
		newTzDto.UtcOffset[1] = ""
	case 1:
		newTzDto.UtcOffset[0] = tzUtcOffset[0]
		newTzDto.UtcOffset[1] = tzUtcOffset[0]
	case 2:
		newTzDto.UtcOffset[0] = tzUtcOffset[0]
		newTzDto.UtcOffset[1] = tzUtcOffset[1]
	default:
		newTzDto.UtcOffset[0] = ""
		newTzDto.UtcOffset[1] = ""
	}


	newTzDto.UtcOffset = tzUtcOffset

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
