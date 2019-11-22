package tzdeclarations

import (
	"fmt"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strconv"
	"strings"
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
//
// func (ameri americaTimeZones) Chicago() string { return "America/Chicago" }
//
func (tZoneDecs TzZoneDeclarations) StandardZoneFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzZoneDeclarations.StandardGrpDeclaration() "

	if tzData == nil {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter 'tzData' is nil!\n")
	}

	b := strings.Builder{}

	b.Grow(2048)

	b.WriteString(tzdatastructs.CommentLead +
		fmt.Sprintf("%v - IANA Time Zone '%v'.\n",
			tzData.TzName,
			tzData.TzCanonicalValue))

	b.WriteString(tzdatastructs.CommentBlankLine)

	var err error

	idx := strings.Index(tzData.TzCanonicalValue, "Etc")

	if idx > -1 {

		if idx > 0 {
			return fmt.Errorf(ePrefix +
				"\nError: Expected index of 'Etc' in 'tzData.TzCanonicalValue' to be '0'\n" +
				"Actual index='%v'\ntzData.TzCanonicalValue='%v'\n",
				idx, tzData.TzCanonicalValue)
		}

		err = tZoneDecs.createEtcComments(&b, tzData, ePrefix)

		if err != nil {
			return err
		}

	}

	b.WriteString(fmt.Sprintf("func (%v %v) %v %v {return %v }\n",
			tzData.FuncSelfReferenceVariable,
			tzData.FuncType,
			tzData.FuncName,
			tzData.FuncReturnType,
			tzData.FuncReturnValue))

	b.WriteString("\n")

	tzData.FuncDeclaration = []byte(b.String())

	return nil
}

// createEtcComments - Creates Function Comments for Iana Etc Time Zones.
//
func (tZoneDecs TzZoneDeclarations) createEtcComments(
	b *strings.Builder,
	tzData *tzdatastructs.TimeZoneDataDto,
	ePrefix string) error {

	ePrefix += "TzZoneDeclarations.createEtcComments() "

	if !strings.HasPrefix(tzData.TzCanonicalValue, "Etc") {
		return fmt.Errorf(ePrefix +
			"\nError: 'tzData.TzCanonicalValue' does NOT contain \"Etc\"\n" +
			"tzData.TzCanonicalValue='%v'\n", tzData.TzCanonicalValue)
	}


	numStrProfile,
	err := strops.StrOps{}.ExtractNumericDigits(
		tzData.TzCanonicalValue,
		0,
		"-+",
		"",
		"")

	if err != nil {
	return fmt.Errorf(ePrefix +
		"\n")
	}

	var numValue int
	var newSign, utcOffset string

	if numStrProfile.NumStrLen == 0 {
		goto etcIanaDocs
	}

	newSign = "+"

	if numStrProfile.LeadingSignChar == "+" {

		newSign = "-"
	}

	numValue, err = strconv.Atoi(numStrProfile.NumStr)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by strconv.Atoi(numStrProfile.NumStr)\n" +
			"numStrProfile.NumStr='%v'\n" +
			"Error='%v'\n", numStrProfile.NumStr, err.Error())
	}

	if numValue < 0 {
		numValue = numValue * -1
	}


	if numValue == 0 {

		goto etcIanaDocs

	} else if numValue < 1000 {
		numValue = numValue * 100
	}


	utcOffset = fmt.Sprintf("UTC%v%04d", newSign, numValue)

	b.WriteString(
		"// This is an 'Etc' IANA Time Zone. The syntax for 'Etc' Time Zones\n" +
		  "// can be confusing. The numeric sign of 'Etc' offsets is the opposite\n")

	b.WriteString("// of the equivalent UTC offset. For example the 'Etc' Time Zone\n")

	b.WriteString(fmt.Sprintf("// %v has a UTC offset of %v. \n",
		tzData.TzCanonicalValue, utcOffset))

	b.WriteString(tzdatastructs.CommentBlankLine)

	etcIanaDocs:
	b.WriteString( tzdatastructs.CommentLead +
		"The 'Etc' time zone group is documented in the IANA Time Zone\n")

	b.WriteString(tzdatastructs.CommentLead +
		"Database at:\n")


	b.WriteString( tzdatastructs.CommentLead +
		"   https://en.wikipedia.org/wiki/Tz_database#Area\n")

	b.WriteString(tzdatastructs.CommentBlankLine)


	return nil
}
