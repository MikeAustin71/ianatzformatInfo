package tzdatastructs

import (
	"time"
)


var ErrorCount  int
var WarningCount int

var ApplicationStartDateTime time.Time
var ApplicationEndDateTime time.Time


//	    A        Alpha Time Zone                   UTC +1
//	    B        Bravo Time Zone                   UTC +2
//	    C        Charlie Time Zone                 UTC +3
//	    D        Delta Time Zone                   UTC +4
//	    E        Echo Time Zone                    UTC +5
//	    F        Foxtrot Time Zone                 UTC +6
//	    G        Golf Time Zone                    UTC +7
//	    H        Hotel Time Zone                   UTC +8
//	    I        India Time Zone                   UTC +9
//	    K        Kilo Time Zone                    UTC +10
//	    L        Lima Time Zone                    UTC +11
//	    M        Mike Time Zone                    UTC +12
//	    N        November Time Zone                UTC -1
//	    O        Oscar Time Zone                   UTC -2
//	    P        Papa Time Zone                    UTC -3
//	    Q        Quebec Time Zone                  UTC -4
//	    R        Romeo Time Zone                   UTC -5
//	    S        Sierra Time Zone                  UTC -6
//	    T        Tango Time Zone                   UTC -7
//	    U        Uniform Time Zone                 UTC -8
//	    V        Victor Time Zone                  UTC -9
//	    W        Whiskey Time Zone                 UTC -10
//	    X        X-ray Time Zone                   UTC -11
//	    Y        Yankee Time Zone                  UTC -12
//	    Z        Zulu Time Zone                    UTC +0
//
var MilitaryTzMap = map[string]string{
	"Alpha":    "Etc/GMT-1",
	"Bravo":    "Etc/GMT-2",
	"Charlie":  "Etc/GMT-3",
	"Delta":    "Etc/GMT-4",
	"Echo":     "Etc/GMT-5",
	"Foxtrot":  "Etc/GMT-6",
	"Golf":     "Etc/GMT-7",
	"Hotel":    "Etc/GMT-8",
	"India":    "Etc/GMT-9",
	"Kilo":     "Etc/GMT-10",
	"Lima":     "Etc/GMT-11",
	"Mike":     "Etc/GMT-12",
	"November": "Etc/GMT+1",
	"Oscar":    "Etc/GMT+2",
	"Papa":     "Etc/GMT+3",
	"Quebec":   "Etc/GMT+4",
	"Romeo":    "Etc/GMT+5",
	"Sierra":   "Etc/GMT+6",
	"Tango":    "Etc/GMT+7",
	"Uniform":  "Etc/GMT+8",
	"Victor":   "Etc/GMT+9",
	"Whiskey":  "Etc/GMT+10",
	"Xray":     "Etc/GMT+11",
	"Yankee":   "Etc/GMT+12",
	"Zulu":     "UTC"}

var MilitaryUTCMap = map[string]string{
	"Alpha":    "UTC+1",
	"Bravo":    "UTC+2",
	"Charlie":  "UTC+3",
	"Delta":    "UTC+4",
	"Echo":     "UTC+5",
	"Foxtrot":  "UTC+6",
	"Golf":     "UTC+7",
	"Hotel":    "UTC+8",
	"India":    "UTC+9",
	"Kilo":     "UTC+10",
	"Lima":     "UTC+11",
	"Mike":     "UTC+12",
	"November": "UTC-1",
	"Oscar":    "UTC-2",
	"Papa":     "UTC-3",
	"Quebec":   "UTC-4",
	"Romeo":    "UTC-5",
	"Sierra":   "UTC-6",
	"Tango":    "UTC-7",
	"Uniform":  "UTC-8",
	"Victor":   "UTC-9",
	"Whiskey":  "UTC-10",
	"Xray":     "UTC-11",
	"Yankee":   "UTC-12",
	"Zulu":     "UTC+0"}

var MilitaryTzUrlReferenceMap = map[string]string {
	"Alpha"    :  "https://www.timeanddate.com/worldclock/timezone/alpha",
	"Bravo"    :  "https://www.timeanddate.com/worldclock/timezone/bravo",
	"Charlie"  :  "https://www.timeanddate.com/worldclock/timezone/charlie",
	"Delta"    :  "https://www.timeanddate.com/worldclock/timezone/delta",
	"Echo"     :  "https://www.timeanddate.com/worldclock/timezone/echo",
	"Foxtrot"  :  "https://www.timeanddate.com/worldclock/timezone/foxtrot",
	"Golf"     :  "https://www.timeanddate.com/worldclock/timezone/golf",
	"Hotel"    :  "https://www.timeanddate.com/worldclock/timezone/hotel",
	"India"    :  "https://www.timeanddate.com/worldclock/timezone/india",
	"Kilo"     :  "https://www.timeanddate.com/worldclock/timezone/kilo",
	"Lima"     :  "https://www.timeanddate.com/worldclock/timezone/lima",
	"Mike"     :  "https://www.timeanddate.com/worldclock/timezone/mike",
	"November" :  "https://www.timeanddate.com/worldclock/timezone/november",
	"Oscar"    :  "https://www.timeanddate.com/worldclock/timezone/oscar",
	"Papa"     :  "https://www.timeanddate.com/worldclock/timezone/papa",
	"Quebec"   :  "https://www.timeanddate.com/worldclock/timezone/quebec",
	"Romeo"    :  "https://www.timeanddate.com/worldclock/timezone/romeo",
	"Sierra"   :  "https://www.timeanddate.com/worldclock/timezone/sierra",
	"Tango"    :  "https://www.timeanddate.com/worldclock/timezone/tango",
	"Uniform"  :  "https://www.timeanddate.com/worldclock/timezone/uniform",
	"Victor"   :  "https://www.timeanddate.com/worldclock/timezone/victor",
	"Whiskey"  :  "https://www.timeanddate.com/worldclock/timezone/whiskey",
	"Xray"     :  "https://www.timeanddate.com/worldclock/timezone/x-ray",
	"Yankee"   :  "https://www.timeanddate.com/worldclock/timezone/yankee",
	"Zulu"     :  "https://www.timeanddate.com/worldclock/timezone/zulu",
}

var MilitaryTzLocationMap = map[string]string{
	"Alpha"    :  "France",
	"Bravo"    :  "Athens, Greece",
	"Charlie"  :  "Arab Standard Time, Iraq, Bahrain, Kuwait, Saudi Arabia, Yemen, Qatar",
	"Delta"    :  "Moscow, Russia and Afghanistan, however, Afghanistan is technically +4:30 from UTC",
	"Echo"     :  "Pakistan, Kazakhstan, Tajikistan, Uzbekistan and Turkmenistan",
	"Foxtrot"  :  "Bangladesh",
	"Golf"     :  "Thailand",
	"Hotel"    :  "Beijing, China",
	"India"    :  "Tokyo, Australia",
	"Kilo"     :  "Brisbane, Australia",
	"Lima"     :  "Sydney, Australia",
	"Mike"     :  "Wellington, New Zealand",
	"November" :  "Azores",
	"Oscar"    :  "Godthab, Greenland",
	"Papa"     :  "Buenos Aires, Argentina",
	"Quebec"   :  "Halifax, Nova Scotia",
	"Romeo"    :  "EST, New York, NY",
	"Sierra"   :  "CST, Dallas, TX",
	"Tango"    :  "MST, Denver, CO",
	"Uniform"  :  "PST, Los Angeles, CA",
	"Victor"   :  "Juneau, AK",
	"Whiskey"  :  "Honolulu, HI",
	"Xray"     :  "American Samoa",
	"Yankee"   :  "e.g. Fiji",
	"Zulu"     :  "Zulu time",
}
// MilitaryTzArray - Array of strings
// describing official Military Time Zones
var MilitaryTzArray = []string{
	"Alpha",
	"Bravo",
	"Charlie",
	"Delta",
	"Echo",
	"Foxtrot",
	"Golf",
	"Hotel",
	"India",
	"Kilo",
	"Lima",
	"Mike",
	"November",
	"Oscar",
	"Papa",
	"Quebec",
	"Romeo",
	"Sierra",
	"Tango",
	"Uniform",
	"Victor",
	"Whiskey",
	"Xray",
	"Yankee",
	"Zulu",
}

var WorldRegions = []string {
	"Africa",
	"America",
	"Antarctica",
	"Asia",
	"Atlantic",
	"Australia",
	"Europe",
	"Indian",
	"Pacific",
	"Etc",
	"Other" }
