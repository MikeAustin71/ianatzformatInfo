package tzdatastructs

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// SortTimeZoneDataByMjrGrpTzName - Sort by GroupName Name, TzName
//
// Example Usage:
//    sort.Sort(SortByTzGroupName(tzMajorGroupDtoArray))
//
type SortTimeZoneDataByMjrGrpTzName []TimeZoneDataDto

// Len - Required by the sort.Interface
func (sortByTzDtoName SortTimeZoneDataByMjrGrpTzName) Len() int {
	return len(sortByTzDtoName)
}

// Swap - Required by the sort.Interface
func (sortByTzDtoName SortTimeZoneDataByMjrGrpTzName) Swap(i, j int) {
	sortByTzDtoName[i], sortByTzDtoName[j] = sortByTzDtoName[j], sortByTzDtoName[i]
}

// Less - required by the sort.Interface
func (sortByTzDtoName SortTimeZoneDataByMjrGrpTzName) Less(i, j int) bool {

	if sortByTzDtoName[i].GroupName == sortByTzDtoName[j].GroupName {
		return sortByTzDtoName[i].TzName < sortByTzDtoName[j].TzName
	}

	return sortByTzDtoName[i].GroupName < sortByTzDtoName[j].GroupName
}


// TimeZoneDataCollection - is a collection TimeZoneDataDto objects.
// The collection effectively encapsulates a TimeZoneDataDto array.
type TimeZoneDataCollection struct {
	tzDataDtos []TimeZoneDataDto
}

// Add - Adds a TimeZoneDataDto object to the collection. This method will add
// duplicate TimeZoneDataDto instances to the collection.
func (tzDataCol *TimeZoneDataCollection) Add(tzDataDto TimeZoneDataDto) error {

	ePrefix := "TimeZoneDataCollection.Add() "

	if tzDataCol.tzDataDtos == nil {
		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)
	}

	if !tzDataDto.IsInitialized() {
		return errors.New(ePrefix + "Input Parameter 'tzDataDto' is uninitialized and INVALID!\n")
	}

	tzDataCol.tzDataDtos = append(tzDataCol.tzDataDtos, tzDataDto)

	return nil
}

// Adds a TimeZoneDataDto object to the collection if an object
// of equal value does NOT already exist in the collection. In
// other words, it will not allow duplicate TimeZoneDataDto
// instances to be added to the collection.
//
func (tzDataCol *TimeZoneDataCollection) AddIfNew(
	tzDataDto TimeZoneDataDto) (isNew bool, err error) {

	ePrefix := "TimeZoneDataCollection.AddIfNew() "

	isNew = false
	err = nil

	if tzDataCol.tzDataDtos == nil {
		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)
	}

	if !tzDataDto.IsInitialized() {
		err = errors.New(ePrefix + "Input Parameter 'tzDataDto' is uninitialized and INVALID!\n")
		return isNew, err
	}

	for i:=0; i < len(tzDataCol.tzDataDtos); i++ {
		if tzDataCol.tzDataDtos[i].EqualValues(tzDataDto){
			return isNew, err
		}
	}

	isNew = true

	tzDataCol.tzDataDtos = append(tzDataCol.tzDataDtos, tzDataDto)

	return isNew, err
}

// ContainsTzName - Determines if the collection contains a TimeZoneDataDto with a group name and
// time zone name equal to input parameters, 'groupName' and 'tzName'.
//
func (tzDataCol *TimeZoneDataCollection) ContainsTzName(
	parentGroupName, groupName, tzName string) (containsTz bool, index int) {

	containsTz = false
	index = -1

	if tzDataCol.tzDataDtos == nil {
		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)
		return containsTz, index
	}

	for i:=0; i < len(tzDataCol.tzDataDtos); i++ {
		if tzDataCol.tzDataDtos[i].ParentGroupName == parentGroupName &&
			tzDataCol.tzDataDtos[i].GroupName == groupName &&
			tzDataCol.tzDataDtos[i].TzName == tzName {
			containsTz = true
			index = i
			return containsTz, index
		}
	}

	return containsTz, index
}

// GetNumberOfTimeZones - Returns the number of time zone elements
// in this collection.
//
func (tzDataCol *TimeZoneDataCollection) GetNumberOfTimeZones() int {

	if tzDataCol.tzDataDtos == nil {
		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)
		return 0
	}

	return len(tzDataCol.tzDataDtos)
}

// GetZoneGroupCol - Returns a collection of time zones which
// match the Time Zone Group input parameter.
//
func (tzDataCol *TimeZoneDataCollection) GetZoneGroupCol(
	tzGroup TimeZoneGroupDto ) (TimeZoneDataCollection, error) {

	ePrefix := "TimeZoneDataCollection) GetZoneGroupCol() "

	if tzDataCol.tzDataDtos == nil {

		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)

	}

	tzCol := TimeZoneDataCollection{}.New()

	lenTzDataDtos := len(tzDataCol.tzDataDtos)

	if lenTzDataDtos == 0 {
		return tzCol,
		errors.New(ePrefix +
			"Time Zone Data Collection is EMPTY!\n")
	}

	for i:=0; i < lenTzDataDtos; i++ {
		if tzDataCol.tzDataDtos[i].ParentGroupName ==
			tzGroup.ParentGroupName &&
			tzDataCol.tzDataDtos[i].GroupName ==
			tzGroup.GroupName {

			err := tzCol.Add(tzDataCol.tzDataDtos[i].CopyOut())

			if err != nil {
				return TimeZoneDataCollection{}.New(),
					fmt.Errorf(ePrefix +
						"Error returned by tzCol.Add()\n" +
						"Group Name='%v'\n" +
						"Tz Name='%v'\n" +
						"Error='%v'\n",
						tzGroup.GroupName,
						tzDataCol.tzDataDtos[i].TzName,
						err.Error())
			}
		}
	}

	return tzCol, nil
}

// GroupExists - Performs a search for on the internal TimeZoneDataDto
// array for a match on TimeZoneDataDto.GroupName. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of 'false' is returned and the
// integer index value is set to -1.
//
// If the input parameter 'caseInsensitiveSearch' is set to 'true', the search for
// TimeZoneDataDto.GroupName will be conducted as a case insensitive search.
// This means that both strings are converted to lower case before the comparison
// is performed.
//
// If the input parameter 'caseInsensitiveSearch' is set to 'false', the search
// is conducted as a case sensitive comparison where upper and lower case
// characters are significant.
//
func (tzDataCol *TimeZoneDataCollection) GroupExists(
	parentGroupName, groupName string) (groupExists bool, index int) {

	groupExists = false
	index = -1

	if tzDataCol.tzDataDtos == nil {

		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)

	}

	lenTzDataDtoArray := len(tzDataCol.tzDataDtos)

	if lenTzDataDtoArray == 0 {
		return groupExists, index
	}


	for i:=0; i < lenTzDataDtoArray; i++ {

		if parentGroupName == tzDataCol.tzDataDtos[i].ParentGroupName &&
			groupName == tzDataCol.tzDataDtos[i].GroupName {
			return true, i
		}

	}

	return false, -1
}

// New - Returns a new and properly initialized instance of
// TimeZoneDataCollection.
func (tzDataCol TimeZoneDataCollection) New() TimeZoneDataCollection {

	newTzCol := TimeZoneDataCollection{}
	newTzCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)

	return newTzCol
}

// Peek - Returns a deep copy of the TimeZoneDataDto located in the internal
// TimeZoneDataDto array at input parameter 'index'.
//
// The internal array is not altered by this method.
//
func (tzDataCol *TimeZoneDataCollection) Peek(index int) (TimeZoneDataDto, error) {

	ePrefix := "TimeZoneDataCollection.Peek() "

	if tzDataCol.tzDataDtos == nil {

		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)

	}

	if index < 0 {
		return TimeZoneDataDto{},
			fmt.Errorf(ePrefix +
				"ERROR: Input parameter 'index' is less than zero and INVALID!\n" +
				"index='%v'", index)
	}

	lenTzDataDtos := len(tzDataCol.tzDataDtos)

	if lenTzDataDtos == 0 {
		return TimeZoneDataDto{}, errors.New(ePrefix +
			"ERROR: The Time Zone Data Collection is EMPTY!")
	}

	if index > (lenTzDataDtos - 1) {
		return TimeZoneDataDto{},
			fmt.Errorf(ePrefix +
				"ERROR: Input paramter 'index' exceeds array upper boundary.\n" +
				"TimeZoneDataDto Array last index='%v'\n" +
				"Input parameter 'index'='%v'\n ", lenTzDataDtos - 1, index )
	}

	return tzDataCol.tzDataDtos[index].CopyOut(), nil
}



// PopAtIndex - Returns the TimeZoneDataDto from collection array at the index
// specified by input parameter, 'index'.  That TimeZoneDataDto object is then
// deleted from the collection.
//
// If the index is out of bounds, this method returns an suitable error.
//
// If the collection is empty, this method returns an error of type 'io.EOF'.
//
func (tzDataCol *TimeZoneDataCollection) PopAtIndex(index int) (TimeZoneDataDto, error) {

	ePrefix := "TimeZoneDataCollection.PopAtIndex() "

	if tzDataCol.tzDataDtos == nil {
		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)
	}

	if index < 0 {
		return TimeZoneDataDto{},
			fmt.Errorf(ePrefix +
				"ERROR: Input parameter 'index' is less than Zero!\n" +
				"index='%v'\n", index)
	}

	arrayLen := len(tzDataCol.tzDataDtos)

	if arrayLen == 0 {
		return TimeZoneDataDto{}, io.EOF
	}

	if index >= arrayLen {
		return TimeZoneDataDto{},
			fmt.Errorf(ePrefix +
				"ERROR: Input parameter 'index' is out-of-bounds!\n" +
				"'index' exceeds the value of the last index in the collection array!\n" +
				"index='%v'\n" +
				"collection last index value='%v'", index, arrayLen-1)
	}

	var newTimeZoneDto TimeZoneDataDto

	if index == 0 {

		newTimeZoneDto = tzDataCol.tzDataDtos[0].CopyOut()
		tzDataCol.tzDataDtos = tzDataCol.tzDataDtos[1:]

	} else if index == arrayLen-1 {

		newTimeZoneDto = tzDataCol.tzDataDtos[arrayLen-1].CopyOut()
		tzDataCol.tzDataDtos = tzDataCol.tzDataDtos[0 : arrayLen-1]

	} else {
		newTimeZoneDto = tzDataCol.tzDataDtos[index].CopyOut()
		tzDataCol.tzDataDtos = append(tzDataCol.tzDataDtos[0:index], tzDataCol.tzDataDtos[index+1:]...)
	}

	return newTimeZoneDto, nil
}

// PopFirst - Returns the first TimeZoneDataDto in the collection and then
// deletes that TimeZoneDataDto from the collection.
//
// If the collection is empty, this method returns an error of type 'io.EOF'.
//
func (tzDataCol *TimeZoneDataCollection) PopFirst() (TimeZoneDataDto, error) {

	if tzDataCol.tzDataDtos == nil {
		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)
	}

	if len(tzDataCol.tzDataDtos) == 0 {
		return TimeZoneDataDto{}, io.EOF
	}

	newTzDataDto := tzDataCol.tzDataDtos[0].CopyOut()

	tzDataCol.tzDataDtos = tzDataCol.tzDataDtos[1:]

	return newTzDataDto, nil
}

// PopFirst - Returns the Last TimeZoneDataDto in the collection and then
// deletes that TimeZoneDataDto from the collection.
//
// If the collection is empty, this method returns an error of type 'io.EOF'.
//
func (tzDataCol *TimeZoneDataCollection) PopLast() (TimeZoneDataDto, error) {

	if tzDataCol.tzDataDtos == nil {
		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)
	}

	arrayLen := len(tzDataCol.tzDataDtos)

	if arrayLen == 0 {
		return TimeZoneDataDto{}, io.EOF
	}

	newTzDataDto := tzDataCol.tzDataDtos[arrayLen-1].CopyOut()

	tzDataCol.tzDataDtos = tzDataCol.tzDataDtos[0 : arrayLen-1]

	return newTzDataDto, nil
}

// SortByGroupTzName - Sort the collection by TimeZone Parent Group, Group and
// Time Zone Name.
//
func (tzDataCol *TimeZoneDataCollection) SortByGroupTzName(caseSensitiveSort bool) {

	if tzDataCol.tzDataDtos == nil {
		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)
	}

	if len(tzDataCol.tzDataDtos) < 2 {
		return
	}

	var less func(i, j int) bool

	if !caseSensitiveSort {
		less = func(i, j int) bool {

			strI := strings.ToLower(tzDataCol.tzDataDtos[i].ParentGroupName) + "/" +
				strings.ToLower(tzDataCol.tzDataDtos[i].GroupName) + "/" +
					strings.ToLower(tzDataCol.tzDataDtos[i].TzSortValue)

			strJ := strings.ToLower(tzDataCol.tzDataDtos[j].ParentGroupName) + "/" +
				strings.ToLower(tzDataCol.tzDataDtos[j].GroupName) + "/" +
					strings.ToLower(tzDataCol.tzDataDtos[j].TzSortValue)

				return strI < strJ
			}
	} else {

		less = func(i, j int) bool {

			strI := tzDataCol.tzDataDtos[i].ParentGroupName + "/" +
				tzDataCol.tzDataDtos[i].GroupName + "/" +
				tzDataCol.tzDataDtos[i].TzSortValue

			strJ := tzDataCol.tzDataDtos[j].ParentGroupName + "/" +
				tzDataCol.tzDataDtos[j].GroupName + "/" +
				tzDataCol.tzDataDtos[j].TzSortValue

			return strI < strJ
		}
	}

	sort.Slice(tzDataCol.tzDataDtos, less)

	return
}