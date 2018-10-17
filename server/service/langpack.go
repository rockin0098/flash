package service

type LangPack struct {
	LangCode          string
	Version           int32
	Strings           []*LangPackString
	StringPluralizeds []*LangPackString
	StringDeleteds    []*LangPackString
}

type LangPackString struct {
	Key        string
	Value      string
	ZeroValue  string
	OneValue   string
	TwoValue   string
	FewValue   string
	ManyValue  string
	OtherValue string
}
