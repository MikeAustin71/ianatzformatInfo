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
	tzGroupsAry [] tzdatastructs.TimeZoneGroupCollection, // Array of Time Zone Group Collections
	tzZonesAry [] tzdatastructs.TimeZoneDataCollection,  // Array of Time Zone Data Collections
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
		tzGroupsAry,
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
		tzGroupsAry,
		tzZonesAry,
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
			"\nError returned by f.CloseThisFile()"))
	}

	if len(errArray) > 0 {
		err = pathfileops.FileHelper{}.ConsolidateErrors(errArray)
	}

	return err
}

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

	_, err2 := outputFileMgr.WriteBytesToFile ([]byte("package main\n\n\n\n\n"))

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

func (tzOut TzOutProcess) writeTimeZones(
	outputFileMgr pathfileops.FileMgr,
	tzGroupsAry [] tzdatastructs.TimeZoneGroupCollection, // Array of Time Zone Group Collections
	tzZonesAry [] tzdatastructs.TimeZoneDataCollection,  // Array of Time Zone Data Collections)
	ePrefix string) error {

	ePrefix += "TzOutProcess.writeLevelOneTimeZones() "

	var grp tzdatastructs.TimeZoneGroupDto
	var tzCol tzdatastructs.TimeZoneDataCollection
	var tZone tzdatastructs.TimeZoneDataDto
	var err error

	for i:=0; i <= tzdatastructs.Level_03_Idx; i++ {

		tzGroupsAry[i].Sort(false)

		tzZonesAry[i].SortByGroupTzName(false)

		lenGrpAry := tzGroupsAry[i].GetNumberOfGroups()

		for j:= 0; j < lenGrpAry; j++ {

			grp, err = tzGroupsAry[i].Peek(j)

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

			tzCol, err = tzZonesAry[i].GetZoneGroupCol(grp)

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

				tZone, err = tzCol.Peek(k)

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

				if tZone.TzType == tzdatastructs.TZType.Standard() ||
						tZone.TzType == tzdatastructs.TZType.SubZone() {
					tzdatastructs.NumberOfTimeZones++
				}

			}
		}
	}

	fmt.Println("Number Of Time Zones Captured: ", tzdatastructs.NumberOfTimeZones)
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
	tzGroupsAry [] tzdatastructs.TimeZoneGroupCollection,
	ePrefix string) error {

	ePrefix += "TzOutProcess.writeTimeZoneMasterType() "

	lenMasterGrps := tzGroupsAry[tzdatastructs.Level_01_Idx].GetNumberOfGroups()

	var err error

	_, err = outputFileMgr.WriteBytesToFile(tzdatastructs.TimeZoneTypeComments)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by outputFileMgr.WriteBytesToFile(typeDeclaration)\n" +
			"Error='%v'\n", err.Error())
	}

	var outBytes []byte

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

	leftMarginStr, err = strops.StrOps{}.MakeSingleCharString(' ', 5)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by StrOps{}.MakeSingleCharString(' ', 5)\n" +
			"Error='%v'\n", err.Error())
	}


	var grp tzdatastructs.TimeZoneGroupDto

	for i:=0; i < lenMasterGrps; i++ {

		grp, err = tzGroupsAry[tzdatastructs.Level_01_Idx].Peek(i)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by tzGroupsAry[tzdatastructs.Level_01_Idx].Peek(i)\n" +
				"i='%v'\n" +
				"Error='%v'\n", i, err.Error())
		}

		centerLen := centerMarginLen - len(grp.GroupName)

		if centerLen < 1 {
			centerLen = 5
		}

		centerMarginStr, err = strops.StrOps{}.MakeSingleCharString(' ', centerLen)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by StrOps{}.MakeSingleCharString(' ', centerLen)\n" +
				"Error='%v'\n", err.Error())
		}

		outBytes = []byte(leftMarginStr + grp.GroupName + centerMarginStr + grp.TypeName + "\n")

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