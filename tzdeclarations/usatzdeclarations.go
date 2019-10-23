package tzdeclarations

import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)

type TzUSADeclarations struct {
	Comments []string
}

// USATzTypeDeclaration - Writes USA time zone type
// ------------------------------------------------------------
//
// uSATimeZones - Defines a collection of time zones in the
// continental United States of America.  These time zones
// are considered deprecated by the IANA Time Zone Database.
// Nevertheless, they remain in continuous every day use within
// the United States of America.
//
// Each member of this collection will be accessed by its common
// name.  These names include 'Eastern', 'Central', 'Mountain'
// and 'Pacific' Time.  However, the returned values associated
// with these common names will be mapped to a current, valid
// IANA time zone such as 'America/New_York', 'America/Chicago',
// 'America/Denver' or 'America/Los_Angeles'.
//
// The 'uSATimeZones' Group is therefore an artificial creation
// outside of the accepted IANA Time Zones and is provided for
// convenience and in recognition of everyday usage and practice
//
// These same USA time zones can also be accessed through the
// 'TimeZones' variable 'Deprecated' which maps to the group,
// 'deprecatedTimeZones'.
//
// For documentation on IANA Time Zones, see type 'TimeZones'.
//
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//
// type uSATimeZones string
//
// USATzTypeDeclaration - Writes USA time zone type
//
func (tzUSADec TzUSADeclarations) USATzTypeDeclaration(
	usaTzType *tzdatastructs.TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TzUSADeclarations.USATzTypeDeclaration() "
	outputStr := "\n"
	outputStr += fmt.Sprintf(
		tzdatastructs.CommentLead + "%v - Defines a collection of time zones in the continental\n",
		usaTzType.TypeName)
	outputStr += tzdatastructs.CommentLead + "United States of America, its territories and associated\n"
	outputStr += tzdatastructs.CommentLead + "commonwealth entities.\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "These time zones are considered deprecated by the IANA Time\n"
	outputStr += tzdatastructs.CommentLead + "Zone Database.  Nevertheless, they remain in continuous, every\n"
	outputStr += tzdatastructs.CommentLead + "day use within the United States of America.\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "Each member of this collection can be accessed by its common\n"
	outputStr += tzdatastructs.CommentLead + "name.  These common names include 'Eastern', 'Central', 'Mountain'\n"
	outputStr += tzdatastructs.CommentLead + "and 'Pacific' Time.  However, the returned values associated with\n"
	outputStr += tzdatastructs.CommentLead + "these common names will be mapped to a current, valid IANA time\n"
	outputStr += tzdatastructs.CommentLead + "zone such as 'America/New_York', 'America/Chicago', 'America/Denver'\n"
	outputStr += tzdatastructs.CommentLead + "or 'America/Los_Angeles'.\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("The '%v'Time Zone Group is therefore an artificial creation\n", usaTzType.GroupName)
	outputStr += tzdatastructs.CommentLead + "outside of the accepted IANA Time Zones and is provided for\n"
	outputStr += tzdatastructs.CommentLead + "convenience and in recognition of everyday usage and practice.\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "These same USA time zones can also be accessed through the\n"
	outputStr += tzdatastructs.CommentLead + "'TimeZones' variable 'Deprecated' which maps to the group,\n"
	outputStr += tzdatastructs.CommentLead + "'deprecatedTimeZones'.\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "For documentation on IANA Time Zones, see type 'TimeZones'.\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "Reference:\n"
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzList
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzDatabase
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefIanaOrgTimeZones
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += "type " + usaTzType.TypeName + "  " + usaTzType.TypeValue + "\n\n"

	usaTzType.TypeDeclaration = append(usaTzType.TypeDeclaration, []byte(outputStr) ...)

	return nil
}


// USAZoneFuncDeclaration - Produces function declarations for
// standard IANA time zones.
//
// Example
//
// --------------------------------------------------------------------
//
// Eastern - USA Eastern Time Zone
// Maps to IANA Time Zone 'America/New_York'.
//
// func (usaTz uSATimeZones) Eastern() string { return "America/New_York" }
//
func (tzUSADec TzUSADeclarations) USAZoneFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzZoneDeclarations.StandardGrpDeclaration() "

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - USA %v Time Zone.\n",
			tzData.TzName,
			tzData.TzName)

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("Maps to IANA Time Zone '%v'\n",
			tzData.TzCanonicalValue)

	outputStr += tzdatastructs.CommentBlankLine
	outputStr += fmt.Sprintf("func (%v %v) %v %v {return %v }\n",
		tzData.FuncSelfReferenceVariable,
		tzData.FuncType,
		tzData.FuncName,
		tzData.FuncReturnType,
		tzData.FuncReturnValue)

	outputStr += "\n"

	tzData.FuncDeclaration = []byte(outputStr)

	return nil
}
