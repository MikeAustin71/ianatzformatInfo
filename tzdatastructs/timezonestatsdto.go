package tzdatastructs

import (
	errors2 "errors"
	"fmt"
	"strings"
)

type TimeZoneStatsDto struct {
	IanaVersion                string
	NumOfLinkConflictResolved  int
	NumOfBackZoneConflicts     int
	NumIanaStdTZones           int
	NumIanaLinkTZones          int
	TotalIanaStdTzLinkZones    int
	NumMilitaryTZones          int
	NumOtherTZones             int
	TotalZones                 int
	NumMajorTZoneGroups        int
	NumMajorLinkGroups         int
	NumMajorMilitaryGroups     int
	NumMajorOtherGroups        int
	TotalMajorGroups           int
	NumLevel2StdSubTZoneGroups int
	NumLevel3StdSubTZoneGroups int
	TotalSubTZoneGroups          int
	NumLevel2LinkSubGroups       int
	NumLevel3LinkSubGroups       int
	TotalLinkSubGroups           int
	TotalSubGroups               int
	NumLevel1TZoneCollections    int
	NumLevel2TZoneCollections    int
	TotalTimeZoneCollections     int
	NumLevel1LinkZoneCollections int
	NumLevel2LinkZoneCollections int
	NumOfLinkZoneCollections     int
	IanaTzRegions                []string
	IanaTzCounters               []int
	IanaLinkCounters             []int
	IanaTotalTzLinkCounters      []int
	IanaTotalTimeZones           int
	IanaTotalLinks               int
	IanaTotalTimeZonesLinks      int
	IanaCapturedTimeZones        TimeZoneDataCollection
	IanaCapturedLinkZones        TimeZoneDataCollection
	TzGroups []TimeZoneGroupCollection
	TzData []TimeZoneDataCollection
	TzLinks []TimeZoneDataCollection
}

func (tzStats *TimeZoneStatsDto) Initialize() {

	lenWorldRegions := len(WorldRegions)

	tzStats.IanaTzRegions = make([]string, lenWorldRegions)

	for i:=0; i < lenWorldRegions; i++ {
		tzStats.IanaTzRegions[i] = WorldRegions[i]
	}

	tzStats.IanaVersion = ""
	tzStats.NumOfLinkConflictResolved = 0
	tzStats.NumOfBackZoneConflicts = 0
	tzStats.NumIanaStdTZones = 0
	tzStats.NumIanaLinkTZones = 0
	tzStats.TotalIanaStdTzLinkZones = 0
	tzStats.NumMilitaryTZones = 0
	tzStats.NumOtherTZones = 0
	tzStats.TotalZones = 0
	tzStats.NumMajorTZoneGroups = 0
	tzStats.NumMajorLinkGroups = 0
	tzStats.NumMajorMilitaryGroups = 0
	tzStats.NumMajorOtherGroups = 0
	tzStats.TotalMajorGroups = 0
	tzStats.NumLevel2StdSubTZoneGroups = 0
	tzStats.NumLevel3StdSubTZoneGroups = 0
	tzStats.TotalSubTZoneGroups = 0
	tzStats.NumLevel2LinkSubGroups = 0
	tzStats.NumLevel3LinkSubGroups = 0
	tzStats.TotalLinkSubGroups = 0
	tzStats.TotalSubGroups = 0
	tzStats.NumLevel1TZoneCollections = 0
	tzStats.NumLevel2TZoneCollections = 0
	tzStats.TotalTimeZoneCollections = 0
	tzStats.NumLevel1LinkZoneCollections = 0
	tzStats.NumLevel2LinkZoneCollections = 0
	tzStats.NumOfLinkZoneCollections = 0

	tzStats.IanaTzCounters = make([]int, lenWorldRegions)
	tzStats.IanaLinkCounters = make([]int, lenWorldRegions)
	tzStats.IanaTotalTzLinkCounters = make([]int, lenWorldRegions)
	tzStats.IanaTotalTimeZones = 0
	tzStats.IanaTotalLinks = 0
	tzStats.IanaTotalTimeZonesLinks = 0
	tzStats.IanaCapturedTimeZones = TimeZoneDataCollection{}.New()
	tzStats.IanaCapturedLinkZones = TimeZoneDataCollection{}.New()
	tzStats.TzGroups = make([]TimeZoneGroupCollection, Level_03_Idx + 1)
	tzStats.TzData = make([]TimeZoneDataCollection, Level_03_Idx + 1)
	tzStats.TzLinks = make([]TimeZoneDataCollection, Level_03_Idx + 1)

	for i:=0; i <= Level_03_Idx; i++ {
		tzStats.TzGroups[i] = TimeZoneGroupCollection{}.New()
		tzStats.TzData[i] = TimeZoneDataCollection{}.New()
		tzStats.TzLinks[i] = TimeZoneDataCollection{}.New()
	}

}

// CountMajorTimeZoneGroup  - Counts, processes and stores
// information for a Iana Major Time Zone Group.
// This is a level-1 Major Group (Level_01_Idx)
func (tzStats *TimeZoneStatsDto) CountMajorTimeZoneGroup(
	tzGroup TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountMajorTimeZoneGroup() "

	err := tzStats.TzGroups[Level_01_Idx].Add(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzGroups[zoneLevel].Add(tzGroup) Error\n" +
			"FileName: %v\n" +
			"Error: %v\n",
			tzGroup.SourceFileNameExt, err.Error() )
	}

	tzStats.NumMajorTZoneGroups++

	return nil
}

// CountLevel2StdSubGroup  - Counts, processes and stores
// information for a Iana Level-1 Standard Sub-Time Zone
// Group. A Sub-Time Group is a place-holder collection
// which contains a collection of standard time zones.
func (tzStats *TimeZoneStatsDto) CountLevel2StdSubGroup(
	tzGroup TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountLevel2StdSubGroup() "

	err := tzStats.TzGroups[Level_02_Idx].Add(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzGroups[zoneLevel].Add(tzGroup) Error\n" +
			"FileName: %v\n" +
			"Error: %v\n",
			tzGroup.SourceFileNameExt, err.Error() )
	}

	tzStats.NumLevel2StdSubTZoneGroups++

	return nil
}

// CountLevel3StdSubGroup  - Counts, processes and stores
// information for a Iana Level-1 Standard Sub-Time Zone
// Group. A Sub-Time Group is a place-holder collection
// which contains a collection of standard time zones.
func (tzStats *TimeZoneStatsDto) CountLevel3StdSubGroup(
	tzGroup TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountLevel3StdSubGroup() "

	err := tzStats.TzGroups[Level_03_Idx].Add(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzGroups[zoneLevel].Add(tzGroup) Error\n" +
			"FileName: %v\n" +
			"Error: %v\n",
			tzGroup.SourceFileNameExt, err.Error() )
	}

	tzStats.NumLevel3StdSubTZoneGroups++

	return nil
}

// CountIanaStdZone  - Counts processes and stores
// information for a standard Iana Time Zone.
func (tzStats *TimeZoneStatsDto)CountIanaStdZone(
	tzDataDto TimeZoneDataDto, zoneLevel int, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountIanaStdZone() "

	lenRegions := len(tzStats.IanaTzRegions) - 1

	if lenRegions < 1 {
		return errors2.New (ePrefix +
			"Error TimeZoneStatsDto Improperly Initialized!\n" +
			"")
	}

	if zoneLevel < 0 || zoneLevel >= len(tzStats.TzData) {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter 'zoneLevel' is out-of-range.\n" +
			"Valid range for zoneLevel is 0-%v\n" +
			"zoneLevel='%v'\n", len(tzStats.TzData) - 1, zoneLevel)
	}

	var err error

	idx := 0

	for ; idx < lenRegions; idx++ {

		if strings.HasPrefix(
			tzDataDto.TzCanonicalValue, tzStats.IanaTzRegions[idx]) {

			goto storeCapturedTimeZones
		}
	}

	idx = lenRegions

	storeCapturedTimeZones:
	
		tzDataDto.ArrayStorageLevel = zoneLevel
	
	// Other IANA Time Zone
	err = tzStats.IanaCapturedTimeZones.Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by tzStats.IanaCapturedTimeZones.Add(tzDataDto)\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzStats.TzData[zoneLevel].Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by tzStats.TzData[zoneLevel].Add(tzDataDto)\n" +
			"zoneLevel='%v'\n" +
			"Error='%v'\n", zoneLevel, err.Error())
	}

	tzStats.IanaTzCounters[idx]++
	tzStats.NumIanaStdTZones++

	return nil
}

// CountLevel1TimeZoneCollection  - Counts processes and stores
// information for a Level-1 Iana Time Zone Collections.
// Time Zone Collections represent a collection of standard
// Iana Time Zones.
func (tzStats *TimeZoneStatsDto) CountLevel1TimeZoneCollection(
	tzDataDto TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountLevel1TimeZoneCollection() "

	tzDataDto.ArrayStorageLevel = Level_01_Idx

	err := tzStats.TzData[Level_01_Idx].Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzData[zoneLevel].Add(tzDataDto) Error\n" +
			"FileName: '%v'\n" +
			"Error: '%v'", tzDataDto.SourceFileNameExt ,err.Error())
	}

	tzStats.NumLevel1TZoneCollections++

	return nil
}

// CountLevel2TimeZoneCollection  - Counts processes and stores
// information for a Level-2 Iana Time Zone Collections.
// Time Zone Collections represent a collection of standard
// Iana Time Zones.
func (tzStats *TimeZoneStatsDto) CountLevel2TimeZoneCollection(
	tzDataDto TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountLevel2TimeZoneCollection() "

	tzDataDto.ArrayStorageLevel = Level_02_Idx

	err := tzStats.TzData[Level_02_Idx].Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzData[zoneLevel].Add(tzDataDto) Error\n" +
			"FileName: '%v'\n" +
			"Error: '%v'", tzDataDto.SourceFileNameExt ,err.Error())
	}

	tzStats.NumLevel2TZoneCollections++

	return nil
}


// CountMajorLinkZoneGroup  - Counts, processes and stores
// information for a Iana Major Time Zone Group.
// This is a level-1 Major Group (Level_01_Idx)
func (tzStats *TimeZoneStatsDto) CountMajorLinkZoneGroup(
	tzGroup TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountMajorLinkZoneGroup() "

	err := tzStats.TzGroups[Level_01_Idx].Add(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzGroups[zoneLevel].Add(tzGroup) Error\n" +
			"FileName: %v\n" +
			"Error: %v\n",
			tzGroup.SourceFileNameExt, err.Error() )
	}

	tzStats.NumMajorLinkGroups++

	return nil
}

// CountLevel2LinkSubGroup  - Counts, processes and stores
// information for a Iana Level-2 Link Sub-Group.
//
func (tzStats *TimeZoneStatsDto) CountLevel2LinkSubGroup(
	tzGroup TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountLevel2LinkSubGroup() "

	err := tzStats.TzGroups[Level_02_Idx].Add(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzGroups[zoneLevel].Add(tzGroup) Error\n" +
			"FileName: %v\n" +
			"Error: %v\n",
			tzGroup.SourceFileNameExt, err.Error() )
	}

	tzStats.NumLevel2LinkSubGroups++

	return nil
}

// CountLevel3LinkSubGroup  - Counts, processes and stores
// information for a Iana Level-3 Link Sub- Group. 
func (tzStats *TimeZoneStatsDto) CountLevel3LinkSubGroup(
	tzGroup TimeZoneGroupDto, ePrefix string) error {

		ePrefix += "TimeZoneStatsDto.CountLevel3LinkSubGroup() "
		
	err := tzStats.TzGroups[Level_03_Idx].Add(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzGroups[zoneLevel].Add(tzGroup) Error\n" +
			"FileName: %v\n" +
			"Error: %v\n",
			tzGroup.SourceFileNameExt, err.Error() )
	}

	tzStats.NumLevel3LinkSubGroups++

	return nil
}


// CountIanaLinkZone  - Counts processes and stores
// information for a Link Iana Time Zone.
func (tzStats *TimeZoneStatsDto)CountIanaLinkZone(
	tzDataDto TimeZoneDataDto, zoneLevel int, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountIanaLinkZone() "

	lenRegions := len(tzStats.IanaTzRegions) - 1

	if lenRegions < 1 {
		return errors2.New (ePrefix +
			"Error TimeZoneStatsDto Improperly Initialized!\n" +
			"")
	}

	if zoneLevel < 0 || zoneLevel >= len(tzStats.TzData) {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter 'zoneLevel' is out-of-range.\n" +
			"Valid range for zoneLevel is 0-%v\n" +
			"zoneLevel='%v'\n", len(tzStats.TzData) - 1, zoneLevel)
	}

	var err error

	idx := 0

	for ; idx < lenRegions; idx++ {

		if strings.HasPrefix(
			tzDataDto.TzCanonicalValue, tzStats.IanaTzRegions[idx]) {

			goto storeCapturedTimeZones
		}
	}

	idx = lenRegions

storeCapturedTimeZones:

	tzDataDto.ArrayStorageLevel = zoneLevel
	
	err = tzStats.IanaCapturedLinkZones.Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by tzStats.IanaCapturedLinkZones.Add(tzDataDto)\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzStats.TzData[zoneLevel].Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by tzStats.TzData[zoneLevel].Add(tzDataDto)\n" +
			"zoneLevel='%v'\n" +
			"Error='%v'\n", zoneLevel, err.Error())
	}

	tzStats.IanaLinkCounters[idx]++
	tzStats.NumIanaLinkTZones++

	return nil
}

// CountLevel1LinkZoneCollection  - Counts processes and stores
// information for a Level-1 Iana Time Zone Collections.
// Time Zone Collections represent a collection of standard
// Iana Time Zones.
func (tzStats *TimeZoneStatsDto) CountLevel1LinkZoneCollection(
	tzDataDto TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountLevel1LinkZoneCollection() "

	tzDataDto.ArrayStorageLevel = Level_01_Idx

	err := tzStats.TzData[Level_01_Idx].Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzData[zoneLevel].Add(tzDataDto) Error\n" +
			"FileName: '%v'\n" +
			"Error: '%v'", tzDataDto.SourceFileNameExt ,err.Error())
	}

	tzStats.NumLevel1LinkZoneCollections++

	return nil
}

// CountLevel2LinkZoneCollection  - Counts processes and stores
// information for a Level-2 Iana Time Zone Collections.
// Time Zone Collections represent a collection of standard
// Iana Time Zones.
func (tzStats *TimeZoneStatsDto) CountLevel2LinkZoneCollection(
	tzDataDto TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountLevel2LinkZoneCollection() "

	tzDataDto.ArrayStorageLevel = Level_02_Idx

	err := tzStats.TzData[Level_02_Idx].Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzData[zoneLevel].Add(tzDataDto) Error\n" +
			"FileName: '%v'\n" +
			"Error: '%v'", tzDataDto.SourceFileNameExt ,err.Error())
	}

	tzStats.NumLevel2LinkZoneCollections++

	return nil
}
