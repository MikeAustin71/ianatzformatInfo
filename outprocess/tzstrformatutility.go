package outprocess

import (
	"fmt"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"time"
)


type TzStrFmt struct {
	InputStr     string
	OutputStr    string
}

// LeftJustifyField - Left Justifies parameter 'str' in a field of length 'fieldLen'.
// If parameter, 'addTrailDots' is 'true', empty spaces to the right of 'str' are
// converted to dots ('.').
func (tzFmtStr TzStrFmt) LeftJustifyField(
	str string,
	fieldLen int,
	addTrailDots bool,
	ePrefix string) (string, error){

	ePrefix += "TzStrFmt.LeftJustifyField() "

	lenStr := len(str)

	if lenStr > fieldLen {
		return "", fmt.Errorf(ePrefix +
			"\nError: Input parameter 'str' exceeds field length parameter, 'fieldLen'.\n" +
			"str='%v'\n" +
			"str length='%v'\n" +
			"fieldLen='%v'\n", str, lenStr, fieldLen)
	}

	if lenStr == fieldLen {
		return str, nil
	}

	var padChar rune
	var padStr string
	var err error

	if !addTrailDots {
		padChar = ' '
	} else {
		padChar = '.'
	}

	padStr, err = strops.StrOps{}.MakeSingleCharString(padChar, fieldLen - lenStr)

	if err != nil {
		return "", fmt.Errorf(ePrefix +
			"\nError returned by StrOps{}.MakeSingleCharString(padChar, fieldLen - lenStr)\n" +
			"padChar='%v'   fieldLen='%v'   lenStr='%v'\n" +
			"Error='%v'\n", padChar, fieldLen, lenStr, err.Error())
	}


	return str + padStr, nil
}

// RightJustifyStr - Right Justifies a string in a field length specified by input
// parameter 'fieldLen'.
func (tzFmtStr TzStrFmt) RightJustifyStr(str string, fieldLen int, ePrefix string) (string, error) {

	ePrefix += "TzStrFmt.RightJustifyStr() "

	lenStr := len(str)

	if lenStr > fieldLen {
		return "", fmt.Errorf(ePrefix +
			"\nError: Input Parameter 'str' exceeds 'fieldLen'.\n" +
			"str length='%v'\n" +
			"fieldLen='%v'\n", lenStr, fieldLen)
	}

	if lenStr == fieldLen {
		return str, nil
	}

	padLen := fieldLen - lenStr

	padStr, err := strops.StrOps{}.MakeSingleCharString(' ', padLen)

	if err != nil {
		return "", fmt.Errorf(ePrefix +
			"\nError returned by strops.StrOps{}.MakeSingleCharString(' ', padLen)\n" +
			"Error='%v'\n", err.Error())
	}

	return padStr + str, nil
}

// RightJustifyNum - Right Justifies an integer number. The number is first
// converted to a string right justified according a numeric field length
// specified by parameter, 'numFieldLen'. This result is then right justified
// in a field of total field length specified by parameter, 'totalFieldLen'.
// If parameter, 'addLeadingDots' is 'true', the numeric portion if the field
// is prefixed with leading dots ('.').
func (tzFmtStr TzStrFmt) RightJustifyNum(
	num int,
	numFieldLen int,
	totalFieldLen int,
	addLeadingDots bool,
	ePrefix string) (string, error) {

	ePrefix += "TzStrFmt.RightJustifyNum() "
	fmtStr := "%" + string(numFieldLen) + "d"
	numStr := fmt.Sprintf(fmtStr, num)

	if totalFieldLen < numFieldLen {
		return "", fmt.Errorf(ePrefix +
			"\nError: Input parameter 'totalFieldLen' is less than parameter 'numFieldLen'.\n" +
			"totalFieldLen='%v'\n" +
			"numFieldLen='%v'\n", totalFieldLen, numFieldLen)
	}

	if totalFieldLen == numFieldLen {
		return numStr, nil
	}

	padLen := totalFieldLen - numFieldLen
	var padStr string
	var err error
	var padChar rune

	if !addLeadingDots {
		padChar = ' '
	} else {
		padChar = '.'
	}

	padStr, err = strops.StrOps{}.MakeSingleCharString(padChar, padLen)

	if err != nil {
		return "", fmt.Errorf(ePrefix +
			"\nError returned by StrOps{}.MakeSingleCharString(padChar, padLen).\n" +
			"padChar='%v'    padLen='%v'\n" +
			"Error='%v'\n", padChar, padLen, err.Error())
	}

	return padStr + numStr, nil
}

// ElapsedTime - Calculates and formats time duration. Usually this means
// subtracting the starting time from the ending time and returning the
// duration result as a string.
func (tzFmtStr TzStrFmt) ElapsedTime(starTime, endTime time.Time) string {

	// microSecondNanoseconds - Number of Nanoseconds in a Microsecond
	// 	A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
	microSecondNanoseconds := int64(time.Microsecond)

	// milliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
	//	 A millisecond is 1/1,000 or 1 one-thousandth of a second
	milliSecondNanoseconds := int64(time.Millisecond)

	// secondNanoseconds - Number of Nanoseconds in a Second
	secondNanoseconds := int64(time.Second)

	// minuteNanoseconds - Number of Nanoseconds in a minute
	minuteNanoseconds := int64(time.Minute)

	// hourNanoseconds - Number of Nanoseconds in an hour
	hourNanoseconds := int64(time.Hour)

	dayNanoseconds := int64(time.Hour) * int64(24)

	t2Dur := endTime.Sub(starTime)

	str := ""

	totalNanoseconds := t2Dur.Nanoseconds()

	var numOfDays, numOfHours, numOfMinutes, numOfSeconds, numOfMilliseconds,
		numOfMicroseconds, numOfNanoseconds int64

	if totalNanoseconds >= dayNanoseconds {
		numOfDays = totalNanoseconds / dayNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfDays * dayNanoseconds)
	}

	if totalNanoseconds >= hourNanoseconds {
		numOfHours = totalNanoseconds / hourNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfHours * hourNanoseconds)
	}

	if totalNanoseconds >= minuteNanoseconds {
		numOfMinutes = totalNanoseconds / minuteNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMinutes * minuteNanoseconds)
	}

	if totalNanoseconds >= secondNanoseconds {
		numOfSeconds = totalNanoseconds / secondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfSeconds * secondNanoseconds)
	}

	if totalNanoseconds >= secondNanoseconds {
		numOfSeconds = totalNanoseconds / secondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfSeconds * secondNanoseconds)
	}

	if totalNanoseconds >= milliSecondNanoseconds {
		numOfMilliseconds = totalNanoseconds / milliSecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMilliseconds * milliSecondNanoseconds)
	}

	if totalNanoseconds >= microSecondNanoseconds {
		numOfMicroseconds = totalNanoseconds / microSecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMicroseconds * microSecondNanoseconds)
	}

	numOfNanoseconds = totalNanoseconds

	if numOfDays > 0 {
		str += fmt.Sprintf("%v-Days ", numOfDays)
	}

	if numOfHours > 0 || str != "" {

		str += fmt.Sprintf("%v-Hours ", numOfHours)

	}

	if numOfMinutes > 0 || str != "" {

		str += fmt.Sprintf("%v-Minutes ", numOfMinutes)

	}

	str += fmt.Sprintf("%v-Seconds ", numOfSeconds)

	str += fmt.Sprintf("%v-Milliseconds ", numOfMilliseconds)

	str += fmt.Sprintf("%v-Microseconds ", numOfMicroseconds)

	str += fmt.Sprintf("%v-Nanoseconds", numOfNanoseconds)

	return str
}
