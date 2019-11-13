package tzdeclarations

import (
	"fmt"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"strings"
)

type TzMilitaryDeclarations struct {
	Comments []string
}

// MilitaryTypeDeclaration - Writes military time zone type
//
func (tzMilDecs TzMilitaryDeclarations) MilitaryTypeDeclaration(
	militaryTzType *tzdatastructs.TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TzMilitaryDeclarations.MilitaryTypeDeclaration() "
	//
	//

	b := strings.Builder{}

	b.Grow(2048)

	b.WriteString("\n")

	b.WriteString(fmt.Sprintf(tzdatastructs.CommentLead +
		"%v - Military Time Zone Names.\n", militaryTzType.GroupName))

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
		"Reference:\n")

	b.WriteString(tzdatastructs.CommentLead +
		"    https://en.wikipedia.org/wiki/List_of_military_time_zones\n")

	b.WriteString(tzdatastructs.CommentLead +
		"    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm\n")

	b.WriteString(tzdatastructs.CommentLead +
		"    https://www.timeanddate.com/time/zones/military\n")

	b.WriteString(tzdatastructs.CommentLead +
		"    https://www.timeanddate.com/worldclock/timezone/alpha\n")

	b.WriteString(tzdatastructs.CommentLead +
		"    https://www.timeanddate.com/time/zones/military\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
		"Military time zones are commonly used in aviation as well as at sea.\n")

	b.WriteString(tzdatastructs.CommentLead +
		"They are also known as nautical or maritime time zones.\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
		"The 'J' (Juliet) Time Zone is occasionally used to refer to the observer's\n")

	b.WriteString(tzdatastructs.CommentLead +
		"local time. Note that Time Zone 'J' (Juliet) is not listed below.\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
		"  Abbreviation Time zone name     Other names    Offset\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
		"      A        Alpha Time Zone                   UTC +1\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      B        Bravo Time Zone                   UTC +2\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      C        Charlie Time Zone                 UTC +3\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      D        Delta Time Zone                   UTC +4\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      E        Echo Time Zone                    UTC +5\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      F        Foxtrot Time Zone                 UTC +6\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      G        Golf Time Zone                    UTC +7\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      H        Hotel Time Zone                   UTC +8\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      I        India Time Zone                   UTC +9\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      K        Kilo Time Zone                    UTC +10\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      L        Lima Time Zone                    UTC +11\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      M        Mike Time Zone                    UTC +12\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      N        November Time Zone                UTC -1\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      O        Oscar Time Zone                   UTC -2\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      P        Papa Time Zone                    UTC -3\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      Q        Quebec Time Zone                  UTC -4\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      R        Romeo Time Zone                   UTC -5\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      S        Sierra Time Zone                  UTC -6\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      T        Tango Time Zone                   UTC -7\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      U        Uniform Time Zone                 UTC -8\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      V        Victor Time Zone                  UTC -9\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      W        Whiskey Time Zone                 UTC -10\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      X        X-ray Time Zone                   UTC -11\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      Y        Yankee Time Zone                  UTC -12\n")

	b.WriteString(tzdatastructs.CommentLead +
		"      Z        Zulu Time Zone                    UTC +0\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
		fmt.Sprintf("The methods associated with type '%v' return the equivalent\n",
			militaryTzType.GroupName))

	b.WriteString(tzdatastructs.CommentLead +
		"IANA time zones. At first this may seem confusing. For example,\n")

	b.WriteString(tzdatastructs.CommentLead +
		"Military Time Zone 'L' or 'Lima' specifies UTC +11-hours.\n")

	b.WriteString(tzdatastructs.CommentLead +
		"However, the equivalent IANA Time Zone is \"Etc/GMT+11\".\n")

	b.WriteString(tzdatastructs.CommentLead +
		"In date time calculations, IANA Time Zone \"Etc/GMT-11\" \n")

	b.WriteString(tzdatastructs.CommentLead +
		"computes as UTC +11 hours.\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
		"  Reference:\n")

	b.WriteString(tzdatastructs.CommentLead +
		"    https://en.wikipedia.org/wiki/Tz_database#Area\n")

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString("type " + militaryTzType.TypeName + "  " + militaryTzType.TypeValue + "\n\n")

	militaryTzType.TypeDeclaration = append(militaryTzType.TypeDeclaration, []byte(b.String()) ...)

	return nil
}

// MilitaryTzFuncDeclaration - Generates function declarations and file output
// for Military Time Zones
// ---------------------------------------------------------------------------
//
// Example:
//   Alpha - Military Time Zone 'A' or 'Alpha' is equivalent TO
//   IANA Time Zone "Etc/GMT-1"
//
//   Offset from UTC is computed at +1 hours.
//
//   func (milTz militaryTimeZones)Alpha() string { return "Etc/GMT-1" }
//
// ----------------------------------------------------------------------------
//
func (tzMilDecs TzMilitaryDeclarations) MilitaryTzFuncDeclaration(
	tzData *tzdatastructs.TimeZoneDataDto, ePrefix string) error {

	ePrefix += "TzMilitaryDeclarations.MilitaryTzFuncDeclaration() "
	firstLetter := tzData.TzName[:1]

	nStrDto, err := strops.StrOps{}.ExtractNumericDigits(
		tzData.TzValue, 0, "+-","", "")

	if err != nil {
		return fmt.Errorf(ePrefix + "\n%v\n", err.Error())
	}

	b := strings.Builder{}
	b.Grow(512)

	utcEquivalent, ok := tzdatastructs.MilitaryUTCMap[tzData.TzName]

	if !ok {
		return fmt.Errorf(ePrefix +
			"\nError: Invlalid Military Time Zone Name!\n" +
			"TzName='%v'\n", tzData.TzName)
	}


	b.WriteString(tzdatastructs.CommentLead +
		fmt.Sprintf("%v - Military Time Zone '%v' or '%v' is equivalent to\n",
			tzData.TzName, firstLetter, tzData.TzName))

	b.WriteString(tzdatastructs.CommentLead +
		fmt.Sprintf("to IANA Time Zone \"%v\".\n", tzData.TzValue))

	utcOffset := utcEquivalent + " hours"

	if nStrDto.NumStr == "+1" || nStrDto.NumStr == "-1" {
		utcOffset = utcEquivalent + " hour"
	}

	b.WriteString(tzdatastructs.CommentBlankLine)

	b.WriteString(tzdatastructs.CommentLead +
			fmt.Sprintf("Offset from Universal Coordinated Time (UTC) is computed at %v.\n",
			utcOffset))

	b.WriteString(tzdatastructs.CommentBlankLine)

	if tzData.TzName != "Zulu" {
		b.WriteString(tzdatastructs.CommentLead +
			fmt.Sprintf("If the reversal of signs necessary to generate UTC%v is\n",utcOffset))

		b.WriteString(tzdatastructs.CommentLead +
			"confusing, see IANA the documentation for the 'ETC' Time Zone Area\n")

		b.WriteString(tzdatastructs.CommentLead +
			"referenced at:\n")

		b.WriteString(tzdatastructs.CommentBlankLine)

		b.WriteString(tzdatastructs.CommentLead +
			"   https://en.wikipedia.org/wiki/Tz_database#Area\n")

		b.WriteString(tzdatastructs.CommentBlankLine)
	}

	b.WriteString(
		fmt.Sprintf("func (%v %v) %v %v {return %v }\n\n",
			tzData.FuncSelfReferenceVariable,
			tzData.FuncType,
			tzData.FuncName,
			tzData.FuncReturnType,
			tzData.FuncReturnValue))

	tzData.FuncDeclaration = append(tzData.FuncDeclaration, []byte(b.String()) ...)

	return nil
}
