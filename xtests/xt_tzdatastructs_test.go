package xtests

import (
	"local.com/amarillomike/ianatzformatInfo/tzdatastructs"
	"testing"
)

func TestTimeZoneDataDto_01(t *testing.T) {

	/*
		parentGroupName,
		groupName,
		tzName,
		tzAliasValue,
		tzCanonicalValue,
		tzValue,
		tzSortValue,
		funcSelfReferenceVariable,
		funcType,
		funcName,
		funcReturnType,
		funcReturnValue,
		srcFileNameExt string,
		tzClass TimeZoneClass,
		tzType TimeZoneType,
		deprecationStatus TimeZoneDeprecationStatus) (TimeZoneDataDto, error) {

	*/
	tzDto, err := tzdatastructs.TimeZoneDataDto{}.New(
		"",
		"America",
		"America/Chicago",
		"",
		"America/Chicago",
		"America/Chicago",
		"UTC-0500",
		"America/Chicago",
		0,
		"",
		"Chicago",
		"string",
		"",
		"America/Chicago",
		"northamerica",
		2,
		tzdatastructs.TZClass.Canonical(),
		tzdatastructs.TZType.Standard(),
		tzdatastructs.TZCat.TimeZone(),
		tzdatastructs.TZSrc.Iana(),
		tzdatastructs.DepStatusCode.Valid())

	if err != nil {
		t.Errorf("Error returned by tzdatastructs.TimeZoneDataDto{}.New()\n" +
			"Error: %v\n", err.Error())
		return
	}

	if tzDto.TzName != "America/Chicago" {
		t.Errorf("Error: Expected tzDto.TzName=='America/Chicago'.\n" +
			"Instead, TzName='%v'\n", tzDto.TzName)
	}

}

func TestTimeZoneGroupDto_01(t *testing.T) {
	/*
		parentGroupName,
		groupName,
		groupSortValue,
		typeName,
		typeValue,
		ianaVariableName,
		sourceFileNameExt string,
		groupType TimeZoneGroupType,
		groupClass TimeZoneGroupClass,
		deprecationStatus TimeZoneDeprecationStatus
	 */

	tzGrpDto, err := tzdatastructs.TimeZoneGroupDto{}.New(
		"",
		"America",
		"America",
		"",
		"",
		"",
		"northamerica",
		tzdatastructs.TzGrpType.Standard(),
		tzdatastructs.TzGrpClass.IANA(),
		tzdatastructs.DepStatusCode.Valid())

	if err != nil {
		t.Errorf("Error returned by tzdatastructs.TimeZoneGroupDto{}.New()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if tzGrpDto.GroupName != "America" {
		t.Errorf("Error: Expected GroupName='America'.\n" +
			"Instead, GroupName='%v'\n", tzGrpDto.GroupName)
	}
}