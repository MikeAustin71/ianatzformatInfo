package main



// TimeZones - This type and its associated methods encapsulate 593 IANA Time
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
// 'TimeZones' type. It is therefore much easier to access any of the 618 time
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
// Africa                                    54        0       54
// America                                  166        0      166
// Antarctica                                12        0       12
// Asia                                      99        0       99
// Atlantic                                  12        0       12
// Australia                                 23        0       23
// Europe                                    63        0       63
// Indian                                    11        0       11
// Pacific                                   43        0       43
// Etc                                       35        0       35
// Other                                     75        0       75
// ==============================================================
//                              Total       593        0      593
//
// ----------------------------------------------------------------------------
// 
// This TimeZones Type is based on IANA Time Zone Database Version: 2019c
// 
//           IANA Standard Time Zones : 593
//           IANA Link Time Zones     :   0
//                                         -------
//                 Sub-Total IANA Time Zones: 593
// 
//                Military Time Zones :  25
//                   Other Time Zones :   0
//                                         -------
//                          Total Time Zones: 618
// 
//       Standard Time Zone Sub-Groups:   4
//           Link Time Zone Sub-Groups:   0
//                                         -------
//                Total Time Zone Sub-Groups:   4
// 
//                  Primary Time Zone Groups:  17
// 
// Type Creation Date: 2019-11-09 Saturday 16:53:16 -0600 CST
// ----------------------------------------------------------------------------
// 
type TimeZones struct {
    Africa                             africaTimeZones
    America                            americaTimeZones
    Antarctica                         antarcticaTimeZones
    Arctic                             arcticTimeZones
    Asia                               asiaTimeZones
    Atlantic                           atlanticTimeZones
    Australia                          australiaTimeZones
    Brazil                             brazilTimeZones
    Canada                             canadaTimeZones
    Chile                              chileTimeZones
    Etc                                etcTimeZones
    Europe                             europeTimeZones
    Indian                             indianTimeZones
    Mexico                             mexicoTimeZones
    Military                           militaryTimeZones
    Other                              otherTimeZones
    Pacific                            pacificTimeZones
    US                                 uSTimeZones
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
// IANA Source File: Abidjan
//  
func (afric africaTimeZones) Abidjan() string {return "Africa/Abidjan" }

// Accra - IANA Time Zone 'Africa/Accra'.
// IANA Source File: Accra
//  
func (afric africaTimeZones) Accra() string {return "Africa/Accra" }

// Addis_Ababa - IANA Time Zone 'Africa/Addis_Ababa'.
// IANA Source File: Addis_Ababa
//  
func (afric africaTimeZones) Addis_Ababa() string {return "Africa/Addis_Ababa" }

// Algiers - IANA Time Zone 'Africa/Algiers'.
// IANA Source File: Algiers
//  
func (afric africaTimeZones) Algiers() string {return "Africa/Algiers" }

// Asmara - IANA Time Zone 'Africa/Asmara'.
// IANA Source File: Asmara
//  
func (afric africaTimeZones) Asmara() string {return "Africa/Asmara" }

// Asmera - IANA Time Zone 'Africa/Asmera'.
// IANA Source File: Asmera
//  
func (afric africaTimeZones) Asmera() string {return "Africa/Asmera" }

// Bamako - IANA Time Zone 'Africa/Bamako'.
// IANA Source File: Bamako
//  
func (afric africaTimeZones) Bamako() string {return "Africa/Bamako" }

// Bangui - IANA Time Zone 'Africa/Bangui'.
// IANA Source File: Bangui
//  
func (afric africaTimeZones) Bangui() string {return "Africa/Bangui" }

// Banjul - IANA Time Zone 'Africa/Banjul'.
// IANA Source File: Banjul
//  
func (afric africaTimeZones) Banjul() string {return "Africa/Banjul" }

// Bissau - IANA Time Zone 'Africa/Bissau'.
// IANA Source File: Bissau
//  
func (afric africaTimeZones) Bissau() string {return "Africa/Bissau" }

// Blantyre - IANA Time Zone 'Africa/Blantyre'.
// IANA Source File: Blantyre
//  
func (afric africaTimeZones) Blantyre() string {return "Africa/Blantyre" }

// Brazzaville - IANA Time Zone 'Africa/Brazzaville'.
// IANA Source File: Brazzaville
//  
func (afric africaTimeZones) Brazzaville() string {return "Africa/Brazzaville" }

// Bujumbura - IANA Time Zone 'Africa/Bujumbura'.
// IANA Source File: Bujumbura
//  
func (afric africaTimeZones) Bujumbura() string {return "Africa/Bujumbura" }

// Cairo - IANA Time Zone 'Africa/Cairo'.
// IANA Source File: Cairo
//  
func (afric africaTimeZones) Cairo() string {return "Africa/Cairo" }

// Casablanca - IANA Time Zone 'Africa/Casablanca'.
// IANA Source File: Casablanca
//  
func (afric africaTimeZones) Casablanca() string {return "Africa/Casablanca" }

// Ceuta - IANA Time Zone 'Africa/Ceuta'.
// IANA Source File: Ceuta
//  
func (afric africaTimeZones) Ceuta() string {return "Africa/Ceuta" }

// Conakry - IANA Time Zone 'Africa/Conakry'.
// IANA Source File: Conakry
//  
func (afric africaTimeZones) Conakry() string {return "Africa/Conakry" }

// Dakar - IANA Time Zone 'Africa/Dakar'.
// IANA Source File: Dakar
//  
func (afric africaTimeZones) Dakar() string {return "Africa/Dakar" }

// Dar_es_Salaam - IANA Time Zone 'Africa/Dar_es_Salaam'.
// IANA Source File: Dar_es_Salaam
//  
func (afric africaTimeZones) Dar_es_Salaam() string {return "Africa/Dar_es_Salaam" }

// Djibouti - IANA Time Zone 'Africa/Djibouti'.
// IANA Source File: Djibouti
//  
func (afric africaTimeZones) Djibouti() string {return "Africa/Djibouti" }

// Douala - IANA Time Zone 'Africa/Douala'.
// IANA Source File: Douala
//  
func (afric africaTimeZones) Douala() string {return "Africa/Douala" }

// El_Aaiun - IANA Time Zone 'Africa/El_Aaiun'.
// IANA Source File: El_Aaiun
//  
func (afric africaTimeZones) El_Aaiun() string {return "Africa/El_Aaiun" }

// Freetown - IANA Time Zone 'Africa/Freetown'.
// IANA Source File: Freetown
//  
func (afric africaTimeZones) Freetown() string {return "Africa/Freetown" }

// Gaborone - IANA Time Zone 'Africa/Gaborone'.
// IANA Source File: Gaborone
//  
func (afric africaTimeZones) Gaborone() string {return "Africa/Gaborone" }

// Harare - IANA Time Zone 'Africa/Harare'.
// IANA Source File: Harare
//  
func (afric africaTimeZones) Harare() string {return "Africa/Harare" }

// Johannesburg - IANA Time Zone 'Africa/Johannesburg'.
// IANA Source File: Johannesburg
//  
func (afric africaTimeZones) Johannesburg() string {return "Africa/Johannesburg" }

// Juba - IANA Time Zone 'Africa/Juba'.
// IANA Source File: Juba
//  
func (afric africaTimeZones) Juba() string {return "Africa/Juba" }

// Kampala - IANA Time Zone 'Africa/Kampala'.
// IANA Source File: Kampala
//  
func (afric africaTimeZones) Kampala() string {return "Africa/Kampala" }

// Khartoum - IANA Time Zone 'Africa/Khartoum'.
// IANA Source File: Khartoum
//  
func (afric africaTimeZones) Khartoum() string {return "Africa/Khartoum" }

// Kigali - IANA Time Zone 'Africa/Kigali'.
// IANA Source File: Kigali
//  
func (afric africaTimeZones) Kigali() string {return "Africa/Kigali" }

// Kinshasa - IANA Time Zone 'Africa/Kinshasa'.
// IANA Source File: Kinshasa
//  
func (afric africaTimeZones) Kinshasa() string {return "Africa/Kinshasa" }

// Lagos - IANA Time Zone 'Africa/Lagos'.
// IANA Source File: Lagos
//  
func (afric africaTimeZones) Lagos() string {return "Africa/Lagos" }

// Libreville - IANA Time Zone 'Africa/Libreville'.
// IANA Source File: Libreville
//  
func (afric africaTimeZones) Libreville() string {return "Africa/Libreville" }

// Lome - IANA Time Zone 'Africa/Lome'.
// IANA Source File: Lome
//  
func (afric africaTimeZones) Lome() string {return "Africa/Lome" }

// Luanda - IANA Time Zone 'Africa/Luanda'.
// IANA Source File: Luanda
//  
func (afric africaTimeZones) Luanda() string {return "Africa/Luanda" }

// Lubumbashi - IANA Time Zone 'Africa/Lubumbashi'.
// IANA Source File: Lubumbashi
//  
func (afric africaTimeZones) Lubumbashi() string {return "Africa/Lubumbashi" }

// Lusaka - IANA Time Zone 'Africa/Lusaka'.
// IANA Source File: Lusaka
//  
func (afric africaTimeZones) Lusaka() string {return "Africa/Lusaka" }

// Malabo - IANA Time Zone 'Africa/Malabo'.
// IANA Source File: Malabo
//  
func (afric africaTimeZones) Malabo() string {return "Africa/Malabo" }

// Maputo - IANA Time Zone 'Africa/Maputo'.
// IANA Source File: Maputo
//  
func (afric africaTimeZones) Maputo() string {return "Africa/Maputo" }

// Maseru - IANA Time Zone 'Africa/Maseru'.
// IANA Source File: Maseru
//  
func (afric africaTimeZones) Maseru() string {return "Africa/Maseru" }

// Mbabane - IANA Time Zone 'Africa/Mbabane'.
// IANA Source File: Mbabane
//  
func (afric africaTimeZones) Mbabane() string {return "Africa/Mbabane" }

// Mogadishu - IANA Time Zone 'Africa/Mogadishu'.
// IANA Source File: Mogadishu
//  
func (afric africaTimeZones) Mogadishu() string {return "Africa/Mogadishu" }

// Monrovia - IANA Time Zone 'Africa/Monrovia'.
// IANA Source File: Monrovia
//  
func (afric africaTimeZones) Monrovia() string {return "Africa/Monrovia" }

// Nairobi - IANA Time Zone 'Africa/Nairobi'.
// IANA Source File: Nairobi
//  
func (afric africaTimeZones) Nairobi() string {return "Africa/Nairobi" }

// Ndjamena - IANA Time Zone 'Africa/Ndjamena'.
// IANA Source File: Ndjamena
//  
func (afric africaTimeZones) Ndjamena() string {return "Africa/Ndjamena" }

// Niamey - IANA Time Zone 'Africa/Niamey'.
// IANA Source File: Niamey
//  
func (afric africaTimeZones) Niamey() string {return "Africa/Niamey" }

// Nouakchott - IANA Time Zone 'Africa/Nouakchott'.
// IANA Source File: Nouakchott
//  
func (afric africaTimeZones) Nouakchott() string {return "Africa/Nouakchott" }

// Ouagadougou - IANA Time Zone 'Africa/Ouagadougou'.
// IANA Source File: Ouagadougou
//  
func (afric africaTimeZones) Ouagadougou() string {return "Africa/Ouagadougou" }

// Porto-Novo - IANA Time Zone 'Africa/Porto-Novo'.
// IANA Source File: Porto-Novo
//  
func (afric africaTimeZones) PortoMinusNovo() string {return "Africa/Porto-Novo" }

// Sao_Tome - IANA Time Zone 'Africa/Sao_Tome'.
// IANA Source File: Sao_Tome
//  
func (afric africaTimeZones) Sao_Tome() string {return "Africa/Sao_Tome" }

// Timbuktu - IANA Time Zone 'Africa/Timbuktu'.
// IANA Source File: Timbuktu
//  
func (afric africaTimeZones) Timbuktu() string {return "Africa/Timbuktu" }

// Tripoli - IANA Time Zone 'Africa/Tripoli'.
// IANA Source File: Tripoli
//  
func (afric africaTimeZones) Tripoli() string {return "Africa/Tripoli" }

// Tunis - IANA Time Zone 'Africa/Tunis'.
// IANA Source File: Tunis
//  
func (afric africaTimeZones) Tunis() string {return "Africa/Tunis" }

// Windhoek - IANA Time Zone 'Africa/Windhoek'.
// IANA Source File: Windhoek
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
// IANA Source File: Adak
//  
func (ameri americaTimeZones) Adak() string {return "America/Adak" }

// Anchorage - IANA Time Zone 'America/Anchorage'.
// IANA Source File: Anchorage
//  
func (ameri americaTimeZones) Anchorage() string {return "America/Anchorage" }

// Anguilla - IANA Time Zone 'America/Anguilla'.
// IANA Source File: Anguilla
//  
func (ameri americaTimeZones) Anguilla() string {return "America/Anguilla" }

// Antigua - IANA Time Zone 'America/Antigua'.
// IANA Source File: Antigua
//  
func (ameri americaTimeZones) Antigua() string {return "America/Antigua" }

// Araguaina - IANA Time Zone 'America/Araguaina'.
// IANA Source File: Araguaina
//  
func (ameri americaTimeZones) Araguaina() string {return "America/Araguaina" }

// Aruba - IANA Time Zone 'America/Aruba'.
// IANA Source File: Aruba
//  
func (ameri americaTimeZones) Aruba() string {return "America/Aruba" }

// Asuncion - IANA Time Zone 'America/Asuncion'.
// IANA Source File: Asuncion
//  
func (ameri americaTimeZones) Asuncion() string {return "America/Asuncion" }

// Atikokan - IANA Time Zone 'America/Atikokan'.
// IANA Source File: Atikokan
//  
func (ameri americaTimeZones) Atikokan() string {return "America/Atikokan" }

// Atka - IANA Time Zone 'America/Atka'.
// IANA Source File: Atka
//  
func (ameri americaTimeZones) Atka() string {return "America/Atka" }

// Bahia - IANA Time Zone 'America/Bahia'.
// IANA Source File: Bahia
//  
func (ameri americaTimeZones) Bahia() string {return "America/Bahia" }

// Bahia_Banderas - IANA Time Zone 'America/Bahia_Banderas'.
// IANA Source File: Bahia_Banderas
//  
func (ameri americaTimeZones) Bahia_Banderas() string {return "America/Bahia_Banderas" }

// Barbados - IANA Time Zone 'America/Barbados'.
// IANA Source File: Barbados
//  
func (ameri americaTimeZones) Barbados() string {return "America/Barbados" }

// Belem - IANA Time Zone 'America/Belem'.
// IANA Source File: Belem
//  
func (ameri americaTimeZones) Belem() string {return "America/Belem" }

// Belize - IANA Time Zone 'America/Belize'.
// IANA Source File: Belize
//  
func (ameri americaTimeZones) Belize() string {return "America/Belize" }

// Blanc-Sablon - IANA Time Zone 'America/Blanc-Sablon'.
// IANA Source File: Blanc-Sablon
//  
func (ameri americaTimeZones) BlancMinusSablon() string {return "America/Blanc-Sablon" }

// Boa_Vista - IANA Time Zone 'America/Boa_Vista'.
// IANA Source File: Boa_Vista
//  
func (ameri americaTimeZones) Boa_Vista() string {return "America/Boa_Vista" }

// Bogota - IANA Time Zone 'America/Bogota'.
// IANA Source File: Bogota
//  
func (ameri americaTimeZones) Bogota() string {return "America/Bogota" }

// Boise - IANA Time Zone 'America/Boise'.
// IANA Source File: Boise
//  
func (ameri americaTimeZones) Boise() string {return "America/Boise" }

// Buenos_Aires - IANA Time Zone 'America/Buenos_Aires'.
// IANA Source File: Buenos_Aires
//  
func (ameri americaTimeZones) Buenos_Aires() string {return "America/Buenos_Aires" }

// Cambridge_Bay - IANA Time Zone 'America/Cambridge_Bay'.
// IANA Source File: Cambridge_Bay
//  
func (ameri americaTimeZones) Cambridge_Bay() string {return "America/Cambridge_Bay" }

// Campo_Grande - IANA Time Zone 'America/Campo_Grande'.
// IANA Source File: Campo_Grande
//  
func (ameri americaTimeZones) Campo_Grande() string {return "America/Campo_Grande" }

// Cancun - IANA Time Zone 'America/Cancun'.
// IANA Source File: Cancun
//  
func (ameri americaTimeZones) Cancun() string {return "America/Cancun" }

// Caracas - IANA Time Zone 'America/Caracas'.
// IANA Source File: Caracas
//  
func (ameri americaTimeZones) Caracas() string {return "America/Caracas" }

// Catamarca - IANA Time Zone 'America/Catamarca'.
// IANA Source File: Catamarca
//  
func (ameri americaTimeZones) Catamarca() string {return "America/Catamarca" }

// Cayenne - IANA Time Zone 'America/Cayenne'.
// IANA Source File: Cayenne
//  
func (ameri americaTimeZones) Cayenne() string {return "America/Cayenne" }

// Cayman - IANA Time Zone 'America/Cayman'.
// IANA Source File: Cayman
//  
func (ameri americaTimeZones) Cayman() string {return "America/Cayman" }

// Chicago - IANA Time Zone 'America/Chicago'.
// IANA Source File: Chicago
//  
func (ameri americaTimeZones) Chicago() string {return "America/Chicago" }

// Chihuahua - IANA Time Zone 'America/Chihuahua'.
// IANA Source File: Chihuahua
//  
func (ameri americaTimeZones) Chihuahua() string {return "America/Chihuahua" }

// Coral_Harbour - IANA Time Zone 'America/Coral_Harbour'.
// IANA Source File: Coral_Harbour
//  
func (ameri americaTimeZones) Coral_Harbour() string {return "America/Coral_Harbour" }

// Cordoba - IANA Time Zone 'America/Cordoba'.
// IANA Source File: Cordoba
//  
func (ameri americaTimeZones) Cordoba() string {return "America/Cordoba" }

// Costa_Rica - IANA Time Zone 'America/Costa_Rica'.
// IANA Source File: Costa_Rica
//  
func (ameri americaTimeZones) Costa_Rica() string {return "America/Costa_Rica" }

// Creston - IANA Time Zone 'America/Creston'.
// IANA Source File: Creston
//  
func (ameri americaTimeZones) Creston() string {return "America/Creston" }

// Cuiaba - IANA Time Zone 'America/Cuiaba'.
// IANA Source File: Cuiaba
//  
func (ameri americaTimeZones) Cuiaba() string {return "America/Cuiaba" }

// Curacao - IANA Time Zone 'America/Curacao'.
// IANA Source File: Curacao
//  
func (ameri americaTimeZones) Curacao() string {return "America/Curacao" }

// Danmarkshavn - IANA Time Zone 'America/Danmarkshavn'.
// IANA Source File: Danmarkshavn
//  
func (ameri americaTimeZones) Danmarkshavn() string {return "America/Danmarkshavn" }

// Dawson - IANA Time Zone 'America/Dawson'.
// IANA Source File: Dawson
//  
func (ameri americaTimeZones) Dawson() string {return "America/Dawson" }

// Dawson_Creek - IANA Time Zone 'America/Dawson_Creek'.
// IANA Source File: Dawson_Creek
//  
func (ameri americaTimeZones) Dawson_Creek() string {return "America/Dawson_Creek" }

// Denver - IANA Time Zone 'America/Denver'.
// IANA Source File: Denver
//  
func (ameri americaTimeZones) Denver() string {return "America/Denver" }

// Detroit - IANA Time Zone 'America/Detroit'.
// IANA Source File: Detroit
//  
func (ameri americaTimeZones) Detroit() string {return "America/Detroit" }

// Dominica - IANA Time Zone 'America/Dominica'.
// IANA Source File: Dominica
//  
func (ameri americaTimeZones) Dominica() string {return "America/Dominica" }

// Edmonton - IANA Time Zone 'America/Edmonton'.
// IANA Source File: Edmonton
//  
func (ameri americaTimeZones) Edmonton() string {return "America/Edmonton" }

// Eirunepe - IANA Time Zone 'America/Eirunepe'.
// IANA Source File: Eirunepe
//  
func (ameri americaTimeZones) Eirunepe() string {return "America/Eirunepe" }

// El_Salvador - IANA Time Zone 'America/El_Salvador'.
// IANA Source File: El_Salvador
//  
func (ameri americaTimeZones) El_Salvador() string {return "America/El_Salvador" }

// Ensenada - IANA Time Zone 'America/Ensenada'.
// IANA Source File: Ensenada
//  
func (ameri americaTimeZones) Ensenada() string {return "America/Ensenada" }

// Fort_Nelson - IANA Time Zone 'America/Fort_Nelson'.
// IANA Source File: Fort_Nelson
//  
func (ameri americaTimeZones) Fort_Nelson() string {return "America/Fort_Nelson" }

// Fort_Wayne - IANA Time Zone 'America/Fort_Wayne'.
// IANA Source File: Fort_Wayne
//  
func (ameri americaTimeZones) Fort_Wayne() string {return "America/Fort_Wayne" }

// Fortaleza - IANA Time Zone 'America/Fortaleza'.
// IANA Source File: Fortaleza
//  
func (ameri americaTimeZones) Fortaleza() string {return "America/Fortaleza" }

// Glace_Bay - IANA Time Zone 'America/Glace_Bay'.
// IANA Source File: Glace_Bay
//  
func (ameri americaTimeZones) Glace_Bay() string {return "America/Glace_Bay" }

// Godthab - IANA Time Zone 'America/Godthab'.
// IANA Source File: Godthab
//  
func (ameri americaTimeZones) Godthab() string {return "America/Godthab" }

// Goose_Bay - IANA Time Zone 'America/Goose_Bay'.
// IANA Source File: Goose_Bay
//  
func (ameri americaTimeZones) Goose_Bay() string {return "America/Goose_Bay" }

// Grand_Turk - IANA Time Zone 'America/Grand_Turk'.
// IANA Source File: Grand_Turk
//  
func (ameri americaTimeZones) Grand_Turk() string {return "America/Grand_Turk" }

// Grenada - IANA Time Zone 'America/Grenada'.
// IANA Source File: Grenada
//  
func (ameri americaTimeZones) Grenada() string {return "America/Grenada" }

// Guadeloupe - IANA Time Zone 'America/Guadeloupe'.
// IANA Source File: Guadeloupe
//  
func (ameri americaTimeZones) Guadeloupe() string {return "America/Guadeloupe" }

// Guatemala - IANA Time Zone 'America/Guatemala'.
// IANA Source File: Guatemala
//  
func (ameri americaTimeZones) Guatemala() string {return "America/Guatemala" }

// Guayaquil - IANA Time Zone 'America/Guayaquil'.
// IANA Source File: Guayaquil
//  
func (ameri americaTimeZones) Guayaquil() string {return "America/Guayaquil" }

// Guyana - IANA Time Zone 'America/Guyana'.
// IANA Source File: Guyana
//  
func (ameri americaTimeZones) Guyana() string {return "America/Guyana" }

// Halifax - IANA Time Zone 'America/Halifax'.
// IANA Source File: Halifax
//  
func (ameri americaTimeZones) Halifax() string {return "America/Halifax" }

// Havana - IANA Time Zone 'America/Havana'.
// IANA Source File: Havana
//  
func (ameri americaTimeZones) Havana() string {return "America/Havana" }

// Hermosillo - IANA Time Zone 'America/Hermosillo'.
// IANA Source File: Hermosillo
//  
func (ameri americaTimeZones) Hermosillo() string {return "America/Hermosillo" }

// Indianapolis - IANA Time Zone 'America/Indianapolis'.
// IANA Source File: Indianapolis
//  
func (ameri americaTimeZones) Indianapolis() string {return "America/Indianapolis" }

// Inuvik - IANA Time Zone 'America/Inuvik'.
// IANA Source File: Inuvik
//  
func (ameri americaTimeZones) Inuvik() string {return "America/Inuvik" }

// Iqaluit - IANA Time Zone 'America/Iqaluit'.
// IANA Source File: Iqaluit
//  
func (ameri americaTimeZones) Iqaluit() string {return "America/Iqaluit" }

// Jamaica - IANA Time Zone 'America/Jamaica'.
// IANA Source File: Jamaica
//  
func (ameri americaTimeZones) Jamaica() string {return "America/Jamaica" }

// Jujuy - IANA Time Zone 'America/Jujuy'.
// IANA Source File: Jujuy
//  
func (ameri americaTimeZones) Jujuy() string {return "America/Jujuy" }

// Juneau - IANA Time Zone 'America/Juneau'.
// IANA Source File: Juneau
//  
func (ameri americaTimeZones) Juneau() string {return "America/Juneau" }

// Knox_IN - IANA Time Zone 'America/Knox_IN'.
// IANA Source File: Knox_IN
//  
func (ameri americaTimeZones) Knox_IN() string {return "America/Knox_IN" }

// Kralendijk - IANA Time Zone 'America/Kralendijk'.
// IANA Source File: Kralendijk
//  
func (ameri americaTimeZones) Kralendijk() string {return "America/Kralendijk" }

// La_Paz - IANA Time Zone 'America/La_Paz'.
// IANA Source File: La_Paz
//  
func (ameri americaTimeZones) La_Paz() string {return "America/La_Paz" }

// Lima - IANA Time Zone 'America/Lima'.
// IANA Source File: Lima
//  
func (ameri americaTimeZones) Lima() string {return "America/Lima" }

// Los_Angeles - IANA Time Zone 'America/Los_Angeles'.
// IANA Source File: Los_Angeles
//  
func (ameri americaTimeZones) Los_Angeles() string {return "America/Los_Angeles" }

// Louisville - IANA Time Zone 'America/Louisville'.
// IANA Source File: Louisville
//  
func (ameri americaTimeZones) Louisville() string {return "America/Louisville" }

// Lower_Princes - IANA Time Zone 'America/Lower_Princes'.
// IANA Source File: Lower_Princes
//  
func (ameri americaTimeZones) Lower_Princes() string {return "America/Lower_Princes" }

// Maceio - IANA Time Zone 'America/Maceio'.
// IANA Source File: Maceio
//  
func (ameri americaTimeZones) Maceio() string {return "America/Maceio" }

// Managua - IANA Time Zone 'America/Managua'.
// IANA Source File: Managua
//  
func (ameri americaTimeZones) Managua() string {return "America/Managua" }

// Manaus - IANA Time Zone 'America/Manaus'.
// IANA Source File: Manaus
//  
func (ameri americaTimeZones) Manaus() string {return "America/Manaus" }

// Marigot - IANA Time Zone 'America/Marigot'.
// IANA Source File: Marigot
//  
func (ameri americaTimeZones) Marigot() string {return "America/Marigot" }

// Martinique - IANA Time Zone 'America/Martinique'.
// IANA Source File: Martinique
//  
func (ameri americaTimeZones) Martinique() string {return "America/Martinique" }

// Matamoros - IANA Time Zone 'America/Matamoros'.
// IANA Source File: Matamoros
//  
func (ameri americaTimeZones) Matamoros() string {return "America/Matamoros" }

// Mazatlan - IANA Time Zone 'America/Mazatlan'.
// IANA Source File: Mazatlan
//  
func (ameri americaTimeZones) Mazatlan() string {return "America/Mazatlan" }

// Mendoza - IANA Time Zone 'America/Mendoza'.
// IANA Source File: Mendoza
//  
func (ameri americaTimeZones) Mendoza() string {return "America/Mendoza" }

// Menominee - IANA Time Zone 'America/Menominee'.
// IANA Source File: Menominee
//  
func (ameri americaTimeZones) Menominee() string {return "America/Menominee" }

// Merida - IANA Time Zone 'America/Merida'.
// IANA Source File: Merida
//  
func (ameri americaTimeZones) Merida() string {return "America/Merida" }

// Metlakatla - IANA Time Zone 'America/Metlakatla'.
// IANA Source File: Metlakatla
//  
func (ameri americaTimeZones) Metlakatla() string {return "America/Metlakatla" }

// Mexico_City - IANA Time Zone 'America/Mexico_City'.
// IANA Source File: Mexico_City
//  
func (ameri americaTimeZones) Mexico_City() string {return "America/Mexico_City" }

// Miquelon - IANA Time Zone 'America/Miquelon'.
// IANA Source File: Miquelon
//  
func (ameri americaTimeZones) Miquelon() string {return "America/Miquelon" }

// Moncton - IANA Time Zone 'America/Moncton'.
// IANA Source File: Moncton
//  
func (ameri americaTimeZones) Moncton() string {return "America/Moncton" }

// Monterrey - IANA Time Zone 'America/Monterrey'.
// IANA Source File: Monterrey
//  
func (ameri americaTimeZones) Monterrey() string {return "America/Monterrey" }

// Montevideo - IANA Time Zone 'America/Montevideo'.
// IANA Source File: Montevideo
//  
func (ameri americaTimeZones) Montevideo() string {return "America/Montevideo" }

// Montreal - IANA Time Zone 'America/Montreal'.
// IANA Source File: Montreal
//  
func (ameri americaTimeZones) Montreal() string {return "America/Montreal" }

// Montserrat - IANA Time Zone 'America/Montserrat'.
// IANA Source File: Montserrat
//  
func (ameri americaTimeZones) Montserrat() string {return "America/Montserrat" }

// Nassau - IANA Time Zone 'America/Nassau'.
// IANA Source File: Nassau
//  
func (ameri americaTimeZones) Nassau() string {return "America/Nassau" }

// New_York - IANA Time Zone 'America/New_York'.
// IANA Source File: New_York
//  
func (ameri americaTimeZones) New_York() string {return "America/New_York" }

// Nipigon - IANA Time Zone 'America/Nipigon'.
// IANA Source File: Nipigon
//  
func (ameri americaTimeZones) Nipigon() string {return "America/Nipigon" }

// Nome - IANA Time Zone 'America/Nome'.
// IANA Source File: Nome
//  
func (ameri americaTimeZones) Nome() string {return "America/Nome" }

// Noronha - IANA Time Zone 'America/Noronha'.
// IANA Source File: Noronha
//  
func (ameri americaTimeZones) Noronha() string {return "America/Noronha" }

// Ojinaga - IANA Time Zone 'America/Ojinaga'.
// IANA Source File: Ojinaga
//  
func (ameri americaTimeZones) Ojinaga() string {return "America/Ojinaga" }

// Panama - IANA Time Zone 'America/Panama'.
// IANA Source File: Panama
//  
func (ameri americaTimeZones) Panama() string {return "America/Panama" }

// Pangnirtung - IANA Time Zone 'America/Pangnirtung'.
// IANA Source File: Pangnirtung
//  
func (ameri americaTimeZones) Pangnirtung() string {return "America/Pangnirtung" }

// Paramaribo - IANA Time Zone 'America/Paramaribo'.
// IANA Source File: Paramaribo
//  
func (ameri americaTimeZones) Paramaribo() string {return "America/Paramaribo" }

// Phoenix - IANA Time Zone 'America/Phoenix'.
// IANA Source File: Phoenix
//  
func (ameri americaTimeZones) Phoenix() string {return "America/Phoenix" }

// Port-au-Prince - IANA Time Zone 'America/Port-au-Prince'.
// IANA Source File: Port-au-Prince
//  
func (ameri americaTimeZones) PortMinusauMinusPrince() string {return "America/Port-au-Prince" }

// Port_of_Spain - IANA Time Zone 'America/Port_of_Spain'.
// IANA Source File: Port_of_Spain
//  
func (ameri americaTimeZones) Port_of_Spain() string {return "America/Port_of_Spain" }

// Porto_Acre - IANA Time Zone 'America/Porto_Acre'.
// IANA Source File: Porto_Acre
//  
func (ameri americaTimeZones) Porto_Acre() string {return "America/Porto_Acre" }

// Porto_Velho - IANA Time Zone 'America/Porto_Velho'.
// IANA Source File: Porto_Velho
//  
func (ameri americaTimeZones) Porto_Velho() string {return "America/Porto_Velho" }

// Puerto_Rico - IANA Time Zone 'America/Puerto_Rico'.
// IANA Source File: Puerto_Rico
//  
func (ameri americaTimeZones) Puerto_Rico() string {return "America/Puerto_Rico" }

// Punta_Arenas - IANA Time Zone 'America/Punta_Arenas'.
// IANA Source File: Punta_Arenas
//  
func (ameri americaTimeZones) Punta_Arenas() string {return "America/Punta_Arenas" }

// Rainy_River - IANA Time Zone 'America/Rainy_River'.
// IANA Source File: Rainy_River
//  
func (ameri americaTimeZones) Rainy_River() string {return "America/Rainy_River" }

// Rankin_Inlet - IANA Time Zone 'America/Rankin_Inlet'.
// IANA Source File: Rankin_Inlet
//  
func (ameri americaTimeZones) Rankin_Inlet() string {return "America/Rankin_Inlet" }

// Recife - IANA Time Zone 'America/Recife'.
// IANA Source File: Recife
//  
func (ameri americaTimeZones) Recife() string {return "America/Recife" }

// Regina - IANA Time Zone 'America/Regina'.
// IANA Source File: Regina
//  
func (ameri americaTimeZones) Regina() string {return "America/Regina" }

// Resolute - IANA Time Zone 'America/Resolute'.
// IANA Source File: Resolute
//  
func (ameri americaTimeZones) Resolute() string {return "America/Resolute" }

// Rio_Branco - IANA Time Zone 'America/Rio_Branco'.
// IANA Source File: Rio_Branco
//  
func (ameri americaTimeZones) Rio_Branco() string {return "America/Rio_Branco" }

// Rosario - IANA Time Zone 'America/Rosario'.
// IANA Source File: Rosario
//  
func (ameri americaTimeZones) Rosario() string {return "America/Rosario" }

// Santa_Isabel - IANA Time Zone 'America/Santa_Isabel'.
// IANA Source File: Santa_Isabel
//  
func (ameri americaTimeZones) Santa_Isabel() string {return "America/Santa_Isabel" }

// Santarem - IANA Time Zone 'America/Santarem'.
// IANA Source File: Santarem
//  
func (ameri americaTimeZones) Santarem() string {return "America/Santarem" }

// Santiago - IANA Time Zone 'America/Santiago'.
// IANA Source File: Santiago
//  
func (ameri americaTimeZones) Santiago() string {return "America/Santiago" }

// Santo_Domingo - IANA Time Zone 'America/Santo_Domingo'.
// IANA Source File: Santo_Domingo
//  
func (ameri americaTimeZones) Santo_Domingo() string {return "America/Santo_Domingo" }

// Sao_Paulo - IANA Time Zone 'America/Sao_Paulo'.
// IANA Source File: Sao_Paulo
//  
func (ameri americaTimeZones) Sao_Paulo() string {return "America/Sao_Paulo" }

// Scoresbysund - IANA Time Zone 'America/Scoresbysund'.
// IANA Source File: Scoresbysund
//  
func (ameri americaTimeZones) Scoresbysund() string {return "America/Scoresbysund" }

// Shiprock - IANA Time Zone 'America/Shiprock'.
// IANA Source File: Shiprock
//  
func (ameri americaTimeZones) Shiprock() string {return "America/Shiprock" }

// Sitka - IANA Time Zone 'America/Sitka'.
// IANA Source File: Sitka
//  
func (ameri americaTimeZones) Sitka() string {return "America/Sitka" }

// St_Barthelemy - IANA Time Zone 'America/St_Barthelemy'.
// IANA Source File: St_Barthelemy
//  
func (ameri americaTimeZones) St_Barthelemy() string {return "America/St_Barthelemy" }

// St_Johns - IANA Time Zone 'America/St_Johns'.
// IANA Source File: St_Johns
//  
func (ameri americaTimeZones) St_Johns() string {return "America/St_Johns" }

// St_Kitts - IANA Time Zone 'America/St_Kitts'.
// IANA Source File: St_Kitts
//  
func (ameri americaTimeZones) St_Kitts() string {return "America/St_Kitts" }

// St_Lucia - IANA Time Zone 'America/St_Lucia'.
// IANA Source File: St_Lucia
//  
func (ameri americaTimeZones) St_Lucia() string {return "America/St_Lucia" }

// St_Thomas - IANA Time Zone 'America/St_Thomas'.
// IANA Source File: St_Thomas
//  
func (ameri americaTimeZones) St_Thomas() string {return "America/St_Thomas" }

// St_Vincent - IANA Time Zone 'America/St_Vincent'.
// IANA Source File: St_Vincent
//  
func (ameri americaTimeZones) St_Vincent() string {return "America/St_Vincent" }

// Swift_Current - IANA Time Zone 'America/Swift_Current'.
// IANA Source File: Swift_Current
//  
func (ameri americaTimeZones) Swift_Current() string {return "America/Swift_Current" }

// Tegucigalpa - IANA Time Zone 'America/Tegucigalpa'.
// IANA Source File: Tegucigalpa
//  
func (ameri americaTimeZones) Tegucigalpa() string {return "America/Tegucigalpa" }

// Thule - IANA Time Zone 'America/Thule'.
// IANA Source File: Thule
//  
func (ameri americaTimeZones) Thule() string {return "America/Thule" }

// Thunder_Bay - IANA Time Zone 'America/Thunder_Bay'.
// IANA Source File: Thunder_Bay
//  
func (ameri americaTimeZones) Thunder_Bay() string {return "America/Thunder_Bay" }

// Tijuana - IANA Time Zone 'America/Tijuana'.
// IANA Source File: Tijuana
//  
func (ameri americaTimeZones) Tijuana() string {return "America/Tijuana" }

// Toronto - IANA Time Zone 'America/Toronto'.
// IANA Source File: Toronto
//  
func (ameri americaTimeZones) Toronto() string {return "America/Toronto" }

// Tortola - IANA Time Zone 'America/Tortola'.
// IANA Source File: Tortola
//  
func (ameri americaTimeZones) Tortola() string {return "America/Tortola" }

// Vancouver - IANA Time Zone 'America/Vancouver'.
// IANA Source File: Vancouver
//  
func (ameri americaTimeZones) Vancouver() string {return "America/Vancouver" }

// Virgin - IANA Time Zone 'America/Virgin'.
// IANA Source File: Virgin
//  
func (ameri americaTimeZones) Virgin() string {return "America/Virgin" }

// Whitehorse - IANA Time Zone 'America/Whitehorse'.
// IANA Source File: Whitehorse
//  
func (ameri americaTimeZones) Whitehorse() string {return "America/Whitehorse" }

// Winnipeg - IANA Time Zone 'America/Winnipeg'.
// IANA Source File: Winnipeg
//  
func (ameri americaTimeZones) Winnipeg() string {return "America/Winnipeg" }

// Yakutat - IANA Time Zone 'America/Yakutat'.
// IANA Source File: Yakutat
//  
func (ameri americaTimeZones) Yakutat() string {return "America/Yakutat" }

// Yellowknife - IANA Time Zone 'America/Yellowknife'.
// IANA Source File: Yellowknife
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
// IANA Source File: Casey
//  
func (antar antarcticaTimeZones) Casey() string {return "Antarctica/Casey" }

// Davis - IANA Time Zone 'Antarctica/Davis'.
// IANA Source File: Davis
//  
func (antar antarcticaTimeZones) Davis() string {return "Antarctica/Davis" }

// DumontDUrville - IANA Time Zone 'Antarctica/DumontDUrville'.
// IANA Source File: DumontDUrville
//  
func (antar antarcticaTimeZones) DumontDUrville() string {return "Antarctica/DumontDUrville" }

// Macquarie - IANA Time Zone 'Antarctica/Macquarie'.
// IANA Source File: Macquarie
//  
func (antar antarcticaTimeZones) Macquarie() string {return "Antarctica/Macquarie" }

// Mawson - IANA Time Zone 'Antarctica/Mawson'.
// IANA Source File: Mawson
//  
func (antar antarcticaTimeZones) Mawson() string {return "Antarctica/Mawson" }

// McMurdo - IANA Time Zone 'Antarctica/McMurdo'.
// IANA Source File: McMurdo
//  
func (antar antarcticaTimeZones) McMurdo() string {return "Antarctica/McMurdo" }

// Palmer - IANA Time Zone 'Antarctica/Palmer'.
// IANA Source File: Palmer
//  
func (antar antarcticaTimeZones) Palmer() string {return "Antarctica/Palmer" }

// Rothera - IANA Time Zone 'Antarctica/Rothera'.
// IANA Source File: Rothera
//  
func (antar antarcticaTimeZones) Rothera() string {return "Antarctica/Rothera" }

// South_Pole - IANA Time Zone 'Antarctica/South_Pole'.
// IANA Source File: South_Pole
//  
func (antar antarcticaTimeZones) South_Pole() string {return "Antarctica/South_Pole" }

// Syowa - IANA Time Zone 'Antarctica/Syowa'.
// IANA Source File: Syowa
//  
func (antar antarcticaTimeZones) Syowa() string {return "Antarctica/Syowa" }

// Troll - IANA Time Zone 'Antarctica/Troll'.
// IANA Source File: Troll
//  
func (antar antarcticaTimeZones) Troll() string {return "Antarctica/Troll" }

// Vostok - IANA Time Zone 'Antarctica/Vostok'.
// IANA Source File: Vostok
//  
func (antar antarcticaTimeZones) Vostok() string {return "Antarctica/Vostok" }


// arcticTimeZones - IANA Time Zones for 'Arctic'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type arcticTimeZones string

// Longyearbyen - IANA Time Zone 'Arctic/Longyearbyen'.
// IANA Source File: Longyearbyen
//  
func (arcti arcticTimeZones) Longyearbyen() string {return "Arctic/Longyearbyen" }


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
// IANA Source File: Aden
//  
func (asiaT asiaTimeZones) Aden() string {return "Asia/Aden" }

// Almaty - IANA Time Zone 'Asia/Almaty'.
// IANA Source File: Almaty
//  
func (asiaT asiaTimeZones) Almaty() string {return "Asia/Almaty" }

// Amman - IANA Time Zone 'Asia/Amman'.
// IANA Source File: Amman
//  
func (asiaT asiaTimeZones) Amman() string {return "Asia/Amman" }

// Anadyr - IANA Time Zone 'Asia/Anadyr'.
// IANA Source File: Anadyr
//  
func (asiaT asiaTimeZones) Anadyr() string {return "Asia/Anadyr" }

// Aqtau - IANA Time Zone 'Asia/Aqtau'.
// IANA Source File: Aqtau
//  
func (asiaT asiaTimeZones) Aqtau() string {return "Asia/Aqtau" }

// Aqtobe - IANA Time Zone 'Asia/Aqtobe'.
// IANA Source File: Aqtobe
//  
func (asiaT asiaTimeZones) Aqtobe() string {return "Asia/Aqtobe" }

// Ashgabat - IANA Time Zone 'Asia/Ashgabat'.
// IANA Source File: Ashgabat
//  
func (asiaT asiaTimeZones) Ashgabat() string {return "Asia/Ashgabat" }

// Ashkhabad - IANA Time Zone 'Asia/Ashkhabad'.
// IANA Source File: Ashkhabad
//  
func (asiaT asiaTimeZones) Ashkhabad() string {return "Asia/Ashkhabad" }

// Atyrau - IANA Time Zone 'Asia/Atyrau'.
// IANA Source File: Atyrau
//  
func (asiaT asiaTimeZones) Atyrau() string {return "Asia/Atyrau" }

// Baghdad - IANA Time Zone 'Asia/Baghdad'.
// IANA Source File: Baghdad
//  
func (asiaT asiaTimeZones) Baghdad() string {return "Asia/Baghdad" }

// Bahrain - IANA Time Zone 'Asia/Bahrain'.
// IANA Source File: Bahrain
//  
func (asiaT asiaTimeZones) Bahrain() string {return "Asia/Bahrain" }

// Baku - IANA Time Zone 'Asia/Baku'.
// IANA Source File: Baku
//  
func (asiaT asiaTimeZones) Baku() string {return "Asia/Baku" }

// Bangkok - IANA Time Zone 'Asia/Bangkok'.
// IANA Source File: Bangkok
//  
func (asiaT asiaTimeZones) Bangkok() string {return "Asia/Bangkok" }

// Barnaul - IANA Time Zone 'Asia/Barnaul'.
// IANA Source File: Barnaul
//  
func (asiaT asiaTimeZones) Barnaul() string {return "Asia/Barnaul" }

// Beirut - IANA Time Zone 'Asia/Beirut'.
// IANA Source File: Beirut
//  
func (asiaT asiaTimeZones) Beirut() string {return "Asia/Beirut" }

// Bishkek - IANA Time Zone 'Asia/Bishkek'.
// IANA Source File: Bishkek
//  
func (asiaT asiaTimeZones) Bishkek() string {return "Asia/Bishkek" }

// Brunei - IANA Time Zone 'Asia/Brunei'.
// IANA Source File: Brunei
//  
func (asiaT asiaTimeZones) Brunei() string {return "Asia/Brunei" }

// Calcutta - IANA Time Zone 'Asia/Calcutta'.
// IANA Source File: Calcutta
//  
func (asiaT asiaTimeZones) Calcutta() string {return "Asia/Calcutta" }

// Chita - IANA Time Zone 'Asia/Chita'.
// IANA Source File: Chita
//  
func (asiaT asiaTimeZones) Chita() string {return "Asia/Chita" }

// Choibalsan - IANA Time Zone 'Asia/Choibalsan'.
// IANA Source File: Choibalsan
//  
func (asiaT asiaTimeZones) Choibalsan() string {return "Asia/Choibalsan" }

// Chongqing - IANA Time Zone 'Asia/Chongqing'.
// IANA Source File: Chongqing
//  
func (asiaT asiaTimeZones) Chongqing() string {return "Asia/Chongqing" }

// Chungking - IANA Time Zone 'Asia/Chungking'.
// IANA Source File: Chungking
//  
func (asiaT asiaTimeZones) Chungking() string {return "Asia/Chungking" }

// Colombo - IANA Time Zone 'Asia/Colombo'.
// IANA Source File: Colombo
//  
func (asiaT asiaTimeZones) Colombo() string {return "Asia/Colombo" }

// Dacca - IANA Time Zone 'Asia/Dacca'.
// IANA Source File: Dacca
//  
func (asiaT asiaTimeZones) Dacca() string {return "Asia/Dacca" }

// Damascus - IANA Time Zone 'Asia/Damascus'.
// IANA Source File: Damascus
//  
func (asiaT asiaTimeZones) Damascus() string {return "Asia/Damascus" }

// Dhaka - IANA Time Zone 'Asia/Dhaka'.
// IANA Source File: Dhaka
//  
func (asiaT asiaTimeZones) Dhaka() string {return "Asia/Dhaka" }

// Dili - IANA Time Zone 'Asia/Dili'.
// IANA Source File: Dili
//  
func (asiaT asiaTimeZones) Dili() string {return "Asia/Dili" }

// Dubai - IANA Time Zone 'Asia/Dubai'.
// IANA Source File: Dubai
//  
func (asiaT asiaTimeZones) Dubai() string {return "Asia/Dubai" }

// Dushanbe - IANA Time Zone 'Asia/Dushanbe'.
// IANA Source File: Dushanbe
//  
func (asiaT asiaTimeZones) Dushanbe() string {return "Asia/Dushanbe" }

// Famagusta - IANA Time Zone 'Asia/Famagusta'.
// IANA Source File: Famagusta
//  
func (asiaT asiaTimeZones) Famagusta() string {return "Asia/Famagusta" }

// Gaza - IANA Time Zone 'Asia/Gaza'.
// IANA Source File: Gaza
//  
func (asiaT asiaTimeZones) Gaza() string {return "Asia/Gaza" }

// Harbin - IANA Time Zone 'Asia/Harbin'.
// IANA Source File: Harbin
//  
func (asiaT asiaTimeZones) Harbin() string {return "Asia/Harbin" }

// Hebron - IANA Time Zone 'Asia/Hebron'.
// IANA Source File: Hebron
//  
func (asiaT asiaTimeZones) Hebron() string {return "Asia/Hebron" }

// Ho_Chi_Minh - IANA Time Zone 'Asia/Ho_Chi_Minh'.
// IANA Source File: Ho_Chi_Minh
//  
func (asiaT asiaTimeZones) Ho_Chi_Minh() string {return "Asia/Ho_Chi_Minh" }

// Hong_Kong - IANA Time Zone 'Asia/Hong_Kong'.
// IANA Source File: Hong_Kong
//  
func (asiaT asiaTimeZones) Hong_Kong() string {return "Asia/Hong_Kong" }

// Hovd - IANA Time Zone 'Asia/Hovd'.
// IANA Source File: Hovd
//  
func (asiaT asiaTimeZones) Hovd() string {return "Asia/Hovd" }

// Irkutsk - IANA Time Zone 'Asia/Irkutsk'.
// IANA Source File: Irkutsk
//  
func (asiaT asiaTimeZones) Irkutsk() string {return "Asia/Irkutsk" }

// Istanbul - IANA Time Zone 'Asia/Istanbul'.
// IANA Source File: Istanbul
//  
func (asiaT asiaTimeZones) Istanbul() string {return "Asia/Istanbul" }

// Jakarta - IANA Time Zone 'Asia/Jakarta'.
// IANA Source File: Jakarta
//  
func (asiaT asiaTimeZones) Jakarta() string {return "Asia/Jakarta" }

// Jayapura - IANA Time Zone 'Asia/Jayapura'.
// IANA Source File: Jayapura
//  
func (asiaT asiaTimeZones) Jayapura() string {return "Asia/Jayapura" }

// Jerusalem - IANA Time Zone 'Asia/Jerusalem'.
// IANA Source File: Jerusalem
//  
func (asiaT asiaTimeZones) Jerusalem() string {return "Asia/Jerusalem" }

// Kabul - IANA Time Zone 'Asia/Kabul'.
// IANA Source File: Kabul
//  
func (asiaT asiaTimeZones) Kabul() string {return "Asia/Kabul" }

// Kamchatka - IANA Time Zone 'Asia/Kamchatka'.
// IANA Source File: Kamchatka
//  
func (asiaT asiaTimeZones) Kamchatka() string {return "Asia/Kamchatka" }

// Karachi - IANA Time Zone 'Asia/Karachi'.
// IANA Source File: Karachi
//  
func (asiaT asiaTimeZones) Karachi() string {return "Asia/Karachi" }

// Kashgar - IANA Time Zone 'Asia/Kashgar'.
// IANA Source File: Kashgar
//  
func (asiaT asiaTimeZones) Kashgar() string {return "Asia/Kashgar" }

// Kathmandu - IANA Time Zone 'Asia/Kathmandu'.
// IANA Source File: Kathmandu
//  
func (asiaT asiaTimeZones) Kathmandu() string {return "Asia/Kathmandu" }

// Katmandu - IANA Time Zone 'Asia/Katmandu'.
// IANA Source File: Katmandu
//  
func (asiaT asiaTimeZones) Katmandu() string {return "Asia/Katmandu" }

// Khandyga - IANA Time Zone 'Asia/Khandyga'.
// IANA Source File: Khandyga
//  
func (asiaT asiaTimeZones) Khandyga() string {return "Asia/Khandyga" }

// Kolkata - IANA Time Zone 'Asia/Kolkata'.
// IANA Source File: Kolkata
//  
func (asiaT asiaTimeZones) Kolkata() string {return "Asia/Kolkata" }

// Krasnoyarsk - IANA Time Zone 'Asia/Krasnoyarsk'.
// IANA Source File: Krasnoyarsk
//  
func (asiaT asiaTimeZones) Krasnoyarsk() string {return "Asia/Krasnoyarsk" }

// Kuala_Lumpur - IANA Time Zone 'Asia/Kuala_Lumpur'.
// IANA Source File: Kuala_Lumpur
//  
func (asiaT asiaTimeZones) Kuala_Lumpur() string {return "Asia/Kuala_Lumpur" }

// Kuching - IANA Time Zone 'Asia/Kuching'.
// IANA Source File: Kuching
//  
func (asiaT asiaTimeZones) Kuching() string {return "Asia/Kuching" }

// Kuwait - IANA Time Zone 'Asia/Kuwait'.
// IANA Source File: Kuwait
//  
func (asiaT asiaTimeZones) Kuwait() string {return "Asia/Kuwait" }

// Macao - IANA Time Zone 'Asia/Macao'.
// IANA Source File: Macao
//  
func (asiaT asiaTimeZones) Macao() string {return "Asia/Macao" }

// Macau - IANA Time Zone 'Asia/Macau'.
// IANA Source File: Macau
//  
func (asiaT asiaTimeZones) Macau() string {return "Asia/Macau" }

// Magadan - IANA Time Zone 'Asia/Magadan'.
// IANA Source File: Magadan
//  
func (asiaT asiaTimeZones) Magadan() string {return "Asia/Magadan" }

// Makassar - IANA Time Zone 'Asia/Makassar'.
// IANA Source File: Makassar
//  
func (asiaT asiaTimeZones) Makassar() string {return "Asia/Makassar" }

// Manila - IANA Time Zone 'Asia/Manila'.
// IANA Source File: Manila
//  
func (asiaT asiaTimeZones) Manila() string {return "Asia/Manila" }

// Muscat - IANA Time Zone 'Asia/Muscat'.
// IANA Source File: Muscat
//  
func (asiaT asiaTimeZones) Muscat() string {return "Asia/Muscat" }

// Nicosia - IANA Time Zone 'Asia/Nicosia'.
// IANA Source File: Nicosia
//  
func (asiaT asiaTimeZones) Nicosia() string {return "Asia/Nicosia" }

// Novokuznetsk - IANA Time Zone 'Asia/Novokuznetsk'.
// IANA Source File: Novokuznetsk
//  
func (asiaT asiaTimeZones) Novokuznetsk() string {return "Asia/Novokuznetsk" }

// Novosibirsk - IANA Time Zone 'Asia/Novosibirsk'.
// IANA Source File: Novosibirsk
//  
func (asiaT asiaTimeZones) Novosibirsk() string {return "Asia/Novosibirsk" }

// Omsk - IANA Time Zone 'Asia/Omsk'.
// IANA Source File: Omsk
//  
func (asiaT asiaTimeZones) Omsk() string {return "Asia/Omsk" }

// Oral - IANA Time Zone 'Asia/Oral'.
// IANA Source File: Oral
//  
func (asiaT asiaTimeZones) Oral() string {return "Asia/Oral" }

// Phnom_Penh - IANA Time Zone 'Asia/Phnom_Penh'.
// IANA Source File: Phnom_Penh
//  
func (asiaT asiaTimeZones) Phnom_Penh() string {return "Asia/Phnom_Penh" }

// Pontianak - IANA Time Zone 'Asia/Pontianak'.
// IANA Source File: Pontianak
//  
func (asiaT asiaTimeZones) Pontianak() string {return "Asia/Pontianak" }

// Pyongyang - IANA Time Zone 'Asia/Pyongyang'.
// IANA Source File: Pyongyang
//  
func (asiaT asiaTimeZones) Pyongyang() string {return "Asia/Pyongyang" }

// Qatar - IANA Time Zone 'Asia/Qatar'.
// IANA Source File: Qatar
//  
func (asiaT asiaTimeZones) Qatar() string {return "Asia/Qatar" }

// Qostanay - IANA Time Zone 'Asia/Qostanay'.
// IANA Source File: Qostanay
//  
func (asiaT asiaTimeZones) Qostanay() string {return "Asia/Qostanay" }

// Qyzylorda - IANA Time Zone 'Asia/Qyzylorda'.
// IANA Source File: Qyzylorda
//  
func (asiaT asiaTimeZones) Qyzylorda() string {return "Asia/Qyzylorda" }

// Rangoon - IANA Time Zone 'Asia/Rangoon'.
// IANA Source File: Rangoon
//  
func (asiaT asiaTimeZones) Rangoon() string {return "Asia/Rangoon" }

// Riyadh - IANA Time Zone 'Asia/Riyadh'.
// IANA Source File: Riyadh
//  
func (asiaT asiaTimeZones) Riyadh() string {return "Asia/Riyadh" }

// Saigon - IANA Time Zone 'Asia/Saigon'.
// IANA Source File: Saigon
//  
func (asiaT asiaTimeZones) Saigon() string {return "Asia/Saigon" }

// Sakhalin - IANA Time Zone 'Asia/Sakhalin'.
// IANA Source File: Sakhalin
//  
func (asiaT asiaTimeZones) Sakhalin() string {return "Asia/Sakhalin" }

// Samarkand - IANA Time Zone 'Asia/Samarkand'.
// IANA Source File: Samarkand
//  
func (asiaT asiaTimeZones) Samarkand() string {return "Asia/Samarkand" }

// Seoul - IANA Time Zone 'Asia/Seoul'.
// IANA Source File: Seoul
//  
func (asiaT asiaTimeZones) Seoul() string {return "Asia/Seoul" }

// Shanghai - IANA Time Zone 'Asia/Shanghai'.
// IANA Source File: Shanghai
//  
func (asiaT asiaTimeZones) Shanghai() string {return "Asia/Shanghai" }

// Singapore - IANA Time Zone 'Asia/Singapore'.
// IANA Source File: Singapore
//  
func (asiaT asiaTimeZones) Singapore() string {return "Asia/Singapore" }

// Srednekolymsk - IANA Time Zone 'Asia/Srednekolymsk'.
// IANA Source File: Srednekolymsk
//  
func (asiaT asiaTimeZones) Srednekolymsk() string {return "Asia/Srednekolymsk" }

// Taipei - IANA Time Zone 'Asia/Taipei'.
// IANA Source File: Taipei
//  
func (asiaT asiaTimeZones) Taipei() string {return "Asia/Taipei" }

// Tashkent - IANA Time Zone 'Asia/Tashkent'.
// IANA Source File: Tashkent
//  
func (asiaT asiaTimeZones) Tashkent() string {return "Asia/Tashkent" }

// Tbilisi - IANA Time Zone 'Asia/Tbilisi'.
// IANA Source File: Tbilisi
//  
func (asiaT asiaTimeZones) Tbilisi() string {return "Asia/Tbilisi" }

// Tehran - IANA Time Zone 'Asia/Tehran'.
// IANA Source File: Tehran
//  
func (asiaT asiaTimeZones) Tehran() string {return "Asia/Tehran" }

// Tel_Aviv - IANA Time Zone 'Asia/Tel_Aviv'.
// IANA Source File: Tel_Aviv
//  
func (asiaT asiaTimeZones) Tel_Aviv() string {return "Asia/Tel_Aviv" }

// Thimbu - IANA Time Zone 'Asia/Thimbu'.
// IANA Source File: Thimbu
//  
func (asiaT asiaTimeZones) Thimbu() string {return "Asia/Thimbu" }

// Thimphu - IANA Time Zone 'Asia/Thimphu'.
// IANA Source File: Thimphu
//  
func (asiaT asiaTimeZones) Thimphu() string {return "Asia/Thimphu" }

// Tokyo - IANA Time Zone 'Asia/Tokyo'.
// IANA Source File: Tokyo
//  
func (asiaT asiaTimeZones) Tokyo() string {return "Asia/Tokyo" }

// Tomsk - IANA Time Zone 'Asia/Tomsk'.
// IANA Source File: Tomsk
//  
func (asiaT asiaTimeZones) Tomsk() string {return "Asia/Tomsk" }

// Ujung_Pandang - IANA Time Zone 'Asia/Ujung_Pandang'.
// IANA Source File: Ujung_Pandang
//  
func (asiaT asiaTimeZones) Ujung_Pandang() string {return "Asia/Ujung_Pandang" }

// Ulaanbaatar - IANA Time Zone 'Asia/Ulaanbaatar'.
// IANA Source File: Ulaanbaatar
//  
func (asiaT asiaTimeZones) Ulaanbaatar() string {return "Asia/Ulaanbaatar" }

// Ulan_Bator - IANA Time Zone 'Asia/Ulan_Bator'.
// IANA Source File: Ulan_Bator
//  
func (asiaT asiaTimeZones) Ulan_Bator() string {return "Asia/Ulan_Bator" }

// Urumqi - IANA Time Zone 'Asia/Urumqi'.
// IANA Source File: Urumqi
//  
func (asiaT asiaTimeZones) Urumqi() string {return "Asia/Urumqi" }

// Ust-Nera - IANA Time Zone 'Asia/Ust-Nera'.
// IANA Source File: Ust-Nera
//  
func (asiaT asiaTimeZones) UstMinusNera() string {return "Asia/Ust-Nera" }

// Vientiane - IANA Time Zone 'Asia/Vientiane'.
// IANA Source File: Vientiane
//  
func (asiaT asiaTimeZones) Vientiane() string {return "Asia/Vientiane" }

// Vladivostok - IANA Time Zone 'Asia/Vladivostok'.
// IANA Source File: Vladivostok
//  
func (asiaT asiaTimeZones) Vladivostok() string {return "Asia/Vladivostok" }

// Yakutsk - IANA Time Zone 'Asia/Yakutsk'.
// IANA Source File: Yakutsk
//  
func (asiaT asiaTimeZones) Yakutsk() string {return "Asia/Yakutsk" }

// Yangon - IANA Time Zone 'Asia/Yangon'.
// IANA Source File: Yangon
//  
func (asiaT asiaTimeZones) Yangon() string {return "Asia/Yangon" }

// Yekaterinburg - IANA Time Zone 'Asia/Yekaterinburg'.
// IANA Source File: Yekaterinburg
//  
func (asiaT asiaTimeZones) Yekaterinburg() string {return "Asia/Yekaterinburg" }

// Yerevan - IANA Time Zone 'Asia/Yerevan'.
// IANA Source File: Yerevan
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
// IANA Source File: Azores
//  
func (atlan atlanticTimeZones) Azores() string {return "Atlantic/Azores" }

// Bermuda - IANA Time Zone 'Atlantic/Bermuda'.
// IANA Source File: Bermuda
//  
func (atlan atlanticTimeZones) Bermuda() string {return "Atlantic/Bermuda" }

// Canary - IANA Time Zone 'Atlantic/Canary'.
// IANA Source File: Canary
//  
func (atlan atlanticTimeZones) Canary() string {return "Atlantic/Canary" }

// Cape_Verde - IANA Time Zone 'Atlantic/Cape_Verde'.
// IANA Source File: Cape_Verde
//  
func (atlan atlanticTimeZones) Cape_Verde() string {return "Atlantic/Cape_Verde" }

// Faeroe - IANA Time Zone 'Atlantic/Faeroe'.
// IANA Source File: Faeroe
//  
func (atlan atlanticTimeZones) Faeroe() string {return "Atlantic/Faeroe" }

// Faroe - IANA Time Zone 'Atlantic/Faroe'.
// IANA Source File: Faroe
//  
func (atlan atlanticTimeZones) Faroe() string {return "Atlantic/Faroe" }

// Jan_Mayen - IANA Time Zone 'Atlantic/Jan_Mayen'.
// IANA Source File: Jan_Mayen
//  
func (atlan atlanticTimeZones) Jan_Mayen() string {return "Atlantic/Jan_Mayen" }

// Madeira - IANA Time Zone 'Atlantic/Madeira'.
// IANA Source File: Madeira
//  
func (atlan atlanticTimeZones) Madeira() string {return "Atlantic/Madeira" }

// Reykjavik - IANA Time Zone 'Atlantic/Reykjavik'.
// IANA Source File: Reykjavik
//  
func (atlan atlanticTimeZones) Reykjavik() string {return "Atlantic/Reykjavik" }

// South_Georgia - IANA Time Zone 'Atlantic/South_Georgia'.
// IANA Source File: South_Georgia
//  
func (atlan atlanticTimeZones) South_Georgia() string {return "Atlantic/South_Georgia" }

// St_Helena - IANA Time Zone 'Atlantic/St_Helena'.
// IANA Source File: St_Helena
//  
func (atlan atlanticTimeZones) St_Helena() string {return "Atlantic/St_Helena" }

// Stanley - IANA Time Zone 'Atlantic/Stanley'.
// IANA Source File: Stanley
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

// ACT - IANA Time Zone 'Australia/ACT'.
// IANA Source File: ACT
//  
func (austr australiaTimeZones) ACT() string {return "Australia/ACT" }

// Adelaide - IANA Time Zone 'Australia/Adelaide'.
// IANA Source File: Adelaide
//  
func (austr australiaTimeZones) Adelaide() string {return "Australia/Adelaide" }

// Brisbane - IANA Time Zone 'Australia/Brisbane'.
// IANA Source File: Brisbane
//  
func (austr australiaTimeZones) Brisbane() string {return "Australia/Brisbane" }

// Broken_Hill - IANA Time Zone 'Australia/Broken_Hill'.
// IANA Source File: Broken_Hill
//  
func (austr australiaTimeZones) Broken_Hill() string {return "Australia/Broken_Hill" }

// Canberra - IANA Time Zone 'Australia/Canberra'.
// IANA Source File: Canberra
//  
func (austr australiaTimeZones) Canberra() string {return "Australia/Canberra" }

// Currie - IANA Time Zone 'Australia/Currie'.
// IANA Source File: Currie
//  
func (austr australiaTimeZones) Currie() string {return "Australia/Currie" }

// Darwin - IANA Time Zone 'Australia/Darwin'.
// IANA Source File: Darwin
//  
func (austr australiaTimeZones) Darwin() string {return "Australia/Darwin" }

// Eucla - IANA Time Zone 'Australia/Eucla'.
// IANA Source File: Eucla
//  
func (austr australiaTimeZones) Eucla() string {return "Australia/Eucla" }

// Hobart - IANA Time Zone 'Australia/Hobart'.
// IANA Source File: Hobart
//  
func (austr australiaTimeZones) Hobart() string {return "Australia/Hobart" }

// LHI - IANA Time Zone 'Australia/LHI'.
// IANA Source File: LHI
//  
func (austr australiaTimeZones) LHI() string {return "Australia/LHI" }

// Lindeman - IANA Time Zone 'Australia/Lindeman'.
// IANA Source File: Lindeman
//  
func (austr australiaTimeZones) Lindeman() string {return "Australia/Lindeman" }

// Lord_Howe - IANA Time Zone 'Australia/Lord_Howe'.
// IANA Source File: Lord_Howe
//  
func (austr australiaTimeZones) Lord_Howe() string {return "Australia/Lord_Howe" }

// Melbourne - IANA Time Zone 'Australia/Melbourne'.
// IANA Source File: Melbourne
//  
func (austr australiaTimeZones) Melbourne() string {return "Australia/Melbourne" }

// North - IANA Time Zone 'Australia/North'.
// IANA Source File: North
//  
func (austr australiaTimeZones) North() string {return "Australia/North" }

// NSW - IANA Time Zone 'Australia/NSW'.
// IANA Source File: NSW
//  
func (austr australiaTimeZones) NSW() string {return "Australia/NSW" }

// Perth - IANA Time Zone 'Australia/Perth'.
// IANA Source File: Perth
//  
func (austr australiaTimeZones) Perth() string {return "Australia/Perth" }

// Queensland - IANA Time Zone 'Australia/Queensland'.
// IANA Source File: Queensland
//  
func (austr australiaTimeZones) Queensland() string {return "Australia/Queensland" }

// South - IANA Time Zone 'Australia/South'.
// IANA Source File: South
//  
func (austr australiaTimeZones) South() string {return "Australia/South" }

// Sydney - IANA Time Zone 'Australia/Sydney'.
// IANA Source File: Sydney
//  
func (austr australiaTimeZones) Sydney() string {return "Australia/Sydney" }

// Tasmania - IANA Time Zone 'Australia/Tasmania'.
// IANA Source File: Tasmania
//  
func (austr australiaTimeZones) Tasmania() string {return "Australia/Tasmania" }

// Victoria - IANA Time Zone 'Australia/Victoria'.
// IANA Source File: Victoria
//  
func (austr australiaTimeZones) Victoria() string {return "Australia/Victoria" }

// West - IANA Time Zone 'Australia/West'.
// IANA Source File: West
//  
func (austr australiaTimeZones) West() string {return "Australia/West" }

// Yancowinna - IANA Time Zone 'Australia/Yancowinna'.
// IANA Source File: Yancowinna
//  
func (austr australiaTimeZones) Yancowinna() string {return "Australia/Yancowinna" }


// brazilTimeZones - IANA Time Zones for 'Brazil'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type brazilTimeZones string

// Acre - IANA Time Zone 'Brazil/Acre'.
// IANA Source File: Acre
//  
func (brazi brazilTimeZones) Acre() string {return "Brazil/Acre" }

// DeNoronha - IANA Time Zone 'Brazil/DeNoronha'.
// IANA Source File: DeNoronha
//  
func (brazi brazilTimeZones) DeNoronha() string {return "Brazil/DeNoronha" }

// East - IANA Time Zone 'Brazil/East'.
// IANA Source File: East
//  
func (brazi brazilTimeZones) East() string {return "Brazil/East" }

// West - IANA Time Zone 'Brazil/West'.
// IANA Source File: West
//  
func (brazi brazilTimeZones) West() string {return "Brazil/West" }


// canadaTimeZones - IANA Time Zones for 'Canada'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type canadaTimeZones string

// Atlantic - IANA Time Zone 'Canada/Atlantic'.
// IANA Source File: Atlantic
//  
func (canad canadaTimeZones) Atlantic() string {return "Canada/Atlantic" }

// Central - IANA Time Zone 'Canada/Central'.
// IANA Source File: Central
//  
func (canad canadaTimeZones) Central() string {return "Canada/Central" }

// Eastern - IANA Time Zone 'Canada/Eastern'.
// IANA Source File: Eastern
//  
func (canad canadaTimeZones) Eastern() string {return "Canada/Eastern" }

// Mountain - IANA Time Zone 'Canada/Mountain'.
// IANA Source File: Mountain
//  
func (canad canadaTimeZones) Mountain() string {return "Canada/Mountain" }

// Newfoundland - IANA Time Zone 'Canada/Newfoundland'.
// IANA Source File: Newfoundland
//  
func (canad canadaTimeZones) Newfoundland() string {return "Canada/Newfoundland" }

// Pacific - IANA Time Zone 'Canada/Pacific'.
// IANA Source File: Pacific
//  
func (canad canadaTimeZones) Pacific() string {return "Canada/Pacific" }

// Saskatchewan - IANA Time Zone 'Canada/Saskatchewan'.
// IANA Source File: Saskatchewan
//  
func (canad canadaTimeZones) Saskatchewan() string {return "Canada/Saskatchewan" }

// Yukon - IANA Time Zone 'Canada/Yukon'.
// IANA Source File: Yukon
//  
func (canad canadaTimeZones) Yukon() string {return "Canada/Yukon" }


// chileTimeZones - IANA Time Zones for 'Chile'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type chileTimeZones string

// Continental - IANA Time Zone 'Chile/Continental'.
// IANA Source File: Continental
//  
func (chile chileTimeZones) Continental() string {return "Chile/Continental" }

// EasterIsland - IANA Time Zone 'Chile/EasterIsland'.
// IANA Source File: EasterIsland
//  
func (chile chileTimeZones) EasterIsland() string {return "Chile/EasterIsland" }


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
// IANA Source File: GMT
//  
func (etcTi etcTimeZones) GMT() string {return "Etc/GMT" }

// GMT+0 - IANA Time Zone 'Etc/GMT+0'.
// IANA Source File: GMT+0
//  
func (etcTi etcTimeZones) GMTPlus00() string {return "Etc/GMT+0" }

// GMT+1 - IANA Time Zone 'Etc/GMT+1'.
// IANA Source File: GMT+1
//  
func (etcTi etcTimeZones) GMTPlus01() string {return "Etc/GMT+1" }

// GMT+2 - IANA Time Zone 'Etc/GMT+2'.
// IANA Source File: GMT+2
//  
func (etcTi etcTimeZones) GMTPlus02() string {return "Etc/GMT+2" }

// GMT+3 - IANA Time Zone 'Etc/GMT+3'.
// IANA Source File: GMT+3
//  
func (etcTi etcTimeZones) GMTPlus03() string {return "Etc/GMT+3" }

// GMT+4 - IANA Time Zone 'Etc/GMT+4'.
// IANA Source File: GMT+4
//  
func (etcTi etcTimeZones) GMTPlus04() string {return "Etc/GMT+4" }

// GMT+5 - IANA Time Zone 'Etc/GMT+5'.
// IANA Source File: GMT+5
//  
func (etcTi etcTimeZones) GMTPlus05() string {return "Etc/GMT+5" }

// GMT+6 - IANA Time Zone 'Etc/GMT+6'.
// IANA Source File: GMT+6
//  
func (etcTi etcTimeZones) GMTPlus06() string {return "Etc/GMT+6" }

// GMT+7 - IANA Time Zone 'Etc/GMT+7'.
// IANA Source File: GMT+7
//  
func (etcTi etcTimeZones) GMTPlus07() string {return "Etc/GMT+7" }

// GMT+8 - IANA Time Zone 'Etc/GMT+8'.
// IANA Source File: GMT+8
//  
func (etcTi etcTimeZones) GMTPlus08() string {return "Etc/GMT+8" }

// GMT+9 - IANA Time Zone 'Etc/GMT+9'.
// IANA Source File: GMT+9
//  
func (etcTi etcTimeZones) GMTPlus09() string {return "Etc/GMT+9" }

// GMT+10 - IANA Time Zone 'Etc/GMT+10'.
// IANA Source File: GMT+10
//  
func (etcTi etcTimeZones) GMTPlus10() string {return "Etc/GMT+10" }

// GMT+11 - IANA Time Zone 'Etc/GMT+11'.
// IANA Source File: GMT+11
//  
func (etcTi etcTimeZones) GMTPlus11() string {return "Etc/GMT+11" }

// GMT+12 - IANA Time Zone 'Etc/GMT+12'.
// IANA Source File: GMT+12
//  
func (etcTi etcTimeZones) GMTPlus12() string {return "Etc/GMT+12" }

// GMT-0 - IANA Time Zone 'Etc/GMT-0'.
// IANA Source File: GMT-0
//  
func (etcTi etcTimeZones) GMTMinus00() string {return "Etc/GMT-0" }

// GMT-1 - IANA Time Zone 'Etc/GMT-1'.
// IANA Source File: GMT-1
//  
func (etcTi etcTimeZones) GMTMinus01() string {return "Etc/GMT-1" }

// GMT-2 - IANA Time Zone 'Etc/GMT-2'.
// IANA Source File: GMT-2
//  
func (etcTi etcTimeZones) GMTMinus02() string {return "Etc/GMT-2" }

// GMT-3 - IANA Time Zone 'Etc/GMT-3'.
// IANA Source File: GMT-3
//  
func (etcTi etcTimeZones) GMTMinus03() string {return "Etc/GMT-3" }

// GMT-4 - IANA Time Zone 'Etc/GMT-4'.
// IANA Source File: GMT-4
//  
func (etcTi etcTimeZones) GMTMinus04() string {return "Etc/GMT-4" }

// GMT-5 - IANA Time Zone 'Etc/GMT-5'.
// IANA Source File: GMT-5
//  
func (etcTi etcTimeZones) GMTMinus05() string {return "Etc/GMT-5" }

// GMT-6 - IANA Time Zone 'Etc/GMT-6'.
// IANA Source File: GMT-6
//  
func (etcTi etcTimeZones) GMTMinus06() string {return "Etc/GMT-6" }

// GMT-7 - IANA Time Zone 'Etc/GMT-7'.
// IANA Source File: GMT-7
//  
func (etcTi etcTimeZones) GMTMinus07() string {return "Etc/GMT-7" }

// GMT-8 - IANA Time Zone 'Etc/GMT-8'.
// IANA Source File: GMT-8
//  
func (etcTi etcTimeZones) GMTMinus08() string {return "Etc/GMT-8" }

// GMT-9 - IANA Time Zone 'Etc/GMT-9'.
// IANA Source File: GMT-9
//  
func (etcTi etcTimeZones) GMTMinus09() string {return "Etc/GMT-9" }

// GMT-10 - IANA Time Zone 'Etc/GMT-10'.
// IANA Source File: GMT-10
//  
func (etcTi etcTimeZones) GMTMinus10() string {return "Etc/GMT-10" }

// GMT-11 - IANA Time Zone 'Etc/GMT-11'.
// IANA Source File: GMT-11
//  
func (etcTi etcTimeZones) GMTMinus11() string {return "Etc/GMT-11" }

// GMT-12 - IANA Time Zone 'Etc/GMT-12'.
// IANA Source File: GMT-12
//  
func (etcTi etcTimeZones) GMTMinus12() string {return "Etc/GMT-12" }

// GMT-13 - IANA Time Zone 'Etc/GMT-13'.
// IANA Source File: GMT-13
//  
func (etcTi etcTimeZones) GMTMinus13() string {return "Etc/GMT-13" }

// GMT-14 - IANA Time Zone 'Etc/GMT-14'.
// IANA Source File: GMT-14
//  
func (etcTi etcTimeZones) GMTMinus14() string {return "Etc/GMT-14" }

// GMT0 - IANA Time Zone 'Etc/GMT0'.
// IANA Source File: GMT0
//  
func (etcTi etcTimeZones) GMT00() string {return "Etc/GMT0" }

// Greenwich - IANA Time Zone 'Etc/Greenwich'.
// IANA Source File: Greenwich
//  
func (etcTi etcTimeZones) Greenwich() string {return "Etc/Greenwich" }

// UCT - IANA Time Zone 'Etc/UCT'.
// IANA Source File: UCT
//  
func (etcTi etcTimeZones) UCT() string {return "Etc/UCT" }

// Universal - IANA Time Zone 'Etc/Universal'.
// IANA Source File: Universal
//  
func (etcTi etcTimeZones) Universal() string {return "Etc/Universal" }

// UTC - IANA Time Zone 'Etc/UTC'.
// IANA Source File: UTC
//  
func (etcTi etcTimeZones) UTC() string {return "Etc/UTC" }

// Zulu - IANA Time Zone 'Etc/Zulu'.
// IANA Source File: Zulu
//  
func (etcTi etcTimeZones) Zulu() string {return "Etc/Zulu" }


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
// IANA Source File: Amsterdam
//  
func (europ europeTimeZones) Amsterdam() string {return "Europe/Amsterdam" }

// Andorra - IANA Time Zone 'Europe/Andorra'.
// IANA Source File: Andorra
//  
func (europ europeTimeZones) Andorra() string {return "Europe/Andorra" }

// Astrakhan - IANA Time Zone 'Europe/Astrakhan'.
// IANA Source File: Astrakhan
//  
func (europ europeTimeZones) Astrakhan() string {return "Europe/Astrakhan" }

// Athens - IANA Time Zone 'Europe/Athens'.
// IANA Source File: Athens
//  
func (europ europeTimeZones) Athens() string {return "Europe/Athens" }

// Belfast - IANA Time Zone 'Europe/Belfast'.
// IANA Source File: Belfast
//  
func (europ europeTimeZones) Belfast() string {return "Europe/Belfast" }

// Belgrade - IANA Time Zone 'Europe/Belgrade'.
// IANA Source File: Belgrade
//  
func (europ europeTimeZones) Belgrade() string {return "Europe/Belgrade" }

// Berlin - IANA Time Zone 'Europe/Berlin'.
// IANA Source File: Berlin
//  
func (europ europeTimeZones) Berlin() string {return "Europe/Berlin" }

// Bratislava - IANA Time Zone 'Europe/Bratislava'.
// IANA Source File: Bratislava
//  
func (europ europeTimeZones) Bratislava() string {return "Europe/Bratislava" }

// Brussels - IANA Time Zone 'Europe/Brussels'.
// IANA Source File: Brussels
//  
func (europ europeTimeZones) Brussels() string {return "Europe/Brussels" }

// Bucharest - IANA Time Zone 'Europe/Bucharest'.
// IANA Source File: Bucharest
//  
func (europ europeTimeZones) Bucharest() string {return "Europe/Bucharest" }

// Budapest - IANA Time Zone 'Europe/Budapest'.
// IANA Source File: Budapest
//  
func (europ europeTimeZones) Budapest() string {return "Europe/Budapest" }

// Busingen - IANA Time Zone 'Europe/Busingen'.
// IANA Source File: Busingen
//  
func (europ europeTimeZones) Busingen() string {return "Europe/Busingen" }

// Chisinau - IANA Time Zone 'Europe/Chisinau'.
// IANA Source File: Chisinau
//  
func (europ europeTimeZones) Chisinau() string {return "Europe/Chisinau" }

// Copenhagen - IANA Time Zone 'Europe/Copenhagen'.
// IANA Source File: Copenhagen
//  
func (europ europeTimeZones) Copenhagen() string {return "Europe/Copenhagen" }

// Dublin - IANA Time Zone 'Europe/Dublin'.
// IANA Source File: Dublin
//  
func (europ europeTimeZones) Dublin() string {return "Europe/Dublin" }

// Gibraltar - IANA Time Zone 'Europe/Gibraltar'.
// IANA Source File: Gibraltar
//  
func (europ europeTimeZones) Gibraltar() string {return "Europe/Gibraltar" }

// Guernsey - IANA Time Zone 'Europe/Guernsey'.
// IANA Source File: Guernsey
//  
func (europ europeTimeZones) Guernsey() string {return "Europe/Guernsey" }

// Helsinki - IANA Time Zone 'Europe/Helsinki'.
// IANA Source File: Helsinki
//  
func (europ europeTimeZones) Helsinki() string {return "Europe/Helsinki" }

// Isle_of_Man - IANA Time Zone 'Europe/Isle_of_Man'.
// IANA Source File: Isle_of_Man
//  
func (europ europeTimeZones) Isle_of_Man() string {return "Europe/Isle_of_Man" }

// Istanbul - IANA Time Zone 'Europe/Istanbul'.
// IANA Source File: Istanbul
//  
func (europ europeTimeZones) Istanbul() string {return "Europe/Istanbul" }

// Jersey - IANA Time Zone 'Europe/Jersey'.
// IANA Source File: Jersey
//  
func (europ europeTimeZones) Jersey() string {return "Europe/Jersey" }

// Kaliningrad - IANA Time Zone 'Europe/Kaliningrad'.
// IANA Source File: Kaliningrad
//  
func (europ europeTimeZones) Kaliningrad() string {return "Europe/Kaliningrad" }

// Kiev - IANA Time Zone 'Europe/Kiev'.
// IANA Source File: Kiev
//  
func (europ europeTimeZones) Kiev() string {return "Europe/Kiev" }

// Kirov - IANA Time Zone 'Europe/Kirov'.
// IANA Source File: Kirov
//  
func (europ europeTimeZones) Kirov() string {return "Europe/Kirov" }

// Lisbon - IANA Time Zone 'Europe/Lisbon'.
// IANA Source File: Lisbon
//  
func (europ europeTimeZones) Lisbon() string {return "Europe/Lisbon" }

// Ljubljana - IANA Time Zone 'Europe/Ljubljana'.
// IANA Source File: Ljubljana
//  
func (europ europeTimeZones) Ljubljana() string {return "Europe/Ljubljana" }

// London - IANA Time Zone 'Europe/London'.
// IANA Source File: London
//  
func (europ europeTimeZones) London() string {return "Europe/London" }

// Luxembourg - IANA Time Zone 'Europe/Luxembourg'.
// IANA Source File: Luxembourg
//  
func (europ europeTimeZones) Luxembourg() string {return "Europe/Luxembourg" }

// Madrid - IANA Time Zone 'Europe/Madrid'.
// IANA Source File: Madrid
//  
func (europ europeTimeZones) Madrid() string {return "Europe/Madrid" }

// Malta - IANA Time Zone 'Europe/Malta'.
// IANA Source File: Malta
//  
func (europ europeTimeZones) Malta() string {return "Europe/Malta" }

// Mariehamn - IANA Time Zone 'Europe/Mariehamn'.
// IANA Source File: Mariehamn
//  
func (europ europeTimeZones) Mariehamn() string {return "Europe/Mariehamn" }

// Minsk - IANA Time Zone 'Europe/Minsk'.
// IANA Source File: Minsk
//  
func (europ europeTimeZones) Minsk() string {return "Europe/Minsk" }

// Monaco - IANA Time Zone 'Europe/Monaco'.
// IANA Source File: Monaco
//  
func (europ europeTimeZones) Monaco() string {return "Europe/Monaco" }

// Moscow - IANA Time Zone 'Europe/Moscow'.
// IANA Source File: Moscow
//  
func (europ europeTimeZones) Moscow() string {return "Europe/Moscow" }

// Nicosia - IANA Time Zone 'Europe/Nicosia'.
// IANA Source File: Nicosia
//  
func (europ europeTimeZones) Nicosia() string {return "Europe/Nicosia" }

// Oslo - IANA Time Zone 'Europe/Oslo'.
// IANA Source File: Oslo
//  
func (europ europeTimeZones) Oslo() string {return "Europe/Oslo" }

// Paris - IANA Time Zone 'Europe/Paris'.
// IANA Source File: Paris
//  
func (europ europeTimeZones) Paris() string {return "Europe/Paris" }

// Podgorica - IANA Time Zone 'Europe/Podgorica'.
// IANA Source File: Podgorica
//  
func (europ europeTimeZones) Podgorica() string {return "Europe/Podgorica" }

// Prague - IANA Time Zone 'Europe/Prague'.
// IANA Source File: Prague
//  
func (europ europeTimeZones) Prague() string {return "Europe/Prague" }

// Riga - IANA Time Zone 'Europe/Riga'.
// IANA Source File: Riga
//  
func (europ europeTimeZones) Riga() string {return "Europe/Riga" }

// Rome - IANA Time Zone 'Europe/Rome'.
// IANA Source File: Rome
//  
func (europ europeTimeZones) Rome() string {return "Europe/Rome" }

// Samara - IANA Time Zone 'Europe/Samara'.
// IANA Source File: Samara
//  
func (europ europeTimeZones) Samara() string {return "Europe/Samara" }

// San_Marino - IANA Time Zone 'Europe/San_Marino'.
// IANA Source File: San_Marino
//  
func (europ europeTimeZones) San_Marino() string {return "Europe/San_Marino" }

// Sarajevo - IANA Time Zone 'Europe/Sarajevo'.
// IANA Source File: Sarajevo
//  
func (europ europeTimeZones) Sarajevo() string {return "Europe/Sarajevo" }

// Saratov - IANA Time Zone 'Europe/Saratov'.
// IANA Source File: Saratov
//  
func (europ europeTimeZones) Saratov() string {return "Europe/Saratov" }

// Simferopol - IANA Time Zone 'Europe/Simferopol'.
// IANA Source File: Simferopol
//  
func (europ europeTimeZones) Simferopol() string {return "Europe/Simferopol" }

// Skopje - IANA Time Zone 'Europe/Skopje'.
// IANA Source File: Skopje
//  
func (europ europeTimeZones) Skopje() string {return "Europe/Skopje" }

// Sofia - IANA Time Zone 'Europe/Sofia'.
// IANA Source File: Sofia
//  
func (europ europeTimeZones) Sofia() string {return "Europe/Sofia" }

// Stockholm - IANA Time Zone 'Europe/Stockholm'.
// IANA Source File: Stockholm
//  
func (europ europeTimeZones) Stockholm() string {return "Europe/Stockholm" }

// Tallinn - IANA Time Zone 'Europe/Tallinn'.
// IANA Source File: Tallinn
//  
func (europ europeTimeZones) Tallinn() string {return "Europe/Tallinn" }

// Tirane - IANA Time Zone 'Europe/Tirane'.
// IANA Source File: Tirane
//  
func (europ europeTimeZones) Tirane() string {return "Europe/Tirane" }

// Tiraspol - IANA Time Zone 'Europe/Tiraspol'.
// IANA Source File: Tiraspol
//  
func (europ europeTimeZones) Tiraspol() string {return "Europe/Tiraspol" }

// Ulyanovsk - IANA Time Zone 'Europe/Ulyanovsk'.
// IANA Source File: Ulyanovsk
//  
func (europ europeTimeZones) Ulyanovsk() string {return "Europe/Ulyanovsk" }

// Uzhgorod - IANA Time Zone 'Europe/Uzhgorod'.
// IANA Source File: Uzhgorod
//  
func (europ europeTimeZones) Uzhgorod() string {return "Europe/Uzhgorod" }

// Vaduz - IANA Time Zone 'Europe/Vaduz'.
// IANA Source File: Vaduz
//  
func (europ europeTimeZones) Vaduz() string {return "Europe/Vaduz" }

// Vatican - IANA Time Zone 'Europe/Vatican'.
// IANA Source File: Vatican
//  
func (europ europeTimeZones) Vatican() string {return "Europe/Vatican" }

// Vienna - IANA Time Zone 'Europe/Vienna'.
// IANA Source File: Vienna
//  
func (europ europeTimeZones) Vienna() string {return "Europe/Vienna" }

// Vilnius - IANA Time Zone 'Europe/Vilnius'.
// IANA Source File: Vilnius
//  
func (europ europeTimeZones) Vilnius() string {return "Europe/Vilnius" }

// Volgograd - IANA Time Zone 'Europe/Volgograd'.
// IANA Source File: Volgograd
//  
func (europ europeTimeZones) Volgograd() string {return "Europe/Volgograd" }

// Warsaw - IANA Time Zone 'Europe/Warsaw'.
// IANA Source File: Warsaw
//  
func (europ europeTimeZones) Warsaw() string {return "Europe/Warsaw" }

// Zagreb - IANA Time Zone 'Europe/Zagreb'.
// IANA Source File: Zagreb
//  
func (europ europeTimeZones) Zagreb() string {return "Europe/Zagreb" }

// Zaporozhye - IANA Time Zone 'Europe/Zaporozhye'.
// IANA Source File: Zaporozhye
//  
func (europ europeTimeZones) Zaporozhye() string {return "Europe/Zaporozhye" }

// Zurich - IANA Time Zone 'Europe/Zurich'.
// IANA Source File: Zurich
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
// IANA Source File: Antananarivo
//  
func (india indianTimeZones) Antananarivo() string {return "Indian/Antananarivo" }

// Chagos - IANA Time Zone 'Indian/Chagos'.
// IANA Source File: Chagos
//  
func (india indianTimeZones) Chagos() string {return "Indian/Chagos" }

// Christmas - IANA Time Zone 'Indian/Christmas'.
// IANA Source File: Christmas
//  
func (india indianTimeZones) Christmas() string {return "Indian/Christmas" }

// Cocos - IANA Time Zone 'Indian/Cocos'.
// IANA Source File: Cocos
//  
func (india indianTimeZones) Cocos() string {return "Indian/Cocos" }

// Comoro - IANA Time Zone 'Indian/Comoro'.
// IANA Source File: Comoro
//  
func (india indianTimeZones) Comoro() string {return "Indian/Comoro" }

// Kerguelen - IANA Time Zone 'Indian/Kerguelen'.
// IANA Source File: Kerguelen
//  
func (india indianTimeZones) Kerguelen() string {return "Indian/Kerguelen" }

// Mahe - IANA Time Zone 'Indian/Mahe'.
// IANA Source File: Mahe
//  
func (india indianTimeZones) Mahe() string {return "Indian/Mahe" }

// Maldives - IANA Time Zone 'Indian/Maldives'.
// IANA Source File: Maldives
//  
func (india indianTimeZones) Maldives() string {return "Indian/Maldives" }

// Mauritius - IANA Time Zone 'Indian/Mauritius'.
// IANA Source File: Mauritius
//  
func (india indianTimeZones) Mauritius() string {return "Indian/Mauritius" }

// Mayotte - IANA Time Zone 'Indian/Mayotte'.
// IANA Source File: Mayotte
//  
func (india indianTimeZones) Mayotte() string {return "Indian/Mayotte" }

// Reunion - IANA Time Zone 'Indian/Reunion'.
// IANA Source File: Reunion
//  
func (india indianTimeZones) Reunion() string {return "Indian/Reunion" }


// mexicoTimeZones - IANA Time Zones for 'Mexico'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type mexicoTimeZones string

// BajaNorte - IANA Time Zone 'Mexico/BajaNorte'.
// IANA Source File: BajaNorte
//  
func (mexic mexicoTimeZones) BajaNorte() string {return "Mexico/BajaNorte" }

// BajaSur - IANA Time Zone 'Mexico/BajaSur'.
// IANA Source File: BajaSur
//  
func (mexic mexicoTimeZones) BajaSur() string {return "Mexico/BajaSur" }

// General - IANA Time Zone 'Mexico/General'.
// IANA Source File: General
//  
func (mexic mexicoTimeZones) General() string {return "Mexico/General" }


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

// CET - IANA Time Zone 'CET'.
// IANA Source File: CET
//  
func (other otherTimeZones) CET() string {return "CET" }

// CST6CDT - IANA Time Zone 'CST6CDT'.
// IANA Source File: CST6CDT
//  
func (other otherTimeZones) CST06CDT() string {return "CST6CDT" }

// Cuba - IANA Time Zone 'Cuba'.
// IANA Source File: Cuba
//  
func (other otherTimeZones) Cuba() string {return "Cuba" }

// EET - IANA Time Zone 'EET'.
// IANA Source File: EET
//  
func (other otherTimeZones) EET() string {return "EET" }

// Egypt - IANA Time Zone 'Egypt'.
// IANA Source File: Egypt
//  
func (other otherTimeZones) Egypt() string {return "Egypt" }

// Eire - IANA Time Zone 'Eire'.
// IANA Source File: Eire
//  
func (other otherTimeZones) Eire() string {return "Eire" }

// EST - IANA Time Zone 'EST'.
// IANA Source File: EST
//  
func (other otherTimeZones) EST() string {return "EST" }

// EST5EDT - IANA Time Zone 'EST5EDT'.
// IANA Source File: EST5EDT
//  
func (other otherTimeZones) EST05EDT() string {return "EST5EDT" }

// Factory - IANA Time Zone 'Factory'.
// IANA Source File: Factory
//  
func (other otherTimeZones) Factory() string {return "Factory" }

// GB - IANA Time Zone 'GB'.
// IANA Source File: GB
//  
func (other otherTimeZones) GB() string {return "GB" }

// GB-Eire - IANA Time Zone 'GB-Eire'.
// IANA Source File: GB-Eire
//  
func (other otherTimeZones) GBMinusEire() string {return "GB-Eire" }

// GMT - IANA Time Zone 'GMT'.
// IANA Source File: GMT
//  
func (other otherTimeZones) GMT() string {return "GMT" }

// GMT+0 - IANA Time Zone 'GMT+0'.
// IANA Source File: GMT+0
//  
func (other otherTimeZones) GMTPlus00() string {return "GMT+0" }

// GMT-0 - IANA Time Zone 'GMT-0'.
// IANA Source File: GMT-0
//  
func (other otherTimeZones) GMTMinus00() string {return "GMT-0" }

// GMT0 - IANA Time Zone 'GMT0'.
// IANA Source File: GMT0
//  
func (other otherTimeZones) GMT00() string {return "GMT0" }

// Greenwich - IANA Time Zone 'Greenwich'.
// IANA Source File: Greenwich
//  
func (other otherTimeZones) Greenwich() string {return "Greenwich" }

// Hongkong - IANA Time Zone 'Hongkong'.
// IANA Source File: Hongkong
//  
func (other otherTimeZones) Hongkong() string {return "Hongkong" }

// HST - IANA Time Zone 'HST'.
// IANA Source File: HST
//  
func (other otherTimeZones) HST() string {return "HST" }

// Iceland - IANA Time Zone 'Iceland'.
// IANA Source File: Iceland
//  
func (other otherTimeZones) Iceland() string {return "Iceland" }

// Iran - IANA Time Zone 'Iran'.
// IANA Source File: Iran
//  
func (other otherTimeZones) Iran() string {return "Iran" }

// Israel - IANA Time Zone 'Israel'.
// IANA Source File: Israel
//  
func (other otherTimeZones) Israel() string {return "Israel" }

// Jamaica - IANA Time Zone 'Jamaica'.
// IANA Source File: Jamaica
//  
func (other otherTimeZones) Jamaica() string {return "Jamaica" }

// Japan - IANA Time Zone 'Japan'.
// IANA Source File: Japan
//  
func (other otherTimeZones) Japan() string {return "Japan" }

// Kwajalein - IANA Time Zone 'Kwajalein'.
// IANA Source File: Kwajalein
//  
func (other otherTimeZones) Kwajalein() string {return "Kwajalein" }

// Libya - IANA Time Zone 'Libya'.
// IANA Source File: Libya
//  
func (other otherTimeZones) Libya() string {return "Libya" }

// MET - IANA Time Zone 'MET'.
// IANA Source File: MET
//  
func (other otherTimeZones) MET() string {return "MET" }

// MST - IANA Time Zone 'MST'.
// IANA Source File: MST
//  
func (other otherTimeZones) MST() string {return "MST" }

// MST7MDT - IANA Time Zone 'MST7MDT'.
// IANA Source File: MST7MDT
//  
func (other otherTimeZones) MST07MDT() string {return "MST7MDT" }

// Navajo - IANA Time Zone 'Navajo'.
// IANA Source File: Navajo
//  
func (other otherTimeZones) Navajo() string {return "Navajo" }

// NZ - IANA Time Zone 'NZ'.
// IANA Source File: NZ
//  
func (other otherTimeZones) NZ() string {return "NZ" }

// NZ-CHAT - IANA Time Zone 'NZ-CHAT'.
// IANA Source File: NZ-CHAT
//  
func (other otherTimeZones) NZMinusCHAT() string {return "NZ-CHAT" }

// Poland - IANA Time Zone 'Poland'.
// IANA Source File: Poland
//  
func (other otherTimeZones) Poland() string {return "Poland" }

// Portugal - IANA Time Zone 'Portugal'.
// IANA Source File: Portugal
//  
func (other otherTimeZones) Portugal() string {return "Portugal" }

// PRC - IANA Time Zone 'PRC'.
// IANA Source File: PRC
//  
func (other otherTimeZones) PRC() string {return "PRC" }

// PST8PDT - IANA Time Zone 'PST8PDT'.
// IANA Source File: PST8PDT
//  
func (other otherTimeZones) PST08PDT() string {return "PST8PDT" }

// ROC - IANA Time Zone 'ROC'.
// IANA Source File: ROC
//  
func (other otherTimeZones) ROC() string {return "ROC" }

// ROK - IANA Time Zone 'ROK'.
// IANA Source File: ROK
//  
func (other otherTimeZones) ROK() string {return "ROK" }

// Singapore - IANA Time Zone 'Singapore'.
// IANA Source File: Singapore
//  
func (other otherTimeZones) Singapore() string {return "Singapore" }

// Turkey - IANA Time Zone 'Turkey'.
// IANA Source File: Turkey
//  
func (other otherTimeZones) Turkey() string {return "Turkey" }

// UCT - IANA Time Zone 'UCT'.
// IANA Source File: UCT
//  
func (other otherTimeZones) UCT() string {return "UCT" }

// Universal - IANA Time Zone 'Universal'.
// IANA Source File: Universal
//  
func (other otherTimeZones) Universal() string {return "Universal" }

// UTC - IANA Time Zone 'UTC'.
// IANA Source File: UTC
//  
func (other otherTimeZones) UTC() string {return "UTC" }

// W-SU - IANA Time Zone 'W-SU'.
// IANA Source File: W-SU
//  
func (other otherTimeZones) WMinusSU() string {return "W-SU" }

// WET - IANA Time Zone 'WET'.
// IANA Source File: WET
//  
func (other otherTimeZones) WET() string {return "WET" }

// Zulu - IANA Time Zone 'Zulu'.
// IANA Source File: Zulu
//  
func (other otherTimeZones) Zulu() string {return "Zulu" }


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
// IANA Source File: Apia
//  
func (pacif pacificTimeZones) Apia() string {return "Pacific/Apia" }

// Auckland - IANA Time Zone 'Pacific/Auckland'.
// IANA Source File: Auckland
//  
func (pacif pacificTimeZones) Auckland() string {return "Pacific/Auckland" }

// Bougainville - IANA Time Zone 'Pacific/Bougainville'.
// IANA Source File: Bougainville
//  
func (pacif pacificTimeZones) Bougainville() string {return "Pacific/Bougainville" }

// Chatham - IANA Time Zone 'Pacific/Chatham'.
// IANA Source File: Chatham
//  
func (pacif pacificTimeZones) Chatham() string {return "Pacific/Chatham" }

// Chuuk - IANA Time Zone 'Pacific/Chuuk'.
// IANA Source File: Chuuk
//  
func (pacif pacificTimeZones) Chuuk() string {return "Pacific/Chuuk" }

// Easter - IANA Time Zone 'Pacific/Easter'.
// IANA Source File: Easter
//  
func (pacif pacificTimeZones) Easter() string {return "Pacific/Easter" }

// Efate - IANA Time Zone 'Pacific/Efate'.
// IANA Source File: Efate
//  
func (pacif pacificTimeZones) Efate() string {return "Pacific/Efate" }

// Enderbury - IANA Time Zone 'Pacific/Enderbury'.
// IANA Source File: Enderbury
//  
func (pacif pacificTimeZones) Enderbury() string {return "Pacific/Enderbury" }

// Fakaofo - IANA Time Zone 'Pacific/Fakaofo'.
// IANA Source File: Fakaofo
//  
func (pacif pacificTimeZones) Fakaofo() string {return "Pacific/Fakaofo" }

// Fiji - IANA Time Zone 'Pacific/Fiji'.
// IANA Source File: Fiji
//  
func (pacif pacificTimeZones) Fiji() string {return "Pacific/Fiji" }

// Funafuti - IANA Time Zone 'Pacific/Funafuti'.
// IANA Source File: Funafuti
//  
func (pacif pacificTimeZones) Funafuti() string {return "Pacific/Funafuti" }

// Galapagos - IANA Time Zone 'Pacific/Galapagos'.
// IANA Source File: Galapagos
//  
func (pacif pacificTimeZones) Galapagos() string {return "Pacific/Galapagos" }

// Gambier - IANA Time Zone 'Pacific/Gambier'.
// IANA Source File: Gambier
//  
func (pacif pacificTimeZones) Gambier() string {return "Pacific/Gambier" }

// Guadalcanal - IANA Time Zone 'Pacific/Guadalcanal'.
// IANA Source File: Guadalcanal
//  
func (pacif pacificTimeZones) Guadalcanal() string {return "Pacific/Guadalcanal" }

// Guam - IANA Time Zone 'Pacific/Guam'.
// IANA Source File: Guam
//  
func (pacif pacificTimeZones) Guam() string {return "Pacific/Guam" }

// Honolulu - IANA Time Zone 'Pacific/Honolulu'.
// IANA Source File: Honolulu
//  
func (pacif pacificTimeZones) Honolulu() string {return "Pacific/Honolulu" }

// Johnston - IANA Time Zone 'Pacific/Johnston'.
// IANA Source File: Johnston
//  
func (pacif pacificTimeZones) Johnston() string {return "Pacific/Johnston" }

// Kiritimati - IANA Time Zone 'Pacific/Kiritimati'.
// IANA Source File: Kiritimati
//  
func (pacif pacificTimeZones) Kiritimati() string {return "Pacific/Kiritimati" }

// Kosrae - IANA Time Zone 'Pacific/Kosrae'.
// IANA Source File: Kosrae
//  
func (pacif pacificTimeZones) Kosrae() string {return "Pacific/Kosrae" }

// Kwajalein - IANA Time Zone 'Pacific/Kwajalein'.
// IANA Source File: Kwajalein
//  
func (pacif pacificTimeZones) Kwajalein() string {return "Pacific/Kwajalein" }

// Majuro - IANA Time Zone 'Pacific/Majuro'.
// IANA Source File: Majuro
//  
func (pacif pacificTimeZones) Majuro() string {return "Pacific/Majuro" }

// Marquesas - IANA Time Zone 'Pacific/Marquesas'.
// IANA Source File: Marquesas
//  
func (pacif pacificTimeZones) Marquesas() string {return "Pacific/Marquesas" }

// Midway - IANA Time Zone 'Pacific/Midway'.
// IANA Source File: Midway
//  
func (pacif pacificTimeZones) Midway() string {return "Pacific/Midway" }

// Nauru - IANA Time Zone 'Pacific/Nauru'.
// IANA Source File: Nauru
//  
func (pacif pacificTimeZones) Nauru() string {return "Pacific/Nauru" }

// Niue - IANA Time Zone 'Pacific/Niue'.
// IANA Source File: Niue
//  
func (pacif pacificTimeZones) Niue() string {return "Pacific/Niue" }

// Norfolk - IANA Time Zone 'Pacific/Norfolk'.
// IANA Source File: Norfolk
//  
func (pacif pacificTimeZones) Norfolk() string {return "Pacific/Norfolk" }

// Noumea - IANA Time Zone 'Pacific/Noumea'.
// IANA Source File: Noumea
//  
func (pacif pacificTimeZones) Noumea() string {return "Pacific/Noumea" }

// Pago_Pago - IANA Time Zone 'Pacific/Pago_Pago'.
// IANA Source File: Pago_Pago
//  
func (pacif pacificTimeZones) Pago_Pago() string {return "Pacific/Pago_Pago" }

// Palau - IANA Time Zone 'Pacific/Palau'.
// IANA Source File: Palau
//  
func (pacif pacificTimeZones) Palau() string {return "Pacific/Palau" }

// Pitcairn - IANA Time Zone 'Pacific/Pitcairn'.
// IANA Source File: Pitcairn
//  
func (pacif pacificTimeZones) Pitcairn() string {return "Pacific/Pitcairn" }

// Pohnpei - IANA Time Zone 'Pacific/Pohnpei'.
// IANA Source File: Pohnpei
//  
func (pacif pacificTimeZones) Pohnpei() string {return "Pacific/Pohnpei" }

// Ponape - IANA Time Zone 'Pacific/Ponape'.
// IANA Source File: Ponape
//  
func (pacif pacificTimeZones) Ponape() string {return "Pacific/Ponape" }

// Port_Moresby - IANA Time Zone 'Pacific/Port_Moresby'.
// IANA Source File: Port_Moresby
//  
func (pacif pacificTimeZones) Port_Moresby() string {return "Pacific/Port_Moresby" }

// Rarotonga - IANA Time Zone 'Pacific/Rarotonga'.
// IANA Source File: Rarotonga
//  
func (pacif pacificTimeZones) Rarotonga() string {return "Pacific/Rarotonga" }

// Saipan - IANA Time Zone 'Pacific/Saipan'.
// IANA Source File: Saipan
//  
func (pacif pacificTimeZones) Saipan() string {return "Pacific/Saipan" }

// Samoa - IANA Time Zone 'Pacific/Samoa'.
// IANA Source File: Samoa
//  
func (pacif pacificTimeZones) Samoa() string {return "Pacific/Samoa" }

// Tahiti - IANA Time Zone 'Pacific/Tahiti'.
// IANA Source File: Tahiti
//  
func (pacif pacificTimeZones) Tahiti() string {return "Pacific/Tahiti" }

// Tarawa - IANA Time Zone 'Pacific/Tarawa'.
// IANA Source File: Tarawa
//  
func (pacif pacificTimeZones) Tarawa() string {return "Pacific/Tarawa" }

// Tongatapu - IANA Time Zone 'Pacific/Tongatapu'.
// IANA Source File: Tongatapu
//  
func (pacif pacificTimeZones) Tongatapu() string {return "Pacific/Tongatapu" }

// Truk - IANA Time Zone 'Pacific/Truk'.
// IANA Source File: Truk
//  
func (pacif pacificTimeZones) Truk() string {return "Pacific/Truk" }

// Wake - IANA Time Zone 'Pacific/Wake'.
// IANA Source File: Wake
//  
func (pacif pacificTimeZones) Wake() string {return "Pacific/Wake" }

// Wallis - IANA Time Zone 'Pacific/Wallis'.
// IANA Source File: Wallis
//  
func (pacif pacificTimeZones) Wallis() string {return "Pacific/Wallis" }

// Yap - IANA Time Zone 'Pacific/Yap'.
// IANA Source File: Yap
//  
func (pacif pacificTimeZones) Yap() string {return "Pacific/Yap" }


// uSTimeZones - IANA Time Zones for 'US'.
//  
// For documentation on IANA Time Zones, see type
// 'TimeZones'.
//  
// Reference:
//   https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//   https://en.wikipedia.org/wiki/Tz_database
//   https://www.iana.org/time-zones
//  
type uSTimeZones string

// Alaska - IANA Time Zone 'US/Alaska'.
// IANA Source File: Alaska
//  
func (uSTim uSTimeZones) Alaska() string {return "US/Alaska" }

// Aleutian - IANA Time Zone 'US/Aleutian'.
// IANA Source File: Aleutian
//  
func (uSTim uSTimeZones) Aleutian() string {return "US/Aleutian" }

// Arizona - IANA Time Zone 'US/Arizona'.
// IANA Source File: Arizona
//  
func (uSTim uSTimeZones) Arizona() string {return "US/Arizona" }

// Central - IANA Time Zone 'US/Central'.
// IANA Source File: Central
//  
func (uSTim uSTimeZones) Central() string {return "US/Central" }

// East-Indiana - IANA Time Zone 'US/East-Indiana'.
// IANA Source File: East-Indiana
//  
func (uSTim uSTimeZones) EastMinusIndiana() string {return "US/East-Indiana" }

// Eastern - IANA Time Zone 'US/Eastern'.
// IANA Source File: Eastern
//  
func (uSTim uSTimeZones) Eastern() string {return "US/Eastern" }

// Hawaii - IANA Time Zone 'US/Hawaii'.
// IANA Source File: Hawaii
//  
func (uSTim uSTimeZones) Hawaii() string {return "US/Hawaii" }

// Indiana-Starke - IANA Time Zone 'US/Indiana-Starke'.
// IANA Source File: Indiana-Starke
//  
func (uSTim uSTimeZones) IndianaMinusStarke() string {return "US/Indiana-Starke" }

// Michigan - IANA Time Zone 'US/Michigan'.
// IANA Source File: Michigan
//  
func (uSTim uSTimeZones) Michigan() string {return "US/Michigan" }

// Mountain - IANA Time Zone 'US/Mountain'.
// IANA Source File: Mountain
//  
func (uSTim uSTimeZones) Mountain() string {return "US/Mountain" }

// Pacific - IANA Time Zone 'US/Pacific'.
// IANA Source File: Pacific
//  
func (uSTim uSTimeZones) Pacific() string {return "US/Pacific" }

// Samoa - IANA Time Zone 'US/Samoa'.
// IANA Source File: Samoa
//  
func (uSTim uSTimeZones) Samoa() string {return "US/Samoa" }


// argentinaTimeZones - A Sub-Group of Time Zones. These are
// IANA Time Zones located in 'Argentina'.
//  
// The Parent Group is 'America'.
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

// Buenos_Aires - IANA Time Zone 'America/Argentina/Buenos_Aires'.
// IANA Source File: Buenos_Aires
//  
func (argen argentinaTimeZones) Buenos_Aires() string {return "America/Argentina/Buenos_Aires" }

// Catamarca - IANA Time Zone 'America/Argentina/Catamarca'.
// IANA Source File: Catamarca
//  
func (argen argentinaTimeZones) Catamarca() string {return "America/Argentina/Catamarca" }

// ComodRivadavia - IANA Time Zone 'America/Argentina/ComodRivadavia'.
// IANA Source File: ComodRivadavia
//  
func (argen argentinaTimeZones) ComodRivadavia() string {return "America/Argentina/ComodRivadavia" }

// Cordoba - IANA Time Zone 'America/Argentina/Cordoba'.
// IANA Source File: Cordoba
//  
func (argen argentinaTimeZones) Cordoba() string {return "America/Argentina/Cordoba" }

// Jujuy - IANA Time Zone 'America/Argentina/Jujuy'.
// IANA Source File: Jujuy
//  
func (argen argentinaTimeZones) Jujuy() string {return "America/Argentina/Jujuy" }

// La_Rioja - IANA Time Zone 'America/Argentina/La_Rioja'.
// IANA Source File: La_Rioja
//  
func (argen argentinaTimeZones) La_Rioja() string {return "America/Argentina/La_Rioja" }

// Mendoza - IANA Time Zone 'America/Argentina/Mendoza'.
// IANA Source File: Mendoza
//  
func (argen argentinaTimeZones) Mendoza() string {return "America/Argentina/Mendoza" }

// Rio_Gallegos - IANA Time Zone 'America/Argentina/Rio_Gallegos'.
// IANA Source File: Rio_Gallegos
//  
func (argen argentinaTimeZones) Rio_Gallegos() string {return "America/Argentina/Rio_Gallegos" }

// Salta - IANA Time Zone 'America/Argentina/Salta'.
// IANA Source File: Salta
//  
func (argen argentinaTimeZones) Salta() string {return "America/Argentina/Salta" }

// San_Juan - IANA Time Zone 'America/Argentina/San_Juan'.
// IANA Source File: San_Juan
//  
func (argen argentinaTimeZones) San_Juan() string {return "America/Argentina/San_Juan" }

// San_Luis - IANA Time Zone 'America/Argentina/San_Luis'.
// IANA Source File: San_Luis
//  
func (argen argentinaTimeZones) San_Luis() string {return "America/Argentina/San_Luis" }

// Tucuman - IANA Time Zone 'America/Argentina/Tucuman'.
// IANA Source File: Tucuman
//  
func (argen argentinaTimeZones) Tucuman() string {return "America/Argentina/Tucuman" }

// Ushuaia - IANA Time Zone 'America/Argentina/Ushuaia'.
// IANA Source File: Ushuaia
//  
func (argen argentinaTimeZones) Ushuaia() string {return "America/Argentina/Ushuaia" }


// indianaTimeZones - A Sub-Group of Time Zones. These are
// IANA Time Zones located in 'Indiana'.
//  
// The Parent Group is 'America'.
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

// Indianapolis - IANA Time Zone 'America/Indiana/Indianapolis'.
// IANA Source File: Indianapolis
//  
func (india indianaTimeZones) Indianapolis() string {return "America/Indiana/Indianapolis" }

// Knox - IANA Time Zone 'America/Indiana/Knox'.
// IANA Source File: Knox
//  
func (india indianaTimeZones) Knox() string {return "America/Indiana/Knox" }

// Marengo - IANA Time Zone 'America/Indiana/Marengo'.
// IANA Source File: Marengo
//  
func (india indianaTimeZones) Marengo() string {return "America/Indiana/Marengo" }

// Petersburg - IANA Time Zone 'America/Indiana/Petersburg'.
// IANA Source File: Petersburg
//  
func (india indianaTimeZones) Petersburg() string {return "America/Indiana/Petersburg" }

// Tell_City - IANA Time Zone 'America/Indiana/Tell_City'.
// IANA Source File: Tell_City
//  
func (india indianaTimeZones) Tell_City() string {return "America/Indiana/Tell_City" }

// Vevay - IANA Time Zone 'America/Indiana/Vevay'.
// IANA Source File: Vevay
//  
func (india indianaTimeZones) Vevay() string {return "America/Indiana/Vevay" }

// Vincennes - IANA Time Zone 'America/Indiana/Vincennes'.
// IANA Source File: Vincennes
//  
func (india indianaTimeZones) Vincennes() string {return "America/Indiana/Vincennes" }

// Winamac - IANA Time Zone 'America/Indiana/Winamac'.
// IANA Source File: Winamac
//  
func (india indianaTimeZones) Winamac() string {return "America/Indiana/Winamac" }


// kentuckyTimeZones - A Sub-Group of Time Zones. These are
// IANA Time Zones located in 'Kentucky'.
//  
// The Parent Group is 'America'.
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

// Louisville - IANA Time Zone 'America/Kentucky/Louisville'.
// IANA Source File: Louisville
//  
func (kentu kentuckyTimeZones) Louisville() string {return "America/Kentucky/Louisville" }

// Monticello - IANA Time Zone 'America/Kentucky/Monticello'.
// IANA Source File: Monticello
//  
func (kentu kentuckyTimeZones) Monticello() string {return "America/Kentucky/Monticello" }


// north_DakotaTimeZones - A Sub-Group of Time Zones. These are
// IANA Time Zones located in 'North_Dakota'.
//  
// The Parent Group is 'America'.
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

// Beulah - IANA Time Zone 'America/North_Dakota/Beulah'.
// IANA Source File: Beulah
//  
func (north north_DakotaTimeZones) Beulah() string {return "America/North_Dakota/Beulah" }

// Center - IANA Time Zone 'America/North_Dakota/Center'.
// IANA Source File: Center
//  
func (north north_DakotaTimeZones) Center() string {return "America/North_Dakota/Center" }

// New_Salem - IANA Time Zone 'America/North_Dakota/New_Salem'.
// IANA Source File: New_Salem
//  
func (north north_DakotaTimeZones) New_Salem() string {return "America/North_Dakota/New_Salem" }

