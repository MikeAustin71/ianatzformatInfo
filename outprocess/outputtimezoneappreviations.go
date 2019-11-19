package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/fileops"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
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

	return nil
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
	b.WriteString("  errors\n")
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
	b.WriteString("func (TzAbbrv *TimeZoneAbbreviationDto) CopyIn(inComing *TzAbbreviationDto) error {\n")
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
	b.WriteString("}  \n")

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