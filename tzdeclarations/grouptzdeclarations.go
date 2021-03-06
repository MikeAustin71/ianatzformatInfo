package tzdeclarations

import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)

type TzGroupDeclarations struct {
	Comments []string
}

// DeprecatedGrpDeclaration() - Generates comments and type declarations
// for the group, 'Deprecated'.
//
// Example Deprecated Group Declaration
//
// ------------------------------------------------------------
//
// deprecatedTimeZones - Defines a collection of IANA
// Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations.
//
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//
// type deprecatedTimeZones string
//
func (tzGrpDecs TzGroupDeclarations) DeprecatedGrpDeclaration(
	tzGroup *tzdatastructs.TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TzGroupDeclarations.DeprecatedGrpDeclaration() "

	outputStr := "\n"
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("%v - Defines a collection of IANA\n",
			tzGroup.GroupName)

	outputStr += tzdatastructs.CommentLead +
		"Time Zones which are obsolete and no longer used as\n"

	outputStr += tzdatastructs.CommentLead +
		"primary and accepted time zone designations. These\n"

	outputStr += tzdatastructs.CommentLead +
		"time zones are classified as 'Link' Time Zones and\n"

	outputStr += tzdatastructs.CommentLead +
		"are mapped to current, valid IANA Time Zones.\n"

	outputStr += tzdatastructs.CommentBlankLine

	outputStr += tzdatastructs.CommentLead +
		"For documentation on IANA Time Zones, see type\n"

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("'%v'.\n",
			tzdatastructs.PrimaryTimeZoneType)

	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "Reference:\n"
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzList
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzDatabase
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefIanaOrgTimeZones
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += fmt.Sprintf("type %v %v\n",
		tzGroup.TypeName, tzGroup.TypeValue)

	outputStr += "\n"

	// tzGroup.TypeDeclaration = append(tzGroup.TypeDeclaration, []byte(outputStr) ...)
	tzGroup.TypeDeclaration = []byte(outputStr)

	return nil
}

// DeprecatedSubGrpDeclaration() - Generates comments and type
// declarations for the Deprecated Sub-Groups.
//
// Example Deprecated Group Declaration
//
// ------------------------------------------------------------
//
// uSDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations.
//
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//
// type uSDeprecatedTimeZones string
//
func (tzGrpDecs TzGroupDeclarations) DeprecatedSubGrpDeclaration(
	tzGroup *tzdatastructs.TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TzGroupDeclarations.DeprecatedGrpDeclaration() "

	outputStr := "\n"
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("%v - Defines a subsidiary collection\n",
			tzGroup.TypeName)

	outputStr += tzdatastructs.CommentLead +
		"of Time Zones which are obsolete and no longer used as\n"

	outputStr += tzdatastructs.CommentLead +
		"primary and accepted time zone designations. These\n"

	outputStr += tzdatastructs.CommentLead +
		"time zones are classified as 'Link' Time Zones and\n"

	outputStr += tzdatastructs.CommentLead +
		"are mapped to current, valid IANA Time Zones.\n"


	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead +
		"For documentation on IANA Time Zones, see type\n"

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("'%v'.\n",
			tzdatastructs.PrimaryTimeZoneType)

	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "Reference:\n"
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzList
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzDatabase
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefIanaOrgTimeZones
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += fmt.Sprintf("type %v %v\n",
		tzGroup.TypeName, tzGroup.TypeValue)

	outputStr += "\n"

	// tzGroup.TypeDeclaration = append(tzGroup.TypeDeclaration, []byte(outputStr) ...)
	tzGroup.TypeDeclaration = []byte(outputStr)

	return nil
}

// StandardGrpDeclaration() - Generates comments and type declarations
// for standard groups.
//
//
// Example Standard Group Declaration
//
// ------------------------------------------------------------
//
// americaTimeZones - IANA Time Zones for 'America'.
//
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//
// type americaTimeZones string
//
func (tzGrpDecs TzGroupDeclarations) StandardGrpDeclaration(
	tzGroup *tzdatastructs.TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TzGroupDeclarations.StandardGrpDeclaration() "

	outputStr := "\n"
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("%v - IANA Time Zones for '%v'.\n",
			tzGroup.TypeName, tzGroup.GroupName)
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead +
		"For documentation on IANA Time Zones, see type\n"
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("'%v'.\n",
			tzdatastructs.PrimaryTimeZoneType)

	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "Reference:\n"
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzList
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzDatabase
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefIanaOrgTimeZones
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += fmt.Sprintf("type %v %v\n",
		tzGroup.TypeName, tzGroup.TypeValue)

	outputStr += "\n"


	//tzGroup.TypeDeclaration = append(tzGroup.TypeDeclaration, []byte(outputStr) ...)
	tzGroup.TypeDeclaration = []byte(outputStr)

	return nil
}

// Example Sub- Group Declaration
//
// ------------------------------------------------------------
//
// argentinaTimeZones - A Sub-Group of Time Zones. These are
// IANA Time Zones in located in 'Argentina'.
//
// The Parent Group is 'America'.
//
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//
// type argentinaTimeZones string
//
func (tzGrpDecs TzGroupDeclarations) SubGroupDeclaration(
	tzGroup *tzdatastructs.TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TzGroupDeclarations) SubGroupDeclaration() "

	outputStr := "\n"
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("%v - A Sub-Group of Time Zones. These are\n",
			tzGroup.TypeName)
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("IANA Time Zones located in '%v'.\n",
			tzGroup.GroupName)
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("The Parent Group is '%v'.\n",
			tzGroup.ParentGroupName)
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead +
		"For documentation on IANA Time Zones, see type\n"
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("'%v'.\n",
			tzdatastructs.PrimaryTimeZoneType)

	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "Reference:\n"
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzList
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzDatabase
	outputStr += tzdatastructs.CommentLead + tzdatastructs.RefIanaOrgTimeZones
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += fmt.Sprintf("type %v %v\n",
		tzGroup.TypeName, tzGroup.TypeValue)

	outputStr += "\n"

	// tzGroup.TypeDeclaration = append(tzGroup.TypeDeclaration, []byte(outputStr) ...)
	tzGroup.TypeDeclaration = []byte(outputStr)

	return nil
}