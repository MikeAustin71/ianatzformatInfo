package outprocess

import (
	"errors"
	"fmt"
	"github.com/MikeAustin71/stringopsgo/strops/v2"
	"strings"
)

type MarginFieldSpec struct {
	MarginStr           string
	MarginLength        int
	MarginChar          rune
}

func (margin *MarginFieldSpec) CopyOut() MarginFieldSpec {

	newMargin := MarginFieldSpec{}

	newMargin.MarginStr = margin.MarginStr
	newMargin.MarginChar = margin.MarginChar
	newMargin.MarginLength = margin.MarginLength

	return newMargin
}

type LineBreakSpec struct {
	CreateLineBreak       bool
	LeadingBlankLines     int
	LeftMargin            MarginFieldSpec
	LeftSpacer            MarginFieldSpec
	LineBreakChar         rune
	LineBreakLength       int
	RightMargin           MarginFieldSpec
	TerminateWithNewLine  bool
	TrailingBlankLines    int
}

func (lineBrk *LineBreakSpec) CopyOut() LineBreakSpec {

	newLineBreak := LineBreakSpec{}

	newLineBreak.CreateLineBreak = lineBrk.CreateLineBreak
	newLineBreak.LeftMargin = lineBrk.LeftMargin.CopyOut()
	newLineBreak.LeftSpacer = lineBrk.LeftSpacer.CopyOut()
	newLineBreak.LineBreakChar = lineBrk.LineBreakChar
	newLineBreak.LineBreakLength = lineBrk.LineBreakLength
	newLineBreak.RightMargin = lineBrk.RightMargin.CopyOut()
	newLineBreak.TerminateWithNewLine = lineBrk.TerminateWithNewLine

	return newLineBreak
}

type StringFieldSpec struct {
	LeftMargin            MarginFieldSpec
	LeftSpacer            MarginFieldSpec
	StrValue              string
	StrFieldLength        int
	StrPadChar            rune
	StrPosition           FieldPositionSpec
	RightSpacer           MarginFieldSpec
	TerminateWithNewLine  bool
}

func (strField *StringFieldSpec) CopyOut() StringFieldSpec {

	newStrField := StringFieldSpec{}

	newStrField.LeftMargin = strField.LeftMargin.CopyOut()
	newStrField.LeftSpacer = strField.LeftSpacer.CopyOut()
	newStrField.StrValue = strField.StrValue
	newStrField.StrFieldLength = strField.StrFieldLength
	newStrField.StrPadChar = strField.StrPadChar
	newStrField.StrPosition = strField.StrPosition
	newStrField.RightSpacer = strField.RightSpacer.CopyOut()
	newStrField.TerminateWithNewLine = strField.TerminateWithNewLine

	return newStrField
}

type NumericIntFieldSpec struct {
	LeftMargin            MarginFieldSpec
	LeftSpacer            MarginFieldSpec
	NumericValue          int
	NumericFieldSpec      string
	NumericFieldLength    int
	NumericPadChar        rune
	NumericPosition       FieldPositionSpec
	RightSpacer           MarginFieldSpec
	TerminateWithNewLine  bool
}

func (numIntField *NumericIntFieldSpec) CopyOut() NumericIntFieldSpec {

	newNumField := NumericIntFieldSpec{}

	newNumField.LeftMargin = numIntField.LeftMargin.CopyOut()
	newNumField.LeftSpacer = numIntField.LeftSpacer.CopyOut()

	newNumField.NumericValue = numIntField.NumericValue
	newNumField.NumericFieldSpec = numIntField.NumericFieldSpec
	newNumField.NumericFieldLength = numIntField.NumericFieldLength
	newNumField.NumericPadChar = numIntField.NumericPadChar
	newNumField.NumericPosition = numIntField.NumericPosition
	newNumField.RightSpacer = numIntField.RightSpacer.CopyOut()
	newNumField.TerminateWithNewLine = numIntField.TerminateWithNewLine

	return newNumField
}


type LabelStrValueLine struct {
	LeadingBlankLines       int
	TopLineBreak            LineBreakSpec
	LabelStr                StringFieldSpec
	ValueStr                StringFieldSpec
	BottomLineBreak         LineBreakSpec
	TrailingBlankLines      int
}

type LabelSingleIntValueLine struct {
	LeadingBlankLines  int
	TopLineBreak       LineBreakSpec
	LabelField         StringFieldSpec
	NumberField        NumericIntFieldSpec
	BottomLineBreak    LineBreakSpec
	TrailingBlankLines int
}

func (labelIntLine *LabelSingleIntValueLine) CopyOut() LabelSingleIntValueLine {

	newLabelInt := LabelSingleIntValueLine{}

	newLabelInt.LeadingBlankLines = labelIntLine.LeadingBlankLines
	newLabelInt.TopLineBreak = labelIntLine.TopLineBreak.CopyOut()
	newLabelInt.LabelField = labelIntLine.LabelField.CopyOut()
	newLabelInt.NumberField = labelIntLine.NumberField.CopyOut()
	newLabelInt.BottomLineBreak = labelIntLine.BottomLineBreak.CopyOut()
	newLabelInt.TrailingBlankLines = labelIntLine.TrailingBlankLines
	return newLabelInt
}

type TzStrBuilder struct {
	Input    string
	Output   string
}


func (tzStrBuilder TzStrBuilder) CreateLabelSingleIntLine(
	element LabelSingleIntValueLine,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzLogOps.CreateLabelSingleIntLine() "

	var outputStr string
	var err error
	strOps := strops.StrOps{}
	return nil
}

func (tzStrBuilder TzStrBuilder) CreateLineBreak(
	lineBreak LineBreakSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzStrBuilder.CreateLineBreak() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	if !lineBreak.CreateLineBreak {
		// Nothing to do
		return nil
	}

	var err error

	if lineBreak.LeadingBlankLines > 0 {
		err = tzStrBuilder.CreateNewLine(
			lineBreak.LeadingBlankLines,
			b,
			ePrefix)

		if err != nil {
			return err
		}
	}

	err = tzStrBuilder.CreateMarginField(
		lineBreak.LeftMargin,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = tzStrBuilder.CreateMarginField(
		lineBreak.LeftSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	if lineBreak.LineBreakLength > 0 &&
		lineBreak.LineBreakChar != 0 {
		for i:=0; i < lineBreak.LineBreakLength; i++ {
			_, err = b.WriteRune(lineBreak.LineBreakChar)

			if err != nil {
				return fmt.Errorf(ePrefix +
					"\nError returned by b.WriteRune(lineBreak.LineBreakChar)\n" +
					"lineBreak.LineBreakChar='%v'\n" +
					"Error='%v'\n", err.Error())
			}
		}
	}

	numOfNewLines := lineBreak.TrailingBlankLines

	if lineBreak.TerminateWithNewLine {
		numOfNewLines++
	}

	if numOfNewLines > 0 {
		err = tzStrBuilder.CreateNewLine(numOfNewLines, b, ePrefix)

		if err != nil {
			return err
		}
	}

	return nil
}

func (tzStrBuilder TzStrBuilder) CreateMarginField(
	margin MarginFieldSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzStrBuilder.CreateMarginField() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	if len(margin.MarginStr) > 0 {
		_, err = b.WriteString(margin.MarginStr)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteString(margin.MarginStr)\n" +
				"margin.MarginStr='%v'\n" +
				"Error='%v'\n", margin.MarginStr, err.Error() )
		}
	}

	if margin.MarginLength == 0 {
		// Nothing to do
		return nil
	}

	if margin.MarginChar == 0 {
		return errors.New(ePrefix +
			"\nError: margin.MarginChar = 0\n")
	}

	xb := strings.Builder{}
	xb.Grow(margin.MarginLength + 2)

	for i:=0; i < margin.MarginLength; i++ {
		_, err = xb.WriteRune(margin.MarginChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by xb.WriteRune(margin.MarginChar)\n" +
				"margin.MarginChar='%v'\n" +
				"Error='%v'\n", margin.MarginChar, err.Error())
		}
	}

	_, err = b.WriteString(xb.String())

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by b.WriteString(xb.String())\n" +
			"xb.String()='%v'\n" +
			"Error='%v'\n", xb.String(), err.Error())
	}

		return nil
}

func (tzStrBuilder TzStrBuilder) CreateNewLine(
	numOfNewLines int,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzStrBuilder.CreateNewLine() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	if numOfNewLines < 1 {
		// Nothing to do
		return nil
	}

	if numOfNewLines > 51 {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter 'numOfNewLines' > 51\n" +
			"numOfNewLines='%v'\n", numOfNewLines)
	}

	var err error

	for i:=0; i < numOfNewLines; i++ {

		_, err = b.WriteRune('\n')

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(newline)\n" +
				"i='%v'\n" +
				"Error='%v'", i, err.Error())
		}
	}

	return nil
}

// CreateStringField - Designed to handle StringFieldSpec specifications.
func (tzStrBuilder TzStrBuilder) CreateStringField(
	strSpec StringFieldSpec,
	b *strings.Builder,
	ePrefix string) error{

	ePrefix += "TzStrBuilder.CreateStringField() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	err = tzStrBuilder.CreateMarginField(
		strSpec.LeftMargin,
		b,
		ePrefix)

	if err != nil {
		return err
	}


	err = tzStrBuilder.CreateMarginField(
		strSpec.LeftSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	switch strSpec.StrPosition {

	case FieldPos.LeftJustify():

		err = tzStrBuilder.LeftJustifyField(
			strSpec.StrValue,
			strSpec.StrFieldLength,
			strSpec.StrPadChar,
			b,
			ePrefix)

	case FieldPos.RightJustify():

		err = tzStrBuilder.RightJustifyField(
			strSpec.StrValue,
			strSpec.StrFieldLength,
			strSpec.StrPadChar,
			b,
			ePrefix)

	case FieldPos.Center():

		err = tzStrBuilder.CenterInField(
			strSpec.StrValue,
			strSpec.StrFieldLength,
			strSpec.StrPadChar,
			b,
			ePrefix)

	default:
		err = fmt.Errorf(ePrefix +
			"\nError: strSpec.StrPosition is invalid!\n" +
			"StrPosition Value='%v'\n", strSpec.StrPosition.UtilityValue())
	}

	if err != nil {
		return err
	}

	err = tzStrBuilder.CreateMarginField(
		strSpec.RightSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	if strSpec.TerminateWithNewLine {
		err = tzStrBuilder.CreateNewLine(1,b, ePrefix)

		if err != nil {
			return err
		}
	}

	return nil
}

func (tzStrBuilder TzStrBuilder) CenterInField(
	strValue string,
	strFieldLen int,
	padChar rune,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzStrBuilder.CenterInField() "


	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	lenStr := len(strValue)

	if lenStr > strFieldLen {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter 'strValue' length exceeds field length parameter, 'strFieldLen'.\n" +
			"strValue='%v'\n" +
			"strValue length='%v'\n" +
			"strFieldLen='%v'\n", strValue, lenStr, strFieldLen)
	}

	var err error

	if lenStr == strFieldLen {

		_, err = b.WriteString(strValue)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteString(strValue)\n" +
				"strValue string length == strFieldLen\n" +
				"strValue='%v'\n" +
				"Error='%v'\n", strValue, err.Error())
		}
	}

	if padChar == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'padChar' is ZERO!\n")
	}

	grossPad := strFieldLen - lenStr
	var leftPad, rightPad int

	leftPad = grossPad / 2
	rightPad = grossPad - leftPad

	for i:= 0; i < leftPad; i++ {

		_, err = b.WriteRune(padChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(left padChar)\n" +
				"padChar='%v'   i='%v'\n" +
				"Error='%v'\n", padChar, i, err.Error())
		}
	}

	_, err = b.WriteString(strValue)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned after left pad by b.WriteString(strValue)\n" +
			"strValue='%v'\n" +
			"Error='%v'\n", strValue, err.Error())
	}

	for j:=0; j < rightPad; j++ {

		_, err = b.WriteRune(padChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(right padChar)\n" +
				"padChar='%v'   j='%v'\n" +
				"Error='%v'\n", padChar, j, err.Error())
		}
	}

	return nil
}



// LeftJustifyField - Left Justifies parameter 'str' in a field of length 'fieldLen'.
func (tzStrBuilder TzStrBuilder) LeftJustifyField(
	strValue string,
	strFieldLen int,
	trailingPadChar rune,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzStrFmt.LeftJustifyField() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	lenStr := len(strValue)

	if lenStr > strFieldLen {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter 'strValue' length exceeds field length parameter, 'strFieldLen'.\n" +
			"strValue='%v'\n" +
			"strValue length='%v'\n" +
			"strFieldLen='%v'\n", strValue, lenStr, strFieldLen)
	}

	var err error

	_, err = b.WriteString(strValue)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by b.WriteString(strValue)\n" +
			"strValue='%v'\n" +
			"Error='%v'\n", strValue, err.Error())
	}

	if lenStr == strFieldLen {
		return nil
	}

	if trailingPadChar == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'trailingPadChar' is Zero!")
	}

	padLen := strFieldLen - lenStr

	for i:=0; i < padLen; i++ {

		_, err = b.WriteRune(trailingPadChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(trailingPadChar)\n" +
				"trailingPadChar='%v'   i='%v'\n" +
				"Error='%v'\n", trailingPadChar, i, err.Error())
		}
	}

	return nil
}

func (tzStrBuilder TzStrBuilder) RightJustifyField(
	strValue string,
	strFieldLen int,
	leadingPadChar rune,
	b *strings.Builder,
	ePrefix string) error {


	ePrefix += "TzStrBuilder.RightJustifyField() "

	lenStr := len(strValue)

	if lenStr > strFieldLen {
		return fmt.Errorf(ePrefix +
			"\nError: Input Parameter 'strValue' length exceeds 'fieldLen'.\n" +
			"strValue='%v'" +
			"strValue length='%v'\n" +
			"fieldLen='%v'\n", strValue, lenStr, strFieldLen)
	}

	var err error

	if lenStr == strFieldLen {

		_, err = b.WriteString(strValue)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"Error returned by b.WriteString(strValue)\n" +
				"strValue='%v'\n" +
				"Error='%v'\n", strValue, err.Error())
		}

		return nil
	}

	if leadingPadChar == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'leadingPadChar' is ZERO!\n")
	}

	padLen := strFieldLen - lenStr

	for i:=0; i < padLen; i++ {

		_, err = b.WriteRune(leadingPadChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(leadingPadChar)\n" +
				"leadingPadChar='%v'   i='%v'\n" +
				"Error='%v'\n", leadingPadChar, i, err.Error())
		}
	}

	_, err = b.WriteString(strValue)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned after padChar by b.WriteString(strValue)\n" +
			"strValue='%v'\n" +
			"Error='%v'\n", strValue, err.Error())
	}

	return nil
}