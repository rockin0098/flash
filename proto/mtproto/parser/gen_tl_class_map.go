package parser

import (
	"fmt"
	"io/ioutil"
	// . "github.com/rockin0098/flash/base/logger"
)

var classmap_output_file = "tl.class.map.go"
var classmap_template = `
package mtproto

type newTLObjectFunc func() TLObject

var tlObjectClassMap = map[int32]newTLObjectFunc{
%v
}

`

func (t *TLLayer) GenerateTLObjectClassMap() {
	lines := t.Lines
	classItems := ""
	for _, line := range lines {

		// 没有id的暂时不解析
		if len(line.ID) == 0 {
			continue
		}

		classItem := fmt.Sprintf("int32(TL_CLASS_%v): func() TLObject { return %v() },\n",
			line.Predicate, convertNewTLFuncName(line.Predicate))
		classItems = classItems + classItem
	}

	filecontent := fmt.Sprintf(classmap_template, classItems)
	file := t.OutputDir + classmap_output_file

	ioutil.WriteFile(file, []byte(filecontent), 0644)
}
