package tzdatastructs

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// SortTimeZoneDataByMjrGrpTzName - Sort by MajorGroup Name, TzName
//
// Example Usage:
//    sort.Sort(SortByTzMajorGroupName(tzMajorGroupDtoArray))
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

	if sortByTzDtoName[i].MajorGroup == sortByTzDtoName[j].MajorGroup {
		return sortByTzDtoName[i].TzName < sortByTzDtoName[j].TzName
	}

	return sortByTzDtoName[i].MajorGroup < sortByTzDtoName[j].MajorGroup
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

// AddIfNewByDetail - Adds a new TimeZoneDataDto to the collection. The
// TimeZoneDataDto instance to be added is first created using the detail
// input parameters. 
//
func (tzDataCol *TimeZoneDataCollection) AddIfNewByDetail(
	majorGroup,
	subTzName,
	tzName,
	tzValue,
	srcFileNameExt string,
	tzClass TimeZoneClass,
	deprecationStatus TimeZoneDeprecationStatus) (isNew bool, err error) {

	ePrefix := "TimeZoneDataCollection.AddIfNewByDetail() "
	isNew = false
	err = nil

	tzDataDto, err2 := TimeZoneDataDto{}.New(
		majorGroup,
		subTzName,
		tzName,
		tzValue,
		srcFileNameExt,
		tzClass,
		deprecationStatus)

	if err2 != nil {
		err = fmt.Errorf(ePrefix +
			"%v\n", err2.Error())
		return isNew, err
	}

	return tzDataCol.AddIfNew(tzDataDto)
}

// MajorGroupExists - Performs a search for on the internal TimeZoneDataDto
// array for a match on TimeZoneDataDto.MajorGroup. If the search is successful,
// this method returns a boolean value of 'true' and the integer index
// value of the found TimeZoneDataDto instance.
//
// If the search fails, a boolean value of 'false' is returned and the
// integer index value is set to -1.
//
// If the input parameter 'caseInsensitiveSearch' is set to 'true', the search for
// TimeZoneDataDto.MajorGroup will be conducted as a case insensitive search.
// This means that both strings are converted to lower case before the comparison
// is performed.
//
// If the input parameter 'caseInsensitiveSearch' is set to 'false', the search
// is conducted as a case sensitive comparison where upper and lower case
// characters are significant.
//
func (tzDataCol *TimeZoneDataCollection) MajorGroupExists(
	majorGroupName string, caseInsensitiveSearch bool) (majorGrpExists bool, index int) {

	majorGrpExists = false
	index = -1

	if tzDataCol.tzDataDtos == nil {

		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)

	}

	lenTzDataDtoArray := len(tzDataCol.tzDataDtos)

	if lenTzDataDtoArray == 0 {
		return false, -1
	}

	if caseInsensitiveSearch {
		majorGroupName = strings.ToLower(majorGroupName)
	}

	for i:=0; i < lenTzDataDtoArray; i++ {

		if caseInsensitiveSearch {
			if majorGroupName == strings.ToLower(tzDataCol.tzDataDtos[i].MajorGroup) {
				return true, i
			}
		} else if majorGroupName == tzDataCol.tzDataDtos[i].MajorGroup {
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

	lenTzDataDtos := len(tzDataCol.tzDataDtos)

	if lenTzDataDtos == 0 {
		return TimeZoneDataDto{}, fmt.Errorf(ePrefix +
			"ERROR: The Time Zone Data Collection is EMPTY!")
	}

	if index < 0 {
		return TimeZoneDataDto{},
			fmt.Errorf(ePrefix +
				"ERROR: Input parameter 'index' is less than zero and INVALID!\n" +
				"index='%v'", index)
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

// SortByMjrGrpTzName - Sort the collection by TimeZone Major Group and
// Time Zone Name.
func (tzDataCol *TimeZoneDataCollection) SortByMjrGrpTzName(caseSensitiveSort bool) {

	if tzDataCol.tzDataDtos == nil {
		tzDataCol.tzDataDtos = make([]TimeZoneDataDto, 0, 500)
	}

	if len(tzDataCol.tzDataDtos) == 0 {
		return
	}

	var less func(i, j int) bool

	if !caseSensitiveSort {
		less = func(i, j int) bool {
			if strings.ToLower(tzDataCol.tzDataDtos[i].MajorGroup) !=
				strings.ToLower(tzDataCol.tzDataDtos[j].MajorGroup) {

				return strings.ToLower(tzDataCol.tzDataDtos[i].MajorGroup) <
					strings.ToLower(tzDataCol.tzDataDtos[j].MajorGroup)
			}

			return strings.ToLower(tzDataCol.tzDataDtos[i].TzName) <
				strings.ToLower(tzDataCol.tzDataDtos[j].TzName)
		}
	} else {
		less = func(i, j int) bool {
			if tzDataCol.tzDataDtos[i].MajorGroup !=
				tzDataCol.tzDataDtos[j].MajorGroup {

				return tzDataCol.tzDataDtos[i].MajorGroup <
					tzDataCol.tzDataDtos[j].MajorGroup
			}

			return tzDataCol.tzDataDtos[i].TzName <
				tzDataCol.tzDataDtos[j].TzName
		}
	}

	sort.Slice(tzDataCol.tzDataDtos, less)

	return
}