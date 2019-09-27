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

/*
  IANA Time Zone Notes

  How to Read the tz Database Source Files
    https://data.iana.org/time-zones/tz-how-to.html

	See File  zic.8.txt, in the code subdirectory.


      Input files use the format described in this section; output files use
      tzfile(5) format.

      Input files should be text files, that is, they should be a series of
      zero or more lines, each ending in a newline byte and containing at
      most 511 bytes, and without any NUL bytes.  The input text's encoding
      is typically UTF-8 or ASCII; it should have a unibyte representation
      for the POSIX Portable Character Set (PPCS) <http://pubs.opengroup.org/
      onlinepubs/9699919799/basedefs/V1_chap06.html> and the encoding's non-
      unibyte characters should consist entirely of non-PPCS bytes.  Non-PPCS
      characters typically occur only in comments: although output file names
      and time zone abbreviations can contain nearly any character, other
      software will work better if these are limited to the restricted syntax
      described under the -v option.


      Input lines are made up of fields.  Fields are separated from one
      another by one or more white space characters.  The white space
      characters are space, form feed, carriage return, newline, tab, and
      vertical tab.  Leading and trailing white space on input lines is
      ignored.  An unquoted sharp character (#) in the input introduces a
      comment which extends to the end of the line the sharp character
      appears on.  White space characters and sharp characters may be
      enclosed in double quotes (") if they're to be used as part of a field.
      Any line that is blank (after comment stripping) is ignored.  Non-blank
      lines are expected to be of one of three types: rule lines, zone lines,
      and link lines.

*/

var tzGroups [] tzdatastructs.TimeZoneGroupCollection

var tzData [] tzdatastructs.TimeZoneDataCollection

/*
  tzMajorGroupArray Format
    Africa,
    America,
    Antarctica,
    Arctic,
    Asia,
    Atlantic,
    Australia,
      ...
    UCT,
    W_SU,
    WET,
    Zulu
    ------------------------------------------

    tzDataCol Format
      Standard Time Zone -
      ====================
          GroupName  = America
          SubTzName   = ""
          TzName      = Chicago
          TzCanonicalValue     = America/Chicago
          TzClass     = 1
                        // 0 = Unknown
                        // 1 = Canonical
                        // 2 = Alias
                        // 3 = Sub-Group
          Deprecated  = false

        Sub-Zone Place Holder -
        =======================
          GroupName  = America
          SubTzName   = Argentina
          TzName      = Argentina
          TzCanonicalValue     = America/Argentina
          TzClass     = 3
                        // 0 = Unknown
                        // 1 = Canonical
                        // 2 = Alias
                        // 3 = Sub-Group
          Deprecated  = false


*/

type ParseIanaTzData struct {
	input string
	output string

}

// ParseTzAndLinks - Parses Time Zone Data from
// IANA Time Zone files.
func (parseTz *ParseIanaTzData) ParseTzAndLinks(
	dirFileInfo pathfileops.FileMgrCollection) (
	[] tzdatastructs.TimeZoneGroupCollection, // Array of Time Zone Group Collections
	[] tzdatastructs.TimeZoneDataCollection,  // Array of Time Zone Data Collections
	error)  {

	ePrefix := "ParseIanaTzData.ParseTzAndLinks() "

	tzGroups =  make([]tzdatastructs.TimeZoneGroupCollection, 3, 10)
	tzData = make([]tzdatastructs.TimeZoneDataCollection, 3, 10)

	numOfFiles := dirFileInfo.GetNumOfFileMgrs()

	fmt.Println("Number of Target Files: ", numOfFiles)

	if numOfFiles < 5 {
		return tzGroups,
			tzData,
			fmt.Errorf(ePrefix+"Number of files is less than 5!\n" +
				"Number of Files='%v'", numOfFiles)
	}

	for i:=0; i < numOfFiles; i++ {

		fMgr, err := dirFileInfo.PeekFileMgrAtIndex(i)

		if err != nil {
			return tzGroups,
				tzData,
				fmt.Errorf(ePrefix+"%v\n", err.Error())
		}

		isSkipFile, err := parseTz.isSkipFile(fMgr)

		if err != nil {
			return tzGroups,
				tzData,
				fmt.Errorf(ePrefix+"%v\n", err.Error())
		}

		if isSkipFile {
			continue
		}

		fmt.Println("Valid File: ", fMgr.GetFileNameExt())

		err =  parseTz.processFileBytes(fMgr)

		if err != nil {
			return tzGroups,
				tzData,
				fmt.Errorf(ePrefix+
					"File Name: %v\n" +
					"Error=%v\n",
					fMgr.GetAbsolutePathFileName(),  err.Error())
		}

	}

	err := parseTz.configMilitaryTimeZones(ePrefix)

	return tzGroups,
		tzData,
		err
}

// configMilitaryTimeZones - Creates and stores Military Time Zones
//
func (parseTz *ParseIanaTzData) configMilitaryTimeZones(ePrefix string) error {
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

	_, err = tzGroups[tzdatastructs.Level_01_Idx].AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzGroups[tzdatastructs.Level_01_Idx] Error\n" +
			"Error: %v\n", err.Error() )
	}

	for i:=0; i < len(tzdatastructs.MilitaryTzArray); i++ {

		// Configure Standard Level-1 Iana Time Zone Data Dto
		// For Military Time Zone
		tzDataDto := tzdatastructs.TimeZoneDataDto{}

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
		tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(tzdatastructs.MilitaryTzArray[i])

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

		_, err = tzData[tzdatastructs.Level_01_Idx].AddIfNew(tzDataDto)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"tzData[tzdatastructs.Level_01_Idx] Error\n" +
				"Military Time Zone: %v\n" +
				"Error: %v\n", tzdatastructs.MilitaryTzArray[i], err.Error())
		}

	}

	return nil
}

// extractLink - Extracts link data from IANA Time Zone files.
// Format for Link:
// Link -> Canonical -> Alias
// Link  America/Panama America/Cayman
// Canonical =  America/Panama
// Link = America/Cayman
func (parseTz *ParseIanaTzData) extractLink(
	fMgr pathfileops.FileMgr ,rawString string) error {

	ePrefix := "ParseIanaTzData.extractLink() "

	startIdx := 0
// Extract Field 1 - Canonical Field
	dFProfile,
	err :=
		strops.StrOps{}.ExtractDataField(
			rawString,
			[]string{tzdatastructs.LinkLabel},
			startIdx,
			tzdatastructs.LeadingFieldSeparators,
			tzdatastructs.TrailingFieldSeparators,
			tzdatastructs.CommentDelimiters,
			tzdatastructs.EndOfLineDelimiters)

	if err != nil {
		return fmt.Errorf(ePrefix + "%v\n", err.Error())
	}

	if dFProfile.DataFieldLength < 1 {
		return nil
	}

	if strings.Index(dFProfile.DataFieldStr, tzdatastructs.ZoneSeparator) == -1 {
		return nil
	}

	tzCanonical := dFProfile.DataFieldStr
	startIdx = dFProfile.NextTargetStrIndex

	// Extract Field 2 - Link Field
	dFProfile,
	err =
		strops.StrOps{}.ExtractDataField(
			rawString,
			[]string{tzdatastructs.LinkLabel},
			startIdx,
			tzdatastructs.LeadingFieldSeparators,
			tzdatastructs.TrailingFieldSeparators,
			tzdatastructs.CommentDelimiters,
			tzdatastructs.EndOfLineDelimiters)

	if err != nil {
		return fmt.Errorf(ePrefix + "%v\n", err.Error())
	}

	if dFProfile.DataFieldLength < 1 {
		return nil
	}

	if strings.Index(dFProfile.DataFieldStr, tzdatastructs.ZoneSeparator) == -1 {
		return nil
	}

	tzLink := dFProfile.DataFieldStr

	linkZoneArray := strings.Split(tzLink, tzdatastructs.ZoneSeparator)

	lenZoneArray := len(linkZoneArray)

	if lenZoneArray < 1 ||
		lenZoneArray > 3 {
		fmt.Printf(ePrefix + "Invalid Link Time Zone!\n" +
			"FileName: %v\n" +
			"Tz Link String: %v\n" +
			"Tz Canonical String: %v\n",
			fMgr.GetFileNameExt(), tzLink, tzCanonical)
		return nil
	}

	if lenZoneArray == 1 {
		// Example: link -> canonical time zone
		//          'Egypt' -> 'Africa/Cairo'
		// Canonical = 'Africa/Cairo'
		// Link      = 'Egypt'
		return parseTz.linkCfgOneElement(fMgr, tzLink, tzCanonical, ePrefix)
	}

	if lenZoneArray == 2 {
		// Example
		// Link -> Canonical -> Alias
		// Link  America/Panama America/Cayman
		// Canonical =  America/Panama
		// Link = America/Cayman
		return parseTz.linkCfgTwoElements(fMgr, linkZoneArray, tzCanonical, ePrefix)
	}

	// Zone Array Length MUST be 3

	// Example
	//
	// Link -> Canonical                  -> Alias
	// Link    America/Argentina/Catamarca   America/Argentina/ComodRivadavia
	// Canonical =  America/Argentina/Catamarca
	// Link      = America/Argentina/ComodRivadavia

	return parseTz.linkCfgThreeElements(fMgr, linkZoneArray, tzCanonical, ePrefix)

}


// extractZone - Extracts standard time zones and sub time zones.
// Data is stored in tzMajorGroupMap, tzDataCol and
// or subTzDataCol.
func (parseTz *ParseIanaTzData) extractZone(
	fMgr pathfileops.FileMgr, rawString string) error {

	ePrefix := "ParseIanaTzData.extractZone() "

	dFProfile,
	err :=
		strops.StrOps{}.ExtractDataField(
			rawString,
			[]string{tzdatastructs.ZoneLabel},
			0,
			tzdatastructs.LeadingFieldSeparators,
			tzdatastructs.TrailingFieldSeparators,
			tzdatastructs.CommentDelimiters,
			tzdatastructs.EndOfLineDelimiters)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	if dFProfile.DataFieldLength < 1 {
		return nil
	}

	if strings.Index(dFProfile.DataFieldStr, tzdatastructs.ZoneSeparator) == -1 {
		return nil
	}

	zoneArray := strings.Split(dFProfile.DataFieldStr, tzdatastructs.ZoneSeparator)

	lenZoneArray := len(zoneArray)

	if lenZoneArray < 2 ||
		lenZoneArray > 3 {
		fmt.Printf(ePrefix+"Invalid Time Zone!\n"+
			"FileName: %v\n"+
			"Raw Zone String: %v\n", fMgr.GetFileNameExt(), rawString)
		return nil
	}

	// Zone Array has two elements
	// Example: 'America/Chicago'
	if lenZoneArray == 2 {

		return parseTz.zoneCfgTwoElements(fMgr, zoneArray, ePrefix)
	}

	// lenZoneArray must == 3
	// This is a sub zone
	// America/Argentina/Buenos_Aires

	return parseTz.zoneCfgThreeElements(fMgr, zoneArray, ePrefix)
}

// linkCfgOneElement - Configures and stores data associated
// with a time zone 'Link' which consists of a single link
// string.
//
// Example: link -> canonical time zone
//          'Egypt' -> 'Africa/Cairo'
//
func (parseTz *ParseIanaTzData) linkCfgOneElement(
	fMgr pathfileops.FileMgr,
	linkZone,
	canonicalZone,
	ePrefix string) error {

	ePrefix += "ParseIanaTzData.linkCfgOneElement() "

	// Configure Level-1 Data
	// Configure Time Zone Level-1 Group
	tzGroup := tzdatastructs.TimeZoneGroupDto{}
	tzGroup.ParentGroupName = ""

	groupAlreadyExists, _ := tzGroups[tzdatastructs.Level_01_Idx].ContainsGroupName(tzdatastructs.DeprecatedTzGroup)

	if !groupAlreadyExists {

		// "Deprecated"
		tzGroup.GroupName = tzdatastructs.DeprecatedTzGroup
		tzGroup.GroupSortValue = tzdatastructs.DeprecatedTzGroup

		// deprecatedTimeZones
		tzGroup.TypeName =
			strops.StrOps{}.LowerCaseFirstLetter(tzdatastructs.DeprecatedTzGroup)  +
				tzdatastructs.MasterGroupTypeSuffix

		tzGroup.TypeValue = "string"

		tzGroup.IanaVariableName = tzGroup.GroupName
		tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
		tzGroup.GroupType = tzdatastructs.TzGrpType.Standard()
		tzGroup.GroupClass = tzdatastructs.TzGrpClass.IANA()
		tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
		tzGroup.SetIsInitialized(true)

		err := tzdeclarations.TzGroupDeclarations{}.DeprecatedGrpDeclaration(&tzGroup, ePrefix)

		if err != nil {
			return err
		}

		err = tzGroups[tzdatastructs.Level_01_Idx].Add(tzGroup)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"tzGroups[tzdatastructs.Level_02_Idx] Error\n" +
				"FileName: %v\n" +
				"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
		}

	}

	containsZone, _ :=
		tzData[tzdatastructs.Level_01_Idx].ContainsTzName(
			tzdatastructs.DeprecatedTzGroup, linkZone)

	if !containsZone {
		// Level - 1 Time Zone Data for single element
		//   Deprecated Time Zone.
		//
		// Configure Deprecated Link - Time Zone Data Dto
		tzDataDto := tzdatastructs.TimeZoneDataDto{}

		tzDataDto.GroupName = tzdatastructs.DeprecatedTzGroup
		tzDataDto.TzName = linkZone // Egypt
		tzDataDto.TzAliasValue = linkZone // Egypt
		tzDataDto.TzCanonicalValue = canonicalZone // 'Africa/Cairo'
		tzDataDto.TzValue = tzDataDto.TzCanonicalValue // 'Africa/Cairo'
		tzDataDto.TzSortValue =
			tzdatastructs.TimeZoneDataDto{}.NewSortValue(linkZone)

		// func (depri deprecatedTimeZones) Egypt() string { return "Africa/Cairo" }

		// Example: deprecatedTimeZones
		tzDataDto.FuncType =
			strops.StrOps{}.LowerCaseFirstLetter(tzDataDto.GroupName) +
				tzdatastructs.MasterGroupTypeSuffix

		// Example: 'depre'
		tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:6]

		// Example: link -> canonical time zone
		//          'Egypt' -> 'Africa/Cairo'
		//
		// Func Name: Egypt()
		tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(linkZone)

		tzDataDto.FuncReturnType = "string"

		// Example Function Return Value = "Africa/Cairo"
		tzDataDto.FuncReturnValue = fmt.Sprintf("\"%v\"", canonicalZone)

		tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
		tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
		tzDataDto.TzType = tzdatastructs.TZType.Standard()
		tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
		tzDataDto.SetIsInitialized(true)

		err := tzdeclarations.TzZoneDeclarations{}.LinkTimeZoneDeclaration(&tzDataDto, ePrefix)

		if err != nil {
			return err
		}

		err = tzData[tzdatastructs.Level_01_Idx].Add(tzDataDto)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"tzData[tzdatastructs.Level_01_Idx] Error\n"+
				"Error: %v\n" +
				"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
		}

	}

	return nil
}

// linkCfgTwoElements - Configures and stores data associated
// with a time zone 'Link' which consists of a 2-element link
// zone string.
//
// Example: link -> canonical time zone
//          'US/Alaska' -> 'America/Anchorage'
//
func (parseTz *ParseIanaTzData) linkCfgTwoElements(
	fMgr pathfileops.FileMgr,
	linkZoneArray []string,
	canonicalZone,
	ePrefix string) error {

	ePrefix += "ParseIanaTzData.linkCfgTwoElements() "


	if len(linkZoneArray) != 2 {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter linkZoneArray length is NOT equal to '2'.\n" +
			"linkZoneArray length='%v'\n", len(linkZoneArray))
	}

	groupAlreadyExists, _ := tzGroups[tzdatastructs.Level_01_Idx].ContainsGroupName(tzdatastructs.DeprecatedTzGroup)

	// Configure Level-1 Data
	// Configure Time Zone Level-1 Group
	// Example: link -> canonical time zone
	//          'US/Alaska' -> 'America/Anchorage'

	if !groupAlreadyExists {

		tzGroup := tzdatastructs.TimeZoneGroupDto{}
		tzGroup.ParentGroupName = ""

		// "Deprecated"
		tzGroup.GroupName = tzdatastructs.DeprecatedTzGroup
		tzGroup.GroupSortValue = tzdatastructs.DeprecatedTzGroup

		// deprecatedTimeZones
		tzGroup.TypeName =
			strops.StrOps{}.LowerCaseFirstLetter(tzdatastructs.DeprecatedTzGroup)  +
				tzdatastructs.MasterGroupTypeSuffix

		tzGroup.TypeValue = "string"

		tzGroup.IanaVariableName = tzGroup.GroupName
		tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
		tzGroup.GroupType = tzdatastructs.TzGrpType.Standard()
		tzGroup.GroupClass = tzdatastructs.TzGrpClass.IANA()
		tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
		tzGroup.SetIsInitialized(true)

		err := tzdeclarations.TzGroupDeclarations{}.DeprecatedGrpDeclaration(&tzGroup, ePrefix)

		if err != nil {
			return err
		}

		err = tzGroups[tzdatastructs.Level_01_Idx].Add(tzGroup)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"tzGroups[tzdatastructs.Level_01_Idx] Error\n" +
				"FileName: %v\n" +
				"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
		}

	}

	containsZone, _ :=
		tzData[tzdatastructs.Level_01_Idx].ContainsTzName(
			tzdatastructs.DeprecatedTzGroup, linkZoneArray[0])

	if !containsZone {
		// Level - 1 Time Zone Data for single element
		//   Deprecated Time Zone Place Holder
		//
		// Configure Deprecated Link - Time Zone Data Dto
		// Example: link -> canonical time zone
		//          'US/Alaska' -> 'America/Anchorage'
		tzDataDto := tzdatastructs.TimeZoneDataDto{}

		tzDataDto.GroupName = tzdatastructs.DeprecatedTzGroup // Deprecated
		tzDataDto.TzName = linkZoneArray[0] // US
		tzDataDto.TzAliasValue =
			linkZoneArray[0] +
				tzdatastructs.ZoneSeparator + linkZoneArray[1] // US/Alaska
		tzDataDto.TzCanonicalValue = linkZoneArray[0] // US
		tzDataDto.TzValue = linkZoneArray[0] // US
		tzDataDto.TzSortValue =
			tzdatastructs.TimeZoneDataDto{}.NewSortValue(linkZoneArray[0])

		// Example:  link       -> canonical time zone
		//          'US/Alaska' -> 'America/Anchorage'
		// func (depre deprecatedTimeZones)
		//     US() uSDeprecatedTimeZones { return uSDeprecatedTimeZones("") }

		// Example: deprecatedTimeZones
		tzDataDto.FuncType =
			strops.StrOps{}.LowerCaseFirstLetter(tzDataDto.GroupName) +
				tzdatastructs.MasterGroupTypeSuffix

		// Example: 'depre'
		tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:6]

		// Example:  link       -> canonical time zone
		//          'US/Alaska' -> 'America/Anchorage'
		// Link            = 'US/Alaska'
		// Canonical Value = 'America/Anchorage'
		// Func Name: US()
		tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(linkZoneArray[0])

		// Example: uSDeprecatedTimeZones
		tzDataDto.FuncReturnType =
			strops.StrOps{}.LowerCaseFirstLetter(
				linkZoneArray[0]) +
				tzdatastructs.DeprecatedTzGroup +
				tzdatastructs.MasterGroupTypeSuffix

		// Example Function Return Value = uSDeprecatedTimeZones("")
		tzDataDto.FuncReturnValue =
			fmt.Sprintf("%v(\"\")", tzDataDto.FuncReturnType)

		tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
		tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
		tzDataDto.TzType = tzdatastructs.TZType.Group()
		tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
		tzDataDto.SetIsInitialized(true)

		err := tzdeclarations.TzZoneDeclarations{}.PlaceHolderLinkFuncDeclaration(&tzDataDto, ePrefix)

		if err != nil {
			return err
		}

		err = tzData[tzdatastructs.Level_01_Idx].Add(tzDataDto)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"tzData[tzdatastructs.Level_01_Idx] Error\n"+
				"Error: %v\n" +
				"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
		}
	}

	// Configure Level-2 Data
	// Example: link -> canonical time zone
	//          'US/Alaska' -> 'America/Anchorage'
	//
	// Configure Level-2 Group Dto
	//

	// linkZoneArray[0] == US
	groupAlreadyExists, _ = tzGroups[tzdatastructs.Level_02_Idx].ContainsGroupName(linkZoneArray[0])

	if !groupAlreadyExists {
		tzGroup := tzdatastructs.TimeZoneGroupDto{}

		// "Deprecated"
		tzGroup.ParentGroupName = tzdatastructs.DeprecatedTzGroup

		tzGroup.GroupName = linkZoneArray[0] // US
		tzGroup.GroupSortValue = linkZoneArray[0] // US

		// Example: uSDeprecatedTimeZones
		tzGroup.TypeName =
			strops.StrOps{}.LowerCaseFirstLetter(
				linkZoneArray[0]) +
				tzdatastructs.DeprecatedTzGroup +
				tzdatastructs.MasterGroupTypeSuffix

		tzGroup.TypeValue = "string"

		tzGroup.IanaVariableName = tzGroup.GroupName
		tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
		tzGroup.GroupType = tzdatastructs.TzGrpType.SubGroup()
		tzGroup.GroupClass = tzdatastructs.TzGrpClass.IANA()
		tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
		tzGroup.SetIsInitialized(true)

		err := tzdeclarations.TzGroupDeclarations{}.DeprecatedSubGrpDeclaration(&tzGroup, ePrefix)

		if err != nil {
			return err
		}

		err = tzGroups[tzdatastructs.Level_02_Idx].Add(tzGroup)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"tzGroups[tzdatastructs.Level_02_Idx] Error\n" +
				"FileName: %v\n" +
				"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
		}
	}

	containsZone, _ =
		tzData[tzdatastructs.Level_02_Idx].ContainsTzName(
			linkZoneArray[0], linkZoneArray[1])

		if !containsZone {
			// Level - 2 Time Zone Data
			//   Deprecated Time Zone Canonical Value
			//
			// Configure Deprecated Link - Time Zone Data Dto
			// Example: link -> canonical time zone
			//          'US/Alaska' -> 'America/Anchorage'

			tzDataDto := tzdatastructs.TimeZoneDataDto{}

			tzDataDto.GroupName = linkZoneArray[0] // US
			tzDataDto.TzName = linkZoneArray[1] // Alaska
			tzDataDto.TzAliasValue =
				linkZoneArray[0] +
					tzdatastructs.ZoneSeparator + linkZoneArray[1] // US/Alaska
			tzDataDto.TzCanonicalValue = canonicalZone // America/Anchorage
			tzDataDto.TzValue = linkZoneArray[1] // Alaska
			tzDataDto.TzSortValue =
				tzdatastructs.TimeZoneDataDto{}.NewSortValue(linkZoneArray[1])

			// Example: link -> canonical time zone
			//          'US/Alaska' -> 'America/Anchorage'
			// Link            = 'US/Alaska'
			// Canonical Value = 'America/Anchorage'
			//
			// func (uSDep uSDeprecatedTimeZones)
			//     Alaska() string { return "America/Anchorage" }

			// Example: uSDeprecatedTimeZones
			tzDataDto.FuncType =
				strops.StrOps{}.LowerCaseFirstLetter(
					linkZoneArray[0]) +
					tzdatastructs.DeprecatedTzGroup +
					tzdatastructs.MasterGroupTypeSuffix

			// Example: 'uSDepr'
			tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:6]

			// Example: link -> canonical time zone
			//          'US/Alaska' -> 'America/Anchorage'
			// Link            = 'US/Alaska'
			// Canonical Value = 'America/Anchorage'
			//
			// Func Name: Alaska()
			tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(linkZoneArray[1])

			// Example: string
			tzDataDto.FuncReturnType = "string"

			// Example Function Return Value = "America/Anchorage"
			tzDataDto.FuncReturnValue =
				fmt.Sprintf("\"%v\"", canonicalZone)

			tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
			tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
			tzDataDto.TzType = tzdatastructs.TZType.SubZone()
			tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
			tzDataDto.SetIsInitialized(true)

			err := tzdeclarations.TzZoneDeclarations{}.LinkTimeZoneDeclaration(&tzDataDto, ePrefix)

			if err != nil {
				return err
			}

			err = tzData[tzdatastructs.Level_02_Idx].Add(tzDataDto)

			if err != nil {
				return fmt.Errorf(ePrefix +
					"tzData[tzdatastructs.Level_01_Idx] Error\n"+
					"Error: %v\n" +
					"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
			}
	}

	return nil
}

// linkCfgThreeElements - Configures and stores data
// associated with a time zone 'Link' which consists
// of a 3-element link string.
//
// Example:
// link                               -> canonical time zone
// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
//
// Link                = 'America/Argentina/ComodRivadavia'
// Canonical Time Zone = 'America/Argentina/Catamarca'
//
func (parseTz *ParseIanaTzData) linkCfgThreeElements(
	fMgr pathfileops.FileMgr,
	linkZoneArray []string,
	canonicalZone,
	ePrefix string) error {

	ePrefix += "ParseIanaTzData.linkCfgThreeElements() "


	if len(linkZoneArray) != 3 {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter linkZoneArray length is NOT equal to '3'.\n" +
			"linkZoneArray length='%v'\n", len(linkZoneArray))
	}

	// Configure Level-1 Data
	// Configure Time Zone Level-1 Group
	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'
	//
	tzGroup := tzdatastructs.TimeZoneGroupDto{}
	tzGroup.ParentGroupName = ""

	// "Deprecated"
	tzGroup.GroupName = tzdatastructs.DeprecatedTzGroup
	tzGroup.GroupSortValue = tzdatastructs.DeprecatedTzGroup

	// deprecatedTimeZones
	tzGroup.TypeName =
		strops.StrOps{}.LowerCaseFirstLetter(tzdatastructs.DeprecatedTzGroup)  +
			tzdatastructs.MasterGroupTypeSuffix

	tzGroup.TypeValue = "string"

	tzGroup.IanaVariableName = tzGroup.GroupName
	tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
	tzGroup.GroupType = tzdatastructs.TzGrpType.Standard()
	tzGroup.GroupClass = tzdatastructs.TzGrpClass.IANA()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
	tzGroup.SetIsInitialized(true)

	_, err := tzGroups[tzdatastructs.Level_01_Idx].AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzGroups[tzdatastructs.Level_02_Idx] Error\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

	// Level - 1 Time Zone Data for single element
	//   Deprecated Time Zone Place Holder
	//
	// Configure Deprecated Link - Time Zone Data Dto
	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'
	tzDataDto := tzdatastructs.TimeZoneDataDto{}

	tzDataDto.GroupName = tzdatastructs.DeprecatedTzGroup
	tzDataDto.TzName = linkZoneArray[0] // America
	tzDataDto.TzAliasValue =
		linkZoneArray[0] +
			tzdatastructs.ZoneSeparator + linkZoneArray[1] // America/Argentina
	tzDataDto.TzCanonicalValue = linkZoneArray[0] // America
	tzDataDto.TzValue = linkZoneArray[0] // America
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(linkZoneArray[0])

	// Configure Deprecated Link - Time Zone Data Dto
	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'

	// func (depre deprecatedTimeZones)
	//  America() americaDeprecatedTimeZones { return americaDeprecatedTimeZones("") }

	// Example: deprecatedTimeZones
	tzDataDto.FuncType =
		strops.StrOps{}.LowerCaseFirstLetter(tzDataDto.GroupName) +
			tzdatastructs.MasterGroupTypeSuffix

	// Example: 'depre'
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:6]

	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'
	// Func Name: America()
	tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(linkZoneArray[0])

	// Example: americaDeprecatedTimeZones
	tzDataDto.FuncReturnType =
		strops.StrOps{}.LowerCaseFirstLetter(
			linkZoneArray[0]) +
			tzdatastructs.DeprecatedTzGroup +
			tzdatastructs.MasterGroupTypeSuffix

	// Example Function Return Value = americaDeprecatedTimeZones("")
	tzDataDto.FuncReturnValue =
		fmt.Sprintf("%v(\"\")", tzDataDto.FuncReturnType)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
	tzDataDto.TzType = tzdatastructs.TZType.Group()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
	tzDataDto.SetIsInitialized(true)

	_, err = tzData[tzdatastructs.Level_01_Idx].AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzData[tzdatastructs.Level_01_Idx] Error\n"+
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Configure Level-2 Data
	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'
	//
	// Configure Level-2 Group Dto
	//
	tzGroup = tzdatastructs.TimeZoneGroupDto{}

	// "Deprecated"
	tzGroup.ParentGroupName = tzdatastructs.DeprecatedTzGroup

	tzGroup.GroupName = linkZoneArray[0] // America
	tzGroup.GroupSortValue = linkZoneArray[0] // America

	// Example: americaDeprecatedTimeZones
	tzGroup.TypeName =
		strops.StrOps{}.LowerCaseFirstLetter(
			linkZoneArray[0]) +
			tzdatastructs.DeprecatedTzGroup +
			tzdatastructs.MasterGroupTypeSuffix

	tzGroup.TypeValue = "string"

	tzGroup.IanaVariableName = tzGroup.GroupName
	tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
	tzGroup.GroupType = tzdatastructs.TzGrpType.SubGroup()
	tzGroup.GroupClass = tzdatastructs.TzGrpClass.IANA()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
	tzGroup.SetIsInitialized(true)

	_, err = tzGroups[tzdatastructs.Level_02_Idx].AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzGroups[tzdatastructs.Level_02_Idx] Error\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

	// Level - 2 Time Zone Data
	//   Deprecated Time Zone Sub-Group Value
	//
	// Configure Deprecated Link - Time Zone Data Dto
	//
	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'
	//
	tzDataDto = tzdatastructs.TimeZoneDataDto{}

	tzDataDto.GroupName = linkZoneArray[0] // America
	tzDataDto.TzName = linkZoneArray[1] // Argentina
	tzDataDto.TzAliasValue =
		linkZoneArray[0] +
			tzdatastructs.ZoneSeparator + linkZoneArray[1] // America/Argentina
	tzDataDto.TzCanonicalValue = linkZoneArray[1] // Argentina
	tzDataDto.TzValue = linkZoneArray[1] // Argentina
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(linkZoneArray[1])

	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'

	// func (ameri americaDeprecatedTimeZones)
	//     Argentina() argentinaDeprecatedTimeZones { return argentinaDeprecatedTimeZones("") }

	// Example: americaDeprecatedTimeZones
	tzDataDto.FuncType =
		strops.StrOps{}.LowerCaseFirstLetter(
			linkZoneArray[0]) +
			tzdatastructs.DeprecatedTzGroup +
			tzdatastructs.MasterGroupTypeSuffix

	// Example: 'ameri'
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:6]

	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'
	// Func Name: Argentina()
	tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(linkZoneArray[1])

	// Example: argentinaDeprecatedTimeZones
	tzDataDto.FuncReturnType =
		strops.StrOps{}.LowerCaseFirstLetter(
			linkZoneArray[0]) +
			tzdatastructs.DeprecatedTzGroup +
			tzdatastructs.MasterGroupTypeSuffix

	// Example Function Return Value = "argentinaDeprecatedTimeZones("")"
	tzDataDto.FuncReturnValue =
		fmt.Sprintf("%v(\"\")", tzDataDto.FuncReturnType)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
	tzDataDto.TzType = tzdatastructs.TZType.Group()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
	tzDataDto.SetIsInitialized(true)

	_, err = tzData[tzdatastructs.Level_02_Idx].AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzData[tzdatastructs.Level_02_Idx] Error\n"+
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Configure Level-3 Data
	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'
	//
	// Configure Level-3 Group Dto
	//
	tzGroup = tzdatastructs.TimeZoneGroupDto{}

	// "America"
	tzGroup.ParentGroupName = linkZoneArray[0] // America

	tzGroup.GroupName = linkZoneArray[1] // Argentina
	tzGroup.GroupSortValue = linkZoneArray[0] // Argentina

	// Example: argentinaDeprecatedTimeZones
	tzGroup.TypeName =
		strops.StrOps{}.LowerCaseFirstLetter(
			linkZoneArray[0]) +
			tzdatastructs.DeprecatedTzGroup +
			tzdatastructs.MasterGroupTypeSuffix

	tzGroup.TypeValue = "string"

	tzGroup.IanaVariableName = tzGroup.GroupName
	tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
	tzGroup.GroupType = tzdatastructs.TzGrpType.SubGroup()
	tzGroup.GroupClass = tzdatastructs.TzGrpClass.IANA()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
	tzGroup.SetIsInitialized(true)

	_, err = tzGroups[tzdatastructs.Level_03_Idx].AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzGroups[tzdatastructs.Level_03_Idx] Error\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

	// Level - 3 Time Zone Data
	//   Deprecated Time Zone Canonical Value
	// Configure Deprecated Link - Time Zone Data Dto
	//
	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'
	//
	tzDataDto = tzdatastructs.TimeZoneDataDto{}

	tzDataDto.GroupName = linkZoneArray[1] // Argentina // US
	tzDataDto.TzName = linkZoneArray[2] // ComodRivadavia

	// America/Argentina/ComodRivadavia
	tzDataDto.TzAliasValue =
		linkZoneArray[0] +
			tzdatastructs.ZoneSeparator + linkZoneArray[1] +
			tzdatastructs.ZoneSeparator + linkZoneArray[2]

	tzDataDto.TzCanonicalValue = canonicalZone // America/Argentina/Catamarca
	tzDataDto.TzValue = linkZoneArray[2] // ComodRivadavia
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(linkZoneArray[2])

	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'

	// func (argen argentinaDeprecatedTimeZones)
	//     ComodRivadavia() string { return "America/Argentina/Catamarca" }

	// Example: argentinaDeprecatedTimeZones
	tzDataDto.FuncType =
		strops.StrOps{}.LowerCaseFirstLetter(
			linkZoneArray[1]) +
			tzdatastructs.DeprecatedTzGroup +
			tzdatastructs.MasterGroupTypeSuffix

	// Example: 'argen'
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:6]

	// Example:
	// link                               -> canonical time zone
	// 'America/Argentina/ComodRivadavia' -> 'America/Argentina/Catamarca'
	//
	// Link                = 'America/Argentina/ComodRivadavia'
	// Canonical Time Zone = 'America/Argentina/Catamarca'
	//
	// Func Name: ComodRivadavia()
	tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(linkZoneArray[2])

	// Example: string
	tzDataDto.FuncReturnType = "string"

	// Example Function Return Value = "America/Argentina/Catamarca"
	tzDataDto.FuncReturnValue =
		fmt.Sprintf("\"%v\"", canonicalZone)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
	tzDataDto.TzType = tzdatastructs.TZType.SubZone()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
	tzDataDto.SetIsInitialized(true)

	_, err = tzData[tzdatastructs.Level_03_Idx].AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzData[tzdatastructs.Level_03_Idx] Error\n"+
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
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
func (parseTz *ParseIanaTzData) zoneCfgTwoElements(
	fMgr pathfileops.FileMgr, zoneArray []string, ePrefix string) error {

	ePrefix += "ParseIanaTzData.zoneCfgTwoElements() "

	if len(zoneArray) != 2 {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter zoneArray length is NOT equal to '2'.\n" +
			"zoneArray length='%v'\n", len(zoneArray))
	}

	groupAlreadyExists, _ := tzGroups[tzdatastructs.Level_01_Idx].ContainsGroupName(zoneArray[0])

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

		err = tzGroups[tzdatastructs.Level_01_Idx].Add(tzGroup)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"tzGroups[tzdatastructs.Level_01_Idx] Error\n" +
				"FileName: %v\n" +
				"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
		}

	}

	containsZone, _ := tzData[tzdatastructs.Level_01_Idx].ContainsTzName(zoneArray[0], zoneArray[1])

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
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:6]

	// Example: 'America/Chicago'
	// FuncName: Chicago()
	tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(zoneArray[1])

	tzDataDto.FuncReturnType = "string"

	// Example Function Return Value = "America/Chicago"
	tzDataDto.FuncReturnValue = 
		fmt.Sprintf("\"%v\"", tzDataDto.TzCanonicalValue)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Canonical()
	tzDataDto.TzType = tzdatastructs.TZType.Standard()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzDataDto.SetIsInitialized(true)

	err := tzdeclarations.TzZoneDeclarations{}.StandardZoneFuncDeclaration(&tzDataDto, ePrefix)

	if err != nil {
		return err
	}

	err = tzData[tzdatastructs.Level_01_Idx].Add(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"tzData[tzdatastructs.Level_01_Idx] Error\n" +
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	return nil
}

// zoneCfgThreeElements - Configures and stores data for a two element time zone
// such as 'America/Argentina/Buenos_Aires'. This method configures both the
// TimeZoneGroupDto and the TimeZoneDataDto.
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
func (parseTz *ParseIanaTzData) zoneCfgThreeElements(
	fMgr pathfileops.FileMgr, zoneArray []string, ePrefix string) error {

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
	groupAlreadyExists, _ := tzGroups[tzdatastructs.Level_01_Idx].ContainsGroupName(zoneArray[0])

	if !groupAlreadyExists {

		// Configure Level-1 Data
		// Configure Time Zone Level-1 Master Group
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

		err = tzGroups[tzdatastructs.Level_01_Idx].Add(tzGroup)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\ntzGroups[tzdatastructs.Level_01_Idx] Error\n" +
				"FileName: %v\n" +
				"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
		}

	}

	containsZone, _ := tzData[tzdatastructs.Level_01_Idx].ContainsTzName(zoneArray[0], zoneArray[1])

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
		//     Argentina() uSDeprecatedTimeZones { return uSDeprecatedTimeZones("") }

		// FuncType: americaTimeZones
		tzDataDto.FuncType =
			strops.StrOps{}.LowerCaseFirstLetter(zoneArray[0]) +
				tzdatastructs.MasterGroupTypeSuffix

		// Example: 'ameri'
		tzDataDto.FuncSelfReferenceVariable =
			tzDataDto.FuncType[:6]

		// Example Time Zone: America/Argentina/Buenos_Aires
		// FuncName: Argentina()
		tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(zoneArray[1])

		// Example: argentinaTimeZones
		tzDataDto.FuncReturnType =
			strops.StrOps{}.LowerCaseFirstLetter(zoneArray[1]) +
				tzdatastructs.MasterGroupTypeSuffix

		// Example: argentinaTimeZones("")
		tzDataDto.FuncReturnValue =
			fmt.Sprintf("%v(\"\")",
				tzDataDto.FuncReturnType)

		tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
		tzDataDto.TzClass = tzdatastructs.TZClass.Canonical()
		tzDataDto.TzType = tzdatastructs.TZType.Group()
		tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
		tzDataDto.SetIsInitialized(true)

		err := tzdeclarations.TzZoneDeclarations{}.PlaceHolderZoneFuncDeclaration(&tzDataDto, ePrefix)

		if err != nil {
			return err
		}

		err = tzData[tzdatastructs.Level_01_Idx].Add(tzDataDto)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"tzData[tzdatastructs.Level_01_Idx] Error\n" +
				"Error: %v\n" +
				"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
		}
	}


	// Configure Level-2 Data
	// Configure Level-2 Secondary Group
	// Example: America/Argentina/Buenos_Aires
	groupAlreadyExists, _ = tzGroups[tzdatastructs.Level_02_Idx].ContainsGroupName(zoneArray[1])

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
		
		err = tzGroups[tzdatastructs.Level_02_Idx].Add(tzGroup)

		if err != nil {
			return fmt.Errorf(ePrefix + "tzGroups[tzdatastructs.SecondaryGroupIdx] Error.\n" +
				"Error: %v\n" +
				"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
		}
	}

	containsZone, _ = tzData[tzdatastructs.Level_02_Idx].ContainsTzName(zoneArray[1], zoneArray[2])

	if !containsZone {

		// Add Sub-Zone Time Zone Data Dto
		// Add the Sub Time Zone to the
		// Sub Time Zone Array (subTzDataCol)
		// Example: America/Argentina/Buenos_Aires
		tzDataDto := tzdatastructs.TimeZoneDataDto{}

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
		tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:6]

		// Example: America/Argentina/Buenos_Aires
		// Func Name: Buenos_Aires()
		tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(zoneArray[2])

		tzDataDto.FuncReturnType = "string"

		// Example Function Return Value = "America/Argentina/Buenos_Aires"
		tzDataDto.FuncReturnValue =
			fmt.Sprintf("\"%v\"", tzDataDto.TzCanonicalValue)

		tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
		tzDataDto.TzClass = tzdatastructs.TZClass.Canonical()
		tzDataDto.TzType = tzdatastructs.TZType.SubZone()
		tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
		tzDataDto.SetIsInitialized(true)
		err := tzdeclarations.TzZoneDeclarations{}.PlaceHolderZoneFuncDeclaration(&tzDataDto, ePrefix)

		if err != nil {
			return err
		}

		err = tzData[tzdatastructs.Level_02_Idx].Add(tzDataDto)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"tzData[tzdatastructs.Level_02_Idx] Error.\n" +
				"FileName: %v\n" +
				"Error: %v\n", err.Error(), fMgr.GetFileNameExt())
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
func (parseTz *ParseIanaTzData) zoneCfgValidFuncName(funcName string) string {
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

// isSkipFile - Examines the file name of a time zone data
// file and determines whether the file should be skipped 
// for processing.
//
func (parseTz *ParseIanaTzData) isSkipFile(fMgr pathfileops.FileMgr) (bool, error) {

	ePrefix := "ParseIanaTzData.isSkipFile() "

	err := fMgr.IsFileMgrValid(ePrefix)

	if err != nil {
		return false, err
	}

	if fMgr.GetFileExt() != "" {
		return true, nil
	}

	fileName := strings.ToLower(fMgr.GetFileName())
	isSkipFile := false

	for k:=0; k < len(tzdatastructs.SkipTzFiles); k++ {
		if fileName == strings.ToLower(tzdatastructs.SkipTzFiles[k]) {
			isSkipFile = true
			break
		}
	}

	return isSkipFile, nil
}

// ProcessFileBytes - Process all the bytes in a time zone file
//
func (parseTz *ParseIanaTzData) processFileBytes(fMgr pathfileops.FileMgr) error {

	ePrefix := "ParseIanaTzData.processFileBytes() "

	err := fMgr.OpenThisFileReadOnly()

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	bytes, err := fMgr.ReadAllFile()

	if err != nil {
		_ = fMgr.CloseThisFile()
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	err = fMgr.CloseThisFile()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error closing file. File='%v' Error='%v'\n",
			fMgr.GetAbsolutePathFileName(), err.Error())
	}

	nextStartIdx := 0
	extractedString := ""
	cntr := 1
	for nextStartIdx > -1 {

		extractedString, nextStartIdx = strops.StrOps{}.ReadStringFromBytes(bytes, nextStartIdx)
		fmt.Printf("str No %v: %v\n", cntr, extractedString)
		cntr++

		cmtIdx := strings.Index(extractedString, tzdatastructs.CommentCharStr)

		zoneIdx := strings.Index(extractedString, tzdatastructs.ZoneLabel)

		linkIdx := strings.Index(extractedString, tzdatastructs.LinkLabel)

		if zoneIdx > -1 {

			if cmtIdx > -1 &&
				cmtIdx < zoneIdx {

				continue
			}

			err = parseTz.extractZone(fMgr, extractedString)

			if err != nil {
				fmt.Printf("Zone Extraction Error: %v\n" +
					"%v\n", fMgr.GetAbsolutePathFileName(), err.Error())
			}

			continue
		}

		if linkIdx > -1 {
			if cmtIdx > -1 &&
				cmtIdx < linkIdx {

				continue
			}

			err = parseTz.extractLink(fMgr, extractedString)
			if err != nil {
				fmt.Printf("Link Extraction Error: %v\n" +
					"%v\n", fMgr.GetAbsolutePathFileName(), err.Error())
			}
		}
	}

	return nil
}