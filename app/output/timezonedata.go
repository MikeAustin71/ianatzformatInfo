package main



// TimeZones - This type and its associated methods encapsulate 0 IANA Time
// Zones, 25-Military Time Zones and 0-Other Time Zones. This type is
// therefore used as a comprehensive enumeration of Global Time Zones.
//
// The Go Programming Language uses IANA Time Zones in date-time calculations.
//  Reference:
//    https://golang.org/pkg/time/
//    https://golang.org/pkg/time/#LoadLocation
//
// The IANA Time Zone database is widely recognized as the the world's leading
// authority on global time zones.
//
// The 'TimeZones' type includes one artificial structure element labeled
// 'Deprecated'. This element encapsulates all of the IANA 'Link' Time Zones.
// 'Link' Time Zones detail those times zones which IANA has classified as
// obsolete and no longer in general use. Each one of these deprecated time
// zones maps to a current, valid IANA time zone.
//
// Reference:
//    https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//    https://en.wikipedia.org/wiki/Tz_database
//
// The IANA Time Zone data base and reference information is located at:
//    https://www.iana.org/time-zones.
//
// For easy access to the all Time Zones it is recommended that you use the
// global variable 'TZones' declared below. This variable instantiates the
// 'TimeZones' type. It is therefore much easier to access any of the 0 time
// zones using dot operators and intellisense (a.k.a. intelligent code completion).
//
// Examples:
// TZones.America.Argentina().Buenos_Aires() - America/Argentina/Buenos_Aires Time Zone
// TZones.America.Chicago()                  - America/Chicago USA Central Time Zone
// TZones.America.New_York()                 - America/New_York USA Eastern Time Zone
// TZones.America.Denver()                   - America/Denver USA Mountain Time Zone
// TZones.America.Los_Angeles()              - America/Los_Angeles USA Pacific Time Zone
// TZones.Europe.London()                    - Europe/London Time Zone
// TZones.Europe.Paris()                     - Europe/Paris  Time Zone
// TZones.Asia.Shanghai()                    - Asia/Shanghai Time Zone
//
// 'TimeZones' has been adapted to function as an enumeration of valid time zone
// values. Since Go does not directly support enumerations, the 'TimeZones' type
// has been configured to function in a manner similar to classic enumerations found
// in other languages like C#. For additional information, reference:
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U 
//
// ----------------------------------------------------------------------------
//                           IANA Time Zones by Region                         
//
//                                         Time     Link    Total
//                                        Zones    Zones    Zones
// --------------------------------------------------------------
// 
// Africa                                     0        0        0
// America                                    0        0        0
// Antarctica                                 0        0        0
// Asia                                       0        0        0
// Atlantic                                   0        0        0
// Australia                                  0        0        0
// Europe                                     0        0        0
// Indian                                     0        0        0
// Pacific                                    0        0        0
// Etc                                        0        0        0
// Other                                     45        0        0
// ==============================================================
//                              Total         0        0        0
//
// ----------------------------------------------------------------------------
// 
// This TimeZones Type is based on IANA Time Zone Database Version: 
// 
//           IANA Standard Time Zones :  45
//           IANA Link Time Zones     :   0
//                                         -------
//                 Sub-Total IANA Time Zones:   0
// 
//                Military Time Zones :  25
//                   Other Time Zones :   0
//                                         -------
//                          Total Time Zones:   0
// 
//       Standard Time Zone Sub-Groups:   0
//           Link Time Zone Sub-Groups:   0
//                                         -------
//                Total Time Zone Sub-Groups:   0
// 
//                  Primary Time Zone Groups:   1
// 
// Type Creation Date: 2019-11-05 Tuesday 17:18:43 -0600 CST
// ----------------------------------------------------------------------------
// 
type TimeZones struct {
    Military                           militaryTimeZones
    Other                              otherTimeZones
}


var TZones = TimeZones{}



// Military - Military Time Zone Names.
//  
// Reference:
//     https://en.wikipedia.org/wiki/List_of_military_time_zones
//     https://www.timeanddate.com/time/zones/military
//     http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//  
// Military time zones are commonly used in aviation as well as at sea.
// They are also known as nautical or maritime time zones.
//  
// The 'J' (Juliet) Time Zone is occasionally used to refer to the observer's
// local time. Note that Time Zone 'J' (Juliet) is not listed below.
//  
//  
//   Abbreviation Time zone name     Other names    Offset
//  
//       A        Alpha Time Zone                   UTC +1
//       B        Bravo Time Zone                   UTC +2
//       C        Charlie Time Zone                 UTC +3
//       D        Delta Time Zone                   UTC +4
//       E        Echo Time Zone                    UTC +5
//       F        Foxtrot Time Zone                 UTC +6
//       G        Golf Time Zone                    UTC +7
//       H        Hotel Time Zone                   UTC +8
//       I        India Time Zone                   UTC +9
//       K        Kilo Time Zone                    UTC +10
//       L        Lima Time Zone                    UTC +11
//       M        Mike Time Zone                    UTC +12
//       N        November Time Zone                UTC -1
//       O        Oscar Time Zone                   UTC -2
//       P        Papa Time Zone                    UTC -3
//       Q        Quebec Time Zone                  UTC -4
//       R        Romeo Time Zone                   UTC -5
//       S        Sierra Time Zone                  UTC -6
//       T        Tango Time Zone                   UTC -7
//       U        Uniform Time Zone                 UTC -8
//       V        Victor Time Zone                  UTC -9
//       W        Whiskey Time Zone                 UTC -10
//       X        X-ray Time Zone                   UTC -11
//       Y        Yankee Time Zone                  UTC -12
//       Z        Zulu Time Zone                    UTC +0
//  
//  
// The methods associated with type 'Military' return the equivalent
// IANA time zones. At first this may seem confusing. For example,
// Military Time Zone 'L' or 'Lima' specifies UTC +11-hours.
// However, the equivalent IANA Time Zone is "Etc/GMT+11".
// In date time calculations, IANA Time Zone "Etc/GMT-11" 
// computes as UTC +11 hours.
//  
//   Reference:
//     https://en.wikipedia.org/wiki/Tz_database#Area
//  
type militaryTimeZones  string

// Alpha - Military Time Zone 'A' or 'Alpha' is equivalent to
// to IANA Time Zone "Etc/GMT+1".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +1 hour.
//  
func (milTz militaryTimeZones) Alpha() string {return "Etc/GMT+1" }

// Bravo - Military Time Zone 'B' or 'Bravo' is equivalent to
// to IANA Time Zone "Etc/GMT+2".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +2 hours.
//  
func (milTz militaryTimeZones) Bravo() string {return "Etc/GMT+2" }

// Charlie - Military Time Zone 'C' or 'Charlie' is equivalent to
// to IANA Time Zone "Etc/GMT+3".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +3 hours.
//  
func (milTz militaryTimeZones) Charlie() string {return "Etc/GMT+3" }

// Delta - Military Time Zone 'D' or 'Delta' is equivalent to
// to IANA Time Zone "Etc/GMT+4".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +4 hours.
//  
func (milTz militaryTimeZones) Delta() string {return "Etc/GMT+4" }

// Echo - Military Time Zone 'E' or 'Echo' is equivalent to
// to IANA Time Zone "Etc/GMT+5".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +5 hours.
//  
func (milTz militaryTimeZones) Echo() string {return "Etc/GMT+5" }

// Foxtrot - Military Time Zone 'F' or 'Foxtrot' is equivalent to
// to IANA Time Zone "Etc/GMT+6".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +6 hours.
//  
func (milTz militaryTimeZones) Foxtrot() string {return "Etc/GMT+6" }

// Golf - Military Time Zone 'G' or 'Golf' is equivalent to
// to IANA Time Zone "Etc/GMT+7".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +7 hours.
//  
func (milTz militaryTimeZones) Golf() string {return "Etc/GMT+7" }

// Hotel - Military Time Zone 'H' or 'Hotel' is equivalent to
// to IANA Time Zone "Etc/GMT+8".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +8 hours.
//  
func (milTz militaryTimeZones) Hotel() string {return "Etc/GMT+8" }

// India - Military Time Zone 'I' or 'India' is equivalent to
// to IANA Time Zone "Etc/GMT+9".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +9 hours.
//  
func (milTz militaryTimeZones) India() string {return "Etc/GMT+9" }

// Kilo - Military Time Zone 'K' or 'Kilo' is equivalent to
// to IANA Time Zone "Etc/GMT+10".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +10 hours.
//  
func (milTz militaryTimeZones) Kilo() string {return "Etc/GMT+10" }

// Lima - Military Time Zone 'L' or 'Lima' is equivalent to
// to IANA Time Zone "Etc/GMT+11".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +11 hours.
//  
func (milTz militaryTimeZones) Lima() string {return "Etc/GMT+11" }

// Mike - Military Time Zone 'M' or 'Mike' is equivalent to
// to IANA Time Zone "Etc/GMT+12".
//  
// Offset from Universal Coordinated Time (UTC) is computed at +12 hours.
//  
func (milTz militaryTimeZones) Mike() string {return "Etc/GMT+12" }

// November - Military Time Zone 'N' or 'November' is equivalent to
// to IANA Time Zone "Etc/GMT-1".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -1 hour.
//  
func (milTz militaryTimeZones) November() string {return "Etc/GMT-1" }

// Oscar - Military Time Zone 'O' or 'Oscar' is equivalent to
// to IANA Time Zone "Etc/GMT-2".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -2 hours.
//  
func (milTz militaryTimeZones) Oscar() string {return "Etc/GMT-2" }

// Papa - Military Time Zone 'P' or 'Papa' is equivalent to
// to IANA Time Zone "Etc/GMT-3".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -3 hours.
//  
func (milTz militaryTimeZones) Papa() string {return "Etc/GMT-3" }

// Quebec - Military Time Zone 'Q' or 'Quebec' is equivalent to
// to IANA Time Zone "Etc/GMT-4".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -4 hours.
//  
func (milTz militaryTimeZones) Quebec() string {return "Etc/GMT-4" }

// Romeo - Military Time Zone 'R' or 'Romeo' is equivalent to
// to IANA Time Zone "Etc/GMT-5".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -5 hours.
//  
func (milTz militaryTimeZones) Romeo() string {return "Etc/GMT-5" }

// Sierra - Military Time Zone 'S' or 'Sierra' is equivalent to
// to IANA Time Zone "Etc/GMT-6".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -6 hours.
//  
func (milTz militaryTimeZones) Sierra() string {return "Etc/GMT-6" }

// Tango - Military Time Zone 'T' or 'Tango' is equivalent to
// to IANA Time Zone "Etc/GMT-7".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -7 hours.
//  
func (milTz militaryTimeZones) Tango() string {return "Etc/GMT-7" }

// Uniform - Military Time Zone 'U' or 'Uniform' is equivalent to
// to IANA Time Zone "Etc/GMT-8".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -8 hours.
//  
func (milTz militaryTimeZones) Uniform() string {return "Etc/GMT-8" }

// Victor - Military Time Zone 'V' or 'Victor' is equivalent to
// to IANA Time Zone "Etc/GMT-9".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -9 hours.
//  
func (milTz militaryTimeZones) Victor() string {return "Etc/GMT-9" }

// Whiskey - Military Time Zone 'W' or 'Whiskey' is equivalent to
// to IANA Time Zone "Etc/GMT-10".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -10 hours.
//  
func (milTz militaryTimeZones) Whiskey() string {return "Etc/GMT-10" }

// Xray - Military Time Zone 'X' or 'Xray' is equivalent to
// to IANA Time Zone "Etc/GMT-11".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -11 hours.
//  
func (milTz militaryTimeZones) Xray() string {return "Etc/GMT-11" }

// Yankee - Military Time Zone 'Y' or 'Yankee' is equivalent to
// to IANA Time Zone "Etc/GMT-12".
//  
// Offset from Universal Coordinated Time (UTC) is computed at -12 hours.
//  
func (milTz militaryTimeZones) Yankee() string {return "Etc/GMT-12" }

// Zulu - Military Time Zone 'Z' or 'Zulu' is equivalent to
// to IANA Time Zone "Etc/UTC".
//  
// Offset from Universal Coordinated Time (UTC) is computed at 0 hours.
//  
func (milTz militaryTimeZones) Zulu() string {return "Etc/UTC" }


// otherTimeZones - IANA Time Zones for 'Other'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type otherTimeZones string

// CET - IANA Time Zone 'Other/CET'.
// IANA Source File: CET
//  
func (other otherTimeZones) CET() string {return "Other/CET" }

// CST6CDT - IANA Time Zone 'Other/CST6CDT'.
// IANA Source File: CST6CDT
//  
func (other otherTimeZones) CST06CDT() string {return "Other/CST6CDT" }

// Cuba - IANA Time Zone 'Other/Cuba'.
// IANA Source File: Cuba
//  
func (other otherTimeZones) Cuba() string {return "Other/Cuba" }

// EET - IANA Time Zone 'Other/EET'.
// IANA Source File: EET
//  
func (other otherTimeZones) EET() string {return "Other/EET" }

// Egypt - IANA Time Zone 'Other/Egypt'.
// IANA Source File: Egypt
//  
func (other otherTimeZones) Egypt() string {return "Other/Egypt" }

// Eire - IANA Time Zone 'Other/Eire'.
// IANA Source File: Eire
//  
func (other otherTimeZones) Eire() string {return "Other/Eire" }

// EST - IANA Time Zone 'Other/EST'.
// IANA Source File: EST
//  
func (other otherTimeZones) EST() string {return "Other/EST" }

// EST5EDT - IANA Time Zone 'Other/EST5EDT'.
// IANA Source File: EST5EDT
//  
func (other otherTimeZones) EST05EDT() string {return "Other/EST5EDT" }

// Factory - IANA Time Zone 'Other/Factory'.
// IANA Source File: Factory
//  
func (other otherTimeZones) Factory() string {return "Other/Factory" }

// GB - IANA Time Zone 'Other/GB'.
// IANA Source File: GB
//  
func (other otherTimeZones) GB() string {return "Other/GB" }

// GB-Eire - IANA Time Zone 'Other/GB-Eire'.
// IANA Source File: GB-Eire
//  
func (other otherTimeZones) GBMinusEire() string {return "Other/GB-Eire" }

// GMT - IANA Time Zone 'Other/GMT'.
// IANA Source File: GMT
//  
func (other otherTimeZones) GMT() string {return "Other/GMT" }

// GMT+0 - IANA Time Zone 'Other/GMT+0'.
// IANA Source File: GMT+0
//  
func (other otherTimeZones) GMTPlus00() string {return "Other/GMT+0" }

// GMT-0 - IANA Time Zone 'Other/GMT-0'.
// IANA Source File: GMT-0
//  
func (other otherTimeZones) GMTMinus00() string {return "Other/GMT-0" }

// GMT0 - IANA Time Zone 'Other/GMT0'.
// IANA Source File: GMT0
//  
func (other otherTimeZones) GMT00() string {return "Other/GMT0" }

// Greenwich - IANA Time Zone 'Other/Greenwich'.
// IANA Source File: Greenwich
//  
func (other otherTimeZones) Greenwich() string {return "Other/Greenwich" }

// Hongkong - IANA Time Zone 'Other/Hongkong'.
// IANA Source File: Hongkong
//  
func (other otherTimeZones) Hongkong() string {return "Other/Hongkong" }

// HST - IANA Time Zone 'Other/HST'.
// IANA Source File: HST
//  
func (other otherTimeZones) HST() string {return "Other/HST" }

// Iceland - IANA Time Zone 'Other/Iceland'.
// IANA Source File: Iceland
//  
func (other otherTimeZones) Iceland() string {return "Other/Iceland" }

// Iran - IANA Time Zone 'Other/Iran'.
// IANA Source File: Iran
//  
func (other otherTimeZones) Iran() string {return "Other/Iran" }

// Israel - IANA Time Zone 'Other/Israel'.
// IANA Source File: Israel
//  
func (other otherTimeZones) Israel() string {return "Other/Israel" }

// Jamaica - IANA Time Zone 'Other/Jamaica'.
// IANA Source File: Jamaica
//  
func (other otherTimeZones) Jamaica() string {return "Other/Jamaica" }

// Japan - IANA Time Zone 'Other/Japan'.
// IANA Source File: Japan
//  
func (other otherTimeZones) Japan() string {return "Other/Japan" }

// Kwajalein - IANA Time Zone 'Other/Kwajalein'.
// IANA Source File: Kwajalein
//  
func (other otherTimeZones) Kwajalein() string {return "Other/Kwajalein" }

// Libya - IANA Time Zone 'Other/Libya'.
// IANA Source File: Libya
//  
func (other otherTimeZones) Libya() string {return "Other/Libya" }

// MET - IANA Time Zone 'Other/MET'.
// IANA Source File: MET
//  
func (other otherTimeZones) MET() string {return "Other/MET" }

// MST - IANA Time Zone 'Other/MST'.
// IANA Source File: MST
//  
func (other otherTimeZones) MST() string {return "Other/MST" }

// MST7MDT - IANA Time Zone 'Other/MST7MDT'.
// IANA Source File: MST7MDT
//  
func (other otherTimeZones) MST07MDT() string {return "Other/MST7MDT" }

// Navajo - IANA Time Zone 'Other/Navajo'.
// IANA Source File: Navajo
//  
func (other otherTimeZones) Navajo() string {return "Other/Navajo" }

// NZ - IANA Time Zone 'Other/NZ'.
// IANA Source File: NZ
//  
func (other otherTimeZones) NZ() string {return "Other/NZ" }

// NZ-CHAT - IANA Time Zone 'Other/NZ-CHAT'.
// IANA Source File: NZ-CHAT
//  
func (other otherTimeZones) NZMinusCHAT() string {return "Other/NZ-CHAT" }

// Poland - IANA Time Zone 'Other/Poland'.
// IANA Source File: Poland
//  
func (other otherTimeZones) Poland() string {return "Other/Poland" }

// Portugal - IANA Time Zone 'Other/Portugal'.
// IANA Source File: Portugal
//  
func (other otherTimeZones) Portugal() string {return "Other/Portugal" }

// PRC - IANA Time Zone 'Other/PRC'.
// IANA Source File: PRC
//  
func (other otherTimeZones) PRC() string {return "Other/PRC" }

// PST8PDT - IANA Time Zone 'Other/PST8PDT'.
// IANA Source File: PST8PDT
//  
func (other otherTimeZones) PST08PDT() string {return "Other/PST8PDT" }

// ROC - IANA Time Zone 'Other/ROC'.
// IANA Source File: ROC
//  
func (other otherTimeZones) ROC() string {return "Other/ROC" }

// ROK - IANA Time Zone 'Other/ROK'.
// IANA Source File: ROK
//  
func (other otherTimeZones) ROK() string {return "Other/ROK" }

// Singapore - IANA Time Zone 'Other/Singapore'.
// IANA Source File: Singapore
//  
func (other otherTimeZones) Singapore() string {return "Other/Singapore" }

// Turkey - IANA Time Zone 'Other/Turkey'.
// IANA Source File: Turkey
//  
func (other otherTimeZones) Turkey() string {return "Other/Turkey" }

// UCT - IANA Time Zone 'Other/UCT'.
// IANA Source File: UCT
//  
func (other otherTimeZones) UCT() string {return "Other/UCT" }

// Universal - IANA Time Zone 'Other/Universal'.
// IANA Source File: Universal
//  
func (other otherTimeZones) Universal() string {return "Other/Universal" }

// UTC - IANA Time Zone 'Other/UTC'.
// IANA Source File: UTC
//  
func (other otherTimeZones) UTC() string {return "Other/UTC" }

// W-SU - IANA Time Zone 'Other/W-SU'.
// IANA Source File: W-SU
//  
func (other otherTimeZones) WMinusSU() string {return "Other/W-SU" }

// WET - IANA Time Zone 'Other/WET'.
// IANA Source File: WET
//  
func (other otherTimeZones) WET() string {return "Other/WET" }

// Zulu - IANA Time Zone 'Other/Zulu'.
// IANA Source File: Zulu
//  
func (other otherTimeZones) Zulu() string {return "Other/Zulu" }

