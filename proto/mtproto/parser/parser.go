package parser

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	. "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

type EnumLineType int

const (
	_ EnumLineType = iota
	LineConstructors
	LineFunctions
	LineDeclaration
)

func calculateCRC32(name, params, restype string) uint32 {
	cleanline := fmt.Sprintf("%s%s= %s", name, params, restype)
	cleanline = regexp.MustCompile(" [a-zA-Z0-9_]+\\:flags\\.[0-9]+\\?true").ReplaceAllString(cleanline, "")
	cleanline = strings.Replace(cleanline, "<", " ", -1)
	cleanline = strings.Replace(cleanline, ">", " ", -1)
	cleanline = strings.Replace(cleanline, "  ", " ", -1)
	cleanline = regexp.MustCompile("^ ").ReplaceAllString(cleanline, "")
	cleanline = regexp.MustCompile(" $").ReplaceAllString(cleanline, "")
	cleanline = strings.Replace(cleanline, ":bytes ", ":string ", -1)
	cleanline = strings.Replace(cleanline, "?bytes ", "?string ", -1)
	cleanline = strings.Replace(cleanline, "{", "", -1)
	cleanline = strings.Replace(cleanline, "}", "", -1)

	// 通过cleanline计算出typeid并进行验证
	return crc32.ChecksumIEEE([]byte(cleanline))
}

func trimLine(line string) string {
	return strings.TrimSpace(line)
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
		return LineConstructors
	} else {
		return LineDeclaration
	}
}

func parseTLParam(param string) *TLParam {

	if len(param) == 0 ||
		len(param) == 1 ||
		param == "#" ||
		!strings.Contains(param, ":") {
		return nil
	}

	if string(param[0]) == "{" {
		return nil
	}

	name := ""
	dtype := ""

	ps := strings.Split(param, ":")
	if len(ps) == 2 {
		name = ps[0]
		dtype = ps[1]
	} else {
		Log.Error("param = ", param)
		return nil
	}

	return &TLParam{
		Name: name,
		Type: dtype,
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

	rtype := trimLine(tail)
	rtype = strings.Replace(rtype, ";", "", -1)

	var params []*TLParam
	ms := strings.Split(mid, " ")
	for _, m := range ms {
		m = trimLine(m)
		if len(m) > 0 {
			para := parseTLParam(m)
			if para != nil {
				params = append(params, para)
			}
		}
	}

	// 手工重新计算crc32
	crc32 := calculateCRC32(name, mid, rtype)

	return &TLLine{
		Predicate: name,
		ID:        uid,
		Params:    params,
		Type:      rtype,
		CRC32:     crc32,
	}
}

func ParseTLLayer(tl string) *TLLayer {

	currLineType := LineConstructors
	var tlConstructors []*TLConstructor
	var tlMethods []*TLMethod
	var tllines []*TLLine
	var layer string
	constructorMap := make(map[string]*TLConstructor)
	rtypeMap := make(map[string]string)

	lines := strings.Split(tl, "\n")
	for _, line := range lines {

		layerAnnotation := "// LAYER"
		if strings.Contains(line, layerAnnotation) {
			idx := strings.Index(line, layerAnnotation)
			layer = strings.TrimSpace(string(line[idx+len(layerAnnotation):]))
		}

		trimline := trimAnnotation(line)
		if len(trimline) > 0 { // 含有内容
			// Log.Info(trimline)

			lineType := checkLineType(line)
			if lineType == LineConstructors {
				currLineType = LineConstructors
			} else if lineType == LineFunctions {
				currLineType = LineFunctions
			} else {
				tlline := parseTLLine(trimline)
				if len(tlline.ID) > 0 { // 校验crc
					iduint32 := convertCRC32(tlline.ID)
					if tlline.CRC32 != iduint32 {
						Log.Errorf("crc32 not match, ID = %v, recalculated crc32 = %v",
							iduint32, tlline.CRC32)
						continue
					}
				} else {
					tlline.ID = fmt.Sprintf("%x", tlline.CRC32)
				}

				tllines = append(tllines, tlline)
				if currLineType == LineConstructors {
					constructor := (*TLConstructor)(tlline)
					tlConstructors = append(tlConstructors, constructor)
					constructorMap[constructor.Predicate] = constructor
					rtypeMap[constructor.Type] = constructor.Predicate
				} else if currLineType == LineFunctions {
					tlMethod := &TLMethod{
						ID:     tlline.ID,
						Method: tlline.Predicate,
						Params: tlline.Params,
						Type:   tlline.Type,
					}
					tlMethods = append(tlMethods, tlMethod)
				}
			}
		}
	}

	schema := &TLSchema{
		Constructors: tlConstructors,
		Methods:      tlMethods,
	}

	return &TLLayer{
		Layer:              layer,
		ConstructorMap:     constructorMap,
		RTypeMap:           rtypeMap,
		ExceptFiledTypeMap: map[string]string{},
		Schema:             schema,
		Lines:              tllines,
	}
}

func ParserEntry() {
	UseMaxCpu()

	Log.Infof("###################################")
	Log.Infof("############ TL Parser ############")
	Log.Infof("###################################")

	input := ""
	output := ""

	if len(os.Args) != 2 {
		Log.Info("\nUseage : tlparser [input file] [output dir]\nexample: tlparser ./proto/mtproto/schema/schema.layer73.tl ./proto/mtproto/\n")
		input = "./proto/mtproto/schema/schema.layer73.tl"
		output = "./proto/mtproto/"
	} else {
		input = os.Args[0]
		output = os.Args[1]
	}

	content, err := ioutil.ReadFile(input)
	if err != nil {
		Log.Warnf("read input file failed, err = %v", err)
		return
	}

	// layer := ParseTLLayer(TLContent)
	layer := ParseTLLayer(string(content))
	layer.OutputDir = output
	// Log.Infof("layer = %v", FormatStruct(layer))

	layer.GenerateTLObjectClassConst()
	layer.GenerateTLObjectAll()

	// for k, _ := range layer.ExceptFiledTypeMap {
	// 	Log.Debugf("except type : %v", k)
	// }
}
