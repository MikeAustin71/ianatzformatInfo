package tzdatastructs

type TimeZoneStatsDto struct {
	IanaVersion           string
	NumStdIanaTZones          int
	NumLinkIanaTZones         int
	TotalIanaTZones           int
	NumOfLinkConflictResolved int
	NumOfBackZoneConflicts    int
	NumMilitaryTZones         int
	NumOtherTZones            int
	TotalTZones               int
	NumPrimaryTZoneGroups     int
	NumSubStdTZoneGroups      int
	NumSubLinkTZoneGroups     int
	TotalSubTZoneGroups       int
}
