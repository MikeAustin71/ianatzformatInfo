package tzdatastructs

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"strconv"
)

type TimeZoneGroupDto struct {
	MajorGroupName    string
	MinorGroupName    string
	GroupNameValue    string
	GroupSortValue    string
	TypeName          string
	IanaVariableName  string
	SourceFileNameExt string
	GroupType         TimeZoneGroupType
	DeprecationStatus TimeZoneDeprecationStatus
	isInitialized     bool
}

// CopyOut - Creates and returns a deep copy of the current
// TimeZoneGroupDto instance.
//
func (tzGrpDto *TimeZoneGroupDto) CopyOut() TimeZoneGroupDto {

	newTzGrpDto := TimeZoneGroupDto{}

	if !tzGrpDto.isInitialized {
		return newTzGrpDto
	}

	newTzGrpDto.MajorGroupName        = tzGrpDto.MajorGroupName
	newTzGrpDto.MinorGroupName        = tzGrpDto.MinorGroupName
	newTzGrpDto.GroupNameValue        = tzGrpDto.GroupNameValue
	newTzGrpDto.GroupSortValue        = tzGrpDto.GroupSortValue
	newTzGrpDto.TypeName              = tzGrpDto.TypeName
	newTzGrpDto.IanaVariableName      = tzGrpDto.IanaVariableName
	newTzGrpDto.SourceFileNameExt     = tzGrpDto.SourceFileNameExt
	newTzGrpDto.GroupType             = tzGrpDto.GroupType
	newTzGrpDto.DeprecationStatus     = tzGrpDto.DeprecationStatus
	newTzGrpDto.isInitialized         = true

	return newTzGrpDto
}

// CopyIn - Receives an input parameter TimeZoneGroupDto instance
// copies all of the data fields to the current TimeZoneGroupDto
// instance.
//
// When complete, both TimeZoneGroupDto instances are equivalent.
//
func (tzGrpDto *TimeZoneGroupDto) CopyIn(
	inGrpDto *TimeZoneGroupDto) {

	tzGrpDto.MajorGroupName      = inGrpDto.MajorGroupName
	tzGrpDto.MinorGroupName      = inGrpDto.MinorGroupName
	tzGrpDto.GroupNameValue      = inGrpDto.GroupNameValue
	tzGrpDto.GroupSortValue      = inGrpDto.GroupSortValue
	tzGrpDto.TypeName            = inGrpDto.TypeName
	tzGrpDto.IanaVariableName    = inGrpDto.IanaVariableName
	tzGrpDto.SourceFileNameExt   = inGrpDto.SourceFileNameExt
	tzGrpDto.GroupType           = inGrpDto.GroupType
	tzGrpDto.DeprecationStatus   = inGrpDto.DeprecationStatus
	tzGrpDto.isInitialized       = inGrpDto.isInitialized
}

// IsInitialized - Returns the value of internal data field
// TimeZoneGroupDto.isInitialized .
func (tzGrpDto *TimeZoneGroupDto) IsInitialized() bool {
	return tzGrpDto.isInitialized
}

// EqualNameValues - Compares the GroupType data values for input
// parameter 'tzMajorGrp2' and the current TimeZoneGroupDto. If
// they are equivalent, this method returns 'true'.
func (tzGrpDto *TimeZoneGroupDto) EqualGroupTypes(
	tzMajorGrp2 TimeZoneGroupDto) bool {

	if tzGrpDto.GroupType == tzMajorGrp2.GroupType {
		return true
	}

	return false
}

// EqualDeprecationStatus - Compares the DeprecationStatus data values
// for input parameter 'tzMajorGrp2' and the current TimeZoneGroupDto.
// If they are equivalent, this method returns 'true'.
func (tzGrpDto *TimeZoneGroupDto) EqualDeprecationStatus(
	tzMajorGrp2 TimeZoneGroupDto) bool {

	if tzGrpDto.DeprecationStatus == tzMajorGrp2.DeprecationStatus {
		return true
	}

	return false
}

// EqualNameValues - Compares the GroupNameValue data values for input
// parameter 'tzGrpDto2' and the current TimeZoneGroupDto. If
// they are equivalent, this method returns 'true'.
func (tzGrpDto *TimeZoneGroupDto) EqualNameValues(
	tzGrpDto2 TimeZoneGroupDto) bool {

		if tzGrpDto.GroupNameValue == tzGrpDto2.GroupNameValue {
			return true
		}

		return false
}

// Creates and returns a new instance of TimeZoneGroupDto.
//
func (tzGrpDto TimeZoneGroupDto) New(
	majorGroupName,
	minorGroupName,
	groupNameValue,
	groupSortValue,
	typeName,
	ianaVariableName,
	sourceFileNameExt string,
	groupType TimeZoneGroupType,
	deprecationStatus TimeZoneDeprecationStatus) (TimeZoneGroupDto, error) {

	ePrefix := "TimeZoneGroupDto.New() "

	newTzGroupDto := TimeZoneGroupDto{}

	if len(groupNameValue) == 0 {
		return newTzGroupDto,
			errors.New(ePrefix + "Input Parameter 'compositeGroupName' is an EMPTY string!\n")
	}

	err := groupType.TypeIsValid()

	if err != nil {
		return TimeZoneGroupDto{},
			fmt.Errorf(ePrefix + "Input Parameter 'groupType' is INVALID!\n" +
				"groupType='%v'", int(groupType))
	}

	err = deprecationStatus.StatusIsValid()

	if err != nil {
		return TimeZoneGroupDto{},
			fmt.Errorf(ePrefix + "Input Parameter 'deprecationStatus' is INVALID!\n" +
				"deprecationStatus='%v'", int(deprecationStatus))
	}

	newTzGroupDto.MajorGroupName = majorGroupName
	newTzGroupDto.MinorGroupName = minorGroupName
	newTzGroupDto.GroupNameValue = groupNameValue
	newTzGroupDto.GroupSortValue = groupSortValue
	newTzGroupDto.TypeName = typeName

	newTzGroupDto.IanaVariableName =
		ianaVariableName

	newTzGroupDto.GroupType = groupType
	newTzGroupDto.DeprecationStatus = deprecationStatus
	newTzGroupDto.isInitialized = true

	return newTzGroupDto, nil
}

// NewSortValue - Creates and returns a new time zone group
// sort value based on a time zone value passed in parameter,
// 'groupValue'.
//
func (tzGrpDto TimeZoneGroupDto) NewSortValue(groupValue string) string {

	numStrProfile,
	err := strops.StrOps{}.ExtractNumericDigits(
		groupValue,
		0,
		"",
		"",
		"")

	if err != nil {
		return groupValue
	}

	if numStrProfile.NumStrLen < 1 {
		return groupValue
	}

	str1 := groupValue[:numStrProfile.FirstNumCharIndex]
	str2 := groupValue[numStrProfile.FirstNumCharIndex + numStrProfile.NumStrLen:]

	number, err := strconv.Atoi(numStrProfile.NumStr)

	if err != nil {
		return groupValue
	}

	sortName := fmt.Sprintf(str1 + "%02d" + str2, number)

	return sortName
}

// SetIsInitialized - sets the value of internal data field
// TimeZoneGroupDto.isInitialized .
func (tzGrpDto *TimeZoneGroupDto) SetIsInitialized(isInitialized bool) {
	tzGrpDto.isInitialized = isInitialized
}


// SortByTzMajorGroupName - This type provides support methods for
// sorting Time Zone Major Group Dto Arrays by Major Group Name.
//
// Example Usage:
//    sort.Sort(SortByTzMajorGroupName(tzMajorGroupDtoArray))
//
type SortByTzMajorGroupName []TimeZoneGroupDto

// Len - Required by the sort.Interface
func (sortMjrGrpName SortByTzMajorGroupName) len() int {
	return len(sortMjrGrpName)
}

// Swap - Required by the sort.Interface
func (sortMjrGrpName SortByTzMajorGroupName) Swap(i, j int) {
	sortMjrGrpName[i], sortMjrGrpName[j] = sortMjrGrpName[j], sortMjrGrpName[i]
}

// Less - Required by the sort.Interface
func (sortMjrGrpName SortByTzMajorGroupName) Less(i, j int) bool {
	return sortMjrGrpName[i].MajorGroupName < sortMjrGrpName[j].MajorGroupName
}

