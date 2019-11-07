package fileops

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
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
