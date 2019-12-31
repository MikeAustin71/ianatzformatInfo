package tzdatastructs

import (
	"errors"
)

// TzAbbreviationDto - encapsulates Time Zone abbreviation
// information. A Time Zone Abbreviation must consist entirely
// of alphabetic characters.
//
// The Id is styled as Abbreviation text plus the UTC offset.
// Example: CST-0600 - Central Standard time with offset UTC-0600.
//
type TzAbbreviationDto struct {
	Id         string
	Abbrv      string
	TzName     string
	Location   string
	UtcOffset  string
	IanaZone   string
}

// CopyOut() - Makes and returns a deep copy of the current TzAbbreviationDto
// object.
//
func (TzAbbrv *TzAbbreviationDto) CopyOut() TzAbbreviationDto {

	newDto := TzAbbreviationDto{}

	newDto.Id = TzAbbrv.Id
	newDto.Abbrv = TzAbbrv.Abbrv
	newDto.TzName = TzAbbrv.TzName
	newDto.Location = TzAbbrv.Location
	newDto.UtcOffset = TzAbbrv.UtcOffset
	newDto.IanaZone = TzAbbrv.IanaZone

	return newDto
}

// CopyIn() - Copies the field values from an incoming TzAbbreviationDto
// object to the current TzAbbreviationDto object.
func (TzAbbrv *TzAbbreviationDto) CopyIn(inComing *TzAbbreviationDto) error {

	ePrefix := "TzAbbreviationDto.CopyIn() "

	if inComing == nil {
		return  errors.New(ePrefix +
			"Error: Input parameter 'incoming' is nil!")
	}

	TzAbbrv.Id = inComing.Id
	TzAbbrv.Abbrv = inComing.Abbrv
	TzAbbrv.TzName = inComing.TzName
	TzAbbrv.Location = inComing.Location
	TzAbbrv.UtcOffset = inComing.UtcOffset
	TzAbbrv.IanaZone = inComing.IanaZone

	return nil
}

// TzAbbreviationReference - Includes a reference map covering all known
// and valid Time Zone abbreviations.
var TzAbbreviationReference = map[string]TzAbbreviationDto{
	"-00+0000"    :{"-00+0000","-00","Unassigned Time Zone Name -00","-00","+0000",""},
	"-01-0100"    :{"-01-0100","-01","Unassigned Time Zone Name -01","-01","-0100",""},
	"-02-0200"    :{"-02-0200","-02","Unassigned Time Zone Name -02","-02","-0200",""},
	"-03-0300"    :{"-03-0300","-03","Unassigned Time Zone Name -03","-03","-0300",""},
	"-04-0400"    :{"-04-0400","-04","Unassigned Time Zone Name -04","-04","-0400",""},
	"-05-0500"    :{"-05-0500","-05","Unassigned Time Zone Name -05","-05","-0500",""},
	"-06-0600"    :{"-06-0600","-06","Unassigned Time Zone Name -06","-06","-0600",""},
	"-07-0700"    :{"-07-0700","-07","Unassigned Time Zone Name -07","-07","-0700",""},
	"-08-0800"    :{"-08-0800","-08","Unassigned Time Zone Name -08","-08","-0800",""},
	"-09-0900"    :{"-09-0900","-09","Unassigned Time Zone Name -09","-09","-0900",""},
	"-0930-0930"    :{"-0930-0930","-0930","Unassigned Time Zone Name -0930","-0930","-0930",""},
	"-10-1000"    :{"-10-1000","-10","Unassigned Time Zone Name -10","-10","-1000",""},
	"-11-1100"    :{"-11-1100","-11","Unassigned Time Zone Name -11","-11","-1100",""},
	"-12-1200"    :{"-12-1200","-12","Unassigned Time Zone Name -12","-12","-1200",""},
	"+00+0000"    :{"+00+0000","+00","Unassigned Time Zone Name +00","+00","+0000",""},
	"+01+0100"    :{"+01+0100","+01","Unassigned Time Zone Name +01","+01","+0100",""},
	"+02+0200"    :{"+02+0200","+02","Unassigned Time Zone Name +02","+02","+0200",""},
	"+03+0300"    :{"+03+0300","+03","Unassigned Time Zone Name +03","+03","+0300",""},
	"+04+0400"    :{"+04+0400","+04","Unassigned Time Zone Name +04","+04","+0400",""},
	"+0430+0430"  :{"+0430+0430","+0430","Unassigned Time Zone Name +0430","+0430","+0430",""},
	"+05+0500"    :{"+05+0500","+05","Unassigned Time Zone Name +05","+05","+0500",""},
	"+0530+0530"  :{"+0530+0530","+0530","Unassigned Time Zone Name +0530","+0530","+0500",""},
	"+0545+0545"  :{"+0545+0545","+0545","Unassigned Time Zone Name +0545","+0545","+0545",""},
	"+06+0600"    :{"+06+0600","+06","Unassigned Time Zone Name +06","+06","+0600",""},
	"+0630+0630"  :{"+0630+0630","+0630","Unassigned Time Zone Name +0630","+0630","+0630",""},
	"+07+0700"    :{"+07+0700","+07","Unassigned Time Zone Name +07","+07","+0700",""},
	"+08+0800"    :{"+08+0800","+08","Unassigned Time Zone Name +08","+08","+0800",""},
	"+0845+0845"  :{"+0845+0845","+0845","Unassigned Time Zone Name +0845","+0845","+0845",""},
	"+09+0900"    :{"+09+0900","+09","Unassigned Time Zone Name +09","+09","+0900",""},
	"+10+1000"    :{"+10+1000","+10","Unassigned Time Zone Name +10","+10","+1000",""},
	"+1030+1030"  :{"+1030+1030","+1030","Unassigned Time Zone Name +1030","+1030","+1030",""},
	"+11+1100"    :{"+11+1100","+11","Unassigned Time Zone Name +11","+11","+1100",""},
	"+12+1200"    :{"+12+1200","+12","Unassigned Time Zone Name +12","+12","+1200",""},
	"+1245+1245"    :{"+1245+1245","+1245","Unassigned Time Zone Name +1245","+1245","+1245",""},
	"+13+1300"    :{"+13+1300","+13","Unassigned Time Zone Name +13","+13","+1300",""},
	"+14+1400"    :{"+14+1400","+14","Unassigned Time Zone Name +14","+14","+1400",""},
	"A+0100"      :{"A+0100","A","Alpha Time Zone","Military","+0100",""},
	"ACDT+1030"   :{"ACDT+1030","ACDT","Australian Central Daylight Time","Australia","+1030",""},
	"ACST+0930"   :{"ACST+0930","ACST","Australian Central Standard Time","Australia","+0930",""},
	"ADT-0300"      :{"ADT-0300","ADT","Atlantic Daylight Time","North America","-0300",""},
	"AEDT+1100"      :{"AEDT+1100","AEDT","Australian Eastern Daylight Time","Australia","+1100",""},
	"AEST+1000"      :{"AEST+1000","AEST","Australian Eastern Standard Time","Australia","+1000",""},
	"AFT+0430"      :{"AFT+0430","AFT","Afghanistan Time","Asia","+0430",""},
	"AKDT-0800"      :{"AKDT-0800","AKDT","Alaska Daylight Time","North America","-0800",""},
	"AKST-0900"      :{"AKST-0900","AKST","Alaska Standard Time","North America","-0900",""},
	"ALMT+0600"      :{"ALMT+0600","ALMT","Alma-Ata Time","Asia","+0600",""},
	"AMST+0500"      :{"AMST+0500","AMST","Armenia Summer Time","Asia","+0500",""},
	"AMST-0300"      :{"AMST-0300","AMST","Amazon Summer Time","South America","-0300",""},
	"AMT+0400"      :{"AMT+0400","AMT","Armenia Time","Asia","+0400",""},
	"AMT-0400"      :{"AMT-0400","AMT","Amazon Time","South America","-0400",""},
	"ANAST+1200"      :{"ANAST+1200","ANAST","Anadyr Summer Time","Asia","+1200",""},
	"ANAT+1200"      :{"ANAT+1200","ANAT","Anadyr Time","Asia","+1200",""},
	"AQTT+0500"      :{"AQTT+0500","AQTT","Aqtobe Time","Asia","+0500",""},
	"ART-0300"      :{"ART-0300","ART","Argentina Time","South America","-0300",""},
	"AST+0300"      :{"AST+0300","AST","Arabia Standard Time","Asia","+0300",""},
	"AST-0400"      :{"AST-0400","AST","Atlantic Standard Time","Atlantic","-0400",""},
	"AWDT+0900"      :{"AWDT+0900","AWDT","Australian Western Daylight Time","Australia","+0900",""},
	"AWST+0800"      :{"AWST+0800","AWST","Australian Western Standard Time","Australia","+0800",""},
	"AZOST+0000"      :{"AZOST+0000","AZOST","Azores Summer Time","Atlantic","+0000",""},
	"AZOT-0100"      :{"AZOT-0100","AZOT","Azores Time","Atlantic","-0100",""},
	"AZST+0500"      :{"AZST+0500","AZST","Azerbaijan Summer Time","Asia","+0500",""},
	"AZT+0400"      :{"AZT+0400","AZT","Azerbaijan Time","Asia","+0400",""},
	"B+0200"      :{"B+0200","B","Bravo Time Zone","Military","+0200",""},
	"BNT+0800"      :{"BNT+0800","BNT","Brunei Darussalam Time","Asia","+0800",""},
	"BOT-0400"      :{"BOT-0400","BOT","Bolivia Time","South America","-0400",""},
	"BRST-0200"      :{"BRST-0200","BRST","Brasilia Summer Time","South America","-0200",""},
	"BRT-0300"      :{"BRT-0300","BRT","Brasília time","South America","-0300",""},
	"BST+0600"      :{"BST+0600","BST","Bangladesh Standard Time","Asia","+0600",""},
	"BST+0100"      :{"BST+0100","BST","British Summer Time","Europe","+0100",""},
	"BTT+0600"      :{"BTT+0600","BTT","Bhutan Time","Asia","+0600",""},
	"C+0300"      :{"C+0300","C","Charlie Time Zone","Military","+0300",""},
	"CAST+0800"      :{"CAST+0800","CAST","Casey Time","Antarctica","+0800",""},
	"CAT+0200"      :{"CAT+0200","CAT","Central Africa Time","Africa","+0200",""},
	"CCT+0630"      :{"CCT+0630","CCT","Cocos Islands Time","Indian Ocean","+0630",""},
	"CDT-0400"      :{"CDT-0400","CDT","Cuba Daylight Time","Caribbean","-0400",""},
	"CDT-0500"      :{"CDT-0500","CDT","Central Daylight Time","North America","-0500",""},
	"CEST+0200"      :{"CEST+0200","CEST","Central European Summer Time","Europe","+0200",""},
	"CET+0100"      :{"CET+0100","CET","Central European Time","Europe","+0100",""},
	"CHADT+1345"      :{"CHADT+1345","CHADT","Chatham Island Daylight Time","Pacific","+1345",""},
	"CHAST+1245"      :{"CHAST+1245","CHAST","Chatham Island Standard Time","Pacific","+1245",""},
	"CKT-1000"      :{"CKT-1000","CKT","Cook Island Time","Pacific","-1000",""},
	"CLST-0300"      :{"CLST-0300","CLST","Chile Summer Time","South America","-0300",""},
	"CLT-0400"      :{"CLT-0400","CLT","Chile Standard Time","South America","-0400",""},
	"COT-0500"      :{"COT-0500","COT","Colombia Time","South America","-0500",""},
	"CST+0800"      :{"CST+0800","CST","China Standard Time","Asia","+0800",""},
	"CST-0500"      :{"CST-0500","CST","Cuba Standard Time","Caribbean","-0500",""},
	"CST-0600"      :{"CST-0600","CST","Central Standard Time","North America","-0600",""},
	"CVT-0100"      :{"CVT-0100","CVT","Cape Verde Time","Africa","-0100",""},
	"CXT+0700"      :{"CXT+0700","CXT","Christmas Island Time","Australia","+0700",""},
	"ChST+1000"      :{"ChST+1000","ChST","Chamorro Standard Time","Pacific","+1000",""},
	"D+0400"      :{"D+0400","D","Delta Time Zone","Military","+0400",""},
	"DAVT+0700"      :{"DAVT+0700","DAVT","Davis Time","Antarctica","+0700",""},
	"E+0500"      :{"E+0500","E","Echo Time Zone","Military","+0500",""},
	"EASST-0500"      :{"EASST-0500","EASST","Easter Island Summer Time","Pacific","-0500",""},
	"EAST-0600"      :{"EAST-0600","EAST","Easter Island Standard Time","Pacific","-0600",""},
	"EAT+0300"      :{"EAT+0300","EAT","Eastern Africa Time","Africa","+0300",""},
	"ECT-0500"      :{"ECT-0500","ECT","Ecuador Time","South America","-0500",""},
	"EDT-0400"      :{"EDT-0400","EDT","Eastern Daylight Time","North America","-0400",""},
	"EDT+1100"      :{"EDT+1100","EDT","Eastern Daylight Time","Pacific","+1100",""},
	"EEST+0300"      :{"EEST+0300","EEST","Eastern European Summer Time","Europe","+0300",""},
	"EET+0200"      :{"EET+0200","EET","Eastern European Time","Europe","+0200",""},
	"EGST+0000"      :{"EGST+0000","EGST","Eastern Greenland Summer Time","North America","+0000",""},
	"EGT-0100"      :{"EGT-0100","EGT","East Greenland Time","North America","-0100",""},
	"EST-0500"      :{"EST-0500","EST","Eastern Standard Time","North America","-0500",""},
	"ET-0500"      :{"ET-0500","ET","Tiempo Del Este","North America","-0500",""},
	"F+0600"      :{"F+0600","F","Foxtrot Time Zone","Military","+0600",""},
	"FJST+1300"      :{"FJST+1300","FJST","Fiji Summer Time","Pacific","+1300",""},
	"FJT+1200"      :{"FJT+1200","FJT","Fiji Time","Pacific","+1200",""},
	"FKST-0300"      :{"FKST-0300","FKST","Falkland Islands Summer Time","South America","-0300",""},
	"FKT-0400"      :{"FKT-0400","FKT","Falkland Island Time","South America","-0400",""},
	"FNT-0200"      :{"FNT-0200","FNT","Fernando de Noronha Time","South America","-0200",""},
	"G+0700"        :{"G+0700","G","Golf Time Zone","Military","+0700",""},
	"GALT-0600"      :{"GALT-0600","GALT","Galapagos Time","Pacific","-0600",""},
	"GAMT-0900"      :{"GAMT-0900","GAMT","Gambier Time","Pacific","-0900",""},
	"GET+0400"      :{"GET+0400","GET","Georgia Standard Time","Asia","+0400",""},
	"GFT-0300"      :{"GFT-0300","GFT","French Guiana Time","South America","-0300",""},
	"GILT+1200"      :{"GILT+1200","GILT","Gilbert Island Time","Pacific","+1200",""},
	"GMT+0000"      :{"GMT+0000","GMT","Greenwich Mean Time","Europe","+0000",""},
	"GST+0400"      :{"GST+0400","GST","Gulf Standard Time","Asia","+0400",""},
	"GYT-0400"      :{"GYT-0400","GYT","Guyana Time","South America","-0400",""},
	"H+0800"      :{"H+0800","H","Hotel Time Zone","Military","+0800",""},
	"HAA-0300"      :{"HAA-0300","HAA","Heure Avancée de l’Atlantique","North America","-0300",""},
	"HAC-0500"      :{"HAC-0500","HAC","Heure Avancée du Centre","North America","-0500",""},
	"HADT-0900"      :{"HADT-0900","HADT","Hawaii-Aleutian Daylight Time","North America","-0900",""},
	"HAE-0400"      :{"HAE-0400","HAE","Heure Avancée de l’Est","North America","-0400",""},
	"HAP-0700"      :{"HAP-0700","HAP","Heure Avancée du Pacifique","North America","-0700",""},
	"HAR-0600"      :{"HAR-0600","HAR","Heure Avancée des Rocheuses","North America","-0600",""},
	"HAST-1000"      :{"HAST-1000","HAST","Hawaii-Aleutian Standard Time","North America","-1000",""},
	"HAT-0230"      :{"HAT-0230","HAT","Heure Avancée de Terre-Neuve","North America","-0230",""},
	"HAY-0800"      :{"HAY-0800","HAY","Heure Avancée du Yukon","North America","-0800",""},
	"HKT+0800"      :{"HKT+0800","HKT","Hong Kong Time","Asia","+0800",""},
	"HLV-0430"      :{"HLV-0430","HLV","Hora Legal de Venezuela","South America","-0430",""},
	"HNA-0400"      :{"HNA-0400","HNA","Heure Normale de l’Atlantique","North America","-0400",""},
	"HNC-0600"      :{"HNC-0600","HNC","Heure Normale du Centre","North America","-0600",""},
	"HNE-0500"      :{"HNE-0500","HNE","Heure Normale de l’Est","North America","-0500",""},
	"HNP-0800"      :{"HNP-0800","HNP","Heure Normale du Pacifique","North America","-0800",""},
	"HNR-0700"      :{"HNR-0700","HNR","Heure Normale des Rocheuses","North America","-0700",""},
	"HNT-0300"      :{"HNT-0300","HNT","Heure Normale de Terre-Neuve","North America","-0300",""},
	"HNY-0900"      :{"HNY-0900","HNY","Heure Normale du Yukon","North America","-0900",""},
	"HOVT+0700"      :{"HOVT+0700","HOVT","Hovd Time","Asia","+0700",""},
	"HDT-0900"      :{"HDT-0900", "HDT", "Hawaii-Aleutian Daylight Time", "Hawaii, Aleutians","-0900", ""},
	"HST-1000"      :{"HST-1000", "HST", "Hawaii-Aleutian Standard Time", "Hawaii, Aleutians", "-1000", ""},
	"I+0900"      :{"I+0900","I","India Time Zone","Military","+0900",""},
	"ICT+0700"      :{"ICT+0700","ICT","Indochina Time","Asia","+0700",""},
	"IDT+0300"      :{"IDT+0300","IDT","Israel Daylight Time","Asia","+0300",""},
	"IOT+0600"      :{"IOT+0600","IOT","Indian Chagos Time","Indian Ocean","+0600",""},
	"IRDT+0430"      :{"IRDT+0430","IRDT","Iran Daylight Time","Asia","+0430",""},
	"IRKST+0900"      :{"IRKST+0900","IRKST","Irkutsk Summer Time","Asia","+0900",""},
	"IRKT+0900"      :{"IRKT+0900","IRKT","Irkutsk Time","Asia","+0900",""},
	"IRST+0330"      :{"IRST+0330","IRST","Iran Standard Time","Asia","+0330",""},
	"IST+0200"      :{"IST+0200","IST","Israel Standard Time","Asia","+0200",""},
	"IST+0530"      :{"IST+0530","IST","India Standard Time","Asia","+0530",""},
	"IST+0100"      :{"IST+0100","IST","Irish Standard Time","Europe","+0100",""},
	"JST+0900"      :{"JST+0900","JST","Japan Standard Time","Asia","+0900",""},
	"K+1000"      :{"K+1000","K","Kilo Time Zone","Military","+1000",""},
	"KGT+0600"      :{"KGT+0600","KGT","Kyrgyzstan Time","Asia","+0600",""},
	"KRAST+0800"      :{"KRAST+0800","KRAST","Krasnoyarsk Summer Time","Asia","+0800",""},
	"KRAT+0800"      :{"KRAT+0800","KRAT","Krasnoyarsk Time","Asia","+0800",""},
	"KST+0900"      :{"KST+0900","KST","Korea Standard Time","Asia","+0900",""},
	"KUYT+0400"      :{"KUYT+0400","KUYT","Kuybyshev Time","Europe","+0400",""},
	"L+1100"      :{"L+1100","L","Lima Time Zone","Military","+1100",""},
	"LHDT+1100"      :{"LHDT+1100","LHDT","Lord Howe Daylight Time","Australia","+1100",""},
	"LHST+1030"      :{"LHST+1030","LHST","Lord Howe Standard Time","Australia","+1030",""},
	"LINT+1400"      :{"LINT+1400","LINT","Line Islands Time","Pacific","+1400",""},
	"M+1200"      :{"M+1200","M","Mike Time Zone","Military","+1200",""},
	"MAGST+1200"      :{"MAGST+1200","MAGST","Magadan Summer Time","Asia","+1200",""},
	"MAGT+1200"      :{"MAGT+1200","MAGT","Magadan Time","Asia","+1200",""},
	"MART-0930"      :{"MART-0930","MART","Marquesas Time","Pacific","-0930",""},
	"MAWT+0500"      :{"MAWT+0500","MAWT","Mawson Time","Antarctica","+0500",""},
	"MDT-0600"       :{"MDT-0600","MDT","Mountain Daylight Time","North America","-0600",""},
	"MET+0100"       :{"MET+0100", "MET", "Middle European Time", "Europe", "+0100", ""},
	"MEST+0200"      :{"MEST+0200", "MEST", "Middle European Summer Time", "Europe", "+0200", ""},
	"MESZ+0200"      :{"MESZ+0200","MESZ","Mitteleuropäische Sommerzeit","Europe","+0200",""},
	"MEZ+0100"      :{"MEZ+0100","MEZ","Mitteleuropäische Zeit","Africa","+0100",""},
	"MHT+1200"      :{"MHT+1200","MHT","Marshall Islands Time","Pacific","+1200",""},
	"MMT+0630"      :{"MMT+0630","MMT","Myanmar Time","Asia","+0630",""},
	"MSD+0400"      :{"MSD+0400","MSD","Moscow Daylight Time","Europe","+0400",""},
	"MSK+0300"      :{"MSK+0300","MSK","Moscow Standard Time","Europe","+0300",""},
	"MST-0700"      :{"MST-0700","MST","Mountain Standard Time","North America","-0700",""},
	"MUT+0400"      :{"MUT+0400","MUT","Mauritius Time","Africa","+0400",""},
	"MVT+0500"      :{"MVT+0500","MVT","Maldives Time","Asia","+0500",""},
	"MYT+0800"      :{"MYT+0800","MYT","Malaysia Time","Asia","+0800",""},
	"N-0100"      :{"N-0100","N","November Time Zone","Military","-0100",""},
	"NCT+1100"      :{"NCT+1100","NCT","New Caledonia Time","Pacific","+1100",""},
	"NDT-0230"      :{"NDT-0230","NDT","Newfoundland Daylight Time","North America","-0230",""},
	"NFT+1130"      :{"NFT+1130","NFT","Norfolk Time","Australia","+1130",""},
	"NOVST+0700"      :{"NOVST+0700","NOVST","Novosibirsk Summer Time","Asia","+0700",""},
	"NOVT+0600"      :{"NOVT+0600","NOVT","Novosibirsk Time","Asia","+0600",""},
	"NPT+0545"      :{"NPT+0545","NPT","Nepal Time","Asia","+0545",""},
	"NST-0330"      :{"NST-0330","NST","Newfoundland Standard Time","North America","-0330",""},
	"NUT-1100"      :{"NUT-1100","NUT","Niue Time","Pacific","-1100",""},
	"NZDT+1300"      :{"NZDT+1300","NZDT","New Zealand Daylight Time","Pacific","+1300",""},
	"NZST+1200"      :{"NZST+1200","NZST","New Zealand Standard Time","Pacific","+1200",""},
	"O-0200"      :{"O-0200","O","Oscar Time Zone","Military","-0200",""},
	"OMSST+0700"      :{"OMSST+0700","OMSST","Omsk Summer Time","Asia","+0700",""},
	"OMST+0700"      :{"OMST+0700","OMST","Omsk Standard Time","Asia","+0700",""},
	"P-0300"      :{"P-0300","P","Papa Time Zone","Military","-0300",""},
	"PDT-0700"      :{"PDT-0700","PDT","Pacific Daylight Time","North America","-0700",""},
	"PET-0500"      :{"PET-0500","PET","Peru Time","South America","-0500",""},
	"PETST+1200"      :{"PETST+1200","PETST","Kamchatka Summer Time","Asia","+1200",""},
	"PETT+1200"      :{"PETT+1200","PETT","Kamchatka Time","Asia","+1200",""},
	"PGT+1000"      :{"PGT+1000","PGT","Papua New Guinea Time","Pacific","+1000",""},
	"PHOT+1300"      :{"PHOT+1300","PHOT","Phoenix Island Time","Pacific","+1300",""},
	"PHT+0800"      :{"PHT+0800","PHT","Philippine Time","Asia","+0800",""},
	"PKT+0500"      :{"PKT+0500","PKT","Pakistan Standard Time","Asia","+0500",""},
	"PMDT-0200"      :{"PMDT-0200","PMDT","Pierre & Miquelon Daylight Time","North America","-0200",""},
	"PMST-0300"      :{"PMST-0300","PMST","Pierre & Miquelon Standard Time","North America","-0300",""},
	"PONT+1100"      :{"PONT+1100","PONT","Pohnpei Standard Time","Pacific","+1100",""},
	"PST-0800"      :{"PST-0800","PST","Pacific Standard Time","North America","-0800",""},
	"PST+0800"      :{"PST+0800","PST","Philippine Standard Time","Asia","+0800",""},
	"PT-0800"      :{"PT-0800","PT","Tiempo del Pacífico","North America","-0800",""},
	"PWT+0900"      :{"PWT+0900","PWT","Palau Time","Pacific","+0900",""},
	"PYST-0300"      :{"PYST-0300","PYST","Paraguay Summer Time","South America","-0300",""},
	"PYT-0400"      :{"PYT-0400","PYT","Paraguay Time","South America","-0400",""},
	"Q-0400"      :{"Q-0400","Q","Quebec Time Zone","Military","-0400",""},
	"R-0500"      :{"R-0500","R","Romeo Time Zone","Military","-0500",""},
	"RET+0400"      :{"RET+0400","RET","Reunion Time","Africa","+0400",""},
	"S-0600"      :{"S-0600","S","Sierra Time Zone","Military","-0600",""},
	"SAMT+0400"      :{"SAMT+0400","SAMT","Samara Time","Europe","+0400",""},
	"SAST+0200"      :{"SAST+0200","SAST","South Africa Standard Time","Africa","+0200",""},
	"SBT+1100"      :{"SBT+1100","SBT","Solomon IslandsTime","Pacific","+1100",""},
	"SCT+0400"      :{"SCT+0400","SCT","Seychelles Time","Africa","+0400",""},
	"SGT+0800"      :{"SGT+0800","SGT","Singapore Time","Asia","+0800",""},
	"SRT-0300"      :{"SRT-0300","SRT","Suriname Time","South America","-0300",""},
	"SST-1100"      :{"SST-1100","SST","Samoa Standard Time","Pacific","-1100",""},
	"T-0700"      :{"T-0700","T","Tango Time Zone","Military","-0700",""},
	"TAHT-1000"      :{"TAHT-1000","TAHT","Tahiti Time","Pacific","-1000",""},
	"TFT+0500"      :{"TFT+0500","TFT","French Southern and Antarctic Time","Indian Ocean","+0500",""},
	"TJT+0500"      :{"TJT+0500","TJT","Tajikistan Time","Asia","+0500",""},
	"TKT+1300"      :{"TKT+1300","TKT","Tokelau Time","Pacific","+1300",""},
	"TLT+0900"      :{"TLT+0900","TLT","East Timor Time","Asia","+0900",""},
	"TMT+0500"      :{"TMT+0500","TMT","Turkmenistan Time","Asia","+0500",""},
	"TVT+1200"      :{"TVT+1200","TVT","Tuvalu Time","Pacific","+1200",""},
	"U-0800"      :{"U-0800","U","Uniform Time Zone","Military","-0800",""},
	"ULAT+0800"      :{"ULAT+0800","ULAT","Ulaanbaatar Time","Asia","+0800",""},
	"UTC+0000"      :{"UTC+0000","UTC","Universal Time Coordinated","Universal","+0000",""},
	"UYST-0200"      :{"UYST-0200","UYST","Uruguay Summer Time","South America","-0200",""},
	"UYT-0300"      :{"UYT-0300","UYT","Uruguay Time","South America","-0300",""},
	"UZT+0500"      :{"UZT+0500","UZT","Uzbekistan Time","Asia","+0500",""},
	"V-0900"      :{"V-0900","V","Victor Time Zone","Military","-0900",""},
	"VET-0430"      :{"VET-0430","VET","Venezuelan Standard Time","South America","-0430",""},
	"VLAST+1100"      :{"VLAST+1100","VLAST","Vladivostok Summer Time","Asia","+1100",""},
	"VLAT+1100"      :{"VLAT+1100","VLAT","Vladivostok Time","Asia","+1100",""},
	"VUT+1100"      :{"VUT+1100","VUT","Vanuatu Time","Pacific","+1100",""},
	"W-1000"      :{"W-1000","W","Whiskey Time Zone","Military","-1000",""},
	"WAST+0200"      :{"WAST+0200","WAST","West Africa Summer Time","Africa","+0200",""},
	"WAT+0100"      :{"WAT+0100","WAT","West Africa Time","Africa","+0100",""},
	"WEST+0100"      :{"WEST+0100","WEST","Western European Summer Time","Europe","+0100",""},
	"WESZ+0100"      :{"WESZ+0100","WESZ","Westeuropäische Sommerzeit","Africa","+0100",""},
	"WET+0000"      :{"WET+0000","WET","Western European Time","Europe","+0000",""},
	"WEZ+0000"      :{"WEZ+0000","WEZ","Westeuropäische Zeit","Europe","+0000",""},
	"WFT+1200"      :{"WFT+1200","WFT","Wallis and Futuna Time","Pacific","+1200",""},
	"WGST-0200"      :{"WGST-0200","WGST","Western Greenland Summer Time","North America","-0200",""},
	"WGT-0300"      :{"WGT-0300","WGT","West Greenland Time","North America","-0300",""},
	"WIB+0700"      :{"WIB+0700","WIB","Western Indonesian Time","Asia","+0700",""},
	"WIT+0900"      :{"WIT+0900","WIT","Eastern Indonesian Time","Asia","+0900",""},
	"WITA+0800"      :{"WITA+0800","WITA","Central Indonesian Time","Asia","+0800",""},
	"WST+0100"      :{"WST+0100","WST","Western Sahara Summer Time","Africa","+0100",""},
	"WST+1300"      :{"WST+1300","WST","West Samoa Time","Pacific","+1300",""},
	"WT+0000"      :{"WT+0000","WT","Western Sahara Standard Time","Africa","+0000",""},
	"X-1100"      :{"X-1100","X","X-ray Time Zone","Military","-1100",""},
	"Y-1200"      :{"Y-1200","Y","Yankee Time Zone","Military","-1200",""},
	"YAKST+1000"      :{"YAKST+1000","YAKST","Yakutsk Summer Time","Asia","+1000",""},
	"YAKT+1000"      :{"YAKT+1000","YAKT","Yakutsk Time","Asia","+1000",""},
	"YAPT+1000"      :{"YAPT+1000","YAPT","Yap Time","Pacific","+1000",""},
	"YEKST+0600"      :{"YEKST+0600","YEKST","Yekaterinburg Summer Time","Asia","+0600",""},
	"YEKT+0500"      :{"YEKT+0500","YEKT","Yekaterinburg Time","Asia","+0500",""},
	"Z+0000"      :{"Z+0000","Z","Zulu Time Zone","Military","+0000",""},
}



