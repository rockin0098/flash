package parser

func convertDataType(in string) string {
	switch in {
	case "string":
		return "string"
	default:
		return "[]byte"
	}
}
