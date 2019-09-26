package tzdatastructs

import (
	"errors"
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
	groupName string) (containsGroupName bool, index int) {

	containsGroupName = false
	index = -1

	if tzGrpCol.tzGroups == nil {
		tzGrpCol.tzGroups = make([]TimeZoneGroupDto, 0, 300)
		return containsGroupName, index
	}

	for i:= 0; i < len(tzGrpCol.tzGroups); i++ {
		if tzGrpCol.tzGroups[i].GroupName == groupName {
			containsGroupName = true
			return containsGroupName, i
		}
	}

	return containsGroupName, index
}

// New - Creates and returns a correctly initialized TimeZoneGroupCollection.
//
func (tzGrpCol TimeZoneGroupCollection) New() TimeZoneGroupCollection {
	newTzGrp := TimeZoneGroupCollection{}

	newTzGrp.tzGroups = make([]TimeZoneGroupDto, 0, 300)

	return newTzGrp
}