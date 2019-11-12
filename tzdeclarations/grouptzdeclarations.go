package tzdeclarations

import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
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

	b := strings.Builder{}

	b.Grow(1024)


	b.WriteString("\n")
	b.WriteString(tzdatastructs.CommentLead +
		fmt.Sprintf("%v - IANA Time Zones for '%v'.\n",
			tzGroup.TypeName, tzGroup.GroupName))

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
		"For documentation on IANA Time Zones, see type\n")

	b.WriteString(tzdatastructs.CommentLead +
		fmt.Sprintf("'%v'.\n",
			tzdatastructs.PrimaryTimeZoneType))

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead + "Reference:\n")

	b.WriteString(tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzList)

	b.WriteString(tzdatastructs.CommentLead + tzdatastructs.RefWikipediaTzDatabase)

	b.WriteString(tzdatastructs.CommentLead + tzdatastructs.RefIanaOrgTimeZones)

	b.WriteString(tzdatastructs.CommentBlankLine)

	if tzGroup.GroupName == "Etc" {
		tzGrpDecs.CreateETCComments(&b)
	}

	if tzGroup.GroupName == "Other" {
		tzGrpDecs.CreateOtherGroupComments(&b)
	}

	b.WriteString(fmt.Sprintf("type %v %v\n",
		tzGroup.TypeName, tzGroup.TypeValue))

	b.WriteString("\n")


	tzGroup.TypeDeclaration = []byte(b.String())

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


// CreateETCComments - Generates comments specific to the 'Etc'
// Group.
func (tzGrpDecs TzGroupDeclarations) CreateETCComments(
	b *strings.Builder) {

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString( tzdatastructs.CommentLead +
		"The 'ETC' group is referenced at the IANA Time Zone Database at:\n")

	b.WriteString( tzdatastructs.CommentLead +
		"   https://en.wikipedia.org/wiki/Tz_database#Area\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
		"The special area of \"Etc\" is used for some administrative zones,\n")

	b.WriteString(tzdatastructs.CommentLead +
		"particularly for \"Etc/UTC\" which represents Coordinated Universal Time.\n")

	b.WriteString(tzdatastructs.CommentLead +
		"In order to conform with the POSIX style, those zone names beginning with\n")

	b.WriteString(tzdatastructs.CommentLead +
		"\"Etc/GMT\" have their sign reversed from the standard ISO 8601 convention.\n")

	b.WriteString(tzdatastructs.CommentLead +
		"In the \"Etc\" area, zones west of GMT have a positive sign and those east have\n")

	b.WriteString(tzdatastructs.CommentLead +
		"a negative sign in their name (e.g \"Etc/GMT-14\" is 14 hours ahead of GMT).\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	return
}

// CreateOtherGroupComments - Generates comments specific to the 'Other' Group of
// IANA Time Zones.
func (tzGrpDecs TzGroupDeclarations) CreateOtherGroupComments(b *strings.Builder) {

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString( tzdatastructs.CommentLead +
		"The 'Other' IANA Time Zone Group contains deprecated or obsolete time zones as\n")

	b.WriteString(tzdatastructs.CommentLead +
		"well as time zone abbreviations.  All deprecated time zones map to current,\n")

	b.WriteString(tzdatastructs.CommentLead +
		"valid time zones.\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	return
}

