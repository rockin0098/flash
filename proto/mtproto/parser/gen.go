package parser

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

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
	// case "bytes":
	// 	return "int32"
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

func convertSpecialDecodeField(in string) string {
	ele, ok := vectorType(in)
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

// 复合类型应该用 TLObject 处理
func convertDecodeField(param *TLParam) string {

	name := param.Name
	tp := param.Type
	fname := convertFieldName(name)

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
	default:
		fn = convertSpecialDecodeField(tp)
		// Log.Warnf("unknown param type = %v", tp)
	}

	s := fmt.Sprintf("t.%s=%s\n", fname, fn)

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
	// case "bytes":
	// 	fn = "ec.Int(%s)"
	default:
		fn = convertSpecialEncodeField(tp)
		// Log.Warnf("unknown param type = %v", tp)
	}

	s := fmt.Sprintf(fn, "t."+convertGetterName(name)+"()") + "\n"

	return s
}
