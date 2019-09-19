package inprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
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

var tzMajorGroupCol tzdatastructs.TimeZoneGroupCollection

var tzMinorGroupCol tzdatastructs.TimeZoneGroupCollection

var tzDataCol tzdatastructs.TimeZoneDataCollection

var subTzDataCol tzdatastructs.TimeZoneDataCollection

var tzLinkDataCol tzdatastructs.TimeZoneDataCollection

/*
  tzMajorGroupArray Format
    Africa,
    America,
    Antarctica,
    Artic,
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
          MajorGroup  = America
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
          MajorGroup  = America
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
	tzdatastructs.TimeZoneGroupCollection, // Time Zone Major Group Collection
	tzdatastructs.TimeZoneGroupCollection, // Time Zone Minor Group Collection
	tzdatastructs.TimeZoneDataCollection,  // Time Zone Data Collection
	tzdatastructs.TimeZoneDataCollection,  // Sub-Zone Data Collection
	tzdatastructs.TimeZoneDataCollection,  // Alias Link Data Collection
	error)  {

	ePrefix := "ParseIanaTzData.ParseTzAndLinks() "

	tzMajorGroupCol = tzdatastructs.TimeZoneGroupCollection{}.New()
	tzMinorGroupCol = tzdatastructs.TimeZoneGroupCollection{}.New()
	tzDataCol = tzdatastructs.TimeZoneDataCollection{}.New()
	subTzDataCol = tzdatastructs.TimeZoneDataCollection{}.New()
	tzLinkDataCol = tzdatastructs.TimeZoneDataCollection{}.New()

	numOfFiles := dirFileInfo.GetNumOfFileMgrs()

	fmt.Println("Number of Target Files: ", numOfFiles)

	if numOfFiles < 5 {
		return tzMajorGroupCol,
			tzMinorGroupCol,
			tzDataCol,
			subTzDataCol,
			tzLinkDataCol,
			fmt.Errorf(ePrefix+"Number of files is less than 5!\n" +
				"Number of Files='%v'", numOfFiles)
	}



	for i:=0; i < numOfFiles; i++ {

		fMgr, err := dirFileInfo.PeekFileMgrAtIndex(i)

		if err != nil {
			return tzMajorGroupCol,
				tzMinorGroupCol,
				tzDataCol,
				subTzDataCol,
				tzLinkDataCol,
				fmt.Errorf(ePrefix+"%v\n", err.Error())
		}

		isSkipFile, err := parseTz.isSkipFile(fMgr)

		if err != nil {
			return tzMajorGroupCol,
				tzMinorGroupCol,
				tzDataCol,
				subTzDataCol,
				tzLinkDataCol,
				fmt.Errorf(ePrefix+"%v\n", err.Error())
		}

		if isSkipFile {
			continue
		}

		fmt.Println("Valid File: ", fMgr.GetFileNameExt())

		err =  parseTz.processFileBytes(fMgr)

		if err != nil {
			return tzMajorGroupCol,
				tzMinorGroupCol,
				tzDataCol,
				subTzDataCol,
				tzLinkDataCol,
				fmt.Errorf(ePrefix+
					"File Name: %v\n" +
					"Error=%v\n",
					fMgr.GetAbsolutePathFileName(),  err.Error())
		}

	}

	return tzMajorGroupCol,
		tzMinorGroupCol,
		tzDataCol,
		subTzDataCol,
		tzLinkDataCol,
		nil
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
		return parseTz.linkCfgOneElement(fMgr, tzLink, tzCanonical, ePrefix)
	}

	if lenZoneArray == 2 {
		return parseTz.linkCfgTwoElements(fMgr, linkZoneArray, tzCanonical, ePrefix)
	}
	
	// Add Major Group if it does not previously exist
	_, err = tzMajorGroupCol.AddIfNewByDetail(
		linkZoneArray[0],
		"",
		linkZoneArray[0],
		fMgr.GetFileNameExt(),
		tzdatastructs.TzGrpType.IANA(),
		tzdatastructs.DepStatusCode.Valid())

	if err != nil {
		return fmt.Errorf(ePrefix + "\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

	if lenZoneArray == 1 {
		// Link Time Zone Array has one element
		tzDataDto := tzdatastructs.TimeZoneDataDto{}
		tzDataDto.MajorGroup = "Link"
		tzDataDto.SubTzName = ""
		tzDataDto.TzName = tzLink
		tzDataDto.TzAliasValue = tzLink
		tzDataDto.TzCanonicalValue = tzCanonical
		tzDataDto.TzValue = tzLink
		tzDataDto.TzSortValue =
			tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzLink)
		tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
		tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
		tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()

		_, err = tzLinkDataCol.AddIfNew(tzDataDto)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"Error: %v\n" +
				"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
		}

		return nil
	}

	// Link Time Zone Array has two elements
	// Example: 'America/Cayman'
	if lenZoneArray == 2 {
		// This is two element link
		tzDataDto := tzdatastructs.TimeZoneDataDto{}
		tzDataDto.MajorGroup = linkZoneArray[0]
		tzDataDto.SubTzName = ""
		tzDataDto.TzName = linkZoneArray[1]
		tzDataDto.TzAliasValue = tzLink
		tzDataDto.TzCanonicalValue = tzCanonical
		tzDataDto.TzValue = tzCanonical
		tzDataDto.TzSortValue =
			tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzCanonical)
		tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
		tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
		tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()

		_, err = tzDataCol.AddIfNew(tzDataDto)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"Error: %v\n" +
				"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
		}

		return nil
	}

	// lenZoneArray must == 3
	// This is a sub zone
	// America/Argentina/Buenos_Aires

	// Add To Sub-Groups
	// America/Argentina
	_,
		err = tzMinorGroupCol.AddIfNewByDetail(
		linkZoneArray[0],                          // majorGroupName
		linkZoneArray[1],                          // minorGroupName
		linkZoneArray[0] + "/" + linkZoneArray[1], // GroupNameValue
		fMgr.GetFileNameExt(),
		tzdatastructs.TzGrpType.SubGroup(),
		tzdatastructs.DepStatusCode.Valid())

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Add Place Holder To main Tz Collection
	// Add to main time zone collection
	//  America/Argentina/Buenos_Aires
	tzDataDto := tzdatastructs.TimeZoneDataDto{}
	tzDataDto.MajorGroup = linkZoneArray[0] // America - majorGroup
	tzDataDto.SubTzName = linkZoneArray[1]  // Argentina - subTzName
	tzDataDto.TzName = linkZoneArray[1]     // Argentina - tzName
	tzDataDto.TzAliasValue = tzLink
	tzDataDto.TzCanonicalValue = tzCanonical
	tzDataDto.TzValue = tzCanonical
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(
			linkZoneArray[0] + tzdatastructs.ZoneSeparator + linkZoneArray[1])
	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()

	_, err = tzDataCol.AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"FileName: %v\n" +
			"Error: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Finally, add the Sub Time Zone to the
	// Sub Time Zone Array (subTzDataCol)
	// America/Argentina/Buenos_Aires
	tzDataDto = tzdatastructs.TimeZoneDataDto{}
	tzDataDto.MajorGroup = linkZoneArray[0] // America - majorGroup
	tzDataDto.SubTzName = linkZoneArray[1]  // Argentina - subTzName
	tzDataDto.TzName = linkZoneArray[2]     // Argentina - tzName
	tzDataDto.TzAliasValue = tzLink
	tzDataDto.TzCanonicalValue = tzCanonical
	tzDataDto.TzValue = tzCanonical
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(
			tzLink)
	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()

	_,
		err = subTzDataCol.AddIfNew(tzDataDto)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"%v\n", err.Error())
		}

	return nil
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

	// Configure Time Zone Major Group
	tzGroup := tzdatastructs.TimeZoneGroupDto{}
	tzGroup.MajorGroupName = "Deprecated"
	tzGroup.MinorGroupName = ""
	tzGroup.GroupNameValue = tzGroup.MajorGroupName
	tzGroup.GroupSortValue = tzGroup.NewSortValue(tzGroup.MajorGroupName)

	// Example: 'deprecatedTimeZones'
	tzGroup.TypeName =
		strops.StrOps{}.LowerCaseFirstLetter(tzGroup.GroupNameValue)  +
			tzdatastructs.TypeSuffix

	tzGroup.IanaVariableName = tzGroup.GroupNameValue
	tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
	tzGroup.GroupType = tzdatastructs.TzGrpType.IANA()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
	tzGroup.SetIsInitialized(true)

	_, err := tzMajorGroupCol.AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix + "\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

	// Configure Deprecated Link - Time Zone Data Dto
	tzDataDto := tzdatastructs.TimeZoneDataDto{}

	tzDataDto.MajorGroup = "Deprecated"
	tzDataDto.SubTzName = ""
	tzDataDto.TzName = linkZone // Egypt
	tzDataDto.TzAliasValue = linkZone // Egypt
	tzDataDto.TzCanonicalValue = canonicalZone // 'Africa/Cairo'
	tzDataDto.TzValue = tzDataDto.TzCanonicalValue // 'Africa/Cairo'
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(linkZone)

	// func (amer deprecatedTimeZones) Egypt() string { return "Africa/Cairo" }

	// Example: deprecatedTimeZones
	tzDataDto.FuncType =
		strops.StrOps{}.LowerCaseFirstLetter(tzDataDto.MajorGroup) +
			tzdatastructs.TypeSuffix

	// Example: 'depr'
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:5]

	// Example: Egypt()
	tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(linkZone)

	tzDataDto.FuncReturnType = "string"

	// Example Function Return Value = "Africa/Cairo"
	tzDataDto.FuncReturnValue = fmt.Sprintf("\"%v\"", canonicalZone)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
	tzDataDto.SetIsInitialized(true)

	_, err = tzDataCol.AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	return nil
}

// linkCfgTwoElements - Configures and stores data associated
// with a time zone 'Link' which consists of a single link
// string.
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

	// Configure Time Zone Major Group
	tzGroup := tzdatastructs.TimeZoneGroupDto{}
	tzGroup.MajorGroupName = linkZoneArray[0]
	tzGroup.MinorGroupName = ""
	tzGroup.GroupNameValue = linkZoneArray[0]
	tzGroup.GroupSortValue = tzGroup.NewSortValue(linkZoneArray[0])

	// Example: 'uSTimeZones'
	tzGroup.TypeName =
		strops.StrOps{}.LowerCaseFirstLetter(linkZoneArray[0])  +
			tzdatastructs.TypeSuffix

	// Example: 'US'
	tzGroup.IanaVariableName =
		strops.StrOps{}.UpperCaseFirstLetter(linkZoneArray[0])

	tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
	tzGroup.GroupType = tzdatastructs.TzGrpType.IANA()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Deprecated()
	tzGroup.SetIsInitialized(true)

	_, err := tzMajorGroupCol.AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix + "\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

	// Configure Standard Iana Time Zone Data Dto
	tzDataDto := tzdatastructs.TimeZoneDataDto{}

	tzDataDto.MajorGroup = linkZoneArray[0] // US - majorGroup
	tzDataDto.SubTzName = ""
	tzDataDto.TzName = linkZoneArray[1] // Alaska - tzName
	tzDataDto.TzAliasValue = linkZoneArray[0] + tzdatastructs.ZoneSeparator + linkZoneArray[1]
	tzDataDto.TzCanonicalValue = canonicalZone // 'America/Anchorage'

	tzDataDto.TzValue = tzDataDto.TzCanonicalValue // 'America/Anchorage'
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzDataDto.TzName)

	// func (uSTi uSTimeZones) Alaska() string { return "America/Anchorage" }
	// Example: uSTimeZones
	tzDataDto.FuncType =
		strops.StrOps{}.LowerCaseFirstLetter(linkZoneArray[0]) +
			tzdatastructs.TypeSuffix

	// Example: 'uSTi'
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:5]

	// Example: Alaska()
	tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(linkZoneArray[1])

	tzDataDto.FuncReturnType = "string"

	// Example Function Return Value = "America/Anchorage"
	tzDataDto.FuncReturnValue = fmt.Sprintf("\"%v\"", tzDataDto.TzCanonicalValue)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Canonical()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzDataDto.SetIsInitialized(true)

	_, err = tzDataCol.AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	return nil
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

	tzGroup := tzdatastructs.TimeZoneGroupDto{}
	tzGroup.MajorGroupName = zoneArray[0]
	tzGroup.MinorGroupName = ""
	tzGroup.GroupNameValue = zoneArray[0]
	tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
	tzGroup.GroupType = tzdatastructs.TzGrpType.IANA()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()

	_, err = tzMajorGroupCol.AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix+"\n"+
			"FileName: %v\n"+
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error())
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

	// Configure Time Zone Major Group
	tzGroup := tzdatastructs.TimeZoneGroupDto{}
	tzGroup.MajorGroupName = zoneArray[0]
	tzGroup.MinorGroupName = ""
	tzGroup.GroupNameValue = zoneArray[0]
	tzGroup.GroupSortValue = tzGroup.NewSortValue(zoneArray[0])

	// Example: 'americaTimeZones'
	tzGroup.TypeName =
		strops.StrOps{}.LowerCaseFirstLetter(zoneArray[0])  +
			tzdatastructs.TypeSuffix

	// Example: 'America'
	tzGroup.IanaVariableName =
		strops.StrOps{}.UpperCaseFirstLetter(zoneArray[0])

	tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
	tzGroup.GroupType = tzdatastructs.TzGrpType.IANA()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzGroup.SetIsInitialized(true)

	_, err := tzMajorGroupCol.AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix + "\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

	// Configure Standard Iana Time Zone Data Dto
	tzDataDto := tzdatastructs.TimeZoneDataDto{}

	tzDataDto.MajorGroup = zoneArray[0] // America - majorGroup
	tzDataDto.SubTzName = ""
	tzDataDto.TzName = zoneArray[1] // Chicago - tzName
	tzDataDto.TzAliasValue = ""
	tzDataDto.TzCanonicalValue =
		zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1]  // 'America/Chicago'
	tzDataDto.TzValue = tzDataDto.TzCanonicalValue // 'America/Chicago'
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzDataDto.TzCanonicalValue)

	// func (amer americaTimeZones) Chicago() string { return "America/Chicago" }

	// Example: americaTimeZones
	tzDataDto.FuncType =
		strops.StrOps{}.LowerCaseFirstLetter(zoneArray[0]) +
			tzdatastructs.TypeSuffix

	// Example: 'amer'
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:5]

	// Example: Chicago()
	tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(zoneArray[1])

	tzDataDto.FuncReturnType = "string"

	// Example Function Return Value = "America/Chicago"
	tzDataDto.FuncReturnValue = fmt.Sprintf("\"%v\"", tzDataDto.TzCanonicalValue)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Canonical()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzDataDto.SetIsInitialized(true)

	_, err = tzDataCol.AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
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
	// This is a sub zone
	// America/Argentina/Buenos_Aires

	// Configure Time Zone Major Group
	tzGroup := tzdatastructs.TimeZoneGroupDto{}
	tzGroup.MajorGroupName = zoneArray[0]
	tzGroup.MinorGroupName = ""
	tzGroup.GroupNameValue = zoneArray[0]
	tzGroup.GroupSortValue = tzGroup.NewSortValue(zoneArray[0])

	// Example: 'americaTimeZones'
	tzGroup.TypeName =
		strops.StrOps{}.LowerCaseFirstLetter(zoneArray[0])  +
			tzdatastructs.TypeSuffix

	// Example: 'America'
	tzGroup.IanaVariableName =
		strops.StrOps{}.UpperCaseFirstLetter(zoneArray[0])

	tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
	tzGroup.GroupType = tzdatastructs.TzGrpType.IANA()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzGroup.SetIsInitialized(true)

	_, err := tzMajorGroupCol.AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix + "\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

	// Configure Minor Group
	// Example: America/Argentina/Buenos_Aires

	tzGroup = tzdatastructs.TimeZoneGroupDto{}
	tzGroup.MajorGroupName = zoneArray[0] // America
	tzGroup.MinorGroupName = zoneArray[1] // Argentina
	tzGroup.GroupNameValue = zoneArray[1] // Argentina

	tzGroup.GroupSortValue = tzGroup.NewSortValue(zoneArray[1])

	// Example: 'argentinaTimeZones'
	tzGroup.TypeName =
		strops.StrOps{}.LowerCaseFirstLetter(zoneArray[1]) +
			tzdatastructs.TypeSuffix

	// Example: ''
	tzGroup.IanaVariableName = ""

	tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
	tzGroup.GroupType = tzdatastructs.TzGrpType.SubGroup()
	tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzGroup.SetIsInitialized(true)

	_, err = tzMinorGroupCol.AddIfNew(tzGroup)
	if err != nil {
		return fmt.Errorf(ePrefix + "tzMinorGroupCol Error.\n" +
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Add Place Holder TimeZoneDataDto
	// Configure Place Holder Time Zone Data Dto
	// Example Time Zone: America/Argentina/Buenos_Aires
	tzDataDto := tzdatastructs.TimeZoneDataDto{}

	tzDataDto.MajorGroup = zoneArray[0] // America - majorGroup
	tzDataDto.SubTzName = ""
	tzDataDto.TzName = zoneArray[1] // Argentina - tzName
	tzDataDto.TzAliasValue = ""
	tzDataDto.TzCanonicalValue =
		zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1]  // 'America/Argentina'
	tzDataDto.TzValue = tzDataDto.TzCanonicalValue // 'America/Argentina'
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzDataDto.TzCanonicalValue)

	// Example: americaTimeZones
	tzDataDto.FuncType =
		strops.StrOps{}.LowerCaseFirstLetter(zoneArray[0]) +
			tzdatastructs.TypeSuffix

	// Example: 'amer'
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[:5]

	// Example: Argentina()
	tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(zoneArray[1])

	// Example: argentinaTimeZones
	tzDataDto.FuncReturnType =
		strops.StrOps{}.LowerCaseFirstLetter(zoneArray[1]) +
			tzdatastructs.TypeSuffix

	// Example: argentinaTimeZones("")
	tzDataDto.FuncReturnValue = fmt.Sprintf("%v(\"\")",
		tzDataDto.FuncReturnType)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Canonical()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzDataDto.SetIsInitialized(true)

	_, err = tzDataCol.AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Add Sub-Zone Time Zone Data Dto
	// Add the Sub Time Zone to the
	// Sub Time Zone Array (subTzDataCol)
	// America/Argentina/Buenos_Aires
	tzDataDto = tzdatastructs.TimeZoneDataDto{}

	tzDataDto.MajorGroup = zoneArray[0] // America - majorGroup
	tzDataDto.SubTzName = zoneArray[1] // Argentina - subTzName
	tzDataDto.TzName = zoneArray[2] // Buenos_Aires - tzName
	tzDataDto.TzAliasValue = ""

	// America/Argentina/Buenos_Aires
	tzDataDto.TzCanonicalValue =
		zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1] +
			tzdatastructs.ZoneSeparator + zoneArray[2]

	// America/Argentina/Buenos_Aires
	tzDataDto.TzValue = tzDataDto.TzCanonicalValue

	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzDataDto.TzValue)

	// argentinaTimeZones
	tzDataDto.FuncType = strops.StrOps{}.LowerCaseFirstLetter(zoneArray[1]) +
			tzdatastructs.TypeSuffix

	// Example: 'arge'
	tzDataDto.FuncSelfReferenceVariable = tzDataDto.FuncType[0:5]

	// Example: Buenos_Aires()
	tzDataDto.FuncName = parseTz.zoneCfgValidFuncName(zoneArray[2])

	tzDataDto.FuncReturnType = "string"

	// Example Function Return Value = "America/Argentina/Buenos_Aires"
	tzDataDto.FuncReturnValue = fmt.Sprintf("\"%v\"", tzDataDto.TzCanonicalValue)

	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.SubTimeZone()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()
	tzDataDto.SetIsInitialized(true)
	_,
		err = subTzDataCol.AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix + "subTzDataCol Error.\n" +
			"FileName: %v\n" +
			"Error: %v\n", err.Error(), fMgr.GetFileNameExt())
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