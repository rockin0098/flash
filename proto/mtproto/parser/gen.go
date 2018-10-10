package parser

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	. "github.com/rockin0098/flash/base/global"
	. "github.com/rockin0098/flash/base/logger"
)

const (
	tl_type_string = "string"
	tl_type_long   = "long"
	tl_type_int    = "int"
	tl_type_double = "double"
	tl_type_int128 = "int128"
	tl_type_int256 = "int256"
)

func convertFieldName(field string) string {
	return "M_" + field
}

func convertCRC32(crc32 string) uint32 {

	if len(crc32) < 8 {
		crc32 = fmt.Sprintf("%08s", crc32)
	}

	idbytes, err := hex.DecodeString(crc32)
	if err != nil {
		Log.Error(err)
		return 0
	}

	crc32int := binary.BigEndian.Uint32(idbytes)

	return crc32int
}

func convertTLObjectName(in string) string {
	tlname := strings.Replace(in, ".", "_", -1)
	objname := "TL_" + tlname

	return objname
}

func convertNewTLFuncName(in string) string {
	tlobj := convertTLObjectName(in)
	return "New_" + tlobj
}

func convertSpecialTLDataType(in string) string {

	if strings.Contains(in, "flags") {
		_, tp2 := parseFlagField(in)
		return convertDataType(tp2)
	}

	ele, ok := vectorType(in)
	if ok {
		nele := convertDataType(ele)
		return "[]" + nele
	} else {
		return "TLObject"
	}
}

func convertDataType(in string) string {
	switch in {
	case "string":
		return "string"
	case "long":
		return "int64"
	case "int":
		return "int32"
	case "double":
		return "float64"
	case "#": // must be "flags"
		return "uint32"
	case "int128",
		"int256":
		return "[]byte"
	default: // including "!X"
		return convertSpecialTLDataType(in)
	}
}

func convertGetterName(name string) string {
	return fmt.Sprintf("Get_%s", name)
}

func convertSetterName(name string) string {
	return fmt.Sprintf("Set_%s", name)
}

func vectorType(vtp string) (string, bool) {
	lowin := strings.ToLower(vtp)

	// vector 类型
	if strings.Contains(lowin, "vector<") {
		si := strings.Index(vtp, "<")
		ei := strings.Index(vtp, ">")
		ele := string(vtp[si+1 : ei])
		// Log.Info("ele : ", ele)

		return ele, true
	}

	return "", false
}

func parseFlagField(tp string) (int, string) {

	ASSERT(strings.Contains(tp, "flags"))

	st := strings.Split(tp, "?")
	ASSERT(len(st) == 2)
	sf := st[0]
	tp2 := st[1]
	sp := strings.Split(sf, ".")
	ASSERT(len(sp) == 2)
	ASSERT(sp[0] == "flags")
	pos, err := strconv.Atoi(sp[1])
	ASSERT(err == nil)

	return pos, tp2
}

func convertSpecialDecodeField(fname string, tp string) string {

	// 处理含 flags 的字段
	if strings.Contains(tp, "flags") {
		pos, tp2 := parseFlagField(tp)
		flagField := `if (t.M_flags &(1 << %v)) != 0 {
			%s
		}`

		return fmt.Sprintf(flagField, pos, convertDecodeType(fname, tp2))

	} else {
		ele, ok := vectorType(tp)
		if ok {
			if ele == tl_type_int {
				return "dc.VectorInt()"
			} else if ele == tl_type_long {
				return "dc.VectorLong()"
			} else if ele == tl_type_string {
				return "dc.VectorString()"
			}
			return "dc.Vector()"
		} else {
			return "dc.TLObject()"
		}
	}
}

func convertDecodeType(fname string, tp string) string {

	fn := ""

	switch tp {
	case "string":
		fn = "dc.String()"
	case "long":
		fn = "dc.Long()"
	case "int":
		fn = "dc.Int()"
	case "double":
		fn = "dc.Double()"
	case "int128":
		fn = "dc.Bytes(16)"
	case "int256":
		fn = "dc.Bytes(32)"
	case "#": // must be "flags"
		fn = "dc.UInt()"
	default:
		fn = convertSpecialDecodeField(fname, tp)
		// Log.Warnf("unknown param type = %v", tp)
	}

	s := fn
	if !strings.Contains(fn, "flags") {
		s = fmt.Sprintf("t.%s=%s", fname, fn)
	}

	return s
}

// 复合类型应该用 TLObject 处理
func convertDecodeField(param *TLParam) string {

	name := param.Name
	tp := param.Type
	fname := convertFieldName(name)

	s := convertDecodeType(fname, tp)

	s = s + "\n"

	return s
}

func convertSpecialEncodeField(in string) string {

	ele, ok := vectorType(in)
	if ok {
		if ele == tl_type_int {
			return "ec.VectorInt(%s)"
		} else if ele == tl_type_long {
			return "ec.VectorLong(%s)"
		} else if ele == tl_type_string {
			return "ec.VectorString(%s)"
		}
		return "ec.Vector(%s)"
	} else {
		return "ec.TLObject(%s)"
	}
}

// 复合类型应该用 TLObject 处理
func convertEncodeField(param *TLParam) string {

	name := param.Name
	tp := param.Type
	// fname := convertFieldName(name)

	fn := ""

	switch tp {
	case "string":
		fn = "ec.String(%s)"
	case "long":
		fn = "ec.Long(%s)"
	case "int":
		fn = "ec.Int(%s)"
	case "double":
		fn = "ec.Double(%s)"
	case "int128":
		fn = "ec.Bytes(%s)"
	case "int256":
		fn = "ec.Bytes(%s)"
	default:
		fn = convertSpecialEncodeField(tp)
		// Log.Warnf("unknown param type = %v", tp)
	}

	s := fmt.Sprintf(fn, "t."+convertGetterName(name)+"()") + "\n"

	return s
}

func convertEncodeFlagsField(params []*TLParam) string {

	ASSERT(len(params) > 0)
	p0 := params[0]
	ASSERT(p0.Name == "flags" && p0.Type == "#")

	var fps []*TLParam
	for _, p := range params {
		if strings.Contains(p.Type, "flags") {
			fps = append(fps, p)
		}
	}

	getRightValue := func(tp2 string) string {
		right := "nil"
		switch tp2 {
		case "int", "long", "double":
			right = "0"
		case "string":
			right = "\"\""
		default:
			right = "nil"
		}

		return right
	}

	res := `
		var flags uint32 = 0
	`
	for _, fp := range fps {

		typ := fp.Type
		pos, tp2 := parseFlagField(typ)

		right := getRightValue(tp2)
		ft := `if t.%v != %v {
			flags |= 1 << %v
		}

		`

		res = res + fmt.Sprintf(ft, convertFieldName(fp.Name), right, pos)
	}

	res = res + "ec.UInt(flags)\n\n"

	p2 := params[1:]
	for _, pp := range p2 {
		if strings.Contains(pp.Type, "flags") {
			_, tp2 := parseFlagField(pp.Type)
			right := getRightValue(tp2)
			tpp := &TLParam{
				Name: pp.Name,
				Type: tp2,
			}

			res = res + fmt.Sprintf(`if t.%v != %v {
				%s}
			`, convertFieldName(pp.Name), right, convertEncodeField(tpp))

		} else {
			res = res + convertEncodeField(pp)
		}
	}

	return res
}
