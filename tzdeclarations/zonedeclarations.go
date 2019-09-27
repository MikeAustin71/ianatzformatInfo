package tzdeclarations

import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)

type TzZoneDeclarations struct {
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

	return nil
}
// StandardZoneFuncDeclaration
//
// Example
//
// --------------------------------------------------------------------
//
// Chicago - IANA Time Zone "America/Chicago"
//
// func (amer americaTimeZones) Chicago() string { return "America/Chicago" }
//
func (tZoneDecs TzZoneDeclarations) StandardZoneFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzZoneDeclarations.StandardGrpDeclaration() "

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - IANA Time Zone '%v'.\n", tzData.TzName, tzData.TzCanonicalValue)
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("func (%v %v) %v %v {return %v }\n",
			tzData.FuncSelfReferenceVariable,
			tzData.FuncDeclaration,
			tzData.FuncName,
			tzData.FuncReturnType,
			tzData.FuncReturnValue)

	outputStr += "\n"

	tzData.FuncDeclaration = append(tzData.FuncDeclaration, []byte(outputStr) ...)

	return nil
}
