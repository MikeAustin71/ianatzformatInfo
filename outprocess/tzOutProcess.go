package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/textlinebuilder"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
)

type TzOutProcess struct {
	input string
}

// WriteOutput - Writes all formatted time zone data to the output file.
func (tzOut TzOutProcess) WriteOutput(
	outputPathDirMgr pathfileops.DirMgr,
	fileNameExt string,
	tzStats *tzdatastructs.TimeZoneStatsDto, // Time Zone Version
	ePrefix string) error {

	ePrefix += "TzOutProcess.WriteOutput() "

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

	err2 := f.FlushBytesToDisk()

	if err2 != nil {
		_ = f.CloseThisFile()
		return fmt.Errorf(ePrefix +
			"\nError returned by f.FlushBytesToDisk()\n" +
			"f='%v'\nError='%v'\n",
			f.GetAbsolutePathFileName(),
			err2.Error())
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
func (tzOut TzOutProcess) createOpenOutputFile(
	outputPathDirMgr pathfileops.DirMgr,
	fileNameExt, ePrefix string) (f pathfileops.FileMgr, err error) {

	ePrefix += "TzOutProcess.CreateOutputFile() "

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
func (tzOut TzOutProcess) createTimeZoneTypeComments(
	tzStats *tzdatastructs.TimeZoneStatsDto, ePrefix string) ([]byte, error) {

	ePrefix += "TzOutProcess.createTimeZoneTypeComments() "


	currDateTimeStr := tzdatastructs.ApplicationStartDateTime.Format(tzdatastructs.FmtDateTimeTzYMD)

	regionalStats, err := tzOut.createIanaRegionalTimeZoneStats(tzStats, ePrefix)

	if err != nil {
		return make([]byte, 0), err
	}

	b := strings.Builder{}
	b.Grow(5120)

	b.WriteString(fmt.Sprintf("\n" +
		"// TimeZones - This type and its associated methods encapsulate %v IANA Time\n",
		tzStats.TotalIanaStdTzLinkZones))

	b.WriteString(fmt.Sprintf(
		"// Zones, %v-Military Time Zones and %v-Other Non-Iana Time Zones. This 'TimeZones'\n",
		tzStats.NumMilitaryTZones,tzStats.NumOtherTZones))

	b.WriteString(
		"// type can therefore be used as a comprehensive enumeration of Global Time Zones.\n")

	b.WriteString("//\n")

	b.WriteString(
		"// The Go Programming Language uses IANA Time Zones in date-time calculations.\n")

	b.WriteString(
		"//  Reference:\n")

	b.WriteString(
		"//    https://golang.org/pkg/time/\n")

	b.WriteString(
		"//    https://golang.org/pkg/time/#LoadLocation\n")

	b.WriteString("//\n")

	b.WriteString(
		"// The IANA Time Zone database is widely recognized as the the world's leading\n")

	b.WriteString(
		"// authority on global time zones.\n")

	b.WriteString("//\n")

	b.WriteString(
		"// Reference:\n")

	b.WriteString(
		"//    https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\n")

	b.WriteString(
			"//    https://en.wikipedia.org/wiki/Tz_database\n")

	b.WriteString("// \n")

			b.WriteString(
		"// The 'TimeZones' type was generated by the application 'ianatzformatInfo.exe'\n")
	b.WriteString(
		"// This application extracts information from a 'zoneinfo.zip' file. The \n")

	b.WriteString(
		"// 'zoneinfo.zip' file was in turn generated from time zone database\n")

	b.WriteString(
			"// files supplied by IANA time zone database.\n")

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
		"//          Other Non-Iana Time Zones : %3d\n",
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

	return []byte(b.String()), nil
}


// createIanaRegionalTimeZoneStats - Configures Iana Time Zone Regional
// Statistics as string returned by the method.
func (tzOut TzOutProcess) createIanaRegionalTimeZoneStats(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) (string, error) {

	ePrefix += "TzOutProcess.createIanaRegionalTimeZoneStats() "

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
func (tzOut TzOutProcess) writeHeadersToOutputFile(
	outputFileMgr pathfileops.FileMgr, ePrefix string) (err error) {

		err = nil

	ePrefix += "TzOutProcess.writeHeadersToOutputFile() "

	if !outputFileMgr.IsInitialized() {
		err = fmt.Errorf(ePrefix +
			"Input parameter 'outputFileMgr' IS NOT INITIALIZED!")
	}

	if !outputFileMgr.IsFilePointerOpen() {
		err = fmt.Errorf(ePrefix +
			"'outputFileMagr IS NOT OPEN!")
	}

	var errorArray []error

	_, err2 := outputFileMgr.WriteBytesToFile ([]byte("package main\n\n\n"))

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
func (tzOut TzOutProcess) writeTimeZones(
	outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzOutProcess.writeLevelOneTimeZones() "

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

		return nil
}

func (tzOut TzOutProcess) writeTimeZoneGlobalType(
	outputFileMgr pathfileops.FileMgr,
	ePrefix string) error {

	ePrefix += "TzOutProcess.writeTimeZoneGlobalType() "

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

// Writes Master Type: type TimeZones struct
func (tzOut TzOutProcess) writeTimeZoneMasterType(
	outputFileMgr pathfileops.FileMgr,
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) error {

	ePrefix += "TzOutProcess.writeTimeZoneMasterType() "

	lenMasterGroups := tzStats.TzGroups[tzdatastructs.Level_01_Idx].GetNumberOfGroups()

	var err error
	var outBytes []byte

	outBytes, err  = tzOut.createTimeZoneTypeComments(tzStats, ePrefix)

	if err != nil {
		return err
	}

	_, err = outputFileMgr.WriteBytesToFile(outBytes)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile(typeDeclaration)\n" +
			"Error='%v'\n", err.Error())
	}


	outBytes = []byte("type TimeZones struct {\n")

	var leftMarginStr string
	var centerMarginStr string
	const centerMarginLen = 35

	_, err = outputFileMgr.WriteBytesToFile(outBytes)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile(typeDeclaration)\n" +
			"Error='%v'\n", err.Error())
	}

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

		outBytes = []byte(leftMarginStr + group.GroupName + centerMarginStr + group.TypeName + "\n")

		_, err = outputFileMgr.WriteBytesToFile(outBytes)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\n Error returned by outputFileMgr.WriteBytesToFile(TzGroupBytes)\n" +
				"Error='%v'\n", err.Error())
		}

	}

	outBytes = []byte("}\n\n\n")

	_, err = outputFileMgr.WriteBytesToFile(outBytes)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\n Error returned by outputFileMgr.WriteBytesToFile(EndOfMasterType)\n" +
			"Error='%v'\n", err.Error())
	}


	return nil
}