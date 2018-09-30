package parser

import (
	"strings"

	. "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

type EnumLineType int

const (
	_ EnumLineType = iota
	LineTypes
	LineFunctions
	LineDeclaration
)

func trimLine(line string) string {
	trimchars := " \t\n\r"
	return strings.Trim(line, trimchars)
}

func trimAnnotation(line string) string {

	idx := strings.Index(line, "//")
	if idx == -1 { // 无注释,返回trim过后的行
		return trimLine(line)
	}

	// 截取注释符号之前的字符串
	sub := string(line[:idx])
	sub = trimLine(sub)

	return sub
}

func checkLineType(line string) EnumLineType {
	if strings.Contains(line, "---functions---") {
		return LineFunctions
	} else if strings.Contains(line, "---types---") {
		return LineTypes
	} else {
		return LineDeclaration
	}
}

func parseField(field string) *TLField {
	return &TLField{
		Name: field,
	}
}

func parseTLLine(line string) *TLLine {

	fi := strings.Index(line, " ")

	first := string(line[:fi])

	name := ""
	uid := ""
	if !strings.Contains(first, "#") {
		name = first
	} else {
		fs := strings.Split(first, "#")
		if len(fs) != 2 {
			Log.Errorf("invalid first head, fs = %v", fs)
			return nil
		}
		name = fs[0]
		uid = fs[1]
	}

	sub := line[fi:]
	si := strings.Index(sub, "=")

	mid := sub[:si]
	tail := sub[si+1:]

	value := trimLine(tail)
	value = strings.Replace(value, ";", "", -1)

	var fields []*TLField
	ms := strings.Split(mid, " ")
	for _, m := range ms {
		m = trimLine(m)
		if len(m) > 0 {
			fd := parseField(m)
			fields = append(fields, fd)
		}
	}

	return &TLLine{
		Name:   name,
		UID:    uid,
		Fields: fields,
		Value:  value,
	}
}

func ParseTL(tl string) *TLSchema {

	currLineType := LineTypes
	var tlTypes []*TLLine
	var tlFunctions []*TLLine

	lines := strings.Split(tl, "\n")
	for _, line := range lines {
		trimline := trimAnnotation(line)
		if len(trimline) > 0 { // 含有内容
			// Log.Info(trimline)
			lineType := checkLineType(line)
			if lineType == LineTypes {
				currLineType = LineTypes
			} else if lineType == LineFunctions {
				currLineType = LineFunctions
			} else {
				tlline := parseTLLine(trimline)
				if currLineType == LineTypes {
					tlTypes = append(tlTypes, tlline)
				} else if currLineType == LineFunctions {
					tlFunctions = append(tlFunctions, tlline)
				}
			}
		}
	}

	return &TLSchema{
		TLTypes:     tlTypes,
		TLFunctions: tlFunctions,
	}
}

func ParserEntry() {
	UseMaxCpu()

	Log.Infof("###################################")
	Log.Infof("############ TL Parser ############")
	Log.Infof("###################################")

	schema := ParseTL(TLContent)
	Log.Infof("schema = %v", FormatStruct(schema))
}
