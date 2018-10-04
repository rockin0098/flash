package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/rockin0098/flash/base/logger"
)

var tlobject_output_file = "tl.object.all.go"
var tlobject_template = `
package tl

%v

`

func (t *TLLayer) generateOneTLObjectFields(params []*TLParam) string {

	res := ""
	for _, p := range params {
		name := p.Name
		tp := convertDataType(p.Type)
		s := fmt.Sprintf("_%s %s\n", name, tp)
		res = res + s
	}

	return res
}

func (t *TLLayer) generateOneTLObject(line *TLLine) string {

	seperator := fmt.Sprintf("//====%v#%v====", line.Predicate, line.ID)

	objname := strings.Replace(line.Predicate, ".", "_", -1)
	objname = "TL_" + objname

	defstr := fmt.Sprintf(`
		type %v struct {
			%v
		}
	`, objname, t.generateOneTLObjectFields(line.Params))

	newfuncname := fmt.Sprintf("New_%v", objname)
	newfuncstr := fmt.Sprintf(`
		func %v() *%v {
			return &%v{}
		}
	`, newfuncname, objname, objname)

	encodestr := fmt.Sprintf(`
		func (t *%v)Encode() []byte {
			return nil
		}
	`, objname)

	decodestr := fmt.Sprintf(`
		func (t *%v)Decode(b []byte) {}
	`, objname)

	res := fmt.Sprintf(`
		%v
		%v
		%v
		%v
		%v
	`, seperator, defstr, newfuncstr, encodestr, decodestr)

	return res
}

func (t *TLLayer) GenerateTLObjectAll() {
	lines := t.Lines

	ss := ""
	for _, line := range lines {
		Log.Infof("predicate = %v", line.Predicate)
		s := t.generateOneTLObject(line)
		ss = ss + s
	}

	filecontent := fmt.Sprintf(tlobject_template, ss)
	file := t.OutputDir + tlobject_output_file

	ioutil.WriteFile(file, []byte(filecontent), 0644)
}
