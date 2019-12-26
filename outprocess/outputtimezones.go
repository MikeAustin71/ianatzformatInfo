package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/textlinebuilder"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
	"time"
)

type OutputTimeZones struct {
	input string
}

// WriteOutput - Writes all formatted time zone data to the output file.
func (tzOut OutputTimeZones) WriteOutput(
	outputPathDirMgr pathfileops.DirMgr,
	fileNameExt string,
	tzStats *tzdatastructs.TimeZoneStatsDto, // Time Zone Version
	ePrefix string) error {

	ePrefix += "OutputTimeZones.WriteOutput() "

	f, err := tzOut.createOpenOutputFile(outputPathDirMgr, fileNameExt, ePrefix)

	if err != nil {
		return err
	}

	err = tzOut.writeHeadersToOutputFile(f, ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	err = tzOut.writeTimeZoneMasterType(
		f,
		tzStats,
		ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	err = tzOut.writeTimeZoneGlobalType(f,ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	err = tzOut.writeTimeZones(
		f,
		tzStats,
		ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	err = tzOut.writeTimeZoneMap(
		f,
		tzStats,
		ePrefix)

	if err != nil {
		_ = f.CloseThisFile()
		return err
	}

	leftMargin := textlinebuilder.MarginSpec{
		MarginStr:    "// ",
		MarginLength: 0,
		MarginChar:   0,
	}

	lineBreakStr := textlinebuilder.LineSpec{
		LineChar:         '-',
		LineLength:       70,
		LineFieldLength:  70,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.LeftJustify(),
	}

	err = TzStrFmt{}.WriteAlphabetizedTimeZoneList(f,leftMargin, lineBreakStr, tzStats, ePrefix)

	if err != nil {
		return err
	}

	err = nil

	errArray := make([]error, 0)

	var err2 error

	err2 = f.FlushBytesToDisk()

	if err2 != nil {

		errArray = append(errArray, fmt.Errorf(ePrefix +
			"\nError returned by f.FlushBytesToDisk()\n" +
			"Error='%v'\n", err2.Error()))
	}

	err2 = f.CloseThisFile()

	if err2 != nil {
		errArray = append(errArray, fmt.Errorf(ePrefix +
			"\nError returned by f.CloseThisFile()\n" +
			"Error='%v'\n", err2.Error()))
	}

	if len(errArray) > 0 {
		err = pathfileops.FileHelper{}.ConsolidateErrors(errArray)
		return err
	}

	return nil
}

// Creates and Opens the Go Source Code output file,
// 'timezonedata.go'.
func (tzOut OutputTimeZones) createOpenOutputFile(
	outputPathDirMgr pathfileops.DirMgr,
	fileNameExt, ePrefix string) (f pathfileops.FileMgr, err error) {

	ePrefix += "OutputTimeZones.CreateOutputFile() "

	f = pathfileops.FileMgr{}
	err = nil
	var err2 error


	f, err2 = pathfileops.FileMgr{}.NewFromDirMgrFileNameExt(outputPathDirMgr, fileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	err = f.IsFileMgrValid(ePrefix)

	if err != nil {
		return f, err
	}

	fileExists, err2 := f.DoesThisFileExist()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	if fileExists {

		err2 = f.DeleteThisFile()

		if err2 != nil {
			err = fmt.Errorf(ePrefix+"%v", err2.Error())
			return f, err
		}

		fileExists, err2 = f.DoesThisFileExist()

		if err2 != nil {
			err = fmt.Errorf(ePrefix+"%v", err2.Error())
			return f, err
		}

		if fileExists {
			err = fmt.Errorf(ePrefix+"Existing Output File FAILED to Delete! "+
				"Output File= '%v' ", f.GetAbsolutePathFileName())
			return f, err
		}

	}

	err2 = f.CreateThisFile()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	err2 = f.OpenThisFileReadWrite()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	err = nil

	return f, err
}

// createTimeZoneTypeComments - Creates comments for the master Time Zone Type.
//
func (tzOut OutputTimeZones) createTimeZoneTypeComments(
	b *strings.Builder,
	tzStats *tzdatastructs.TimeZoneStatsDto, ePrefix string) error {

	ePrefix += "OutputTimeZones.createTimeZoneTypeComments() "


	currDateTimeStr := tzdatastructs.ApplicationStartDateTime.Format(tzdatastructs.FmtDateTimeTzYMD)

	regionalStats, err := tzOut.createIanaRegionalTimeZoneStats(tzStats, ePrefix)

	if err != nil {
		return err
	}

	b.WriteString(fmt.Sprintf("\n" +
		"// TimeZones - This type and its associated methods encapsulate %v IANA Time\n",
		tzStats.TotalIanaStdTzLinkZones))

	b.WriteString(fmt.Sprintf(
		"// Zones, %v-Military Time Zones and %v-Other Non-Iana Time Zone. This 'TimeZones'\n",
		tzStats.NumMilitaryTZones,tzStats.NumOtherTZones))

	b.WriteString(
		"// type can therefore be used as a comprehensive enumeration of Global Time Zones.\n")

	b.WriteString("//\n")

	b.WriteString("// The Time Zones Type encapsulates data elements used to access specific\n")
	b.WriteString("// time zones. These data elements classify time zones by geographic region\n")
	b.WriteString("// and type. For example, classifications like, 'America', 'Asia' and 'Europe'\n")
	b.WriteString("// provide access to zones located in those geographic regions. The 'Military'\n")
	b.WriteString("// classification provides access to time zones used exclusively by military,\n")
	b.WriteString("// aviation or maritime organizations.\n")

	b.WriteString("// \n")

	b.WriteString(
		"// The classification 'Other' includes many Time Zones which have been deprecated\n")
	b.WriteString(
		"// by the IANA Time Zone database. Deprecated IANA Time Zones are mapped internally\n")
	b.WriteString(
		"// to valid, current time zones. \n")

	b.WriteString("// \n")

	b.WriteString(
		"// The Go Programming Language uses IANA Time Zones in date-time calculations.\n")

	b.WriteString(
		"// Reference:\n")

	b.WriteString(
		"//    https://golang.org/pkg/time/\n")

	b.WriteString(
		"//    https://golang.org/pkg/time/#LoadLocation\n")

	b.WriteString("// \n")

	b.WriteString(
		"// The IANA Time Zone database is widely recognized as a leading authority on global\n")

	b.WriteString(
		"// time zones.\n")

	b.WriteString("// \n")

	b.WriteString(
		"// Reference: \n")

	b.WriteString(
		"//    https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\n")

	b.WriteString(
		"//    https://en.wikipedia.org/wiki/Tz_database\n")

	b.WriteString(
		"//    https://en.wikipedia.org/wiki/List_of_military_time_zones\n")

	b.WriteString("// \n")

			b.WriteString(
		"// The 'TimeZones' type was generated by the application 'ianatzformatInfo.exe'\n")
	b.WriteString(
		"// This application extracts information from a 'zoneinfo.zip' file. The \n")

	b.WriteString(
		"// 'zoneinfo.zip' file was in turn generated from time zone database file supplied\n")

	b.WriteString(
			"// by the IANA time zone database.\n")

	b.WriteString("// \n")

	b.WriteString(
		"// For information on the IANA Time Zone Database, reference:\n")

	b.WriteString("// \n")

	b.WriteString(
		"//    https://www.iana.org/time-zones\n")

	b.WriteString(
		"//    https://data.iana.org/time-zones/releases/\n")

	b.WriteString("// \n")


	b.WriteString(
		"// For information on the application, 'ianatzformatInfo.exe', reference:\n")

	b.WriteString("// \n")

	b.WriteString(
		"//    https://github.com/MikeAustin71/ianatzformatInfo\n")

	b.WriteString("// \n")

	b.WriteString(
		"// For additional information on locating or creating the 'zoneinfo.zip' file,\n")

	b.WriteString(
		"// reference:\n")

	b.WriteString("// \n")

	b.WriteString(
		"//    https://github.com/MikeAustin71/ianatzformatInfo/blob/master/xtechnotes/TimeZoneDatabaseUpdates.md\n")

	b.WriteString("// \n")

	b.WriteString(
		"// For easy access to all time zones, use the global variable, 'TZones', declared below.\n")

	b.WriteString(
		"// This variable instantiates the 'TimeZones' type. 'TZones' allows for much easier access\n")

	b.WriteString( fmt.Sprintf(
		"// to any of the %v time zones using dot operators and intellisense (a.k.a. intelligent code completion).\n",
		tzStats.TotalZones))

	b.WriteString("// \n")

	b.WriteString(
	"// Examples:\n")

	b.WriteString(
		"// TZones.America.Argentina().Buenos_Aires() - America/Argentina/Buenos_Aires Time Zone\n")

	b.WriteString(
		"// TZones.America.Chicago()                  - America/Chicago USA Central Time Zone\n")

	b.WriteString(
		"// TZones.America.New_York()                 - America/New_York USA Eastern Time Zone\n")

	b.WriteString(
		"// TZones.America.Denver()                   - America/Denver USA Mountain Time Zone\n")

	b.WriteString(
		"// TZones.America.Los_Angeles()              - America/Los_Angeles USA Pacific Time Zone\n")

	b.WriteString(
		"// TZones.Europe.London()                    - Europe/London Time Zone\n")

	b.WriteString(
		"// TZones.Europe.Paris()                     - Europe/Paris  Time Zone\n")

	b.WriteString(
		"// TZones.Asia.Shanghai()                    - Asia/Shanghai Time Zone\n")

	b.WriteString(
		"// TZones.Local()                            - Time Zone used on host computer\n")

	b.WriteString(
		"// TZones.UTC()                              - Coordinated Universal Time, UTC+0000\n")

	b.WriteString(
		"// TZones.Zulu()                             - Military Time Zone, UTC+0000.\n")

	b.WriteString("//\n")

	b.WriteString(
		"// 'TimeZones' has been adapted to function as an enumeration of valid time zone\n")

	b.WriteString(
		"// values. Since Go does not directly support enumerations, the 'TimeZones' type\n")

	b.WriteString(
		"// has been configured to function in a manner similar to classic enumerations found\n")

	b.WriteString(
		"// in other languages like C#. For additional information, reference:\n")

	b.WriteString("// \n")

	b.WriteString(
		"//    Jeffrey Richter Using Reflection to implement enumerated types\n")

	b.WriteString(
		"//    https://www.youtube.com/watch?v=DyXJy_0v0_U \n")

	b.WriteString("//\n")
	b.WriteString("//\n")

	b.WriteString(fmt.Sprintf(
		"// A complete alphabetic listing of all %v time zones is provided at the end\n",
		tzStats.TotalZones))

	b.WriteString("// of this source file.\n")

	b.WriteString(
		"// ----------------------------------------------------------------------------\n")

	b.WriteString(
		"//                           IANA Time Zones by Region                         \n")

	b.WriteString("//\n")

	b.WriteString(regionalStats)

	b.WriteString("//\n")

	b.WriteString("//\n")

	b.WriteString("// Note that the 'Other' time zone classification includes deprecated \n")

	b.WriteString("// or obsolete time zones as well as time zone abbreviations.  All\n")

	b.WriteString("// 'deprecated' time zones map to valid current time zones.\n")

	b.WriteString(
		"// ----------------------------------------------------------------------------\n")

	b.WriteString("// \n")

	b.WriteString(fmt.Sprintf(
		"// This 'TimeZones' Type is based on IANA Time Zone Database Version: %v\n",
		tzStats.IanaVersion))

	b.WriteString("// \n")

	b.WriteString(fmt.Sprintf(
		"//           IANA Standard Time Zones : %3d\n",
		tzStats.NumIanaStdTZones))

	b.WriteString(fmt.Sprintf(
		"//                Military Time Zones : %3d\n",
			tzStats.NumMilitaryTZones))

	b.WriteString(fmt.Sprintf(
		"//          Other Non-Iana Time Zones : %3d  ('Local' Time Zone)\n",
		tzStats.NumOtherTZones))

	b.WriteString(
		"//                                         -------\n")

	b.WriteString(fmt.Sprintf(
		"//                          Total Time Zones: %3d\n",
		tzStats.TotalZones))

	b.WriteString("// \n")

	b.WriteString(fmt.Sprintf(
		"//       Standard Time Zone Sub-Groups: %3d\n",
		tzStats.NumLevel2StdSubTZoneGroups))

	b.WriteString("// \n")

	b.WriteString(fmt.Sprintf(
		"//            Primary Time Zone Groups: %3d\n",
		tzStats.NumMajorTZoneGroups))

	b.WriteString("// \n")

	b.WriteString(fmt.Sprintf(
		"// Type Creation Date: %v\n",
			currDateTimeStr))

	b.WriteString(
		"// ----------------------------------------------------------------------------\n")

	b.WriteString("// \n")

	return nil
}


// createIanaRegionalTimeZoneStats - Configures Iana Time Zone Regional
// Statistics as string returned by the method.
func (tzOut OutputTimeZones) createIanaRegionalTimeZoneStats(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) (string, error) {

	ePrefix += "OutputTimeZones.createIanaRegionalTimeZoneStats() "

	b := strings.Builder{}

	b.Grow(2048)

	leftMargin := textlinebuilder.MarginSpec{
		MarginStr:    "// ",
		MarginLength: 0,
		MarginChar:   0,
	}

	leftSpacer := textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: 35,
		MarginChar:   ' ',
	}

	leftPad := textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: 5,
		MarginChar:   ' ',
	}

	strSpec1 := textlinebuilder.StringSpec{
		StrValue:       "Time",
		StrFieldLength: 4,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	newLine := textlinebuilder.NewLineSpec{AddNewLine:true}

	err := textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		leftSpacer,
		leftPad,
		strSpec1,
		newLine)

	if err != nil {
		return "", err
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "Zones",
		StrFieldLength: 5,
		StrPadChar:     ' ',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	leftPad.MarginLength = 4

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		leftSpacer,
		leftPad,
		strSpec1,
		newLine)

	if err != nil {
		return "", err
	}

lineBreak := textlinebuilder.LineSpec{
	LineChar:         '-',
	LineLength:       62,
	LineFieldLength:  62,
	LineFieldPadChar: ' ',
	LinePosition:     textlinebuilder.FieldPos.LeftJustify(),
}

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		lineBreak,
		newLine)

	if err != nil {
		return "", err
	}

	intSpec2 := textlinebuilder.IntegerSpec{
		NumericValue:       99,
		NumericFieldSpec:   "%4d",
		NumericFieldLength: 4,
		NumericPadChar:     ' ',
		NumericPosition:    textlinebuilder.FieldPos.RightJustify(),
	}

	strSpec1 = textlinebuilder.StringSpec{
		StrValue:       "xx",
		StrFieldLength: 35,
		StrPadChar:     '.',
		StrPosition:    textlinebuilder.FieldPos.LeftJustify(),
	}

	leftSpacer = textlinebuilder.MarginSpec{
		MarginStr:    "",
		MarginLength: 5,
		MarginChar:   '.',
	}

	for i:=0; i < len(tzStats.IanaTzRegions); i++ {

		strSpec1.StrValue =  tzStats.IanaTzRegions[i]

		intSpec2.NumericValue = tzStats.IanaTzCounters[i]

		err = textlinebuilder.TextLineBuilder{}.Build(
			&b,
			ePrefix,
			leftMargin,
			strSpec1,
			leftSpacer,
			intSpec2,
			newLine)

		if err != nil {
			return "", err
		}

	}

	lineBreak.LineChar = '='

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		lineBreak,
		newLine)

	if err != nil {
		return "", err
	}

	strSpec1.StrValue = "Total "
	strSpec1.StrFieldLength = 35
	strSpec1.StrPosition = textlinebuilder.FieldPos.RightJustify()

	leftSpacer.MarginLength = 5
	leftSpacer.MarginChar  = ' '

	intSpec2.NumericValue = tzStats.IanaTotalTimeZones

	err = textlinebuilder.TextLineBuilder{}.Build(
		&b,
		ePrefix,
		leftMargin,
		strSpec1,
		leftSpacer,
		intSpec2,
		newLine,
		leftMargin,
		newLine)

	if err != nil {
		return "", err
	}


	return b.String(), nil
}

// writeHeadersToOutputFile - Writes header information to the
// output file. This includes the 'package' statement.
//
// Input parameter 'outputFileMgr' MUST be open and ready for
// Write operations.
//
func (tzOut OutputTimeZones) writeHeadersToOutputFile(
	outputFileMgr pathfileops.FileMgr, ePrefix string) (err error) {

		err = nil

	ePrefix += "OutputTimeZones.writeHeadersToOutputFile() "

	if !outputFileMgr.IsInitialized() {
		err = fmt.Errorf(ePrefix +
			"Input parameter 'outputFileMgr' IS NOT INITIALIZED!")
	}

	if !outputFileMgr.IsFilePointerOpen() {
		err = fmt.Errorf(ePrefix +
			"'outputFileMagr IS NOT OPEN!")
	}

	var errorArray []error
	b := strings.Builder{}

	b.Grow(1024)

	b.WriteString("package main\n\n\n")

	b.WriteString("import (\n")
	b.WriteString("      \"fmt\"\n")
	b.WriteString("      \"strings\"\n")
	b.WriteString("      \"sync\"\n")
	b.WriteString("      \"time\"\n")
	b.WriteString(")\n\n\n")

	_, err2 := outputFileMgr.WriteBytesToFile ([]byte(b.String()))

	if err2 != nil {

		errorArray = append(errorArray, fmt.Errorf(ePrefix+"Line1: %v", err2.Error()))

		err2 = outputFileMgr.CloseThisFile()

		if err2 != nil {
			errorArray = append(errorArray, err2)
			err = pathfileops.FileHelper{}.ConsolidateErrors(errorArray)
		}

		return err
	}


	err = nil
	return err
}

// writeTimeZones - Writes all time zones and link zones to
// the output file timezonedata.go.
func (tzOut OutputTimeZones) writeTimeZones(
	outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "OutputTimeZones.writeLevelOneTimeZones() "

	var grp *tzdatastructs.TimeZoneGroupDto
	var tzCol tzdatastructs.TimeZoneDataCollection
	var tZone *tzdatastructs.TimeZoneDataDto
	var err error

	for i:=0; i <= tzdatastructs.Level_03_Idx; i++ {

		tzStats.TzGroups[i].Sort(false)

		tzStats.TzData[i].SortByGroupTzName(false)

		lenGrpAry := tzStats.TzGroups[i].GetNumberOfGroups()

		for j:= 0; j < lenGrpAry; j++ {

			grp, err = tzStats.TzGroups[i].PeekPtr(j)

			if err != nil {
				return fmt.Errorf(ePrefix +
					"\nError returned by tzGroupsAry[i].Peek(j).\n" +
					"i='%v' j='%v'\n" +
					"Error='%v'\n", i, j, err.Error())
			}

			if len(grp.TypeDeclaration) == 0 {
				return fmt.Errorf(ePrefix +
					"\nError: Group Type Declaration has Zero Bytes!\n" +
					"i='%v' j='%v'\n" +
					"Parent Group Name='%v'\n" +
					"Group Name='%v'\n", i, j, grp.ParentGroupName, grp.GroupName)
			}

			_, err = outputFileMgr.WriteBytesToFile(grp.TypeDeclaration)

			if err != nil {
				return fmt.Errorf(ePrefix +
					"\nError returned by outputFileMgr.WriteBytesToFile()\n" +
					"i='%v' j='%v'\n" +
					"Parent Group Name='%v'\n" +
					"Group Name='%v'\n" +
					"Error='%v'\n",
					i, j, grp.ParentGroupName, grp.GroupName, err.Error())
			}

			tzCol, err = tzStats.TzData[i].GetZoneGroupCol(grp)

			if err != nil {
				return fmt.Errorf(ePrefix)
			}

			tzCol.SortByGroupTzName(false)

			numOfTimeZones := tzCol.GetNumberOfTimeZones()

			if numOfTimeZones == 0 {
				return fmt.Errorf(ePrefix +
					"\nTime Zone Collection is EMPTY!\n" +
					"Parent Group='%v'\n" +
					"Group Name='%v'\n",
					grp.ParentGroupName, grp.GroupName)
			}

			for k:=0; k < numOfTimeZones; k++ {

				tZone, err = tzCol.PeekPtr(k)

				if err != nil {
					return fmt.Errorf(ePrefix +
						"\nError returned by tzCol.Peek(k)\n" +
						"k='%v'\n" +
						"Parent Group='%v'\n" +
						"Group Name='%v'\n",
						k, grp.ParentGroupName, grp.GroupName)
				}

				if len(tZone.FuncDeclaration) == 0 {
					return fmt.Errorf(ePrefix +
						"\nTime Zone Func Declaration has Zero Bytes!\n" +
						"Time Zone Name='%v'\n" +
						"Parent Group='%v'\n" +
						"Group Name='%v'\n",
						tZone.TzName, grp.ParentGroupName, grp.GroupName)
				}

				_, err = outputFileMgr.WriteBytesToFile(tZone.FuncDeclaration)

				if err != nil {
					return fmt.Errorf(ePrefix +
						"\nError returned by outputFileMgr.WriteBytesToFile(tZone.FuncDeclaration)\n" +
						"tZone.FuncDeclaration: %v\n" +
						"Error:'%v'\n", string(tZone.FuncDeclaration), err.Error())
				}

			}
		}
	}

	err = outputFileMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

		return nil
}

func (tzOut OutputTimeZones) writeTimeZoneGlobalType(
	outputFileMgr pathfileops.FileMgr,
	ePrefix string) error {

	ePrefix += "OutputTimeZones.writeTimeZoneGlobalType() "

	var err error

	outBytes := []byte("var TZones = TimeZones{}\n\n\n")

	_, err = outputFileMgr.WriteBytesToFile(outBytes)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile(outBytes)\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

// Writes a map of all Time Zones to include: 'Local' time zone,
// 'IANA' Time Zones and Military Time Zones.
//
func (tzOut OutputTimeZones) writeTimeZoneMap(
	outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "OutputTimeZones.writeTimeZoneMap() "

	numOfTimeZones := tzStats.IanaCapturedTimeZones.GetNumberOfTimeZones()

	if numOfTimeZones < 1 {

		return fmt.Errorf(ePrefix +
			"\nError: Number of tzStats.IanaCapturedTimeZones='%v'\n", numOfTimeZones)
	}

	b := strings.Builder{}

	b.Grow(5120)

	b.WriteString("\n\n")

	b.WriteString("// TimeZoneUtcOffsetReference - Provides thread safe access to\n")
	b.WriteString("// all Time Zone zone names and their associated UTC offsets.\n")
	b.WriteString("//  \n")
	b.WriteString("type TimeZoneUtcOffsetReference struct {\n")
	b.WriteString("	Input                  string\n")
	b.WriteString("	Output                 string\n")
	b.WriteString("}\n\n")

	b.WriteString("// GetTimeZoneUtcOffset - Returns a UTC offset expressed in accordance \n")
	b.WriteString("// with the following examples: 'UTC+0500', UTC-0500'\n")
	b.WriteString("// \n")
	b.WriteString("// The lookup operation is performed based on input string parameter\n")
	b.WriteString("// 'timeZoneName'.\n")
	b.WriteString("//  \n")
	b.WriteString("func(tzUtcOffset TimeZoneUtcOffsetReference) GetTimeZoneUtcOffset(\n")
	b.WriteString("  timeZoneName string) (string, error) {\n" )
	b.WriteString("\n")

	b.WriteString("  lockMapAllTimeZonesToUtcOffsets.Lock()\n\n")

	b.WriteString("  defer lockMapAllTimeZonesToUtcOffsets.Unlock()\n\n")

	b.WriteString("  ePrefix := \"TimeZoneUtcOffsetReference.GetTimeZoneUtcOffset() \" \n\n")

	b.WriteString("  testStr := strings.ToLower(timeZoneName)\n\n")

	b.WriteString("  if testStr == \"local\"{\n\n")

	b.WriteString("    t := time.Now().In(time.Local)\n\n")

	b.WriteString("    tStr := t.Format(\"2006-01-02 15:04:05 -0700 MST\")\n\n")

	b.WriteString("    lenLeadStr := len(\"2006-01-02 15:04:05 \")\n\n")

	b.WriteString("    return \"UTC\" + tStr[lenLeadStr: lenLeadStr + 5], nil \n")
	b.WriteString("  }\n\n")

	b.WriteString("  utcOffset, ok := mapAllTimeZonesToUtcOffsets[timeZoneName]\n\n")

	b.WriteString("  if !ok {\n")
	b.WriteString("    return \"\", fmt.Errorf(ePrefix + \n")
	b.WriteString("      \"\\nInvalid 'timeZoneName'!\\n\" + \n")
	b.WriteString("      \"timeZoneName='%v'\\n\", timeZoneName)\n")
	b.WriteString("  }\n\n")

	b.WriteString("  return utcOffset, nil\n")
	b.WriteString("}\n\n")

	b.WriteString("// mapAllTimeZonesToUtcOffsets - A reference map including all\n")

	b.WriteString("// valid time zones and their associated UTC offsets.\n")

	b.WriteString("//\n\n")

	b.WriteString("var lockMapAllTimeZonesToUtcOffsets sync.Mutex\n\n")

	xSpacer := strings.Repeat(" ", 41)

	b.WriteString("var mapAllTimeZonesToUtcOffsets = map[string]string{\n")

	t := time.Now().In(time.Local)

	testTimeStr := t.Format("2006-01-02 15:04:05 -0700 MST")
	lenLeadStr := len("2006-01-02 15:04:05 ")
	localOffset := "UTC" + testTimeStr[lenLeadStr: lenLeadStr + 5]


	b.WriteString("  \"Local\" :" + xSpacer + "\"" + localOffset + "\",\n")

	tzStats.IanaCapturedTimeZones.SortByCanonicalValue()

	var tz * tzdatastructs.TimeZoneDataDto
	var err error

	for i :=0; i < numOfTimeZones; i++ {

		tz, err = tzStats.IanaCapturedTimeZones.PeekPtr(i)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by tzStats.IanaCapturedTimeZones.PeekPtr(i)\n" +
				"i='%v'\nError='%v'\n", i, err.Error())
		}


		xSpacer = strings.Repeat(" ", 46 - len(tz.TzCanonicalValue ))

		b.WriteString("  \"" + tz.TzCanonicalValue + "\" :" + xSpacer + "\"" + tz.UtcOffset + "\",\n" )

	}

	numOfMilitaryTimeZones := tzStats.CapturedMilitaryZones.GetNumberOfTimeZones()

	if numOfMilitaryTimeZones < 1 {
		return fmt.Errorf(ePrefix +
			"\nNumber of Military Time Zones is INVALID!\n" +
			"numOfMilitaryTimeZones='%v'\n", numOfMilitaryTimeZones)
	}

	for j:=0; j < numOfMilitaryTimeZones; j++ {
		tz, err = tzStats.CapturedMilitaryZones.PeekPtr(j)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by tzStats.CapturedMilitaryZones.PeekPtr(j)\n"+
				"j='%v'\n"+
				"Error='%v'\n", j, err.Error())
		}

		if tz.TzName == "Zulu" {
			continue
		}

		xSpacer = strings.Repeat(" ", 46 - len(tz.TzName))

		b.WriteString("  \"" + tz.TzName + "\" :" + xSpacer + "\"" + tz.UtcOffset + "\",\n" )
	}
	b.WriteString("}\n\n\n")

	_, err = outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile([]byte(b.String()))\n" +
			"Error='%v'\n", err.Error())
	}

	err = outputFileMgr.FlushBytesToDisk()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.FlushBytesToDisk()\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}

// Writes Master Type: type TimeZones struct
func (tzOut OutputTimeZones) writeTimeZoneMasterType(
	outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "OutputTimeZones.writeTimeZoneMasterType() "

	lenMasterGroups := tzStats.TzGroups[tzdatastructs.Level_01_Idx].GetNumberOfGroups()

	b := strings.Builder{}

	b.Grow(5120)

	var err error

	err  = tzOut.createTimeZoneTypeComments( &b, tzStats, ePrefix)

	if err != nil {
		return err
	}

	b.WriteString("type TimeZones struct {\n")

	var leftMarginStr string
	var centerMarginStr string
	const centerMarginLen = 35

	_, err = outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile(typeDeclaration)\n" +
			"Error='%v'\n", err.Error())
	}

	b.Reset()

	b.Grow(5120)

	leftMarginStr, err = strops.StrOps{}.MakeSingleCharString(' ', 4)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by StrOps{}.MakeSingleCharString(' ', 4)\n" +
			"Error='%v'\n", err.Error())
	}

	tzStats.TzGroups[tzdatastructs.Level_01_Idx].Sort(true)

	var group *tzdatastructs.TimeZoneGroupDto

	for i:=0; i < lenMasterGroups; i++ {

		group, err = tzStats.TzGroups[tzdatastructs.Level_01_Idx].PeekPtr(i)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by tzGroupsAry[tzdatastructs.Level_01_Idx].Peek(i)\n" +
				"i='%v'\n" +
				"Error='%v'\n", i, err.Error())
		}

		centerLen := centerMarginLen - len(group.GroupName)

		if centerLen < 1 {
			centerLen = 5
		}

		centerMarginStr, err = strops.StrOps{}.MakeSingleCharString(' ', centerLen)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by StrOps{}.MakeSingleCharString(' ', centerLen)\n" +
				"Error='%v'\n", err.Error())
		}

		b.WriteString(leftMarginStr + group.GroupName + centerMarginStr + group.TypeName + "\n")

	}

	b.WriteString ("}\n\n")

	_, err = outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by outputFileMgr.WriteBytesToFile(EndOfMasterType)\n" +
			"Error='%v'\n", err.Error())
	}

	return tzOut.writeMasterTypeFunctions(outputFileMgr, ePrefix)
}

// writeMasterTypeFunctions - Writes type TimeZone functions to the
// output file. These are top level functions directly accessible
// by type 'TimeZones'. Currently, one function is provided, 'Local'.
// 'Local' returns the time zone of the host computer running the code.
//
func (tzOut OutputTimeZones) writeMasterTypeFunctions(
	outputFileMgr pathfileops.FileMgr,
	ePrefix string) error {

	ePrefix += "OutputTimeZones.writeTimeZoneFunctions() "
	var err error

	b := strings.Builder{}

	b.Grow(5120)

	b.WriteString("// Local - Returns the local time zone on the host computer where this\n")
	b.WriteString("// code is executed.\n")
	b.WriteString("// \n")
	b.WriteString("// For documentation, reference:\n")
	b.WriteString("//    https://golang.org/pkg/time/#LoadLocation\n")
	b.WriteString("// \n")
	b.WriteString("func(tZones TimeZones) Local() string {return \"Local\"}\n")
	b.WriteString("\n\n")

	b.WriteString("// UCT - A time zone equivalent to Coordinated Universal Time. \n")
	b.WriteString("// Coordinated Universal Time is the primary time standard by which the world regulates\n")
	b.WriteString("// regulates clocks and time. It is within about 1 second of mean solar time at 0°\n")
	b.WriteString("// longitude, and is not adjusted for daylight saving time. In some countries, the\n")
	b.WriteString("// term Greenwich Mean Time is used.\n ")
	b.WriteString("// \n")
	b.WriteString("// UCT is equivalent to a zero offset: UTC+0000. For additional information, reference:\n")
	b.WriteString("//     https://en.wikipedia.org/wiki/Coordinated_Universal_Time\n" )
	b.WriteString("// \n")
	b.WriteString("func (tZones TimeZones) UCT()  string { return \"UCT\" }\n")
	b.WriteString("\n\n")


	b.WriteString("// UTC - Coordinated Universal Time. \n")
	b.WriteString("// Coordinated Universal Time (or UTC) is the primary time standard by which the\n")
	b.WriteString("// world regulates clocks and time. It is within about 1 second of mean solar time\n")
	b.WriteString("// at 0° longitude, and is not adjusted for daylight saving time. In some countries,\n")
	b.WriteString("// the term Greenwich Mean Time is used.\n ")
	b.WriteString("// \n")
	b.WriteString("// UTC is equivalent to a zero offset: UTC+0000. For additional information, reference:\n")
	b.WriteString("//     https://en.wikipedia.org/wiki/Coordinated_Universal_Time\n" )
	b.WriteString("// \n")
	b.WriteString("func (tZones TimeZones) UTC()  string { return \"UTC\" }\n")
	b.WriteString("\n\n")

	b.WriteString("// Zulu - Zulu Time Zone (Z) has no offset from Coordinated Universal Time (UTC).\n")
	b.WriteString("// This time zone is a military time zone. ")
	b.WriteString("// \n")
	b.WriteString("// Zulu Time Zone is often used in aviation and the military as another name for UTC +0.\n")
	b.WriteString("// Zulu Time Zone is also commonly used at sea between longitudes 7.5° West and 7.5° East.\n")
	b.WriteString("// The letter Z may be used as a suffix to denote a time being in the Zulu Time Zone,\n")
	b.WriteString("// such as 08:00Z or 0800Z. This is spoken as \"zero eight hundred Zulu\".\n")
	b.WriteString("// \n")
	b.WriteString("// The US Military, Chinese Military, and several others have adopted a unique list of\n")
	b.WriteString("// names for the time zones across the world. The names use the NATO phonetic alphabet which\n")
	b.WriteString("// calls for the 26 letters of the alphabet to be designated by easily recognizable words.\n")
	b.WriteString("// The Zulu time zone is one of these.\n")
	b.WriteString("// \n")
	b.WriteString("// For additional information, reference:\n")
	b.WriteString("//     https://www.timeanddate.com/time/zones/z\n")
	b.WriteString("//     https://www.timeanddate.com/time/zone/timezone/zulu\n")
	b.WriteString("//      \n")
	b.WriteString("func (tZones TimeZones) Zulu()  string { return \"UTC\" }\n")
	b.WriteString("\n\n\n")

		_, err = outputFileMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile(EndOfMasterType)\n" +
			"Error='%v'\n", err.Error())
	}

	return nil
}