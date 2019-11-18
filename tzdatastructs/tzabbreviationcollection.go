package tzdatastructs

import (
	"errors"
	"fmt"
	"sort"
)

type TimeZoneAbbreviationCollection struct {
	tzAbbreviations []TzAbbreviationDto
}


// AddIfNew - Add a new TzAbbreviationDto to the collection
// provided it does not previously exist in the collection
func (tzAbbrv *TimeZoneAbbreviationCollection) AddIfNew(
	tzAbbreviation TzAbbreviationDto) {

		if tzAbbrv.tzAbbreviations == nil {

			tzAbbrv.tzAbbreviations = make([]TzAbbreviationDto, 1)

			tzAbbrv.tzAbbreviations[0] = tzAbbreviation.CopyOut()

			return
		}

		lenAbrvs := len(tzAbbrv.tzAbbreviations)

		if lenAbrvs == 0 {
			tzAbbrv.tzAbbreviations = append(tzAbbrv.tzAbbreviations, tzAbbreviation.CopyOut())
			return
		}

		for i:=0; i < lenAbrvs; i++ {
			if tzAbbrv.tzAbbreviations[i].Id == tzAbbreviation.Id {
				return
			}
		}

	tzAbbrv.tzAbbreviations = append(tzAbbrv.tzAbbreviations, tzAbbreviation.CopyOut())

	return
}

// New - Returns a new properly initialized instance of a TimeZoneAbbreviationCollection
func (tzAbbrv TimeZoneAbbreviationCollection) New() (TimeZoneAbbreviationCollection) {

	newAbbrvDto := TimeZoneAbbreviationCollection{}
	newAbbrvDto.tzAbbreviations = make([]TzAbbreviationDto, 0)

	return newAbbrvDto
}

// PeekPtr - Returns a pointer to the TzAbbreviationDto object
// located in the internal tzAbbreviations array at the index
// specified by input parameter 'index'.
func (tzAbbrv *TimeZoneAbbreviationCollection) PeekPtr(
	index int) (*TzAbbreviationDto, error) {

	ePrefix := "TimeZoneAbbreviationCollection.PeekPtr() "

	if tzAbbrv.tzAbbreviations == nil {
		tzAbbrv.tzAbbreviations = make([]TzAbbreviationDto, 0)
	}

	if index < 0 {
		return &TzAbbreviationDto{},
			fmt.Errorf(ePrefix +
				"Error: Input parameter 'index' is less than Zero!\n" +
				"index='%v'\n", index)
	}

	lenTzAbbrvDtos := len(tzAbbrv.tzAbbreviations)

	if lenTzAbbrvDtos == 0 {
		return &TzAbbreviationDto{},
			errors.New(ePrefix +
				"Error: The Time Zone Abbreviation Collection is EMPTY!\n")
	}

	if index > (lenTzAbbrvDtos - 1) {
		return &TzAbbreviationDto{},
		fmt.Errorf(ePrefix +
			"Error: Input parameter 'index' is INVALID!\n" +
			"'index' exceeds the maximum index for this collection.\n" +
			"Collection Maximum Index='%v'\n" +
			"index='%v'\n", lenTzAbbrvDtos, index)
	}

	return &tzAbbrv.tzAbbreviations[index], nil
}

func (tzAbbrv *TimeZoneAbbreviationCollection) SortByAbbrv() {

	if tzAbbrv.tzAbbreviations == nil {
		tzAbbrv.tzAbbreviations = make([]TzAbbreviationDto, 0)
		return
	}

	if len(tzAbbrv.tzAbbreviations) < 2 {
		return
	}

	var less func(i, j int) bool

	less = func(i, j int) bool {

		if tzAbbrv.tzAbbreviations[i].Abbrv ==
			tzAbbrv.tzAbbreviations[j].Abbrv {

			return tzAbbrv.tzAbbreviations[i].UtcOffset <
				tzAbbrv.tzAbbreviations[j].UtcOffset

		} else {

			return tzAbbrv.tzAbbreviations[i].Abbrv <
				tzAbbrv.tzAbbreviations[j].Abbrv

		}
	}

	sort.Slice(tzAbbrv.tzAbbreviations, less)

	return
}