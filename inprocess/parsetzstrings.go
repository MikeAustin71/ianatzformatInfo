package inprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
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

	zoneArray := strings.Split(tzLink, tzdatastructs.ZoneSeparator)

	lenZoneArray := len(zoneArray)

	if lenZoneArray < 1 ||
		lenZoneArray > 3 {
		fmt.Printf(ePrefix + "Invalid Link Time Zone!\n" +
			"FileName: %v\n" +
			"Tz Link String: %v\n" +
			"Tz Canonical String: %v\n",
			fMgr.GetFileNameExt(), tzLink, tzCanonical)
		return nil
	}

	// Add Major Group if it does not previously exist
	_, err = tzMajorGroupCol.AddIfNewByDetail(
		zoneArray[0],
		"",
		zoneArray[0],
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
		tzDataDto.MajorGroup = zoneArray[0]
		tzDataDto.SubTzName = ""
		tzDataDto.TzName = zoneArray[1]
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
		zoneArray[0],  // majorGroupName
		zoneArray[1],  // minorGroupName
		zoneArray[0] + "/" + zoneArray[1], // CompositeGroupName
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
	tzDataDto.MajorGroup = zoneArray[0] // America - majorGroup
	tzDataDto.SubTzName = zoneArray[1] // Argentina - subTzName
	tzDataDto.TzName = zoneArray[1] // Argentina - tzName
	tzDataDto.TzAliasValue = tzLink
	tzDataDto.TzCanonicalValue = tzCanonical
	tzDataDto.TzValue = tzCanonical
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(
			zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1])
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
	tzDataDto.MajorGroup = zoneArray[0] // America - majorGroup
	tzDataDto.SubTzName = zoneArray[1] // Argentina - subTzName
	tzDataDto.TzName = zoneArray[2] // Argentina - tzName
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

// extractZone - Extracts standard time zones and sub time zones.
// Data is stored in tzMajorGroupMap, tzDataCol and
// or subTzDataCol.
func (parseTz *ParseIanaTzData) extractZone(
	fMgr pathfileops.FileMgr, rawString string) error  {

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
		return fmt.Errorf(ePrefix + "%v\n", err.Error())
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
		fmt.Printf(ePrefix + "Invalid Time Zone!\n" +
			"FileName: %v\n" +
			"Raw Zone String: %v\n", fMgr.GetFileNameExt(), rawString)
		return nil
	}

tzGroup := tzdatastructs.TimeZoneGroupDto{}
tzGroup.MajorGroupName = zoneArray[0]
tzGroup.MinorGroupName = ""
tzGroup.CompositeGroupName = zoneArray[0]
tzGroup.SourceFileNameExt = fMgr.GetFileNameExt()
tzGroup.GroupType = tzdatastructs.TzGrpType.IANA()
tzGroup.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()

	_, err = tzMajorGroupCol.AddIfNew(tzGroup)

	if err != nil {
		return fmt.Errorf(ePrefix + "\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

// Zone Array has two elements
// Example: 'America/Chicago'
	if lenZoneArray == 2 {

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
		tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
		tzDataDto.TzClass = tzdatastructs.TZClass.Canonical()
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
	zoneArray[0],  // majorGroupName
	zoneArray[1],  // minorGroupName
	zoneArray[0] + "/" + zoneArray[1], // CompositeGroupName
	fMgr.GetFileNameExt(),
	tzdatastructs.TzGrpType.SubGroup(),
	tzdatastructs.DepStatusCode.Valid())

	if err != nil {
		return fmt.Errorf(ePrefix + "tzMinorGroupCol Error.\n" +
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Add Place Holder To main Tz Collection
	// Add to main time zone collection
	//  America/Argentina/Buenos_Aires
	tzDataDto := tzdatastructs.TimeZoneDataDto{}

	tzDataDto.MajorGroup = zoneArray[0] // America - majorGroup
	tzDataDto.SubTzName = zoneArray[1] // Argentina - subTzName
	tzDataDto.TzName = zoneArray[1] // Argentina - tzName
	tzDataDto.TzAliasValue = ""
	tzDataDto.TzCanonicalValue =
		zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1]
	tzDataDto.TzValue = zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1]
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(
			zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1])
	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.Alias()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()

	_, err = tzDataCol.AddIfNew(tzDataDto)

		if err != nil {
			return fmt.Errorf(ePrefix + "tzDataCol Error.\n" +
				"FileName: %v\n" +
				"Error: %v\n", err.Error(), fMgr.GetFileNameExt())
		}

	// Finally, add the Sub Time Zone to the
	// Sub Time Zone Array (subTzDataCol)
	// America/Argentina/Buenos_Aires
	tzDataDto = tzdatastructs.TimeZoneDataDto{}

	tzDataDto.MajorGroup = zoneArray[0] // America - majorGroup
	tzDataDto.SubTzName = zoneArray[1] // Argentina - subTzName
	tzDataDto.TzName = zoneArray[1] // Argentina - tzName
	tzDataDto.TzAliasValue = ""
	tzDataDto.TzCanonicalValue =
		zoneArray[0] + tzdatastructs.ZoneSeparator + zoneArray[1]
	tzDataDto.TzValue = zoneArray[0] + tzdatastructs.ZoneSeparator +
		zoneArray[1] + tzdatastructs.ZoneSeparator + zoneArray[2]
	tzDataDto.TzSortValue =
		tzdatastructs.TimeZoneDataDto{}.NewSortValue(tzDataDto.TzValue)
	tzDataDto.SourceFileNameExt = fMgr.GetFileNameExt()
	tzDataDto.TzClass = tzdatastructs.TZClass.SubTimeZone()
	tzDataDto.DeprecationStatus = tzdatastructs.DepStatusCode.Valid()

	_,
	err = subTzDataCol.AddIfNew(tzDataDto)

	if err != nil {
		return fmt.Errorf(ePrefix + "subTzDataCol Error.\n" +
			"FileName: %v\n" +
			"Error: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	return nil
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