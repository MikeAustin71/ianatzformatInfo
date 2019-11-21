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
	b.WriteString("  Id         string\n")
	b.WriteString("  Abbrv      string\n")
	b.WriteString("  TzName     string\n")
	b.WriteString("  Location   string\n")
	b.WriteString("  UtcOffset  string\n")
	b.WriteString("}\n\n\n")

	b.WriteString("// CopyOut() - Makes and returns a deep copy of the current TimeZoneAbbreviationDto\n")
	b.WriteString("// object.\n")
	b.WriteString("// \n")

	b.WriteString("func (TzAbbrv *TimeZoneAbbreviationDto) CopyOut() TimeZoneAbbreviationDto {\n")
	b.WriteString("  \n")
	b.WriteString("  newDto := TimeZoneAbbreviationDto{}\n")
	b.WriteString("  newDto.Id = TzAbbrv.Id\n")
	b.WriteString("  newDto.Abbrv = TzAbbrv.Abbrv\n")
	b.WriteString("  newDto.TzName = TzAbbrv.TzName\n")
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
	b.WriteString("  TzAbbrv.TzName = inComing.TzName\n")
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

// writeMapTzAbbreviationReference - Writes MapTzAbbreviationReference to the
// Time Zone Abbreviations output file.
func (outTzAbbrvs OutputTimeZoneAbbreviations) writeMapTzAbbreviationReference(
	fMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

		ePrefix += "OutputTimeZoneAbbreviations.writeMapTzAbbreviationReference() "

		b := strings.Builder{}

		b.Grow(5120)

		b.WriteString("// MapTzAbbreviationReference - A reference map including all valid\n")
		b.WriteString("// alphabetic Time Zone abbreviations.\n")
		b.WriteString("//\n")

		b.WriteString("var MapTzAbbreviationReference = map[string]TimeZoneAbbreviationDto{\n")

	tzStats.TzAbbreviations.SortByAbbrv()

	numOfAbbrvs := tzStats.TzAbbreviations.GetNumOfAbbreviations()

	if numOfAbbrvs < 1 {
		return fmt.Errorf(ePrefix +
			"\nError: tzStats.TzAbbreviations is EMPTY!\n" +
			"Number of Abbreviations: '%v'\n", numOfAbbrvs)
	}

	var outputStr string

	for i:=0; i < numOfAbbrvs; i++ {

		tzAbrv, err := tzStats.TzAbbreviations.PeekPtr(i)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError Returned by tzStats.TzAbbreviations.PeekPtr(i)\n" +
				"Error='%v'\n", err.Error())
		}

		outputStr = fmt.Sprintf(
			"\"%v\"     :{\"%v\",\"%v\",\"%v\",\"%v\",\"%v\"},\n",
			tzAbrv.Id,
			tzAbrv.Id,
			tzAbrv.Abbrv,
			tzAbrv.TzName,
			tzAbrv.Location,
			tzAbrv.UtcOffset)

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

	b.WriteString("// MapTzAbbrvsToTimeZones - A cross reference that maps\n")
	b.WriteString("// Time Zone Abbreviations to Time Zone Canonical Values.\n")
	b.WriteString("// \n")
	b.WriteString("var MapTzAbbrvsToTimeZones = map[string][]string {\n")

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

	b.WriteString("// MapTimeZonesToTzAbbrvs - A cross reference that maps\n")
	b.WriteString("// Time Zone Canonical Values to Time Zone Abbreviations.\n")
	b.WriteString("// \n")
	b.WriteString("var MapTimeZonesToTzAbbrvs = map[string][]string {\n")

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