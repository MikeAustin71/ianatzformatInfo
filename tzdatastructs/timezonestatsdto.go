package tzdatastructs

type TimeZoneStatsDto struct {
	IanaVersion           string
	NumStdIanaTZones      int
	NumLinkIanaTZones     int
	TotalIanaTZones       int
	NumMilitaryTZones     int
	NumOtherTZones        int
	TotalTZones           int
	NumPrimaryTZoneGroups int
	NumSubStdTZoneGroups  int
	NumSubLinkTZoneGroups int
	TotalSubTZoneGroups   int
}
