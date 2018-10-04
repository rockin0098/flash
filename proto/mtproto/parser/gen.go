package parser

import (
	"fmt"
	// . "github.com/rockin0098/flash/base/logger"
)

func convertFieldName(field string) string {
	return "_" + field
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
	case "Bool":
		return "bool"
	default:
		return "[]byte"
	}
}

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
	case "Bool":
		fn = "dc.Bool()"
	default:
		fn = "dc.Bytes(16)"
		// Log.Warnf("unknown param type = %v", tp)
	}

	s := fmt.Sprintf("t.%s=%s\n", fname, fn)

	return s
}

func convertGetterName(name string) string {
	return fmt.Sprintf("Get_%s", name)
}

func convertSetterName(name string) string {
	return fmt.Sprintf("Set_%s", name)
}

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
	case "Bool":
		fn = "ec.Bool(%s)"
	// case "ipPort": //没有crc值
	// fn = "ec.Bytes(%s)"
	default:
		fn = "ec.Bytes(%s)"
		// Log.Warnf("unknown param type = %v", tp)
	}

	s := fmt.Sprintf(fn, "t."+convertGetterName(name)+"()") + "\n"

	return s
}
