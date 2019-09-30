package tzdatastructs

var timeZoneTypeComments =
  []byte("// TimeZones - This type and its associated methods encapsulate all 590+\n" +
  "// IANA Time Zones plus Military Time Zones. This type is therefore used as an\n" +
  "// enumeration of the Global Time Zones.\n" +
  "//\n" +
  "// The Go Programming Language uses IANA Time Zones in date-time calculations.\n" +
  "//  Reference:\n" +
  "//    https://golang.org/pkg/time/#LoadLocation\n" +
  "//\n" +
  "// IANA Time Zones are widely recognized as the the world's leading authority on\n" +
  "// time zones.\n" +
  "//\n"+
  "// Reference:\n" +
  "//    https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\n" +
  "//    https://en.wikipedia.org/wiki/Tz_database\n" +
  "//\n" +
  "// The IANA Time Zone data base and reference information is located at:\n" +
  "//    https://www.iana.org/time-zones.\n" +
  "//\n" +
  "// For easy access to the IANA Time Zones it is recommended that you use\n" +
  "// the global variable 'TZones' declared below. This variable instantiates the\n" +
  "// 'IanaTimeZones' type. It is therefore much easier to access any of the 590+\n" +
  "// IANA time zones using dot operators and intelliSense (a.k.a. intelligent code\n" +
  "// completion).\n" +
  "//\n" +
  "// Examples:\n" +
  "//   TZones.America.Argentina().Buenos_Aires() - America/Argentina/Buenos_Aires Time Zone\n" +
  "//   TZones.America.Chicago()                  - USA Central Time Zone\n" +
  "//   TZones.America.New_York()                 - USA Eastern Time Zone\n" +
  "//   TZones.America.Denver()                   - USA Mountain Time Zone\n" +
  "//   TZones.America.Los_Angeles()              - USA Pacific Time Zone\n" +
  "//   TZones.Europe.London()                    - Europe/London Time Zone\n" +
  "//   TZones.Europe.Paris()                     - Europe/Paris  Time Zone\n" +
  "//   TZones.Asia.Shanghai()                    - Asia/Shanghai Time Zone\n" +
  "//\n" +
  "// 'TimeZones' has been adapted to function as an enumeration of valid time zone\n" +
  "// values. Since Go does not directly support enumerations, the 'TimeZones' type\n" +
  "// has been configured to function in a manner similar to classic enumerations found\n" +
  "// in other languages like C#.\n" +
  "// For additional information, reference:\n" +
  "//      Jeffrey Richter Using Reflection to implement enumerated types\n" +
  "//             https://www.youtube.com/watch?v=DyXJy_0v0_U \n" +
  "//\n")
