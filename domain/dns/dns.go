package dns

import (
	"fmt"
	"strings"

	"github.com/arnaldomf/dns/domain/dns/header"
	"github.com/arnaldomf/dns/domain/dns/question"
)

type DNS struct {
	Header   *header.Header
	Question *question.Question
}

func New(pkg []byte) (*DNS, error) {
	h, err := header.New(pkg)
	if err != nil {
		return nil, err
	}
	q, err := question.New(pkg, header.HeaderLength)
	if err != nil {
		return nil, err
	}
	return &DNS{
		Header:   h,
		Question: q,
	}, nil
}

func (d *DNS) String() string {
	var sb strings.Builder

	sb.WriteString("DNS={")
	sb.WriteString("\n\t")
	sb.WriteString(fmt.Sprintf("Header=%s,", d.Header))
	sb.WriteString("\n\t")
	sb.WriteString(fmt.Sprintf("Question=%s", d.Question))
	sb.WriteString("\n")
	sb.WriteRune('}')

	return sb.String()
}
