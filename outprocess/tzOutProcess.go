package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
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

	err = nil

	errArray := make([]error, 0)

	err2 := f.FlushBytesToDisk()

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


	currDateTimeStr := tzdatastructs.ApplicationStartDateTime.Format(tzdatastructs.FmtDateTime)

	regionalStats, err := tzOut.createIanaRegionalTimeZoneStats(tzStats, ePrefix)

	if err != nil {
		return make([]byte, 0), err
	}

	outputStr := fmt.Sprintf("\n" +
		"// TimeZones - This type and its associated methods encapsulate %v IANA Time\n" +
		"// Zones, %v-Military Time Zones and %v-Other Time Zones. This type is\n" +
		"// therefore used as a comprehensive enumeration of Global Time Zones.\n" +
		"//\n" +
		"// The Go Programming Language uses IANA Time Zones in date-time calculations.\n" +
		"//  Reference:\n" +
		"//    https://golang.org/pkg/time/\n" +
		"//    https://golang.org/pkg/time/#LoadLocation\n" +
		"//\n" +
		"// The IANA Time Zone database is widely recognized as the the world's leading\n" +
		"// authority on global time zones.\n" +
		"//\n" +
		"// The 'TimeZones' type includes one artificial structure element labeled\n" +
		"// 'Deprecated'. This element encapsulates all of the IANA 'Link' Time Zones.\n" +
		"// 'Link' Time Zones detail those times zones which IANA has classified as\n" +
		"// obsolete and no longer in general use. Each one of these deprecated time\n" +
		"// zones maps to a current, valid IANA time zone.\n"+
		"//\n" +
		"// Reference:\n" +
		"//    https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\n" +
		"//    https://en.wikipedia.org/wiki/Tz_database\n" +
		"//\n" +
		"// The IANA Time Zone data base and reference information is located at:\n" +
		"//    https://www.iana.org/time-zones.\n" +
		"//\n" +
		"// For easy access to the all Time Zones it is recommended that you use the\n" +
		"// global variable 'TZones' declared below. This variable instantiates the\n" +
		"// 'TimeZones' type. It is therefore much easier to access any of the %v time\n" +
		"// zones using dot operators and intellisense (a.k.a. intelligent code completion).\n" +
		"//\n" +
		"// Examples:\n" +
		"// TZones.America.Argentina().Buenos_Aires() - America/Argentina/Buenos_Aires Time Zone\n" +
		"// TZones.America.Chicago()                  - America/Chicago USA Central Time Zone\n" +
		"// TZones.America.New_York()                 - America/New_York USA Eastern Time Zone\n" +
		"// TZones.America.Denver()                   - America/Denver USA Mountain Time Zone\n" +
		"// TZones.America.Los_Angeles()              - America/Los_Angeles USA Pacific Time Zone\n" +
		"// TZones.Europe.London()                    - Europe/London Time Zone\n" +
		"// TZones.Europe.Paris()                     - Europe/Paris  Time Zone\n" +
		"// TZones.Asia.Shanghai()                    - Asia/Shanghai Time Zone\n" +
		"//\n" +
		"// 'TimeZones' has been adapted to function as an enumeration of valid time zone\n" +
		"// values. Since Go does not directly support enumerations, the 'TimeZones' type\n" +
		"// has been configured to function in a manner similar to classic enumerations found\n" +
		"// in other languages like C#. For additional information, reference:\n" +
		"//      Jeffrey Richter Using Reflection to implement enumerated types\n" +
		"//             https://www.youtube.com/watch?v=DyXJy_0v0_U \n" +
		"//\n" +
		"// ----------------------------------------------------------------------------\n" +
		"//                           IANA Time Zones by Region                         \n" +
		"//\n" +
		regionalStats +
		"//\n" +
		"// ----------------------------------------------------------------------------\n" +
		"// \n" +
		"// This TimeZones Type is based on IANA Time Zone Database Version: %v\n" +
		"// \n" +
		"//           IANA Standard Time Zones : %3d\n" +
		"//           IANA Link Time Zones     : %3d\n" +
		"//                                         -------\n" +
		"//                 Sub-Total IANA Time Zones: %3d\n" +
		"// \n" +
		"//                Military Time Zones : %3d\n" +
		"//                   Other Time Zones : %3d\n" +
		"//                                         -------\n" +
		"//                          Total Time Zones: %3d\n" +
		"// \n" +
		"//       Standard Time Zone Sub-Groups: %3d\n" +
		"//           Link Time Zone Sub-Groups: %3d\n" +
		"//                                         -------\n" +
		"//                Total Time Zone Sub-Groups: %3d\n" +
		"// \n" +
		"//                  Primary Time Zone Groups: %3d\n" +
		"// \n" +
		"// Type Creation Date: %v\n" +
		"// ----------------------------------------------------------------------------\n" +
		"// \n",
		tzStats.TotalIanaStdTzLinkZones,
		tzStats.NumMilitaryTZones,
		tzStats.NumOtherTZones,
		tzStats.TotalZones,
		tzStats.IanaVersion,
		tzStats.NumIanaStdTZones,
		tzStats.NumIanaLinkTZones,
		tzStats.TotalIanaStdTzLinkZones,
		tzStats.NumMilitaryTZones,
		tzStats.NumOtherTZones,
		tzStats.TotalZones,
		tzStats.NumLevel2StdSubTZoneGroups,
		tzStats.NumLevel2LinkSubGroups,
		tzStats.TotalSubTZoneGroups,
		tzStats.NumMajorTZoneGroups,
		currDateTimeStr)

	return []byte(outputStr), nil
}

// createIanaRegionalTimeZoneStats - Configures Iana Time Zone Regional
// Statistics as string returned by the method.
func (tzOut TzOutProcess) createIanaRegionalTimeZoneStats(
	tzStats *tzdatastructs.TimeZoneStatsDto,
	ePrefix string) (string, error) {

	ePrefix += "TzOutProcess.createIanaRegionalTimeZoneStats() "
	strOps := strops.StrOps{}

	outputStats := ""

	temp, err := strOps.MakeSingleCharString(' ', 35)

	if err != nil {
		return outputStats,
		fmt.Errorf(ePrefix +
			"Error returned by err := strOps.MakeSingleCharString(' ', 35)\n" +
			"Error='%v'\n", err.Error())
	}



	outputStats += "// " + temp + "     " + "Time" + "     " + "Link" + "    " + "Total" + "\n"
	outputStats += "// " + temp + "    " + "Zones" + "    " + "Zones" + "    " + "Zones" + "\n"

	temp, err = strOps.MakeSingleCharString('-', 62)

	if err != nil {
		return outputStats,
			fmt.Errorf(ePrefix +
				"Error returned by err := strOps.MakeSingleCharString('-', 62)\n" +
				"Error='%v'\n", err.Error())
	}

	outputStats += "// " + temp + "\n"
	outputStats += "// \n"

	for i:=0; i < len(tzStats.IanaTzRegions); i++ {

		region, err := strOps.StrLeftJustify(tzStats.IanaTzRegions[i], 35)

		if err != nil {
			return outputStats,
				fmt.Errorf(ePrefix +
					"\nError returned by strOps.StrLeftJustify(tzStats.IanaTzRegions[i], 35)\n" +
					"Error='%v'\n", err.Error())
		}

		tzCount := "     " +
			fmt.Sprintf("%4d", tzStats.IanaTzCounters[i])

		linkCount := "     " +
			fmt.Sprintf("%4d", tzStats.IanaLinkCounters[i])

		linkTzCount := "     " +
			fmt.Sprintf("%4d", tzStats.IanaTotalTimeZoneLinkCounters[i])

		outputStats += "// " +
			region +
			tzCount +
			linkCount +
			linkTzCount + "\n"

	}

	totalLine, err := strOps.MakeSingleCharString('=', 62)

	if err != nil {
		return "",
		fmt.Errorf(ePrefix +
			"\nError returned by strOps.MakeSingleCharString('=', 62)\n" +
			"Error='%v'\n", err.Error())
	}

	outputStats += "// " + totalLine + "\n"

	totalValues := ""

	totalValues, err = strOps.StrRightJustify("Total ", 35)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix +
				"\nError returned by strOps.StrRightJustify(\"Total \", 35)\n" +
				"Error='%v'\n", err.Error())
	}

	totalValues += "     " +
		fmt.Sprintf("%4d", tzStats.IanaTotalTimeZones)

	totalValues += "     " +
		fmt.Sprintf("%4d", tzStats.IanaTotalLinks)

	totalValues += "     " +
		fmt.Sprintf("%4d", tzStats.IanaTotalTimeZonesLinks)

	outputStats += "// " + totalValues + "\n"

	return outputStats, nil
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