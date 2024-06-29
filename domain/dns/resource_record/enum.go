package resourcerecord

type Type byte

const (
	_      = iota
	A Type = iota
	NS
	MD
	MF
	CNAME
	SOA
	MB
	MG
	MR
	NULL
	WKS
	PTR
	HINFO
	MINFO
	MX
	TXT
)

func (t Type) String() string {
	types := []string{"", "A", "NS", "MD", "MF", "CNAME", "SOA", "MB", "MG", "MR", "NULL", "WKS", "PTR", "HINFO", "MINFO", "MX", "TXT", "INVALID"}
	index := t
	if index == 0 || index > TXT {
		index = TXT + 1
	}

	return types[index]
}

type Class byte

const (
	_        = iota
	IN Class = iota
	CS
	CH
	HS
	AnyClass Class = 255
)

func (c Class) String() string {
	classes := []string{"", "IN", "CS", "CH", "HS", "INVALID", "*"}
	index := c
	if index == 0 || (index > HS && index != AnyClass) {
		index = HS + 1
	}
	if index == AnyClass {
		index = Class(len(classes) - 1)
	}

	return classes[index]
}
