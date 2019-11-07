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
// America                                   26        0        0
// Antarctica                                 0        0        0
// Asia                                       0        0        0
// Atlantic                                   1        0        0
// Australia                                  0        0        0
// Europe                                     0        0        0
// Indian                                     2        0        0
// Pacific                                    2        0        0
// Etc                                        0        0        0
// Other                                    562        0        0
// ==============================================================
//                              Total         0        0        0
//
// ----------------------------------------------------------------------------
// 
// This TimeZones Type is based on IANA Time Zone Database Version: 
// 
//           IANA Standard Time Zones : 593
//           IANA Link Time Zones     :   0
//                                         -------
//                 Sub-Total IANA Time Zones:   0
// 
//                Military Time Zones :  25
//                   Other Time Zones :   0
//                                         -------
//                          Total Time Zones:   0
// 
//       Standard Time Zone Sub-Groups:   4
//           Link Time Zone Sub-Groups:   0
//                                         -------
//                Total Time Zone Sub-Groups:   0
// 
//                  Primary Time Zone Groups:  17
// 
// Type Creation Date: 2019-11-06 Wednesday 18:01:37 -0600 CST
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

// Abidjan - IANA Time Zone 'Abidjan'.
// IANA Source File: Abidjan
//  
func (afric africaTimeZones) Abidjan() string {return "Abidjan" }

// Accra - IANA Time Zone 'Accra'.
// IANA Source File: Accra
//  
func (afric africaTimeZones) Accra() string {return "Accra" }

// Addis_Ababa - IANA Time Zone 'Addis_Ababa'.
// IANA Source File: Addis_Ababa
//  
func (afric africaTimeZones) Addis_Ababa() string {return "Addis_Ababa" }

// Algiers - IANA Time Zone 'Algiers'.
// IANA Source File: Algiers
//  
func (afric africaTimeZones) Algiers() string {return "Algiers" }

// Asmara - IANA Time Zone 'Asmara'.
// IANA Source File: Asmara
//  
func (afric africaTimeZones) Asmara() string {return "Asmara" }

// Asmera - IANA Time Zone 'Asmera'.
// IANA Source File: Asmera
//  
func (afric africaTimeZones) Asmera() string {return "Asmera" }

// Bamako - IANA Time Zone 'Bamako'.
// IANA Source File: Bamako
//  
func (afric africaTimeZones) Bamako() string {return "Bamako" }

// Bangui - IANA Time Zone 'Bangui'.
// IANA Source File: Bangui
//  
func (afric africaTimeZones) Bangui() string {return "Bangui" }

// Banjul - IANA Time Zone 'Banjul'.
// IANA Source File: Banjul
//  
func (afric africaTimeZones) Banjul() string {return "Banjul" }

// Bissau - IANA Time Zone 'Bissau'.
// IANA Source File: Bissau
//  
func (afric africaTimeZones) Bissau() string {return "Bissau" }

// Blantyre - IANA Time Zone 'Blantyre'.
// IANA Source File: Blantyre
//  
func (afric africaTimeZones) Blantyre() string {return "Blantyre" }

// Brazzaville - IANA Time Zone 'Brazzaville'.
// IANA Source File: Brazzaville
//  
func (afric africaTimeZones) Brazzaville() string {return "Brazzaville" }

// Bujumbura - IANA Time Zone 'Bujumbura'.
// IANA Source File: Bujumbura
//  
func (afric africaTimeZones) Bujumbura() string {return "Bujumbura" }

// Cairo - IANA Time Zone 'Cairo'.
// IANA Source File: Cairo
//  
func (afric africaTimeZones) Cairo() string {return "Cairo" }

// Casablanca - IANA Time Zone 'Casablanca'.
// IANA Source File: Casablanca
//  
func (afric africaTimeZones) Casablanca() string {return "Casablanca" }

// Ceuta - IANA Time Zone 'Ceuta'.
// IANA Source File: Ceuta
//  
func (afric africaTimeZones) Ceuta() string {return "Ceuta" }

// Conakry - IANA Time Zone 'Conakry'.
// IANA Source File: Conakry
//  
func (afric africaTimeZones) Conakry() string {return "Conakry" }

// Dakar - IANA Time Zone 'Dakar'.
// IANA Source File: Dakar
//  
func (afric africaTimeZones) Dakar() string {return "Dakar" }

// Dar_es_Salaam - IANA Time Zone 'Dar_es_Salaam'.
// IANA Source File: Dar_es_Salaam
//  
func (afric africaTimeZones) Dar_es_Salaam() string {return "Dar_es_Salaam" }

// Djibouti - IANA Time Zone 'Djibouti'.
// IANA Source File: Djibouti
//  
func (afric africaTimeZones) Djibouti() string {return "Djibouti" }

// Douala - IANA Time Zone 'Douala'.
// IANA Source File: Douala
//  
func (afric africaTimeZones) Douala() string {return "Douala" }

// El_Aaiun - IANA Time Zone 'El_Aaiun'.
// IANA Source File: El_Aaiun
//  
func (afric africaTimeZones) El_Aaiun() string {return "El_Aaiun" }

// Freetown - IANA Time Zone 'Freetown'.
// IANA Source File: Freetown
//  
func (afric africaTimeZones) Freetown() string {return "Freetown" }

// Gaborone - IANA Time Zone 'Gaborone'.
// IANA Source File: Gaborone
//  
func (afric africaTimeZones) Gaborone() string {return "Gaborone" }

// Harare - IANA Time Zone 'Harare'.
// IANA Source File: Harare
//  
func (afric africaTimeZones) Harare() string {return "Harare" }

// Johannesburg - IANA Time Zone 'Johannesburg'.
// IANA Source File: Johannesburg
//  
func (afric africaTimeZones) Johannesburg() string {return "Johannesburg" }

// Juba - IANA Time Zone 'Juba'.
// IANA Source File: Juba
//  
func (afric africaTimeZones) Juba() string {return "Juba" }

// Kampala - IANA Time Zone 'Kampala'.
// IANA Source File: Kampala
//  
func (afric africaTimeZones) Kampala() string {return "Kampala" }

// Khartoum - IANA Time Zone 'Khartoum'.
// IANA Source File: Khartoum
//  
func (afric africaTimeZones) Khartoum() string {return "Khartoum" }

// Kigali - IANA Time Zone 'Kigali'.
// IANA Source File: Kigali
//  
func (afric africaTimeZones) Kigali() string {return "Kigali" }

// Kinshasa - IANA Time Zone 'Kinshasa'.
// IANA Source File: Kinshasa
//  
func (afric africaTimeZones) Kinshasa() string {return "Kinshasa" }

// Lagos - IANA Time Zone 'Lagos'.
// IANA Source File: Lagos
//  
func (afric africaTimeZones) Lagos() string {return "Lagos" }

// Libreville - IANA Time Zone 'Libreville'.
// IANA Source File: Libreville
//  
func (afric africaTimeZones) Libreville() string {return "Libreville" }

// Lome - IANA Time Zone 'Lome'.
// IANA Source File: Lome
//  
func (afric africaTimeZones) Lome() string {return "Lome" }

// Luanda - IANA Time Zone 'Luanda'.
// IANA Source File: Luanda
//  
func (afric africaTimeZones) Luanda() string {return "Luanda" }

// Lubumbashi - IANA Time Zone 'Lubumbashi'.
// IANA Source File: Lubumbashi
//  
func (afric africaTimeZones) Lubumbashi() string {return "Lubumbashi" }

// Lusaka - IANA Time Zone 'Lusaka'.
// IANA Source File: Lusaka
//  
func (afric africaTimeZones) Lusaka() string {return "Lusaka" }

// Malabo - IANA Time Zone 'Malabo'.
// IANA Source File: Malabo
//  
func (afric africaTimeZones) Malabo() string {return "Malabo" }

// Maputo - IANA Time Zone 'Maputo'.
// IANA Source File: Maputo
//  
func (afric africaTimeZones) Maputo() string {return "Maputo" }

// Maseru - IANA Time Zone 'Maseru'.
// IANA Source File: Maseru
//  
func (afric africaTimeZones) Maseru() string {return "Maseru" }

// Mbabane - IANA Time Zone 'Mbabane'.
// IANA Source File: Mbabane
//  
func (afric africaTimeZones) Mbabane() string {return "Mbabane" }

// Mogadishu - IANA Time Zone 'Mogadishu'.
// IANA Source File: Mogadishu
//  
func (afric africaTimeZones) Mogadishu() string {return "Mogadishu" }

// Monrovia - IANA Time Zone 'Monrovia'.
// IANA Source File: Monrovia
//  
func (afric africaTimeZones) Monrovia() string {return "Monrovia" }

// Nairobi - IANA Time Zone 'Nairobi'.
// IANA Source File: Nairobi
//  
func (afric africaTimeZones) Nairobi() string {return "Nairobi" }

// Ndjamena - IANA Time Zone 'Ndjamena'.
// IANA Source File: Ndjamena
//  
func (afric africaTimeZones) Ndjamena() string {return "Ndjamena" }

// Niamey - IANA Time Zone 'Niamey'.
// IANA Source File: Niamey
//  
func (afric africaTimeZones) Niamey() string {return "Niamey" }

// Nouakchott - IANA Time Zone 'Nouakchott'.
// IANA Source File: Nouakchott
//  
func (afric africaTimeZones) Nouakchott() string {return "Nouakchott" }

// Ouagadougou - IANA Time Zone 'Ouagadougou'.
// IANA Source File: Ouagadougou
//  
func (afric africaTimeZones) Ouagadougou() string {return "Ouagadougou" }

// Porto-Novo - IANA Time Zone 'Porto-Novo'.
// IANA Source File: Porto-Novo
//  
func (afric africaTimeZones) PortoMinusNovo() string {return "Porto-Novo" }

// Sao_Tome - IANA Time Zone 'Sao_Tome'.
// IANA Source File: Sao_Tome
//  
func (afric africaTimeZones) Sao_Tome() string {return "Sao_Tome" }

// Timbuktu - IANA Time Zone 'Timbuktu'.
// IANA Source File: Timbuktu
//  
func (afric africaTimeZones) Timbuktu() string {return "Timbuktu" }

// Tripoli - IANA Time Zone 'Tripoli'.
// IANA Source File: Tripoli
//  
func (afric africaTimeZones) Tripoli() string {return "Tripoli" }

// Tunis - IANA Time Zone 'Tunis'.
// IANA Source File: Tunis
//  
func (afric africaTimeZones) Tunis() string {return "Tunis" }

// Windhoek - IANA Time Zone 'Windhoek'.
// IANA Source File: Windhoek
//  
func (afric africaTimeZones) Windhoek() string {return "Windhoek" }


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

// Adak - IANA Time Zone 'Adak'.
// IANA Source File: Adak
//  
func (ameri americaTimeZones) Adak() string {return "Adak" }

// Anchorage - IANA Time Zone 'Anchorage'.
// IANA Source File: Anchorage
//  
func (ameri americaTimeZones) Anchorage() string {return "Anchorage" }

// Anguilla - IANA Time Zone 'Anguilla'.
// IANA Source File: Anguilla
//  
func (ameri americaTimeZones) Anguilla() string {return "Anguilla" }

// Antigua - IANA Time Zone 'Antigua'.
// IANA Source File: Antigua
//  
func (ameri americaTimeZones) Antigua() string {return "Antigua" }

// Araguaina - IANA Time Zone 'Araguaina'.
// IANA Source File: Araguaina
//  
func (ameri americaTimeZones) Araguaina() string {return "Araguaina" }

// Argentina - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (ameri americaTimeZones) Argentina() argentinaTimeZones {return "" }

// Aruba - IANA Time Zone 'Aruba'.
// IANA Source File: Aruba
//  
func (ameri americaTimeZones) Aruba() string {return "Aruba" }

// Asuncion - IANA Time Zone 'Asuncion'.
// IANA Source File: Asuncion
//  
func (ameri americaTimeZones) Asuncion() string {return "Asuncion" }

// Atikokan - IANA Time Zone 'Atikokan'.
// IANA Source File: Atikokan
//  
func (ameri americaTimeZones) Atikokan() string {return "Atikokan" }

// Atka - IANA Time Zone 'Atka'.
// IANA Source File: Atka
//  
func (ameri americaTimeZones) Atka() string {return "Atka" }

// Bahia - IANA Time Zone 'Bahia'.
// IANA Source File: Bahia
//  
func (ameri americaTimeZones) Bahia() string {return "Bahia" }

// Bahia_Banderas - IANA Time Zone 'Bahia_Banderas'.
// IANA Source File: Bahia_Banderas
//  
func (ameri americaTimeZones) Bahia_Banderas() string {return "Bahia_Banderas" }

// Barbados - IANA Time Zone 'Barbados'.
// IANA Source File: Barbados
//  
func (ameri americaTimeZones) Barbados() string {return "Barbados" }

// Belem - IANA Time Zone 'Belem'.
// IANA Source File: Belem
//  
func (ameri americaTimeZones) Belem() string {return "Belem" }

// Belize - IANA Time Zone 'Belize'.
// IANA Source File: Belize
//  
func (ameri americaTimeZones) Belize() string {return "Belize" }

// Blanc-Sablon - IANA Time Zone 'Blanc-Sablon'.
// IANA Source File: Blanc-Sablon
//  
func (ameri americaTimeZones) BlancMinusSablon() string {return "Blanc-Sablon" }

// Boa_Vista - IANA Time Zone 'Boa_Vista'.
// IANA Source File: Boa_Vista
//  
func (ameri americaTimeZones) Boa_Vista() string {return "Boa_Vista" }

// Bogota - IANA Time Zone 'Bogota'.
// IANA Source File: Bogota
//  
func (ameri americaTimeZones) Bogota() string {return "Bogota" }

// Boise - IANA Time Zone 'Boise'.
// IANA Source File: Boise
//  
func (ameri americaTimeZones) Boise() string {return "Boise" }

// Buenos_Aires - IANA Time Zone 'Buenos_Aires'.
// IANA Source File: Buenos_Aires
//  
func (ameri americaTimeZones) Buenos_Aires() string {return "Buenos_Aires" }

// Cambridge_Bay - IANA Time Zone 'Cambridge_Bay'.
// IANA Source File: Cambridge_Bay
//  
func (ameri americaTimeZones) Cambridge_Bay() string {return "Cambridge_Bay" }

// Campo_Grande - IANA Time Zone 'Campo_Grande'.
// IANA Source File: Campo_Grande
//  
func (ameri americaTimeZones) Campo_Grande() string {return "Campo_Grande" }

// Cancun - IANA Time Zone 'Cancun'.
// IANA Source File: Cancun
//  
func (ameri americaTimeZones) Cancun() string {return "Cancun" }

// Caracas - IANA Time Zone 'Caracas'.
// IANA Source File: Caracas
//  
func (ameri americaTimeZones) Caracas() string {return "Caracas" }

// Catamarca - IANA Time Zone 'Catamarca'.
// IANA Source File: Catamarca
//  
func (ameri americaTimeZones) Catamarca() string {return "Catamarca" }

// Cayenne - IANA Time Zone 'Cayenne'.
// IANA Source File: Cayenne
//  
func (ameri americaTimeZones) Cayenne() string {return "Cayenne" }

// Cayman - IANA Time Zone 'Cayman'.
// IANA Source File: Cayman
//  
func (ameri americaTimeZones) Cayman() string {return "Cayman" }

// Chicago - IANA Time Zone 'Chicago'.
// IANA Source File: Chicago
//  
func (ameri americaTimeZones) Chicago() string {return "Chicago" }

// Chihuahua - IANA Time Zone 'Chihuahua'.
// IANA Source File: Chihuahua
//  
func (ameri americaTimeZones) Chihuahua() string {return "Chihuahua" }

// Coral_Harbour - IANA Time Zone 'Coral_Harbour'.
// IANA Source File: Coral_Harbour
//  
func (ameri americaTimeZones) Coral_Harbour() string {return "Coral_Harbour" }

// Cordoba - IANA Time Zone 'Cordoba'.
// IANA Source File: Cordoba
//  
func (ameri americaTimeZones) Cordoba() string {return "Cordoba" }

// Costa_Rica - IANA Time Zone 'Costa_Rica'.
// IANA Source File: Costa_Rica
//  
func (ameri americaTimeZones) Costa_Rica() string {return "Costa_Rica" }

// Creston - IANA Time Zone 'Creston'.
// IANA Source File: Creston
//  
func (ameri americaTimeZones) Creston() string {return "Creston" }

// Cuiaba - IANA Time Zone 'Cuiaba'.
// IANA Source File: Cuiaba
//  
func (ameri americaTimeZones) Cuiaba() string {return "Cuiaba" }

// Curacao - IANA Time Zone 'Curacao'.
// IANA Source File: Curacao
//  
func (ameri americaTimeZones) Curacao() string {return "Curacao" }

// Danmarkshavn - IANA Time Zone 'Danmarkshavn'.
// IANA Source File: Danmarkshavn
//  
func (ameri americaTimeZones) Danmarkshavn() string {return "Danmarkshavn" }

// Dawson - IANA Time Zone 'Dawson'.
// IANA Source File: Dawson
//  
func (ameri americaTimeZones) Dawson() string {return "Dawson" }

// Dawson_Creek - IANA Time Zone 'Dawson_Creek'.
// IANA Source File: Dawson_Creek
//  
func (ameri americaTimeZones) Dawson_Creek() string {return "Dawson_Creek" }

// Denver - IANA Time Zone 'Denver'.
// IANA Source File: Denver
//  
func (ameri americaTimeZones) Denver() string {return "Denver" }

// Detroit - IANA Time Zone 'Detroit'.
// IANA Source File: Detroit
//  
func (ameri americaTimeZones) Detroit() string {return "Detroit" }

// Dominica - IANA Time Zone 'Dominica'.
// IANA Source File: Dominica
//  
func (ameri americaTimeZones) Dominica() string {return "Dominica" }

// Edmonton - IANA Time Zone 'Edmonton'.
// IANA Source File: Edmonton
//  
func (ameri americaTimeZones) Edmonton() string {return "Edmonton" }

// Eirunepe - IANA Time Zone 'Eirunepe'.
// IANA Source File: Eirunepe
//  
func (ameri americaTimeZones) Eirunepe() string {return "Eirunepe" }

// El_Salvador - IANA Time Zone 'El_Salvador'.
// IANA Source File: El_Salvador
//  
func (ameri americaTimeZones) El_Salvador() string {return "El_Salvador" }

// Ensenada - IANA Time Zone 'Ensenada'.
// IANA Source File: Ensenada
//  
func (ameri americaTimeZones) Ensenada() string {return "Ensenada" }

// Fort_Nelson - IANA Time Zone 'Fort_Nelson'.
// IANA Source File: Fort_Nelson
//  
func (ameri americaTimeZones) Fort_Nelson() string {return "Fort_Nelson" }

// Fort_Wayne - IANA Time Zone 'Fort_Wayne'.
// IANA Source File: Fort_Wayne
//  
func (ameri americaTimeZones) Fort_Wayne() string {return "Fort_Wayne" }

// Fortaleza - IANA Time Zone 'Fortaleza'.
// IANA Source File: Fortaleza
//  
func (ameri americaTimeZones) Fortaleza() string {return "Fortaleza" }

// Glace_Bay - IANA Time Zone 'Glace_Bay'.
// IANA Source File: Glace_Bay
//  
func (ameri americaTimeZones) Glace_Bay() string {return "Glace_Bay" }

// Godthab - IANA Time Zone 'Godthab'.
// IANA Source File: Godthab
//  
func (ameri americaTimeZones) Godthab() string {return "Godthab" }

// Goose_Bay - IANA Time Zone 'Goose_Bay'.
// IANA Source File: Goose_Bay
//  
func (ameri americaTimeZones) Goose_Bay() string {return "Goose_Bay" }

// Grand_Turk - IANA Time Zone 'Grand_Turk'.
// IANA Source File: Grand_Turk
//  
func (ameri americaTimeZones) Grand_Turk() string {return "Grand_Turk" }

// Grenada - IANA Time Zone 'Grenada'.
// IANA Source File: Grenada
//  
func (ameri americaTimeZones) Grenada() string {return "Grenada" }

// Guadeloupe - IANA Time Zone 'Guadeloupe'.
// IANA Source File: Guadeloupe
//  
func (ameri americaTimeZones) Guadeloupe() string {return "Guadeloupe" }

// Guatemala - IANA Time Zone 'Guatemala'.
// IANA Source File: Guatemala
//  
func (ameri americaTimeZones) Guatemala() string {return "Guatemala" }

// Guayaquil - IANA Time Zone 'Guayaquil'.
// IANA Source File: Guayaquil
//  
func (ameri americaTimeZones) Guayaquil() string {return "Guayaquil" }

// Guyana - IANA Time Zone 'Guyana'.
// IANA Source File: Guyana
//  
func (ameri americaTimeZones) Guyana() string {return "Guyana" }

// Halifax - IANA Time Zone 'Halifax'.
// IANA Source File: Halifax
//  
func (ameri americaTimeZones) Halifax() string {return "Halifax" }

// Havana - IANA Time Zone 'Havana'.
// IANA Source File: Havana
//  
func (ameri americaTimeZones) Havana() string {return "Havana" }

// Hermosillo - IANA Time Zone 'Hermosillo'.
// IANA Source File: Hermosillo
//  
func (ameri americaTimeZones) Hermosillo() string {return "Hermosillo" }

// Indiana - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (ameri americaTimeZones) Indiana() indianaTimeZones {return "" }

// Indianapolis - IANA Time Zone 'Indianapolis'.
// IANA Source File: Indianapolis
//  
func (ameri americaTimeZones) Indianapolis() string {return "Indianapolis" }

// Inuvik - IANA Time Zone 'Inuvik'.
// IANA Source File: Inuvik
//  
func (ameri americaTimeZones) Inuvik() string {return "Inuvik" }

// Iqaluit - IANA Time Zone 'Iqaluit'.
// IANA Source File: Iqaluit
//  
func (ameri americaTimeZones) Iqaluit() string {return "Iqaluit" }

// Jamaica - IANA Time Zone 'Jamaica'.
// IANA Source File: Jamaica
//  
func (ameri americaTimeZones) Jamaica() string {return "Jamaica" }

// Jujuy - IANA Time Zone 'Jujuy'.
// IANA Source File: Jujuy
//  
func (ameri americaTimeZones) Jujuy() string {return "Jujuy" }

// Juneau - IANA Time Zone 'Juneau'.
// IANA Source File: Juneau
//  
func (ameri americaTimeZones) Juneau() string {return "Juneau" }

// Kentucky - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (ameri americaTimeZones) Kentucky() kentuckyTimeZones {return "" }

// Knox_IN - IANA Time Zone 'Knox_IN'.
// IANA Source File: Knox_IN
//  
func (ameri americaTimeZones) Knox_IN() string {return "Knox_IN" }

// Kralendijk - IANA Time Zone 'Kralendijk'.
// IANA Source File: Kralendijk
//  
func (ameri americaTimeZones) Kralendijk() string {return "Kralendijk" }

// La_Paz - IANA Time Zone 'La_Paz'.
// IANA Source File: La_Paz
//  
func (ameri americaTimeZones) La_Paz() string {return "La_Paz" }

// Lima - IANA Time Zone 'Lima'.
// IANA Source File: Lima
//  
func (ameri americaTimeZones) Lima() string {return "Lima" }

// Los_Angeles - IANA Time Zone 'Los_Angeles'.
// IANA Source File: Los_Angeles
//  
func (ameri americaTimeZones) Los_Angeles() string {return "Los_Angeles" }

// Louisville - IANA Time Zone 'Louisville'.
// IANA Source File: Louisville
//  
func (ameri americaTimeZones) Louisville() string {return "Louisville" }

// Lower_Princes - IANA Time Zone 'Lower_Princes'.
// IANA Source File: Lower_Princes
//  
func (ameri americaTimeZones) Lower_Princes() string {return "Lower_Princes" }

// Maceio - IANA Time Zone 'Maceio'.
// IANA Source File: Maceio
//  
func (ameri americaTimeZones) Maceio() string {return "Maceio" }

// Managua - IANA Time Zone 'Managua'.
// IANA Source File: Managua
//  
func (ameri americaTimeZones) Managua() string {return "Managua" }

// Manaus - IANA Time Zone 'Manaus'.
// IANA Source File: Manaus
//  
func (ameri americaTimeZones) Manaus() string {return "Manaus" }

// Marigot - IANA Time Zone 'Marigot'.
// IANA Source File: Marigot
//  
func (ameri americaTimeZones) Marigot() string {return "Marigot" }

// Martinique - IANA Time Zone 'Martinique'.
// IANA Source File: Martinique
//  
func (ameri americaTimeZones) Martinique() string {return "Martinique" }

// Matamoros - IANA Time Zone 'Matamoros'.
// IANA Source File: Matamoros
//  
func (ameri americaTimeZones) Matamoros() string {return "Matamoros" }

// Mazatlan - IANA Time Zone 'Mazatlan'.
// IANA Source File: Mazatlan
//  
func (ameri americaTimeZones) Mazatlan() string {return "Mazatlan" }

// Mendoza - IANA Time Zone 'Mendoza'.
// IANA Source File: Mendoza
//  
func (ameri americaTimeZones) Mendoza() string {return "Mendoza" }

// Menominee - IANA Time Zone 'Menominee'.
// IANA Source File: Menominee
//  
func (ameri americaTimeZones) Menominee() string {return "Menominee" }

// Merida - IANA Time Zone 'Merida'.
// IANA Source File: Merida
//  
func (ameri americaTimeZones) Merida() string {return "Merida" }

// Metlakatla - IANA Time Zone 'Metlakatla'.
// IANA Source File: Metlakatla
//  
func (ameri americaTimeZones) Metlakatla() string {return "Metlakatla" }

// Mexico_City - IANA Time Zone 'Mexico_City'.
// IANA Source File: Mexico_City
//  
func (ameri americaTimeZones) Mexico_City() string {return "Mexico_City" }

// Miquelon - IANA Time Zone 'Miquelon'.
// IANA Source File: Miquelon
//  
func (ameri americaTimeZones) Miquelon() string {return "Miquelon" }

// Moncton - IANA Time Zone 'Moncton'.
// IANA Source File: Moncton
//  
func (ameri americaTimeZones) Moncton() string {return "Moncton" }

// Monterrey - IANA Time Zone 'Monterrey'.
// IANA Source File: Monterrey
//  
func (ameri americaTimeZones) Monterrey() string {return "Monterrey" }

// Montevideo - IANA Time Zone 'Montevideo'.
// IANA Source File: Montevideo
//  
func (ameri americaTimeZones) Montevideo() string {return "Montevideo" }

// Montreal - IANA Time Zone 'Montreal'.
// IANA Source File: Montreal
//  
func (ameri americaTimeZones) Montreal() string {return "Montreal" }

// Montserrat - IANA Time Zone 'Montserrat'.
// IANA Source File: Montserrat
//  
func (ameri americaTimeZones) Montserrat() string {return "Montserrat" }

// Nassau - IANA Time Zone 'Nassau'.
// IANA Source File: Nassau
//  
func (ameri americaTimeZones) Nassau() string {return "Nassau" }

// New_York - IANA Time Zone 'New_York'.
// IANA Source File: New_York
//  
func (ameri americaTimeZones) New_York() string {return "New_York" }

// Nipigon - IANA Time Zone 'Nipigon'.
// IANA Source File: Nipigon
//  
func (ameri americaTimeZones) Nipigon() string {return "Nipigon" }

// Nome - IANA Time Zone 'Nome'.
// IANA Source File: Nome
//  
func (ameri americaTimeZones) Nome() string {return "Nome" }

// Noronha - IANA Time Zone 'Noronha'.
// IANA Source File: Noronha
//  
func (ameri americaTimeZones) Noronha() string {return "Noronha" }

// North_Dakota - A place holder which defines a sub-group
// of IANA Time Zones.
//  
func (ameri americaTimeZones) North_Dakota() north_DakotaTimeZones {return "" }

// Ojinaga - IANA Time Zone 'Ojinaga'.
// IANA Source File: Ojinaga
//  
func (ameri americaTimeZones) Ojinaga() string {return "Ojinaga" }

// Panama - IANA Time Zone 'Panama'.
// IANA Source File: Panama
//  
func (ameri americaTimeZones) Panama() string {return "Panama" }

// Pangnirtung - IANA Time Zone 'Pangnirtung'.
// IANA Source File: Pangnirtung
//  
func (ameri americaTimeZones) Pangnirtung() string {return "Pangnirtung" }

// Paramaribo - IANA Time Zone 'Paramaribo'.
// IANA Source File: Paramaribo
//  
func (ameri americaTimeZones) Paramaribo() string {return "Paramaribo" }

// Phoenix - IANA Time Zone 'Phoenix'.
// IANA Source File: Phoenix
//  
func (ameri americaTimeZones) Phoenix() string {return "Phoenix" }

// Port-au-Prince - IANA Time Zone 'Port-au-Prince'.
// IANA Source File: Port-au-Prince
//  
func (ameri americaTimeZones) PortMinusauMinusPrince() string {return "Port-au-Prince" }

// Port_of_Spain - IANA Time Zone 'Port_of_Spain'.
// IANA Source File: Port_of_Spain
//  
func (ameri americaTimeZones) Port_of_Spain() string {return "Port_of_Spain" }

// Porto_Acre - IANA Time Zone 'Porto_Acre'.
// IANA Source File: Porto_Acre
//  
func (ameri americaTimeZones) Porto_Acre() string {return "Porto_Acre" }

// Porto_Velho - IANA Time Zone 'Porto_Velho'.
// IANA Source File: Porto_Velho
//  
func (ameri americaTimeZones) Porto_Velho() string {return "Porto_Velho" }

// Puerto_Rico - IANA Time Zone 'Puerto_Rico'.
// IANA Source File: Puerto_Rico
//  
func (ameri americaTimeZones) Puerto_Rico() string {return "Puerto_Rico" }

// Punta_Arenas - IANA Time Zone 'Punta_Arenas'.
// IANA Source File: Punta_Arenas
//  
func (ameri americaTimeZones) Punta_Arenas() string {return "Punta_Arenas" }

// Rainy_River - IANA Time Zone 'Rainy_River'.
// IANA Source File: Rainy_River
//  
func (ameri americaTimeZones) Rainy_River() string {return "Rainy_River" }

// Rankin_Inlet - IANA Time Zone 'Rankin_Inlet'.
// IANA Source File: Rankin_Inlet
//  
func (ameri americaTimeZones) Rankin_Inlet() string {return "Rankin_Inlet" }

// Recife - IANA Time Zone 'Recife'.
// IANA Source File: Recife
//  
func (ameri americaTimeZones) Recife() string {return "Recife" }

// Regina - IANA Time Zone 'Regina'.
// IANA Source File: Regina
//  
func (ameri americaTimeZones) Regina() string {return "Regina" }

// Resolute - IANA Time Zone 'Resolute'.
// IANA Source File: Resolute
//  
func (ameri americaTimeZones) Resolute() string {return "Resolute" }

// Rio_Branco - IANA Time Zone 'Rio_Branco'.
// IANA Source File: Rio_Branco
//  
func (ameri americaTimeZones) Rio_Branco() string {return "Rio_Branco" }

// Rosario - IANA Time Zone 'Rosario'.
// IANA Source File: Rosario
//  
func (ameri americaTimeZones) Rosario() string {return "Rosario" }

// Santa_Isabel - IANA Time Zone 'Santa_Isabel'.
// IANA Source File: Santa_Isabel
//  
func (ameri americaTimeZones) Santa_Isabel() string {return "Santa_Isabel" }

// Santarem - IANA Time Zone 'Santarem'.
// IANA Source File: Santarem
//  
func (ameri americaTimeZones) Santarem() string {return "Santarem" }

// Santiago - IANA Time Zone 'Santiago'.
// IANA Source File: Santiago
//  
func (ameri americaTimeZones) Santiago() string {return "Santiago" }

// Santo_Domingo - IANA Time Zone 'Santo_Domingo'.
// IANA Source File: Santo_Domingo
//  
func (ameri americaTimeZones) Santo_Domingo() string {return "Santo_Domingo" }

// Sao_Paulo - IANA Time Zone 'Sao_Paulo'.
// IANA Source File: Sao_Paulo
//  
func (ameri americaTimeZones) Sao_Paulo() string {return "Sao_Paulo" }

// Scoresbysund - IANA Time Zone 'Scoresbysund'.
// IANA Source File: Scoresbysund
//  
func (ameri americaTimeZones) Scoresbysund() string {return "Scoresbysund" }

// Shiprock - IANA Time Zone 'Shiprock'.
// IANA Source File: Shiprock
//  
func (ameri americaTimeZones) Shiprock() string {return "Shiprock" }

// Sitka - IANA Time Zone 'Sitka'.
// IANA Source File: Sitka
//  
func (ameri americaTimeZones) Sitka() string {return "Sitka" }

// St_Barthelemy - IANA Time Zone 'St_Barthelemy'.
// IANA Source File: St_Barthelemy
//  
func (ameri americaTimeZones) St_Barthelemy() string {return "St_Barthelemy" }

// St_Johns - IANA Time Zone 'St_Johns'.
// IANA Source File: St_Johns
//  
func (ameri americaTimeZones) St_Johns() string {return "St_Johns" }

// St_Kitts - IANA Time Zone 'St_Kitts'.
// IANA Source File: St_Kitts
//  
func (ameri americaTimeZones) St_Kitts() string {return "St_Kitts" }

// St_Lucia - IANA Time Zone 'St_Lucia'.
// IANA Source File: St_Lucia
//  
func (ameri americaTimeZones) St_Lucia() string {return "St_Lucia" }

// St_Thomas - IANA Time Zone 'St_Thomas'.
// IANA Source File: St_Thomas
//  
func (ameri americaTimeZones) St_Thomas() string {return "St_Thomas" }

// St_Vincent - IANA Time Zone 'St_Vincent'.
// IANA Source File: St_Vincent
//  
func (ameri americaTimeZones) St_Vincent() string {return "St_Vincent" }

// Swift_Current - IANA Time Zone 'Swift_Current'.
// IANA Source File: Swift_Current
//  
func (ameri americaTimeZones) Swift_Current() string {return "Swift_Current" }

// Tegucigalpa - IANA Time Zone 'Tegucigalpa'.
// IANA Source File: Tegucigalpa
//  
func (ameri americaTimeZones) Tegucigalpa() string {return "Tegucigalpa" }

// Thule - IANA Time Zone 'Thule'.
// IANA Source File: Thule
//  
func (ameri americaTimeZones) Thule() string {return "Thule" }

// Thunder_Bay - IANA Time Zone 'Thunder_Bay'.
// IANA Source File: Thunder_Bay
//  
func (ameri americaTimeZones) Thunder_Bay() string {return "Thunder_Bay" }

// Tijuana - IANA Time Zone 'Tijuana'.
// IANA Source File: Tijuana
//  
func (ameri americaTimeZones) Tijuana() string {return "Tijuana" }

// Toronto - IANA Time Zone 'Toronto'.
// IANA Source File: Toronto
//  
func (ameri americaTimeZones) Toronto() string {return "Toronto" }

// Tortola - IANA Time Zone 'Tortola'.
// IANA Source File: Tortola
//  
func (ameri americaTimeZones) Tortola() string {return "Tortola" }

// Vancouver - IANA Time Zone 'Vancouver'.
// IANA Source File: Vancouver
//  
func (ameri americaTimeZones) Vancouver() string {return "Vancouver" }

// Virgin - IANA Time Zone 'Virgin'.
// IANA Source File: Virgin
//  
func (ameri americaTimeZones) Virgin() string {return "Virgin" }

// Whitehorse - IANA Time Zone 'Whitehorse'.
// IANA Source File: Whitehorse
//  
func (ameri americaTimeZones) Whitehorse() string {return "Whitehorse" }

// Winnipeg - IANA Time Zone 'Winnipeg'.
// IANA Source File: Winnipeg
//  
func (ameri americaTimeZones) Winnipeg() string {return "Winnipeg" }

// Yakutat - IANA Time Zone 'Yakutat'.
// IANA Source File: Yakutat
//  
func (ameri americaTimeZones) Yakutat() string {return "Yakutat" }

// Yellowknife - IANA Time Zone 'Yellowknife'.
// IANA Source File: Yellowknife
//  
func (ameri americaTimeZones) Yellowknife() string {return "Yellowknife" }


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

// Casey - IANA Time Zone 'Casey'.
// IANA Source File: Casey
//  
func (antar antarcticaTimeZones) Casey() string {return "Casey" }

// Davis - IANA Time Zone 'Davis'.
// IANA Source File: Davis
//  
func (antar antarcticaTimeZones) Davis() string {return "Davis" }

// DumontDUrville - IANA Time Zone 'DumontDUrville'.
// IANA Source File: DumontDUrville
//  
func (antar antarcticaTimeZones) DumontDUrville() string {return "DumontDUrville" }

// Macquarie - IANA Time Zone 'Macquarie'.
// IANA Source File: Macquarie
//  
func (antar antarcticaTimeZones) Macquarie() string {return "Macquarie" }

// Mawson - IANA Time Zone 'Mawson'.
// IANA Source File: Mawson
//  
func (antar antarcticaTimeZones) Mawson() string {return "Mawson" }

// McMurdo - IANA Time Zone 'McMurdo'.
// IANA Source File: McMurdo
//  
func (antar antarcticaTimeZones) McMurdo() string {return "McMurdo" }

// Palmer - IANA Time Zone 'Palmer'.
// IANA Source File: Palmer
//  
func (antar antarcticaTimeZones) Palmer() string {return "Palmer" }

// Rothera - IANA Time Zone 'Rothera'.
// IANA Source File: Rothera
//  
func (antar antarcticaTimeZones) Rothera() string {return "Rothera" }

// South_Pole - IANA Time Zone 'South_Pole'.
// IANA Source File: South_Pole
//  
func (antar antarcticaTimeZones) South_Pole() string {return "South_Pole" }

// Syowa - IANA Time Zone 'Syowa'.
// IANA Source File: Syowa
//  
func (antar antarcticaTimeZones) Syowa() string {return "Syowa" }

// Troll - IANA Time Zone 'Troll'.
// IANA Source File: Troll
//  
func (antar antarcticaTimeZones) Troll() string {return "Troll" }

// Vostok - IANA Time Zone 'Vostok'.
// IANA Source File: Vostok
//  
func (antar antarcticaTimeZones) Vostok() string {return "Vostok" }


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

// Longyearbyen - IANA Time Zone 'Longyearbyen'.
// IANA Source File: Longyearbyen
//  
func (arcti arcticTimeZones) Longyearbyen() string {return "Longyearbyen" }


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

// Aden - IANA Time Zone 'Aden'.
// IANA Source File: Aden
//  
func (asiaT asiaTimeZones) Aden() string {return "Aden" }

// Almaty - IANA Time Zone 'Almaty'.
// IANA Source File: Almaty
//  
func (asiaT asiaTimeZones) Almaty() string {return "Almaty" }

// Amman - IANA Time Zone 'Amman'.
// IANA Source File: Amman
//  
func (asiaT asiaTimeZones) Amman() string {return "Amman" }

// Anadyr - IANA Time Zone 'Anadyr'.
// IANA Source File: Anadyr
//  
func (asiaT asiaTimeZones) Anadyr() string {return "Anadyr" }

// Aqtau - IANA Time Zone 'Aqtau'.
// IANA Source File: Aqtau
//  
func (asiaT asiaTimeZones) Aqtau() string {return "Aqtau" }

// Aqtobe - IANA Time Zone 'Aqtobe'.
// IANA Source File: Aqtobe
//  
func (asiaT asiaTimeZones) Aqtobe() string {return "Aqtobe" }

// Ashgabat - IANA Time Zone 'Ashgabat'.
// IANA Source File: Ashgabat
//  
func (asiaT asiaTimeZones) Ashgabat() string {return "Ashgabat" }

// Ashkhabad - IANA Time Zone 'Ashkhabad'.
// IANA Source File: Ashkhabad
//  
func (asiaT asiaTimeZones) Ashkhabad() string {return "Ashkhabad" }

// Atyrau - IANA Time Zone 'Atyrau'.
// IANA Source File: Atyrau
//  
func (asiaT asiaTimeZones) Atyrau() string {return "Atyrau" }

// Baghdad - IANA Time Zone 'Baghdad'.
// IANA Source File: Baghdad
//  
func (asiaT asiaTimeZones) Baghdad() string {return "Baghdad" }

// Bahrain - IANA Time Zone 'Bahrain'.
// IANA Source File: Bahrain
//  
func (asiaT asiaTimeZones) Bahrain() string {return "Bahrain" }

// Baku - IANA Time Zone 'Baku'.
// IANA Source File: Baku
//  
func (asiaT asiaTimeZones) Baku() string {return "Baku" }

// Bangkok - IANA Time Zone 'Bangkok'.
// IANA Source File: Bangkok
//  
func (asiaT asiaTimeZones) Bangkok() string {return "Bangkok" }

// Barnaul - IANA Time Zone 'Barnaul'.
// IANA Source File: Barnaul
//  
func (asiaT asiaTimeZones) Barnaul() string {return "Barnaul" }

// Beirut - IANA Time Zone 'Beirut'.
// IANA Source File: Beirut
//  
func (asiaT asiaTimeZones) Beirut() string {return "Beirut" }

// Bishkek - IANA Time Zone 'Bishkek'.
// IANA Source File: Bishkek
//  
func (asiaT asiaTimeZones) Bishkek() string {return "Bishkek" }

// Brunei - IANA Time Zone 'Brunei'.
// IANA Source File: Brunei
//  
func (asiaT asiaTimeZones) Brunei() string {return "Brunei" }

// Calcutta - IANA Time Zone 'Calcutta'.
// IANA Source File: Calcutta
//  
func (asiaT asiaTimeZones) Calcutta() string {return "Calcutta" }

// Chita - IANA Time Zone 'Chita'.
// IANA Source File: Chita
//  
func (asiaT asiaTimeZones) Chita() string {return "Chita" }

// Choibalsan - IANA Time Zone 'Choibalsan'.
// IANA Source File: Choibalsan
//  
func (asiaT asiaTimeZones) Choibalsan() string {return "Choibalsan" }

// Chongqing - IANA Time Zone 'Chongqing'.
// IANA Source File: Chongqing
//  
func (asiaT asiaTimeZones) Chongqing() string {return "Chongqing" }

// Chungking - IANA Time Zone 'Chungking'.
// IANA Source File: Chungking
//  
func (asiaT asiaTimeZones) Chungking() string {return "Chungking" }

// Colombo - IANA Time Zone 'Colombo'.
// IANA Source File: Colombo
//  
func (asiaT asiaTimeZones) Colombo() string {return "Colombo" }

// Dacca - IANA Time Zone 'Dacca'.
// IANA Source File: Dacca
//  
func (asiaT asiaTimeZones) Dacca() string {return "Dacca" }

// Damascus - IANA Time Zone 'Damascus'.
// IANA Source File: Damascus
//  
func (asiaT asiaTimeZones) Damascus() string {return "Damascus" }

// Dhaka - IANA Time Zone 'Dhaka'.
// IANA Source File: Dhaka
//  
func (asiaT asiaTimeZones) Dhaka() string {return "Dhaka" }

// Dili - IANA Time Zone 'Dili'.
// IANA Source File: Dili
//  
func (asiaT asiaTimeZones) Dili() string {return "Dili" }

// Dubai - IANA Time Zone 'Dubai'.
// IANA Source File: Dubai
//  
func (asiaT asiaTimeZones) Dubai() string {return "Dubai" }

// Dushanbe - IANA Time Zone 'Dushanbe'.
// IANA Source File: Dushanbe
//  
func (asiaT asiaTimeZones) Dushanbe() string {return "Dushanbe" }

// Famagusta - IANA Time Zone 'Famagusta'.
// IANA Source File: Famagusta
//  
func (asiaT asiaTimeZones) Famagusta() string {return "Famagusta" }

// Gaza - IANA Time Zone 'Gaza'.
// IANA Source File: Gaza
//  
func (asiaT asiaTimeZones) Gaza() string {return "Gaza" }

// Harbin - IANA Time Zone 'Harbin'.
// IANA Source File: Harbin
//  
func (asiaT asiaTimeZones) Harbin() string {return "Harbin" }

// Hebron - IANA Time Zone 'Hebron'.
// IANA Source File: Hebron
//  
func (asiaT asiaTimeZones) Hebron() string {return "Hebron" }

// Ho_Chi_Minh - IANA Time Zone 'Ho_Chi_Minh'.
// IANA Source File: Ho_Chi_Minh
//  
func (asiaT asiaTimeZones) Ho_Chi_Minh() string {return "Ho_Chi_Minh" }

// Hong_Kong - IANA Time Zone 'Hong_Kong'.
// IANA Source File: Hong_Kong
//  
func (asiaT asiaTimeZones) Hong_Kong() string {return "Hong_Kong" }

// Hovd - IANA Time Zone 'Hovd'.
// IANA Source File: Hovd
//  
func (asiaT asiaTimeZones) Hovd() string {return "Hovd" }

// Irkutsk - IANA Time Zone 'Irkutsk'.
// IANA Source File: Irkutsk
//  
func (asiaT asiaTimeZones) Irkutsk() string {return "Irkutsk" }

// Istanbul - IANA Time Zone 'Istanbul'.
// IANA Source File: Istanbul
//  
func (asiaT asiaTimeZones) Istanbul() string {return "Istanbul" }

// Jakarta - IANA Time Zone 'Jakarta'.
// IANA Source File: Jakarta
//  
func (asiaT asiaTimeZones) Jakarta() string {return "Jakarta" }

// Jayapura - IANA Time Zone 'Jayapura'.
// IANA Source File: Jayapura
//  
func (asiaT asiaTimeZones) Jayapura() string {return "Jayapura" }

// Jerusalem - IANA Time Zone 'Jerusalem'.
// IANA Source File: Jerusalem
//  
func (asiaT asiaTimeZones) Jerusalem() string {return "Jerusalem" }

// Kabul - IANA Time Zone 'Kabul'.
// IANA Source File: Kabul
//  
func (asiaT asiaTimeZones) Kabul() string {return "Kabul" }

// Kamchatka - IANA Time Zone 'Kamchatka'.
// IANA Source File: Kamchatka
//  
func (asiaT asiaTimeZones) Kamchatka() string {return "Kamchatka" }

// Karachi - IANA Time Zone 'Karachi'.
// IANA Source File: Karachi
//  
func (asiaT asiaTimeZones) Karachi() string {return "Karachi" }

// Kashgar - IANA Time Zone 'Kashgar'.
// IANA Source File: Kashgar
//  
func (asiaT asiaTimeZones) Kashgar() string {return "Kashgar" }

// Kathmandu - IANA Time Zone 'Kathmandu'.
// IANA Source File: Kathmandu
//  
func (asiaT asiaTimeZones) Kathmandu() string {return "Kathmandu" }

// Katmandu - IANA Time Zone 'Katmandu'.
// IANA Source File: Katmandu
//  
func (asiaT asiaTimeZones) Katmandu() string {return "Katmandu" }

// Khandyga - IANA Time Zone 'Khandyga'.
// IANA Source File: Khandyga
//  
func (asiaT asiaTimeZones) Khandyga() string {return "Khandyga" }

// Kolkata - IANA Time Zone 'Kolkata'.
// IANA Source File: Kolkata
//  
func (asiaT asiaTimeZones) Kolkata() string {return "Kolkata" }

// Krasnoyarsk - IANA Time Zone 'Krasnoyarsk'.
// IANA Source File: Krasnoyarsk
//  
func (asiaT asiaTimeZones) Krasnoyarsk() string {return "Krasnoyarsk" }

// Kuala_Lumpur - IANA Time Zone 'Kuala_Lumpur'.
// IANA Source File: Kuala_Lumpur
//  
func (asiaT asiaTimeZones) Kuala_Lumpur() string {return "Kuala_Lumpur" }

// Kuching - IANA Time Zone 'Kuching'.
// IANA Source File: Kuching
//  
func (asiaT asiaTimeZones) Kuching() string {return "Kuching" }

// Kuwait - IANA Time Zone 'Kuwait'.
// IANA Source File: Kuwait
//  
func (asiaT asiaTimeZones) Kuwait() string {return "Kuwait" }

// Macao - IANA Time Zone 'Macao'.
// IANA Source File: Macao
//  
func (asiaT asiaTimeZones) Macao() string {return "Macao" }

// Macau - IANA Time Zone 'Macau'.
// IANA Source File: Macau
//  
func (asiaT asiaTimeZones) Macau() string {return "Macau" }

// Magadan - IANA Time Zone 'Magadan'.
// IANA Source File: Magadan
//  
func (asiaT asiaTimeZones) Magadan() string {return "Magadan" }

// Makassar - IANA Time Zone 'Makassar'.
// IANA Source File: Makassar
//  
func (asiaT asiaTimeZones) Makassar() string {return "Makassar" }

// Manila - IANA Time Zone 'Manila'.
// IANA Source File: Manila
//  
func (asiaT asiaTimeZones) Manila() string {return "Manila" }

// Muscat - IANA Time Zone 'Muscat'.
// IANA Source File: Muscat
//  
func (asiaT asiaTimeZones) Muscat() string {return "Muscat" }

// Nicosia - IANA Time Zone 'Nicosia'.
// IANA Source File: Nicosia
//  
func (asiaT asiaTimeZones) Nicosia() string {return "Nicosia" }

// Novokuznetsk - IANA Time Zone 'Novokuznetsk'.
// IANA Source File: Novokuznetsk
//  
func (asiaT asiaTimeZones) Novokuznetsk() string {return "Novokuznetsk" }

// Novosibirsk - IANA Time Zone 'Novosibirsk'.
// IANA Source File: Novosibirsk
//  
func (asiaT asiaTimeZones) Novosibirsk() string {return "Novosibirsk" }

// Omsk - IANA Time Zone 'Omsk'.
// IANA Source File: Omsk
//  
func (asiaT asiaTimeZones) Omsk() string {return "Omsk" }

// Oral - IANA Time Zone 'Oral'.
// IANA Source File: Oral
//  
func (asiaT asiaTimeZones) Oral() string {return "Oral" }

// Phnom_Penh - IANA Time Zone 'Phnom_Penh'.
// IANA Source File: Phnom_Penh
//  
func (asiaT asiaTimeZones) Phnom_Penh() string {return "Phnom_Penh" }

// Pontianak - IANA Time Zone 'Pontianak'.
// IANA Source File: Pontianak
//  
func (asiaT asiaTimeZones) Pontianak() string {return "Pontianak" }

// Pyongyang - IANA Time Zone 'Pyongyang'.
// IANA Source File: Pyongyang
//  
func (asiaT asiaTimeZones) Pyongyang() string {return "Pyongyang" }

// Qatar - IANA Time Zone 'Qatar'.
// IANA Source File: Qatar
//  
func (asiaT asiaTimeZones) Qatar() string {return "Qatar" }

// Qostanay - IANA Time Zone 'Qostanay'.
// IANA Source File: Qostanay
//  
func (asiaT asiaTimeZones) Qostanay() string {return "Qostanay" }

// Qyzylorda - IANA Time Zone 'Qyzylorda'.
// IANA Source File: Qyzylorda
//  
func (asiaT asiaTimeZones) Qyzylorda() string {return "Qyzylorda" }

// Rangoon - IANA Time Zone 'Rangoon'.
// IANA Source File: Rangoon
//  
func (asiaT asiaTimeZones) Rangoon() string {return "Rangoon" }

// Riyadh - IANA Time Zone 'Riyadh'.
// IANA Source File: Riyadh
//  
func (asiaT asiaTimeZones) Riyadh() string {return "Riyadh" }

// Saigon - IANA Time Zone 'Saigon'.
// IANA Source File: Saigon
//  
func (asiaT asiaTimeZones) Saigon() string {return "Saigon" }

// Sakhalin - IANA Time Zone 'Sakhalin'.
// IANA Source File: Sakhalin
//  
func (asiaT asiaTimeZones) Sakhalin() string {return "Sakhalin" }

// Samarkand - IANA Time Zone 'Samarkand'.
// IANA Source File: Samarkand
//  
func (asiaT asiaTimeZones) Samarkand() string {return "Samarkand" }

// Seoul - IANA Time Zone 'Seoul'.
// IANA Source File: Seoul
//  
func (asiaT asiaTimeZones) Seoul() string {return "Seoul" }

// Shanghai - IANA Time Zone 'Shanghai'.
// IANA Source File: Shanghai
//  
func (asiaT asiaTimeZones) Shanghai() string {return "Shanghai" }

// Singapore - IANA Time Zone 'Singapore'.
// IANA Source File: Singapore
//  
func (asiaT asiaTimeZones) Singapore() string {return "Singapore" }

// Srednekolymsk - IANA Time Zone 'Srednekolymsk'.
// IANA Source File: Srednekolymsk
//  
func (asiaT asiaTimeZones) Srednekolymsk() string {return "Srednekolymsk" }

// Taipei - IANA Time Zone 'Taipei'.
// IANA Source File: Taipei
//  
func (asiaT asiaTimeZones) Taipei() string {return "Taipei" }

// Tashkent - IANA Time Zone 'Tashkent'.
// IANA Source File: Tashkent
//  
func (asiaT asiaTimeZones) Tashkent() string {return "Tashkent" }

// Tbilisi - IANA Time Zone 'Tbilisi'.
// IANA Source File: Tbilisi
//  
func (asiaT asiaTimeZones) Tbilisi() string {return "Tbilisi" }

// Tehran - IANA Time Zone 'Tehran'.
// IANA Source File: Tehran
//  
func (asiaT asiaTimeZones) Tehran() string {return "Tehran" }

// Tel_Aviv - IANA Time Zone 'Tel_Aviv'.
// IANA Source File: Tel_Aviv
//  
func (asiaT asiaTimeZones) Tel_Aviv() string {return "Tel_Aviv" }

// Thimbu - IANA Time Zone 'Thimbu'.
// IANA Source File: Thimbu
//  
func (asiaT asiaTimeZones) Thimbu() string {return "Thimbu" }

// Thimphu - IANA Time Zone 'Thimphu'.
// IANA Source File: Thimphu
//  
func (asiaT asiaTimeZones) Thimphu() string {return "Thimphu" }

// Tokyo - IANA Time Zone 'Tokyo'.
// IANA Source File: Tokyo
//  
func (asiaT asiaTimeZones) Tokyo() string {return "Tokyo" }

// Tomsk - IANA Time Zone 'Tomsk'.
// IANA Source File: Tomsk
//  
func (asiaT asiaTimeZones) Tomsk() string {return "Tomsk" }

// Ujung_Pandang - IANA Time Zone 'Ujung_Pandang'.
// IANA Source File: Ujung_Pandang
//  
func (asiaT asiaTimeZones) Ujung_Pandang() string {return "Ujung_Pandang" }

// Ulaanbaatar - IANA Time Zone 'Ulaanbaatar'.
// IANA Source File: Ulaanbaatar
//  
func (asiaT asiaTimeZones) Ulaanbaatar() string {return "Ulaanbaatar" }

// Ulan_Bator - IANA Time Zone 'Ulan_Bator'.
// IANA Source File: Ulan_Bator
//  
func (asiaT asiaTimeZones) Ulan_Bator() string {return "Ulan_Bator" }

// Urumqi - IANA Time Zone 'Urumqi'.
// IANA Source File: Urumqi
//  
func (asiaT asiaTimeZones) Urumqi() string {return "Urumqi" }

// Ust-Nera - IANA Time Zone 'Ust-Nera'.
// IANA Source File: Ust-Nera
//  
func (asiaT asiaTimeZones) UstMinusNera() string {return "Ust-Nera" }

// Vientiane - IANA Time Zone 'Vientiane'.
// IANA Source File: Vientiane
//  
func (asiaT asiaTimeZones) Vientiane() string {return "Vientiane" }

// Vladivostok - IANA Time Zone 'Vladivostok'.
// IANA Source File: Vladivostok
//  
func (asiaT asiaTimeZones) Vladivostok() string {return "Vladivostok" }

// Yakutsk - IANA Time Zone 'Yakutsk'.
// IANA Source File: Yakutsk
//  
func (asiaT asiaTimeZones) Yakutsk() string {return "Yakutsk" }

// Yangon - IANA Time Zone 'Yangon'.
// IANA Source File: Yangon
//  
func (asiaT asiaTimeZones) Yangon() string {return "Yangon" }

// Yekaterinburg - IANA Time Zone 'Yekaterinburg'.
// IANA Source File: Yekaterinburg
//  
func (asiaT asiaTimeZones) Yekaterinburg() string {return "Yekaterinburg" }

// Yerevan - IANA Time Zone 'Yerevan'.
// IANA Source File: Yerevan
//  
func (asiaT asiaTimeZones) Yerevan() string {return "Yerevan" }


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

// Azores - IANA Time Zone 'Azores'.
// IANA Source File: Azores
//  
func (atlan atlanticTimeZones) Azores() string {return "Azores" }

// Bermuda - IANA Time Zone 'Bermuda'.
// IANA Source File: Bermuda
//  
func (atlan atlanticTimeZones) Bermuda() string {return "Bermuda" }

// Canary - IANA Time Zone 'Canary'.
// IANA Source File: Canary
//  
func (atlan atlanticTimeZones) Canary() string {return "Canary" }

// Cape_Verde - IANA Time Zone 'Cape_Verde'.
// IANA Source File: Cape_Verde
//  
func (atlan atlanticTimeZones) Cape_Verde() string {return "Cape_Verde" }

// Faeroe - IANA Time Zone 'Faeroe'.
// IANA Source File: Faeroe
//  
func (atlan atlanticTimeZones) Faeroe() string {return "Faeroe" }

// Faroe - IANA Time Zone 'Faroe'.
// IANA Source File: Faroe
//  
func (atlan atlanticTimeZones) Faroe() string {return "Faroe" }

// Jan_Mayen - IANA Time Zone 'Jan_Mayen'.
// IANA Source File: Jan_Mayen
//  
func (atlan atlanticTimeZones) Jan_Mayen() string {return "Jan_Mayen" }

// Madeira - IANA Time Zone 'Madeira'.
// IANA Source File: Madeira
//  
func (atlan atlanticTimeZones) Madeira() string {return "Madeira" }

// Reykjavik - IANA Time Zone 'Reykjavik'.
// IANA Source File: Reykjavik
//  
func (atlan atlanticTimeZones) Reykjavik() string {return "Reykjavik" }

// South_Georgia - IANA Time Zone 'South_Georgia'.
// IANA Source File: South_Georgia
//  
func (atlan atlanticTimeZones) South_Georgia() string {return "South_Georgia" }

// St_Helena - IANA Time Zone 'St_Helena'.
// IANA Source File: St_Helena
//  
func (atlan atlanticTimeZones) St_Helena() string {return "St_Helena" }

// Stanley - IANA Time Zone 'Stanley'.
// IANA Source File: Stanley
//  
func (atlan atlanticTimeZones) Stanley() string {return "Stanley" }


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

// ACT - IANA Time Zone 'ACT'.
// IANA Source File: ACT
//  
func (austr australiaTimeZones) ACT() string {return "ACT" }

// Adelaide - IANA Time Zone 'Adelaide'.
// IANA Source File: Adelaide
//  
func (austr australiaTimeZones) Adelaide() string {return "Adelaide" }

// Brisbane - IANA Time Zone 'Brisbane'.
// IANA Source File: Brisbane
//  
func (austr australiaTimeZones) Brisbane() string {return "Brisbane" }

// Broken_Hill - IANA Time Zone 'Broken_Hill'.
// IANA Source File: Broken_Hill
//  
func (austr australiaTimeZones) Broken_Hill() string {return "Broken_Hill" }

// Canberra - IANA Time Zone 'Canberra'.
// IANA Source File: Canberra
//  
func (austr australiaTimeZones) Canberra() string {return "Canberra" }

// Currie - IANA Time Zone 'Currie'.
// IANA Source File: Currie
//  
func (austr australiaTimeZones) Currie() string {return "Currie" }

// Darwin - IANA Time Zone 'Darwin'.
// IANA Source File: Darwin
//  
func (austr australiaTimeZones) Darwin() string {return "Darwin" }

// Eucla - IANA Time Zone 'Eucla'.
// IANA Source File: Eucla
//  
func (austr australiaTimeZones) Eucla() string {return "Eucla" }

// Hobart - IANA Time Zone 'Hobart'.
// IANA Source File: Hobart
//  
func (austr australiaTimeZones) Hobart() string {return "Hobart" }

// LHI - IANA Time Zone 'LHI'.
// IANA Source File: LHI
//  
func (austr australiaTimeZones) LHI() string {return "LHI" }

// Lindeman - IANA Time Zone 'Lindeman'.
// IANA Source File: Lindeman
//  
func (austr australiaTimeZones) Lindeman() string {return "Lindeman" }

// Lord_Howe - IANA Time Zone 'Lord_Howe'.
// IANA Source File: Lord_Howe
//  
func (austr australiaTimeZones) Lord_Howe() string {return "Lord_Howe" }

// Melbourne - IANA Time Zone 'Melbourne'.
// IANA Source File: Melbourne
//  
func (austr australiaTimeZones) Melbourne() string {return "Melbourne" }

// North - IANA Time Zone 'North'.
// IANA Source File: North
//  
func (austr australiaTimeZones) North() string {return "North" }

// NSW - IANA Time Zone 'NSW'.
// IANA Source File: NSW
//  
func (austr australiaTimeZones) NSW() string {return "NSW" }

// Perth - IANA Time Zone 'Perth'.
// IANA Source File: Perth
//  
func (austr australiaTimeZones) Perth() string {return "Perth" }

// Queensland - IANA Time Zone 'Queensland'.
// IANA Source File: Queensland
//  
func (austr australiaTimeZones) Queensland() string {return "Queensland" }

// South - IANA Time Zone 'South'.
// IANA Source File: South
//  
func (austr australiaTimeZones) South() string {return "South" }

// Sydney - IANA Time Zone 'Sydney'.
// IANA Source File: Sydney
//  
func (austr australiaTimeZones) Sydney() string {return "Sydney" }

// Tasmania - IANA Time Zone 'Tasmania'.
// IANA Source File: Tasmania
//  
func (austr australiaTimeZones) Tasmania() string {return "Tasmania" }

// Victoria - IANA Time Zone 'Victoria'.
// IANA Source File: Victoria
//  
func (austr australiaTimeZones) Victoria() string {return "Victoria" }

// West - IANA Time Zone 'West'.
// IANA Source File: West
//  
func (austr australiaTimeZones) West() string {return "West" }

// Yancowinna - IANA Time Zone 'Yancowinna'.
// IANA Source File: Yancowinna
//  
func (austr australiaTimeZones) Yancowinna() string {return "Yancowinna" }


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

// Acre - IANA Time Zone 'Acre'.
// IANA Source File: Acre
//  
func (brazi brazilTimeZones) Acre() string {return "Acre" }

// DeNoronha - IANA Time Zone 'DeNoronha'.
// IANA Source File: DeNoronha
//  
func (brazi brazilTimeZones) DeNoronha() string {return "DeNoronha" }

// East - IANA Time Zone 'East'.
// IANA Source File: East
//  
func (brazi brazilTimeZones) East() string {return "East" }

// West - IANA Time Zone 'West'.
// IANA Source File: West
//  
func (brazi brazilTimeZones) West() string {return "West" }


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

// Atlantic - IANA Time Zone 'Atlantic'.
// IANA Source File: Atlantic
//  
func (canad canadaTimeZones) Atlantic() string {return "Atlantic" }

// Central - IANA Time Zone 'Central'.
// IANA Source File: Central
//  
func (canad canadaTimeZones) Central() string {return "Central" }

// Eastern - IANA Time Zone 'Eastern'.
// IANA Source File: Eastern
//  
func (canad canadaTimeZones) Eastern() string {return "Eastern" }

// Mountain - IANA Time Zone 'Mountain'.
// IANA Source File: Mountain
//  
func (canad canadaTimeZones) Mountain() string {return "Mountain" }

// Newfoundland - IANA Time Zone 'Newfoundland'.
// IANA Source File: Newfoundland
//  
func (canad canadaTimeZones) Newfoundland() string {return "Newfoundland" }

// Pacific - IANA Time Zone 'Pacific'.
// IANA Source File: Pacific
//  
func (canad canadaTimeZones) Pacific() string {return "Pacific" }

// Saskatchewan - IANA Time Zone 'Saskatchewan'.
// IANA Source File: Saskatchewan
//  
func (canad canadaTimeZones) Saskatchewan() string {return "Saskatchewan" }

// Yukon - IANA Time Zone 'Yukon'.
// IANA Source File: Yukon
//  
func (canad canadaTimeZones) Yukon() string {return "Yukon" }


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

// Continental - IANA Time Zone 'Continental'.
// IANA Source File: Continental
//  
func (chile chileTimeZones) Continental() string {return "Continental" }

// EasterIsland - IANA Time Zone 'EasterIsland'.
// IANA Source File: EasterIsland
//  
func (chile chileTimeZones) EasterIsland() string {return "EasterIsland" }


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

// GMT - IANA Time Zone 'GMT'.
// IANA Source File: GMT
//  
func (etcTi etcTimeZones) GMT() string {return "GMT" }

// GMT+0 - IANA Time Zone 'GMT+0'.
// IANA Source File: GMT+0
//  
func (etcTi etcTimeZones) GMTPlus00() string {return "GMT+0" }

// GMT+1 - IANA Time Zone 'GMT+1'.
// IANA Source File: GMT+1
//  
func (etcTi etcTimeZones) GMTPlus01() string {return "GMT+1" }

// GMT+2 - IANA Time Zone 'GMT+2'.
// IANA Source File: GMT+2
//  
func (etcTi etcTimeZones) GMTPlus02() string {return "GMT+2" }

// GMT+3 - IANA Time Zone 'GMT+3'.
// IANA Source File: GMT+3
//  
func (etcTi etcTimeZones) GMTPlus03() string {return "GMT+3" }

// GMT+4 - IANA Time Zone 'GMT+4'.
// IANA Source File: GMT+4
//  
func (etcTi etcTimeZones) GMTPlus04() string {return "GMT+4" }

// GMT+5 - IANA Time Zone 'GMT+5'.
// IANA Source File: GMT+5
//  
func (etcTi etcTimeZones) GMTPlus05() string {return "GMT+5" }

// GMT+6 - IANA Time Zone 'GMT+6'.
// IANA Source File: GMT+6
//  
func (etcTi etcTimeZones) GMTPlus06() string {return "GMT+6" }

// GMT+7 - IANA Time Zone 'GMT+7'.
// IANA Source File: GMT+7
//  
func (etcTi etcTimeZones) GMTPlus07() string {return "GMT+7" }

// GMT+8 - IANA Time Zone 'GMT+8'.
// IANA Source File: GMT+8
//  
func (etcTi etcTimeZones) GMTPlus08() string {return "GMT+8" }

// GMT+9 - IANA Time Zone 'GMT+9'.
// IANA Source File: GMT+9
//  
func (etcTi etcTimeZones) GMTPlus09() string {return "GMT+9" }

// GMT+10 - IANA Time Zone 'GMT+10'.
// IANA Source File: GMT+10
//  
func (etcTi etcTimeZones) GMTPlus10() string {return "GMT+10" }

// GMT+11 - IANA Time Zone 'GMT+11'.
// IANA Source File: GMT+11
//  
func (etcTi etcTimeZones) GMTPlus11() string {return "GMT+11" }

// GMT+12 - IANA Time Zone 'GMT+12'.
// IANA Source File: GMT+12
//  
func (etcTi etcTimeZones) GMTPlus12() string {return "GMT+12" }

// GMT-0 - IANA Time Zone 'GMT-0'.
// IANA Source File: GMT-0
//  
func (etcTi etcTimeZones) GMTMinus00() string {return "GMT-0" }

// GMT-1 - IANA Time Zone 'GMT-1'.
// IANA Source File: GMT-1
//  
func (etcTi etcTimeZones) GMTMinus01() string {return "GMT-1" }

// GMT-2 - IANA Time Zone 'GMT-2'.
// IANA Source File: GMT-2
//  
func (etcTi etcTimeZones) GMTMinus02() string {return "GMT-2" }

// GMT-3 - IANA Time Zone 'GMT-3'.
// IANA Source File: GMT-3
//  
func (etcTi etcTimeZones) GMTMinus03() string {return "GMT-3" }

// GMT-4 - IANA Time Zone 'GMT-4'.
// IANA Source File: GMT-4
//  
func (etcTi etcTimeZones) GMTMinus04() string {return "GMT-4" }

// GMT-5 - IANA Time Zone 'GMT-5'.
// IANA Source File: GMT-5
//  
func (etcTi etcTimeZones) GMTMinus05() string {return "GMT-5" }

// GMT-6 - IANA Time Zone 'GMT-6'.
// IANA Source File: GMT-6
//  
func (etcTi etcTimeZones) GMTMinus06() string {return "GMT-6" }

// GMT-7 - IANA Time Zone 'GMT-7'.
// IANA Source File: GMT-7
//  
func (etcTi etcTimeZones) GMTMinus07() string {return "GMT-7" }

// GMT-8 - IANA Time Zone 'GMT-8'.
// IANA Source File: GMT-8
//  
func (etcTi etcTimeZones) GMTMinus08() string {return "GMT-8" }

// GMT-9 - IANA Time Zone 'GMT-9'.
// IANA Source File: GMT-9
//  
func (etcTi etcTimeZones) GMTMinus09() string {return "GMT-9" }

// GMT-10 - IANA Time Zone 'GMT-10'.
// IANA Source File: GMT-10
//  
func (etcTi etcTimeZones) GMTMinus10() string {return "GMT-10" }

// GMT-11 - IANA Time Zone 'GMT-11'.
// IANA Source File: GMT-11
//  
func (etcTi etcTimeZones) GMTMinus11() string {return "GMT-11" }

// GMT-12 - IANA Time Zone 'GMT-12'.
// IANA Source File: GMT-12
//  
func (etcTi etcTimeZones) GMTMinus12() string {return "GMT-12" }

// GMT-13 - IANA Time Zone 'GMT-13'.
// IANA Source File: GMT-13
//  
func (etcTi etcTimeZones) GMTMinus13() string {return "GMT-13" }

// GMT-14 - IANA Time Zone 'GMT-14'.
// IANA Source File: GMT-14
//  
func (etcTi etcTimeZones) GMTMinus14() string {return "GMT-14" }

// GMT0 - IANA Time Zone 'GMT0'.
// IANA Source File: GMT0
//  
func (etcTi etcTimeZones) GMT00() string {return "GMT0" }

// Greenwich - IANA Time Zone 'Greenwich'.
// IANA Source File: Greenwich
//  
func (etcTi etcTimeZones) Greenwich() string {return "Greenwich" }

// UCT - IANA Time Zone 'UCT'.
// IANA Source File: UCT
//  
func (etcTi etcTimeZones) UCT() string {return "UCT" }

// Universal - IANA Time Zone 'Universal'.
// IANA Source File: Universal
//  
func (etcTi etcTimeZones) Universal() string {return "Universal" }

// UTC - IANA Time Zone 'UTC'.
// IANA Source File: UTC
//  
func (etcTi etcTimeZones) UTC() string {return "UTC" }

// Zulu - IANA Time Zone 'Zulu'.
// IANA Source File: Zulu
//  
func (etcTi etcTimeZones) Zulu() string {return "Zulu" }


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

// Amsterdam - IANA Time Zone 'Amsterdam'.
// IANA Source File: Amsterdam
//  
func (europ europeTimeZones) Amsterdam() string {return "Amsterdam" }

// Andorra - IANA Time Zone 'Andorra'.
// IANA Source File: Andorra
//  
func (europ europeTimeZones) Andorra() string {return "Andorra" }

// Astrakhan - IANA Time Zone 'Astrakhan'.
// IANA Source File: Astrakhan
//  
func (europ europeTimeZones) Astrakhan() string {return "Astrakhan" }

// Athens - IANA Time Zone 'Athens'.
// IANA Source File: Athens
//  
func (europ europeTimeZones) Athens() string {return "Athens" }

// Belfast - IANA Time Zone 'Belfast'.
// IANA Source File: Belfast
//  
func (europ europeTimeZones) Belfast() string {return "Belfast" }

// Belgrade - IANA Time Zone 'Belgrade'.
// IANA Source File: Belgrade
//  
func (europ europeTimeZones) Belgrade() string {return "Belgrade" }

// Berlin - IANA Time Zone 'Berlin'.
// IANA Source File: Berlin
//  
func (europ europeTimeZones) Berlin() string {return "Berlin" }

// Bratislava - IANA Time Zone 'Bratislava'.
// IANA Source File: Bratislava
//  
func (europ europeTimeZones) Bratislava() string {return "Bratislava" }

// Brussels - IANA Time Zone 'Brussels'.
// IANA Source File: Brussels
//  
func (europ europeTimeZones) Brussels() string {return "Brussels" }

// Bucharest - IANA Time Zone 'Bucharest'.
// IANA Source File: Bucharest
//  
func (europ europeTimeZones) Bucharest() string {return "Bucharest" }

// Budapest - IANA Time Zone 'Budapest'.
// IANA Source File: Budapest
//  
func (europ europeTimeZones) Budapest() string {return "Budapest" }

// Busingen - IANA Time Zone 'Busingen'.
// IANA Source File: Busingen
//  
func (europ europeTimeZones) Busingen() string {return "Busingen" }

// Chisinau - IANA Time Zone 'Chisinau'.
// IANA Source File: Chisinau
//  
func (europ europeTimeZones) Chisinau() string {return "Chisinau" }

// Copenhagen - IANA Time Zone 'Copenhagen'.
// IANA Source File: Copenhagen
//  
func (europ europeTimeZones) Copenhagen() string {return "Copenhagen" }

// Dublin - IANA Time Zone 'Dublin'.
// IANA Source File: Dublin
//  
func (europ europeTimeZones) Dublin() string {return "Dublin" }

// Gibraltar - IANA Time Zone 'Gibraltar'.
// IANA Source File: Gibraltar
//  
func (europ europeTimeZones) Gibraltar() string {return "Gibraltar" }

// Guernsey - IANA Time Zone 'Guernsey'.
// IANA Source File: Guernsey
//  
func (europ europeTimeZones) Guernsey() string {return "Guernsey" }

// Helsinki - IANA Time Zone 'Helsinki'.
// IANA Source File: Helsinki
//  
func (europ europeTimeZones) Helsinki() string {return "Helsinki" }

// Isle_of_Man - IANA Time Zone 'Isle_of_Man'.
// IANA Source File: Isle_of_Man
//  
func (europ europeTimeZones) Isle_of_Man() string {return "Isle_of_Man" }

// Istanbul - IANA Time Zone 'Istanbul'.
// IANA Source File: Istanbul
//  
func (europ europeTimeZones) Istanbul() string {return "Istanbul" }

// Jersey - IANA Time Zone 'Jersey'.
// IANA Source File: Jersey
//  
func (europ europeTimeZones) Jersey() string {return "Jersey" }

// Kaliningrad - IANA Time Zone 'Kaliningrad'.
// IANA Source File: Kaliningrad
//  
func (europ europeTimeZones) Kaliningrad() string {return "Kaliningrad" }

// Kiev - IANA Time Zone 'Kiev'.
// IANA Source File: Kiev
//  
func (europ europeTimeZones) Kiev() string {return "Kiev" }

// Kirov - IANA Time Zone 'Kirov'.
// IANA Source File: Kirov
//  
func (europ europeTimeZones) Kirov() string {return "Kirov" }

// Lisbon - IANA Time Zone 'Lisbon'.
// IANA Source File: Lisbon
//  
func (europ europeTimeZones) Lisbon() string {return "Lisbon" }

// Ljubljana - IANA Time Zone 'Ljubljana'.
// IANA Source File: Ljubljana
//  
func (europ europeTimeZones) Ljubljana() string {return "Ljubljana" }

// London - IANA Time Zone 'London'.
// IANA Source File: London
//  
func (europ europeTimeZones) London() string {return "London" }

// Luxembourg - IANA Time Zone 'Luxembourg'.
// IANA Source File: Luxembourg
//  
func (europ europeTimeZones) Luxembourg() string {return "Luxembourg" }

// Madrid - IANA Time Zone 'Madrid'.
// IANA Source File: Madrid
//  
func (europ europeTimeZones) Madrid() string {return "Madrid" }

// Malta - IANA Time Zone 'Malta'.
// IANA Source File: Malta
//  
func (europ europeTimeZones) Malta() string {return "Malta" }

// Mariehamn - IANA Time Zone 'Mariehamn'.
// IANA Source File: Mariehamn
//  
func (europ europeTimeZones) Mariehamn() string {return "Mariehamn" }

// Minsk - IANA Time Zone 'Minsk'.
// IANA Source File: Minsk
//  
func (europ europeTimeZones) Minsk() string {return "Minsk" }

// Monaco - IANA Time Zone 'Monaco'.
// IANA Source File: Monaco
//  
func (europ europeTimeZones) Monaco() string {return "Monaco" }

// Moscow - IANA Time Zone 'Moscow'.
// IANA Source File: Moscow
//  
func (europ europeTimeZones) Moscow() string {return "Moscow" }

// Nicosia - IANA Time Zone 'Nicosia'.
// IANA Source File: Nicosia
//  
func (europ europeTimeZones) Nicosia() string {return "Nicosia" }

// Oslo - IANA Time Zone 'Oslo'.
// IANA Source File: Oslo
//  
func (europ europeTimeZones) Oslo() string {return "Oslo" }

// Paris - IANA Time Zone 'Paris'.
// IANA Source File: Paris
//  
func (europ europeTimeZones) Paris() string {return "Paris" }

// Podgorica - IANA Time Zone 'Podgorica'.
// IANA Source File: Podgorica
//  
func (europ europeTimeZones) Podgorica() string {return "Podgorica" }

// Prague - IANA Time Zone 'Prague'.
// IANA Source File: Prague
//  
func (europ europeTimeZones) Prague() string {return "Prague" }

// Riga - IANA Time Zone 'Riga'.
// IANA Source File: Riga
//  
func (europ europeTimeZones) Riga() string {return "Riga" }

// Rome - IANA Time Zone 'Rome'.
// IANA Source File: Rome
//  
func (europ europeTimeZones) Rome() string {return "Rome" }

// Samara - IANA Time Zone 'Samara'.
// IANA Source File: Samara
//  
func (europ europeTimeZones) Samara() string {return "Samara" }

// San_Marino - IANA Time Zone 'San_Marino'.
// IANA Source File: San_Marino
//  
func (europ europeTimeZones) San_Marino() string {return "San_Marino" }

// Sarajevo - IANA Time Zone 'Sarajevo'.
// IANA Source File: Sarajevo
//  
func (europ europeTimeZones) Sarajevo() string {return "Sarajevo" }

// Saratov - IANA Time Zone 'Saratov'.
// IANA Source File: Saratov
//  
func (europ europeTimeZones) Saratov() string {return "Saratov" }

// Simferopol - IANA Time Zone 'Simferopol'.
// IANA Source File: Simferopol
//  
func (europ europeTimeZones) Simferopol() string {return "Simferopol" }

// Skopje - IANA Time Zone 'Skopje'.
// IANA Source File: Skopje
//  
func (europ europeTimeZones) Skopje() string {return "Skopje" }

// Sofia - IANA Time Zone 'Sofia'.
// IANA Source File: Sofia
//  
func (europ europeTimeZones) Sofia() string {return "Sofia" }

// Stockholm - IANA Time Zone 'Stockholm'.
// IANA Source File: Stockholm
//  
func (europ europeTimeZones) Stockholm() string {return "Stockholm" }

// Tallinn - IANA Time Zone 'Tallinn'.
// IANA Source File: Tallinn
//  
func (europ europeTimeZones) Tallinn() string {return "Tallinn" }

// Tirane - IANA Time Zone 'Tirane'.
// IANA Source File: Tirane
//  
func (europ europeTimeZones) Tirane() string {return "Tirane" }

// Tiraspol - IANA Time Zone 'Tiraspol'.
// IANA Source File: Tiraspol
//  
func (europ europeTimeZones) Tiraspol() string {return "Tiraspol" }

// Ulyanovsk - IANA Time Zone 'Ulyanovsk'.
// IANA Source File: Ulyanovsk
//  
func (europ europeTimeZones) Ulyanovsk() string {return "Ulyanovsk" }

// Uzhgorod - IANA Time Zone 'Uzhgorod'.
// IANA Source File: Uzhgorod
//  
func (europ europeTimeZones) Uzhgorod() string {return "Uzhgorod" }

// Vaduz - IANA Time Zone 'Vaduz'.
// IANA Source File: Vaduz
//  
func (europ europeTimeZones) Vaduz() string {return "Vaduz" }

// Vatican - IANA Time Zone 'Vatican'.
// IANA Source File: Vatican
//  
func (europ europeTimeZones) Vatican() string {return "Vatican" }

// Vienna - IANA Time Zone 'Vienna'.
// IANA Source File: Vienna
//  
func (europ europeTimeZones) Vienna() string {return "Vienna" }

// Vilnius - IANA Time Zone 'Vilnius'.
// IANA Source File: Vilnius
//  
func (europ europeTimeZones) Vilnius() string {return "Vilnius" }

// Volgograd - IANA Time Zone 'Volgograd'.
// IANA Source File: Volgograd
//  
func (europ europeTimeZones) Volgograd() string {return "Volgograd" }

// Warsaw - IANA Time Zone 'Warsaw'.
// IANA Source File: Warsaw
//  
func (europ europeTimeZones) Warsaw() string {return "Warsaw" }

// Zagreb - IANA Time Zone 'Zagreb'.
// IANA Source File: Zagreb
//  
func (europ europeTimeZones) Zagreb() string {return "Zagreb" }

// Zaporozhye - IANA Time Zone 'Zaporozhye'.
// IANA Source File: Zaporozhye
//  
func (europ europeTimeZones) Zaporozhye() string {return "Zaporozhye" }

// Zurich - IANA Time Zone 'Zurich'.
// IANA Source File: Zurich
//  
func (europ europeTimeZones) Zurich() string {return "Zurich" }


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

// Antananarivo - IANA Time Zone 'Antananarivo'.
// IANA Source File: Antananarivo
//  
func (india indianTimeZones) Antananarivo() string {return "Antananarivo" }

// Chagos - IANA Time Zone 'Chagos'.
// IANA Source File: Chagos
//  
func (india indianTimeZones) Chagos() string {return "Chagos" }

// Christmas - IANA Time Zone 'Christmas'.
// IANA Source File: Christmas
//  
func (india indianTimeZones) Christmas() string {return "Christmas" }

// Cocos - IANA Time Zone 'Cocos'.
// IANA Source File: Cocos
//  
func (india indianTimeZones) Cocos() string {return "Cocos" }

// Comoro - IANA Time Zone 'Comoro'.
// IANA Source File: Comoro
//  
func (india indianTimeZones) Comoro() string {return "Comoro" }

// Kerguelen - IANA Time Zone 'Kerguelen'.
// IANA Source File: Kerguelen
//  
func (india indianTimeZones) Kerguelen() string {return "Kerguelen" }

// Mahe - IANA Time Zone 'Mahe'.
// IANA Source File: Mahe
//  
func (india indianTimeZones) Mahe() string {return "Mahe" }

// Maldives - IANA Time Zone 'Maldives'.
// IANA Source File: Maldives
//  
func (india indianTimeZones) Maldives() string {return "Maldives" }

// Mauritius - IANA Time Zone 'Mauritius'.
// IANA Source File: Mauritius
//  
func (india indianTimeZones) Mauritius() string {return "Mauritius" }

// Mayotte - IANA Time Zone 'Mayotte'.
// IANA Source File: Mayotte
//  
func (india indianTimeZones) Mayotte() string {return "Mayotte" }

// Reunion - IANA Time Zone 'Reunion'.
// IANA Source File: Reunion
//  
func (india indianTimeZones) Reunion() string {return "Reunion" }


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

// BajaNorte - IANA Time Zone 'BajaNorte'.
// IANA Source File: BajaNorte
//  
func (mexic mexicoTimeZones) BajaNorte() string {return "BajaNorte" }

// BajaSur - IANA Time Zone 'BajaSur'.
// IANA Source File: BajaSur
//  
func (mexic mexicoTimeZones) BajaSur() string {return "BajaSur" }

// General - IANA Time Zone 'General'.
// IANA Source File: General
//  
func (mexic mexicoTimeZones) General() string {return "General" }


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

// Apia - IANA Time Zone 'Apia'.
// IANA Source File: Apia
//  
func (pacif pacificTimeZones) Apia() string {return "Apia" }

// Auckland - IANA Time Zone 'Auckland'.
// IANA Source File: Auckland
//  
func (pacif pacificTimeZones) Auckland() string {return "Auckland" }

// Bougainville - IANA Time Zone 'Bougainville'.
// IANA Source File: Bougainville
//  
func (pacif pacificTimeZones) Bougainville() string {return "Bougainville" }

// Chatham - IANA Time Zone 'Chatham'.
// IANA Source File: Chatham
//  
func (pacif pacificTimeZones) Chatham() string {return "Chatham" }

// Chuuk - IANA Time Zone 'Chuuk'.
// IANA Source File: Chuuk
//  
func (pacif pacificTimeZones) Chuuk() string {return "Chuuk" }

// Easter - IANA Time Zone 'Easter'.
// IANA Source File: Easter
//  
func (pacif pacificTimeZones) Easter() string {return "Easter" }

// Efate - IANA Time Zone 'Efate'.
// IANA Source File: Efate
//  
func (pacif pacificTimeZones) Efate() string {return "Efate" }

// Enderbury - IANA Time Zone 'Enderbury'.
// IANA Source File: Enderbury
//  
func (pacif pacificTimeZones) Enderbury() string {return "Enderbury" }

// Fakaofo - IANA Time Zone 'Fakaofo'.
// IANA Source File: Fakaofo
//  
func (pacif pacificTimeZones) Fakaofo() string {return "Fakaofo" }

// Fiji - IANA Time Zone 'Fiji'.
// IANA Source File: Fiji
//  
func (pacif pacificTimeZones) Fiji() string {return "Fiji" }

// Funafuti - IANA Time Zone 'Funafuti'.
// IANA Source File: Funafuti
//  
func (pacif pacificTimeZones) Funafuti() string {return "Funafuti" }

// Galapagos - IANA Time Zone 'Galapagos'.
// IANA Source File: Galapagos
//  
func (pacif pacificTimeZones) Galapagos() string {return "Galapagos" }

// Gambier - IANA Time Zone 'Gambier'.
// IANA Source File: Gambier
//  
func (pacif pacificTimeZones) Gambier() string {return "Gambier" }

// Guadalcanal - IANA Time Zone 'Guadalcanal'.
// IANA Source File: Guadalcanal
//  
func (pacif pacificTimeZones) Guadalcanal() string {return "Guadalcanal" }

// Guam - IANA Time Zone 'Guam'.
// IANA Source File: Guam
//  
func (pacif pacificTimeZones) Guam() string {return "Guam" }

// Honolulu - IANA Time Zone 'Honolulu'.
// IANA Source File: Honolulu
//  
func (pacif pacificTimeZones) Honolulu() string {return "Honolulu" }

// Johnston - IANA Time Zone 'Johnston'.
// IANA Source File: Johnston
//  
func (pacif pacificTimeZones) Johnston() string {return "Johnston" }

// Kiritimati - IANA Time Zone 'Kiritimati'.
// IANA Source File: Kiritimati
//  
func (pacif pacificTimeZones) Kiritimati() string {return "Kiritimati" }

// Kosrae - IANA Time Zone 'Kosrae'.
// IANA Source File: Kosrae
//  
func (pacif pacificTimeZones) Kosrae() string {return "Kosrae" }

// Kwajalein - IANA Time Zone 'Kwajalein'.
// IANA Source File: Kwajalein
//  
func (pacif pacificTimeZones) Kwajalein() string {return "Kwajalein" }

// Majuro - IANA Time Zone 'Majuro'.
// IANA Source File: Majuro
//  
func (pacif pacificTimeZones) Majuro() string {return "Majuro" }

// Marquesas - IANA Time Zone 'Marquesas'.
// IANA Source File: Marquesas
//  
func (pacif pacificTimeZones) Marquesas() string {return "Marquesas" }

// Midway - IANA Time Zone 'Midway'.
// IANA Source File: Midway
//  
func (pacif pacificTimeZones) Midway() string {return "Midway" }

// Nauru - IANA Time Zone 'Nauru'.
// IANA Source File: Nauru
//  
func (pacif pacificTimeZones) Nauru() string {return "Nauru" }

// Niue - IANA Time Zone 'Niue'.
// IANA Source File: Niue
//  
func (pacif pacificTimeZones) Niue() string {return "Niue" }

// Norfolk - IANA Time Zone 'Norfolk'.
// IANA Source File: Norfolk
//  
func (pacif pacificTimeZones) Norfolk() string {return "Norfolk" }

// Noumea - IANA Time Zone 'Noumea'.
// IANA Source File: Noumea
//  
func (pacif pacificTimeZones) Noumea() string {return "Noumea" }

// Pago_Pago - IANA Time Zone 'Pago_Pago'.
// IANA Source File: Pago_Pago
//  
func (pacif pacificTimeZones) Pago_Pago() string {return "Pago_Pago" }

// Palau - IANA Time Zone 'Palau'.
// IANA Source File: Palau
//  
func (pacif pacificTimeZones) Palau() string {return "Palau" }

// Pitcairn - IANA Time Zone 'Pitcairn'.
// IANA Source File: Pitcairn
//  
func (pacif pacificTimeZones) Pitcairn() string {return "Pitcairn" }

// Pohnpei - IANA Time Zone 'Pohnpei'.
// IANA Source File: Pohnpei
//  
func (pacif pacificTimeZones) Pohnpei() string {return "Pohnpei" }

// Ponape - IANA Time Zone 'Ponape'.
// IANA Source File: Ponape
//  
func (pacif pacificTimeZones) Ponape() string {return "Ponape" }

// Port_Moresby - IANA Time Zone 'Port_Moresby'.
// IANA Source File: Port_Moresby
//  
func (pacif pacificTimeZones) Port_Moresby() string {return "Port_Moresby" }

// Rarotonga - IANA Time Zone 'Rarotonga'.
// IANA Source File: Rarotonga
//  
func (pacif pacificTimeZones) Rarotonga() string {return "Rarotonga" }

// Saipan - IANA Time Zone 'Saipan'.
// IANA Source File: Saipan
//  
func (pacif pacificTimeZones) Saipan() string {return "Saipan" }

// Samoa - IANA Time Zone 'Samoa'.
// IANA Source File: Samoa
//  
func (pacif pacificTimeZones) Samoa() string {return "Samoa" }

// Tahiti - IANA Time Zone 'Tahiti'.
// IANA Source File: Tahiti
//  
func (pacif pacificTimeZones) Tahiti() string {return "Tahiti" }

// Tarawa - IANA Time Zone 'Tarawa'.
// IANA Source File: Tarawa
//  
func (pacif pacificTimeZones) Tarawa() string {return "Tarawa" }

// Tongatapu - IANA Time Zone 'Tongatapu'.
// IANA Source File: Tongatapu
//  
func (pacif pacificTimeZones) Tongatapu() string {return "Tongatapu" }

// Truk - IANA Time Zone 'Truk'.
// IANA Source File: Truk
//  
func (pacif pacificTimeZones) Truk() string {return "Truk" }

// Wake - IANA Time Zone 'Wake'.
// IANA Source File: Wake
//  
func (pacif pacificTimeZones) Wake() string {return "Wake" }

// Wallis - IANA Time Zone 'Wallis'.
// IANA Source File: Wallis
//  
func (pacif pacificTimeZones) Wallis() string {return "Wallis" }

// Yap - IANA Time Zone 'Yap'.
// IANA Source File: Yap
//  
func (pacif pacificTimeZones) Yap() string {return "Yap" }


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

// Alaska - IANA Time Zone 'Alaska'.
// IANA Source File: Alaska
//  
func (uSTim uSTimeZones) Alaska() string {return "Alaska" }

// Aleutian - IANA Time Zone 'Aleutian'.
// IANA Source File: Aleutian
//  
func (uSTim uSTimeZones) Aleutian() string {return "Aleutian" }

// Arizona - IANA Time Zone 'Arizona'.
// IANA Source File: Arizona
//  
func (uSTim uSTimeZones) Arizona() string {return "Arizona" }

// Central - IANA Time Zone 'Central'.
// IANA Source File: Central
//  
func (uSTim uSTimeZones) Central() string {return "Central" }

// East-Indiana - IANA Time Zone 'East-Indiana'.
// IANA Source File: East-Indiana
//  
func (uSTim uSTimeZones) EastMinusIndiana() string {return "East-Indiana" }

// Eastern - IANA Time Zone 'Eastern'.
// IANA Source File: Eastern
//  
func (uSTim uSTimeZones) Eastern() string {return "Eastern" }

// Hawaii - IANA Time Zone 'Hawaii'.
// IANA Source File: Hawaii
//  
func (uSTim uSTimeZones) Hawaii() string {return "Hawaii" }

// Indiana-Starke - IANA Time Zone 'Indiana-Starke'.
// IANA Source File: Indiana-Starke
//  
func (uSTim uSTimeZones) IndianaMinusStarke() string {return "Indiana-Starke" }

// Michigan - IANA Time Zone 'Michigan'.
// IANA Source File: Michigan
//  
func (uSTim uSTimeZones) Michigan() string {return "Michigan" }

// Mountain - IANA Time Zone 'Mountain'.
// IANA Source File: Mountain
//  
func (uSTim uSTimeZones) Mountain() string {return "Mountain" }

// Pacific - IANA Time Zone 'Pacific'.
// IANA Source File: Pacific
//  
func (uSTim uSTimeZones) Pacific() string {return "Pacific" }

// Samoa - IANA Time Zone 'Samoa'.
// IANA Source File: Samoa
//  
func (uSTim uSTimeZones) Samoa() string {return "Samoa" }


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

