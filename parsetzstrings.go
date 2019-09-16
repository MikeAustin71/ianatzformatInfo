package main

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
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

var tzMajorGroupCol TimeZoneGroupCollection

var tzMinorGroupCol TimeZoneGroupCollection

var tzDataCol TimeZoneDataCollection

var subTzDataCol TimeZoneDataCollection

var tzLinkDataCol TimeZoneDataCollection

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
	TimeZoneGroupCollection, // Time Zone Major Group Collection
	TimeZoneGroupCollection, // Time Zone Minor Group Collection
	TimeZoneDataCollection,  // Time Zone Data Collection
	TimeZoneDataCollection,  // Sub-Zone Data Collection
	TimeZoneDataCollection,  // Alias Link Data Collection
	error)  {

	ePrefix := "ParseIanaTzData.ParseTzAndLinks() "

	tzMajorGroupCol = TimeZoneGroupCollection{}.New()
	tzMinorGroupCol = TimeZoneGroupCollection{}.New()
	tzDataCol = TimeZoneDataCollection{}.New()
	subTzDataCol = TimeZoneDataCollection{}.New()
	tzLinkDataCol = TimeZoneDataCollection{}.New()

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
			[]string{LinkLabel},
			startIdx,
			LeadingFieldSeparators,
			TrailingFieldSeparators,
			CommentDelimiters,
			EndOfLineDelimiters)

	if err != nil {
		return fmt.Errorf(ePrefix + "%v\n", err.Error())
	}

	if dFProfile.DataFieldLength < 1 {
		return nil
	}

	if strings.Index(dFProfile.DataFieldStr, ZoneSeparator) == -1 {
		return nil
	}

	tzCanonical := dFProfile.DataFieldStr
	startIdx = dFProfile.NextTargetStrIndex

	// Extract Field 2 - Link Field
	dFProfile,
	err =
		strops.StrOps{}.ExtractDataField(
			rawString,
			[]string{LinkLabel},
			startIdx,
			LeadingFieldSeparators,
			TrailingFieldSeparators,
			CommentDelimiters,
			EndOfLineDelimiters)

	if err != nil {
		return fmt.Errorf(ePrefix + "%v\n", err.Error())
	}

	if dFProfile.DataFieldLength < 1 {
		return nil
	}

	if strings.Index(dFProfile.DataFieldStr, ZoneSeparator) == -1 {
		return nil
	}

	tzLink := dFProfile.DataFieldStr

	zoneArray := strings.Split(tzLink, ZoneSeparator)

	lenZoneArray := len(zoneArray)

	if lenZoneArray < 2 ||
		lenZoneArray > 3 {
		fmt.Printf(ePrefix + "Invalid Link Time Zone!\n" +
			"FileName: %v\n" +
			"Tz Link String: %v\n" +
			"Tz Canonical String: %v\n",
			fMgr.GetFileNameExt(), tzLink, tzCanonical)
		return nil
	}

	_, err = tzMajorGroupCol.AddIfNewByDetail(
		zoneArray[0],
		"",
		zoneArray[0],
		fMgr.GetFileNameExt(),
		TzGrpType.IANA(),
		DepStatusCode.Valid())

	if err != nil {
		return fmt.Errorf(ePrefix + "\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

	// Link Time Zone Array has two elements
	// Example: 'America/Cayman'
	if lenZoneArray == 2 {

		_, err = tzDataCol.AddIfNewByDetail(
			zoneArray[0],
			"",
			zoneArray[1],
			tzCanonical,
			fMgr.GetFileNameExt(),
			TZClass.Alias(),
			DepStatusCode.Valid())

		if err != nil {
			return fmt.Errorf(ePrefix +
				"Error: %v\n" +
				"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
		}

		_, err = tzLinkDataCol.AddIfNewByDetail(
			zoneArray[0],
			"",
			zoneArray[1],
			tzCanonical,
			fMgr.GetFileNameExt(),
			TZClass.Alias(),
			DepStatusCode.Valid())

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
		TzGrpType.SubGroup(),
		DepStatusCode.Valid())

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Add Place Holder To main Tz Collection
	// Add to main time zone collection
	//  America/Argentina/Buenos_Aires
	_, err = tzDataCol.AddIfNewByDetail(
		zoneArray[0], // America - majorGroup
		zoneArray[1], // Argentina - subTzName
		zoneArray[1], // Argentina - tzName
		tzCanonical, // America/Argentina - tzValue
		fMgr.GetFileNameExt(),
		TZClass.SubGroup(),
		DepStatusCode.Valid())

	if err != nil {
		return fmt.Errorf(ePrefix +
			"FileName: %v\n" +
			"Error: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Finally, add the Sub Time Zone to the
	// Sub Time Zone Array (subTzDataCol)
	// America/Argentina/Buenos_Aires

	_,
		err = subTzDataCol.AddIfNewByDetail(
		zoneArray[0],  // America - majorGroup
		zoneArray[1],  // Argentina - subTzName
		zoneArray[0] + "/" + zoneArray[1],  // America/Argentina - tzName
		tzCanonical, // America/Argentina/Buenos_Aires - tzValue
		fMgr.GetFileNameExt(),
		TZClass.SubTimeZone(),
		DepStatusCode.Valid())

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
			[]string{ZoneLabel},
			0,
			LeadingFieldSeparators,
			TrailingFieldSeparators,
			CommentDelimiters,
			EndOfLineDelimiters)

	if err != nil {
		return fmt.Errorf(ePrefix + "%v\n", err.Error())
	}

	if dFProfile.DataFieldLength < 1 {
		return nil
	}

	if strings.Index(dFProfile.DataFieldStr, ZoneSeparator) == -1 {
		return nil
	}

	zoneArray := strings.Split(dFProfile.DataFieldStr, ZoneSeparator)

	lenZoneArray := len(zoneArray)

	if lenZoneArray < 2 ||
		lenZoneArray > 3 {
		fmt.Printf(ePrefix + "Invalid Time Zone!\n" +
			"FileName: %v\n" +
			"Raw Zone String: %v\n", fMgr.GetFileNameExt(), rawString)
		return nil
	}


	_, err = tzMajorGroupCol.AddIfNewByDetail(
		zoneArray[0],
		"",
		zoneArray[0],
		fMgr.GetFileNameExt(),
		TzGrpType.IANA(),
		DepStatusCode.Valid())

	if err != nil {
		return fmt.Errorf(ePrefix + "\n" +
			"FileName: %v\n" +
			"Error: %v\n", fMgr.GetFileNameExt(), err.Error() )
	}

// Zone Array has two elements
// Example: 'America/Chicago'
	if lenZoneArray == 2 {

		_, err = tzDataCol.AddIfNewByDetail(
			zoneArray[0],
			"",
			zoneArray[1],
			zoneArray[0] + "/" + zoneArray[1],
			fMgr.GetFileNameExt(),
			TZClass.Canonical(),
			DepStatusCode.Valid())

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
	TzGrpType.SubGroup(),
	DepStatusCode.Valid())

	if err != nil {
		return fmt.Errorf(ePrefix + "tzMinorGroupCol Error.\n" +
			"Error: %v\n" +
			"FileName: %v\n", err.Error(), fMgr.GetFileNameExt())
	}

	// Add Place Holder To main Tz Collection
	// Add to main time zone collection
	//  America/Argentina/Buenos_Aires
		_, err = tzDataCol.AddIfNewByDetail(
			zoneArray[0], // America - majorGroup
			zoneArray[1], // Argentina - subTzName
			zoneArray[1], // Argentina - tzName
			zoneArray[0] + "/" + zoneArray[1], // America/Argentina - tzValue
			fMgr.GetFileNameExt(),
			TZClass.SubGroup(),
			DepStatusCode.Valid())

		if err != nil {
			return fmt.Errorf(ePrefix + "tzDataCol Error.\n" +
				"FileName: %v\n" +
				"Error: %v\n", err.Error(), fMgr.GetFileNameExt())
		}

	// Finally, add the Sub Time Zone to the
	// Sub Time Zone Array (subTzDataCol)
	// America/Argentina/Buenos_Aires

	_,
	err = subTzDataCol.AddIfNewByDetail(
		zoneArray[0],  // America - majorGroup
		zoneArray[1],  // Argentina - subTzName
		zoneArray[0] + "/" + zoneArray[1],  // America/Argentina - tzName
		zoneArray[0] + "/" + zoneArray[1] + "/" + zoneArray[2], // America/Argentina/Buenos_Aires - tzValue
		fMgr.GetFileNameExt(),
		TZClass.SubTimeZone(),
		DepStatusCode.Valid())

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

	for k:=0; k < len(skipTzFiles); k++ {
		if fileName == strings.ToLower(skipTzFiles[k]) {
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

		cmtIdx := strings.Index(extractedString, CommentCharStr)

		zoneIdx := strings.Index(extractedString, ZoneLabel)

		linkIdx := strings.Index(extractedString, LinkLabel)

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