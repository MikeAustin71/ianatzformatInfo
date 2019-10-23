package tzdeclarations

import (
	"fmt"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
)

type TzMilitaryDeclarations struct {
	Comments []string
}

// MilitaryTypeDeclaration - Writes military time zone type
//
func (tzMilDecs TzMilitaryDeclarations) MilitaryTypeDeclaration(
	militaryTzType *tzdatastructs.TimeZoneGroupDto, ePrefix string) error {

	ePrefix += "TzMilitaryDeclarations.MilitaryTypeDeclaration() "

	outputStr := "\n"
	outputStr += fmt.Sprintf(tzdatastructs.CommentLead + "%v - Military Time Zone Names.\n", militaryTzType.GroupName)
	outputStr += tzdatastructs.CommentBlankLine

	outputStr += tzdatastructs.CommentLead + "Reference:\n"
	outputStr += tzdatastructs.CommentLead + "    https://en.wikipedia.org/wiki/List_of_military_time_zones\n"
	outputStr += tzdatastructs.CommentLead + "    https://www.timeanddate.com/time/zones/military\n"
	outputStr += tzdatastructs.CommentLead + "    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm\n"
	outputStr += tzdatastructs.CommentBlankLine

	outputStr += tzdatastructs.CommentLead + "Military time zones are commonly used in aviation as well as at sea.\n"
	outputStr += tzdatastructs.CommentLead + "They are also known as nautical or maritime time zones.\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "The 'J' (Juliet) Time Zone is occasionally used to refer to the observer's\n"
	outputStr += tzdatastructs.CommentLead + "local time. Note that Time Zone 'J' (Juliet) is not listed below.\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "  Abbreviation Time zone name     Other names    Offset\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "      A        Alpha Time Zone                   UTC +1\n"
	outputStr += tzdatastructs.CommentLead + "      B        Bravo Time Zone                   UTC +2\n"
	outputStr += tzdatastructs.CommentLead + "      C        Charlie Time Zone                 UTC +3\n"
	outputStr += tzdatastructs.CommentLead + "      D        Delta Time Zone                   UTC +4\n"
	outputStr += tzdatastructs.CommentLead + "      E        Echo Time Zone                    UTC +5\n"
	outputStr += tzdatastructs.CommentLead + "      F        Foxtrot Time Zone                 UTC +6\n"
	outputStr += tzdatastructs.CommentLead + "      G        Golf Time Zone                    UTC +7\n"
	outputStr += tzdatastructs.CommentLead + "      H        Hotel Time Zone                   UTC +8\n"
	outputStr += tzdatastructs.CommentLead + "      I        India Time Zone                   UTC +9\n"
	outputStr += tzdatastructs.CommentLead + "      K        Kilo Time Zone                    UTC +10\n"
	outputStr += tzdatastructs.CommentLead + "      L        Lima Time Zone                    UTC +11\n"
	outputStr += tzdatastructs.CommentLead + "      M        Mike Time Zone                    UTC +12\n"
	outputStr += tzdatastructs.CommentLead + "      N        November Time Zone                UTC -1\n"
	outputStr += tzdatastructs.CommentLead + "      O        Oscar Time Zone                   UTC -2\n"
	outputStr += tzdatastructs.CommentLead + "      P        Papa Time Zone                    UTC -3\n"
	outputStr += tzdatastructs.CommentLead + "      Q        Quebec Time Zone                  UTC -4\n"
	outputStr += tzdatastructs.CommentLead + "      R        Romeo Time Zone                   UTC -5\n"
	outputStr += tzdatastructs.CommentLead + "      S        Sierra Time Zone                  UTC -6\n"
	outputStr += tzdatastructs.CommentLead + "      T        Tango Time Zone                   UTC -7\n"
	outputStr += tzdatastructs.CommentLead + "      U        Uniform Time Zone                 UTC -8\n"
	outputStr += tzdatastructs.CommentLead + "      V        Victor Time Zone                  UTC -9\n"
	outputStr += tzdatastructs.CommentLead + "      W        Whiskey Time Zone                 UTC -10\n"
	outputStr += tzdatastructs.CommentLead + "      X        X-ray Time Zone                   UTC -11\n"
	outputStr += tzdatastructs.CommentLead + "      Y        Yankee Time Zone                  UTC -12\n"
	outputStr += tzdatastructs.CommentLead + "      Z        Zulu Time Zone                    UTC +0\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("The methods associated with type '%v' return the equivalent\n",
			militaryTzType.GroupName)

	outputStr += tzdatastructs.CommentLead + "IANA time zones. At first this may seem confusing. For example,\n"
	outputStr += tzdatastructs.CommentLead + "Military Time Zone 'L' or 'Lima' specifies UTC +11-hours.\n"
	outputStr += tzdatastructs.CommentLead + "However, the equivalent IANA Time Zone is \"Etc/GMT+11\".\n"
	outputStr += tzdatastructs.CommentLead + "In date time calculations, IANA Time Zone \"Etc/GMT-11\" \n"
	outputStr += tzdatastructs.CommentLead + "computes as UTC +11 hours.\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead + "  Reference:\n"
	outputStr += tzdatastructs.CommentLead + "    https://en.wikipedia.org/wiki/Tz_database#Area\n"
	outputStr += tzdatastructs.CommentBlankLine
	outputStr += "type " + militaryTzType.TypeName + "  " + militaryTzType.TypeValue + "\n\n"

	militaryTzType.TypeDeclaration = append(militaryTzType.TypeDeclaration, []byte(outputStr) ...)

	return nil
}

// MilitaryTzFuncDeclaration - Generates function declarations and file output
// for Military Time Zones
// ---------------------------------------------------------------------------
//
// Example:
//   Alpha - Military Time Zone 'A' or 'Alpha' is equivalent
//   IANA Time Zone "Etc/GMT+1"
//
//   Offset from UTC is computed at +1 hours.
//
//   func (milTz militaryTimeZones)Alpha() string { return "Etc/GMT+1" }
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

	if nStrDto.NumStrLen == 0 {
		nStrDto.NumStr = "0"
	}

	outputStr := tzdatastructs.CommentLead +
		fmt.Sprintf("%v - Military Time Zone '%v' or '%v' is equivalent to\n",
			tzData.TzName, firstLetter, tzData.TzName)

	outputStr += tzdatastructs.CommentLead +
		fmt.Sprintf("to IANA Time Zone \"%v\".\n", tzData.TzValue)

	utcOffset := nStrDto.NumStr + " hours."

	if nStrDto.NumStr == "+1" || nStrDto.NumStr == "-1" {
		utcOffset = nStrDto.NumStr + " hour."
	}

	outputStr += tzdatastructs.CommentBlankLine
	outputStr += tzdatastructs.CommentLead +
			fmt.Sprintf("Offset from Universal Coordinated Time (UTC) is computed at %v\n",
			utcOffset)

	outputStr += tzdatastructs.CommentBlankLine
	outputStr +=
		fmt.Sprintf("func (%v %v) %v %v {return %v }\n\n",
			tzData.FuncSelfReferenceVariable,
			tzData.FuncType,
			tzData.FuncName,
			tzData.FuncReturnType,
			tzData.FuncReturnValue)

	tzData.FuncDeclaration = append(tzData.FuncDeclaration, []byte(outputStr) ...)

	return nil
}
