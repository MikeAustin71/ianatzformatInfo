package tzdatastructs

import 	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"


var TimeZoneGroups [] TimeZoneGroupCollection
var TimeZones [] TimeZoneDataCollection
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
