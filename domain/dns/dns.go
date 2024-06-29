package dns

import (
	"fmt"
	"strings"

	"github.com/arnaldomf/dns/domain/dns/header"
)

type DNS struct {
	Header *header.Header
}

func New(pkg []byte) (*DNS, error) {
	h, err := header.New(pkg)
	if err != nil {
		return nil, err
	}
	return &DNS{Header: h}, nil
}

func (d *DNS) String() string {
	var sb strings.Builder

	sb.WriteString("DNS={")
	sb.WriteString(fmt.Sprintf("Header=%s", d.Header))
	sb.WriteRune('}')

	return sb.String()
}
