package tzdeclarations

import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)

type TzZoneDeclarations struct {
	Comments []string
}

// LinkTimeZoneOneElementDeclaration - Produces function declarations for
// 'Link' time zones. 'Link' time zones identify deprecated or
// obsolete time zones. These time zone links are mapped to valid
// current time zones.
//
// Example
//
//
// --------------------------------------------------------------------
//
// Egypt - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//
// Linked Deprecated Time Zone: 'Egypt'
// Maps To Valid Time Zone: 'Africa/Cairo'
//
// func (depri deprecatedTimeZones) Egypt() string { return "Africa/Cairo" }
//
func (tZoneDecs TzZoneDeclarations) LinkTimeZoneOneElementDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzZoneDeclarations.LinkTimeZoneOneElementDeclaration() "

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - This is an IANA 'Link' Time Zone. 'Link' Time Zones\n",
			tzData.TzName)

	outputStr += tzdatastructs.CommentLead +
		"Zones identify deprecated or obsolete time zones. These\n"

	outputStr += tzdatastructs.CommentLead +
		"obsolete time zones are mapped to valid current time zones.\n"

	outputStr += tzdatastructs.CommentBlankLine

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("Linked Deprecated Time Zone: '%v'\n",
			tzData.TzName)

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("Maps To Valid Time Zone: '%v'\n",
			tzData.TzCanonicalValue)

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("           IANA Source File: %v\n", tzData.SourceFileNameExt)

	outputStr += tzdatastructs.CommentBlankLine +
	fmt.Sprintf("func (%v %v) %v %v { return %v }\n",
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

// LinkTimeZoneTwoElementDeclaration - Produces function declarations
// for standard two element 'Link' time zones (i.e. 'US/Alaska').
// 'Link' time zones identify deprecated or obsolete time zones.
// These time zone links are mapped to valid current time zones.
//
// Example
//
// --------------------------------------------------------------------
//  Link                      canonical time zone
//  Obsolete Time Zone    ->  Valid Current Time Zone
// 'US/Alaska'            -> 'America/Anchorage'
//
// 'US/Alaska' - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//
// Linked Deprecated Time Zone: 'US/Alaska'
//     Maps To Valid Time Zone: 'America/Anchorage'
//            IANA Source File: backward
//
// func (uSDep uSDeprecatedTimeZones) Alaska() string { return "America/Anchorage" }
//
func (tZoneDecs TzZoneDeclarations) LinkTimeZoneTwoElementDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzZoneDeclarations.LinkTimeZoneOneElementDeclaration() "

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - This is an IANA 'Link' Time Zone. 'Link' Time Zones\n",
			tzData.TzName)

	outputStr += tzdatastructs.CommentLead +
		"identify deprecated or obsolete time zones. These obsolete\n"

	outputStr += tzdatastructs.CommentLead +
		"time zones are mapped to valid current time zones.\n"

	outputStr += tzdatastructs.CommentBlankLine

	// Linked Deprecated Time Zone: 'US/Alaska'
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("Linked Deprecated Time Zone: '%v/%v'\n",
			tzData.GroupName, tzData.TzName)

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("    Maps To Valid Time Zone: '%v'\n",
			tzData.TzCanonicalValue)

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("           IANA Source File: %v\n", tzData.SourceFileNameExt)

	outputStr += tzdatastructs.CommentBlankLine +
	fmt.Sprintf("func (%v %v) %v %v { return %v }\n",
		tzData.FuncSelfReferenceVariable,
		tzData.FuncType,
		tzData.FuncName,
		tzData.FuncReturnType,
		tzData.FuncReturnValue)

	outputStr += "\n"

	tzData.FuncDeclaration = []byte(outputStr)

	return nil
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
func (tZoneDecs TzZoneDeclarations) PlaceHolderZoneFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzZoneDeclarations.PlaceHolderZoneFuncDeclaration() "

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
func (tZoneDecs TzZoneDeclarations) PlaceHolderLinkFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzZoneDeclarations.PlaceHolderZoneFuncDeclaration() "

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
// func (amer americaTimeZones) Chicago() string { return "America/Chicago" }
//
func (tZoneDecs TzZoneDeclarations) StandardZoneFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzZoneDeclarations.StandardGrpDeclaration() "

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

	// tzData.FuncDeclaration = append(tzData.FuncDeclaration, []byte(outputStr) ...)
	tzData.FuncDeclaration = []byte(outputStr)

	return nil
}
