package tzdatastructs

import 	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"


var CurWorkingDirectory pathfileops.DirMgr


var DEBUG = 0


// For IANA Time Zone Files the white space characters which delimit fields
// are space, form feed, carriage return, newline, tab, and  vertical tab.
//
// ' ' = space
// \f = form feed
// \r = carriage return
// \n = new line
// \t = tab
// \v = vertical tab

var LeadingFieldSeparators = []string {
	" ",
	"\t",
	"\v"}


var TrailingFieldSeparators = []string {
	" ",
	"\t",
	"\v"}


var EndOfLineDelimiters = []string{
	"\n",  // new line 0x0A - 10
	"\r",  // carriage return 0x0C - 12
	"\f" } // form feed 0x0D - 13

var CommentDelimiters = []string {CommentCharStr}

var SkipTzFiles = []string{
	"checklinks.awk",
	"checktab.awk",
	"calendars",
	"CONTRIBUTING",
	"date.1",
	"date.1.txt",
	"difftime.c",
	"factory",
	"iso3166.tab",  // Two digit Country Codes
	"leapseconds",
	"leapseconds.awk",
	"leap-seconds.list",
	"LICENSE",
	"localtime.c",
	"Makefile",
	"newctime.3",
	"newctime.3.txt",
	"NEWS",
	"newstrftime.3",
	"newstrftime.3.txt",
	"newtzset.3.txt",
	"private.h",
	"README",
	"strftime.c",
	"systemv",
	"theory.html",
	"time2posix.3",
	"time2posix.3.txt",
	"to2050.tzs",       // Contains a lot of valid links
	"tz-art.html",
	"tzdata.zi",
	"tzfile.5",
	"tzfile.5.txt",
	"tzfile.h",
	"tz-how-to.html",
	"tz-link.html",
	"tzselect.8",
	"tzselect.8.txt",
	"tzselect.ksh",
	"version",
	"workman.sh",
	"yearistype.sh",
	"zdump.8",
	"zdump.8.txt",
	"zdump.c",
	"zic.8",
	"zic.8.txt",
	"zic.c",
	"ziguard.awk",
	"zishrink.awk",
	"zone.tab", // Country Zone Tab Deprecated
	"zone1970.tab", // Country Zone New
	"zoneinfo2tdf.pl" }


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
	"Alpha":    "Etc/GMT+1",
	"Bravo":    "Etc/GMT+2",
	"Charlie":  "Etc/GMT+3",
	"Delta":    "Etc/GMT+4",
	"Echo":     "Etc/GMT+5",
	"Foxtrot":  "Etc/GMT+6",
	"Golf":     "Etc/GMT+7",
	"Hotel":    "Etc/GMT+8",
	"India":    "Etc/GMT+9",
	"Kilo":     "Etc/GMT+10",
	"Lima":     "Etc/GMT+11",
	"Mike":     "Etc/GMT+12",
	"November": "Etc/GMT-1",
	"Oscar":    "Etc/GMT-2",
	"Papa":     "Etc/GMT-3",
	"Quebec":   "Etc/GMT-4",
	"Romeo":    "Etc/GMT-5",
	"Sierra":   "Etc/GMT-6",
	"Tango":    "Etc/GMT-7",
	"Uniform":  "Etc/GMT-8",
	"Victor":   "Etc/GMT-9",
	"Whiskey":  "Etc/GMT-10",
	"Xray":     "Etc/GMT-11",
	"Yankee":   "Etc/GMT-12",
	"Zulu":     "Etc/UTC"}


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


