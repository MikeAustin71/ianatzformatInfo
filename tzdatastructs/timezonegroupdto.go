package tzdatastructs

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"strconv"
)

type TimeZoneGroupDto struct {
	ParentGroupName    string
	GroupName          string
	GroupSortValue     string
	TypeName           string
	TypeValue          string
	IanaVariableName   string
	SourceFileNameExt  string
	GroupType          TimeZoneGroupType
	GroupClass         TimeZoneGroupClass
	DeprecationStatus  TimeZoneDeprecationStatus
	isInitialized      bool
}

// CopyOut - Creates and returns a deep copy of the current
// TimeZoneGroupDto instance.
//
func (tzGrpDto *TimeZoneGroupDto) CopyOut() TimeZoneGroupDto {

	newTzGrpDto := TimeZoneGroupDto{}

	if !tzGrpDto.isInitialized {
		return newTzGrpDto
	}

	newTzGrpDto.ParentGroupName       = tzGrpDto.ParentGroupName
	newTzGrpDto.GroupName             = tzGrpDto.GroupName
	newTzGrpDto.GroupSortValue        = tzGrpDto.GroupSortValue
	newTzGrpDto.TypeName              = tzGrpDto.TypeName
	newTzGrpDto.TypeValue             = tzGrpDto.TypeValue
	newTzGrpDto.IanaVariableName      = tzGrpDto.IanaVariableName
	newTzGrpDto.SourceFileNameExt     = tzGrpDto.SourceFileNameExt
	newTzGrpDto.GroupType             = tzGrpDto.GroupType
	newTzGrpDto.GroupClass             = tzGrpDto.GroupClass
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

	tzGrpDto.ParentGroupName     = inGrpDto.ParentGroupName
	tzGrpDto.GroupName           = inGrpDto.GroupName
	tzGrpDto.GroupSortValue      = inGrpDto.GroupSortValue
	tzGrpDto.TypeName            = inGrpDto.TypeName
	tzGrpDto.TypeValue           = inGrpDto.TypeValue
	tzGrpDto.IanaVariableName    = inGrpDto.IanaVariableName
	tzGrpDto.SourceFileNameExt   = inGrpDto.SourceFileNameExt
	tzGrpDto.GroupType           = inGrpDto.GroupType
	tzGrpDto.GroupClass           = inGrpDto.GroupClass
	tzGrpDto.DeprecationStatus   = inGrpDto.DeprecationStatus
	tzGrpDto.isInitialized       = inGrpDto.isInitialized
}

// IsInitialized - Returns the value of internal data field
// TimeZoneGroupDto.isInitialized .
func (tzGrpDto *TimeZoneGroupDto) IsInitialized() bool {
	return tzGrpDto.isInitialized
}

// EqualGroupTypes - Compares the GroupType data values for input
// parameter 'tzGrp2' and the current TimeZoneGroupDto. If
// they are equivalent, this method returns 'true'.
func (tzGrpDto *TimeZoneGroupDto) EqualGroupTypes(
	tzGrp2 TimeZoneGroupDto) bool {

	if tzGrpDto.GroupType == tzGrp2.GroupType {
		return true
	}

	return false
}

// EqualGroupClass - Compares the GroupClass data values for input
// parameter 'tzGrp2' and the current TimeZoneGroupDto. If
// they are equivalent, this method returns 'true'.
func (tzGrpDto *TimeZoneGroupDto) EqualGroupClass(
	tzGrp2 TimeZoneGroupDto) bool {

	if tzGrpDto.GroupClass == tzGrp2.GroupClass {
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

		if tzGrpDto.GroupName == tzGrpDto2.GroupName {
			return true
		}

		return false
}

// Creates and returns a new instance of TimeZoneGroupDto.
//
func (tzGrpDto TimeZoneGroupDto) New(
	parentGroupName,
	groupName,
	groupSortValue,
	typeName,
	typeValue,
	ianaVariableName,
	sourceFileNameExt string,
	groupType TimeZoneGroupType,
	groupClass TimeZoneGroupClass,
	deprecationStatus TimeZoneDeprecationStatus) (TimeZoneGroupDto, error) {

	ePrefix := "TimeZoneGroupDto.New() "

	newTzGroupDto := TimeZoneGroupDto{}

	if len(groupName) == 0 {
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

	newTzGroupDto.ParentGroupName = parentGroupName
	newTzGroupDto.GroupName = groupName
	newTzGroupDto.GroupSortValue = groupSortValue
	newTzGroupDto.TypeName = typeName
	newTzGroupDto.TypeValue = typeValue
	newTzGroupDto.IanaVariableName = ianaVariableName
	newTzGroupDto.GroupType = groupType
	newTzGroupDto.GroupClass = groupClass
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

// SortByTzGroupName - This type provides support methods for
// sorting Time Zone Major Group Dto Arrays by Major Group Name.
//
// Example Usage:
//    sort.Sort(SortByTzGroupName(tzMajorGroupDtoArray))
//
type SortByTzGroupName []TimeZoneGroupDto

// Len - Required by the sort.Interface
func (sortMjrGrpName SortByTzGroupName) len() int {
	return len(sortMjrGrpName)
}

// Swap - Required by the sort.Interface
func (sortMjrGrpName SortByTzGroupName) Swap(i, j int) {
	sortMjrGrpName[i], sortMjrGrpName[j] = sortMjrGrpName[j], sortMjrGrpName[i]
}

// Less - Required by the sort.Interface
func (sortMjrGrpName SortByTzGroupName) Less(i, j int) bool {
	return sortMjrGrpName[i].GroupName < sortMjrGrpName[j].GroupName
}

