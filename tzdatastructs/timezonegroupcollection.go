package tzdatastructs

import (
	"errors"
	"fmt"
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
		if tzGrpCol.tzGroups[i].EqualNameValues(tzGrpDto) {
			isNew = false
			return isNew, err
		}
	}

	isNew = true

	tzGrpCol.tzGroups = append(tzGrpCol.tzGroups, tzGrpDto)

	return isNew, err
	}

func (tzGrpCol *TimeZoneGroupCollection) AddIfNewByDetail(
	majorGroupName,
	minorGroupName,
	compositeGroupName,
	sourceFileNameExt string,
	groupType TimeZoneGroupType,
	deprecationStatus TimeZoneDeprecationStatus) (isNew bool, err error) {

	ePrefix := "TimeZoneGroupCollection.AddIfNewByDetail() "

	isNew = false
	err = nil

	tzGrpDto, err2 := TimeZoneGroupDto{}.New(
		majorGroupName,
		minorGroupName,
		compositeGroupName,
		sourceFileNameExt,
		groupType,
		deprecationStatus)

	if err2 != nil {
		err = fmt.Errorf(ePrefix +
			"%v\n", err2.Error())
		return isNew, err
	}

	return tzGrpCol.AddIfNew(tzGrpDto)
}

// Adds a TimeZoneGroupDto to the collection. This allows
// duplicate TimeZoneGroupDto objects to be added to the
// collection.
func (tzGrpCol *TimeZoneGroupCollection) Add(
	tzGrpDto TimeZoneGroupDto) (isNew bool, err error) {

	ePrefix := "TimeZoneGroupCollection.Add() "
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
		if tzGrpCol.tzGroups[i].EqualNameValues(tzGrpDto) {
			isNew = false
			return isNew, err
		}
	}

	isNew = true

	tzGrpCol.tzGroups = append(tzGrpCol.tzGroups, tzGrpDto)

	return isNew, err
	}

// New - Creates and returns a correctly initialized TimeZoneGroupCollection.
//
func (tzGrpCol TimeZoneGroupCollection) New() TimeZoneGroupCollection {
	newTzGrp := TimeZoneGroupCollection{}

	newTzGrp.tzGroups = make([]TimeZoneGroupDto, 0, 300)

	return newTzGrp
}