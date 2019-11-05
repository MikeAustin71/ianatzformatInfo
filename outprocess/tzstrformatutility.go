package outprocess

import (
	"fmt"
	"time"
)


type TzStrFmt struct {
	InputStr     string
	OutputStr    string
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
