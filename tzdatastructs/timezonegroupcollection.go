package tzdatastructs

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// TimeZoneGroupCollection a collection of TimeZoneGroupDto
// objects.
type TimeZoneGroupCollection struct {
	tzGroups []TimeZoneGroupDto
}

// Adds a TimeZoneGroupDto to the collection if, and
// only if, this is a new Time Zone Major Group. This method
// will NOT allow duplicate TimeZoneGroupDto objects
// to be added to the collection.
func (tzGrpCol *TimeZoneGroupCollection) AddIfNew(
	tzGrpDto TimeZoneGroupDto) (isNew bool, err error) {

	ePrefix := "TimeZoneGroupCollection.AddIfNew() "
	isNew = true
	err = nil

	if tzGrpCol.tzGroups == nil {
		tzGrpCol.tzGroups = make([]TimeZoneGroupDto, 0, 300)
	}

	if !tzGrpDto.IsInitialized() {
		err = errors.New(ePrefix +
			"Input Parameter 'tzGrpDto' is uninitialized and INVALID!\n")
		return isNew, err
	}

	for i:= 0; i < len(tzGrpCol.tzGroups); i++ {
		if tzGrpCol.tzGroups[i].EqualNames(&tzGrpDto) {
			isNew = false
			return isNew, err
		}
	}

	isNew = true

	tzGrpCol.tzGroups = append(tzGrpCol.tzGroups, tzGrpDto)

	return isNew, err
	}

// Adds a TimeZoneGroupDto to the collection. This allows
// duplicate TimeZoneGroupDto objects to be added to the
// collection because no test is performed to determine
// if the new TimeZoneGroupDto previously exists in the
// collection.
//
func (tzGrpCol *TimeZoneGroupCollection) Add(
	tzGrpDto TimeZoneGroupDto) error {

	ePrefix := "TimeZoneGroupCollection.Add() "

	if tzGrpCol.tzGroups == nil {
		tzGrpCol.tzGroups = make([]TimeZoneGroupDto, 0, 300)
	}

	if !tzGrpDto.IsInitialized() {
		return errors.New(ePrefix +
			"Input Parameter 'tzGrpDto' is uninitialized and INVALID!\n")
	}

	tzGrpCol.tzGroups = append(tzGrpCol.tzGroups, tzGrpDto)

	return nil
}

// ContainsGroup - Returns true if the group in question is
// already part of the group.collection.
func (tzGrpCol *TimeZoneGroupCollection) ContainsGroup(
	tzGroup TimeZoneGroupDto) (containsGroup bool, index int) {

	containsGroup = false
	index = -1

	if tzGrpCol.tzGroups == nil {
		tzGrpCol.tzGroups = make([]TimeZoneGroupDto, 0, 300)
		return containsGroup, index
	}

	for i:= 0; i < len(tzGrpCol.tzGroups); i++ {

		if tzGrpCol.tzGroups[i].EqualNames(&tzGroup) {
			containsGroup = true
			return containsGroup, i
		}
	}

	return containsGroup, index
}

// ContainsGroupName - Returns true if the group name in question is
// already part of the group collection.
func (tzGrpCol *TimeZoneGroupCollection) ContainsGroupName(
	parentGroupName, groupName string) (containsGroupName bool, index int) {

	containsGroupName = false
	index = -1

	if tzGrpCol.tzGroups == nil {
		tzGrpCol.tzGroups = make([]TimeZoneGroupDto, 0, 300)
		return containsGroupName, index
	}

	for i:= 0; i < len(tzGrpCol.tzGroups); i++ {

		if tzGrpCol.tzGroups[i].GroupName == groupName &&
			tzGrpCol.tzGroups[i].ParentGroupName == parentGroupName {
			containsGroupName = true
			return containsGroupName, i
		}
	}

	return containsGroupName, index
}

// GetNumberOfGroups - Returns the number of group elements
// in this collection.
//
func (tzGrpCol *TimeZoneGroupCollection) GetNumberOfGroups() int {

	if tzGrpCol.tzGroups == nil {
		tzGrpCol.tzGroups = make([]TimeZoneGroupDto, 0, 300)
		return 0
	}

	return len(tzGrpCol.tzGroups)
}

// New - Creates and returns a correctly initialized TimeZoneGroupCollection.
//
func (tzGrpCol TimeZoneGroupCollection) New() TimeZoneGroupCollection {
	newTzGrp := TimeZoneGroupCollection{}

	newTzGrp.tzGroups = make([]TimeZoneGroupDto, 0, 300)

	return newTzGrp
}

// PeekPtr - Returns a pointer to the TimeZoneGroupDto
// element located at input parameter 'index' within
// the collection.
//
func (tzGrpCol *TimeZoneGroupCollection) PeekPtr(
	index int) (*TimeZoneGroupDto, error) {

	ePrefix := "TimeZoneGroupCollection.PeekPtr() "

	if tzGrpCol.tzGroups == nil {
		tzGrpCol.tzGroups = make([]TimeZoneGroupDto, 0, 300)
	}

	lenGroupAry := len(tzGrpCol.tzGroups)

	if lenGroupAry == 0 {
		return &TimeZoneGroupDto{},
			errors.New (ePrefix + "Collection is EMPTY!\n")
	}

	if index >= lenGroupAry {
		return &TimeZoneGroupDto{},
			fmt.Errorf(ePrefix + "Error: Input parameter 'index'\n" +
				"exceeds collection upper limit.\n" +
				"TimeZoneGroupDto Array last index='%v'\n" +
				"Input parameter 'index'='%v'\n",
				lenGroupAry-1, index)
	}

	return &tzGrpCol.tzGroups[index], nil
}

// Sort - Sorts the collection (TimeZoneGroupCollection.tzGroups)
// by Parent Group and Group.
//
// If input parameter, 'caseSensitive' is false, all strings will
// be converted to lower case before string comparisons are executed.
//
func (tzGrpCol *TimeZoneGroupCollection) Sort(caseSensitiveSort bool) {

	if tzGrpCol.tzGroups == nil {
		tzGrpCol.tzGroups = make([]TimeZoneGroupDto, 0, 300)
		return
	}

	if len(tzGrpCol.tzGroups) < 2 {
		return
	}

	var less func(i, j int) bool

	if !caseSensitiveSort {
		less = func(i, j int) bool {

			strI := strings.ToLower(tzGrpCol.tzGroups[i].ParentGroupName) + "/" +
				strings.ToLower(tzGrpCol.tzGroups[i].GroupSortValue)

			strJ := strings.ToLower(tzGrpCol.tzGroups[j].ParentGroupName) + "/" +
				strings.ToLower(tzGrpCol.tzGroups[j].GroupSortValue)

			return strI < strJ
		}
	} else {

		less = func(i, j int) bool {

			strI := tzGrpCol.tzGroups[i].ParentGroupName + "/" +
				tzGrpCol.tzGroups[i].GroupSortValue

			strJ := tzGrpCol.tzGroups[j].ParentGroupName + "/" +
				tzGrpCol.tzGroups[j].GroupSortValue

			return strI < strJ
		}
	}

	sort.Slice(tzGrpCol.tzGroups, less)

	return
}