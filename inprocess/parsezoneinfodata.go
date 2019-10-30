package inprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"local.com/amarillomike/ianatzformatInfo/tzdeclarations"
	"strconv"
	"strings"
)

type ParseZoneInfoData struct {
	input string
	output string
}

// ParseZoneInfo - Parses 'zoneinfo' data and generates lists of time zones.
//
func (parseZInfo ParseZoneInfoData) ParseZoneInfo(
	zInfoDto ZoneInfoDataDto, ePrefix string) (tzdatastructs.TimeZoneStatsDto, error) {

	ePrefix += "ParseZoneInfoData.ParseZoneInfo()"

	tzStats := tzdatastructs.TimeZoneStatsDto{}

	zInfoDto.ZoneInfoDirFileInfo.SortByAbsPathFileName(true)

	err := parseZInfo.parseIanaTimeZoneFiles(zInfoDto, &tzStats, ePrefix)

	if err != nil {
		return tzStats, err
	}

	err = parseZInfo.configMilitaryTimeZones( &tzStats, ePrefix)

	return tzStats, nil
}


// configMilitaryTimeZones - Creates and stores Military Time Zones
//
func (parseZInfo ParseZoneInfoData) configMilitaryTimeZones(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "ParseIanaTzData.configMilitaryTimeZones() "

	// Configure Time Zone Level-1 Major Group
	// for Military Time Zones
	tzGroup := tzdatastructs.TimeZoneGroupDto{}
	tzGroup.ParentGroupName = ""
	tzGroup.GroupName = "Military"
	tzGroup.GroupSortValue = "Military"

	// Example: 'militaryTimeZones'
	tzGroup.TypeName = "military" +
		tzdatastructs.MasterGroupTypeSuffix

	tzGroup.TypeValue = "string"

	// Example: 'America'
	tzGroup.IanaVariableName = "Military"

	tzGroup.SourceFileNameExt = "None"
	tzGroup.GroupType = tzdatastructs.TzGrpType.Standard()
	tzGroup.GroupClass = tzdatastructs.TzGrpClass.Artificial()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzGroup.SetIsInitialized(true)

	err := tzdeclarations.TzMilitaryDeclarations{}.MilitaryTypeDeclaration(&tzGroup, ePrefix)

	if err != nil {
		return err
	}

	err = tzStats.CountMajorMilitaryTimeZoneGroup(tzGroup, ePrefix)

	if err != nil {
		return err
	}

	for i:=0; i < len(tzdatastructs.MilitaryTzArray); i++ {

		// Configure Standard Level-1 Iana Time Zone Data Dto
		// For Military Time Zone
		tzDataDto := tzdatastructs.TimeZoneDataDto{}
		tzDataDto.ParentGroupName = ""
		tzDataDto.GroupName = "Military" // Military - majorGroup
		tzDataDto.TzName = tzdatastructs.MilitaryTzArray[i] // Alpha - tzName
		tzDataDto.TzAliasValue = ""
		canonicalValue, ok := tzdatastructs.MilitaryTzMap[tzdatastructs.MilitaryTzArray[i]]

		if !ok {
			return fmt.Errorf(ePrefix +
				"tzdatastructs.MilitaryTzMap[] look-up error. Military Time Zone missing!\n" +
				"Military Time Zone: %v\n", tzdatastructs.MilitaryTzArray[i])
		}

		tzDataDto.TzCanonicalValue = canonicalValue
		tzDataDto.TzValue = tzDataDto.TzCanonicalValue // 'Etc/GMT-1'
		tzDataDto.TzSortValue =
			tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzdatastructs.MilitaryTzArray[i])

		// Example func signature
		// func (milTz militaryTimeZones) Alpha() string { return "Etc/GMT+1" }

		// Example: militaryTimeZones
		tzDataDto.FuncType =
			"military" +
				tzdatastructs.MasterGroupTypeSuffix

		// Example: 'milTz'
		tzDataDto.FuncSelfReferenceVariable = "milTz"

		// FuncName: Alpha()
		tzDataDto.FuncName = parseZInfo.zoneCfgValidFuncName(tzdatastructs.MilitaryTzArray[i])

		tzDataDto.FuncReturnType = "string"

		// Example Function Return Value = "Etc/GMT+1"
		tzDataDto.FuncReturnValue =
			fmt.Sprintf("\"%v\"", tzDataDto.TzCanonicalValue)

		tzDataDto.SourceFileNameExt = "None"
		tzDataDto.TzClass = tzdatastructs.TZClass.Artificial()
		tzDataDto.TzType = tzdatastructs.TZType.Standard()
		tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
		tzDataDto.SetIsInitialized(true)

		err = tzdeclarations.TzMilitaryDeclarations{}.MilitaryTzFuncDeclaration(&tzDataDto, ePrefix)

		if err != nil {
			return err
		}

		err = tzStats.CountMilitaryZone(tzDataDto, tzdatastructs.Level_01_Idx, ePrefix)

		if err != nil {
			return err
		}
	}

	return nil
}


// parseIanaTimeZoneFiles - Parses the directory and file structure
// of the 'zoneinfo' directory tree and generates IANA time zones.
//
func (parseZInfo ParseZoneInfoData) parseIanaTimeZoneFiles(
	zInfoDto ZoneInfoDataDto,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "ParseZoneInfoData.parseIanaTimeZoneFiles() "

	numOfZoneInfoFMgrs := zInfoDto.ZoneInfoDirFileInfo.GetNumOfFileMgrs()

	if numOfZoneInfoFMgrs < 30 {
		return fmt.Errorf(ePrefix+
			"Error: Number of ZoneInfo files is less than 30.\n"+
			"Total Zone Info files='%v'\n", numOfZoneInfoFMgrs)
	}

	topDir := zInfoDto.ZoneInfoDirMgr.GetAbsolutePathWithSeparator()

	lenTopDir := len(topDir)

	for i := 0; i < numOfZoneInfoFMgrs; i++ {

		fMgr, err := zInfoDto.ZoneInfoDirFileInfo.PeekFileMgrAtIndex(i)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error returned by zInfoDto.ZoneInfoDirFileInfo.PeekFileMgrAtIndex(i)\n"+
				"i='%v'\n"+
				"Error='%v'\n", i, err.Error())
		}

		absPathFileName := fMgr.GetAbsolutePathFileName()
		tZone := absPathFileName[lenTopDir:]
		tZone = strings.Replace(tZone, "\\", "/", -1)
		zoneArray := strings.Split(tZone, "/")

		switch len(zoneArray) {
		case 1:
			zoneArray = make([]string, 2)
			zoneArray[0] = "Other"
			zoneArray[1] = tZone
			err = parseZInfo.zoneConfigTwoElements(fMgr, zoneArray, tzStats, ePrefix)

			if err != nil {
				return err
			}
		case 2:
			err = parseZInfo.zoneConfigTwoElements(fMgr, zoneArray, tzStats, ePrefix)

			if err != nil {
				return err
			}

		case 3:
			err = parseZInfo.zoneConfigThreeElements(fMgr, zoneArray, tzStats, ePrefix)

			if err != nil {
				return err
			}

		default:
			return fmt.Errorf(ePrefix +
				"\nError Invalid 'zoneArray' length. Expected length is > 0 and < 4.\n" +
				"Actual 'zoneArray' length='%v'\n" +
				"fMgr='%v'\n", len(zoneArray), fMgr.GetAbsolutePathFileName())
		}
	}

	return nil
}


// zoneCfgTwoElements - Configures and stores data for a two element time zone
// such as 'America/Chicago'. This method configures both the TimeZoneGroupDto and
// the TimeZoneDataDto.
//
// The TimeZoneGroupDto is added to the TimeZoneGroupCollection, 'tzMajorGroupCol'.
// The TimeZoneDataDto is added to the TimeZoneData Collection, 'tzDataCol'
//
func (parseZInfo ParseZoneInfoData) zoneConfigTwoElements(
	fMgr pathfileops.FileMgr,
	zoneArray []string,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "ParseZoneInfoData.zoneConfigTwoElements() "

	if len(zoneArray) != 2 {
		return fmt.Errorf(ePrefix +
			"\nError: Input Parameter length 'zoneArray' is NOT equal to '2'.\n" +
			"'zoneArray' length='%v'\n", len(zoneArray))
	}


	groupAlreadyExists, _ :=
		tzStats.TzGroups[tzdatastructs.Level_01_Idx].ContainsGroupName(
			"", // Parent Group Name - ""
			zoneArray[0]) // Group Name - 'America'

	if !groupAlreadyExists{

		// Configure Time Zone Level-1 Major Group
		// Example: 'America/Chicago'
		tzGroup := tzdatastructs.TimeZoneGroupDto{}
		tzGroup.ParentGroupName = ""
		tzGroup.GroupName = zoneArray[0] // America

		tzGroup.GroupSortValue = tzGroup.NewSortValue(zoneArray[0])

		// Example: 'americaTimeZones'
		tzGroup.TypeName =
			strops.StrOps{}.LowerCaseFirstLetter(zoneArray[0])  +
				tzdatastructs.MasterGroupTypeSuffix

		tzGroup.TypeValue = "string"

		// Example: 'America'
		tzGroup.IanaVariableName =
			strops.StrOps{}.UpperCaseFirstLetter(zoneArray[0])

		tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
		tzGroup.GroupType = tzdatastructs.TzGrpType.Standard()
		tzGroup.GroupClass = tzdatastructs.TzGrpClass.IANA()
		tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
		tzGroup.SetIsInitialized(true)
		err := tzdeclarations.TzGroupDeclarations{}.StandardGrpDeclaration(&tzGroup, ePrefix)

		if err != nil {
			return err
		}

		err = tzStats.CountMajorTimeZoneGroup(tzGroup, ePrefix)

		if err != nil {
			return err
		}
	}

	containsZone, _ := tzStats.TzData[tzdatastructs.Level_01_Idx].ContainsTzName(
		"", // Parent Group Name - ""
		zoneArray[0], // Group Name - 'America'
		zoneArray[1]) // Tz = 'Chicago'

	if containsZone {
		return nil
	}

	// Configure Standard Level-1 Iana Time Zone Data Dto
	tzDataDto := tzdatastructs.TimeZoneDataDto{}

	tzDataDto.GroupName = zoneArray[0] // America - majorGroup
	tzDataDto.TzName = zoneArray[1] // Chicago - tzName
	tzDataDto.TzAliasValue = ""
	tzDataDto.TzCanonicalValue =
		zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1]  // 'America/Chicago'
	tzDataDto.TzValue = tzDataDto.TzCanonicalValue // 'America/Chicago'
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzDataDto.TzCanonicalValue)

	// Example func signature
	// func (amer americaTimeZones) Chicago() string { return "America/Chicago" }

	// Example: americaTimeZones
	tzDataDto.FuncType =
		strops.StrOps{}.LowerCaseFirstLetter(zoneArray[0]) +
			tzdatastructs.MasterGroupTypeSuffix

	// Example: 'ameri'
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:5]

	// Example: 'America/Chicago'
	// FuncName: Chicago()
	tzDataDto.FuncName = parseZInfo.zoneCfgValidFuncName(zoneArray[1])

	tzDataDto.FuncReturnType = "string"

	// Example Function Return Value = "America/Chicago"
	tzDataDto.FuncReturnValue =
		fmt.Sprintf("\"%v\"", tzDataDto.TzCanonicalValue)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Canonical()
	tzDataDto.TzType = tzdatastructs.TZType.Standard()
	tzDataDto.TzCategory = tzdatastructs.TZCat.TimeZone()
	tzDataDto.TzSource = tzdatastructs.TZSrc.Iana()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzDataDto.SetIsInitialized(true)

	err := tzdeclarations.TzZoneDeclarations{}.StandardZoneFuncDeclaration(&tzDataDto, ePrefix)

	if err != nil {
		return err
	}

	err = tzStats.CountIanaStdZone(tzDataDto, tzdatastructs.Level_01_Idx, ePrefix)

	if err != nil {
		return err
	}


	return nil
}

// zoneConfigThreeElements - Configures and stores data for a two element time
// zone such as 'America/Argentina/Buenos_Aires'. This method configures both
// the TimeZoneGroupDto and the TimeZoneDataDto.
//
// The Major TimeZoneGroupDto is added to the TimeZoneGroupCollection,
// 'tzMajorGroupCol'.
//
// The Minor TimeZoneGroupDto is added to the TimeZoneGroupCollection,
// 'tzMinorGroupCol'.
//
// The Place Holder TimeZoneDataDto is added to the TimeZoneData Collection,
// 'tzDataCol'
//
// The Sub-Zone TimeZoneDataDto is added to the TimeZoneData Collection,
// 'subTzDataCol'
//
func (parseZInfo ParseZoneInfoData) zoneConfigThreeElements(fMgr pathfileops.FileMgr,
	zoneArray []string,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix = ePrefix + "ParseIanaTzData.zoneCfgThreeElements() "

	if len(zoneArray) != 3 {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter zoneArray length is NOT equal to '3'.\n" +
			"zoneArray length='%v'\n", len(zoneArray))
	}

	// lenZoneArray must == 3-elements
	// This is a standard IANA sub zone
	// America/Argentina/Buenos_Aires

	// zoneArray[0] == America
	groupAlreadyExists, _ :=
		tzStats.TzGroups[tzdatastructs.Level_01_Idx].ContainsGroupName(
			"", // Parent Group Name - ""
			zoneArray[0]) // Group Name - 'America'

	if !groupAlreadyExists {

		// Configure Level-1 Data
		// Configure Time Zone Level-1 Major Group
		tzGroup := tzdatastructs.TimeZoneGroupDto{}
		tzGroup.ParentGroupName = ""
		tzGroup.GroupName = zoneArray[0]  // America
		tzGroup.GroupSortValue = tzGroup.NewSortValue(zoneArray[0])

		// Example: 'americaTimeZones'
		tzGroup.TypeName =
			strops.StrOps{}.LowerCaseFirstLetter(zoneArray[0])  +
				tzdatastructs.MasterGroupTypeSuffix

		tzGroup.TypeValue = "string"

		// Example: 'America'
		tzGroup.IanaVariableName =
			strops.StrOps{}.UpperCaseFirstLetter(zoneArray[0])

		tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
		tzGroup.GroupType = tzdatastructs.TzGrpType.Standard()
		tzGroup.GroupClass = tzdatastructs.TzGrpClass.IANA()
		tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
		tzGroup.SetIsInitialized(true)
		err := tzdeclarations.TzGroupDeclarations{}.StandardGrpDeclaration(&tzGroup, ePrefix)

		if err != nil {
			return err
		}

		err = tzStats.CountMajorTimeZoneGroup(tzGroup, ePrefix)

		if err != nil {
			return err
		}

	}

	containsZone, _ :=
		tzStats.TzData[tzdatastructs.Level_01_Idx].ContainsTzName(
			"", // Parent Group Name - ""
			zoneArray[0], // Group Name - 'America'
			zoneArray[1]) // Tz - 'Argentina'

	if !containsZone {
		// Add Level-1 Place Holder TimeZoneDataDto
		// Configure Place Holder Time Zone Data Dto
		// Example Time Zone: America/Argentina/Buenos_Aires
		tzDataDto := tzdatastructs.TimeZoneDataDto{}

		tzDataDto.GroupName = zoneArray[0] // America - majorGroup
		tzDataDto.TzName = zoneArray[1] // Argentina - tzName
		tzDataDto.TzAliasValue = ""

		// 'argentinaTimeZones'
		tzDataDto.TzCanonicalValue =
			strops.StrOps{}.LowerCaseFirstLetter(zoneArray[1]) +
				tzdatastructs.MasterGroupTypeSuffix

		// Example Time Zone: America/Argentina/Buenos_Aires
		// 'Argentina'
		tzDataDto.TzValue = zoneArray[1]

		// Example: Argentina
		tzDataDto.TzSortValue =
			tzdatastructs.TimeZoneDataDto{}.NewSortValue(zoneArray[1])

		// Example: America/Argentina/Buenos_Aires

		// func (argen argentinaTimeZones)
		//     Argentina() argentinaTimeZones { return "" }

		// FuncType: americaTimeZones
		tzDataDto.FuncType =
			strops.StrOps{}.LowerCaseFirstLetter(zoneArray[0]) +
				tzdatastructs.MasterGroupTypeSuffix

		// Example: 'ameri'
		tzDataDto.FuncSelfReferenceVariable =
			tzDataDto.FuncType[:5]

		// Example Time Zone: America/Argentina/Buenos_Aires
		// FuncName: Argentina()
		tzDataDto.FuncName = parseZInfo.zoneCfgValidFuncName(zoneArray[1])

		// Example: argentinaTimeZones
		tzDataDto.FuncReturnType =
			strops.StrOps{}.LowerCaseFirstLetter(zoneArray[1]) +
				tzdatastructs.MasterGroupTypeSuffix

		// Example:  ""
		tzDataDto.FuncReturnValue = "\"\""

		tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
		tzDataDto.TzClass = tzdatastructs.TZClass.Artificial()
		tzDataDto.TzType = tzdatastructs.TZType.Group()
		tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
		tzDataDto.SetIsInitialized(true)

		err := tzdeclarations.TzZoneDeclarations{}.PlaceHolderZoneFuncDeclaration(&tzDataDto, ePrefix)

		if err != nil {
			return err
		}

		err = tzStats.CountLevel1TimeZoneCollection(tzDataDto, ePrefix)

		if err != nil {
			return err
		}
	}


	// Configure Level-2 Data
	// Configure Level-2 Secondary Group
	// Example: America/Argentina/Buenos_Aires
	groupAlreadyExists, _ =
		tzStats.TzGroups[tzdatastructs.Level_02_Idx].ContainsGroupName(
			zoneArray[0], // Parent Name - 'America'
			zoneArray[1]) // Group Name - 'Argentina'

	if !groupAlreadyExists {

		tzGroup := tzdatastructs.TimeZoneGroupDto{}
		tzGroup.ParentGroupName = zoneArray[0] // America
		tzGroup.GroupName = zoneArray[1] // Argentina

		// Argentina
		tzGroup.GroupSortValue = tzGroup.NewSortValue(zoneArray[1])

		// Example: 'argentinaTimeZones'
		tzGroup.TypeName =
			strops.StrOps{}.LowerCaseFirstLetter(zoneArray[1]) +
				tzdatastructs.MasterGroupTypeSuffix

		tzGroup.TypeValue = "string"

		tzGroup.IanaVariableName = ""
		tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
		tzGroup.GroupType = tzdatastructs.TzGrpType.SubGroup()
		tzGroup.GroupClass = tzdatastructs.TzGrpClass.IANA()
		tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
		tzGroup.SetIsInitialized(true)
		err := tzdeclarations.TzGroupDeclarations{}.SubGroupDeclaration(&tzGroup, ePrefix)

		if err != nil {
			return err
		}

		err = tzStats.CountLevel2StdSubGroup(tzGroup,ePrefix)

		if err != nil {
			return err
		}
	}

	containsZone, _ =
		tzStats.TzData[tzdatastructs.Level_02_Idx].ContainsTzName(
			zoneArray[0], // Parent Name - 'America'
			zoneArray[1], // Group Name - 'Argentina'
			zoneArray[2]) // Tz - Buenos_Aires

	if !containsZone {

		// Add Sub-Zone Time Zone Data Dto
		// Add the Sub Time Zone to the
		// Sub Time Zone Array (subTzDataCol)
		// Example: America/Argentina/Buenos_Aires
		tzDataDto := tzdatastructs.TimeZoneDataDto{}
		tzDataDto.ParentGroupName = zoneArray[0] // America
		tzDataDto.GroupName = zoneArray[1] // Argentina
		tzDataDto.TzName = zoneArray[2]    // Buenos_Aires - tzName
		tzDataDto.TzAliasValue = ""

		// America/Argentina/Buenos_Aires
		tzDataDto.TzCanonicalValue =
			zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1] +
				tzdatastructs.ZoneSeparator + zoneArray[2]

		// Example: America/Argentina/Buenos_Aires
		tzDataDto.TzValue = tzDataDto.TzCanonicalValue

		tzDataDto.TzSortValue =
			tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzDataDto.TzValue)

		// Example: America/Argentina/Buenos_Aires
		// func (argen argentinaTimeZones)
		//   Buenos_Aires() string { return "America/Argentina/Buenos_Aires" }

		// argentinaTimeZones
		tzDataDto.FuncType = strops.StrOps{}.LowerCaseFirstLetter(zoneArray[1]) +
			tzdatastructs.MasterGroupTypeSuffix

		// Example: 'argen'
		tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:5]

		// Example: America/Argentina/Buenos_Aires
		// Func Name: Buenos_Aires()
		tzDataDto.FuncName = parseZInfo.zoneCfgValidFuncName(zoneArray[2])

		tzDataDto.FuncReturnType = "string"

		// Example Function Return Value = "America/Argentina/Buenos_Aires"
		tzDataDto.FuncReturnValue =
			fmt.Sprintf("\"%v\"", tzDataDto.TzCanonicalValue)

		tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
		tzDataDto.TzClass = tzdatastructs.TZClass.Canonical()
		tzDataDto.TzType = tzdatastructs.TZType.SubZone()
		tzDataDto.TzCategory = tzdatastructs.TZCat.TimeZone()
		tzDataDto.TzSource = tzdatastructs.TZSrc.Iana()
		tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
		tzDataDto.SetIsInitialized(true)
		err := tzdeclarations.TzZoneDeclarations{}.StandardZoneFuncDeclaration(&tzDataDto, ePrefix)

		if err != nil {
			return err
		}

		err = tzStats.CountIanaStdZone(tzDataDto, tzdatastructs.Level_02_Idx, ePrefix)

		if err != nil {
			return err
		}
	}

	return nil
}


// zoneCfgValidFuncName - Generates and returns a valid enum function name for
// a time zone.
//
// This method replaces minus signs ('-') with the word 'Minus'. If the 'funcName'
// includes a plus sign ('+'), it is replaced by the word 'Plus'.  Single digit
// numbers are reformatted with a leading zero.
//
// If the 'funcName' does not have a closed parenthesis ('()'), one will be added.
//
func (parseZInfo ParseZoneInfoData) zoneCfgValidFuncName(funcName string) string {
	/*
		actualFuncName := strings.Replace(funcName, "GMT-", "GMT_Minus_", -1)
		actualFuncName = strings.Replace(actualFuncName, "GMT+", "GMT_Plus_", -1)
	*/
	var actualFuncName string

	if strings.Index(funcName, "-") > -1 {
		actualFuncName = strings.Replace(funcName, "-", "Minus", -1)
	} else if strings.Index(funcName, "+") > -1 {
		actualFuncName = strings.Replace(funcName, "+", "Plus", -1)
	} else {
		actualFuncName = funcName
	}

	if !strings.HasSuffix(actualFuncName,"()") {
		actualFuncName += "()"
	}

	numStrProfile,
	err := strops.StrOps{}.ExtractNumericDigits(
		actualFuncName,
		0,
		"",
		"",
		"")

	if err != nil {
		return actualFuncName
	}

	if numStrProfile.NumStrLen < 1 {
		return actualFuncName
	}

	str1 := actualFuncName[:numStrProfile.FirstNumCharIndex]
	str2 := actualFuncName[numStrProfile.FirstNumCharIndex + numStrProfile.NumStrLen:]

	number, err := strconv.Atoi(numStrProfile.NumStr)

	if err != nil {
		return actualFuncName
	}

	actualFuncName = fmt.Sprintf(str1 + "%02d" + str2, number)

	return actualFuncName
}
