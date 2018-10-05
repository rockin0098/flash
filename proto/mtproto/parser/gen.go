package parser

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	. "github.com/rockin0098/flash/base/logger"
)

func convertFieldName(field string) string {
	return "_" + field
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
	case "bytes":
		return "int32"
	case "int128",
		"int256":
		return "[]byte"
	default: // including "!X"
		return "TLObject"
	}
}

func convertGetterName(name string) string {
	return fmt.Sprintf("Get_%s", name)
}

func convertSetterName(name string) string {
	return fmt.Sprintf("Set_%s", name)
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
	case "bytes":
		fn = "dc.Int()"
	default:
		fn = "dc.TLObject()"
		// Log.Warnf("unknown param type = %v", tp)
	}

	s := fmt.Sprintf("t.%s=%s\n", fname, fn)

	return s
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
	case "bytes":
		fn = "ec.Int(%s)"
	default:
		fn = "ec.TLObject(%s)"
		// Log.Warnf("unknown param type = %v", tp)
	}

	s := fmt.Sprintf(fn, "t."+convertGetterName(name)+"()") + "\n"

	return s
}
