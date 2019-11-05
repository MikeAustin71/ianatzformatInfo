package inprocess

import "github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"

type ZoneInfoDataDto struct {
	ZoneInfoInputDir    string
	AppOutputDir        string
	AppLogfilePathName  string
	IanaTimeZoneVersion string
	ZoneInfoDirFileInfo pathfileops.FileMgrCollection
	ZoneInfoDirMgr      pathfileops.DirMgr
	AppOutputDirMgr     pathfileops.DirMgr
	AppLogFileMgr       pathfileops.FileMgr
}

