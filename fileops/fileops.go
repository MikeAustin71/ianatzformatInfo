package fileops

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"os"
)

type FileOps struct {
	input      string
	output     string
}


// CreateOpenFile - Creates and opens a file for Read/Write operations.
//
func (fOps FileOps) CreateOpenFile(
	pathDirMgr pathfileops.DirMgr,
	fileNameExt string,
	ePrefix string) (f pathfileops.FileMgr, err error) {

	ePrefix += "TzOutProcess.CreateOpenFile() "

	f = pathfileops.FileMgr{}
	err = nil
	var err2 error

	f, err2 = pathfileops.FileMgr{}.NewFromDirMgrFileNameExt(pathDirMgr, fileNameExt)

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
		_ = f.CloseThisFile()
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	err = nil

	return f, err
}

// GetApplicationDirectory - Returns the current working directory in which
// this application resides.
func (fOps FileOps) GetApplicationDirectory(ePrefix string) (pathfileops.DirMgr, error) {

	ePrefix += "FileOps.GetApplicationCurrentWorkingDirectory() "

	exeDirMgr := pathfileops.DirMgr{}

	exePathFileNameExt, err :=  pathfileops.FileHelper{}.GetExecutablePathFileName()

	if err != nil {
		return exeDirMgr,
		fmt.Errorf(ePrefix +
			"\nError returned by pathfileops.FileHelper{}.GetExecutablePathFileName()\n" +
			"Error='%v'\n", err.Error())
	}

	var exeFMgr pathfileops.FileMgr

	exeFMgr, err = pathfileops.FileMgr{}.New(exePathFileNameExt)

	if err != nil {
		return exeDirMgr,
			fmt.Errorf(ePrefix +
				"\nError returned by FileMgr{}.New(exePathFileNameExt)\n" +
				"exePathFileNameExt='%v'\n" +
				"Error='%v'\n", exePathFileNameExt, err.Error())
	}

	var filePathDoesExist bool

	filePathDoesExist, err = exeFMgr.DoesThisFileExist()

	if err != nil {
		return exeDirMgr, fmt.Errorf(ePrefix +
			"\nNon-Path Error returned by exeFMgr.DoesThisFileExist()\n" +
			"exeFMgr='%v'\n" +
			"Error='%v'\n", exeFMgr.GetAbsolutePathFileName(), err.Error())
	}

	if !filePathDoesExist {
		return exeDirMgr, fmt.Errorf(ePrefix +
			"\nError: The executable path and file name do NOT exist!\n" +
			"exeFMgr='%v'\n", exeFMgr.GetAbsolutePathFileName())
	}

	exeDirMgr = exeFMgr.GetDirMgr()

	return exeDirMgr, nil
}

// GetCurrentWorkingDirectory - Returns the current working directory encapsulated in a
// 'DirMgr' type.
func (fOps FileOps) GetCurrentWorkingDirectory(ePrefix string) (pathfileops.DirMgr, error) {

	ePrefix += "FileOps.GetCurrentWorkingDirectory() "

	pathStr, err := pathfileops.FileHelper{}.GetCurrentDir()

	if err != nil {
		return pathfileops.DirMgr{},
			fmt.Errorf(ePrefix +
				"\nError returned by pathfileops.FileHelper{}.GetCurrentDir()\n" +
				"Error='%v'\n", err.Error())
	}

	var curDirMgr pathfileops.DirMgr

	curDirMgr, err = pathfileops.DirMgr{}.New(pathStr)

	if err != nil {
		return pathfileops.DirMgr{},
			fmt.Errorf(ePrefix +
				"\nError returned by pathfileops.DirMgr{}.New(pathStr)\n" +
				"Error='%v'\n", err.Error())
	}

	return curDirMgr, nil
}

// GetBaseDirectory - Conducts a search and returns a pathfileops.DirMgr which
//encapsulates the directory where ianatzformatInfo.exe is located.
//
func (fOps FileOps) GetBaseDirectory(ePrefix string) (pathfileops.DirMgr, error) {

	ePrefix += "FileOps.GetBaseDirectory() "

	var exeDirMgr, currDirMgr pathfileops.DirMgr
	var err error
	var pathFileNameExt string
	var fMgr pathfileops.FileMgr

	exeDirMgr, err = fOps.GetApplicationDirectory(ePrefix)

	if err != nil {
		goto checkCurrWorkingDir
	}

	pathFileNameExt = exeDirMgr.GetAbsolutePathWithSeparator() + "input" +
		string(os.PathSeparator) + tzdatastructs.AppInputParametersFileName


	fMgr, err = pathfileops.FileMgr{}.New(pathFileNameExt)

	if err != nil {
		return pathfileops.DirMgr{},
		fmt.Errorf(ePrefix +
			"\nError returned by exeDirMgr pathfileops.FileMgr{}.New(pathFileNameExt)\n" +
			"pathFileNameExt='%v'\nError='%v'\n", err.Error())
	}

	if fMgr.DoesFileExist() {
		return exeDirMgr, nil
	}

	checkCurrWorkingDir:

	currDirMgr, err = fOps.GetCurrentWorkingDirectory(ePrefix)

	if err != nil {
		return pathfileops.DirMgr{},
		fmt.Errorf(ePrefix +
			"\nError returned by fOps.GetCurrentWorkingDirectory(ePrefix)\n" +
			"Error='%v'\n", err.Error())
	}

	pathFileNameExt = currDirMgr.GetAbsolutePathWithSeparator() + "input"  +
		string(os.PathSeparator) + tzdatastructs.AppInputParametersFileName

	fMgr, err = pathfileops.FileMgr{}.New(pathFileNameExt)

	if err != nil {
		return pathfileops.DirMgr{},
			fmt.Errorf(ePrefix +
				"\nError returned by Current Directory pathfileops.FileMgr{}.New(pathFileNameExt)\n" +
				"pathFileNameExt='%v'\nError='%v'\n", err.Error())
	}

	if !fMgr.DoesFileExist() {
		return pathfileops.DirMgr{},
			errors.New(ePrefix +
				"\nError: Could NOT locate Base or Application Directory!\n")
	}

	return currDirMgr, nil
}