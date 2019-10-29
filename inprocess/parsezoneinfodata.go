package inprocess

import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)

type ParseZoneInfoData struct {
	input string
	output string
}

func (parseZInfo ParseZoneInfoData) ParseZoneInfo(
	zInfoDto ZoneInfoDataDto, ePrefix string) (tzdatastructs.TimeZoneStatsDto, error) {

	ePrefix += "ParseZoneInfoData.ParseZoneInfo()"
	tzStats := tzdatastructs.TimeZoneStatsDto{}

	numOfZoneInfoFMgrs := zInfoDto.ZoneInfoDirFileInfo.GetNumOfFileMgrs()

	if numOfZoneInfoFMgrs < 30 {
		return tzStats, fmt.Errorf(ePrefix +
			"Error: Number of ZoneInfo files is less than 30.\n" +
			"Total Zone Info files='%v'\n", numOfZoneInfoFMgrs)
	}


}
