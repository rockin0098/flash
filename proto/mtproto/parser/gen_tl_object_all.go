package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	// . "github.com/rockin0098/meow/base/logger"
	. "github.com/rockin0098/meow/base/global"
)

var m_class_id = "M_classID"
var tlobject_output_file = "tl.object.all.go"
var tlobject_template = `
package mtproto

import "fmt"
import "encoding/hex"

%v

`

func (t *TLLayer) generateOneTLObjectFields(params []*TLParam) string {

	res := m_class_id + " int32\n"
	for _, p := range params {
		name := p.Name
		tp := convertDataType(p.Type)
		s := fmt.Sprintf("%s %s\n", convertFieldName(name), tp)
		res = res + s

		_, ok := t.RTypeMap[p.Type]
		if !ok {
			t.ExceptFiledTypeMap[p.Type] = p.Type
		}
	}

	return res
}

func (t *TLLayer) generateOneTLObjectSetterGetter(objName string, params []*TLParam) string {

	res := `
		func (t *%v)ClassID() int32 {
			return t.%v
		}
	`
	res = fmt.Sprintf(res, objName, m_class_id)
	for _, p := range params {
		name := p.Name
		fname := convertFieldName(name)
		tp := convertDataType(p.Type)
		s := fmt.Sprintf(`
		func (t *%v)%s(%s %s) {
			t.%s = %s
		}
		`, objName, convertSetterName(name), fname, tp, fname, fname)

		g := fmt.Sprintf(`
		func (t *%v)%s() %s {
			return t.%s
		}
		`, objName, convertGetterName(name), tp, fname)

		res = res + s + g
	}

	return res
}

func (t *TLLayer) generateOneTLObjectDecode(params []*TLParam) string {

	if len(params) == 0 {
		return "return nil"
	}

	res := "dc := NewMTPDecodeBuffer(b)\n\n"
	for _, p := range params {
		// if p.Type == "#" {
		// 	continue
		// }
		res = res + convertDecodeField(p)
	}

	res = res + "\nreturn dc.err"

	return res
}

func (t *TLLayer) generateOneTLObjectEncode(tlname string, params []*TLParam) string {

	res := "ec := NewMTPEncodeBuffer(512)\n\n"
	res = res + fmt.Sprintf("ec.Int(int32(TL_CLASS_%v))\n", tlname)

	if len(params) > 0 && params[0].Type == "#" {
		ASSERT(params[0].Name == "flags")
		res = res + convertEncodeFlagsField(params)
	} else {
		for _, p := range params {
			ASSERT(p.Type != "#")
			res = res + convertEncodeField(p)
		}
	}

	res = res + "\nreturn ec.GetBuffer()"

	return res
}

func (t *TLLayer) generateOneTLObjectString(objname string, line *TLLine) string {

	res := `
	func (t *%v)String() string {
	`
	res = fmt.Sprintf(res, objname)

	formatstr := fmt.Sprintf(`- %v#%v\n`, line.Predicate, line.ID)
	res = res + `return fmt.Sprintf("` + formatstr

	params := line.Params
	if len(params) > 0 {

		for _, p := range params {
			res = res + fmt.Sprintf("\t-- %v : ", p.Name)
			res = res + "%v\\n"
		}

		res = res + `",`

		for _, p := range params {
			cvType := convertDataType(p.Type)
			if cvType == "[]byte" {
				res = res + fmt.Sprintf("hex.EncodeToString(t.%v),", convertFieldName(p.Name))
			} else {
				res = res + fmt.Sprintf("t.%v,", convertFieldName(p.Name))
			}
		}

		// 去掉尾巴
		c := string(res[len(res)-1])
		if c == "," {
			res = res[:len(res)-1]
		}

		res = res + ")}"
	} else {
		res = res + `")}`
	}

	return res
}

func (t *TLLayer) generateOneTLObject(line *TLLine) string {

	tlname := strings.Replace(line.Predicate, ".", "_", -1)
	objname := convertTLObjectName(line.Predicate)

	defstr := fmt.Sprintf(`
		// %v#%v
		type %v struct {
			%v
		}
	`, line.Predicate, line.ID, objname, t.generateOneTLObjectFields(line.Params))

	sgstr := t.generateOneTLObjectSetterGetter(objname, line.Params)

	// newfuncname := fmt.Sprintf("New_%v", objname)
	newfuncname := convertNewTLFuncName(line.Predicate)
	newfuncstr := fmt.Sprintf(`
		func %v() *%v {
			return &%v{
				%v: %v,
			}
		}
	`, newfuncname, objname, objname, m_class_id, "TL_CLASS_"+tlname)

	encodestr := fmt.Sprintf(`
		func (t *%v)Encode() []byte {
			%s
		}
	`, objname, t.generateOneTLObjectEncode(tlname, line.Params))

	decodestr := fmt.Sprintf(`
		func (t *%v)Decode(b []byte) error {
			%s
		}
	`, objname, t.generateOneTLObjectDecode(line.Params))

	stringstr := t.generateOneTLObjectString(objname, line)

	res := fmt.Sprintf(`
		%v
		%v
		%v
		%v
		%v
		%v
	`, defstr, sgstr, newfuncstr, encodestr, decodestr, stringstr)

	return res
}

func (t *TLLayer) GenerateTLObjectAll() {
	lines := t.Lines

	ss := ""
	for _, line := range lines {
		// Log.Infof("predicate = %v", line.Predicate)
		s := t.generateOneTLObject(line)
		ss = ss + s
	}

	filecontent := fmt.Sprintf(tlobject_template, ss)
	file := t.OutputDir + tlobject_output_file

	ioutil.WriteFile(file, []byte(filecontent), 0644)
}
