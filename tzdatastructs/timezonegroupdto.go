package tzdatastructs

import (
	"errors"
	"fmt"
)

type TimeZoneGroupDto struct {
	MajorGroupName     string
	MinorGroupName     string
	CompositeGroupName string
	SourceFileNameExt  string
	GroupType          TimeZoneGroupType
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

	newTzGrpDto.MajorGroupName        = tzGrpDto.MajorGroupName
	newTzGrpDto.MinorGroupName        = tzGrpDto.MinorGroupName
	newTzGrpDto.CompositeGroupName    = tzGrpDto.CompositeGroupName
	newTzGrpDto.GroupType             = tzGrpDto.GroupType
	newTzGrpDto.SourceFileNameExt     = tzGrpDto.SourceFileNameExt
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
	tzGrpDto.CompositeGroupName  = inGrpDto.CompositeGroupName
	tzGrpDto.GroupType           = inGrpDto.GroupType
	tzGrpDto.SourceFileNameExt   = inGrpDto.SourceFileNameExt
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

// EqualNameValues - Compares the CompositeGroupName data values for input
// parameter 'tzGrpDto2' and the current TimeZoneGroupDto. If
// they are equivalent, this method returns 'true'.
func (tzGrpDto *TimeZoneGroupDto) EqualNameValues(
	tzGrpDto2 TimeZoneGroupDto) bool {

		if tzGrpDto.CompositeGroupName == tzGrpDto2.CompositeGroupName {
			return true
		}

		return false
}

// Creates and returns a new instance of TimeZoneGroupDto.
//
func (tzGrpDto TimeZoneGroupDto) New(
	majorGroupName,
	minorGroupName,
	compositeGroupName,
	sourceFileNameExt string,
	groupType TimeZoneGroupType,
	deprecationStatus TimeZoneDeprecationStatus) (TimeZoneGroupDto, error) {

	ePrefix := "TimeZoneGroupDto.New() "

	newTzGroupDto := TimeZoneGroupDto{}

	if len(compositeGroupName) == 0 {
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
	newTzGroupDto.CompositeGroupName = compositeGroupName
	newTzGroupDto.GroupType = groupType
	newTzGroupDto.DeprecationStatus = deprecationStatus
	newTzGroupDto.isInitialized = true

	return newTzGroupDto, nil
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

