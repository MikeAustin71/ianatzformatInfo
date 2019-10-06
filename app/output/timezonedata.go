package main





// TimeZones - This type and its associated methods encapsulate 663 IANA Time
// Zones plus 25-Military Time Zones. This type is therefore used as a 
// comprehensive enumeration of Global Time Zones.
//
// The Go Programming Language uses IANA Time Zones in date-time calculations.
//  Reference:
//    https://golang.org/pkg/time/#LoadLocation
//
// IANA Time Zones are widely recognized as the the world's leading authority
// on time zones.
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
// 'TimeZones' type. It is therefore much easier to access any of the 688 time
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
// 
// This TimeZones Type is based on IANA Time Zone Database Version: 2019c
// 
//           IANA Standard Time Zones : 457
//           IANA Link Time Zones     : 206
//                                         -------
//                 Sub-Total IANA Time Zones: 663
// 
//                Military Time Zones :  25
//                   Other Time Zones :   0
//                                         -------
//                          Total Time Zones: 688
// 
//       Standard Time Zone Sub-Groups:   4
//           Link Time Zone Sub-Groups:  16
//                                         -------
//                Total Time Zone Sub-Groups:  20
// 
//                  Primary Time Zone Groups:  12
// 
// ----------------------------------------------------------------------------
// 
type TimeZones struct {
     Africa                             africaTimeZones
     Atlantic                           atlanticTimeZones
     Deprecated                         deprecatedTimeZones
     Indian                             indianTimeZones
     Antarctica                         antarcticaTimeZones
     Asia                               asiaTimeZones
     Australia                          australiaTimeZones
     Pacific                            pacificTimeZones
     America                            americaTimeZones
     Europe                             europeTimeZones
     Etc                                etcTimeZones
     Military                           militaryTimeZones
}


var TZones = TimeZones{}


// africaTimeZones - IANA Time Zones for 'Africa'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type africaTimeZones string

// Abidjan - IANA Time Zone 'Africa/Abidjan'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Abidjan() string {return "Africa/Abidjan" }

// Accra - IANA Time Zone 'Africa/Accra'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Accra() string {return "Africa/Accra" }

// Addis_Ababa - IANA Time Zone 'Africa/Addis_Ababa'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Addis_Ababa() string {return "Africa/Addis_Ababa" }

// Algiers - IANA Time Zone 'Africa/Algiers'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Algiers() string {return "Africa/Algiers" }

// Asmara - IANA Time Zone 'Africa/Asmara'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Asmara() string {return "Africa/Asmara" }

// Bamako - IANA Time Zone 'Africa/Bamako'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Bamako() string {return "Africa/Bamako" }

// Bangui - IANA Time Zone 'Africa/Bangui'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Bangui() string {return "Africa/Bangui" }

// Banjul - IANA Time Zone 'Africa/Banjul'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Banjul() string {return "Africa/Banjul" }

// Bissau - IANA Time Zone 'Africa/Bissau'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Bissau() string {return "Africa/Bissau" }

// Blantyre - IANA Time Zone 'Africa/Blantyre'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Blantyre() string {return "Africa/Blantyre" }

// Brazzaville - IANA Time Zone 'Africa/Brazzaville'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Brazzaville() string {return "Africa/Brazzaville" }

// Bujumbura - IANA Time Zone 'Africa/Bujumbura'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Bujumbura() string {return "Africa/Bujumbura" }

// Cairo - IANA Time Zone 'Africa/Cairo'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Cairo() string {return "Africa/Cairo" }

// Casablanca - IANA Time Zone 'Africa/Casablanca'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Casablanca() string {return "Africa/Casablanca" }

// Ceuta - IANA Time Zone 'Africa/Ceuta'.
// IANA Source File: europe
//  
func (afric africaTimeZones) Ceuta() string {return "Africa/Ceuta" }

// Conakry - IANA Time Zone 'Africa/Conakry'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Conakry() string {return "Africa/Conakry" }

// Dakar - IANA Time Zone 'Africa/Dakar'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Dakar() string {return "Africa/Dakar" }

// Dar_es_Salaam - IANA Time Zone 'Africa/Dar_es_Salaam'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Dar_es_Salaam() string {return "Africa/Dar_es_Salaam" }

// Djibouti - IANA Time Zone 'Africa/Djibouti'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Djibouti() string {return "Africa/Djibouti" }

// Douala - IANA Time Zone 'Africa/Douala'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Douala() string {return "Africa/Douala" }

// El_Aaiun - IANA Time Zone 'Africa/El_Aaiun'.
// IANA Source File: africa
//  
func (afric africaTimeZones) El_Aaiun() string {return "Africa/El_Aaiun" }

// Freetown - IANA Time Zone 'Africa/Freetown'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Freetown() string {return "Africa/Freetown" }

// Gaborone - IANA Time Zone 'Africa/Gaborone'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Gaborone() string {return "Africa/Gaborone" }

// Harare - IANA Time Zone 'Africa/Harare'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Harare() string {return "Africa/Harare" }

// Johannesburg - IANA Time Zone 'Africa/Johannesburg'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Johannesburg() string {return "Africa/Johannesburg" }

// Juba - IANA Time Zone 'Africa/Juba'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Juba() string {return "Africa/Juba" }

// Kampala - IANA Time Zone 'Africa/Kampala'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Kampala() string {return "Africa/Kampala" }

// Khartoum - IANA Time Zone 'Africa/Khartoum'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Khartoum() string {return "Africa/Khartoum" }

// Kigali - IANA Time Zone 'Africa/Kigali'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Kigali() string {return "Africa/Kigali" }

// Kinshasa - IANA Time Zone 'Africa/Kinshasa'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Kinshasa() string {return "Africa/Kinshasa" }

// Lagos - IANA Time Zone 'Africa/Lagos'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Lagos() string {return "Africa/Lagos" }

// Libreville - IANA Time Zone 'Africa/Libreville'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Libreville() string {return "Africa/Libreville" }

// Lome - IANA Time Zone 'Africa/Lome'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Lome() string {return "Africa/Lome" }

// Luanda - IANA Time Zone 'Africa/Luanda'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Luanda() string {return "Africa/Luanda" }

// Lubumbashi - IANA Time Zone 'Africa/Lubumbashi'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Lubumbashi() string {return "Africa/Lubumbashi" }

// Lusaka - IANA Time Zone 'Africa/Lusaka'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Lusaka() string {return "Africa/Lusaka" }

// Malabo - IANA Time Zone 'Africa/Malabo'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Malabo() string {return "Africa/Malabo" }

// Maputo - IANA Time Zone 'Africa/Maputo'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Maputo() string {return "Africa/Maputo" }

// Maseru - IANA Time Zone 'Africa/Maseru'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Maseru() string {return "Africa/Maseru" }

// Mbabane - IANA Time Zone 'Africa/Mbabane'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Mbabane() string {return "Africa/Mbabane" }

// Mogadishu - IANA Time Zone 'Africa/Mogadishu'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Mogadishu() string {return "Africa/Mogadishu" }

// Monrovia - IANA Time Zone 'Africa/Monrovia'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Monrovia() string {return "Africa/Monrovia" }

// Nairobi - IANA Time Zone 'Africa/Nairobi'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Nairobi() string {return "Africa/Nairobi" }

// Ndjamena - IANA Time Zone 'Africa/Ndjamena'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Ndjamena() string {return "Africa/Ndjamena" }

// Niamey - IANA Time Zone 'Africa/Niamey'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Niamey() string {return "Africa/Niamey" }

// Nouakchott - IANA Time Zone 'Africa/Nouakchott'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Nouakchott() string {return "Africa/Nouakchott" }

// Ouagadougou - IANA Time Zone 'Africa/Ouagadougou'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Ouagadougou() string {return "Africa/Ouagadougou" }

// Porto-Novo - IANA Time Zone 'Africa/Porto-Novo'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) PortoMinusNovo() string {return "Africa/Porto-Novo" }

// Sao_Tome - IANA Time Zone 'Africa/Sao_Tome'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Sao_Tome() string {return "Africa/Sao_Tome" }

// Timbuktu - IANA Time Zone 'Africa/Timbuktu'.
// IANA Source File: backzone
//  
func (afric africaTimeZones) Timbuktu() string {return "Africa/Timbuktu" }

// Tripoli - IANA Time Zone 'Africa/Tripoli'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Tripoli() string {return "Africa/Tripoli" }

// Tunis - IANA Time Zone 'Africa/Tunis'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Tunis() string {return "Africa/Tunis" }

// Windhoek - IANA Time Zone 'Africa/Windhoek'.
// IANA Source File: africa
//  
func (afric africaTimeZones) Windhoek() string {return "Africa/Windhoek" }

// americaTimeZones - IANA Time Zones for 'America'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type americaTimeZones string

// Adak - IANA Time Zone 'America/Adak'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Adak() string {return "America/Adak" }

// Anchorage - IANA Time Zone 'America/Anchorage'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Anchorage() string {return "America/Anchorage" }

// Anguilla - IANA Time Zone 'America/Anguilla'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Anguilla() string {return "America/Anguilla" }

// Antigua - IANA Time Zone 'America/Antigua'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Antigua() string {return "America/Antigua" }

// Araguaina - IANA Time Zone 'America/Araguaina'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Araguaina() string {return "America/Araguaina" }

// Aruba - IANA Time Zone 'America/Aruba'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Aruba() string {return "America/Aruba" }

// Asuncion - IANA Time Zone 'America/Asuncion'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Asuncion() string {return "America/Asuncion" }

// Atikokan - IANA Time Zone 'America/Atikokan'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Atikokan() string {return "America/Atikokan" }

// Bahia - IANA Time Zone 'America/Bahia'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Bahia() string {return "America/Bahia" }

// Bahia_Banderas - IANA Time Zone 'America/Bahia_Banderas'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Bahia_Banderas() string {return "America/Bahia_Banderas" }

// Barbados - IANA Time Zone 'America/Barbados'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Barbados() string {return "America/Barbados" }

// Belem - IANA Time Zone 'America/Belem'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Belem() string {return "America/Belem" }

// Belize - IANA Time Zone 'America/Belize'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Belize() string {return "America/Belize" }

// Blanc-Sablon - IANA Time Zone 'America/Blanc-Sablon'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) BlancMinusSablon() string {return "America/Blanc-Sablon" }

// Boa_Vista - IANA Time Zone 'America/Boa_Vista'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Boa_Vista() string {return "America/Boa_Vista" }

// Bogota - IANA Time Zone 'America/Bogota'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Bogota() string {return "America/Bogota" }

// Boise - IANA Time Zone 'America/Boise'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Boise() string {return "America/Boise" }

// Cambridge_Bay - IANA Time Zone 'America/Cambridge_Bay'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Cambridge_Bay() string {return "America/Cambridge_Bay" }

// Campo_Grande - IANA Time Zone 'America/Campo_Grande'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Campo_Grande() string {return "America/Campo_Grande" }

// Cancun - IANA Time Zone 'America/Cancun'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Cancun() string {return "America/Cancun" }

// Caracas - IANA Time Zone 'America/Caracas'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Caracas() string {return "America/Caracas" }

// Cayenne - IANA Time Zone 'America/Cayenne'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Cayenne() string {return "America/Cayenne" }

// Cayman - IANA Time Zone 'America/Cayman'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Cayman() string {return "America/Cayman" }

// Chicago - IANA Time Zone 'America/Chicago'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Chicago() string {return "America/Chicago" }

// Chihuahua - IANA Time Zone 'America/Chihuahua'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Chihuahua() string {return "America/Chihuahua" }

// Coral_Harbour - IANA Time Zone 'America/Coral_Harbour'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Coral_Harbour() string {return "America/Coral_Harbour" }

// Costa_Rica - IANA Time Zone 'America/Costa_Rica'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Costa_Rica() string {return "America/Costa_Rica" }

// Creston - IANA Time Zone 'America/Creston'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Creston() string {return "America/Creston" }

// Cuiaba - IANA Time Zone 'America/Cuiaba'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Cuiaba() string {return "America/Cuiaba" }

// Curacao - IANA Time Zone 'America/Curacao'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Curacao() string {return "America/Curacao" }

// Danmarkshavn - IANA Time Zone 'America/Danmarkshavn'.
// IANA Source File: europe
//  
func (ameri americaTimeZones) Danmarkshavn() string {return "America/Danmarkshavn" }

// Dawson - IANA Time Zone 'America/Dawson'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Dawson() string {return "America/Dawson" }

// Dawson_Creek - IANA Time Zone 'America/Dawson_Creek'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Dawson_Creek() string {return "America/Dawson_Creek" }

// Denver - IANA Time Zone 'America/Denver'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Denver() string {return "America/Denver" }

// Detroit - IANA Time Zone 'America/Detroit'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Detroit() string {return "America/Detroit" }

// Dominica - IANA Time Zone 'America/Dominica'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Dominica() string {return "America/Dominica" }

// Edmonton - IANA Time Zone 'America/Edmonton'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Edmonton() string {return "America/Edmonton" }

// Eirunepe - IANA Time Zone 'America/Eirunepe'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Eirunepe() string {return "America/Eirunepe" }

// El_Salvador - IANA Time Zone 'America/El_Salvador'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) El_Salvador() string {return "America/El_Salvador" }

// Ensenada - IANA Time Zone 'America/Ensenada'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Ensenada() string {return "America/Ensenada" }

// Fort_Nelson - IANA Time Zone 'America/Fort_Nelson'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Fort_Nelson() string {return "America/Fort_Nelson" }

// Fortaleza - IANA Time Zone 'America/Fortaleza'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Fortaleza() string {return "America/Fortaleza" }

// Glace_Bay - IANA Time Zone 'America/Glace_Bay'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Glace_Bay() string {return "America/Glace_Bay" }

// Godthab - IANA Time Zone 'America/Godthab'.
// IANA Source File: europe
//  
func (ameri americaTimeZones) Godthab() string {return "America/Godthab" }

// Goose_Bay - IANA Time Zone 'America/Goose_Bay'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Goose_Bay() string {return "America/Goose_Bay" }

// Grand_Turk - IANA Time Zone 'America/Grand_Turk'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Grand_Turk() string {return "America/Grand_Turk" }

// Grenada - IANA Time Zone 'America/Grenada'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Grenada() string {return "America/Grenada" }

// Guadeloupe - IANA Time Zone 'America/Guadeloupe'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Guadeloupe() string {return "America/Guadeloupe" }

// Guatemala - IANA Time Zone 'America/Guatemala'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Guatemala() string {return "America/Guatemala" }

// Guayaquil - IANA Time Zone 'America/Guayaquil'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Guayaquil() string {return "America/Guayaquil" }

// Guyana - IANA Time Zone 'America/Guyana'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Guyana() string {return "America/Guyana" }

// Halifax - IANA Time Zone 'America/Halifax'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Halifax() string {return "America/Halifax" }

// Havana - IANA Time Zone 'America/Havana'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Havana() string {return "America/Havana" }

// Hermosillo - IANA Time Zone 'America/Hermosillo'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Hermosillo() string {return "America/Hermosillo" }

// Inuvik - IANA Time Zone 'America/Inuvik'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Inuvik() string {return "America/Inuvik" }

// Iqaluit - IANA Time Zone 'America/Iqaluit'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Iqaluit() string {return "America/Iqaluit" }

// Jamaica - IANA Time Zone 'America/Jamaica'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Jamaica() string {return "America/Jamaica" }

// Juneau - IANA Time Zone 'America/Juneau'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Juneau() string {return "America/Juneau" }

// La_Paz - IANA Time Zone 'America/La_Paz'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) La_Paz() string {return "America/La_Paz" }

// Lima - IANA Time Zone 'America/Lima'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Lima() string {return "America/Lima" }

// Los_Angeles - IANA Time Zone 'America/Los_Angeles'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Los_Angeles() string {return "America/Los_Angeles" }

// Maceio - IANA Time Zone 'America/Maceio'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Maceio() string {return "America/Maceio" }

// Managua - IANA Time Zone 'America/Managua'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Managua() string {return "America/Managua" }

// Manaus - IANA Time Zone 'America/Manaus'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Manaus() string {return "America/Manaus" }

// Martinique - IANA Time Zone 'America/Martinique'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Martinique() string {return "America/Martinique" }

// Matamoros - IANA Time Zone 'America/Matamoros'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Matamoros() string {return "America/Matamoros" }

// Mazatlan - IANA Time Zone 'America/Mazatlan'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Mazatlan() string {return "America/Mazatlan" }

// Menominee - IANA Time Zone 'America/Menominee'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Menominee() string {return "America/Menominee" }

// Merida - IANA Time Zone 'America/Merida'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Merida() string {return "America/Merida" }

// Metlakatla - IANA Time Zone 'America/Metlakatla'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Metlakatla() string {return "America/Metlakatla" }

// Mexico_City - IANA Time Zone 'America/Mexico_City'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Mexico_City() string {return "America/Mexico_City" }

// Miquelon - IANA Time Zone 'America/Miquelon'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Miquelon() string {return "America/Miquelon" }

// Moncton - IANA Time Zone 'America/Moncton'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Moncton() string {return "America/Moncton" }

// Monterrey - IANA Time Zone 'America/Monterrey'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Monterrey() string {return "America/Monterrey" }

// Montevideo - IANA Time Zone 'America/Montevideo'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Montevideo() string {return "America/Montevideo" }

// Montreal - IANA Time Zone 'America/Montreal'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Montreal() string {return "America/Montreal" }

// Montserrat - IANA Time Zone 'America/Montserrat'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Montserrat() string {return "America/Montserrat" }

// Nassau - IANA Time Zone 'America/Nassau'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Nassau() string {return "America/Nassau" }

// New_York - IANA Time Zone 'America/New_York'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) New_York() string {return "America/New_York" }

// Nipigon - IANA Time Zone 'America/Nipigon'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Nipigon() string {return "America/Nipigon" }

// Nome - IANA Time Zone 'America/Nome'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Nome() string {return "America/Nome" }

// Noronha - IANA Time Zone 'America/Noronha'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Noronha() string {return "America/Noronha" }

// Ojinaga - IANA Time Zone 'America/Ojinaga'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Ojinaga() string {return "America/Ojinaga" }

// Panama - IANA Time Zone 'America/Panama'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Panama() string {return "America/Panama" }

// Pangnirtung - IANA Time Zone 'America/Pangnirtung'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Pangnirtung() string {return "America/Pangnirtung" }

// Paramaribo - IANA Time Zone 'America/Paramaribo'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Paramaribo() string {return "America/Paramaribo" }

// Phoenix - IANA Time Zone 'America/Phoenix'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Phoenix() string {return "America/Phoenix" }

// Port-au-Prince - IANA Time Zone 'America/Port-au-Prince'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) PortMinusauMinusPrince() string {return "America/Port-au-Prince" }

// Port_of_Spain - IANA Time Zone 'America/Port_of_Spain'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Port_of_Spain() string {return "America/Port_of_Spain" }

// Porto_Velho - IANA Time Zone 'America/Porto_Velho'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Porto_Velho() string {return "America/Porto_Velho" }

// Puerto_Rico - IANA Time Zone 'America/Puerto_Rico'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Puerto_Rico() string {return "America/Puerto_Rico" }

// Punta_Arenas - IANA Time Zone 'America/Punta_Arenas'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Punta_Arenas() string {return "America/Punta_Arenas" }

// Rainy_River - IANA Time Zone 'America/Rainy_River'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Rainy_River() string {return "America/Rainy_River" }

// Rankin_Inlet - IANA Time Zone 'America/Rankin_Inlet'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Rankin_Inlet() string {return "America/Rankin_Inlet" }

// Recife - IANA Time Zone 'America/Recife'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Recife() string {return "America/Recife" }

// Regina - IANA Time Zone 'America/Regina'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Regina() string {return "America/Regina" }

// Resolute - IANA Time Zone 'America/Resolute'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Resolute() string {return "America/Resolute" }

// Rio_Branco - IANA Time Zone 'America/Rio_Branco'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Rio_Branco() string {return "America/Rio_Branco" }

// Rosario - IANA Time Zone 'America/Rosario'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Rosario() string {return "America/Rosario" }

// Santarem - IANA Time Zone 'America/Santarem'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Santarem() string {return "America/Santarem" }

// Santiago - IANA Time Zone 'America/Santiago'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Santiago() string {return "America/Santiago" }

// Santo_Domingo - IANA Time Zone 'America/Santo_Domingo'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Santo_Domingo() string {return "America/Santo_Domingo" }

// Sao_Paulo - IANA Time Zone 'America/Sao_Paulo'.
// IANA Source File: southamerica
//  
func (ameri americaTimeZones) Sao_Paulo() string {return "America/Sao_Paulo" }

// Scoresbysund - IANA Time Zone 'America/Scoresbysund'.
// IANA Source File: europe
//  
func (ameri americaTimeZones) Scoresbysund() string {return "America/Scoresbysund" }

// Sitka - IANA Time Zone 'America/Sitka'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Sitka() string {return "America/Sitka" }

// St_Johns - IANA Time Zone 'America/St_Johns'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) St_Johns() string {return "America/St_Johns" }

// St_Kitts - IANA Time Zone 'America/St_Kitts'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) St_Kitts() string {return "America/St_Kitts" }

// St_Lucia - IANA Time Zone 'America/St_Lucia'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) St_Lucia() string {return "America/St_Lucia" }

// St_Thomas - IANA Time Zone 'America/St_Thomas'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) St_Thomas() string {return "America/St_Thomas" }

// St_Vincent - IANA Time Zone 'America/St_Vincent'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) St_Vincent() string {return "America/St_Vincent" }

// Swift_Current - IANA Time Zone 'America/Swift_Current'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Swift_Current() string {return "America/Swift_Current" }

// Tegucigalpa - IANA Time Zone 'America/Tegucigalpa'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Tegucigalpa() string {return "America/Tegucigalpa" }

// Thule - IANA Time Zone 'America/Thule'.
// IANA Source File: europe
//  
func (ameri americaTimeZones) Thule() string {return "America/Thule" }

// Thunder_Bay - IANA Time Zone 'America/Thunder_Bay'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Thunder_Bay() string {return "America/Thunder_Bay" }

// Tijuana - IANA Time Zone 'America/Tijuana'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Tijuana() string {return "America/Tijuana" }

// Toronto - IANA Time Zone 'America/Toronto'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Toronto() string {return "America/Toronto" }

// Tortola - IANA Time Zone 'America/Tortola'.
// IANA Source File: backzone
//  
func (ameri americaTimeZones) Tortola() string {return "America/Tortola" }

// Vancouver - IANA Time Zone 'America/Vancouver'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Vancouver() string {return "America/Vancouver" }

// Whitehorse - IANA Time Zone 'America/Whitehorse'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Whitehorse() string {return "America/Whitehorse" }

// Winnipeg - IANA Time Zone 'America/Winnipeg'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Winnipeg() string {return "America/Winnipeg" }

// Yakutat - IANA Time Zone 'America/Yakutat'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Yakutat() string {return "America/Yakutat" }

// Yellowknife - IANA Time Zone 'America/Yellowknife'.
// IANA Source File: northamerica
//  
func (ameri americaTimeZones) Yellowknife() string {return "America/Yellowknife" }

// Argentina - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (ameri americaTimeZones) Argentina() argentinaTimeZones {return "" }

// Indiana - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (ameri americaTimeZones) Indiana() indianaTimeZones {return "" }

// Kentucky - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (ameri americaTimeZones) Kentucky() kentuckyTimeZones {return "" }

// North_Dakota - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (ameri americaTimeZones) North_Dakota() north_DakotaTimeZones {return "" }

// antarcticaTimeZones - IANA Time Zones for 'Antarctica'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type antarcticaTimeZones string

// Casey - IANA Time Zone 'Antarctica/Casey'.
// IANA Source File: antarctica
//  
func (antar antarcticaTimeZones) Casey() string {return "Antarctica/Casey" }

// Davis - IANA Time Zone 'Antarctica/Davis'.
// IANA Source File: antarctica
//  
func (antar antarcticaTimeZones) Davis() string {return "Antarctica/Davis" }

// DumontDUrville - IANA Time Zone 'Antarctica/DumontDUrville'.
// IANA Source File: antarctica
//  
func (antar antarcticaTimeZones) DumontDUrville() string {return "Antarctica/DumontDUrville" }

// Macquarie - IANA Time Zone 'Antarctica/Macquarie'.
// IANA Source File: australasia
//  
func (antar antarcticaTimeZones) Macquarie() string {return "Antarctica/Macquarie" }

// Mawson - IANA Time Zone 'Antarctica/Mawson'.
// IANA Source File: antarctica
//  
func (antar antarcticaTimeZones) Mawson() string {return "Antarctica/Mawson" }

// McMurdo - IANA Time Zone 'Antarctica/McMurdo'.
// IANA Source File: backzone
//  
func (antar antarcticaTimeZones) McMurdo() string {return "Antarctica/McMurdo" }

// Palmer - IANA Time Zone 'Antarctica/Palmer'.
// IANA Source File: southamerica
//  
func (antar antarcticaTimeZones) Palmer() string {return "Antarctica/Palmer" }

// Rothera - IANA Time Zone 'Antarctica/Rothera'.
// IANA Source File: antarctica
//  
func (antar antarcticaTimeZones) Rothera() string {return "Antarctica/Rothera" }

// Syowa - IANA Time Zone 'Antarctica/Syowa'.
// IANA Source File: antarctica
//  
func (antar antarcticaTimeZones) Syowa() string {return "Antarctica/Syowa" }

// Troll - IANA Time Zone 'Antarctica/Troll'.
// IANA Source File: antarctica
//  
func (antar antarcticaTimeZones) Troll() string {return "Antarctica/Troll" }

// Vostok - IANA Time Zone 'Antarctica/Vostok'.
// IANA Source File: antarctica
//  
func (antar antarcticaTimeZones) Vostok() string {return "Antarctica/Vostok" }

// asiaTimeZones - IANA Time Zones for 'Asia'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type asiaTimeZones string

// Aden - IANA Time Zone 'Asia/Aden'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Aden() string {return "Asia/Aden" }

// Almaty - IANA Time Zone 'Asia/Almaty'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Almaty() string {return "Asia/Almaty" }

// Amman - IANA Time Zone 'Asia/Amman'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Amman() string {return "Asia/Amman" }

// Anadyr - IANA Time Zone 'Asia/Anadyr'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Anadyr() string {return "Asia/Anadyr" }

// Aqtau - IANA Time Zone 'Asia/Aqtau'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Aqtau() string {return "Asia/Aqtau" }

// Aqtobe - IANA Time Zone 'Asia/Aqtobe'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Aqtobe() string {return "Asia/Aqtobe" }

// Ashgabat - IANA Time Zone 'Asia/Ashgabat'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Ashgabat() string {return "Asia/Ashgabat" }

// Atyrau - IANA Time Zone 'Asia/Atyrau'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Atyrau() string {return "Asia/Atyrau" }

// Baghdad - IANA Time Zone 'Asia/Baghdad'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Baghdad() string {return "Asia/Baghdad" }

// Bahrain - IANA Time Zone 'Asia/Bahrain'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Bahrain() string {return "Asia/Bahrain" }

// Baku - IANA Time Zone 'Asia/Baku'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Baku() string {return "Asia/Baku" }

// Bangkok - IANA Time Zone 'Asia/Bangkok'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Bangkok() string {return "Asia/Bangkok" }

// Barnaul - IANA Time Zone 'Asia/Barnaul'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Barnaul() string {return "Asia/Barnaul" }

// Beirut - IANA Time Zone 'Asia/Beirut'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Beirut() string {return "Asia/Beirut" }

// Bishkek - IANA Time Zone 'Asia/Bishkek'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Bishkek() string {return "Asia/Bishkek" }

// Brunei - IANA Time Zone 'Asia/Brunei'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Brunei() string {return "Asia/Brunei" }

// Chita - IANA Time Zone 'Asia/Chita'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Chita() string {return "Asia/Chita" }

// Choibalsan - IANA Time Zone 'Asia/Choibalsan'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Choibalsan() string {return "Asia/Choibalsan" }

// Chongqing - IANA Time Zone 'Asia/Chongqing'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Chongqing() string {return "Asia/Chongqing" }

// Colombo - IANA Time Zone 'Asia/Colombo'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Colombo() string {return "Asia/Colombo" }

// Damascus - IANA Time Zone 'Asia/Damascus'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Damascus() string {return "Asia/Damascus" }

// Dhaka - IANA Time Zone 'Asia/Dhaka'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Dhaka() string {return "Asia/Dhaka" }

// Dili - IANA Time Zone 'Asia/Dili'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Dili() string {return "Asia/Dili" }

// Dubai - IANA Time Zone 'Asia/Dubai'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Dubai() string {return "Asia/Dubai" }

// Dushanbe - IANA Time Zone 'Asia/Dushanbe'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Dushanbe() string {return "Asia/Dushanbe" }

// Famagusta - IANA Time Zone 'Asia/Famagusta'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Famagusta() string {return "Asia/Famagusta" }

// Gaza - IANA Time Zone 'Asia/Gaza'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Gaza() string {return "Asia/Gaza" }

// Hanoi - IANA Time Zone 'Asia/Hanoi'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Hanoi() string {return "Asia/Hanoi" }

// Harbin - IANA Time Zone 'Asia/Harbin'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Harbin() string {return "Asia/Harbin" }

// Hebron - IANA Time Zone 'Asia/Hebron'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Hebron() string {return "Asia/Hebron" }

// Ho_Chi_Minh - IANA Time Zone 'Asia/Ho_Chi_Minh'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Ho_Chi_Minh() string {return "Asia/Ho_Chi_Minh" }

// Hong_Kong - IANA Time Zone 'Asia/Hong_Kong'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Hong_Kong() string {return "Asia/Hong_Kong" }

// Hovd - IANA Time Zone 'Asia/Hovd'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Hovd() string {return "Asia/Hovd" }

// Irkutsk - IANA Time Zone 'Asia/Irkutsk'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Irkutsk() string {return "Asia/Irkutsk" }

// Jakarta - IANA Time Zone 'Asia/Jakarta'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Jakarta() string {return "Asia/Jakarta" }

// Jayapura - IANA Time Zone 'Asia/Jayapura'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Jayapura() string {return "Asia/Jayapura" }

// Jerusalem - IANA Time Zone 'Asia/Jerusalem'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Jerusalem() string {return "Asia/Jerusalem" }

// Kabul - IANA Time Zone 'Asia/Kabul'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Kabul() string {return "Asia/Kabul" }

// Kamchatka - IANA Time Zone 'Asia/Kamchatka'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Kamchatka() string {return "Asia/Kamchatka" }

// Karachi - IANA Time Zone 'Asia/Karachi'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Karachi() string {return "Asia/Karachi" }

// Kashgar - IANA Time Zone 'Asia/Kashgar'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Kashgar() string {return "Asia/Kashgar" }

// Kathmandu - IANA Time Zone 'Asia/Kathmandu'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Kathmandu() string {return "Asia/Kathmandu" }

// Khandyga - IANA Time Zone 'Asia/Khandyga'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Khandyga() string {return "Asia/Khandyga" }

// Kolkata - IANA Time Zone 'Asia/Kolkata'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Kolkata() string {return "Asia/Kolkata" }

// Krasnoyarsk - IANA Time Zone 'Asia/Krasnoyarsk'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Krasnoyarsk() string {return "Asia/Krasnoyarsk" }

// Kuala_Lumpur - IANA Time Zone 'Asia/Kuala_Lumpur'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Kuala_Lumpur() string {return "Asia/Kuala_Lumpur" }

// Kuching - IANA Time Zone 'Asia/Kuching'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Kuching() string {return "Asia/Kuching" }

// Kuwait - IANA Time Zone 'Asia/Kuwait'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Kuwait() string {return "Asia/Kuwait" }

// Macau - IANA Time Zone 'Asia/Macau'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Macau() string {return "Asia/Macau" }

// Magadan - IANA Time Zone 'Asia/Magadan'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Magadan() string {return "Asia/Magadan" }

// Makassar - IANA Time Zone 'Asia/Makassar'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Makassar() string {return "Asia/Makassar" }

// Manila - IANA Time Zone 'Asia/Manila'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Manila() string {return "Asia/Manila" }

// Muscat - IANA Time Zone 'Asia/Muscat'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Muscat() string {return "Asia/Muscat" }

// Nicosia - IANA Time Zone 'Asia/Nicosia'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Nicosia() string {return "Asia/Nicosia" }

// Novokuznetsk - IANA Time Zone 'Asia/Novokuznetsk'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Novokuznetsk() string {return "Asia/Novokuznetsk" }

// Novosibirsk - IANA Time Zone 'Asia/Novosibirsk'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Novosibirsk() string {return "Asia/Novosibirsk" }

// Omsk - IANA Time Zone 'Asia/Omsk'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Omsk() string {return "Asia/Omsk" }

// Oral - IANA Time Zone 'Asia/Oral'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Oral() string {return "Asia/Oral" }

// Phnom_Penh - IANA Time Zone 'Asia/Phnom_Penh'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Phnom_Penh() string {return "Asia/Phnom_Penh" }

// Pontianak - IANA Time Zone 'Asia/Pontianak'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Pontianak() string {return "Asia/Pontianak" }

// Pyongyang - IANA Time Zone 'Asia/Pyongyang'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Pyongyang() string {return "Asia/Pyongyang" }

// Qatar - IANA Time Zone 'Asia/Qatar'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Qatar() string {return "Asia/Qatar" }

// Qostanay - IANA Time Zone 'Asia/Qostanay'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Qostanay() string {return "Asia/Qostanay" }

// Qyzylorda - IANA Time Zone 'Asia/Qyzylorda'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Qyzylorda() string {return "Asia/Qyzylorda" }

// Riyadh - IANA Time Zone 'Asia/Riyadh'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Riyadh() string {return "Asia/Riyadh" }

// Sakhalin - IANA Time Zone 'Asia/Sakhalin'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Sakhalin() string {return "Asia/Sakhalin" }

// Samarkand - IANA Time Zone 'Asia/Samarkand'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Samarkand() string {return "Asia/Samarkand" }

// Seoul - IANA Time Zone 'Asia/Seoul'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Seoul() string {return "Asia/Seoul" }

// Shanghai - IANA Time Zone 'Asia/Shanghai'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Shanghai() string {return "Asia/Shanghai" }

// Singapore - IANA Time Zone 'Asia/Singapore'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Singapore() string {return "Asia/Singapore" }

// Srednekolymsk - IANA Time Zone 'Asia/Srednekolymsk'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Srednekolymsk() string {return "Asia/Srednekolymsk" }

// Taipei - IANA Time Zone 'Asia/Taipei'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Taipei() string {return "Asia/Taipei" }

// Tashkent - IANA Time Zone 'Asia/Tashkent'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Tashkent() string {return "Asia/Tashkent" }

// Tbilisi - IANA Time Zone 'Asia/Tbilisi'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Tbilisi() string {return "Asia/Tbilisi" }

// Tehran - IANA Time Zone 'Asia/Tehran'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Tehran() string {return "Asia/Tehran" }

// Tel_Aviv - IANA Time Zone 'Asia/Tel_Aviv'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Tel_Aviv() string {return "Asia/Tel_Aviv" }

// Thimphu - IANA Time Zone 'Asia/Thimphu'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Thimphu() string {return "Asia/Thimphu" }

// Tokyo - IANA Time Zone 'Asia/Tokyo'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Tokyo() string {return "Asia/Tokyo" }

// Tomsk - IANA Time Zone 'Asia/Tomsk'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Tomsk() string {return "Asia/Tomsk" }

// Ulaanbaatar - IANA Time Zone 'Asia/Ulaanbaatar'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Ulaanbaatar() string {return "Asia/Ulaanbaatar" }

// Urumqi - IANA Time Zone 'Asia/Urumqi'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Urumqi() string {return "Asia/Urumqi" }

// Ust-Nera - IANA Time Zone 'Asia/Ust-Nera'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) UstMinusNera() string {return "Asia/Ust-Nera" }

// Vientiane - IANA Time Zone 'Asia/Vientiane'.
// IANA Source File: backzone
//  
func (asiaT asiaTimeZones) Vientiane() string {return "Asia/Vientiane" }

// Vladivostok - IANA Time Zone 'Asia/Vladivostok'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Vladivostok() string {return "Asia/Vladivostok" }

// Yakutsk - IANA Time Zone 'Asia/Yakutsk'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Yakutsk() string {return "Asia/Yakutsk" }

// Yangon - IANA Time Zone 'Asia/Yangon'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Yangon() string {return "Asia/Yangon" }

// Yekaterinburg - IANA Time Zone 'Asia/Yekaterinburg'.
// IANA Source File: europe
//  
func (asiaT asiaTimeZones) Yekaterinburg() string {return "Asia/Yekaterinburg" }

// Yerevan - IANA Time Zone 'Asia/Yerevan'.
// IANA Source File: asia
//  
func (asiaT asiaTimeZones) Yerevan() string {return "Asia/Yerevan" }

// atlanticTimeZones - IANA Time Zones for 'Atlantic'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type atlanticTimeZones string

// Azores - IANA Time Zone 'Atlantic/Azores'.
// IANA Source File: europe
//  
func (atlan atlanticTimeZones) Azores() string {return "Atlantic/Azores" }

// Bermuda - IANA Time Zone 'Atlantic/Bermuda'.
// IANA Source File: northamerica
//  
func (atlan atlanticTimeZones) Bermuda() string {return "Atlantic/Bermuda" }

// Canary - IANA Time Zone 'Atlantic/Canary'.
// IANA Source File: europe
//  
func (atlan atlanticTimeZones) Canary() string {return "Atlantic/Canary" }

// Cape_Verde - IANA Time Zone 'Atlantic/Cape_Verde'.
// IANA Source File: africa
//  
func (atlan atlanticTimeZones) Cape_Verde() string {return "Atlantic/Cape_Verde" }

// Faroe - IANA Time Zone 'Atlantic/Faroe'.
// IANA Source File: europe
//  
func (atlan atlanticTimeZones) Faroe() string {return "Atlantic/Faroe" }

// Jan_Mayen - IANA Time Zone 'Atlantic/Jan_Mayen'.
// IANA Source File: backzone
//  
func (atlan atlanticTimeZones) Jan_Mayen() string {return "Atlantic/Jan_Mayen" }

// Madeira - IANA Time Zone 'Atlantic/Madeira'.
// IANA Source File: europe
//  
func (atlan atlanticTimeZones) Madeira() string {return "Atlantic/Madeira" }

// Reykjavik - IANA Time Zone 'Atlantic/Reykjavik'.
// IANA Source File: europe
//  
func (atlan atlanticTimeZones) Reykjavik() string {return "Atlantic/Reykjavik" }

// South_Georgia - IANA Time Zone 'Atlantic/South_Georgia'.
// IANA Source File: southamerica
//  
func (atlan atlanticTimeZones) South_Georgia() string {return "Atlantic/South_Georgia" }

// St_Helena - IANA Time Zone 'Atlantic/St_Helena'.
// IANA Source File: backzone
//  
func (atlan atlanticTimeZones) St_Helena() string {return "Atlantic/St_Helena" }

// Stanley - IANA Time Zone 'Atlantic/Stanley'.
// IANA Source File: southamerica
//  
func (atlan atlanticTimeZones) Stanley() string {return "Atlantic/Stanley" }

// australiaTimeZones - IANA Time Zones for 'Australia'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type australiaTimeZones string

// Adelaide - IANA Time Zone 'Australia/Adelaide'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Adelaide() string {return "Australia/Adelaide" }

// Brisbane - IANA Time Zone 'Australia/Brisbane'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Brisbane() string {return "Australia/Brisbane" }

// Broken_Hill - IANA Time Zone 'Australia/Broken_Hill'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Broken_Hill() string {return "Australia/Broken_Hill" }

// Currie - IANA Time Zone 'Australia/Currie'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Currie() string {return "Australia/Currie" }

// Darwin - IANA Time Zone 'Australia/Darwin'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Darwin() string {return "Australia/Darwin" }

// Eucla - IANA Time Zone 'Australia/Eucla'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Eucla() string {return "Australia/Eucla" }

// Hobart - IANA Time Zone 'Australia/Hobart'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Hobart() string {return "Australia/Hobart" }

// Lindeman - IANA Time Zone 'Australia/Lindeman'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Lindeman() string {return "Australia/Lindeman" }

// Lord_Howe - IANA Time Zone 'Australia/Lord_Howe'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Lord_Howe() string {return "Australia/Lord_Howe" }

// Melbourne - IANA Time Zone 'Australia/Melbourne'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Melbourne() string {return "Australia/Melbourne" }

// Perth - IANA Time Zone 'Australia/Perth'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Perth() string {return "Australia/Perth" }

// Sydney - IANA Time Zone 'Australia/Sydney'.
// IANA Source File: australasia
//  
func (austr australiaTimeZones) Sydney() string {return "Australia/Sydney" }

// Deprecated - Defines a collection of IANA
// Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type deprecatedTimeZones string

// Africa - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Africa() africaDeprecatedTimeZones {return "" }

// America - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) America() americaDeprecatedTimeZones {return "" }

// Antarctica - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Antarctica() antarcticaDeprecatedTimeZones {return "" }

// Arctic - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Arctic() arcticDeprecatedTimeZones {return "" }

// Asia - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Asia() asiaDeprecatedTimeZones {return "" }

// Atlantic - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Atlantic() atlanticDeprecatedTimeZones {return "" }

// Australia - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Australia() australiaDeprecatedTimeZones {return "" }

// Brazil - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Brazil() brazilDeprecatedTimeZones {return "" }

// Canada - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Canada() canadaDeprecatedTimeZones {return "" }

// Chile - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Chile() chileDeprecatedTimeZones {return "" }

// Cuba - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Cuba'
// Maps To Valid Time Zone: 'America/Havana'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Cuba() string { return "America/Havana" }

// Egypt - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Egypt'
// Maps To Valid Time Zone: 'Africa/Cairo'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Egypt() string { return "Africa/Cairo" }

// Eire - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Eire'
// Maps To Valid Time Zone: 'Europe/Dublin'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Eire() string { return "Europe/Dublin" }

// Etc - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Etc() etcDeprecatedTimeZones {return "" }

// Europe - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Europe() europeDeprecatedTimeZones {return "" }

// GB - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'GB'
// Maps To Valid Time Zone: 'Europe/London'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) GB() string { return "Europe/London" }

// GB-Eire - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'GB-Eire'
// Maps To Valid Time Zone: 'Europe/London'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) GBMinusEire() string { return "Europe/London" }

// GMT - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'GMT'
// Maps To Valid Time Zone: 'Etc/GMT'
//            IANA Source File: etcetera
//  
func (depre deprecatedTimeZones) GMT() string { return "Etc/GMT" }

// GMT+0 - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'GMT+0'
// Maps To Valid Time Zone: 'Etc/GMT'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) GMTPlus00() string { return "Etc/GMT" }

// GMT-0 - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'GMT-0'
// Maps To Valid Time Zone: 'Etc/GMT'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) GMTMinus00() string { return "Etc/GMT" }

// GMT0 - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'GMT0'
// Maps To Valid Time Zone: 'Etc/GMT'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) GMT00() string { return "Etc/GMT" }

// Greenwich - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Greenwich'
// Maps To Valid Time Zone: 'Etc/GMT'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Greenwich() string { return "Etc/GMT" }

// Hongkong - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Hongkong'
// Maps To Valid Time Zone: 'Asia/Hong_Kong'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Hongkong() string { return "Asia/Hong_Kong" }

// Iceland - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Iceland'
// Maps To Valid Time Zone: 'Atlantic/Reykjavik'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Iceland() string { return "Atlantic/Reykjavik" }

// Indian - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Indian() indianDeprecatedTimeZones {return "" }

// Iran - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Iran'
// Maps To Valid Time Zone: 'Asia/Tehran'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Iran() string { return "Asia/Tehran" }

// Israel - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Israel'
// Maps To Valid Time Zone: 'Asia/Jerusalem'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Israel() string { return "Asia/Jerusalem" }

// Jamaica - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Jamaica'
// Maps To Valid Time Zone: 'America/Jamaica'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Jamaica() string { return "America/Jamaica" }

// Japan - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Japan'
// Maps To Valid Time Zone: 'Asia/Tokyo'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Japan() string { return "Asia/Tokyo" }

// Kwajalein - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Kwajalein'
// Maps To Valid Time Zone: 'Pacific/Kwajalein'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Kwajalein() string { return "Pacific/Kwajalein" }

// Libya - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Libya'
// Maps To Valid Time Zone: 'Africa/Tripoli'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Libya() string { return "Africa/Tripoli" }

// Mexico - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Mexico() mexicoDeprecatedTimeZones {return "" }

// Navajo - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Navajo'
// Maps To Valid Time Zone: 'America/Denver'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Navajo() string { return "America/Denver" }

// NZ - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'NZ'
// Maps To Valid Time Zone: 'Pacific/Auckland'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) NZ() string { return "Pacific/Auckland" }

// NZ-CHAT - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'NZ-CHAT'
// Maps To Valid Time Zone: 'Pacific/Chatham'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) NZMinusCHAT() string { return "Pacific/Chatham" }

// Pacific - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) Pacific() pacificDeprecatedTimeZones {return "" }

// Poland - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Poland'
// Maps To Valid Time Zone: 'Europe/Warsaw'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Poland() string { return "Europe/Warsaw" }

// Portugal - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Portugal'
// Maps To Valid Time Zone: 'Europe/Lisbon'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Portugal() string { return "Europe/Lisbon" }

// PRC - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'PRC'
// Maps To Valid Time Zone: 'Asia/Shanghai'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) PRC() string { return "Asia/Shanghai" }

// ROC - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'ROC'
// Maps To Valid Time Zone: 'Asia/Taipei'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) ROC() string { return "Asia/Taipei" }

// ROK - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'ROK'
// Maps To Valid Time Zone: 'Asia/Seoul'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) ROK() string { return "Asia/Seoul" }

// Singapore - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Singapore'
// Maps To Valid Time Zone: 'Asia/Singapore'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Singapore() string { return "Asia/Singapore" }

// Turkey - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Turkey'
// Maps To Valid Time Zone: 'Europe/Istanbul'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Turkey() string { return "Europe/Istanbul" }

// UCT - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'UCT'
// Maps To Valid Time Zone: 'Etc/UTC'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) UCT() string { return "Etc/UTC" }

// Universal - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Universal'
// Maps To Valid Time Zone: 'Etc/UTC'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Universal() string { return "Etc/UTC" }

// US - A place holder which defines a sub-group
// of IANA 'Link' Time Zones.
//  
func (depre deprecatedTimeZones) US() uSDeprecatedTimeZones {return "" }

// UTC - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'UTC'
// Maps To Valid Time Zone: 'Etc/UTC'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) UTC() string { return "Etc/UTC" }

// W-SU - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'W-SU'
// Maps To Valid Time Zone: 'Europe/Moscow'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) WMinusSU() string { return "Europe/Moscow" }

// Zulu - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Zulu'
// Maps To Valid Time Zone: 'Etc/UTC'
//            IANA Source File: backward
//  
func (depre deprecatedTimeZones) Zulu() string { return "Etc/UTC" }

// etcTimeZones - IANA Time Zones for 'Etc'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type etcTimeZones string

// GMT - IANA Time Zone 'Etc/GMT'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMT() string {return "Etc/GMT" }

// GMT+1 - IANA Time Zone 'Etc/GMT+1'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus01() string {return "Etc/GMT+1" }

// GMT+2 - IANA Time Zone 'Etc/GMT+2'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus02() string {return "Etc/GMT+2" }

// GMT+3 - IANA Time Zone 'Etc/GMT+3'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus03() string {return "Etc/GMT+3" }

// GMT+4 - IANA Time Zone 'Etc/GMT+4'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus04() string {return "Etc/GMT+4" }

// GMT+5 - IANA Time Zone 'Etc/GMT+5'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus05() string {return "Etc/GMT+5" }

// GMT+6 - IANA Time Zone 'Etc/GMT+6'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus06() string {return "Etc/GMT+6" }

// GMT+7 - IANA Time Zone 'Etc/GMT+7'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus07() string {return "Etc/GMT+7" }

// GMT+8 - IANA Time Zone 'Etc/GMT+8'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus08() string {return "Etc/GMT+8" }

// GMT+9 - IANA Time Zone 'Etc/GMT+9'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus09() string {return "Etc/GMT+9" }

// GMT+10 - IANA Time Zone 'Etc/GMT+10'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus10() string {return "Etc/GMT+10" }

// GMT+11 - IANA Time Zone 'Etc/GMT+11'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus11() string {return "Etc/GMT+11" }

// GMT+12 - IANA Time Zone 'Etc/GMT+12'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTPlus12() string {return "Etc/GMT+12" }

// GMT-1 - IANA Time Zone 'Etc/GMT-1'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus01() string {return "Etc/GMT-1" }

// GMT-2 - IANA Time Zone 'Etc/GMT-2'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus02() string {return "Etc/GMT-2" }

// GMT-3 - IANA Time Zone 'Etc/GMT-3'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus03() string {return "Etc/GMT-3" }

// GMT-4 - IANA Time Zone 'Etc/GMT-4'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus04() string {return "Etc/GMT-4" }

// GMT-5 - IANA Time Zone 'Etc/GMT-5'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus05() string {return "Etc/GMT-5" }

// GMT-6 - IANA Time Zone 'Etc/GMT-6'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus06() string {return "Etc/GMT-6" }

// GMT-7 - IANA Time Zone 'Etc/GMT-7'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus07() string {return "Etc/GMT-7" }

// GMT-8 - IANA Time Zone 'Etc/GMT-8'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus08() string {return "Etc/GMT-8" }

// GMT-9 - IANA Time Zone 'Etc/GMT-9'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus09() string {return "Etc/GMT-9" }

// GMT-10 - IANA Time Zone 'Etc/GMT-10'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus10() string {return "Etc/GMT-10" }

// GMT-11 - IANA Time Zone 'Etc/GMT-11'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus11() string {return "Etc/GMT-11" }

// GMT-12 - IANA Time Zone 'Etc/GMT-12'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus12() string {return "Etc/GMT-12" }

// GMT-13 - IANA Time Zone 'Etc/GMT-13'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus13() string {return "Etc/GMT-13" }

// GMT-14 - IANA Time Zone 'Etc/GMT-14'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) GMTMinus14() string {return "Etc/GMT-14" }

// UTC - IANA Time Zone 'Etc/UTC'.
// IANA Source File: etcetera
//  
func (etcTi etcTimeZones) UTC() string {return "Etc/UTC" }

// europeTimeZones - IANA Time Zones for 'Europe'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type europeTimeZones string

// Amsterdam - IANA Time Zone 'Europe/Amsterdam'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Amsterdam() string {return "Europe/Amsterdam" }

// Andorra - IANA Time Zone 'Europe/Andorra'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Andorra() string {return "Europe/Andorra" }

// Astrakhan - IANA Time Zone 'Europe/Astrakhan'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Astrakhan() string {return "Europe/Astrakhan" }

// Athens - IANA Time Zone 'Europe/Athens'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Athens() string {return "Europe/Athens" }

// Belfast - IANA Time Zone 'Europe/Belfast'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Belfast() string {return "Europe/Belfast" }

// Belgrade - IANA Time Zone 'Europe/Belgrade'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Belgrade() string {return "Europe/Belgrade" }

// Berlin - IANA Time Zone 'Europe/Berlin'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Berlin() string {return "Europe/Berlin" }

// Brussels - IANA Time Zone 'Europe/Brussels'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Brussels() string {return "Europe/Brussels" }

// Bucharest - IANA Time Zone 'Europe/Bucharest'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Bucharest() string {return "Europe/Bucharest" }

// Budapest - IANA Time Zone 'Europe/Budapest'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Budapest() string {return "Europe/Budapest" }

// Chisinau - IANA Time Zone 'Europe/Chisinau'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Chisinau() string {return "Europe/Chisinau" }

// Copenhagen - IANA Time Zone 'Europe/Copenhagen'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Copenhagen() string {return "Europe/Copenhagen" }

// Dublin - IANA Time Zone 'Europe/Dublin'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Dublin() string {return "Europe/Dublin" }

// Gibraltar - IANA Time Zone 'Europe/Gibraltar'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Gibraltar() string {return "Europe/Gibraltar" }

// Guernsey - IANA Time Zone 'Europe/Guernsey'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Guernsey() string {return "Europe/Guernsey" }

// Helsinki - IANA Time Zone 'Europe/Helsinki'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Helsinki() string {return "Europe/Helsinki" }

// Isle_of_Man - IANA Time Zone 'Europe/Isle_of_Man'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Isle_of_Man() string {return "Europe/Isle_of_Man" }

// Istanbul - IANA Time Zone 'Europe/Istanbul'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Istanbul() string {return "Europe/Istanbul" }

// Jersey - IANA Time Zone 'Europe/Jersey'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Jersey() string {return "Europe/Jersey" }

// Kaliningrad - IANA Time Zone 'Europe/Kaliningrad'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Kaliningrad() string {return "Europe/Kaliningrad" }

// Kiev - IANA Time Zone 'Europe/Kiev'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Kiev() string {return "Europe/Kiev" }

// Kirov - IANA Time Zone 'Europe/Kirov'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Kirov() string {return "Europe/Kirov" }

// Lisbon - IANA Time Zone 'Europe/Lisbon'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Lisbon() string {return "Europe/Lisbon" }

// Ljubljana - IANA Time Zone 'Europe/Ljubljana'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Ljubljana() string {return "Europe/Ljubljana" }

// London - IANA Time Zone 'Europe/London'.
// IANA Source File: europe
//  
func (europ europeTimeZones) London() string {return "Europe/London" }

// Luxembourg - IANA Time Zone 'Europe/Luxembourg'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Luxembourg() string {return "Europe/Luxembourg" }

// Madrid - IANA Time Zone 'Europe/Madrid'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Madrid() string {return "Europe/Madrid" }

// Malta - IANA Time Zone 'Europe/Malta'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Malta() string {return "Europe/Malta" }

// Minsk - IANA Time Zone 'Europe/Minsk'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Minsk() string {return "Europe/Minsk" }

// Monaco - IANA Time Zone 'Europe/Monaco'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Monaco() string {return "Europe/Monaco" }

// Moscow - IANA Time Zone 'Europe/Moscow'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Moscow() string {return "Europe/Moscow" }

// Oslo - IANA Time Zone 'Europe/Oslo'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Oslo() string {return "Europe/Oslo" }

// Paris - IANA Time Zone 'Europe/Paris'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Paris() string {return "Europe/Paris" }

// Prague - IANA Time Zone 'Europe/Prague'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Prague() string {return "Europe/Prague" }

// Riga - IANA Time Zone 'Europe/Riga'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Riga() string {return "Europe/Riga" }

// Rome - IANA Time Zone 'Europe/Rome'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Rome() string {return "Europe/Rome" }

// Samara - IANA Time Zone 'Europe/Samara'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Samara() string {return "Europe/Samara" }

// Sarajevo - IANA Time Zone 'Europe/Sarajevo'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Sarajevo() string {return "Europe/Sarajevo" }

// Saratov - IANA Time Zone 'Europe/Saratov'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Saratov() string {return "Europe/Saratov" }

// Simferopol - IANA Time Zone 'Europe/Simferopol'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Simferopol() string {return "Europe/Simferopol" }

// Skopje - IANA Time Zone 'Europe/Skopje'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Skopje() string {return "Europe/Skopje" }

// Sofia - IANA Time Zone 'Europe/Sofia'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Sofia() string {return "Europe/Sofia" }

// Stockholm - IANA Time Zone 'Europe/Stockholm'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Stockholm() string {return "Europe/Stockholm" }

// Tallinn - IANA Time Zone 'Europe/Tallinn'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Tallinn() string {return "Europe/Tallinn" }

// Tirane - IANA Time Zone 'Europe/Tirane'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Tirane() string {return "Europe/Tirane" }

// Tiraspol - IANA Time Zone 'Europe/Tiraspol'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Tiraspol() string {return "Europe/Tiraspol" }

// Ulyanovsk - IANA Time Zone 'Europe/Ulyanovsk'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Ulyanovsk() string {return "Europe/Ulyanovsk" }

// Uzhgorod - IANA Time Zone 'Europe/Uzhgorod'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Uzhgorod() string {return "Europe/Uzhgorod" }

// Vaduz - IANA Time Zone 'Europe/Vaduz'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Vaduz() string {return "Europe/Vaduz" }

// Vienna - IANA Time Zone 'Europe/Vienna'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Vienna() string {return "Europe/Vienna" }

// Vilnius - IANA Time Zone 'Europe/Vilnius'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Vilnius() string {return "Europe/Vilnius" }

// Volgograd - IANA Time Zone 'Europe/Volgograd'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Volgograd() string {return "Europe/Volgograd" }

// Warsaw - IANA Time Zone 'Europe/Warsaw'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Warsaw() string {return "Europe/Warsaw" }

// Zagreb - IANA Time Zone 'Europe/Zagreb'.
// IANA Source File: backzone
//  
func (europ europeTimeZones) Zagreb() string {return "Europe/Zagreb" }

// Zaporozhye - IANA Time Zone 'Europe/Zaporozhye'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Zaporozhye() string {return "Europe/Zaporozhye" }

// Zurich - IANA Time Zone 'Europe/Zurich'.
// IANA Source File: europe
//  
func (europ europeTimeZones) Zurich() string {return "Europe/Zurich" }

// indianTimeZones - IANA Time Zones for 'Indian'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type indianTimeZones string

// Antananarivo - IANA Time Zone 'Indian/Antananarivo'.
// IANA Source File: backzone
//  
func (india indianTimeZones) Antananarivo() string {return "Indian/Antananarivo" }

// Chagos - IANA Time Zone 'Indian/Chagos'.
// IANA Source File: asia
//  
func (india indianTimeZones) Chagos() string {return "Indian/Chagos" }

// Christmas - IANA Time Zone 'Indian/Christmas'.
// IANA Source File: australasia
//  
func (india indianTimeZones) Christmas() string {return "Indian/Christmas" }

// Cocos - IANA Time Zone 'Indian/Cocos'.
// IANA Source File: australasia
//  
func (india indianTimeZones) Cocos() string {return "Indian/Cocos" }

// Comoro - IANA Time Zone 'Indian/Comoro'.
// IANA Source File: backzone
//  
func (india indianTimeZones) Comoro() string {return "Indian/Comoro" }

// Kerguelen - IANA Time Zone 'Indian/Kerguelen'.
// IANA Source File: antarctica
//  
func (india indianTimeZones) Kerguelen() string {return "Indian/Kerguelen" }

// Mahe - IANA Time Zone 'Indian/Mahe'.
// IANA Source File: africa
//  
func (india indianTimeZones) Mahe() string {return "Indian/Mahe" }

// Maldives - IANA Time Zone 'Indian/Maldives'.
// IANA Source File: asia
//  
func (india indianTimeZones) Maldives() string {return "Indian/Maldives" }

// Mauritius - IANA Time Zone 'Indian/Mauritius'.
// IANA Source File: africa
//  
func (india indianTimeZones) Mauritius() string {return "Indian/Mauritius" }

// Mayotte - IANA Time Zone 'Indian/Mayotte'.
// IANA Source File: backzone
//  
func (india indianTimeZones) Mayotte() string {return "Indian/Mayotte" }

// Reunion - IANA Time Zone 'Indian/Reunion'.
// IANA Source File: africa
//  
func (india indianTimeZones) Reunion() string {return "Indian/Reunion" }

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

// pacificTimeZones - IANA Time Zones for 'Pacific'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type pacificTimeZones string

// Apia - IANA Time Zone 'Pacific/Apia'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Apia() string {return "Pacific/Apia" }

// Auckland - IANA Time Zone 'Pacific/Auckland'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Auckland() string {return "Pacific/Auckland" }

// Bougainville - IANA Time Zone 'Pacific/Bougainville'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Bougainville() string {return "Pacific/Bougainville" }

// Chatham - IANA Time Zone 'Pacific/Chatham'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Chatham() string {return "Pacific/Chatham" }

// Chuuk - IANA Time Zone 'Pacific/Chuuk'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Chuuk() string {return "Pacific/Chuuk" }

// Easter - IANA Time Zone 'Pacific/Easter'.
// IANA Source File: southamerica
//  
func (pacif pacificTimeZones) Easter() string {return "Pacific/Easter" }

// Efate - IANA Time Zone 'Pacific/Efate'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Efate() string {return "Pacific/Efate" }

// Enderbury - IANA Time Zone 'Pacific/Enderbury'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Enderbury() string {return "Pacific/Enderbury" }

// Fakaofo - IANA Time Zone 'Pacific/Fakaofo'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Fakaofo() string {return "Pacific/Fakaofo" }

// Fiji - IANA Time Zone 'Pacific/Fiji'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Fiji() string {return "Pacific/Fiji" }

// Funafuti - IANA Time Zone 'Pacific/Funafuti'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Funafuti() string {return "Pacific/Funafuti" }

// Galapagos - IANA Time Zone 'Pacific/Galapagos'.
// IANA Source File: southamerica
//  
func (pacif pacificTimeZones) Galapagos() string {return "Pacific/Galapagos" }

// Gambier - IANA Time Zone 'Pacific/Gambier'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Gambier() string {return "Pacific/Gambier" }

// Guadalcanal - IANA Time Zone 'Pacific/Guadalcanal'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Guadalcanal() string {return "Pacific/Guadalcanal" }

// Guam - IANA Time Zone 'Pacific/Guam'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Guam() string {return "Pacific/Guam" }

// Honolulu - IANA Time Zone 'Pacific/Honolulu'.
// IANA Source File: northamerica
//  
func (pacif pacificTimeZones) Honolulu() string {return "Pacific/Honolulu" }

// Johnston - IANA Time Zone 'Pacific/Johnston'.
// IANA Source File: backzone
//  
func (pacif pacificTimeZones) Johnston() string {return "Pacific/Johnston" }

// Kiritimati - IANA Time Zone 'Pacific/Kiritimati'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Kiritimati() string {return "Pacific/Kiritimati" }

// Kosrae - IANA Time Zone 'Pacific/Kosrae'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Kosrae() string {return "Pacific/Kosrae" }

// Kwajalein - IANA Time Zone 'Pacific/Kwajalein'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Kwajalein() string {return "Pacific/Kwajalein" }

// Majuro - IANA Time Zone 'Pacific/Majuro'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Majuro() string {return "Pacific/Majuro" }

// Marquesas - IANA Time Zone 'Pacific/Marquesas'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Marquesas() string {return "Pacific/Marquesas" }

// Midway - IANA Time Zone 'Pacific/Midway'.
// IANA Source File: backzone
//  
func (pacif pacificTimeZones) Midway() string {return "Pacific/Midway" }

// Nauru - IANA Time Zone 'Pacific/Nauru'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Nauru() string {return "Pacific/Nauru" }

// Niue - IANA Time Zone 'Pacific/Niue'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Niue() string {return "Pacific/Niue" }

// Norfolk - IANA Time Zone 'Pacific/Norfolk'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Norfolk() string {return "Pacific/Norfolk" }

// Noumea - IANA Time Zone 'Pacific/Noumea'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Noumea() string {return "Pacific/Noumea" }

// Pago_Pago - IANA Time Zone 'Pacific/Pago_Pago'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Pago_Pago() string {return "Pacific/Pago_Pago" }

// Palau - IANA Time Zone 'Pacific/Palau'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Palau() string {return "Pacific/Palau" }

// Pitcairn - IANA Time Zone 'Pacific/Pitcairn'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Pitcairn() string {return "Pacific/Pitcairn" }

// Pohnpei - IANA Time Zone 'Pacific/Pohnpei'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Pohnpei() string {return "Pacific/Pohnpei" }

// Port_Moresby - IANA Time Zone 'Pacific/Port_Moresby'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Port_Moresby() string {return "Pacific/Port_Moresby" }

// Rarotonga - IANA Time Zone 'Pacific/Rarotonga'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Rarotonga() string {return "Pacific/Rarotonga" }

// Saipan - IANA Time Zone 'Pacific/Saipan'.
// IANA Source File: backzone
//  
func (pacif pacificTimeZones) Saipan() string {return "Pacific/Saipan" }

// Tahiti - IANA Time Zone 'Pacific/Tahiti'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Tahiti() string {return "Pacific/Tahiti" }

// Tarawa - IANA Time Zone 'Pacific/Tarawa'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Tarawa() string {return "Pacific/Tarawa" }

// Tongatapu - IANA Time Zone 'Pacific/Tongatapu'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Tongatapu() string {return "Pacific/Tongatapu" }

// Wake - IANA Time Zone 'Pacific/Wake'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Wake() string {return "Pacific/Wake" }

// Wallis - IANA Time Zone 'Pacific/Wallis'.
// IANA Source File: australasia
//  
func (pacif pacificTimeZones) Wallis() string {return "Pacific/Wallis" }

// IANA Time Zones in located in 'Argentina'.
//  
// The Parent Group is 'America'
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type argentinaTimeZones string

// Buenos_Aires - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) Buenos_Aires() string {return "America/Argentina/Buenos_Aires" }

// Catamarca - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) Catamarca() string {return "America/Argentina/Catamarca" }

// ComodRivadavia - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) ComodRivadavia() string {return "America/Argentina/ComodRivadavia" }

// Cordoba - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) Cordoba() string {return "America/Argentina/Cordoba" }

// Jujuy - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) Jujuy() string {return "America/Argentina/Jujuy" }

// La_Rioja - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) La_Rioja() string {return "America/Argentina/La_Rioja" }

// Mendoza - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) Mendoza() string {return "America/Argentina/Mendoza" }

// Rio_Gallegos - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) Rio_Gallegos() string {return "America/Argentina/Rio_Gallegos" }

// Salta - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) Salta() string {return "America/Argentina/Salta" }

// San_Juan - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) San_Juan() string {return "America/Argentina/San_Juan" }

// San_Luis - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) San_Luis() string {return "America/Argentina/San_Luis" }

// Tucuman - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) Tucuman() string {return "America/Argentina/Tucuman" }

// Ushuaia - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (argen argentinaTimeZones) Ushuaia() string {return "America/Argentina/Ushuaia" }

// IANA Time Zones in located in 'Indiana'.
//  
// The Parent Group is 'America'
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type indianaTimeZones string

// Indianapolis - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (india indianaTimeZones) Indianapolis() string {return "America/Indiana/Indianapolis" }

// Knox - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (india indianaTimeZones) Knox() string {return "America/Indiana/Knox" }

// Marengo - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (india indianaTimeZones) Marengo() string {return "America/Indiana/Marengo" }

// Petersburg - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (india indianaTimeZones) Petersburg() string {return "America/Indiana/Petersburg" }

// Tell_City - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (india indianaTimeZones) Tell_City() string {return "America/Indiana/Tell_City" }

// Vevay - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (india indianaTimeZones) Vevay() string {return "America/Indiana/Vevay" }

// Vincennes - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (india indianaTimeZones) Vincennes() string {return "America/Indiana/Vincennes" }

// Winamac - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (india indianaTimeZones) Winamac() string {return "America/Indiana/Winamac" }

// IANA Time Zones in located in 'Kentucky'.
//  
// The Parent Group is 'America'
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type kentuckyTimeZones string

// Louisville - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (kentu kentuckyTimeZones) Louisville() string {return "America/Kentucky/Louisville" }

// Monticello - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (kentu kentuckyTimeZones) Monticello() string {return "America/Kentucky/Monticello" }

// IANA Time Zones in located in 'North_Dakota'.
//  
// The Parent Group is 'America'
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type north_DakotaTimeZones string

// Beulah - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (north north_DakotaTimeZones) Beulah() string {return "America/North_Dakota/Beulah" }

// Center - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (north north_DakotaTimeZones) Center() string {return "America/North_Dakota/Center" }

// New_Salem - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (north north_DakotaTimeZones) New_Salem() string {return "America/North_Dakota/New_Salem" }

// africaDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type africaDeprecatedTimeZones string

// Addis_Ababa - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Addis_Ababa'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Addis_Ababa() string { return "Africa/Nairobi" }

// Asmara - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Asmara'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Asmara() string { return "Africa/Nairobi" }

// Asmera - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Asmera'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: backward
//  
func (afric africaDeprecatedTimeZones) Asmera() string { return "Africa/Nairobi" }

// Bamako - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Bamako'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Bamako() string { return "Africa/Abidjan" }

// Bangui - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Bangui'
//     Maps To Valid Time Zone: 'Africa/Lagos'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Bangui() string { return "Africa/Lagos" }

// Banjul - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Banjul'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Banjul() string { return "Africa/Abidjan" }

// Blantyre - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Blantyre'
//     Maps To Valid Time Zone: 'Africa/Maputo'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Blantyre() string { return "Africa/Maputo" }

// Brazzaville - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Brazzaville'
//     Maps To Valid Time Zone: 'Africa/Lagos'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Brazzaville() string { return "Africa/Lagos" }

// Bujumbura - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Bujumbura'
//     Maps To Valid Time Zone: 'Africa/Maputo'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Bujumbura() string { return "Africa/Maputo" }

// Conakry - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Conakry'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Conakry() string { return "Africa/Abidjan" }

// Dakar - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Dakar'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Dakar() string { return "Africa/Abidjan" }

// Dar_es_Salaam - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Dar_es_Salaam'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Dar_es_Salaam() string { return "Africa/Nairobi" }

// Djibouti - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Djibouti'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Djibouti() string { return "Africa/Nairobi" }

// Douala - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Douala'
//     Maps To Valid Time Zone: 'Africa/Lagos'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Douala() string { return "Africa/Lagos" }

// Freetown - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Freetown'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Freetown() string { return "Africa/Abidjan" }

// Gaborone - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Gaborone'
//     Maps To Valid Time Zone: 'Africa/Maputo'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Gaborone() string { return "Africa/Maputo" }

// Harare - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Harare'
//     Maps To Valid Time Zone: 'Africa/Maputo'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Harare() string { return "Africa/Maputo" }

// Kampala - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Kampala'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Kampala() string { return "Africa/Nairobi" }

// Kigali - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Kigali'
//     Maps To Valid Time Zone: 'Africa/Maputo'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Kigali() string { return "Africa/Maputo" }

// Kinshasa - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Kinshasa'
//     Maps To Valid Time Zone: 'Africa/Lagos'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Kinshasa() string { return "Africa/Lagos" }

// Libreville - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Libreville'
//     Maps To Valid Time Zone: 'Africa/Lagos'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Libreville() string { return "Africa/Lagos" }

// Lome - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Lome'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Lome() string { return "Africa/Abidjan" }

// Luanda - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Luanda'
//     Maps To Valid Time Zone: 'Africa/Lagos'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Luanda() string { return "Africa/Lagos" }

// Lubumbashi - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Lubumbashi'
//     Maps To Valid Time Zone: 'Africa/Maputo'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Lubumbashi() string { return "Africa/Maputo" }

// Lusaka - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Lusaka'
//     Maps To Valid Time Zone: 'Africa/Maputo'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Lusaka() string { return "Africa/Maputo" }

// Malabo - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Malabo'
//     Maps To Valid Time Zone: 'Africa/Lagos'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Malabo() string { return "Africa/Lagos" }

// Maseru - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Maseru'
//     Maps To Valid Time Zone: 'Africa/Johannesburg'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Maseru() string { return "Africa/Johannesburg" }

// Mbabane - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Mbabane'
//     Maps To Valid Time Zone: 'Africa/Johannesburg'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Mbabane() string { return "Africa/Johannesburg" }

// Mogadishu - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Mogadishu'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Mogadishu() string { return "Africa/Nairobi" }

// Niamey - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Niamey'
//     Maps To Valid Time Zone: 'Africa/Lagos'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Niamey() string { return "Africa/Lagos" }

// Nouakchott - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Nouakchott'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Nouakchott() string { return "Africa/Abidjan" }

// Ouagadougou - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Ouagadougou'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) Ouagadougou() string { return "Africa/Abidjan" }

// Porto-Novo - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Porto-Novo'
//     Maps To Valid Time Zone: 'Africa/Lagos'
//            IANA Source File: africa
//  
func (afric africaDeprecatedTimeZones) PortoMinusNovo() string { return "Africa/Lagos" }

// Timbuktu - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Africa/Timbuktu'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: backward
//  
func (afric africaDeprecatedTimeZones) Timbuktu() string { return "Africa/Abidjan" }

// americaDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type americaDeprecatedTimeZones string

// Anguilla - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Anguilla'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Anguilla() string { return "America/Port_of_Spain" }

// Antigua - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Antigua'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Antigua() string { return "America/Port_of_Spain" }

// Argentina - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// Zones identify deprecated or obsolete time zones. These
// obsolete time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Argentina'
// Maps To Valid Time Zone: 'Argentina'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Argentina() argentinaDeprecatedTimeZones { return "" }

// Aruba - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Aruba'
//     Maps To Valid Time Zone: 'America/Curacao'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Aruba() string { return "America/Curacao" }

// Atka - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Atka'
//     Maps To Valid Time Zone: 'America/Adak'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Atka() string { return "America/Adak" }

// Buenos_Aires - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Buenos_Aires'
//     Maps To Valid Time Zone: 'America/Argentina/Buenos_Aires'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Buenos_Aires() string { return "America/Argentina/Buenos_Aires" }

// Catamarca - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Catamarca'
//     Maps To Valid Time Zone: 'America/Argentina/Catamarca'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Catamarca() string { return "America/Argentina/Catamarca" }

// Cayman - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Cayman'
//     Maps To Valid Time Zone: 'America/Panama'
//            IANA Source File: northamerica
//  
func (ameri americaDeprecatedTimeZones) Cayman() string { return "America/Panama" }

// Coral_Harbour - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Coral_Harbour'
//     Maps To Valid Time Zone: 'America/Atikokan'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Coral_Harbour() string { return "America/Atikokan" }

// Cordoba - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Cordoba'
//     Maps To Valid Time Zone: 'America/Argentina/Cordoba'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Cordoba() string { return "America/Argentina/Cordoba" }

// Dominica - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Dominica'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Dominica() string { return "America/Port_of_Spain" }

// Ensenada - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Ensenada'
//     Maps To Valid Time Zone: 'America/Tijuana'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Ensenada() string { return "America/Tijuana" }

// Fort_Wayne - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Fort_Wayne'
//     Maps To Valid Time Zone: 'America/Indiana/Indianapolis'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Fort_Wayne() string { return "America/Indiana/Indianapolis" }

// Grenada - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Grenada'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Grenada() string { return "America/Port_of_Spain" }

// Guadeloupe - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Guadeloupe'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Guadeloupe() string { return "America/Port_of_Spain" }

// Indianapolis - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Indianapolis'
//     Maps To Valid Time Zone: 'America/Indiana/Indianapolis'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Indianapolis() string { return "America/Indiana/Indianapolis" }

// Jujuy - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Jujuy'
//     Maps To Valid Time Zone: 'America/Argentina/Jujuy'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Jujuy() string { return "America/Argentina/Jujuy" }

// Knox_IN - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Knox_IN'
//     Maps To Valid Time Zone: 'America/Indiana/Knox'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Knox_IN() string { return "America/Indiana/Knox" }

// Kralendijk - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Kralendijk'
//     Maps To Valid Time Zone: 'America/Curacao'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Kralendijk() string { return "America/Curacao" }

// Louisville - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Louisville'
//     Maps To Valid Time Zone: 'America/Kentucky/Louisville'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Louisville() string { return "America/Kentucky/Louisville" }

// Lower_Princes - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Lower_Princes'
//     Maps To Valid Time Zone: 'America/Curacao'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Lower_Princes() string { return "America/Curacao" }

// Marigot - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Marigot'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Marigot() string { return "America/Port_of_Spain" }

// Mendoza - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Mendoza'
//     Maps To Valid Time Zone: 'America/Argentina/Mendoza'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Mendoza() string { return "America/Argentina/Mendoza" }

// Montreal - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Montreal'
//     Maps To Valid Time Zone: 'America/Toronto'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Montreal() string { return "America/Toronto" }

// Montserrat - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Montserrat'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Montserrat() string { return "America/Port_of_Spain" }

// Porto_Acre - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Porto_Acre'
//     Maps To Valid Time Zone: 'America/Rio_Branco'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Porto_Acre() string { return "America/Rio_Branco" }

// Rosario - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Rosario'
//     Maps To Valid Time Zone: 'America/Argentina/Cordoba'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Rosario() string { return "America/Argentina/Cordoba" }

// Santa_Isabel - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Santa_Isabel'
//     Maps To Valid Time Zone: 'America/Tijuana'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Santa_Isabel() string { return "America/Tijuana" }

// Shiprock - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Shiprock'
//     Maps To Valid Time Zone: 'America/Denver'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Shiprock() string { return "America/Denver" }

// St_Barthelemy - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/St_Barthelemy'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) St_Barthelemy() string { return "America/Port_of_Spain" }

// St_Kitts - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/St_Kitts'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) St_Kitts() string { return "America/Port_of_Spain" }

// St_Lucia - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/St_Lucia'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) St_Lucia() string { return "America/Port_of_Spain" }

// St_Thomas - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/St_Thomas'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) St_Thomas() string { return "America/Port_of_Spain" }

// St_Vincent - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/St_Vincent'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) St_Vincent() string { return "America/Port_of_Spain" }

// Tortola - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Tortola'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: southamerica
//  
func (ameri americaDeprecatedTimeZones) Tortola() string { return "America/Port_of_Spain" }

// Virgin - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'America/Virgin'
//     Maps To Valid Time Zone: 'America/Port_of_Spain'
//            IANA Source File: backward
//  
func (ameri americaDeprecatedTimeZones) Virgin() string { return "America/Port_of_Spain" }

// antarcticaDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type antarcticaDeprecatedTimeZones string

// McMurdo - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Antarctica/McMurdo'
//     Maps To Valid Time Zone: 'Pacific/Auckland'
//            IANA Source File: australasia
//  
func (antar antarcticaDeprecatedTimeZones) McMurdo() string { return "Pacific/Auckland" }

// South_Pole - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Antarctica/South_Pole'
//     Maps To Valid Time Zone: 'Pacific/Auckland'
//            IANA Source File: backward
//  
func (antar antarcticaDeprecatedTimeZones) South_Pole() string { return "Pacific/Auckland" }

// arcticDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type arcticDeprecatedTimeZones string

// Longyearbyen - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Arctic/Longyearbyen'
//     Maps To Valid Time Zone: 'Europe/Oslo'
//            IANA Source File: europe
//  
func (arcti arcticDeprecatedTimeZones) Longyearbyen() string { return "Europe/Oslo" }

// asiaDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type asiaDeprecatedTimeZones string

// Aden - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Aden'
//     Maps To Valid Time Zone: 'Asia/Riyadh'
//            IANA Source File: asia
//  
func (asiaD asiaDeprecatedTimeZones) Aden() string { return "Asia/Riyadh" }

// Ashkhabad - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Ashkhabad'
//     Maps To Valid Time Zone: 'Asia/Ashgabat'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Ashkhabad() string { return "Asia/Ashgabat" }

// Bahrain - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Bahrain'
//     Maps To Valid Time Zone: 'Asia/Qatar'
//            IANA Source File: asia
//  
func (asiaD asiaDeprecatedTimeZones) Bahrain() string { return "Asia/Qatar" }

// Calcutta - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Calcutta'
//     Maps To Valid Time Zone: 'Asia/Kolkata'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Calcutta() string { return "Asia/Kolkata" }

// Chongqing - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Chongqing'
//     Maps To Valid Time Zone: 'Asia/Shanghai'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Chongqing() string { return "Asia/Shanghai" }

// Chungking - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Chungking'
//     Maps To Valid Time Zone: 'Asia/Shanghai'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Chungking() string { return "Asia/Shanghai" }

// Dacca - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Dacca'
//     Maps To Valid Time Zone: 'Asia/Dhaka'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Dacca() string { return "Asia/Dhaka" }

// Harbin - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Harbin'
//     Maps To Valid Time Zone: 'Asia/Shanghai'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Harbin() string { return "Asia/Shanghai" }

// Istanbul - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Istanbul'
//     Maps To Valid Time Zone: 'Europe/Istanbul'
//            IANA Source File: europe
//  
func (asiaD asiaDeprecatedTimeZones) Istanbul() string { return "Europe/Istanbul" }

// Kashgar - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Kashgar'
//     Maps To Valid Time Zone: 'Asia/Urumqi'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Kashgar() string { return "Asia/Urumqi" }

// Katmandu - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Katmandu'
//     Maps To Valid Time Zone: 'Asia/Kathmandu'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Katmandu() string { return "Asia/Kathmandu" }

// Kuwait - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Kuwait'
//     Maps To Valid Time Zone: 'Asia/Riyadh'
//            IANA Source File: asia
//  
func (asiaD asiaDeprecatedTimeZones) Kuwait() string { return "Asia/Riyadh" }

// Macao - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Macao'
//     Maps To Valid Time Zone: 'Asia/Macau'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Macao() string { return "Asia/Macau" }

// Muscat - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Muscat'
//     Maps To Valid Time Zone: 'Asia/Dubai'
//            IANA Source File: asia
//  
func (asiaD asiaDeprecatedTimeZones) Muscat() string { return "Asia/Dubai" }

// Phnom_Penh - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Phnom_Penh'
//     Maps To Valid Time Zone: 'Asia/Bangkok'
//            IANA Source File: asia
//  
func (asiaD asiaDeprecatedTimeZones) Phnom_Penh() string { return "Asia/Bangkok" }

// Rangoon - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Rangoon'
//     Maps To Valid Time Zone: 'Asia/Yangon'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Rangoon() string { return "Asia/Yangon" }

// Saigon - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Saigon'
//     Maps To Valid Time Zone: 'Asia/Ho_Chi_Minh'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Saigon() string { return "Asia/Ho_Chi_Minh" }

// Tel_Aviv - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Tel_Aviv'
//     Maps To Valid Time Zone: 'Asia/Jerusalem'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Tel_Aviv() string { return "Asia/Jerusalem" }

// Thimbu - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Thimbu'
//     Maps To Valid Time Zone: 'Asia/Thimphu'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Thimbu() string { return "Asia/Thimphu" }

// Ujung_Pandang - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Ujung_Pandang'
//     Maps To Valid Time Zone: 'Asia/Makassar'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Ujung_Pandang() string { return "Asia/Makassar" }

// Ulan_Bator - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Ulan_Bator'
//     Maps To Valid Time Zone: 'Asia/Ulaanbaatar'
//            IANA Source File: backward
//  
func (asiaD asiaDeprecatedTimeZones) Ulan_Bator() string { return "Asia/Ulaanbaatar" }

// Vientiane - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Asia/Vientiane'
//     Maps To Valid Time Zone: 'Asia/Bangkok'
//            IANA Source File: asia
//  
func (asiaD asiaDeprecatedTimeZones) Vientiane() string { return "Asia/Bangkok" }

// atlanticDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type atlanticDeprecatedTimeZones string

// Faeroe - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Atlantic/Faeroe'
//     Maps To Valid Time Zone: 'Atlantic/Faroe'
//            IANA Source File: backward
//  
func (atlan atlanticDeprecatedTimeZones) Faeroe() string { return "Atlantic/Faroe" }

// Jan_Mayen - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Atlantic/Jan_Mayen'
//     Maps To Valid Time Zone: 'Europe/Oslo'
//            IANA Source File: backward
//  
func (atlan atlanticDeprecatedTimeZones) Jan_Mayen() string { return "Europe/Oslo" }

// St_Helena - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Atlantic/St_Helena'
//     Maps To Valid Time Zone: 'Africa/Abidjan'
//            IANA Source File: africa
//  
func (atlan atlanticDeprecatedTimeZones) St_Helena() string { return "Africa/Abidjan" }

// australiaDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type australiaDeprecatedTimeZones string

// ACT - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/ACT'
//     Maps To Valid Time Zone: 'Australia/Sydney'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) ACT() string { return "Australia/Sydney" }

// Canberra - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/Canberra'
//     Maps To Valid Time Zone: 'Australia/Sydney'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) Canberra() string { return "Australia/Sydney" }

// LHI - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/LHI'
//     Maps To Valid Time Zone: 'Australia/Lord_Howe'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) LHI() string { return "Australia/Lord_Howe" }

// North - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/North'
//     Maps To Valid Time Zone: 'Australia/Darwin'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) North() string { return "Australia/Darwin" }

// NSW - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/NSW'
//     Maps To Valid Time Zone: 'Australia/Sydney'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) NSW() string { return "Australia/Sydney" }

// Queensland - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/Queensland'
//     Maps To Valid Time Zone: 'Australia/Brisbane'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) Queensland() string { return "Australia/Brisbane" }

// South - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/South'
//     Maps To Valid Time Zone: 'Australia/Adelaide'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) South() string { return "Australia/Adelaide" }

// Tasmania - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/Tasmania'
//     Maps To Valid Time Zone: 'Australia/Hobart'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) Tasmania() string { return "Australia/Hobart" }

// Victoria - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/Victoria'
//     Maps To Valid Time Zone: 'Australia/Melbourne'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) Victoria() string { return "Australia/Melbourne" }

// West - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/West'
//     Maps To Valid Time Zone: 'Australia/Perth'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) West() string { return "Australia/Perth" }

// Yancowinna - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Australia/Yancowinna'
//     Maps To Valid Time Zone: 'Australia/Broken_Hill'
//            IANA Source File: backward
//  
func (austr australiaDeprecatedTimeZones) Yancowinna() string { return "Australia/Broken_Hill" }

// brazilDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type brazilDeprecatedTimeZones string

// Acre - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Brazil/Acre'
//     Maps To Valid Time Zone: 'America/Rio_Branco'
//            IANA Source File: backward
//  
func (brazi brazilDeprecatedTimeZones) Acre() string { return "America/Rio_Branco" }

// DeNoronha - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Brazil/DeNoronha'
//     Maps To Valid Time Zone: 'America/Noronha'
//            IANA Source File: backward
//  
func (brazi brazilDeprecatedTimeZones) DeNoronha() string { return "America/Noronha" }

// East - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Brazil/East'
//     Maps To Valid Time Zone: 'America/Sao_Paulo'
//            IANA Source File: backward
//  
func (brazi brazilDeprecatedTimeZones) East() string { return "America/Sao_Paulo" }

// West - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Brazil/West'
//     Maps To Valid Time Zone: 'America/Manaus'
//            IANA Source File: backward
//  
func (brazi brazilDeprecatedTimeZones) West() string { return "America/Manaus" }

// canadaDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type canadaDeprecatedTimeZones string

// Atlantic - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Canada/Atlantic'
//     Maps To Valid Time Zone: 'America/Halifax'
//            IANA Source File: backward
//  
func (canad canadaDeprecatedTimeZones) Atlantic() string { return "America/Halifax" }

// Central - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Canada/Central'
//     Maps To Valid Time Zone: 'America/Winnipeg'
//            IANA Source File: backward
//  
func (canad canadaDeprecatedTimeZones) Central() string { return "America/Winnipeg" }

// Eastern - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Canada/Eastern'
//     Maps To Valid Time Zone: 'America/Toronto'
//            IANA Source File: backward
//  
func (canad canadaDeprecatedTimeZones) Eastern() string { return "America/Toronto" }

// Mountain - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Canada/Mountain'
//     Maps To Valid Time Zone: 'America/Edmonton'
//            IANA Source File: backward
//  
func (canad canadaDeprecatedTimeZones) Mountain() string { return "America/Edmonton" }

// Newfoundland - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Canada/Newfoundland'
//     Maps To Valid Time Zone: 'America/St_Johns'
//            IANA Source File: backward
//  
func (canad canadaDeprecatedTimeZones) Newfoundland() string { return "America/St_Johns" }

// Pacific - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Canada/Pacific'
//     Maps To Valid Time Zone: 'America/Vancouver'
//            IANA Source File: backward
//  
func (canad canadaDeprecatedTimeZones) Pacific() string { return "America/Vancouver" }

// Saskatchewan - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Canada/Saskatchewan'
//     Maps To Valid Time Zone: 'America/Regina'
//            IANA Source File: backward
//  
func (canad canadaDeprecatedTimeZones) Saskatchewan() string { return "America/Regina" }

// Yukon - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Canada/Yukon'
//     Maps To Valid Time Zone: 'America/Whitehorse'
//            IANA Source File: backward
//  
func (canad canadaDeprecatedTimeZones) Yukon() string { return "America/Whitehorse" }

// chileDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type chileDeprecatedTimeZones string

// Continental - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Chile/Continental'
//     Maps To Valid Time Zone: 'America/Santiago'
//            IANA Source File: backward
//  
func (chile chileDeprecatedTimeZones) Continental() string { return "America/Santiago" }

// EasterIsland - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Chile/EasterIsland'
//     Maps To Valid Time Zone: 'Pacific/Easter'
//            IANA Source File: backward
//  
func (chile chileDeprecatedTimeZones) EasterIsland() string { return "Pacific/Easter" }

// etcDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type etcDeprecatedTimeZones string

// GMT+0 - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Etc/GMT+0'
//     Maps To Valid Time Zone: 'Etc/GMT'
//            IANA Source File: etcetera
//  
func (etcDe etcDeprecatedTimeZones) GMTPlus00() string { return "Etc/GMT" }

// GMT-0 - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Etc/GMT-0'
//     Maps To Valid Time Zone: 'Etc/GMT'
//            IANA Source File: etcetera
//  
func (etcDe etcDeprecatedTimeZones) GMTMinus00() string { return "Etc/GMT" }

// GMT0 - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Etc/GMT0'
//     Maps To Valid Time Zone: 'Etc/GMT'
//            IANA Source File: etcetera
//  
func (etcDe etcDeprecatedTimeZones) GMT00() string { return "Etc/GMT" }

// Greenwich - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Etc/Greenwich'
//     Maps To Valid Time Zone: 'Etc/GMT'
//            IANA Source File: etcetera
//  
func (etcDe etcDeprecatedTimeZones) Greenwich() string { return "Etc/GMT" }

// UCT - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Etc/UCT'
//     Maps To Valid Time Zone: 'Etc/UTC'
//            IANA Source File: backward
//  
func (etcDe etcDeprecatedTimeZones) UCT() string { return "Etc/UTC" }

// Universal - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Etc/Universal'
//     Maps To Valid Time Zone: 'Etc/UTC'
//            IANA Source File: etcetera
//  
func (etcDe etcDeprecatedTimeZones) Universal() string { return "Etc/UTC" }

// Zulu - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Etc/Zulu'
//     Maps To Valid Time Zone: 'Etc/UTC'
//            IANA Source File: etcetera
//  
func (etcDe etcDeprecatedTimeZones) Zulu() string { return "Etc/UTC" }

// europeDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type europeDeprecatedTimeZones string

// Belfast - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Belfast'
//     Maps To Valid Time Zone: 'Europe/London'
//            IANA Source File: backward
//  
func (europ europeDeprecatedTimeZones) Belfast() string { return "Europe/London" }

// Bratislava - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Bratislava'
//     Maps To Valid Time Zone: 'Europe/Prague'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Bratislava() string { return "Europe/Prague" }

// Busingen - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Busingen'
//     Maps To Valid Time Zone: 'Europe/Zurich'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Busingen() string { return "Europe/Zurich" }

// Guernsey - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Guernsey'
//     Maps To Valid Time Zone: 'Europe/London'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Guernsey() string { return "Europe/London" }

// Isle_of_Man - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Isle_of_Man'
//     Maps To Valid Time Zone: 'Europe/London'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Isle_of_Man() string { return "Europe/London" }

// Jersey - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Jersey'
//     Maps To Valid Time Zone: 'Europe/London'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Jersey() string { return "Europe/London" }

// Ljubljana - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Ljubljana'
//     Maps To Valid Time Zone: 'Europe/Belgrade'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Ljubljana() string { return "Europe/Belgrade" }

// Mariehamn - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Mariehamn'
//     Maps To Valid Time Zone: 'Europe/Helsinki'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Mariehamn() string { return "Europe/Helsinki" }

// Nicosia - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Nicosia'
//     Maps To Valid Time Zone: 'Asia/Nicosia'
//            IANA Source File: asia
//  
func (europ europeDeprecatedTimeZones) Nicosia() string { return "Asia/Nicosia" }

// Podgorica - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Podgorica'
//     Maps To Valid Time Zone: 'Europe/Belgrade'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Podgorica() string { return "Europe/Belgrade" }

// San_Marino - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/San_Marino'
//     Maps To Valid Time Zone: 'Europe/Rome'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) San_Marino() string { return "Europe/Rome" }

// Sarajevo - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Sarajevo'
//     Maps To Valid Time Zone: 'Europe/Belgrade'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Sarajevo() string { return "Europe/Belgrade" }

// Skopje - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Skopje'
//     Maps To Valid Time Zone: 'Europe/Belgrade'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Skopje() string { return "Europe/Belgrade" }

// Tiraspol - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Tiraspol'
//     Maps To Valid Time Zone: 'Europe/Chisinau'
//            IANA Source File: backward
//  
func (europ europeDeprecatedTimeZones) Tiraspol() string { return "Europe/Chisinau" }

// Vaduz - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Vaduz'
//     Maps To Valid Time Zone: 'Europe/Zurich'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Vaduz() string { return "Europe/Zurich" }

// Vatican - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Vatican'
//     Maps To Valid Time Zone: 'Europe/Rome'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Vatican() string { return "Europe/Rome" }

// Zagreb - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Europe/Zagreb'
//     Maps To Valid Time Zone: 'Europe/Belgrade'
//            IANA Source File: europe
//  
func (europ europeDeprecatedTimeZones) Zagreb() string { return "Europe/Belgrade" }

// indianDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type indianDeprecatedTimeZones string

// Antananarivo - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Indian/Antananarivo'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: africa
//  
func (india indianDeprecatedTimeZones) Antananarivo() string { return "Africa/Nairobi" }

// Comoro - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Indian/Comoro'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: africa
//  
func (india indianDeprecatedTimeZones) Comoro() string { return "Africa/Nairobi" }

// Mayotte - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Indian/Mayotte'
//     Maps To Valid Time Zone: 'Africa/Nairobi'
//            IANA Source File: africa
//  
func (india indianDeprecatedTimeZones) Mayotte() string { return "Africa/Nairobi" }

// mexicoDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type mexicoDeprecatedTimeZones string

// BajaNorte - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Mexico/BajaNorte'
//     Maps To Valid Time Zone: 'America/Tijuana'
//            IANA Source File: backward
//  
func (mexic mexicoDeprecatedTimeZones) BajaNorte() string { return "America/Tijuana" }

// BajaSur - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Mexico/BajaSur'
//     Maps To Valid Time Zone: 'America/Mazatlan'
//            IANA Source File: backward
//  
func (mexic mexicoDeprecatedTimeZones) BajaSur() string { return "America/Mazatlan" }

// General - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Mexico/General'
//     Maps To Valid Time Zone: 'America/Mexico_City'
//            IANA Source File: backward
//  
func (mexic mexicoDeprecatedTimeZones) General() string { return "America/Mexico_City" }

// pacificDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type pacificDeprecatedTimeZones string

// Johnston - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Pacific/Johnston'
//     Maps To Valid Time Zone: 'Pacific/Honolulu'
//            IANA Source File: backward
//  
func (pacif pacificDeprecatedTimeZones) Johnston() string { return "Pacific/Honolulu" }

// Midway - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Pacific/Midway'
//     Maps To Valid Time Zone: 'Pacific/Pago_Pago'
//            IANA Source File: australasia
//  
func (pacif pacificDeprecatedTimeZones) Midway() string { return "Pacific/Pago_Pago" }

// Ponape - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Pacific/Ponape'
//     Maps To Valid Time Zone: 'Pacific/Pohnpei'
//            IANA Source File: backward
//  
func (pacif pacificDeprecatedTimeZones) Ponape() string { return "Pacific/Pohnpei" }

// Saipan - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Pacific/Saipan'
//     Maps To Valid Time Zone: 'Pacific/Guam'
//            IANA Source File: australasia
//  
func (pacif pacificDeprecatedTimeZones) Saipan() string { return "Pacific/Guam" }

// Samoa - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Pacific/Samoa'
//     Maps To Valid Time Zone: 'Pacific/Pago_Pago'
//            IANA Source File: backward
//  
func (pacif pacificDeprecatedTimeZones) Samoa() string { return "Pacific/Pago_Pago" }

// Truk - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Pacific/Truk'
//     Maps To Valid Time Zone: 'Pacific/Chuuk'
//            IANA Source File: backward
//  
func (pacif pacificDeprecatedTimeZones) Truk() string { return "Pacific/Chuuk" }

// Yap - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Pacific/Yap'
//     Maps To Valid Time Zone: 'Pacific/Chuuk'
//            IANA Source File: backward
//  
func (pacif pacificDeprecatedTimeZones) Yap() string { return "Pacific/Chuuk" }

// uSDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type uSDeprecatedTimeZones string

// Alaska - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Alaska'
//     Maps To Valid Time Zone: 'America/Anchorage'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Alaska() string { return "America/Anchorage" }

// Aleutian - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Aleutian'
//     Maps To Valid Time Zone: 'America/Adak'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Aleutian() string { return "America/Adak" }

// Arizona - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Arizona'
//     Maps To Valid Time Zone: 'America/Phoenix'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Arizona() string { return "America/Phoenix" }

// Central - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Central'
//     Maps To Valid Time Zone: 'America/Chicago'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Central() string { return "America/Chicago" }

// East-Indiana - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/East-Indiana'
//     Maps To Valid Time Zone: 'America/Indiana/Indianapolis'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) EastMinusIndiana() string { return "America/Indiana/Indianapolis" }

// Eastern - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Eastern'
//     Maps To Valid Time Zone: 'America/New_York'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Eastern() string { return "America/New_York" }

// Hawaii - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Hawaii'
//     Maps To Valid Time Zone: 'Pacific/Honolulu'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Hawaii() string { return "Pacific/Honolulu" }

// Indiana-Starke - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Indiana-Starke'
//     Maps To Valid Time Zone: 'America/Indiana/Knox'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) IndianaMinusStarke() string { return "America/Indiana/Knox" }

// Michigan - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Michigan'
//     Maps To Valid Time Zone: 'America/Detroit'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Michigan() string { return "America/Detroit" }

// Mountain - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Mountain'
//     Maps To Valid Time Zone: 'America/Denver'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Mountain() string { return "America/Denver" }

// Pacific - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Pacific'
//     Maps To Valid Time Zone: 'America/Los_Angeles'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Pacific() string { return "America/Los_Angeles" }

// Pacific-New - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Pacific-New'
//     Maps To Valid Time Zone: 'America/Los_Angeles'
//            IANA Source File: pacificnew
//  
func (uSDep uSDeprecatedTimeZones) PacificMinusNew() string { return "America/Los_Angeles" }

// Samoa - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'US/Samoa'
//     Maps To Valid Time Zone: 'Pacific/Pago_Pago'
//            IANA Source File: backward
//  
func (uSDep uSDeprecatedTimeZones) Samoa() string { return "Pacific/Pago_Pago" }

// argentinaDeprecatedTimeZones - Defines a subsidiary collection
// of Time Zones which are obsolete and no longer used as
// primary and accepted time zone designations. These
// time zones are classified as 'Link' Time Zones and
// are mapped to current, valid IANA Time Zones.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type argentinaDeprecatedTimeZones string

// ComodRivadavia - This is an IANA 'Link' Time Zone. 'Link' Time Zones
// identify deprecated or obsolete time zones. These obsolete
// time zones are mapped to valid current time zones.
//  
// Linked Deprecated Time Zone: 'Argentina/ComodRivadavia'
//     Maps To Valid Time Zone: 'America/Argentina/Catamarca'
//            IANA Source File: backward
//  
func (argen argentinaDeprecatedTimeZones) ComodRivadavia() string { return "America/Argentina/Catamarca" }

