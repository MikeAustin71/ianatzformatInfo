package tzdatastructs

import (
	errors2 "errors"
	"fmt"
	"strconv"
	"time"
)

type TimeZoneStatsDto struct {
	IanaVersion                  string
	NumOfLinkConflictResolved    int
	NumOfBackZoneConflicts       int
	NumIanaStdTZones             int
	NumIanaLinkTZones            int
	TotalIanaStdTzLinkZones      int
	NumMilitaryTZones            int
	NumOtherTZones               int
	TotalZones                   int
	NumMajorTZoneGroups          int
	NumMajorLinkGroups           int
	NumMajorMilitaryGroups       int
	NumMajorOtherGroups          int
	TotalMajorGroups             int
	NumLevel2StdSubTZoneGroups   int
	NumLevel3StdSubTZoneGroups   int
	TotalSubTZoneGroups          int
	NumLevel2LinkSubGroups       int
	NumLevel3LinkSubGroups       int
	TotalLinkSubGroups           int
	TotalSubGroups                int
	NumLevel1TZoneCollections     int
	NumLevel2TZoneCollections     int
	TotalTimeZoneCollections      int
	NumLevel1LinkZoneCollections  int
	NumLevel2LinkZoneCollections  int
	TotalLinkZoneCollections      int
	TotalZoneCollections          int
	IanaTzRegions                 []string
	IanaTzCounters                []int
	IanaLinkCounters              []int
	IanaTotalTimeZoneLinkCounters []int
	IanaTotalTimeZones            int
	IanaTotalLinks                int
	IanaTotalTimeZonesLinks       int
	IanaCapturedTimeZones         TimeZoneDataCollection
	CapturedMilitaryZones         TimeZoneDataCollection
	TzGroups                      []TimeZoneGroupCollection
	TzData                        []TimeZoneDataCollection
	TzAbbreviations               TimeZoneAbbreviationCollection
	MapTzAbbrvsToTimeZones        map[string][]string
	// Index Abbreviation+Utc Offset to Time Zone Canonical Value
	MapTimeZonesToTzAbbrvs        map[string][]string
	// Index TimeZone Canonical Value to Time Zone Abbreviation+Utc Offset
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
	tzStats.NumOtherTZones = 1
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
	tzStats.TotalLinkZoneCollections = 0
	tzStats.TotalZoneCollections = 0

	tzStats.IanaTzCounters = make([]int, lenWorldRegions)
	tzStats.IanaLinkCounters = make([]int, lenWorldRegions)
	tzStats.IanaTotalTimeZoneLinkCounters = make([]int, lenWorldRegions)
	tzStats.IanaTotalTimeZones = 0
	tzStats.IanaTotalLinks = 0
	tzStats.IanaTotalTimeZonesLinks = 0
	tzStats.IanaCapturedTimeZones = TimeZoneDataCollection{}.New()
	tzStats.CapturedMilitaryZones = TimeZoneDataCollection{}.New()
	tzStats.TzAbbreviations = TimeZoneAbbreviationCollection{}.New()
	tzStats.TzGroups = make([]TimeZoneGroupCollection, Level_03_Idx + 1)
	tzStats.TzData = make([]TimeZoneDataCollection, Level_03_Idx + 1)

	for i:=0; i <= Level_03_Idx; i++ {
		tzStats.TzGroups[i] = TimeZoneGroupCollection{}.New()
		tzStats.TzData[i] = TimeZoneDataCollection{}.New()
	}

	tzStats.MapTzAbbrvsToTimeZones = make(map[string][]string)

	tzStats.MapTimeZonesToTzAbbrvs = make(map[string][]string)
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

	tzDataDto.WorldRegionSortCode = -1

	for ; idx < lenRegions; idx++ {

	//	if strings.HasPrefix(
	//		tzDataDto.TzCanonicalValue, tzStats.IanaTzRegions[idx]) {
		if tzDataDto.ParentGroupName == tzStats.IanaTzRegions[idx] ||
			tzDataDto.GroupName == tzStats.IanaTzRegions[idx] {

			tzDataDto.WorldRegionSortCode = idx
			goto storeCapturedTimeZones
		}
	}

	idx = lenRegions
	tzDataDto.WorldRegionSortCode = idx

	storeCapturedTimeZones:
	
	tzDataDto.ArrayStorageLevel = zoneLevel

	var testLocation *time.Location

	testLocation, err = time.LoadLocation(tzDataDto.TzCanonicalValue)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by time.LoadLocation(tzDataDto.TzCanonicalValue).\n"+
			"tzLocation='%v'\n"+
			"Error:'%v'\n", tzDataDto.TzCanonicalValue, err.Error())
	}

	testTime := time.Now().In(testLocation)

	testTimeStr := testTime.Format("2006-01-02 15:04:05 -0700 MST")
	lenLeadStr := len("2006-01-02 15:04:05 ")

	tzDataDto.UtcOffset = "UTC" + testTimeStr[lenLeadStr: lenLeadStr + 5]

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

// CountMajorOtherNonIanaTimeZoneGroup - Counts, processes
// and stores information for a Major Other Time Zone
// Group.
//
// This is a level-1 Major Group (Level_01_Idx)
func (tzStats *TimeZoneStatsDto) CountMajorOtherNonIanaTimeZoneGroup(
	tzGroup TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountMajorOtherNonIanaTimeZoneGroup() "

	err := tzStats.TzGroups[Level_01_Idx].Add(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzGroups[zoneLevel].Add(tzGroup) Error\n" +
			"FileName: %v\n" +
			"Error: %v\n",
			tzGroup.SourceFileNameExt, err.Error() )
	}

	tzStats.NumMajorOtherGroups++

	return nil
}

// CountOtherNonIanaTimeZone - Counts processes and stores
// information for a standard Other Time Zone Type.
func (tzStats *TimeZoneStatsDto) CountOtherNonIanaTimeZone(
	tzDataDto TimeZoneDataDto, zoneLevel int, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountOtherNonIanaTimeZone() "

	if zoneLevel < 0 || zoneLevel >= len(tzStats.TzData) {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter 'zoneLevel' is out-of-range.\n" +
			"Valid range for zoneLevel is 0-%v.\n" +
			"zoneLevel='%v'\n", len(tzStats.TzData) - 1, zoneLevel)
	}

	tzDataDto.ArrayStorageLevel = zoneLevel

	err := tzStats.TzData[zoneLevel].Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by tzStats.TzData[zoneLevel].Add(tzDataDto)\n" +
			"zoneLevel='%v'\n" +
			"Error='%v'\n", zoneLevel, err.Error())
	}

	tzStats.NumOtherTZones++

	return nil
}

// CountMajorMilitaryTimeZoneGroup - - Counts, processes and stores
//// information for a Major Military Time Zone Group.
//// This is a level-1 Major Group (Level_01_Idx)
func (tzStats *TimeZoneStatsDto) CountMajorMilitaryTimeZoneGroup(
	tzGroup TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountMilitaryZone() "

	err := tzStats.TzGroups[Level_01_Idx].Add(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzStats.TzGroups[zoneLevel].Add(tzGroup) Error\n" +
			"FileName: %v\n" +
			"Error: %v\n",
			tzGroup.SourceFileNameExt, err.Error() )
	}

	tzStats.NumMajorMilitaryGroups++

	return nil
}

// CountMilitaryZone - Counts processes and stores
// information for a standard Military Time Zone.
func (tzStats *TimeZoneStatsDto) CountMilitaryZone(
	tzDataDto TimeZoneDataDto, zoneLevel int, ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.CountMilitaryZone() "

	if zoneLevel < 0 || zoneLevel >= len(tzStats.TzData) {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter 'zoneLevel' is out-of-range.\n" +
			"Valid range for zoneLevel is 0-%v.\n" +
			"zoneLevel='%v'\n", len(tzStats.TzData) - 1, zoneLevel)
	}

	tzDataDto.ArrayStorageLevel = zoneLevel

	utcEquivalent, ok := MilitaryUTCMap[tzDataDto.TzName]

	if !ok {
		return fmt.Errorf(ePrefix +
			"Error: MilitaryUTCMap[tzDataDto.TzName] FAILED!\n" +
			"tzDataDto.TzName='%v'", tzDataDto.TzName)
	}
		// "Whiskey":  "UTC-10",
	utcPrefix := utcEquivalent[0:4]

	var utcHours int
	var err error

	utcHours, err = strconv.Atoi(utcEquivalent[4:])

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by strconv.Atoi(utcEquivalent[4:])\n" +
			"utcEquivalent[4:]='%v'\n" +
			"Error='%v'\n", utcEquivalent[4:], err.Error())
	}

	tzDataDto.UtcOffset = fmt.Sprintf(utcPrefix + "%02d00", utcHours )

	err = tzStats.CapturedMilitaryZones.Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by tzStats.CapturedMilitaryZones.Add(tzDataDto)\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzStats.TzData[zoneLevel].Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by tzStats.TzData[zoneLevel].Add(tzDataDto)\n" +
			"zoneLevel='%v'\n" +
			"Error='%v'\n", zoneLevel, err.Error())
	}

	tzStats.NumMilitaryTZones++


	return nil
}

// RunTotals - Computes all statistical totals and sub totals.
func (tzStats *TimeZoneStatsDto) RunTotals(ePrefix string) error {

	ePrefix += "TimeZoneStatsDto.RunTotals() "

	tzStats.TotalIanaStdTzLinkZones =
		tzStats.NumIanaStdTZones +
		tzStats.NumIanaLinkTZones

	tzStats.TotalZones =
		tzStats.TotalIanaStdTzLinkZones +
		tzStats.NumMilitaryTZones +
		tzStats.NumOtherTZones

	tzStats.TotalMajorGroups =
		tzStats.NumMajorTZoneGroups +
		tzStats.NumMajorLinkGroups +
		tzStats.NumMajorMilitaryGroups +
		tzStats.NumMajorOtherGroups

	tzStats.TotalSubTZoneGroups =
		tzStats.NumLevel2StdSubTZoneGroups +
		tzStats.NumLevel3StdSubTZoneGroups

	tzStats.TotalLinkSubGroups =
		tzStats.NumLevel2LinkSubGroups +
		tzStats.NumLevel3LinkSubGroups

	tzStats.TotalSubGroups =
		tzStats.TotalSubTZoneGroups +
		tzStats.TotalLinkSubGroups

	tzStats.TotalTimeZoneCollections =
		tzStats.NumLevel1TZoneCollections +
		tzStats.NumLevel2TZoneCollections

	tzStats.TotalLinkZoneCollections =
		tzStats.NumLevel1LinkZoneCollections +
		tzStats.NumLevel2LinkZoneCollections

	tzStats.TotalZoneCollections =
		tzStats.TotalTimeZoneCollections +
			tzStats.TotalLinkZoneCollections

	lenWorldRegions := len(WorldRegions)

	if len(tzStats.IanaTzRegions) != lenWorldRegions {

		return fmt.Errorf(ePrefix +
			"Error: Lenght of tzStats.IanaTzRegions is invalid.\n" +
			"tzStats.IanaTzRegions should equal '%v'.\n" +
			"Instead, tzStats.IanaTzRegions equals '%v'\n",
			lenWorldRegions, len(tzStats.IanaTzRegions))
	}

	if len(tzStats.IanaTzCounters) != lenWorldRegions {
		return fmt.Errorf(ePrefix +
			"Error: Lenght of tzStats.IanaTzCounters is invalid.\n" +
			"tzStats.IanaTzCounters should equal '%v'.\n" +
			"Instead, tzStats.IanaTzCounters equals '%v'\n",
			lenWorldRegions, len(tzStats.IanaTzCounters))
	}

	if len(tzStats.IanaLinkCounters) != lenWorldRegions {
		return fmt.Errorf(ePrefix +
			"Error: Lenght of tzStats.IanaLinkCounters is invalid.\n" +
			"tzStats.IanaLinkCounters should equal '%v'.\n" +
			"Instead, tzStats.IanaLinkCounters equals '%v'\n",
			lenWorldRegions, len(tzStats.IanaLinkCounters))
	}

	if len(tzStats.IanaTotalTimeZoneLinkCounters) != lenWorldRegions {
		return fmt.Errorf(ePrefix +
			"Error: Lenght of tzStats.IanaTotalTimeZoneLinkCounters is invalid.\n" +
			"tzStats.IanaTotalTimeZoneLinkCounters should equal '%v'.\n" +
			"Instead, tzStats.IanaTotalTimeZoneLinkCounters equals '%v'\n",
			lenWorldRegions, len(tzStats.IanaTotalTimeZoneLinkCounters))
	}

	tzStats.IanaTotalTimeZones = 0
	tzStats.IanaTotalLinks = 0
	tzStats.IanaTotalTimeZonesLinks = 0

	for i:=0; i < lenWorldRegions; i++ {
		tzStats.IanaTotalTimeZoneLinkCounters[i] =
		tzStats.IanaTzCounters[i] +
		tzStats.IanaLinkCounters[i]

		tzStats.IanaTotalTimeZones +=
			tzStats.IanaTzCounters[i]

		tzStats.IanaTotalLinks +=
			tzStats.IanaLinkCounters[i]

		tzStats.IanaTotalTimeZonesLinks +=
			tzStats.IanaTzCounters[i]

		tzStats.IanaTotalTimeZonesLinks +=
			tzStats.IanaLinkCounters[i]

	}

	return nil
}
