package tzdeclarations

import (
	"fmt"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)

type TzZoneDeclarations struct {
	Comments []string
}

//
// Example
//
// --------------------------------------------------------------------
// Chicago - IANA Time Zone "America/Chicago"
//
// func (amer americaTimeZones) Chicago() string { return "America/Chicago" }
//
func (tZoneDecs TzZoneDeclarations) StandardGrpDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzZoneDeclarations.StandardGrpDeclaration() "

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - IANA Time Zone '%v'.\n\n", tzData.TzName, tzData.TzCanonicalValue)
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("func (%v %v) %v %v {return %v }",
			tzData.FuncSelfReferenceVariable,
			tzData.FuncDeclaration,
			tzData.FuncName,
			tzData.FuncReturnType,
			tzData.FuncReturnValue)

	tzData.FuncDeclaration = append(tzData.FuncDeclaration, []byte(outputStr) ...)

	return nil
}