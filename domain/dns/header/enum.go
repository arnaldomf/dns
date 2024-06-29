package header

type QR byte

const (
	Query QR = iota
	Response
)

func (q QR) String() string {
	qnames := []string{"query", "response", "invalid"}
	index := q

	if q > Response {
		index = Response + 1
	}

	return qnames[index]
}

type OPCode byte

const (
	StandardQuery OPCode = iota
	InverseQuery
	ServerStatus
	OPReserved
)

func (op OPCode) String() string {
	opnames := []string{"QUERY", "IQUERY", "STATUS", "RESERVED"}
	index := op

	if op > OPReserved {
		index = OPReserved
	}

	return opnames[index]
}

type RCode byte

const (
	NoError RCode = iota
	FormatError
	ServerFailure
	NameError
	NotImplemented
	Refused
	RCReserved
)

func (rc RCode) String() string {
	rcodes := []string{"NoError", "FormatError", "ServerFailure", "NameError", "NotImplemented", "Refused", "Reserved"}
	index := rc

	if rc > RCReserved {
		index = RCReserved
	}

	return rcodes[index]
}
