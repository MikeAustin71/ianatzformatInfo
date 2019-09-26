package tzdeclarations

import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)

type TzGroupDeclarations struct {
	Comments []string
}


// StandardGrpDeclaration() - Generates comments and type declarations
// for standard groups.
//
// Example Standard Group Declaration
// ------------------------------------------------------------
// americaTimeZones - IANA Time Zones for 'America'.
//
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//
// type americaTimeZones string

func (tzGrpDecs TzGroupDeclarations) StandardGrpDeclaration(
	tzGroup *tzdatastructs.TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TzGroupDeclarations.StandardGrpDeclaration() "

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - IANA Time Zones for '%v'.\n",
			tzGroup.TypeName, tzGroup.GroupName)

	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "Reference:\n"
	outputStr += tzdatastructs.CommentLead + "  https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\n"
	outputStr += tzdatastructs.CommentLead + "  https://en.wikipedia.org/wiki/Tz_database\n"
	outputStr += tzdatastructs.CommentLead + "  https://www.iana.org/time-zones\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += fmt.Sprintf("type %v %v\n\n",
		tzGroup.TypeName, tzGroup.TypeValue)

	tzGroup.TypeDeclaration = append(tzGroup.TypeDeclaration, []byte(outputStr) ...)

	return nil
}