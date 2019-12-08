package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/fileops"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"sort"
	"strings"
)

type OutputTimeZoneAbbreviations struct {
	input          string
	output         string
}

// WriteOutput - Writes Time Zone Abbreviations data structures and
// types to output file 'timezoneabbreviations.go'.
func (outTzAbbrvs OutputTimeZoneAbbreviations) WriteOutput(
	zoneInfoDto tzdatastructs.ZoneInfoDataDto,
	tzStats *tzdatastructs.TimeZoneStatsDto, // Time Zone Version
	ePrefix string) error {

	ePrefix += "OutputTimeZoneAbbreviations.WriteOutput() "

	f,
	err := fileops.FileOps{}.CreateOpenFile(
		zoneInfoDto.AppOutputDirMgr,
		tzdatastructs.TzAbbrvDataOutputFileName,
		ePrefix)

	if err != nil {
		return err
	}

	err = outTzAbbrvs.writeHeader(f, ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	err = outTzAbbrvs.writeTimeZoneAbbreviationDto(f, ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	err = outTzAbbrvs.writeStdTZoneAbbreviationsType(f,ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}


	err = outTzAbbrvs.writeMapTzAbbreviationReference(f, tzStats, ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	err = outTzAbbrvs.writeMapTzAbbrvsToTimeZones(f, tzStats, ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	err = outTzAbbrvs.writeMapTimeZonesToTzAbbrvs(f, tzStats, ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	errs := make([]error, 0)

	err = f.FlushBytesToDisk()

	if err != nil {
		errs = append(errs,
			fmt.Errorf(ePrefix +
				"\nError returned by f.FlushBytesToDisk()\n" +
				"Error='%v'\n", err.Error()))
	}

	err = f.CloseThisFile()

	if err != nil {
		errs = append(errs,
			fmt.Errorf(ePrefix +
				"\nError returned by f.CloseThisFile()\n" +
				"Error='%v'\n",err.Error()))
	}

	return pathfileops.FileHelper{}.ConsolidateErrors(errs)
}

// writeHeader - Writes header information to the Time Zone
// Abbreviation output file.
func (outTzAbbrvs OutputTimeZoneAbbreviations) writeHeader(
	fMgr pathfileops.FileMgr,
	ePrefix string) error {

	ePrefix += "OutputTimeZoneAbbreviations.writeHeader() "

	b := strings.Builder{}

	b.Grow(256)

	b.WriteString("package main\n\n")
	b.WriteString("import (\n")
	b.WriteString("  \"errors\"\n")
	b.WriteString("  \"sync\"\n")
	b.WriteString(")\n\n\n")

	_, err := fMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = fMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

// writeTimeZoneAbbreviationDto - Writes type TimeZoneAbbreviationDto to the
// output file.
func (outTzAbbrvs OutputTimeZoneAbbreviations) writeTimeZoneAbbreviationDto(
	fMgr pathfileops.FileMgr,
	ePrefix string) error {

	ePrefix += "OutputTimeZoneAbbreviations.writeTimeZoneAbbreviationDto() "

	b := strings.Builder{}

	b.Grow(2048)

	b.WriteString("// TzAbbreviationDto - encapsulates Time Zone abbreviation\n")

	b.WriteString("// information. A Time Zone Abbreviation must consist entirely\n")

	b.WriteString("// of alphabetic characters.\n")

	b.WriteString("// \n")

	b.WriteString("// The Id is styled as Abbreviation text plus the UTC offset.\n")

	b.WriteString("// Example: CST-0600 - Central Standard time with offset UTC-0600.\n")

	b.WriteString("// \n")

	b.WriteString("type TimeZoneAbbreviationDto struct {\n")
	b.WriteString("  Id                 string  // Example: \"CST-0600\"\n")
	b.WriteString("  Abbrv              string  // Example: \"CST\"\n")
	b.WriteString("  AbbrvDescription   string  // Example: \"Central Standard Time\"\n")
	b.WriteString("  Location           string  // Example: \"North America\"\n")
	b.WriteString("  UtcOffset          string  // Example: \"-0600\"\n")
	b.WriteString("}\n\n\n")

	b.WriteString("// CopyOut() - Makes and returns a deep copy of the current TimeZoneAbbreviationDto\n")
	b.WriteString("// object.\n")
	b.WriteString("// \n")

	b.WriteString("func (TzAbbrv *TimeZoneAbbreviationDto) CopyOut() TimeZoneAbbreviationDto {\n")
	b.WriteString("  \n")
	b.WriteString("  newDto := TimeZoneAbbreviationDto{}\n")
	b.WriteString("  newDto.Id = TzAbbrv.Id\n")
	b.WriteString("  newDto.Abbrv = TzAbbrv.Abbrv\n")
	b.WriteString("  newDto.AbbrvDescription = TzAbbrv.AbbrvDescription\n")
	b.WriteString("  newDto.Location = TzAbbrv.Location\n")
	b.WriteString("  newDto.UtcOffset = TzAbbrv.UtcOffset\n")
	b.WriteString("  \n")
	b.WriteString("  return newDto\n")
	b.WriteString("}\n\n\n")


	b.WriteString("// CopyIn() - Copies the field values from an incoming TimeZoneAbbreviationDto\n")
	b.WriteString("// object to the current TimeZoneAbbreviationDto object.\n")
	b.WriteString("// \n")
	b.WriteString("func (TzAbbrv *TimeZoneAbbreviationDto) CopyIn(inComing *TimeZoneAbbreviationDto) error {\n")
	b.WriteString("  \n")
	b.WriteString("  ePrefix := \"TzAbbreviationDto.CopyIn()\" \n")
	b.WriteString("  \n")
	b.WriteString("  if inComing == nil {\n")
	b.WriteString("    return  errors.New(ePrefix +\n")
	b.WriteString("      \"Error: Input parameter 'incoming' is nil!\")")
	b.WriteString("  }\n")
	b.WriteString("  \n")

	b.WriteString("  TzAbbrv.Id = inComing.Id\n")
	b.WriteString("  TzAbbrv.Abbrv = inComing.Abbrv\n")
	b.WriteString("  TzAbbrv.AbbrvDescription = inComing.AbbrvDescription\n")
	b.WriteString("  TzAbbrv.Location = inComing.Location\n")
	b.WriteString("  TzAbbrv.UtcOffset = inComing.UtcOffset\n")
	b.WriteString("  return nil\n")
	b.WriteString("}  \n\n\n")

	_, err := fMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = fMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

func (outTzAbbrvs OutputTimeZoneAbbreviations) writeStdTZoneAbbreviationsType(
	fMgr pathfileops.FileMgr,
	ePrefix string) error {

	ePrefix += "OutputTimeZoneAbbreviations.writeStdTZoneAbbreviationsType() "

	b := strings.Builder{}

	b.Grow(5120)

	b.WriteString(
		"// StdTZoneAbbreviations - Provides thread safe access to\n")

	b.WriteString(
		"// standard IANA Time Zone abbreviations, abbreviation\n")

	b.WriteString(
		"// descriptions and UTC Offsets.\n//  \n")

	b.WriteString(
		"type StdTZoneAbbreviations struct {\n")

	b.WriteString(
		"  Input                  TimeZoneAbbreviationDto\n")

	b.WriteString(
		"  Output                 TimeZoneAbbreviationDto\n")

	b.WriteString(
		"  lock                   sync.Mutex\n")

	b.WriteString("}\n\n\n")

	b.WriteString(
		"// AbbrvOffsetToTzReference - This method returns a type\n")

	b.WriteString(
		"// 'TimeZoneAbbreviation' describing a specific time zone\n")

	b.WriteString(
		"// abbreviation based on an input parameter consisting of\n")

	b.WriteString(
		"// an alphabetic time zone abbreviation and an UTC offset\n")

	b.WriteString(
		"// parameter.\n")

	b.WriteString(
		"//  \n")

	b.WriteString(
		"// The Time Zone Abbreviation Offset parameter, 'abbrvOffset',\n")

	b.WriteString(
		"// must be formatted with a time zone abbreviation in all\n")

	b.WriteString(
		"// upper case characters followed by the UTC Offset expressed\n")

	b.WriteString(
		"// in hours and minutes.\n")

	b.WriteString(
		"\n")

	b.WriteString(
		"// For example, to return a 'TimeZoneAbbreviationDto' describing\n")

	b.WriteString(
		"// North America Central Standard Time, the 'abbrvOffset' input\n")

	b.WriteString(
		"// parameter must be formatted as 'CST-0600'. Note: the UTC\n")

	b.WriteString(
		"// offset for North America Central Standard Time is 'UTC-0600'.\n")

	b.WriteString(
		"\n")

	b.WriteString(
		"// If the Abbreviation Offset parameter is invalid or if no\n")

	b.WriteString(
		"// 'TimeZoneAbbreviationDto' exists for the Abbreviation Offset\n")

	b.WriteString(
		"// parameter, this method will return a boolean value of 'false'.\n")

	b.WriteString(
		"//  \n")

	b.WriteString(
		"func (stdTzAbbrvs *StdTZoneAbbreviations) AbbrvOffsetToTzReference(\n")

	b.WriteString(
		"		abbrvOffset string) (TimeZoneAbbreviationDto, bool) {\n\n")

	b.WriteString(
		"	stdTzAbbrvs.lock.Lock()\n\n")

	b.WriteString(
		"	defer stdTzAbbrvs.lock.Unlock()\n\n")


	b.WriteString(
		"	result, ok := mapTzAbbreviationReference[abbrvOffset]\n\n")


	b.WriteString(
		"	return result, ok\n")


	b.WriteString(
		"}\n\n")


	b.WriteString(
		"// AbbrvOffsetToTimeZones - Returns a string array consisting of\n")


	b.WriteString(
		"// all standard time zones associated with a specific time zone\n")


	b.WriteString(
		"// abbreviation based on an input parameter consisting of an\n")


	b.WriteString(
		"// alphabetic time zone abbreviation and an UTC offset parameter.\n")


	b.WriteString(
		"// \n")


	b.WriteString(
		"// The Time Zone Abbreviation Offset parameter, 'abbrvOffset'\n")


	b.WriteString(
		"// must be formatted with a time zone abbreviation in all\n")


	b.WriteString(
		"// upper case characters followed by the UTC Offset expressed\n")


	b.WriteString(
		"// in hours and minutes.\n")


	b.WriteString(
		"// \n")


	b.WriteString(
		"// For example, to return a string array containing all standard\n")


	b.WriteString(
		"// time zones associated with North America Central Standard\n")


	b.WriteString(
		"// Time, the 'abbrvOffset' input parameter must be formatted as\n")


	b.WriteString(
		"// 'CST-0600'. Note: the UTC offset for North America Central\n")


	b.WriteString(
		"// Standard Time is 'UTC-0600'.\n")


	b.WriteString(
		"// \n")


	b.WriteString(
		"// If the Time Zone Abbreviation Offset parameter is invalid or\n")


	b.WriteString(
		"// if no string array exists for the Abbreviation Offset\n")


	b.WriteString(
		"// parameter, this method will return a boolean value of 'false'.\n")


	b.WriteString(
		"// \n")


	b.WriteString(
		"func (stdTzAbbrvs *StdTZoneAbbreviations) AbbrvOffsetToTimeZones(\n")


	b.WriteString(
		"		abbrvOffset string) ([]string, bool) {\n\n")


	b.WriteString(
		"	stdTzAbbrvs.lock.Lock()\n\n")


	b.WriteString(
		"	defer stdTzAbbrvs.lock.Unlock()\n\n")


	b.WriteString(
		"	result, ok :=  mapTzAbbrvsToTimeZones[abbrvOffset]\n\n")

	b.WriteString(
		"	return result, ok\n")

	b.WriteString(
		"}\n\n")

	b.WriteString(
		"// TimeZonesToAbbrvs - Returns a string array consisting of\n")

	b.WriteString(
		"// all time zone abbreviations associated with a standard,\n")

	b.WriteString(
		"// IANA time zone name passed as an input parameter.  This\n")

	b.WriteString(
		"// input parameter, 'timeZone', must be formatted as a\n")

	b.WriteString(
		"// standard IANA time zone name using upper and lower case\n")

	b.WriteString(
		"// characters as specified in the IANA Time Zone Database.\n")

	b.WriteString(
		"// \n")

	b.WriteString(
		"// The returned string array actually contains Time Zone\n")

	b.WriteString(
		"// Abbreviation and UTC Offset pairs.\n")

	b.WriteString(
		"// \n")

	b.WriteString(
		"// For example, the standard IANA Time Zone, 'America/Chicago'\n")

	b.WriteString(
		"// will return a string array consisting of two strings:\n")

	b.WriteString(
		"// \"CDT-0500\" and \"CST-0600\". These two strings describe the\n")

	b.WriteString(
		"// two time zone abbreviations associated with 'America/Chicago'\n")

	b.WriteString(
		"// (a.k.a. North America Central Time). \"CDT-0500\" stands for\n")

	b.WriteString(
		"// 'Central Daylight Time' and UTC-0500 (a UTC offset of\n")

	b.WriteString(
		"// -5 hours). Likewise, \"CST-0600\" identifies 'Central Standard\n")

	b.WriteString(
		"// Time' with an UTC offset of -6 hours (UTC-0600).\n")

	b.WriteString(
		"// \n")

	b.WriteString(
		"// If the Time Zone input parameter is invalid or if no string\n")

	b.WriteString(
		"// array exists for the Time Zone input parameter, this method\n")

	b.WriteString(
		"// will return a boolean value of 'false'.\n")

	b.WriteString(
		"// \n")

	b.WriteString(
		"func (stdTzAbbrvs *StdTZoneAbbreviations) TimeZonesToAbbrvs(\n")

	b.WriteString(
		"		timeZone string) ([]string, bool) {\n\n")

	b.WriteString(
		"	stdTzAbbrvs.lock.Lock()\n\n")

	b.WriteString(
		"	defer stdTzAbbrvs.lock.Unlock()\n\n")

	b.WriteString(
		"	result, ok := mapTimeZonesToTzAbbrvs[timeZone]\n\n")

	b.WriteString(
		"	return result, ok\n")

	b.WriteString(
		"}\n\n\n")

	_, err := fMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = fMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

// writeMapTzAbbreviationReference - Writes MapTzAbbreviationReference to the
// Time Zone Abbreviations output file.
func (outTzAbbrvs OutputTimeZoneAbbreviations) writeMapTzAbbreviationReference(
	fMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

		ePrefix += "OutputTimeZoneAbbreviations.writeMapTzAbbreviationReference() "

		b := strings.Builder{}

		b.Grow(5120)

		b.WriteString("// mapTzAbbreviationReference - A reference map including all valid\n")
		b.WriteString("// alphabetic Time Zone abbreviations.\n")
		b.WriteString("//\n")

		b.WriteString("var mapTzAbbreviationReference = map[string]TimeZoneAbbreviationDto{\n")

	tzStats.TzAbbreviations.SortByAbbrv()

	numOfAbbrvs := tzStats.TzAbbreviations.GetNumOfAbbreviations()

	if numOfAbbrvs < 1 {
		return fmt.Errorf(ePrefix +
			"\nError: tzStats.TzAbbreviations is EMPTY!\n" +
			"Number of Abbreviations: '%v'\n", numOfAbbrvs)
	}

	var outputStr string

	for i:=0; i < numOfAbbrvs; i++ {

		tzAbbrv, err := tzStats.TzAbbreviations.PeekPtr(i)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError Returned by tzStats.TzAbbreviations.PeekPtr(i)\n" +
				"Error='%v'\n", err.Error())
		}

		outputStr = fmt.Sprintf(
			"\"%v\"     :{\"%v\",\"%v\",\"%v\",\"%v\",\"%v\"},\n",
			tzAbbrv.Id,
			tzAbbrv.Id,
			tzAbbrv.Abbrv,
			tzAbbrv.TzName,
			tzAbbrv.Location,
			tzAbbrv.UtcOffset)

		b.WriteString(outputStr)
	}

	b.WriteString("}\n\n\n")

	_, err := fMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = fMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

// writeMapTzAbbrvsToTimeZones - A cross reference that maps
// Time Zone Abbreviations to Time Zone Canonical Values.
//
func (outTzAbbrvs OutputTimeZoneAbbreviations) writeMapTzAbbrvsToTimeZones(
	fMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "OutputTimeZoneAbbreviations.writeMapTzAbbrvsToTimeZones() "

	b := strings.Builder{}

	b.Grow(5120)

	b.WriteString("// mapTzAbbrvsToTimeZones - A cross reference that maps\n")
	b.WriteString("// Time Zone Abbreviations to Time Zone Canonical Values.\n")
	b.WriteString("// \n")
	b.WriteString("var mapTzAbbrvsToTimeZones = map[string][]string {\n")

	abbrvIds := make([]string, 0)

	for idx := range tzStats.MapTzAbbrvsToTimeZones {

		abbrvIds = append(abbrvIds, idx)

	}

	sort.Strings(abbrvIds)

	for i:=0; i < len(abbrvIds); i++ {

		tzCanonicalValues, ok := tzStats.MapTzAbbrvsToTimeZones[abbrvIds[i]]

		if !ok {
			return fmt.Errorf(ePrefix +
				"\nError: Expected Valid Map Entry for Tz Abbreviation '%v'.\n" +
				"However, this entry is INVALID!\n")
		}

		outputStr := fmt.Sprintf("\"%v\"     :{ ",
			abbrvIds[i])

		lenTzCanonicalValues := len(tzCanonicalValues)

		for j:=0; j < lenTzCanonicalValues; j++ {

			tzStr := fmt.Sprintf("\"%v\"", tzCanonicalValues[j])

			if j == lenTzCanonicalValues - 1 {
				tzStr += "},\n"
			} else {
				tzStr += ","
			}

			outputStr += tzStr
		}

		b.WriteString(outputStr)
	}

	b.WriteString("}\n\n\n")

	_, err := fMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
	return fmt.Errorf(ePrefix +
	"\n Error returned by fMgr.WriteBytesToFile([]byte(b.String()))\n" +
	"Error='%v'\n", err.Error())
	}

	err = fMgr.FlushBytesToDisk()

	if err != nil {
	return fmt.Errorf(ePrefix +
	"\n Error returned by fMgr.FlushBytesToDisk()\n" +
	"Error='%v'\n", err.Error())
	}

	return nil
}

// writeMapTimeZonesToTzAbbrvs - A cross reference that maps
// Time Zone Canonical Values to Time Zone Abbreviations.
//
func (outTzAbbrvs OutputTimeZoneAbbreviations) writeMapTimeZonesToTzAbbrvs(
	fMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "OutputTimeZoneAbbreviations.writeMapTimeZonesToTzAbbrvs() "

	b := strings.Builder{}

	b.Grow(5120)

	b.WriteString("// mapTimeZonesToTzAbbrvs - A cross reference that maps\n")
	b.WriteString("// Time Zone Canonical Values to Time Zone Abbreviations.\n")
	b.WriteString("// \n")
	b.WriteString("var mapTimeZonesToTzAbbrvs = map[string][]string {\n")

	timeZoneCanonicalValues := make([]string ,0)

	for idx := range tzStats.MapTimeZonesToTzAbbrvs {

		timeZoneCanonicalValues = append(timeZoneCanonicalValues, idx)

	}

	sort.Strings(timeZoneCanonicalValues)

	for i:=0; i < len(timeZoneCanonicalValues); i++ {

		tzAbbrvValues, ok :=
			tzStats.MapTimeZonesToTzAbbrvs[timeZoneCanonicalValues[i]]

		if !ok {
			return fmt.Errorf(ePrefix +
				"\nError: Expected Valid Map Entry for Tz Abbreviation '%v'.\n" +
				"However, this entry is INVALID!\n")
		}

		outputStr := fmt.Sprintf("\"%v\"     :{ ",
			timeZoneCanonicalValues[i])

		lenTzAbbrvValues := len(tzAbbrvValues)

		for j:=0; j < lenTzAbbrvValues; j++ {

			tzStr := fmt.Sprintf("\"%v\"", tzAbbrvValues[j])

			if j == lenTzAbbrvValues - 1 {
				tzStr += "},\n"
			} else {
				tzStr += ","
			}

			outputStr += tzStr
		}

		b.WriteString(outputStr)
	}

	b.WriteString("}\n\n\n")

	_, err := fMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = fMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by fMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}