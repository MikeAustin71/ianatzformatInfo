package tzdeclarations

import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)

type ZoneInfoZoneDeclarations struct {
	Comments []string
}

// PlaceHolderZoneFuncDeclaration
// Example
//
// --------------------------------------------------------------------
//
// Argentina - A place holder which defines a sub-group
// of IANA Time Zones.
//
// func (amer americaTimeZones) Argentina() argentinaTimeZones { return argentinaTimeZones("") }
//
func (zoneInfoDecs ZoneInfoZoneDeclarations) PlaceHolderZoneFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "ZoneInfoZoneDeclarations.PlaceHolderZoneFuncDeclaration() "

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - A place holder which defines a sub-group\n",
			tzData.TzName)

	outputStr += tzdatastructs.CommentLead +
		"of IANA Time Zones.\n"

	outputStr += tzdatastructs.CommentBlankLine

	outputStr += fmt.Sprintf("func (%v %v) %v %v {return %v }\n",
		tzData.FuncSelfReferenceVariable,
		tzData.FuncType,
		tzData.FuncName,
		tzData.FuncReturnType,
		tzData.FuncReturnValue)

	outputStr += "\n"

	// tzData.FuncDeclaration = append(tzData.FuncDeclaration, []byte(outputStr) ...)
	tzData.FuncDeclaration = []byte(outputStr)

	return nil
}

// PlaceHolderLinkFuncDeclaration - Generates a function
// declaration for a 'Link' Zone which defines a group of
// 'link' zones.
//
// 'Link' Time Zones identify deprecated or obsolete time
// zones. These obsolete time zones are mapped to valid
// current time zones.
//
// --------------------------------------------------------------------
//
// Example:
//
//   Link        -> canonical time zone
//   'US/Alaska' -> 'America/Anchorage'
//
//   func (depre deprecatedTimeZones)
//    US() uSDeprecatedTimeZones { return uSDeprecatedTimeZones("") }
//
func (zoneInfoDecs ZoneInfoZoneDeclarations) PlaceHolderLinkFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "ZoneInfoZoneDeclarations.PlaceHolderZoneFuncDeclaration() "

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - A place holder which defines a sub-group\n",
			tzData.TzName)

	outputStr += tzdatastructs.CommentLead +
		"of IANA 'Link' Time Zones.\n"

	outputStr += tzdatastructs.CommentBlankLine

	outputStr += fmt.Sprintf("func (%v %v) %v %v {return %v }\n",
		tzData.FuncSelfReferenceVariable,
		tzData.FuncType,
		tzData.FuncName,
		tzData.FuncReturnType,
		tzData.FuncReturnValue)

	outputStr += "\n"

	// tzData.FuncDeclaration = append(tzData.FuncDeclaration, []byte(outputStr) ...)

	tzData.FuncDeclaration = []byte(outputStr)

	return nil
}

// StandardZoneFuncDeclaration - Produces function declarations for
// standard IANA time zones.
//
// Example
//
// --------------------------------------------------------------------
//
// Chicago - IANA Time Zone "America/Chicago"
// IANA Source File: northamerica
//
// func (ameri americaTimeZones) Chicago() string { return "America/Chicago" }
//
func (zoneInfoDecs ZoneInfoZoneDeclarations) StandardZoneFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "ZoneInfoZoneDeclarations.StandardGrpDeclaration() "

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - IANA Time Zone '%v'.\n",
			tzData.TzName,
			tzData.TzCanonicalValue)

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("IANA Source File: %v\n",
			tzData.SourceFileNameExt)

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
