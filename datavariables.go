package main




const CommentCharStr = "#"
const ZoneLabel = "Zone"
const lenZoneLabel = len(ZoneLabel)
const LinkLabel = "Link"
const lenLinkLabel = len(LinkLabel)

// For IANA Time Zone Files the white space characters which delimit fields
// are space, form feed, carriage return, newline, tab, and  vertical tab.
//
// ' ' = space
// \f = form feed
// \r = carriage return
// \n = new line
// \t = tab
// \v = vertical tab

var FieldSeparators = []rune{
  ' ',
  '\f',
  '\r',
  '\n',
  '\t',
  '\v'}

var fieldSeparatorsLen = len(FieldSeparators)

var tzMajorGroupArray = make([]TimeZoneMajorGroupDto, 0, 100)



var skipTzFiles = []string{
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
  "zone.tab", // Country Zone Tab
  "zone1970.tab", // Country Zone
  "zoneinfo2tdf.pl" }


var filesToDoLast = []string{
  "backward",
  "backzone" }